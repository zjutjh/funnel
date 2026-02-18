package sessionpool

import (
	"context"
	"errors"
	"funnel/comm"
	zfClient "funnel/httpclient/zf"
	"sync"
	"time"

	"github.com/zjutjh/mygo/nlog"
)

type CaptchaCookie struct {
	Cookie    *zfClient.ZFCookie
	CreatedAt time.Time
}

// SessionPool 用于缓存一批通过验证吗的Cookie, 便于登录的时候直接获取
type SessionPool struct {
	cookies chan *CaptchaCookie // 维护Cookie的channel
	maxSize int                 // channel容量
	ttl     time.Duration       // cookie过期时间
	workers int                 // 填充pool的并发数
	ctx     context.Context
}

var (
	once sync.Once
	pool *SessionPool
)

func New() *SessionPool {
	once.Do(func() {
		conf := comm.BizConf.ZF.SessionPool
		workers := conf.FillWorkers
		if workers <= 0 {
			workers = 3
		}
		maxSize := conf.MaxSize
		if maxSize <= 0 {
			maxSize = 64
		}
		ttlMinute := conf.TTLMinute
		if ttlMinute <= 0 {
			ttlMinute = 30
		}
		pool = &SessionPool{
			cookies: make(chan *CaptchaCookie, maxSize),
			maxSize: maxSize,
			ttl:     time.Minute * time.Duration(ttlMinute),
			workers: workers,
			ctx:     context.Background(),
		}
	})
	return pool
}

func (p *SessionPool) WithContext(ctx context.Context) *SessionPool {
	p.ctx = ctx
	return p
}

func Pick() *SessionPool {
	return New()
}

// Get 获取一个有效 cookie, 阻塞直到超时
func (p *SessionPool) Get(ctx context.Context) (*zfClient.ZFCookie, error) {
	timeout := time.After(5 * time.Second)
	for {
		select {
		case c := <-p.cookies:
			if time.Since(c.CreatedAt) > p.ttl {
				continue
			}
			nlog.Pick().WithContext(ctx).Info("命中 session pool")
			return c.Cookie, nil
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-timeout:
			return nil, errors.New("session pool: 获取超时")
		}
	}
}

// TryGet 非阻塞获取一个有效 cookie, 没有可用的立马返回错误
func (p *SessionPool) TryGet(ctx context.Context) (*zfClient.ZFCookie, error) {
	for {
		select {
		case c := <-p.cookies:
			if time.Since(c.CreatedAt) > p.ttl {
				continue
			}
			nlog.Pick().WithContext(ctx).Info("命中 session pool")
			return c.Cookie, nil
		default:
			nlog.Pick().WithContext(ctx).Errorf("未命中 session pool, 当前可用 %d/%d", len(p.cookies), p.maxSize)
			return nil, errors.New("session pool: 无可用 cookie")
		}
	}
}

// Run 启动 session pool 后台填充循环
func (p *SessionPool) Run() {
	// 预热
	logger := nlog.Pick()
	logger.Info("session pool: 预热填充中")
	p.fill(p.ctx)
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			p.fill(p.ctx)
		case <-p.ctx.Done():
			nlog.Pick().WithContext(p.ctx).Info("session pool: 停止后台填充")
			return
		}
	}
}

// fill 清理过期 cookie 并补充到 maxSize
func (p *SessionPool) fill(ctx context.Context) {
	logger := nlog.Pick()
	var valid []*CaptchaCookie
	// 排空
draining:
	for {
		select {
		case c := <-p.cookies:
			if time.Since(c.CreatedAt) <= p.ttl {
				valid = append(valid, c)
			}
		default:
			break draining
		}
	}
	// 回填
	for _, c := range valid {
		select {
		case p.cookies <- c:
		default:
		}
	}

	// 补充
	deficit := p.maxSize - len(p.cookies) + p.workers
	if deficit <= 0 {
		return
	}
	logger.Infof("session pool: 开始填充 (当前 %d/%d)", len(p.cookies), p.maxSize)
	// Semaphore 并发控制
	sem := make(chan struct{}, p.workers)
	var wg sync.WaitGroup

	for range deficit {
		sem <- struct{}{}
		wg.Add(1)
		go func() {
			defer wg.Done()
			defer func() { <-sem }()
			p.fillOne(ctx)
		}()
	}
	wg.Wait()
	logger.Infof("session pool: 填充完成 (当前 %d/%d)", len(p.cookies), p.maxSize)
}

// fillOne 填充单个 cookie, 并采取指数退避的重试策略
func (p *SessionPool) fillOne(ctx context.Context) {
	logger := nlog.Pick().WithContext(ctx)
	backoff := time.Second
	const maxBackoff = 30 * time.Second // 最大退避时间
	const maxRetries = 5                // 最大重试次数
	for range maxRetries {
		if ctx.Err() != nil {
			return
		}

		cookie, err := zfClient.New(ctx).BypassCaptcha()
		if err != nil {
			// 指数退避重试
			logger.Warnf("session pool: BypassCaptcha 失败: %v, %v 后重试", err, backoff)
			select {
			case <-time.After(backoff):
			case <-ctx.Done():
				return
			}
			backoff *= 2
			if backoff > maxBackoff {
				backoff = maxBackoff
			}
			continue
		}

		select {
		case p.cookies <- &CaptchaCookie{Cookie: cookie, CreatedAt: time.Now()}:
			logger.Debug("session pool: 成功填充一个 cookie")
		default:
			logger.Debug("session pool: channel 已满，丢弃")
		}
		return
	}
}

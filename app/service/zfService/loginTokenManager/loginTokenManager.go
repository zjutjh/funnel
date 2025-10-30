package loginTokenManager

import (
	"context"
	"fmt"
	"log"
	"sync"
	"sync/atomic"
	"time"
)

type LoginToken struct {
	RequestTime    time.Time
	JSESSIONID     string
	Route          string
	CSRFToken      string
	CryptoModulus  string
	CryptoExponent string
}

type LoginTokenManager struct {
	// CONFIGS
	preferredTokenCount     int32  // 期望保持的 Token 数量
	currentRunningResisters int32  // 当前正在运行的注册器数量
	maxResisters            int32  // 最大注册器数量
	tokenExpirationTimeSec  int32  // Token 过期时间
	routineMaintenanceSec   int32  // 两次例行池维护的间隔
	activeMaintenanceEvery  int32  // n 次 get 请求触发一次池维护
	hostUrl                 string // 注册器目标主机地址
	// VARS
	tokenPool         *loginTokenPool // Token 池实例
	getAfterLastActMt int32           // 上次主动维护后进行的请求数
}

func (m *LoginTokenManager) Init(host string, ctx context.Context, wg *sync.WaitGroup) {
	m.tokenPool = &loginTokenPool{}
	m.hostUrl = host
	m.preferredTokenCount = 100
	m.maxResisters = 20
	m.currentRunningResisters = 0
	m.tokenExpirationTimeSec = 2 * 60 * 60
	m.routineMaintenanceSec = 10
	m.activeMaintenanceEvery = 5
	m.getAfterLastActMt = 0

	go m.runPoolRoutineMt(ctx, wg)
}

func (m *LoginTokenManager) GetAToken() (*LoginToken, error) {
	log.Println("Pool Remain:", m.tokenPool.Size())
	tk, err := m.getTokenFromPool()
	if tk != nil && err == nil { //从池取 token 成功, 直接返回
		return tk, nil
	}
	ptk, err := RunRegister(m.hostUrl) // 开始尝试即时获取 token
	if err == nil {
		return &ptk, nil
	}
	return nil, fmt.Errorf("cannot get a token anyway")
}

// 从池中取 token, 一旦成功, 增加计数器的值
func (m *LoginTokenManager) getTokenFromPool() (*LoginToken, error) {
	if m.tokenPool.Size() < 1 {
		return nil, fmt.Errorf("there is no token remain")
	}
	atomic.AddInt32(&m.getAfterLastActMt, 1)
	if m.getAfterLastActMt >= m.activeMaintenanceEvery {
		atomic.StoreInt32(&m.getAfterLastActMt, 0)
		log.Println("Active Mt")
		m.poolMaintenance()
	}
	return m.tokenPool.Get(), nil
}

// 启动定时池维护
func (m *LoginTokenManager) runPoolRoutineMt(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	ticker := time.NewTicker(time.Duration(m.routineMaintenanceSec) * time.Second) // 定时触发
	defer ticker.Stop()                                                            // 确保 ticker 在函数退出时被停止

	for {
		select {
		case <-ctx.Done(): // 监听退出信号
			return
		case <-ticker.C:
			log.Println("Routine Mt")
			m.poolMaintenance()
		}
	}
}

// 运行池维护
func (m *LoginTokenManager) poolMaintenance() {

	log.Printf("Pool Mt %d\n", m.tokenPool.Size())

	m.tokenPool.RemoveExpired(time.Now().Add(-time.Duration(m.tokenExpirationTimeSec) * time.Second)) // 移除过期 token
	currentPoolSize := m.tokenPool.Size()                                                             // 获取当前池中 Token 数量
	if currentPoolSize < m.preferredTokenCount {                                                      // 需要增加注册器的情况
		var pfRegister int32
		if currentPoolSize > int32(float32(m.preferredTokenCount)*0.25) { // 池填充 25% 以上
			// 启动部分注册器 根据池的当前大小动态调整
			pfRegister = int32(float32(m.maxResisters) * (1 - float32(currentPoolSize)/float32(m.preferredTokenCount)))
			if pfRegister < 1 { // 最少启动一个注册器
				pfRegister = 1
			}
		} else { // 池填充 25% 以下 全部启动!
			pfRegister = m.maxResisters // 启动所有注册器
		}
		m.callRegister(pfRegister)
	}
}

// 调用注册器
func (m *LoginTokenManager) callRegister(preferredCount int32) {
	// 限制最大同时启用的注册器数量
	if m.currentRunningResisters+preferredCount > m.maxResisters {
		preferredCount = m.maxResisters - m.currentRunningResisters
	}
	for i := int32(0); i < preferredCount; i++ {
		go func() {
			atomic.AddInt32(&m.currentRunningResisters, 1) // 正在运行的注册器数量加一
			log.Println("RUN REG")
			token, err := RunRegister(m.hostUrl)
			if err == nil {
				m.tokenPool.Add(&token)
				log.Println("POOL ADD NOW:", m.tokenPool.Size())
			} // TODO: ERROR LOG NEEDED
			defer atomic.AddInt32(&m.currentRunningResisters, -1) // 正在运行的注册器数量减一
		}()
	}
}

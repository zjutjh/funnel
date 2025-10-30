package loginTokenManager

import (
	"container/list"
	"sync"
	"time"
)

type loginTokenPool struct {
	mu sync.Mutex
	l  list.List
}

// Get 从池中获取首个 LoginToken, 若池为空则返回 nil
func (p *loginTokenPool) Get() *LoginToken {
	p.mu.Lock()
	defer p.mu.Unlock()
	if e := p.l.Front(); e != nil {
		return p.l.Remove(e).(*LoginToken)
	}
	return nil
}

func (p *loginTokenPool) Add(v *LoginToken) {
	p.mu.Lock()
	p.l.PushBack(v)
	p.mu.Unlock()
}

func (p *loginTokenPool) Size() int32 {
	size := p.l.Len()
	return int32(size)
}

func (p *loginTokenPool) RemoveExpired(before time.Time) {
	p.mu.Lock()
	defer p.mu.Unlock()
	for {
		if e := p.l.Front(); e != nil {
			if e.Value.(*LoginToken).RequestTime.Before(before) {
				p.l.Remove(e)
			} else {
				break // 遇到未过期 token 结束
			}
		} else {
			break // 池为空结束
		}
	}
}

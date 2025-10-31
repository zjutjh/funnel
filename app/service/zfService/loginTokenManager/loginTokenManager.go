package loginTokenManager

import (
	"fmt"
	"log"
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
	preferredTokenCount     int64  // 期望保持的 Token 数量
	currentRunningResisters int64  // 当前正在运行的注册器数量
	maxResisters            int64  // 最大注册器数量
	tokenExpirationTimeSec  int32  // Token 过期时间
	routineMaintenanceSec   int32  // 两次例行池维护的间隔
	activeMaintenanceEvery  int32  // n 次 get 请求触发一次池维护
	hostUrl                 string // 注册器目标主机地址
	// VARS
	tokenPool         *loginTokenPool // Token 池实例
	getAfterLastActMt int32           // 上次主动维护后进行的请求数
}

var LoginTkMgr *LoginTokenManager = &LoginTokenManager{}

func Init(host string) {
	LoginTkMgr.Init(host)
}

func (m *LoginTokenManager) Init(host string) {
	m.tokenPool = &loginTokenPool{}
	m.hostUrl = host
	m.preferredTokenCount = 10
	m.maxResisters = 10
	m.currentRunningResisters = 0
	m.tokenExpirationTimeSec = 2 * 60 * 60
	m.routineMaintenanceSec = 15
	m.activeMaintenanceEvery = 5
	m.getAfterLastActMt = 0

	go m.runPoolRoutineMt()
}

func (m *LoginTokenManager) GetAToken() (*LoginToken, error) {
	log.Println("Pool Remain:", m.tokenPool.Size())
	tk, err := m.getTokenFromPool()
	if tk != nil && err == nil { //从池取 token 成功, 直接返回
		return tk, nil
	}
	log.Println("WARN Cannot get token in pool")
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
func (m *LoginTokenManager) runPoolRoutineMt() {
	ticker := time.NewTicker(time.Duration(m.routineMaintenanceSec) * time.Second) // 定时触发
	defer ticker.Stop()                                                            // 确保 ticker 在函数退出时被停止

	for range ticker.C {
		log.Println("Routine Mt")
		m.poolMaintenance()
	}
}

// 运行池维护 以下三种情况会触发池维护:
// 1. 定时维护
// 2. 主动维护 (每 n 次 get 请求触发一次)
// 3. 注册器完成后检查池状态发现池不满时
func (m *LoginTokenManager) poolMaintenance() {

	log.Printf("Pool Mt %d\n", m.tokenPool.Size())

	m.tokenPool.RemoveExpired(time.Now().Add(-time.Duration(m.tokenExpirationTimeSec) * time.Second)) // 移除过期 token
	currentPoolSize := m.tokenPool.Size()                                                             // 获取当前池中 Token 数量
	if currentPoolSize < m.preferredTokenCount {                                                      // 需要增加注册器的情况
		var pfRegister int64
		if currentPoolSize > int64(float64(m.preferredTokenCount)*0.25) { // 池填充 25% 以上
			// 启动部分注册器 根据池的当前大小动态调整
			pfRegister = int64(float32(m.maxResisters) * (1 - float32(currentPoolSize)/float32(m.preferredTokenCount)))
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
func (m *LoginTokenManager) callRegister(preferredCount int64) {
	// 限制最大同时启用的注册器数量
	if m.currentRunningResisters+preferredCount > m.maxResisters {
		preferredCount = m.maxResisters - m.currentRunningResisters
	}
	for i := int64(0); i < preferredCount; i++ {
		go func() {
			atomic.AddInt64(&m.currentRunningResisters, 1) // 正在运行的注册器数量加一
			log.Println("RUN REG NOW:", m.currentRunningResisters)
			token, err := RunRegister(m.hostUrl)
			if err == nil {
				m.tokenPool.Add(&token)
				log.Println("POOL ADD NOW:", m.tokenPool.Size())
			} // TODO: ERROR LOG NEEDED
			log.Printf("REG FINISH: %v\n", err)
			if m.tokenPool.Size() < m.preferredTokenCount &&
				atomic.LoadInt64(&m.currentRunningResisters) < m.maxResisters {
				m.poolMaintenance() // 池不满且有注册器名额时 继续维护池
			}
			defer atomic.AddInt64(&m.currentRunningResisters, -1) // 正在运行的注册器数量减一
		}()
	}
}

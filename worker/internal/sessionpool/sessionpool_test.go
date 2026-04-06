package sessionpool

import (
	"funnel/worker/internal/config"
	"funnel/worker/internal/svc"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/zeromicro/go-zero/core/conf"
)

type testSessionPoolConfig struct {
	ZF config.ZFConfig `json:"zf,optional"`
}

func loadSessionPoolTestConfig(t *testing.T) config.ZFConfig {
	t.Helper()

	var cfg testSessionPoolConfig
	require.NoError(t, conf.Load("../../etc/config.yaml", &cfg))

	if cfg.ZF.BaseURL == "" {
		t.Skip("integration test requires zf.base_url in etc/config.yaml")
	}

	return cfg.ZF
}

func TestSessionPool1(t *testing.T) {
	cfg := loadSessionPoolTestConfig(t)
	p := New(svc.NewServiceContext(config.Config{ZF: cfg}))

	go func() {
		p.Run(t.Context())
	}()
	for range 200 {
		cookie, err := p.Get(t.Context())
		require.NoError(t, err)
		require.NotNil(t, cookie)
		assert.NotZero(t, cookie.JSessionID)
		assert.NotZero(t, cookie.Route)
		t.Logf("cookie = %v", cookie)
	}
}

func TestSessionPool2(t *testing.T) {
	cfg := loadSessionPoolTestConfig(t)
	p := New(svc.NewServiceContext(config.Config{ZF: cfg}))

	go func() {
		p.Run(t.Context())
	}()
	// 等待预热完毕
	for !(len(p.cookies) == p.maxSize) {
	}
	// 模拟大量请求, 远超pool的容量
	for range 2 * p.maxSize {
		cookie, err := p.Get(t.Context())
		require.NoError(t, err)
		require.NotNil(t, cookie)
		assert.NotZero(t, cookie.JSessionID)
		assert.NotZero(t, cookie.Route)
		t.Logf("cookie = %v", cookie)
	}
}

func TestSessionPool3(t *testing.T) {
	cfg := loadSessionPoolTestConfig(t)
	p := New(svc.NewServiceContext(config.Config{ZF: cfg}))

	go func() {
		p.Run(t.Context())
	}()
	// 等待预热完毕
	for !(len(p.cookies) == p.maxSize) {
	}
	for range p.maxSize {
		cookie, err := p.Get(t.Context())
		require.NoError(t, err)
		require.NotNil(t, cookie)
		assert.NotZero(t, cookie.JSessionID)
		assert.NotZero(t, cookie.Route)
		t.Logf("cookie = %v", cookie)
	}
	// 上面已经把pool里的cookie都取完了, 再使用TryGet会直接返回错误
	_, err := p.TryGet(t.Context())
	require.Error(t, err)
}

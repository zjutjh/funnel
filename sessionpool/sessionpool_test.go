package sessionpool

import (
	"funnel/register"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/zjutjh/mygo/foundation/kernel"
)

func TestMain(m *testing.M) {
	kernel.Bootstrap("../conf", register.Boot)
	m.Run()
}

func TestSessionPool1(t *testing.T) {
	go func() {
		New().WithContext(t.Context()).Run()
	}()
	for range 200 {
		cookie, err := Pick().Get(t.Context())
		require.NoError(t, err)
		require.NotNil(t, cookie)
		assert.NotZero(t, cookie.JSessionID)
		assert.NotZero(t, cookie.Route)
		t.Logf("cookie = %v", cookie)
	}
}

func TestSessionPool2(t *testing.T) {
	go func() {
		New().WithContext(t.Context()).Run()
	}()
	p := Pick()
	// 等待预热完毕
	for !(len(p.cookies) == p.maxSize) {
	}
	// 模拟大量请求, 远超pool的容量
	for range 2 * p.maxSize {
		cookie, err := Pick().Get(t.Context())
		require.NoError(t, err)
		require.NotNil(t, cookie)
		assert.NotZero(t, cookie.JSessionID)
		assert.NotZero(t, cookie.Route)
		t.Logf("cookie = %v", cookie)
	}
}

func TestSessionPool3(t *testing.T) {
	go func() {
		New().WithContext(t.Context()).Run()
	}()
	p := Pick()
	// 等待预热完毕
	for !(len(p.cookies) == p.maxSize) {
	}
	for range p.maxSize {
		cookie, err := Pick().Get(t.Context())
		require.NoError(t, err)
		require.NotNil(t, cookie)
		assert.NotZero(t, cookie.JSessionID)
		assert.NotZero(t, cookie.Route)
		t.Logf("cookie = %v", cookie)
	}
	// 上面已经把pool里的cookie都取完了, 再使用TryGet会直接返回错误
	_, err := Pick().TryGet(t.Context())
	require.Error(t, err)
}

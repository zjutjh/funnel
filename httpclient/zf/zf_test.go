package zfClient

import (
	"funnel/comm"
	"funnel/register"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/zjutjh/mygo/foundation/kernel"
)

func TestMain(m *testing.M) {
	kernel.Bootstrap("../../conf", register.Boot)
	m.Run()
}

func TestBypassCaptcha(t *testing.T) {
	sem := make(chan struct{}, 10)
	var wg sync.WaitGroup

	for range 100 {
		sem <- struct{}{}
		wg.Add(1)
		go func() {
			defer wg.Done()
			defer func() { <-sem }()
			result, err := New(t.Context()).BypassCaptcha()
			require.NoError(t, err)
			require.NotNil(t, result)
			assert.NotEmpty(t, result.JSessionID)
			assert.NotEmpty(t, result.Route)
		}()
	}
	wg.Wait()
}

func TestLoginByCaptcha(t *testing.T) {
	zf := New(t.Context())
	username := comm.BizConf.ZF.Public.Username
	password := comm.BizConf.ZF.Public.Password
	cookies, err := zf.
		LoginByCaptcha(username, password)
	assert.NoError(t, err)
	t.Logf("验证成功: %s", cookies)
}

func TestGetCurrentSchoolTerm(t *testing.T) {
	zf := New(t.Context())
	username := comm.BizConf.ZF.Public.Username
	password := comm.BizConf.ZF.Public.Password
	cookies, err := zf.LoginByCaptcha(username, password)

	assert.NoError(t, err)
	info, err := zf.GetCurrentSchoolTerm(cookies)
	assert.NoError(t, err)
	assert.NotEmpty(t, info.Term)
	assert.NotEmpty(t, info.Year)
	t.Logf("当前学年学期: %s", info)
}

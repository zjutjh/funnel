package test

import (
	"funnel/worker/internal/config"
	zfClient "funnel/worker/internal/httpclient/zf"
	"funnel/worker/internal/svc"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/zeromicro/go-zero/core/conf"
)

func loadZFTestConfig(t *testing.T) config.Config {
	t.Helper()

	var cfg config.Config
	require.NoError(t, conf.Load("../../etc/config.yaml", &cfg))

	if cfg.ZF.BaseURL == "" {
		t.Skip("integration test requires zf.base_url in etc/config.yaml")
	}

	return cfg
}

func TestBypassCaptcha(t *testing.T) {
	cfg := loadZFTestConfig(t)
	svcCtx := svc.NewServiceContext(cfg)

	sem := make(chan struct{}, 2)
	var wg sync.WaitGroup

	for range 10 {
		sem <- struct{}{}
		wg.Add(1)
		go func() {
			defer wg.Done()
			defer func() { <-sem }()
			result, err := zfClient.New(svcCtx).BypassCaptcha()
			require.NoError(t, err)
			require.NotNil(t, result)
			assert.NotEmpty(t, result.JSessionID)
			assert.NotEmpty(t, result.Route)
		}()
	}
	wg.Wait()
}

func TestLoginByCaptcha(t *testing.T) {
	cfg := loadZFTestConfig(t)
	svcCtx := svc.NewServiceContext(cfg)

	username := cfg.ZF.Public.Username
	password := cfg.ZF.Public.Password
	if username == "" || password == "" {
		t.Skip("integration test requires zf.public.username and zf.public.password")
	}

	zf := zfClient.New(svcCtx)
	cookies, err := zf.
		LoginByCaptcha(username, password)
	assert.NoError(t, err)
	t.Logf("验证成功: %s", cookies)
}

func TestGetCurrentSchoolTerm(t *testing.T) {
	cfg := loadZFTestConfig(t)
	svcCtx := svc.NewServiceContext(cfg)

	username := cfg.ZF.Public.Username
	password := cfg.ZF.Public.Password
	if username == "" || password == "" {
		t.Skip("integration test requires zf.public.username and zf.public.password")
	}

	zf := zfClient.New(svcCtx)
	cookies, err := zf.LoginByCaptcha(username, password)

	assert.NoError(t, err)
	info, err := zf.GetCurrentSchoolTerm(cookies)
	assert.NoError(t, err)
	assert.NotEmpty(t, info.Term)
	assert.NotEmpty(t, info.Year)
	t.Logf("当前学年学期: %s", info)
}

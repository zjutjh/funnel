package zfClient

import (
	"funnel/comm"
	"funnel/register"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zjutjh/mygo/foundation/kernel"
)

func TestMain(m *testing.M) {
	kernel.Bootstrap("../../conf", register.Boot)
	m.Run()
}

func TestBypassCaptcha(t *testing.T) {
	result, err := New(t.Context()).BypassCaptcha()
	if err != nil {
		t.Fatalf("验证失败: %v", err)
	}
	t.Logf("验证成功: %s", result)
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

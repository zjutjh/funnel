package zfClient

import (
	"context"
	"funnel/register"
	"testing"

	"github.com/zjutjh/mygo/foundation/kernel"
)

func TestMain(m *testing.M) {
	kernel.Bootstrap("../../conf", register.Boot)
	m.Run()
}

func TestBypassCaptcha(t *testing.T) {
	login, err := New(context.Background()).BypassCaptcha()
	if err != nil {
		t.Fatalf("验证失败: %v", err)
	}
	t.Logf("验证成功: %s", login)
}

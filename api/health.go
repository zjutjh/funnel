package api

import (
	"reflect"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zjutjh/mygo/foundation/reply"
	"github.com/zjutjh/mygo/kit"
	"github.com/zjutjh/mygo/nlog"
	"github.com/zjutjh/mygo/swagger"

	"app/comm"
)

// HealthHandler Api router注册点
func HealthHandler() gin.HandlerFunc {
	api := HealthApi{}
	swagger.CM[runtime.FuncForPC(reflect.ValueOf(hfHealth).Pointer()).Name()] = api
	return hfHealth
}

type HealthApi struct {
	Info     struct{} `name:"健康检查接口" desc:"健康检查接口"`
	Request  HealthApiRequest
	Response HealthApiResponse
}

type HealthApiRequest struct {
}

type HealthApiResponse struct {
	Now int64 `json:"now" desc:"当前服务器时间戳(Unix秒级)"`
}

// Run Api业务逻辑执行点
func (h *HealthApi) Run(ctx *gin.Context) kit.Code {
	h.Response.Now = time.Now().Unix()
	return comm.CodeOK
}

// Run Api初始化 进行参数校验和绑定
func (h *HealthApi) Init(ctx *gin.Context) (err error) {
	return err
}

// hfHealth Api执行入口
func hfHealth(ctx *gin.Context) {
	api := &HealthApi{}
	err := api.Init(ctx)
	if err != nil {
		nlog.Pick().WithContext(ctx).WithError(err).Warn("参数绑定校验错误")
		reply.Fail(ctx, comm.CodeParameterInvalid)
		return
	}
	code := api.Run(ctx)
	if !ctx.IsAborted() {
		if code == comm.CodeOK {
			reply.Success(ctx, api.Response)
		} else {
			reply.Fail(ctx, code)
		}
	}
}

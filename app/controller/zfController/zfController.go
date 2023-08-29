package zfController

import (
	"encoding/json"
	"funnel/app/controller"
	"funnel/app/errors"
	"funnel/app/model"
	"funnel/app/service/zfService"
	"funnel/app/utils"
	"github.com/gin-gonic/gin"
)

// GetScoreDetail
//
//		@Summary 正方教务详细成绩
//		@Description 正方教务详细成绩
//		@Tags 正方
//		@Produce  json
//		@Param term body string true "学期"
//		@Param year body string true "年份"
//		@Param username body string true "用户名"
//		@Param password body string true "密码"
//	    @Param type body string true "登录类型"
//		@Success 200 json  {"code":200,"data":{object},"msg":"OK"}
//		@Failure 400 json  {"code":400,"data":null,"msg":""}
//		@Router /student/zfService/score/info [post]
func GetScoreDetail(context *gin.Context) {
	_, _ = ZFTermInfoHandle(context, zfService.GetScoreDetail)
	return
}

// GetScore
//
//		@Summary 正方教务成绩
//		@Description 正方教务成绩
//		@Tags 正方
//		@Produce  json
//		@Param term body string true "学期"
//		@Param year body string true "年份"
//		@Param username body string true "用户名"
//		@Param password body string true "密码"
//	    @Param type body string true "登录类型"
//		@Success 200 json  {"code":200,"data":{object},"msg":"OK"}
//		@Failure 400 json  {"code":400,"data":null,"msg":""}
//		@Router /student/zfService/score/info [post]
func GetScore(context *gin.Context) {
	_, _ = ZFTermInfoHandle(context, zfService.GetScore)
	return
}

// GetMidTermScore
//
//		@Summary 正方教务期中成绩
//		@Description 正方教务期中成绩
//		@Tags 正方
//		@Produce  json
//		@Param term body string true "学期"
//		@Param year body string true "年份"
//		@Param username body string true "用户名"
//		@Param password body string true "密码"
//	    @Param type body string true "登录类型"
//		@Success 200 json  {"code":200,"data":{object},"msg":"OK"}
//		@Failure 400 json  {"code":400,"data":null,"msg":""}
//		@Router /student/zfService/score/info [post]
func GetMidTermScore(context *gin.Context) {
	_, _ = ZFTermInfoHandle(context, zfService.GetMidTermScore)
	return
}

// GetExamInfo
//
//		@Summary 正方教务考试信息
//		@Description 正方教务考试信息
//		@Tags 正方
//		@Produce  json
//		@Param term body string true "学期"
//		@Param year body string true "年份"
//		@Param username body string true "用户名"
//		@Param password body string true "密码"
//	    @Param type body string true "登录类型"
//		@Success 200 json  {"code":200,"data":{object},"msg":"OK"}
//		@Failure 400 json  {"code":400,"data":null,"msg":""}
//		@Router /student/zfService/exam [post]
func GetExamInfo(context *gin.Context) {
	_, _ = ZFTermInfoHandle(context, zfService.GetExamInfo)
	return
}

// GetClassTable
//
//		@Summary 正方教务课表
//		@Description 正方教务课表
//		@Tags 正方
//		@Produce  json
//		@Param term body string true "学期"
//		@Param year body string true "年份"
//		@Param username body string true "用户名"
//		@Param password body string true "密码"
//	    @Param type body string true "登录类型"
//		@Success 200 json  {"code":200,"data":{object},"msg":"OK"}
//		@Failure 400 json  {"code":400,"data":null,"msg":""}
//		@Router /student/zfService/score [post]
func GetClassTable(context *gin.Context) {
	_, _ = ZFTermInfoHandle(context, zfService.GetLessonsTable)
	return
}

// GetProgInfo
//
//		@Summary 正方教务考试信息
//		@Description 正方教务考试信息
//		@Tags 正方
//		@Produce  json
//		@Param username body string true "用户名"
//		@Param password body string true "密码"
//	    @Param type body string true "登录类型"
//		@Success 200 json  {"code":200,"data":{object},"msg":"OK"}
//		@Failure 400 json  {"code":400,"data":null,"msg":""}
//		@Router /student/zfService/program [post]
func GetProgInfo(context *gin.Context) {

	user, err := controller.LoginHandle(context, zfService.GetUser, false)
	if err != nil {
		return
	}
	result, err := zfService.GetTrainingPrograms(user)
	if err == errors.ERR_SESSION_EXPIRES {
		user, err = controller.LoginHandle(context, zfService.GetUser, false)
		if err != nil {
			return
		}
		result, err = zfService.GetTrainingPrograms(user)
	}

	if err != nil {
		controller.ErrorHandle(context, err)
		return
	}

	context.Data(200, "application/pdf", result)
	return
}

// GetRoomInfo
//
//		@Summary 正方教务考试信息
//		@Description 正方教务考试信息
//		@Tags 正方
//		@Produce  json
//		@Param term body string true "学期"
//		@Param year body string true "年份"
//		@Param campus body string true "校区"
//		@Param weekday body string true "星期几 1，2，3"
//		@Param week body string true "第几周的2次幂的和"
//		@Param sections body string true "第几节课的2次幂的和"
//		@Param username body string true "用户名"
//		@Param password body string true "密码"
//	    @Param type body string true "登录类型"
//		@Success 200 json  {"code":200,"data":{object},"msg":"OK"}
//		@Failure 400 json  {"code":400,"data":null,"msg":""}
//		@Router /student/zfService/room [post]
func GetRoomInfo(context *gin.Context) {
	user, err := controller.LoginHandle(context, zfService.GetUser, false)
	if err != nil {
		return
	}
	isValid := utils.CheckPostFormEmpty(
		context,
		[]string{"week", "sections", "campus", "weekday"},
	)

	if !isValid {
		utils.ContextDataResponseJson(context, utils.FailResponseJson(errors.InvalidArgs, nil))
		return
	}

	result, err := zfService.GetEmptyRoomInfo(user, context.PostForm("year"), context.PostForm("term"), context.PostForm("campus"), context.PostForm("weekday"), context.PostForm("week"), context.PostForm("sections"))
	if err == errors.ERR_SESSION_EXPIRES {
		user, err = controller.LoginHandle(context, zfService.GetUser, false)
		if err != nil {
			return
		}
		result, err = zfService.GetEmptyRoomInfo(user, context.PostForm("year"), context.PostForm("term"), context.PostForm("campus"), context.PostForm("weekday"), context.PostForm("week"), context.PostForm("sections"))
	}

	if err != nil {
		controller.ErrorHandle(context, err)
		return
	}

	var f model.EmptyRoomRawInfo
	err = json.Unmarshal([]byte(result), &f)
	if err != nil {
		controller.ErrorHandle(context, err)
		return
	}
	utils.ContextDataResponseJson(context, utils.SuccessResponseJson(model.TransformEmptyRoom(&f)))
	return

}

package controller

import (
	"encoding/json"
	"funnel/app/errors"
	"funnel/app/helps"
	"funnel/app/model"
	"funnel/app/service"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var cardSystem = service.CardSystem{}

func CardLogin(context *gin.Context) {
	err := cardLoginHandle(context)
	if err == nil {
		helps.ContextDataResponseJson(context, helps.SuccessResponseJson(nil))
	}
}

func CardBalance(context *gin.Context) {
	session := sessions.Default(context)
	cardJson := session.Get("card").([]byte)

	if string(cardJson) == "{}" {
		helps.ContextDataResponseJson(context, helps.FailResponseJson(errors.NotLogin, nil))
		return
	}

	user := model.CardUser{}
	_ = json.Unmarshal(cardJson, &user)
	balance := cardSystem.GetCurrentBalance(&user)
	helps.ContextDataResponseJson(context, helps.SuccessResponseJson(balance))
}
func CardToday(context *gin.Context) {
	session := sessions.Default(context)
	cardJson := session.Get("card").([]byte)

	if string(cardJson) == "{}" {
		helps.ContextDataResponseJson(context, helps.FailResponseJson(errors.NotLogin, nil))
		return
	}

	user := model.CardUser{}
	_ = json.Unmarshal(cardJson, &user)
	balance := cardSystem.GetCardToday(&user)
	helps.ContextDataResponseJson(context, helps.SuccessResponseJson(balance))
}
func CardHistory(context *gin.Context) {
	isValid := helps.CheckPostFormEmpty(
		context,
		[]string{"year", "month"},
	)

	if !isValid {
		helps.ContextDataResponseJson(context, helps.FailResponseJson(errors.InvalidArgs, nil))
		return
	}

	session := sessions.Default(context)
	cardJson := session.Get("card").([]byte)

	if string(cardJson) == "{}" {
		helps.ContextDataResponseJson(context, helps.FailResponseJson(errors.NotLogin, nil))
		return
	}

	user := model.CardUser{}
	_ = json.Unmarshal(cardJson, &user)
	balance := cardSystem.GetCardHistory(&user, context.PostForm("year"), context.PostForm("month"))
	helps.ContextDataResponseJson(context, helps.SuccessResponseJson(balance))
}
func cardLoginHandle(context *gin.Context) error {
	isValid := helps.CheckPostFormEmpty(
		context,
		[]string{"username", "password"},
	)

	if !isValid {
		helps.ContextDataResponseJson(context, helps.FailResponseJson(errors.InvalidArgs, nil))
		return errors.ERR_INVALID_ARGS
	}

	user := model.CardUser{Username: context.PostForm("username"), Password: context.PostForm("password")}
	err := cardSystem.Login(&user)

	if err == errors.ERR_WRONG_PASSWORD {
		helps.ContextDataResponseJson(context, helps.FailResponseJson(errors.WrongPassword, nil))
		return err
	}
	if err != nil {
		helps.ContextDataResponseJson(context, helps.FailResponseJson(errors.UnKnown, nil))
		return err
	}

	session := sessions.Default(context)
	cardJson, _ := json.Marshal(user)
	session.Set("card", cardJson)
	_ = session.Save()

	return nil
}

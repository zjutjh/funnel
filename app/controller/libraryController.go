package controller

import (
	"encoding/json"
	"funnel/app/errors"
	"funnel/app/helps"
	"funnel/app/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var system = model.LibrarySystem{}

func LibraryLogin(context *gin.Context) {
	user := model.LibraryUser{Username: context.PostForm("username"), Password: context.PostForm("password")}
	err := system.Login(&user)

	if err != nil {
		context.Data(200, "application/json", helps.FailResponseJson(errors.WrongPassword, nil))
		return
	}

	session := sessions.Default(context)
	libraryJson, _ := json.Marshal(user)
	session.Set("library", libraryJson)
	_ = session.Save()

	context.Data(200, "application/json", helps.SuccessResponseJson(nil))
}

func LibraryBorrowHistory(context *gin.Context) {
	session := sessions.Default(context)
	libraryJson := session.Get("library").([]byte)
	if string(libraryJson) == "{}" {
		context.Data(403, "application/json", helps.FailResponseJson(errors.NotLogin, nil))
		return
	}

	user := model.LibraryUser{}
	_ = json.Unmarshal(libraryJson, &user)
	books := system.GetBorrowHistory(&user)
	context.Data(200, "application/json", helps.SuccessResponseJson(books))
}

func LibraryCurrentBorrow(context *gin.Context) {
	session := sessions.Default(context)
	libraryJson := session.Get("library").([]byte)
	if string(libraryJson) == "{}" {
		context.Data(403, "application/json", helps.FailResponseJson(errors.NotLogin, nil))
		return
	}

	user := model.LibraryUser{}
	_ = json.Unmarshal(libraryJson, &user)
	books := system.GetCurrentBorrow(&user)
	context.Data(200, "application/json", helps.SuccessResponseJson(books))
}

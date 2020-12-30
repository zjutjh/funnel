package controller

import (
	"encoding/json"
	"funnel/app/errors"
	"funnel/app/helps"
	"funnel/app/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var system = model.LibrarySystem{RootUrl: "http://172.16.19.163/jwglxt/"}

func LibraryLogin(context *gin.Context) {
	user := model.LibraryUser{Username: context.PostForm("username"), Password: context.PostForm("password")}
	_ = system.Login(&user)

	session := sessions.Default(context)

	libraryJson, _ := json.Marshal(session)
	session.Set("library", libraryJson)
	_ = session.Save()
	context.Data(200, "application/json", helps.SuccessResponseJson(nil))
}

func LibraryBorrowHistory(context *gin.Context) {
	session := sessions.Default(context)
	libraryJson := session.Get("library").([]byte)

	if libraryJson == nil {
		context.Data(403, "application/json", helps.FailResponseJson(errors.NotLogin,nil))
	}

	user := model.LibraryUser{}
	_ = json.Unmarshal(libraryJson, &user)
	str := system.GetBorrowHistory(&user)

	context.Data(200, "application/json", helps.SuccessResponseJson(str))
}

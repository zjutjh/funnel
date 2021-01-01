package helps

import (
	"encoding/json"
	"funnel/app/errors"
	"funnel/app/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"time"
)

func CheckPostFormEmpty(context *gin.Context, items []string) bool {
	for i := range items {
		_, has := context.GetPostForm(items[i])
		if has == false {
			return false
		}
	}
	return true
}

func CheckZFSession(context *gin.Context, sessionName string) (*model.ZFUser, error) {
	session := sessions.Default(context)
	ZFSession := session.Get(sessionName).([]byte)
	if ZFSession == nil {
		return nil, errors.ERR_SESSION_NOT_EXIST
	}

	user := &model.ZFUser{}
	err := json.Unmarshal(ZFSession, user)

	if err != nil {
		return nil, errors.ERR_JSON_DESER
	}
	if user.Session.Expires.Unix() >= time.Now().Unix() {
		return nil, errors.ERR_SESSION_EXPIRES
	}
	return user, nil
}
func CheckLibrarySession(context *gin.Context, sessionName string) (*model.LibraryUser, error) {
	session := sessions.Default(context)
	LibrarySession := session.Get(sessionName).([]byte)
	if LibrarySession == nil {
		return nil, errors.ERR_SESSION_NOT_EXIST
	}
	user := &model.LibraryUser{}
	err := json.Unmarshal(LibrarySession, user)
	if err != nil {
		return nil, errors.ERR_JSON_DESER
	}
	if user.Session.Expires.Unix() >= time.Now().Unix() {
		return nil, errors.ERR_SESSION_EXPIRES
	}
	return user, nil
}

func CheckCardSession(context *gin.Context, sessionName string) (*model.CardUser, error) {
	session := sessions.Default(context)
	CardSession := session.Get(sessionName).([]byte)
	if CardSession == nil {
		return nil, errors.ERR_SESSION_NOT_EXIST
	}
	user := &model.CardUser{}
	err := json.Unmarshal(CardSession, user)
	if err != nil {
		return nil, errors.ERR_JSON_DESER
	}
	if user.Session.Expires.Unix() >= time.Now().Unix() {
		return nil, errors.ERR_SESSION_EXPIRES
	}
	return user, nil
}

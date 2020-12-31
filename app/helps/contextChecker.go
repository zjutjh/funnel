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

func CheckSession(context *gin.Context, sessionName string) (*model.ZFUser, error) {
	session := sessions.Default(context)
	ZFSession := session.Get(sessionName).([]byte)
	if ZFSession == nil {
		return nil, errors.ERR_SESSION_NOT_EXIST
	}
	user := &model.ZFUser{}
	_ = json.Unmarshal(ZFSession, user)
	if user.Session.Expires.Unix() >= time.Now().Unix() {
		return nil, errors.ERR_SESSION_EXPIRES
	}
	return user, nil
}

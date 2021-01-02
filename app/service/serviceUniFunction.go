package service

import (
	"encoding/json"
	funnel "funnel/app"
	"funnel/app/errors"
	"funnel/app/model"
)

var ZFPrefix = "funnel-ZF-"
var LibraryPrefix = "funnel-Library-"
var CardPrefix = "funnel-Card-"

func GetUser(prefix string, username string) (*model.User, error) {
	Session := funnel.Redis.Get(prefix + username)
	if Session == nil {
		return nil, errors.ERR_SESSION_NOT_EXIST
	}
	user := &model.User{}
	err := json.Unmarshal([]byte(Session.Val()), user)

	if err != nil {
		funnel.Redis.Del(prefix + username)
		return nil, errors.ERR_JSON_DESER
	}
	return user, nil
}

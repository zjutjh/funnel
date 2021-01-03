package service

import (
	"encoding/json"
	funnel "funnel/app"
	"funnel/app/errors"
	"funnel/app/model"
	"hash/crc32"
	"net/http"
	"strconv"
	"time"
)

var ZFPrefix = "funnel-ZF-"
var LibraryPrefix = "funnel-Library-"
var CardPrefix = "funnel-Card-"

func GetUser(prefix string, username string, password string) (*model.User, error) {
	Session := funnel.Redis.Get(prefix + username + strconv.Itoa(int(crc32.ChecksumIEEE([]byte(password)))))
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

func SetUser(prefix string, username string, password string, cookie *http.Cookie) (*model.User, error) {
	user := model.User{Username: username, Password: password, Session: *cookie}
	userJson, _ := json.Marshal(user)
	funnel.Redis.Set(prefix+username+strconv.Itoa(int(crc32.ChecksumIEEE([]byte(password)))), string(userJson), cookie.Expires.Sub(time.Now().Add(time.Minute*5)))
	return &user, nil
}

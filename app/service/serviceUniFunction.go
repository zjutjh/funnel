package service

import (
	"encoding/json"
	"funnel/app/errors"
	"funnel/app/model"
	"funnel/config"
	"hash/crc32"
	"net/http"
	"strconv"
	"time"
)

var ZFPrefix = "funnel-ZF-"
var LibraryPrefix = "funnel-Library-"
var CardPrefix = "funnel-Card-"

func getRediskey(prefix, username, password string) string {
	return prefix + username + strconv.Itoa(int(crc32.ChecksumIEEE([]byte(password))))
}

func GetUser(prefix string, username string, password string) (*model.User, error) {
	Session := config.Redis.Get(getRediskey(prefix, username, password))
	if Session == nil {
		return nil, errors.ERR_SESSION_NOT_EXIST
	}
	user := &model.User{}
	err := json.Unmarshal([]byte(Session.Val()), user)

	if err != nil {
		config.Redis.Del(getRediskey(prefix, username, password))
		return nil, errors.ERR_JSON_DESER
	}
	return user, nil
}

func SetUser(prefix string, username string, password string, cookie *http.Cookie) (*model.User, error) {
	user := model.User{Username: username, Password: password, Session: *cookie}
	userJson, _ := json.Marshal(user)
	config.Redis.Set(getRediskey(prefix, username, password), string(userJson), cookie.Expires.Sub(time.Now().Add(time.Minute*5)))
	return &user, nil
}

func ForgetUser(prefix string, username string, password string) {
	config.Redis.Del(getRediskey(prefix, username, password))
}
func ForgetAllUser(prefix string) {
	res, _ := config.Redis.Keys(prefix + "*").Result()
	for _, v := range res {
		config.Redis.Del(v)
	}
}

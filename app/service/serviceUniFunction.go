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
var ZFLoginTkPrefix = "funnel-ZFLoginTk"

func genRedisKeyUser(prefix, username, password string) string {
	return prefix + username + strconv.Itoa(int(crc32.ChecksumIEEE([]byte(password))))
}

func GetUser(prefix string, username string, password string) (*model.User, error) {
	Session := config.Redis.Get(genRedisKeyUser(prefix, username, password))
	if Session == nil {
		return nil, errors.ERR_SESSION_NOT_EXIST
	}
	user := &model.User{}
	err := json.Unmarshal([]byte(Session.Val()), user)

	if err != nil {
		config.Redis.Del(genRedisKeyUser(prefix, username, password))
		return nil, errors.ERR_JSON_DESER
	}
	return user, nil
}

func SetUser(prefix string, username string, password string, sessionCookie *http.Cookie, routeCookie *http.Cookie) (*model.User, error) {
	user := model.User{Username: username, Password: password, Session: *sessionCookie, Route: *routeCookie}
	userJson, _ := json.Marshal(user)
	config.Redis.Set(genRedisKeyUser(prefix, username, password), string(userJson), 1*time.Second)
	return &user, nil
}

func ForgetUser(prefix string, username string, password string) {
	config.Redis.Del(genRedisKeyUser(prefix, username, password))
}

func ForgetAllUser(prefix string) {
	res, _ := config.Redis.Keys(prefix + "*").Result()
	for _, v := range res {
		config.Redis.Del(v)
	}
}

func ForgetUserByUsername(prefix, username string) {
	res, _ := config.Redis.Keys(prefix + username + "*").Result()
	for _, v := range res {
		config.Redis.Del(v)
	}
}

package loginTokenManager

import (
	"encoding/json"
	"funnel/app/service"
	"funnel/config"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

type loginTokenPool struct {
}

// Get 从池中获取首个 LoginToken, 若池为空则返回 nil
func (p *loginTokenPool) Get() *LoginToken {
	// 取出一个时间值(即分值)最小的元素
	tk, err := config.Redis.ZPopMin(service.ZFLoginTkPrefix).Result()
	if err != nil || len(tk) == 0 { // 空与错误判断
		return nil
	}
	loginToken := &LoginToken{}
	d, ok := tk[0].Member.(string) // 类型断言 如果此步出错应为设计问题
	if !ok {
		return nil
	}
	err = json.Unmarshal([]byte(d), loginToken) // 反序列化
	if err != nil {
		return nil
	}
	return loginToken
}

func (p *loginTokenPool) Add(v *LoginToken) {
	tokenJson, err := json.Marshal(v)
	if err != nil {
		return
	}
	config.Redis.ZAdd(service.ZFLoginTkPrefix,
		redis.Z{
			Score:  float64(v.RequestTime.UnixMicro()),
			Member: tokenJson}) // 按照时间排序的有序集合
}

func (p *loginTokenPool) Size() int64 {
	size, err := config.Redis.ZCard(service.ZFLoginTkPrefix).Result()
	if err != nil {
		return 0
	}
	return int64(size)
}

func (p *loginTokenPool) RemoveExpired(before time.Time) {
	config.Redis.ZRemRangeByScore(service.ZFLoginTkPrefix, "0", strconv.FormatInt(before.UnixMicro(), 10))
}

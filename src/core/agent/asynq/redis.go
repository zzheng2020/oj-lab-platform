package asynqAgent

import "github.com/OJ-lab/oj-lab-services/src/core"

const (
	redisHostProp = "redis.host"
)

var redisHost string

func init() {
	redisHost = core.AppConfig.GetString(redisHostProp)
	if redisHost == "" {
		panic("redis addr is not set")
	}
}
package db

import (
	"fmt"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"short_url/base/config"
	"short_url/base/tool"
	"time"
)

func initRedis() {
	redisDb=redis.NewClient(
		&redis.Options{
			Addr:         fmt.Sprintf("%s:%s", config.GetRedisConfig().GetIP(), config.GetRedisConfig().GetPort()),
			DialTimeout:  10 * time.Second, //超时时间
			ReadTimeout:  30 * time.Second,//都去邪恶如实现超时错误
			WriteTimeout: 30 * time.Second,
			Password:     config.GetRedisConfig().GetPass(),
			PoolSize:     config.GetRedisConfig().GetMaxOpen(), //最大连接数
		},
		)
	err =redisDb.Ping().Err()
	if nil!= err{
		tool.GetLogger().Error("ping redis err:", zap.Error(err))
		panic(err)
	}
	tool.GetLogger().Debug("redis : " + fmt.Sprintf("%s:%s", config.GetRedisConfig().GetIP(), config.GetRedisConfig().GetPort()))

}


func closeRedis() {
	if redisDb != nil {
		_ = redisDb.Close()
	}
}
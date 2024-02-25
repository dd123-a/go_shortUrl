package short_url

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/go-xorm/xorm"
	"short_url/base/db"
	"short_url/model/sequence"
	"sync"
)

type service struct {
	m        *xorm.Engine      // 数据库引擎实例
	r        *redis.Client     // Redis 客户端实例
	sequence sequence.Service // 序列号服务实例
}

type Service interface {
	CreateLinks(url, key string, status int) (string, error) // 创建短链接
	GetLinksByUrl(url string) (*Links, error)                // 根据 URL 获取短链接信息
	GetLinksByKeyword(keyword string) (*Links, error)        // 根据关键字获取短链接信息
}

var (
	s *service //全局服务实例
	m sync.Once
)

// Init 初始化短链接服务，获取数据库和 Redis 连接，并初始化序列号服务
func Init() {
	var err error
	m.Do(func() {
		s = &service{}                              // 创建服务实例
		s.m = db.GetMySqlDb()                      // 获取 MySQL 数据库连接
		s.r = db.GetRedisDb()                      // 获取 Redis 连接
		s.sequence, err = sequence.GetServer()     // 获取序列号服务实例
		if err != nil {                            // 如果获取序列号服务失败，则抛出错误
			panic(err)
		}
	})
}

// GetServer 返回短链接服务实例
func GetServer() (Service, error) {
	if s == nil { // 如果服务实例未初始化，则返回错误
		return nil, fmt.Errorf("[short_url] no init")
	}
	return s, nil // 否则返回服务实例
}
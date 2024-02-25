package sequence

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/go-xorm/xorm"
	"short_url/base/db"
	"sync"
)

var (
	s *service
	m sync.Once
)

type service struct {
	m *xorm.Engine
	r *redis.Client
}

type Service interface {
	GetBorrowOrder()(int64,error)
}

func Init() {
	m.Do(func() {
		s=&service{}
		s.m=db.GetMySqlDb()
		s.r=db.GetRedisDb()
	})
}

//于获取序列号服务实例。
func GetServer() (Service, error) {
	if s == nil {
		return nil, fmt.Errorf("[sequence] no init")
	}
	return s, nil
}




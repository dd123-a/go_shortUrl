package sequence

import (
	"short_url/base/tool"
)

//func (s *service) GetBorrowOrder() (int64, error) {
//	var (
//		err error
//		data int64
//	)
//	// 递增 Redis 中的指定键的值
//	if err = s.r.Incr(base.REDIS_SEQUENCE_GENERATOR).Err(); err != nil {
//		// 如果递增操作失败，记录错误日志并返回错误
//		tool.GetLogger().Error("[GetBorrowOrder] s.r.Incr " + base.REDIS_SEQUENCE_GENERATOR)
//		return 0, err
//	}
//	// 获取递增后的值（借用订单号）
//	if data, err = s.r.Get(base.REDIS_SEQUENCE_GENERATOR).Int64(); err != nil {
//		// 如果获取数据失败，记录错误日志并返回错误
//		tool.GetLogger().Error("[GetBorrowOrder] s.r.Get " + base.REDIS_SEQUENCE_GENERATOR)
//		return 0, err
//	}
//
//	return data,nil
//}

func (s *service) GetBorrowOrder() (int64, error) {
	data :=tool.GetWorker().GetId()
	return data,nil
}

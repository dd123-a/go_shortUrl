package short_url

import (
	"errors"
	"go.uber.org/zap"
	"short_url/base"
	"short_url/base/tool"
)

type Links struct {
	Id         int    `json:"id"`
	Url        string `json:"url"`
	Keyword    string `json:"keyword"`
	Status     int    `json:"status"`
	CreateTime int64  `json:"create_time"`
	UpdateTime int64  `json:"update_time"`
}

//生成短链接地址
//status 1系统分配 2用户自定义
func (s *service) CreateLinks(url, key string, status int) (string, error) {
	var (
		err error
		order int64
		links   *Links
		keyword string
	)
	//判断是否存在短码
	if links,err =s.GetLinksByUrl(url);err !=nil{
		tool.GetLogger().Error("[CreateLinks] GetLinksByUrl", zap.Error(err))
		return keyword,err
	}
	if links !=nil{
		return links.Keyword,nil
	}
	// 根据状态判断是系统分配还是用户自定义关键字
	if status == base.LINKS_CREATE_SYSTEM {
		// 如果是系统分配，从序列号生成器获取一个订单号
		if order, err = s.sequence.GetBorrowOrder(); err != nil {
			// 如果获取订单号时发生错误，记录错误日志并返回错误
			tool.GetLogger().Error("[CreateLinks] GetBorrowOrder", zap.Error(err))
			return keyword, err
		}
		// 将订单号转换为 62 进制，作为关键字
		keyword = tool.TenToAny(int(order), 62)
	} else if status == base.LINKS_CREATE_CUSTOM {
		// 如果是用户自定义，检查该关键字是否已存在
		if links, err = s.GetLinksByKeyword(key); err != nil {
			// 如果查询过程中发生错误，记录错误日志并返回错误
			tool.GetLogger().Error("[CreateLinks] GetLinksByUrl", zap.Error(err))
			return keyword, err
		}
		// 如果关键字已存在，直接返回错误
		if links != nil {
			return keyword, errors.New("already exists")
		}
		keyword = key // 使用用户提供的关键字
	} else {
		// 如果状态未定义，记录警告日志并返回错误
		tool.GetLogger().Warn("[GetLinksByUrl] undefined", zap.Any("status", status))
		return keyword, errors.New("undefined")
	}
	links=new(Links)
	links.Status=status  //设置链接状态
	links.Keyword=keyword //设置链接关键字
	links.CreateTime = tool.GetTime() // 设置创建时间
	links.UpdateTime = tool.GetTime() // 设置更新时间
	links.Url = url                 // 设置原始 URL

	if _, err = s.m.InsertOne(links); err != nil {
		// 如果插入数据库时发生错误，记录错误日志并返回错误
		tool.GetLogger().Warn("[GetLinksByUrl] undefined", zap.Any("status", status))
		return keyword, err
	}
	return keyword,nil
}

//查找url直接
func (s *service) GetLinksByUrl(url string) (*Links, error) {
	var (
		err   error
		has   bool
		links = &Links{}
	)
	has, err = s.m.Where("url = ?", url).Get(links)
	if err != nil {
		tool.GetLogger().Error("[GetLinksByUrl] Where", zap.Error(err))
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return links, nil
}

func (s *service) GetLinksByKeyword(keyword string) (*Links, error) {
	var (
		err   error
		has   bool
		links = &Links{}
	)
	has, err = s.m.Where("keyword = ?", keyword).Get(links)
	if err != nil {
		tool.GetLogger().Error("[GetLinksByKeyword] Where", zap.Error(err))
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return links, nil
}


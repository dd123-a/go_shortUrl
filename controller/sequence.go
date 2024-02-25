package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"short_url/base"
	"strconv"
	"strings"
)

// CreateLink 处理创建短链接的请求
func (b *BaseController) CreateLink(g *gin.Context) {
	var (
		url     string // 存储用户提交的原始 URL
		err     error  // 存储可能发生的错误
		order   string // 存储生成的短链接关键字
		status  string // 存储用户提交的链接状态
		keyword string // 存储用户指定的自定义关键字
	)
	// 获取用户提交的原始 URL、链接状态和自定义关键字
	url = g.PostForm("url")
	status = g.PostForm("status")
	keyword = g.PostForm("keyword")
	// 如果原始 URL 不包含 "http" 前缀，则添加上前缀
	if !strings.Contains(url, "http") {
		url = "http://" + url
	}
	// 将字符串类型的链接状态转换为整数类型
	statusInt, err := strconv.Atoi(status)
	if err != nil {
		// 如果转换失败，返回错误响应
		b.ResponseDataFailure(g, err.Error())
		return
	}
	// 调用服务层方法生成短链接，并获取生成的关键字
	if order, err = b.shortUrl.CreateLinks(url, keyword, statusInt); err != nil {
		// 如果生成短链接时发生错误，返回错误响应
		b.ResponseDataFailure(g, err.Error())
		return
	}
	// 构建完整的短链接地址，并返回成功响应
	b.ResponseData(g, fmt.Sprintf("http://%s/%s", base.ServerUrl, order))
}

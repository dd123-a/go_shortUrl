package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"short_url/controller"
	"short_url/model/short_url"
	"strings"
)

//设置路由
func Route(r *gin.Engine) {
	// 使用跨域中间件处理跨域请求
	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods: []string{"*"},
		AllowHeaders:  []string{"Origin", "Authorization", "Access-Control-Allow-Origin"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
	}))

	//获取基础控制器对象
	controller,err :=controller.GetBaseController()
	if err!=nil{
		panic(err)
	}
	//注册全局中间件
	r.Use(Middleware)
	// 创建路由组 "/v1"，用于处理 "/v1" 开头的请求
	v1:=r.Group("/v1")
	v1.POST("/create_url",controller.CreateLink)
}

func Middleware(c *gin.Context) {
	//获取请求的URL路径
	urlPath :=c.Request.URL.Path
	if strings.Contains(urlPath,"create_url"){
		c.Next()
		return
	}else{
		//获取短链接信息
		if tool,err :=short_url.GetServer();err ==nil{
			if links, err := tool.GetLinksByKeyword(strings.ReplaceAll(urlPath, "/", "")); err == nil {
				// 如果存在对应的短链接信息，则重定向到原始 URL 地址
				if links != nil {
					c.Redirect(301, links.Url)
					return
				}
			}
		}
		// 如果获取失败或者数据库中不存在对应的短链接，则返回 404 Not Found 状态码
		c.Status(404)
		c.Abort()
	}
}
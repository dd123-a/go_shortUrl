package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"short_url/base"
	"short_url/base/tool"
	"time"
)

func Run(addr string) {
	base.ServerUrl=addr
	gin.SetMode(gin.ReleaseMode)
	g := gin.Default()
	Route(g)
	s := &http.Server{
		Handler:        g, //请求处理函数
		Addr:           addr, //请求地址
		WriteTimeout:   10 * time.Second,//超时时间
		ReadTimeout:    10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	tool.GetLogger().Info("server start success : " + addr)
	err :=s.ListenAndServe()
	if err !=nil{
		panic(err)
	}
}
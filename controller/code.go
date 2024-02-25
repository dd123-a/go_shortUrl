package controller

import "github.com/gin-gonic/gin"

func (c *BaseController) ResponseData(g *gin.Context, data interface{}) {
	g.JSON(200, gin.H{
		"code": 200,
		"data": data,
		"msg":  "success",
	})
}

func (c *BaseController) ResponseDataFailure(g *gin.Context, data interface{}) {
	g.JSON(200, gin.H{
		"code": 200,
		"data": data,
		"msg":  "failure",
	})
}
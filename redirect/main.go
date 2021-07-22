//重定向
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//站外转发
	r.GET("/direct", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")
	})
	//站内转发

	r.GET("/a", func(c *gin.Context) {
		c.Request.URL.Path = "/b" //这里是把请求的URL参数改了
		r.HandleContext(c)        //继续后续的处理
	})
	r.GET("/b", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "you are great man",
		})
	})

	r.Run(":9090")
}

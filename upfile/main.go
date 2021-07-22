//文件上传
package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./index.html")

	r.GET("/select", func(c *gin.Context) {

		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.POST("/fileup", func(c *gin.Context) {
		//从请求中读取文件
		f, err := c.FormFile("f1")
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"state": err.Error(),
			})
		} else {
			dst := fmt.Sprintf("./%s", f.Filename) //获取存取文件地址
			fmt.Println("地址：", dst)
			//dst:=path.Join("./upfile/%s",f.Filename)
			//将读取到的数据保存在本地
			err := c.SaveUploadedFile(f, dst)
			if err != nil {
				fmt.Println("mistake")
			}
			fmt.Println("地址：", dst)
			c.JSON(http.StatusOK, gin.H{
				"status": "OK",
			})
		}

	})

	r.Run(":9090")
}

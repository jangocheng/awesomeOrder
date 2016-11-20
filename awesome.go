package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK,"hello gin")
	})
	rgv1 := r.Group("v1")
	{
		rgv1.GET("/order", func(c *gin.Context) {
			c.JSON(http.StatusOK,map[string]interface{}{
				"oid":time.Now().Format(time.RFC3339),
			})
		})
		rgv1.GET("/alipay", func(c *gin.Context) {
			c.String(http.StatusOK,"alipay")
		})
	}

	r.Run(":8080")
}

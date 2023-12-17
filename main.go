package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		user_ip := c.ClientIP()
		c.String(http.StatusOK, "%s", user_ip)
	})

	r.Run()
}

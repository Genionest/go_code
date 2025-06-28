package main

import (
	"bytes"
	"io"
	"log"

	"github.com/gin-gonic/gin"
)

// 自定义中间件：记录请求详细信息
func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 记录请求基本信息
		log.Printf("Request Method: %s", c.Request.Method)
		log.Printf("Request URL: %s", c.Request.URL.String())
		// 获取ip:port
		log.Printf("Host Header: %s", c.Request.Host)
		log.Printf("Remote Addr: %s", c.Request.RemoteAddr)

		log.Printf("Client IP: %s", c.ClientIP())

		// 记录请求头
		headers := c.Request.Header
		log.Println("Request Headers:")
		for key, values := range headers {
			log.Printf("  %s: %s", key, values)
		}

		// 记录查询参数
		queryParams := c.Request.URL.Query()
		log.Println("Query Parameters:")
		for key, values := range queryParams {
			log.Printf("  %s: %s", key, values)
		}

		// 记录请求体（仅限非 GET 请求）
		if c.Request.Method != "GET" && c.Request.Body != nil {
			bodyBytes, _ := io.ReadAll(c.Request.Body)
			// 将 body 重新写回请求（后续处理需要）
			c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
			log.Printf("Request Body: %s", string(bodyBytes))
		}

		// 继续处理请求
		c.Next()
	}
}

func main() {
	r := gin.Default()

	// 注册自定义中间件
	r.Use(RequestLogger())

	r.POST("/api/data", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "OK"})
	})

	r.Run(":3000")
}

package middleware

import "github.com/gin-gonic/gin"

// 设置应用名称和应用版本号
func AppInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("app_name", "blog-service")
		c.Set("app_version", "1.0.0")
		c.Next()
	}
}

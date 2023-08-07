package cors

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

//允许跨域

func Cors() gin.HandlerFunc {
	handlerFunc := cors.New(cors.Config{
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"Authentication"}, //此处设置非默认之外的请求头(自定义请求头),否则会出现跨域问题
		AllowAllOrigins:  true,
		AllowCredentials: true,
		MaxAge:           24 * time.Hour,
	})
	return handlerFunc
}
func CorsV2() gin.HandlerFunc {
	handlerFunc := cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Authentication"}, //此处设置非默认之外的请求头(自定义请求头),否则会出现跨域问题
		AllowAllOrigins:  true,
		AllowCredentials: true,
		MaxAge:           24 * time.Hour,
	})
	return handlerFunc
}

// 处理跨域请求,支持options访问
func ECors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, PATCH")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}

// 处理跨域请求,支持options访问
func CorsPlus() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "*")
		c.Header("Access-Control-Allow-Methods", "*")
		c.Header("Access-Control-Expose-Headers", "*")
		c.Header("Access-Control-Allow-Credentials", "true")

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}

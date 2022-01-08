package cors

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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

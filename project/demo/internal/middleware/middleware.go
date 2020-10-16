package middleware

import (
	"demo/global"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

//func GinLogger() gin.HandlerFunc {
//	file, err := os.OpenFile("/Users/linmingliang/workspace/go/project/demo/storage/logger/blog.logger", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
//	if err != nil {
//		//TODO 错误处理
//	}
//	loggerConfig:=gin.LoggerConfig{
//		Output: file,
//	}
//	return gin.LoggerWithConfig(loggerConfig)
//}

func BeforeRequest() gin.HandlerFunc {
	startTime := time.Now()
	return func(context *gin.Context) {
		path := context.Request.URL.Path
		raw := context.Request.URL.RawQuery
		if raw != "" {
			path = path + "?" + raw
		}

		s := fmt.Sprintf("[%9s]| %3v | %s | %s | %13v |%d \n",
			"GIN",
			startTime.Format("2006/01/02 - 15:04:05"),
			context.ClientIP(),
			context.Request.Method,
			path,
			context.Writer.Status(),
		)
		global.Logger.Print(s)
		context.Next()
	}
}

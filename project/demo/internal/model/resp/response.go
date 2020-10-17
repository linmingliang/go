package resp

import (
	"demo/pkg/err"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Success(c *gin.Context, v ...interface{}) {
	var data interface{}
	if len(v) > 0 {
		data = v[0]
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": data,
	})
	return
}
func Error(c *gin.Context, err *err.Error) {
	c.JSON(err.Code(), gin.H{
		"code": err.Code(),
		"data": err.Msg(),
	})
	return
}

package middles

import (
	"apigin/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !utils.Checklogic(c) {
			noAuth(c, "Please login first")
			return
		}

		c.Next()
		return
	}
}

//验证失败返回信息
func noAuth(c *gin.Context, msg string) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"error": msg,
	})
	c.Abort()
}

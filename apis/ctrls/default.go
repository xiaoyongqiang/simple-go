package ctrls

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexApi(c *gin.Context) {
	//utils.PreSession(c, "1")
	//utils.Clear(c, "1")
	c.HTML(http.StatusOK, "default.html", gin.H{
		"title": "测试gin界面",
	})
}

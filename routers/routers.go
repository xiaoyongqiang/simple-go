package routers

import (
	"net/http"

	"apigin/apis/ctrls"
	"apigin/apis/middles"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	router := gin.Default()

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))
	// router.MaxMultipartMemory = 8 << 20  // 8 MiB
	//router.Static("/favicon.ico", "/favicon.ico")
	router.StaticFS("/public", http.Dir("public"))
	router.LoadHTMLGlob("apis/views/*")
	router.GET("/", ctrls.IndexApi)
	v0 := router.Group("/v0")
	v0.Use(middles.Auth())
	{
		v0.GET("/member/:id", ctrls.TestDb)
		v0.GET("/member", ctrls.TestValidator)
		v0.POST("/member", ctrls.TestFile)
	}

	return router
}

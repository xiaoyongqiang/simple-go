package main

import (
	"apigin/config"
	"apigin/routers"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/koding/multiconfig"
)

func main() {
	var err error
	m := multiconfig.New()
	config.Config = new(config.CmdConfig)
	err = m.Load(config.Config)
	if err != nil {
		log.Fatalf("Load configuration failed. Error: %s\n", err.Error())
	}
	m.MustLoad(config.Config)

	err = config.InitializeConn()
	if err != nil {
		log.Fatalf("config.InitialzeConn() failed. Error info: %s\n", err.Error())
	}
	defer func() {
		config.DBHandle.Close()
		config.RedisHandle.Close()
	}()

	gin.SetMode(gin.DebugMode)
	// Logging to a file.
	// gin.DisableConsoleColor()
	// f, _ := os.Create("info.log")
	// gin.DefaultWriter = io.MultiWriter(f)

	router := routers.InitRouter()
	router.Run(fmt.Sprintf(":%d", config.Config.ApiConf.Port))
}

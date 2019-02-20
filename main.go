package main

import (
	"bitbucket.org/marktohark/nfuusrsystem/kernel/config"
	"bitbucket.org/marktohark/nfuusrsystem/kernel/router"
	"bitbucket.org/marktohark/nfuusrsystem/middleware/global"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func main() {
	// Logging to a file.
	gin.DisableConsoleColor()
	f, _ := os.Create("./log/gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	//read .env
	config.Init()

	//get engine
	r := gin.Default()

	//set view path
	r.LoadHTMLGlob("views/*")
	//set static
	static_name, _ := config.Get("STATIC_URL_NAME")
	r.Static(static_name, "static")

	//global middleware reg
	global.RegGlobalMiddleware(r)

	//RouterRegister setting
	_, err := router.New(r)
	if err != nil {
		fmt.Println(err)
	}

	port, _ := config.Get("port")
	//run
	r.Run(":" + port)
}
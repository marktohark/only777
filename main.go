package main

import (
	"bitbucket.org/marktohark/nfuusrsystem/kernel/config"
	"bitbucket.org/marktohark/nfuusrsystem/kernel/router"
	"bitbucket.org/marktohark/nfuusrsystem/middleware/global"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	//read .env
	config.Init()

	//get engine
	r := gin.Default()

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
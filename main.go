package main

import (
	"bitbucket.org/marktohark/nfuusrsystem/kernel/router"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	Rr, err := router.New(r)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(Rr)
	r.Run()
}
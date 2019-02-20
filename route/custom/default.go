package custom

import (
	"bitbucket.org/marktohark/nfuusrsystem/controller/HomeController"
	"github.com/gin-gonic/gin"
)

func Default(r *gin.Engine) error {
	r.GET("/", HomeController.Index)
	return nil
}
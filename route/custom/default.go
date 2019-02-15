package custom

import (
	"github.com/gin-gonic/gin"
)

func Default(r *gin.Engine) error {
	r.GET("test", func(c *gin.Context) {
		c.String(200, "hello")
	})
	return nil
}
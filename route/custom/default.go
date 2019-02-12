package custom

import "github.com/gin-gonic/gin"

func Default(r *gin.Engine) error {
	r.GET("/welecome", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "welecome",
		})
	})
	return nil
}
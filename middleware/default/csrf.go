package _default

import (
	"bitbucket.org/marktohark/nfuusrsystem/kernel/config"
	"github.com/gin-gonic/gin"
	"github.com/marktohark/gin-csrf"
)

func CsrfMiddleware(r *gin.Engine)  {
	secret, _ := config.Get("CSRF_SECRET")
	r.Use(csrf.Middleware(csrf.Options{
		Secret: secret,
		ErrorFunc: func(c *gin.Context) {
			c.String(400, "CSRF token mismatch")
			c.Abort()
		},
	}))
}
package global

import (
	"bitbucket.org/marktohark/nfuusrsystem/middleware/default"
	"github.com/gin-gonic/gin"
)

func RegGlobalMiddleware(r *gin.Engine) {
	/**
		r.Use or func(r)
	 */
	_default.SessionMiddleware(r)
	_default.CsrfMiddleware(r)
}
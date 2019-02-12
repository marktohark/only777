package route

import (
	"bitbucket.org/marktohark/nfuusrsystem/route/custom"
	"github.com/gin-gonic/gin"
)

func RegList() []func(*gin.Engine)(error) {
	return []func(*gin.Engine)(error) {
		custom.Default,
	}
}
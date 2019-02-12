package router

import (
	"bitbucket.org/marktohark/nfuusrsystem/route"
	"github.com/gin-gonic/gin"
)

type RouterRegister struct {
	 RList []func(*gin.Engine)(error)
	 Engine *gin.Engine
}

func New(r *gin.Engine) (*RouterRegister, error) {
	Rr := &RouterRegister{}
	if err :=  Rr.init(r); err != nil {
		return nil, err
	}
	return Rr, nil
}

func (this *RouterRegister) init(r *gin.Engine) error {
	this.Engine = r
	return this.register()
}

func(this *RouterRegister) register() error {
	this.RList = route.RegList()

	for _, f := range this.RList {
		err := f(this.Engine)
		if err != nil {
			return err
		}
	}
	return nil
}
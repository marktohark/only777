package _default

import (
	"bitbucket.org/marktohark/nfuusrsystem/kernel/config"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func SessionMiddleware(r *gin.Engine)  {
	secret, _ := config.Get("COOKIE_SECRET")
	store := cookie.NewStore([]byte(secret))
	sessionName, _ := config.Get("SESSION_NAME")
	r.Use(sessions.Sessions(sessionName, store))
}
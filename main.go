package main

import (
	"GoExample/routers"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

/**
*
* @author yth
* @language go
* @since 2022/11/16 15:44
 */

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "static")
	store := cookie.NewStore([]byte("secret111"))
	r.Use(sessions.Sessions("mysession", store))
	routers.LrInit(r)
	routers.PatientRoutersInit(r)
	routers.UserRoutersInit(r)
	r.Run()
}

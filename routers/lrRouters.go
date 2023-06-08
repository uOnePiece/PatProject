package routers

import (
	"GoExample/controllers"
	"GoExample/middlewares"
	"github.com/gin-gonic/gin"
)

/**
*
* @author yth
* @language go
* @since 2022/11/16 20:38
 */

func LrInit(r *gin.Engine) {
	userRouters := r.Group("/", middlewares.InitMiddleware)
	{
		userRouters.GET("/login", controllers.LrController{}.Login)
		userRouters.GET("/register", controllers.LrController{}.Register)
	}
}

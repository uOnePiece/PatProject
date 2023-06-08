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

func UserRoutersInit(r *gin.Engine) {
	userRouters := r.Group("/user")
	{
		userRouters.GET("/", controllers.UserController{}.UserIndex)
		userRouters.POST("/login", middlewares.InitMiddleware, controllers.UserController{}.UserLogin)
		userRouters.POST("/register", controllers.UserController{}.UserRegister)
		userRouters.POST("/regPatient", controllers.UserController{}.RegPatient)
		userRouters.GET("/NowUser", controllers.UserController{}.NowUser)
		userRouters.POST("/UpUser", controllers.UserController{}.UpUser)
		userRouters.POST("/UpPa", controllers.UserController{}.UpPa)
		userRouters.GET("/Logout", controllers.UserController{}.Logout)
		userRouters.GET("/Pr", controllers.UserController{}.Pr)
		userRouters.GET("/SelByCon", controllers.UserController{}.SelByCon)
		userRouters.GET("/SelByType", controllers.UserController{}.SelByType)
		userRouters.GET("/GetCity", controllers.UserController{}.GetCity)
		userRouters.POST("/AddArea", controllers.UserController{}.AddArea)
		userRouters.GET("/UserUpdate", controllers.UserController{}.UserUpdate)
		userRouters.POST("/UpInfo", controllers.UserController{}.UpInfo)
		userRouters.GET("/De", controllers.UserController{}.De)
	}
}

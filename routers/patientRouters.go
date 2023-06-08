package routers

import (
	"GoExample/controllers"
	"github.com/gin-gonic/gin"
)

/**
*
* @author yth
* @language go
* @since 2022/11/16 20:38
 */

func PatientRoutersInit(r *gin.Engine) {
	userRouters := r.Group("/patient")
	{
		userRouters.GET("/typeAll", controllers.PatientController{}.TypeAll)
		userRouters.POST("/SelByCon", controllers.PatientController{}.SelByCon)
		userRouters.POST("/SelByRe", controllers.PatientController{}.SelByRe)
		userRouters.POST("/SelByTy", controllers.PatientController{}.SelByTy)
		userRouters.POST("/SelByDate", controllers.PatientController{}.SelByDate)
		userRouters.GET("/SelAll", controllers.PatientController{}.SelAll)
		userRouters.GET("/SelPage", controllers.PatientController{}.SelPage)
		userRouters.GET("/SelReAll", controllers.PatientController{}.SelReAll)
		userRouters.GET("/SelReCon", controllers.PatientController{}.SelReCon)
		userRouters.POST("/SelReTy", controllers.PatientController{}.SelReTy)
		userRouters.GET("/SelPageCon", controllers.PatientController{}.SelPageCon)
		userRouters.GET("/updatePage", controllers.PatientController{}.UpdatePage)
		userRouters.POST("/UpInfo", controllers.PatientController{}.UpInfo)
		userRouters.GET("/De", controllers.PatientController{}.De)
		userRouters.POST("/AddType", controllers.PatientController{}.AddType)
	}
}

package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
*
* @author yth
* @language go
* @since 2022/11/16 21:30
 */

type LrController struct {
}

func (lr LrController) Login(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
}

func (lr LrController) Register(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", gin.H{})
}

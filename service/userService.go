package service

import (
	"GoExample/dao"
	"GoExample/models"
	"github.com/gin-gonic/gin"
	"path"
)

/**
*
* @author yth
* @language go
* @since 2022/11/16 20:57
 */

type UserService struct {
}

func (user UserService) UserLoginService(c *gin.Context) models.User {
	username := c.PostForm("username")
	password := c.PostForm("password")
	return dao.UserDao{}.UserLoginDao(username, password)

}

func (user UserService) UserRegisterService(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	typeN := c.PostForm("typeN")
	sex := c.PostForm("sex")
	identity := c.PostForm("identity")
	address := c.PostForm("address")
	phone := c.PostForm("phone")
	img, err := c.FormFile("headImg")
	dst := path.Join("GoExample/static/upload", img.Filename)
	if err == nil {
		c.SaveUploadedFile(img, dst)
	}
	dao.UserDao{}.UserRegisterDao(username, password, typeN, sex, identity, address, phone, img.Filename)
}

func (user UserService) GetTypeN(userId int) models.UserRoleConn {
	urConn := dao.UserDao{}.GetTypeN(userId)
	return urConn
}

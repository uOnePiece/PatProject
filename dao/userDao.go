package dao

import (
	"GoExample/models"
	"gorm.io/gorm"
	"time"
)

/**
*
* @author yth
* @language go
* @since 2022/11/16 20:58
 */

type UserDao struct {
}

func (user UserDao) UserLoginDao(username, password string) models.User {
	var u = models.User{}
	models.DB.Where("username = ? and password = ?", username, password).Find(&u)
	return u
}

func (user UserDao) GetTypeN(userId int) models.UserRoleConn {
	var urConn = models.UserRoleConn{}
	models.DB.Where("user_id = ?", userId).Find(&urConn)
	return urConn
}

func (user UserDao) UserRegisterDao(username, password, typeN, sex, identity, address, phone, img string) {
	u := models.User{
		Username:     username,
		Password:     password,
		Identity:     identity,
		Gender:       sex,
		Address:      address,
		RegisterTime: models.FormatTime(time.Now()),
		Phone:        phone,
		Img:          img,
	}
	models.DB.Transaction(func(tx *gorm.DB) error {
		//往用户表中插入数据
		if err := tx.Create(&u).Error; err != nil {
			return err
		}
		t := user.UserLoginDao(username, password)
		ur := models.UserRoleConn{
			UserId: t.Id,
		}
		if typeN == "1" {
			ur.RoleId = 1
		} else {
			ur.RoleId = 2
		}
		//往用户角色连接表中插入数据
		if err := tx.Create(&ur).Error; err != nil {
			return err
		}
		return nil
	})

}

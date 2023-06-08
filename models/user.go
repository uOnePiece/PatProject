package models

/**
*
* @author yth
* @language go
* @since 2022/11/16 16:18
 */

type User struct {
	Id           int
	Username     string
	Password     string
	Identity     string
	Gender       string
	Address      string
	RegisterTime string
	Phone        string
	Img          string
	IsDelete     int
}

func (User) TableName() string {
	return "user_info"
}

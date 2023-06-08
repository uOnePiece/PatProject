package models

/**
*
* @author yth
* @language go
* @since 2022/11/16 16:32
 */

type Role struct {
	Id       int
	RoleName string
}

func (Role) TableName() string {
	return "role_info"
}

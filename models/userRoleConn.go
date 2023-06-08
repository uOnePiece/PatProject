package models

/**
*
* @author yth
* @language go
* @since 2022/11/12 14:42
 */

type UserRoleConn struct {
	Id     int
	UserId int
	RoleId int
}

func (UserRoleConn) TableName() string {
	return "user_role_conn"
}

package models

/**
*
* @author yth
* @language go
* @since 2022/11/17 16:17
 */

type TypeInfo struct {
	Id       int
	TypeName string
}

func (TypeInfo) TableName() string {
	return "type_info"
}

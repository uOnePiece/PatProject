package models

/**
*
* @author yth
* @language go
* @since 2022/11/19 10:06
 */

type AreaInfo struct {
	Id       int
	Province string
	City     string
}

func (AreaInfo) TableName() string {
	return "area"
}

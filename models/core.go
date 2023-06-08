package models

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/**
*
* @author yth
* @language go
* @since 2022/11/12 14:33
 */

var DB *gorm.DB
var err error

func init() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/go_example?charset=utf8&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
}

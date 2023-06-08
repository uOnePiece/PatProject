package models

import (
	"gorm.io/gorm"
	"time"
)

/**
*
* @author yth
* @language go
* @since 2022/11/17 22:43
 */

func GetNowTime() time.Time {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	bDate, _ := time.ParseInLocation("2006-01-02 15:04:05", time.Now().Format("2006-01-02 15:04:05"), loc)
	return bDate
}

func FormatTime(t time.Time) string {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	bDate, _ := time.ParseInLocation("2006-01-02 15:04:05", t.Format("2006-01-02 15:04:05"), loc)
	return bDate.Format("2006-01-02 15:04:05")
}

// Paginate 分页封装
func Paginate(page int, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

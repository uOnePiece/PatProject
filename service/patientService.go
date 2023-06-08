package service

import (
	"GoExample/dao"
	"github.com/gin-gonic/gin"
)

/**
*
* @author yth
* @language go
* @since 2022/11/16 20:57
 */

type PatientService struct {
}

func (patient PatientService) TypeAll(c *gin.Context) {
	dao.PatientDao{}.TypeAll()
}

package controllers

import (
	"GoExample/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"strings"
	"time"
)

/**
*
* @author yth
* @language go
* @since 2022/11/17 16:24
 */

type PatientController struct {
}

func (patient PatientController) TypeAll(c *gin.Context) {
	var typeInfo []models.TypeInfo
	models.DB.Find(&typeInfo)
	c.JSON(http.StatusOK, gin.H{
		"r": typeInfo,
	})
}

// SelByCon 模糊查询病人信息
func (patient PatientController) SelByCon(c *gin.Context) {
	json := make(map[string]interface{})
	c.BindJSON(&json)
	con := json["con"].(string)
	var p []models.Patient
	err := models.DB.Where("(identity like ? or name like ? or gender like ? or type_disease like ? or phone like ?) and is_delete = 0", "%"+con+"%", "%"+con+"%", "%"+con+"%", "%"+con+"%", "%"+con+"%").Find(&p).Error
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"data": p,
	})
}

// SelByRe 通过判断是否恢复查询
func (patient PatientController) SelByRe(c *gin.Context) {
	json := make(map[string]interface{})
	c.BindJSON(&json)
	isRe := json["isRe"]
	if json["beginDate"].(string) == "1" {
		var p []models.Patient
		err := models.DB.Where("is_recover = ? and is_delete = 0", isRe).Find(&p).Error
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, gin.H{
			"data": p,
		})
	} else {
		beginDate := strings.Replace(json["beginDate"].(string), "T", " ", 1) + ":00"
		endDate := strings.Replace(json["endDate"].(string), "T", " ", 1) + ":00"
		bDate, _ := time.Parse("2006-01-02 15:04:05", beginDate)
		eDate, _ := time.Parse("2006-01-02 15:04:05", endDate)
		if bDate.After(eDate) {
			c.JSON(http.StatusOK, gin.H{
				"data": "查询条件不合理",
			})
		} else {
			var p []models.Patient
			err := models.DB.Where("begin_date > ? and end_date < ? and is_recover = ? and is_delete = 0", bDate, eDate, isRe).Find(&p).Error
			if err != nil {
				panic(err)
			}
			c.JSON(http.StatusOK, gin.H{
				"data": p,
			})
		}
	}

}

// SelByTy 通过病情类型查询
func (patient PatientController) SelByTy(c *gin.Context) {
	json := make(map[string]interface{})
	c.BindJSON(&json)
	typeInfo := json["typeInfo"].(string)
	if json["beginDate"].(string) == "1" {
		var p []models.Patient
		err := models.DB.Where("type_disease = ? and is_delete = 0", typeInfo).Find(&p).Error
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, gin.H{
			"data": p,
		})
	} else {
		beginDate := strings.Replace(json["beginDate"].(string), "T", " ", 1) + ":00"
		endDate := strings.Replace(json["endDate"].(string), "T", " ", 1) + ":00"
		bDate, _ := time.Parse("2006-01-02 15:04:05", beginDate)
		eDate, _ := time.Parse("2006-01-02 15:04:05", endDate)
		if bDate.After(eDate) {
			c.JSON(http.StatusOK, gin.H{
				"data": "查询条件不合理",
			})
		} else {
			var p []models.Patient
			err := models.DB.Where("begin_date > ? and end_date < ? and type_disease = ? and is_delete = 0", bDate, eDate, typeInfo).Find(&p).Error
			if err != nil {
				panic(err)
			}
			c.JSON(http.StatusOK, gin.H{
				"data": p,
			})
		}
	}
}

// SelByDate 通过日期查询
func (patient PatientController) SelByDate(c *gin.Context) {
	json := make(map[string]interface{})
	c.BindJSON(&json)
	isRecover := json["isRe"]
	isTy := json["isTy"]
	var isRe int
	if isRecover == "已恢复" {
		isRe = 1
	} else {
		isRe = 1
	}
	beginDate := strings.Replace(json["beginDate"].(string), "T", " ", 1) + ":00"
	endDate := strings.Replace(json["endDate"].(string), "T", " ", 1) + ":00"
	bDate, _ := time.Parse("2006-01-02 15:04:05", beginDate)
	eDate, _ := time.Parse("2006-01-02 15:04:05", endDate)
	if bDate.After(eDate) {
		c.JSON(http.StatusOK, gin.H{
			"data": "查询条件不合理",
		})
	} else {
		var p []models.Patient
		err := models.DB.Where("begin_date > ? and end_date < ? and type_disease = ? and is_recover = ? and is_delete = 0", beginDate, endDate, isTy, isRe).Find(&p).Error
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, gin.H{
			"data": p,
		})
	}
}

// SelAll 查询所有病人
func (patient PatientController) SelAll(c *gin.Context) {
	var p []models.Patient
	err := models.DB.Where("is_delete = 0").Find(&p).Error
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"data": p,
	})
}

// SelPage 分页查询病人信息
func (patient PatientController) SelPage(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	var p []models.Patient
	err := models.DB.Where("is_delete = 0").Scopes(models.Paginate(page, pageSize)).Find(&p).Error
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"data": p,
	})
}

// SelPageCon 分页查询病人信息
func (patient PatientController) SelPageCon(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	var p []models.Patient
	err := models.DB.Scopes(models.Paginate(page, pageSize)).Where("is_recover = 1 and is_delete = 0").Find(&p).Error
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"data": p,
	})
}

// SelReAll 查询所有痊愈病人
func (patient PatientController) SelReAll(c *gin.Context) {
	var p []models.Patient
	err := models.DB.Where("is_recover = ? and is_delete = 0", 1).Find(&p).Error
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"data": p,
	})
}

// SelReCon 模糊查询所有痊愈病人
func (patient PatientController) SelReCon(c *gin.Context) {
	con := c.Query("v")
	var p []models.Patient
	err := models.DB.Where("(name like ? or gender like ? or type_disease like ? or phone like ?) and is_recover = 1 and is_delete = 0", "%"+con+"%", "%"+con+"%", "%"+con+"%", "%"+con+"%").Find(&p).Error
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"data": p,
	})
}

// SelReTy 根据类型查询痊愈病人
func (patient PatientController) SelReTy(c *gin.Context) {
	json := make(map[string]interface{})
	c.BindJSON(&json)
	t := json["t"]
	var p []models.Patient
	err := models.DB.Where("type_disease = ? and is_recover = 1 and is_delete = 0", t).Find(&p).Error
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"data": p,
	})
}

// UpdatePage 进入修改界面
func (patient PatientController) UpdatePage(c *gin.Context) {
	c.HTML(http.StatusOK, "patient_update.html", gin.H{})
}

// UpInfo 病人信息修改
func (patient PatientController) UpInfo(c *gin.Context) {
	name := c.PostForm("uName")
	isId, _ := strconv.Atoi(c.PostForm("isId"))
	uIdentity := c.PostForm("uIdentity")
	uPhone := c.PostForm("uPhone")
	uGender := c.PostForm("uGender")
	uType := c.PostForm("uType")
	isRe := c.PostForm("isRe")
	u := models.Patient{Id: isId}
	models.DB.Find(&u)
	u.Name = name
	u.Identity = uIdentity
	u.Phone = uPhone
	u.Gender = uGender
	u.TypeDisease = uType
	u.IsRecover, _ = strconv.Atoi(isRe)
	if u.IsRecover == 1 {
		u.EndDate = models.FormatTime(time.Now())
	}
	models.DB.Save(&u)
	c.HTML(http.StatusOK, "user_main.html", gin.H{})
}

// De 删除病人信息
func (patient PatientController) De(c *gin.Context) {
	pId, _ := strconv.Atoi(c.Query("pId"))
	u := models.Patient{Id: pId}
	models.DB.Find(&u)
	u.IsDelete = 1
	models.DB.Save(&u)
}

// AddType 删除病人信息
func (patient PatientController) AddType(c *gin.Context) {
	fType := c.PostForm("f_type")
	var newType models.TypeInfo
	newType.TypeName = fType
	models.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&newType).Error; err != nil {
			return err
		}
		return nil
	})
	c.HTML(http.StatusOK, "user_main.html", gin.H{})
}

package controllers

import (
	"GoExample/models"
	"GoExample/service"
	"encoding/gob"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"time"
)

/**
*
* @author yth
* @language go
* @since 2022/11/16 20:40
 */

type UserController struct {
}

func (user UserController) UserIndex(c *gin.Context) {

}

func (user UserController) UserLogin(c *gin.Context) {

	u := service.UserService{}.UserLoginService(c)
	urConn := service.UserService{}.GetTypeN(u.Id)
	gob.Register(models.User{})
	if u != (models.User{}) {
		session := sessions.Default(c)
		session.Set("user", u)
		err := session.Save()
		if err != nil {
			panic(err)
		}
		if urConn.RoleId == 2 {
			c.HTML(http.StatusOK, "admin_main.html", gin.H{
				"c":    session.Get("c"),
				"user": u,
			})
		} else {
			c.HTML(http.StatusOK, "user_main.html", gin.H{
				"user": u,
			})
		}
	} else {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"msg": "账号或密码有误",
		})
	}
}

func (user UserController) UserRegister(c *gin.Context) {
	service.UserService{}.UserRegisterService(c)
	c.HTML(http.StatusOK, "login.html", gin.H{
		"msg": "注册成功，请登录~",
	})
}

// RegPatient 登记病人
func (user UserController) RegPatient(c *gin.Context) {
	gob.Register(models.User{})
	session := sessions.Default(c)
	a := session.Get("user").(models.User)
	p := models.Patient{
		Name:        c.PostForm("pName"),
		Identity:    c.PostForm("pIdentity"),
		Person:      a.Username,
		Gender:      c.PostForm("pGender"),
		TypeDisease: c.PostForm("pType"),
		Introduce:   c.PostForm("pIntroduce"),
		Phone:       c.PostForm("pPhone"),
		BeginDate:   models.FormatTime(time.Now()),
	}
	//向病人表中插入数据
	models.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&p).Error; err != nil {
			return err
		}
		t := models.TypeInfo{}
		if err := models.DB.Where("type_name = ?", c.PostForm("pType")).Find(&t).Error; err != nil {
			return err
		}
		pc := models.TyPatConn{
			TypeId:    t.Id,
			PatientId: p.Id,
		}
		//向病人与类型连接表中插入数据
		if err := tx.Create(&pc).Error; err != nil {
			return err
		}
		return nil
	})
}

// NowUser 获取当前用户信息
func (user UserController) NowUser(c *gin.Context) {
	gob.Register(models.User{})
	session := sessions.Default(c)
	a := session.Get("user").(models.User)
	c.JSON(http.StatusOK, gin.H{
		"name":     a.Username,
		"phone":    a.Phone,
		"address":  a.Address,
		"password": a.Password,
	})
}

// UpUser 获取当前用户信息
func (user UserController) UpUser(c *gin.Context) {

	gob.Register(models.User{})
	session := sessions.Default(c)
	a := session.Get("user").(models.User)
	u := models.User{Id: a.Id}
	models.DB.Find(&u)
	u.Username = c.PostForm("dName")
	u.Identity = c.PostForm("dIdentity")
	u.Address = c.PostForm("dAddress")
	models.DB.Save(&u)
	c.HTML(http.StatusOK, "user_main.html", gin.H{})
}

// UpPa 获取当前用户密码
func (user UserController) UpPa(c *gin.Context) {
	json := make(map[string]interface{})
	c.BindJSON(&json)
	gob.Register(models.User{})
	session := sessions.Default(c)
	a := session.Get("user").(models.User)
	u := models.User{Id: a.Id}
	models.DB.Find(&u)
	u.Password = json["newP"].(string)
	models.DB.Save(&u)
	c.JSON(http.StatusOK, gin.H{
		"data": "修改成功",
	})
}

// Logout 获取当前用户密码
func (user UserController) Logout(c *gin.Context) {
	gob.Register(models.User{})
	session := sessions.Default(c)
	session.Delete("user")
	c.JSON(http.StatusOK, gin.H{
		"data": "退出成功",
	})
}

// Pr 回显省、市
func (user UserController) Pr(c *gin.Context) {
	var area []models.AreaInfo
	models.DB.Find(&area)
	c.JSON(http.StatusOK, gin.H{
		"data": area,
	})
}

// SelByCon 根据文本查询医院
func (user UserController) SelByCon(c *gin.Context) {
	con := c.Query("con")
	var u []models.User
	if err := models.DB.Where("(username like ? or gender like ? or identity like ? or phone like ? or address like ?) and is_delete = 0", "%"+con+"%", "%"+con+"%", "%"+con+"%", "%"+con+"%", "%"+con+"%").Find(&u).Error; err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"data": u,
	})
}

// SelByType 根据省、区查询医院
func (user UserController) SelByType(c *gin.Context) {
	t1 := c.Query("t1")
	t2 := c.Query("t2")
	var u []models.User
	if err := models.DB.Where("address like ?", "%"+t1+t2+"%").Find(&u).Error; err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"data": u,
	})
}

// GetCity 获取省下市区
func (user UserController) GetCity(c *gin.Context) {
	p := c.Query("p")
	var a []models.AreaInfo
	models.DB.Where("province = ?", p).Find(&a)
	c.JSON(http.StatusOK, gin.H{
		"data": a,
	})
}

// AddArea 增加区域信息（省、市）
func (user UserController) AddArea(c *gin.Context) {
	json := make(map[string]interface{})
	c.BindJSON(&json)
	p := json["p"].(string)
	a := json["a"].(string)
	var newArea = models.AreaInfo{
		Province: p,
		City:     a,
	}
	models.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&newArea).Error; err != nil {
			return err
		}
		return nil
	})
	c.JSON(http.StatusOK, gin.H{
		"data": "添加成功",
	})
}

// UserUpdate 进入用户修改界面（回显数据）
func (user UserController) UserUpdate(c *gin.Context) {
	c.HTML(http.StatusOK, "user_update.html", gin.H{})
}

// UpInfo 更改医院（用户）信息
func (user UserController) UpInfo(c *gin.Context) {
	name := c.PostForm("uName")
	identity := c.PostForm("uIdentity")
	phone := c.PostForm("uPhone")
	address := c.PostForm("uAddress")
	gender := c.PostForm("uGender")
	uId, _ := strconv.Atoi(c.PostForm("uId"))
	models.DB.Transaction(func(tx *gorm.DB) error {
		u := models.User{Id: uId}
		if err := tx.Find(&u).Error; err != nil {
			return err
		}
		u.Username = name
		u.Identity = identity
		u.Phone = phone
		u.Address = address
		u.Gender = gender
		if err := tx.Save(&u).Error; err != nil {
			return err
		}
		return nil
	})
	c.HTML(http.StatusOK, "admin_main.html", gin.H{})
}

// De 删除用户信息
func (user UserController) De(c *gin.Context) {
	uId, _ := strconv.Atoi(c.Query("userId"))
	models.DB.Transaction(func(tx *gorm.DB) error {
		u := models.User{Id: uId}
		if err := tx.Find(&u).Error; err != nil {
			return err
		}
		u.IsDelete = 1
		if err := tx.Save(&u).Error; err != nil {
			return err
		}
		return nil
	})
	c.JSON(http.StatusOK, gin.H{
		"data": "删除成功",
	})

}

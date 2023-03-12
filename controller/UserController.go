package controller

import (
	"ginDemo/common"
	"ginDemo/dto"
	"ginDemo/model"
	"ginDemo/response"
	"ginDemo/util"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

func Register(c *gin.Context) {
	DB := common.GetDB()
	//获取参数
	name := c.PostForm("name")
	telephone := c.PostForm("telephone")
	password := c.PostForm("password")
	//数据验证
	if len(telephone) != 11 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位", false)
		return
	}

	if len(password) < 6 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "密码不能少于6位", false)
		return
	}

	if len(name) == 0 {
		name = util.RandomString(10)
	}
	//判断手机号是否存在
	if isTelephoneExist(DB, telephone) {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "用户已存在", false)
		return
	}
	//创建用户
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(c, http.StatusInternalServerError, 500, nil, "加密失败！", false)
		return
	}
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hashedPassword),
	}

	DB.Create(&newUser)

	response.Success(c, nil, "注册成功")
	log.Println(name, telephone, password)
}

func Login(c *gin.Context) {
	DB := common.GetDB()
	//获取参数
	// name := c.PostForm("name")
	telephone := c.PostForm("telephone")
	password := c.PostForm("password")
	//数据验证
	if len(telephone) != 11 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位", false)
		return
	}

	if len(password) < 6 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "密码不能少于6位", false)
		return
	}

	//验证手机号是否存在
	var user model.User
	DB.Where("telephone = ?", telephone).First(&user)
	if user.ID == 0 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "用户不存在", false)
		return
	}

	//判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		response.Response(c, 400, 400, nil, "密码错误", false)
		return
	}

	//发放token
	token, err := common.ReleaseToken(user)
	if err != nil {
		response.Response(c, http.StatusInternalServerError, 500, nil, "系统异常", false)
		log.Printf("token generate error: %v", err)
		return
	}

	//返回结果
	response.Success(c, gin.H{"token": token}, "登陆成功")
}

func Info(c *gin.Context) {
	user, _ := c.Get("user")

	c.JSON(200, gin.H{"code": 200, "data": gin.H{"user": dto.ToUserDto(user.(model.User))}})
}

func ReName(c *gin.Context) {
	DB := common.GetDB()
	user, _ := c.Get("user")
	newName := c.PostForm("newName")
	user1 := user.(model.User)

	DB.Model(&model.User{}).Where("id = ?", user1.ID).Update("name", newName)

	response.Success(c, gin.H{"newName": newName}, "修改成功")
}

func RePassword(c *gin.Context) {
	DB := common.GetDB()
	user, _ := c.Get("user")
	newPwd := c.PostForm("newPwd")
	if len(newPwd) <= 6 {
		response.Fail(c, gin.H{}, "密码长度过短")
		return
	}

	user1 := user.(model.User)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPwd), bcrypt.DefaultCost)
	if err != nil {
		response.Fail(c, gin.H{}, "加密失败")
		return
	}

	DB.Model(&model.User{}).Where("id = ?", user1.ID).Update("password", string(hashedPassword))
	response.Success(c, gin.H{}, "修改成功")
}

func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}

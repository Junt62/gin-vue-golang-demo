package controller

import (
	"OceanLearn/common"
	"OceanLearn/dto"
	"OceanLearn/model"
	"OceanLearn/response"
	"OceanLearn/util"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"regexp"
)

func Register(ctx *gin.Context) {

	// 数据库对象
	DB := common.GetDB()

	// 使用结构体来解析body内的参数
	var requestUser = model.User{}
	err := ctx.Bind(&requestUser)

	name := requestUser.Name
	telephone := requestUser.Telephone
	password := requestUser.Password

	// 数据验证
	telRE := regexp.MustCompile("^1\\d{10}$")
	if !telRE.MatchString(telephone) {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号必须为1开头的11位数字")
		return
	}
	passRE := regexp.MustCompile("^[A-Za-z0-9]{6,18}$")
	if !passRE.MatchString(password) {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码必须为6-18位的字母与数字的组合")
		return
	}

	// 判断是否要生成随机姓名
	if len(name) == 0 {
		name = util.RandomString(10)
	}

	// 判断手机号是否存在
	if isTelephoneExist(DB, telephone) {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户已存在")
		return
	}

	//创建用户
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "加密出现错误")
		return
	}
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hasedPassword),
	}
	DB.Create(&newUser)

	// 发放token
	token, err := common.ReleaseToken(newUser)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "系统异常")
		//ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "系统异常"})
		log.Printf("token generate error : %v", err)
		return
	}

	// 返回结果
	response.Success(ctx, gin.H{"token": token}, "注册成功")

}

func Login(ctx *gin.Context) {

	// 数据库对象
	DB := common.GetDB()

	// 使用结构体来解析body内的参数
	var requestUser = model.User{}
	err := ctx.Bind(&requestUser)

	telephone := requestUser.Telephone
	password := requestUser.Password

	// 数据验证
	telRE := regexp.MustCompile("^1\\d{10}$")
	if !telRE.MatchString(telephone) {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号必须为1开头的11位数字")
		return
	}
	passRE := regexp.MustCompile("^[A-Za-z0-9]{6,18}$")
	if !passRE.MatchString(password) {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码必须为6-18位的字母与数字的组合")
		return
	}

	// 判断用户存在
	var user model.User
	DB.Where("telephone = ?", telephone).First(&user)
	if user.ID == 0 {
		response.Response(ctx, http.StatusUnprocessableEntity, 400, nil, "用户不存在")
		return
	}

	// 判断密码正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		response.Response(ctx, http.StatusBadRequest, 400, nil, "密码错误")
		return
	}

	// 发放token
	token, err := common.ReleaseToken(user)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "系统异常")
		log.Printf("token generate error : %v", err)
		return
	}

	// 返回结果
	response.Success(ctx, gin.H{"token": token}, "登录成功")

}

func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{"user": dto.ToUserDto(user.(model.User))}})
}

func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	return user.ID != 0
}

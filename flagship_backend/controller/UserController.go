package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"xiangerfer.com/db_flagship/common"
	"xiangerfer.com/db_flagship/dto"
	"xiangerfer.com/db_flagship/model"
	"xiangerfer.com/db_flagship/response"
	"xiangerfer.com/db_flagship/utils"
)

func Register(ctx *gin.Context) {

	DB := common.GetDB()
	// 获取参数
	name := ctx.PostForm("name")
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")

	// 数据验证
	if len(telephone) != 11{
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号不能低于11位")
	}

	if len(password) < 6{
		response.Response(ctx, http.StatusUnprocessableEntity, 423, nil, "密码不能少于16位")
		return
	}

	// 如果用户名为空，随机生成一个。
	if len(name) == 0{
		name = utils.RandomString(10)
	}

	// 判断手机号是否存在
	if IsTelephoneExist(DB, telephone){
		response.Response(ctx, http.StatusUnprocessableEntity, 425, nil, "手机号已经存在")
		return
	}

	// 创建用户
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil{
		response.Response(ctx, http.StatusUnprocessableEntity, 500, nil, "密码加密错误，联系站长")
		return
	}
	newUser := model.User{
		Name: name,
		Telephone: telephone,
		Password : string(hasedPassword)}
	DB.Create(&newUser)
	// 返回结果

	response.Success(ctx, 200, nil, "注册成功")
}


func Login(ctx *gin.Context){
	DB := common.GetDB()
	// 获取参数
	//name := ctx.PostForm("name")
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")

	// 数据验证
	if len(telephone) != 11{
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号不能低于11位")
		return
	}

	if len(password) < 6{
		response.Response(ctx, http.StatusUnprocessableEntity, 423, nil, "密码不能少于6位")
		return
	}

	// 数据验证
	var user model.User;
	DB.Where("Telephone = ?", telephone).First(&user)
	if user.ID == 0{
		response.Response(ctx, http.StatusUnprocessableEntity, 424, nil, "用户不存在")
		return
	}

	// 密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil{
		response.Response(ctx, http.StatusBadRequest, 400, nil, "密码错误")
		return
	}

	// 发放token
	token, err := common.ReleaseToken(user)
	if err != nil{
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "生成token失败")
		return
	}

	// 返回结果


	response.Success(ctx, 200, gin.H{"token": token}, "登陆成功")



}

func Info(ctx *gin.Context){

	user, _ := ctx.Get("user")

	response.Success(ctx, 200, gin.H{"user": dto.ToUserDto(user.(model.User))}, "获取用户信息成功")

}

func IsTelephoneExist(db *gorm.DB, telephone string) bool{
	var user model.User;
	db.Where("Telephone = ?", telephone).First(&user)
	if user.ID != 0{
		return true
	}
	return false
}
package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"xiangerfer.com/db_flagship/common"
	"xiangerfer.com/db_flagship/model"
)

func AuthMiddleware() gin.HandlerFunc{
	return func(ctx *gin.Context){
		// 获取authorization Header
		tokenString := ctx.GetHeader("Authorization")

		// 验证
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer"){
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg": "权限不足, token没有",
			})
			ctx.Abort()
			return
		}
		// 除去Bearer
		tokenString = tokenString[7:]

		token, claims, err := common.ParseToken(tokenString)

		if err != nil || !token.Valid{
			//ctx.Set("user", nil)

			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg": "权限不足, token有问题",
			})
			ctx.Abort()
			return
		}

		// 验证通过后，通过claims 获取Userid
		userId := claims.UserId
		DB := common.GetDB()
		var user model.User
		DB.First(&user, userId)


		// 用户不存在
		if user.ID == 0{
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg": "权限不足,用户不存在",
			})
		}

		// 用户存在,将用户信息写入上下文
		ctx.Set("user", user)
		// 这一句话是将控制权交给下个handler
		ctx.Next()

	}
}

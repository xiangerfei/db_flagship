package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"xiangerfer.com/db_flagship/common"
	"xiangerfer.com/db_flagship/response"
)


var logger = common.Logger{}


func ApiParam(ctx *gin.Context){
	name := ctx.Param("name")
	action := ctx.Param("action")
	other := ctx.Param("other")
	ctx.JSON(http.StatusOK, gin.H{
		"message": name + action + other,
	})
}

func UrlQuery( ctx *gin.Context){
	/*
	请求：127.0.0.1:8081/api/v1/user/url?username=yixiang&password=helloworld
	返回：{
	    "email": "475@qq.com",
	    "password": "helloworld",
	    "username": "yixiang"
	}
	*/
	username := ctx.Query("username")
	password := ctx.Query("password")
	email := ctx.DefaultQuery("email", "475@qq.com")
	ctx.JSON(http.StatusOK, gin.H{
		"username": username,
		"password": password,
		"email": email,
	})
}


func FormPostParam(ctx *gin.Context){
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	email := ctx.DefaultPostForm("email", "18340016418@163.com")

	ctx.JSON(http.StatusOK, gin.H{
		"username": username,
		"password": password,
		"email": email,
	})
}


/*pictures_file*/
func FormPostFile(ctx *gin.Context){
	//var pictures_file *multipart.FileHeader
	//var err error
	pictures_file, err:= ctx.FormFile("pictures_file")
	if err != nil{

	}
	open_file, err := pictures_file.Open()
	defer open_file.Close()

	cur_dir, _ := os.Getwd()
	fmt.Println("当前文件夹 :", cur_dir)
	new_filename := cur_dir + "/" + "files/" + pictures_file.Filename
	new_file, err := os.OpenFile(new_filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	defer new_file.Close()

	if err != nil{
		logger.Info("保存上传文件-->打开文件失败")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg": "失败",
			"data": nil,
		})
		return
	}
	buf := make([]byte, 1024) //一次读取1024字节
	/* 循环读取文件内容 */
	for {
		n, err := open_file.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
			logger.Info("保存上传文件-->保存文件内容失败")
		}
		if err == io.EOF{
			break
		}
		//io.WriteString(new_file, string(buf))
		new_file.Write(buf[0:n])
	}

	response.Success(ctx, http.StatusOK, nil, "保存文件成功")
}

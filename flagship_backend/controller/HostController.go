package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/* 查询主机信息 */
func HostInfo(ctx *gin.Context){
	if host_id, is_exist := ctx.Get("host_id"); is_exist{
		ctx.String(http.StatusOK, "%s%s", "查询指定主机信息", host_id)
	}
	ctx.String(http.StatusOK, "%s", "查询所有主机信息")

}
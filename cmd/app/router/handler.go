package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"rbac-demo/pkg/httputils"
	"rbac-demo/pkg/types"
)

func Login(c *gin.Context) {
	var login types.Login
	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(400, gin.H{"status": 400, "msg": "解析login json失败"})
		return
	}
	fmt.Println(login)
	token, err := httputils.GenerateToken(1, []byte("jixingxing"))
	if err != nil {
		c.JSON(400, gin.H{"status": 400, "msg": "获取token 失败"})
		return
	}
	c.JSON(200, gin.H{"status": 200, "msg": "获取token 成功", "token": token})

}

func Metrics(c *gin.Context) {
	c.JSON(200, gin.H{"status": 200, "msg": "success"})
}

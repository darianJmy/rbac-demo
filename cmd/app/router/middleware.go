package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"rbac-demo/pkg/httputils"
	"strings"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
	}
}

func HandleToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println(c.Request.URL.Path)
		if c.Request.URL.Path == "/login" {
			return
		}
		tokenString := c.GetHeader("Authorization")

		if len(tokenString) == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{"status": 401, "msg": "权限不足"})
			c.Abort()
			return
		}

		fields := strings.Fields(tokenString)
		if len(fields) != 2 {
			c.JSON(http.StatusUnauthorized, gin.H{"status": 401, "msg": "权限不足"})
			c.Abort()
			return
		}

		if fields[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": 401, "msg": "权限不足"})
			c.Abort()
			return
		}

		accessToken := fields[1]
		jwtKey := "jixingxing"
		token, err := httputils.ParseToken(accessToken, []byte(jwtKey))
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"status": 401, "msg": "权限不足"})
			c.Abort()
			return
		}

		fmt.Println(token)
		c.Next()
	}
}

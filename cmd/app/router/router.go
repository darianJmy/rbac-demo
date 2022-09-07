package router

import "github.com/gin-gonic/gin"

func RegistryRoutes(ginEngine *gin.Engine) {
	//ginEngine.Use(LoggerToFile(), IPWhiteList)
	ginEngine.Use(Cors(), HandleToken())
	{
		//ginEngine.Group("/users")
		ginEngine.POST("/login", Login)
		ginEngine.GET("/metrics", Metrics)
	}
}

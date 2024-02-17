package server

import (
	"github.com/gin-gonic/gin"
	"lease/middleware"
	"lease/service"
	"log"
)

func Start() {
	e := gin.Default()

	////  解决跨域请求
	//mwCORS := cors.New(cors.Config{
	//	AllowOrigins:     []string{"http://localhost:8081/"},
	//	AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
	//	AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
	//	ExposeHeaders:    []string{"Content-Type"},
	//	AllowCredentials: true,
	//	AllowOriginFunc: func(origin string) bool {
	//		return true
	//	},
	//	MaxAge: 24 * time.Hour,
	//})
	//e.Use(mwCORS)

	e.POST("/hello", service.HelloWorld)
	e.POST("/login", service.Login)
	e.POST("/register", service.Register)
	i := e.Group("/index", middleware.JWT())
	{
		i.POST("/logout", service.Logout)
		i.GET("/getUserInfo", service.GetUserInfo)
		i.POST("/upload", service.Upload)

		i.GET("/address", service.GetAddress)
		i.GET("/address/:id", service.GetAddressByID)
		i.PUT("/address/:id/default", service.SetDefaultAddress)
		i.POST("/address", service.AddAddress)
		i.PUT("/address/:id", service.UpdateAddress)
		i.DELETE("/address/:id", service.DeleteAddress)
	}

	err := e.Run(":8080")
	if err != nil {
		log.Println("服务器启动失败")
		return
	}
}

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

	e.POST("/login", service.Login)
	e.POST("/register", service.Register)
	i := e.Group("/index", middleware.JWT())
	{
		i.POST("/logout", service.Logout)
		i.GET("/getUserInfo", service.GetUserInfo)
		i.GET("/getUserInfo/:id", service.GetOtherUserInfo)
		i.PUT("/updateUser", service.UpdateUser)
		i.POST("/upload", service.Upload)

		i.GET("/address", service.GetAddress)
		i.GET("/address/:id", service.GetAddressByID)
		i.PUT("/address/:id/default", service.SetDefaultAddress)
		i.POST("/address", service.AddAddress)
		i.PUT("/address/:id", service.UpdateAddress)
		i.DELETE("/address/:id", service.DeleteAddress)

		i.POST("/product", service.AddProduct)
		i.GET("/product/:id", service.GetProduct)
		i.PUT("/product/status", service.UpdateProductStatus)
		i.GET("/product/my/:status", service.GetMyProduct)
		i.GET("/product/list", service.GetProductList)
		i.POST("/product/mainImage/:id", service.UploadProductMainImage)
		i.POST("/product/image/:id", service.UploadProductImage)

		i.GET("collection/isCollection/:product_id", service.IsCollection)
		i.POST("collection/addOrCancel", service.Collection)
		i.GET("collection", service.GetCollection)

		i.POST("/message", service.SendMessage)
		i.GET("/message/user", service.GetChatUser)
		i.GET("/message/list/:toId", service.GetMessage)

		i.GET("/comment/two/:productId", service.GetTwoComment)
		i.GET("/comment/root/:productId", service.GetRootComment)
		i.GET("/comment/list/:commentId", service.GetCommentList)
		i.POST("/comment", service.AddComment)

		i.POST("order", service.AddOrder)
		i.GET("order/:id", service.GetOrder)
		i.DELETE("order/:id", service.CancelOrder)
		i.GET("order/my", service.GetMyOrder)
		i.PUT("order/receive/my/:id", service.IReceiveOrder)
		i.PUT("order/return/my/:id", service.IReturnOrder)
		i.GET("order/his", service.GetHisOrder)
		i.PUT("order/delivery/his/:id", service.HeDeliveryOrder)
		i.PUT("order/receive/his/:id", service.HeReceiveOrder)
		i.PUT("order/inspect/problem/his/:id", service.HeInspectOrderHasProblem)
		i.PUT("order/inspect/ok/his/:id", service.HeInspectOrderWithoutProblem)
		i.PUT("order/solve/:id", service.SolveOrder)
		i.GET("order/count", service.GetOrderCount)

		i.GET("/alipay/:id", service.PayUrl)
		i.GET("payment", service.GetPayment)
		i.POST("withdraw", service.Withdraw)
	}
	e.GET("/category/:parentId", service.GetCategory)

	e.GET("/callback", service.Callback)
	e.POST("/notify", service.Notify)

	err := e.Run(":8080")
	if err != nil {
		log.Println("服务器启动失败")
		return
	}
}

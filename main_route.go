package main

import (
	"shopfood/component/appctx"
	"shopfood/middleware"
	"shopfood/module/restaurant/transport/ginrestaurant"
	"shopfood/module/upload/transport/ginupload"
	"shopfood/module/user/transport/ginuser"

	"github.com/gin-gonic/gin"
)

func SetupRouter(appContext appctx.AppContext, v1 *gin.RouterGroup) {
	// API /upload file
	v1.POST("/upload", ginupload.Upload(appContext))

	// API /users
	users := v1.Group("/users")
	{
		users.POST("/register", ginuser.Register(appContext))
		users.POST("/authenticate", ginuser.Login(appContext))
		users.GET("/profile", middleware.RequiredAuth(appContext), ginuser.Profile(appContext))
	}

	// API /restaurants
	restaurants := v1.Group("/restaurants", middleware.RequiredAuth(appContext))
	{
		restaurants.POST("", ginrestaurant.CreateRestaurant(appContext))
		restaurants.DELETE("/:id", ginrestaurant.DeleteRestaurant(appContext))
		restaurants.GET("", ginrestaurant.ListRestaurant(appContext))
		restaurants.GET("/:id", ginrestaurant.DetailRestaurant(appContext))
		restaurants.PATCH("/:id", ginrestaurant.UpdateRestaurant(appContext))
	}

}

package main

import (
	"shopfood/component/appctx"
	"shopfood/middleware"
	"shopfood/module/user/transport/ginuser"

	"github.com/gin-gonic/gin"
)

func SetupAdminRouter(appContext appctx.AppContext, v1 *gin.RouterGroup) {
	// API /admin
	admin := v1.Group("/admin",
		middleware.RequiredAuth(appContext),
		middleware.RoleRequired(appContext, "admin", "mod"),
	)

	{
		admin.GET("/profile", ginuser.Profile(appContext))
	}
}

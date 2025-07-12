package router

import (
	"github.com/gin-gonic/gin"
	"github.com/liuzhaomax/go-maxms/internal/core"
	"github.com/liuzhaomax/go-maxms/internal/middleware"
	"github.com/liuzhaomax/go-maxms/src/api_user/handler"
)

func Register(root *gin.RouterGroup, handler *handler.HandlerUser, mw *middleware.Middleware) {
	root.GET("/login", core.WrapperRes(func(c *gin.Context) (any, error) {
		return handler.GetPuk(c)
	}))
	root.POST("/login", core.WrapperRes(func(c *gin.Context) (any, error) {
		return handler.PostLogin(c)
	}))

	root.Use(mw.Auth.ValidateToken())
	root.POST("/user/avatar", core.WrapperRes(func(c *gin.Context) (any, error) {
		return handler.PostAvatar(c)
	}))
	root.PATCH("/user/nickName", core.WrapperRes(func(c *gin.Context) (any, error) {
		return handler.PatchNickName(c)
	}))
	root.GET("/user/user", core.WrapperRes(func(c *gin.Context) (any, error) {
		return handler.GetUser(c)
	}))
	root.POST("/group", core.WrapperRes(func(c *gin.Context) (any, error) {
		return handler.PostGroup(c)
	}))
	root.PATCH("/group/join", core.WrapperRes(func(c *gin.Context) (any, error) {
		return handler.PatchJoinGroup(c)
	}))
	root.DELETE("/group/quit", core.WrapperRes(func(c *gin.Context) (any, error) {
		return handler.DeleteQuitGroup(c)
	}))
}

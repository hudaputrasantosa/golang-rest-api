package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initUserRoute(route *gin.Engine) {
	groupRoute := route.Group("api/v1")

	groupRoute.GET("/users", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Get User!",
		})
	})
}

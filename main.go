package main

import (
	"log"

	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := SetupRouter()

	// r.GET("/", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	// })

	log.Fatal(router.Run())
}

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"*"},
		AllowHeaders:  []string{"*"},
		AllowWildcard: true,
	}))

	// var routes router.InitUserRoute(router)

	router.Use(helmet.Default())
	routes.initUserRoute(router)
	return router
}

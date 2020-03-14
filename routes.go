package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	ginlogrus "github.com/toorop/gin-logrus"
	"github.com/weblair/ag7if/controllers"
)

func configureRoutingGroup(g *gin.RouterGroup)  {
	// Visit {host}/api/v1/swagger/index.html to see the API documentation.
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	services := controllers.NewServicesController()
	{
		g.POST("/services", services.Create)
		g.GET("/services", services.List)
		g.GET("/services/:id", services.Fetch)
		g.PUT("/services/:id", services.Update)
		g.DELETE("/services/:id", services.Delete)
	}
}

func newRouter(version string, logger *logrus.Logger) *gin.Engine {
	router := gin.New()
	router.Use(ginlogrus.Logger(logger), gin.Recovery())

	v1 := router.Group("/api/v1")
	health := controllers.NewHealthController(version)
	{
		v1.GET("/health", health.Check)
	}
	configureRoutingGroup(v1)

	return router
}

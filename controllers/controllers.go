package controllers

import "github.com/gin-gonic/gin"

type Controller interface {
	Create(c *gin.Context)
	List(c *gin.Context)
	Fetch(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

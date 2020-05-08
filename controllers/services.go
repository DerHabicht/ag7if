package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/weblair/ag7if/db"
	"github.com/weblair/ag7if/models"
	"net/http"
	"strconv"
)

type ServicesController struct{
	resourceModel models.Service
}

func NewServicesController() ServicesController {
	return ServicesController{}
}

func (s ServicesController) ResourceModel() interface{} {
	return s.resourceModel
}

func (s ServicesController) Create(c *gin.Context) {

}

func (s ServicesController) List(c *gin.Context) {
	ListRecords(c, s)
}

func (s ServicesController) Fetch(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err,
			"id":    c.Param("id"),
		}).Fatal("Invalid ID passed for lookup.")
	}

	var svc models.Service
	db.Tx.Preload("Bands").Find(&svc, id)

	c.JSON(http.StatusOK, svc)
}

func (s ServicesController) Update(c *gin.Context) {

}

func (s ServicesController) Delete(c *gin.Context) {

}

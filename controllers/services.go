package controllers

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"github.com/weblair/ag7if/db"
	"github.com/weblair/ag7if/models"
	"net/http"
)

type ServicesController struct {}

func NewServicesController() ServicesController {
	return ServicesController{}
}

func (s ServicesController) Create(c *gin.Context) {

}

func (s ServicesController) List(c *gin.Context) {

}

func (s ServicesController) Fetch(c *gin.Context) {
	public_id, err := uuid.FromString(c.Param("id"))
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err,
		}).Fatal("Failed to parse UUID")
	}

	svc := &models.Service{}
	db.DB.Where("public_id = ?", public_id).Preload("Bands").First(svc)

	c.JSON(http.StatusOK, svc)
}

func (s ServicesController) Update(c *gin.Context) {

}

func (s ServicesController) Delete(c *gin.Context) {

}

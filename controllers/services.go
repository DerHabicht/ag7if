package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/weblair/maricopa/models"
)

type ServicesController struct{
	resourceModel models.Service
}

func NewServicesController() ServicesController {
	return ServicesController{}
}

func (s ServicesController) ResourceModel() models.Model {
	return s.resourceModel
}

func (s ServicesController) Create(c *gin.Context) {

}

func (s ServicesController) List(c *gin.Context) {
	ListRecords(c, s, false)
}

func (s ServicesController) Fetch(c *gin.Context) {
	FetchRecord(c, s, "public_id", []string{"Bands"})
}

func (s ServicesController) Update(c *gin.Context) {

}

func (s ServicesController) Delete(c *gin.Context) {

}

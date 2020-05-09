package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"reflect"

	"github.com/weblair/ag7if/database"
	"github.com/weblair/ag7if/models"
)

type Controller interface {
	ResourceModel() models.Model
	Create(c *gin.Context)
	List(c *gin.Context)
	Fetch(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

func CreateRecord(c *gin.Context) {

}

// ListRecords uses reflection to fetch a list of model instances from the database.
func ListRecords(ctx *gin.Context, ctl Controller) {
	// Create resource slice
	t := reflect.TypeOf(ctl.ResourceModel())
	s := reflect.MakeSlice(reflect.SliceOf(t), 0, 0)

	// Create pointer to resource slice
	p := reflect.New(s.Type())
	p.Elem().Set(s)

	// Get the actual resource slice
	r := p.Interface()

	logrus.WithFields(logrus.Fields{
		"resource_method": "list",
		"controller": reflect.TypeOf(ctl).String(),
		"resource_type": reflect.TypeOf(ctl.ResourceModel()).String(),
		"resource_list_type": reflect.TypeOf(r).String(),
	}).Debug("Fetching list of resources from database")

	database.DB.Find(r)

	ctx.JSON(http.StatusOK, r)
}

func FetchRecord(c *gin.Context, ctl Controller, param string, preload []string) {
	tx := database.DB.Where(fmt.Sprintf("%s = ?", param), c.Param(param))

	for _, v := range preload {
		tx = tx.Preload(v)
	}

	r := reflect.New(reflect.TypeOf(ctl.ResourceModel())).Interface()

	logrus.WithFields(logrus.Fields{
		"resource_method": "fetch",
		"controller": reflect.TypeOf(ctl).String(),
		"resource_type": reflect.TypeOf(r).String(),
		"lookup_param": param,
		"lookup_value": c.Param(param),
	}).Debug("Fetching resource from database")

	tx.First(r)

	c.JSON(http.StatusOK, r)
}

func UpdateRecord(c *gin.Context) {

}

func DeleteRecord(c *gin.Context) {

}

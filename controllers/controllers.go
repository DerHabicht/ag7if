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

	// Reflection on resource
	rt := reflect.TypeOf(ctl.ResourceModel())
	r := reflect.New(rt)

	// Reflection on resource ID field
	idt := reflect.TypeOf(ctl.ResourceModel().GetID())
	idz := reflect.Zero(idt).Interface()
	idn := ctl.ResourceModel().GetIDFieldName()

	tx.First(r.Interface())

	logrus.WithFields(logrus.Fields{
		"resource_method": "fetch",
		"controller": reflect.TypeOf(ctl).String(),
	}).Info("Fetching resource from database")

	logrus.WithFields(logrus.Fields{
		"resource_type": rt.String(),
		"lookup_param": param,
		"lookup_value": c.Param(param),
		"is_valid": r.IsValid(),
	}).Debug("Resource reflection info")

	logrus.WithFields(logrus.Fields{
		"resource_id_value": r.Elem().FieldByName(idn).Interface(),
		"resource_id_type": idt.String(),
		"resource_id_name": idn,
		"resource_id_zero_value": idz,
	}).Debug("Resource ID field reflection info")

	if r.Elem().FieldByName(idn).Interface() == idz {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "resource not found",
			"id": c.Param(param),
		})
		return
	}

	c.JSON(http.StatusOK, r.Interface())
}

func UpdateRecord(c *gin.Context) {

}

func DeleteRecord(c *gin.Context) {

}

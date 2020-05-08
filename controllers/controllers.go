package controllers

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"

	"github.com/weblair/ag7if/db"
)

type Controller interface {
	ResourceModel() interface{}
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
		"resource_type": reflect.TypeOf(ctl.Resource()).String(),
		"resource_list_type": reflect.TypeOf(r).String(),
	}).Debug("Fetching list of resources from database")

	db.Tx.Find(r)

	ctx.JSON(http.StatusOK, r)
}

func FetchRecord(c *gin.Context) {

}

func UpdateRecord(c *gin.Context) {

}

func DeleteRecord(c *gin.Context) {

}

package main

import (
	"fmt"
	"net/http"

	"./controller"
	"./models"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func main() {
	m1 := models.Metric{Moisture: 30.0, Temp: 10.0}
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if db == nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&models.Metric{})
	fmt.Println(m1)
	db.Create(&m1)

	c1 := controller.MetricHandler{Db: db}

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello running Server Iot central")
	})

	r.GET("/metricsgetlist", c1.GetMetricsController)

	r.POST("/createmetric", PostMetricCreate)
	r.POST("/deletemetric", DeleteMetricCreate)
	r.POST("/updatemetric", UpdateMetricCreate)

	r.Run()
}

//PostMetricCreate controller for create a new register metric
func PostMetricCreate(c *gin.Context) {

	var metric models.Metric
	if c.BindJSON(&metric) == nil {
		db.Create(&metric)
		c.JSON(200, metric)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error processing"})
	}
}

//DeleteMetricCreate controller for create a new register metric
func DeleteMetricCreate(c *gin.Context) {

	var listMetric []models.Metric
	var listID models.List

	if c.BindJSON(&listID) == nil {
		if db.Find(&listMetric, listID.ListID); len(listMetric) != 0 {
			db.Delete(&listMetric, listID.ListID)
			c.JSON(200, gin.H{"success Deleted": (listID.ListID)})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "there are not those ids"})
		}

	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error processing"})
	}
}

//UpdateMetricCreate controller for create a new register metric
func UpdateMetricCreate(c *gin.Context) {

	var metric models.Metric
	if c.BindJSON(&metric) == nil {

		db.Save(&metric)
		c.JSON(200, metric)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error processing"})
	}
}

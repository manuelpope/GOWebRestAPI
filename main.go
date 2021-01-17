package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func main() {
	m1 := Metric{Moisture: 30.0, Temp: 10.0}
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if db == nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Metric{})
	fmt.Println(m1)
	db.Create(&m1)

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello running Server Iot central")
	})

	r.GET("/metricsgetlist", GetMetricsController)

	r.POST("/createmetric", PostMetricCreate)
	r.POST("/deletemetric", DeleteMetricCreate)
	r.POST("/updatemetric", UpdateMetricCreate)

	r.Run()
}

//GetMetricsController controller for retrieve all data Metric
func GetMetricsController(c *gin.Context) {
	var metrics []Metric
	if err := db.Find(&metrics).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, metrics)
	}
}

//PostMetricCreate controller for create a new register metric
func PostMetricCreate(c *gin.Context) {

	var metric Metric
	if c.BindJSON(&metric) == nil {
		db.Create(&metric)
		c.JSON(200, metric)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error processing"})
	}
}

//DeleteMetricCreate controller for create a new register metric
func DeleteMetricCreate(c *gin.Context) {

	var listMetric []Metric
	var listID List

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

	var metric Metric
	if c.BindJSON(&metric) == nil {

		db.Save(&metric)
		c.JSON(200, metric)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error processing"})
	}
}

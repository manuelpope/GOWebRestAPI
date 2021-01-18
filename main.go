package main

import (
	"fmt"

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

	r.POST("/createmetric", c1.PostMetricCreate)
	r.POST("/deletemetric", c1.DeleteMetricCreate)
	r.POST("/updatemetric", c1.UpdateMetricCreate)

	r.Run()
}

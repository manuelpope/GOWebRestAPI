package controller

import (
	"fmt"
	"log"

	"../models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

//MetricHandler func 2 hanlder router
type MetricHandler struct {
	Db *gorm.DB
}

//GetMetricsController controller for retrieve all data Metric
func (mh *MetricHandler) GetMetricsController(c *gin.Context) {
	log.Println("listing all metrics")
	var metrics []models.Metric
	if err := mh.Db.Find(&metrics).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, metrics)
	}
}

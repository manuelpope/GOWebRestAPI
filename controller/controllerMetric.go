package controller

import (
	"fmt"
	"log"
	"net/http"

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

//PostMetricCreate controller for create a new register metric
func (mh *MetricHandler) PostMetricCreate(c *gin.Context) {

	var metric models.Metric
	if c.BindJSON(&metric) == nil {
		mh.Db.Create(&metric)
		c.JSON(200, metric)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error processing"})
	}
}

//DeleteMetricCreate controller for create a new register metric
func (mh *MetricHandler) DeleteMetricCreate(c *gin.Context) {

	var listMetric []models.Metric
	var listID models.List

	if c.BindJSON(&listID) == nil {
		if mh.Db.Find(&listMetric, listID.ListID); len(listMetric) != 0 {
			mh.Db.Delete(&listMetric, listID.ListID)
			c.JSON(200, gin.H{"success Deleted": (listID.ListID)})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "there are not those ids"})
		}

	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error processing"})
	}
}

//UpdateMetricCreate controller for create a new register metric
func (mh *MetricHandler) UpdateMetricCreate(c *gin.Context) {

	var metric models.Metric
	if c.BindJSON(&metric) == nil {

		mh.Db.Save(&metric)
		c.JSON(200, metric)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error processing"})
	}
}

package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yehezkiel1086/go-gin-employees-training-enrollment-system/internal/core/port"
)

type StatisticsHandler struct {
	svc port.StatisticsService
}

func InitStatisticsHandler(svc port.StatisticsService) *StatisticsHandler {
	return &StatisticsHandler{
		svc: svc,
	}
}

func (sh *StatisticsHandler) GetTrainingStatistics(c *gin.Context) {
	stats, err := sh.svc.GetTrainingStatistics()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, stats)
}

func (sh *StatisticsHandler) GetTrainingsByCategories(c *gin.Context) {
	stats, err := sh.svc.GetTrainingsByCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, stats)
}

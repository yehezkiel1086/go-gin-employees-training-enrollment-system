package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yehezkiel1086/go-gin-employees-training-enrollment-system/internal/core/domain"
	"github.com/yehezkiel1086/go-gin-employees-training-enrollment-system/internal/core/port"
	"github.com/yehezkiel1086/go-gin-employees-training-enrollment-system/internal/core/util"
)

type TrainingHandler struct {
	svc port.TrainingService
}

func InitTrainingHandler(svc port.TrainingService) *TrainingHandler {
	return &TrainingHandler{
		svc: svc,
	}
}

type CreateTrainingReq struct {
	Title string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Date string `json:"date" binding:"required"` // start date
	Duration int `json:"duration" binding:"required"` // duration in days
	Instructor string `json:"instructor" binding:"required"`
}

func (th *TrainingHandler) CreateTraining(c *gin.Context) {
	var req CreateTrainingReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("title, description, date, duration and instructor are required: %v", err.Error()),
		})
		return
	}

	// parse string date
	dateLayout := "2006-01-02"
	date, err := util.ParseDate(dateLayout, req.Date)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("invalid date type: %v", err.Error()),
		})
		return
	}

	training, err := th.svc.CreateTraining(c, &domain.Training{
		Title: req.Title,
		Description: req.Description,
		Date: date,
		Duration: req.Duration,
		Instructor: req.Instructor,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, training)
}

func (th *TrainingHandler) GetTrainings(c *gin.Context) {
	trainings, err := th.svc.GetTrainings(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, trainings)
}

func (th *TrainingHandler) GetTrainingByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id is required",
		})
		return
	}

	training, err := th.svc.GetTrainingByID(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, training)
}
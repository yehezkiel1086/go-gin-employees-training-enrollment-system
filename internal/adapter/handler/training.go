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
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Date        string `json:"date" binding:"required"` // start date
	Duration    int    `json:"duration" binding:"required"` // duration in days
	Instructor  string `json:"instructor" binding:"required"`
	CategoryID  uint   `json:"category_id" binding:"required"`
}

type UpdateTrainingReq struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Date        string `json:"date"` // start date
	Duration    int    `json:"duration"`
	Instructor  string `json:"instructor"`
	CategoryID  uint   `json:"category_id"`
}

func (th *TrainingHandler) CreateTraining(c *gin.Context) {
	var req CreateTrainingReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("title, description, date, duration, instructor and category_id are required: %v", err.Error()),
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
		CategoryID: req.CategoryID,
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

func (th *TrainingHandler) DeleteTraining(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id is required",
		})
		return
	}

	// get training by id
	training, err := th.svc.GetTrainingByID(c, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "training not found",
		})
		return
	}

	if err := th.svc.DeleteTraining(c, training); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "training deleted successfully",
	})
}

func (th *TrainingHandler) UpdateTraining(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id is required",
		})
		return
	}

	var req UpdateTrainingReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	training, err := th.svc.GetTrainingByID(c, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "training not found",
		})
		return
	}

	// Update fields if they are provided in the request
	if req.Title != "" {
		training.Title = req.Title
	}
	if req.Description != "" {
		training.Description = req.Description
	}
	if req.Date != "" {
		dateLayout := "2006-01-02"
		date, err := util.ParseDate(dateLayout, req.Date)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("invalid date type: %v", err.Error())})
			return
		}
		training.Date = date
	}
	if req.Duration != 0 {
		training.Duration = req.Duration
	}
	if req.Instructor != "" {
		training.Instructor = req.Instructor
	}
	if req.CategoryID != 0 {
		training.CategoryID = req.CategoryID
	}

	updatedTraining, err := th.svc.UpdateTraining(c, training)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedTraining)
}
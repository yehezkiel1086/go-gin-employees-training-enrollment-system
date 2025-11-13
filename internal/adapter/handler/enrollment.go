package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yehezkiel1086/go-gin-employees-training-enrollment-system/internal/core/port"
)

type EnrollmentHandler struct {
	svc port.EnrollmentService
}

func InitEnrollmentHandler(svc port.EnrollmentService) *EnrollmentHandler {
	return &EnrollmentHandler{
		svc: svc,
	}
}

type CreateEnrollmentReq struct {
	Email string `json:"email" binding:"required,email"`
	TrainingID uint `json:"training_id" binding:"required"`
}

func (eh *EnrollmentHandler) CreateEnrollment(c *gin.Context) {
	var req CreateEnrollmentReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// set enrolled at to now
	enrolledAt := time.Now()

	if err := eh.svc.CreateEnrollment(c, req.Email, req.TrainingID, enrolledAt); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "enrollment created successfully",
	})
}

func (eh *EnrollmentHandler) GetEnrollments(c *gin.Context) {
	enrollments, err := eh.svc.GetEnrollments(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, enrollments)
}

func (eh *EnrollmentHandler) GetEnrollmentsByEmail(c *gin.Context) {
	email := c.Param("email")
	if email == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id parameter is required",
		})
	}

	enrollments, err := eh.svc.GetEnrollmentsByEmail(c, email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, enrollments)
}

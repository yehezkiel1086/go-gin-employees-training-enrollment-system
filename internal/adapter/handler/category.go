package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yehezkiel1086/go-gin-employees-training-enrollment-system/internal/core/domain"
	"github.com/yehezkiel1086/go-gin-employees-training-enrollment-system/internal/core/port"
)

type CategoryHandler struct {
	svc port.CategoryService
}

func InitCategoryHandler(svc port.CategoryService) *CategoryHandler {
	return &CategoryHandler{
		svc: svc,
	}
}

type CreateCategoryReq struct {
	Name string `json:"name" binding:"required"`
}

func (ch *CategoryHandler) CreateCategory(c *gin.Context) {
	var req CreateCategoryReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	category, err := ch.svc.CreateCategory(c, &domain.Category{
		Name: req.Name,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, category)
}

func (ch *CategoryHandler) GetCategories(c *gin.Context) {
	categories, err := ch.svc.GetCategories(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, categories)
}

func (ch *CategoryHandler) DeleteCategory(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id is required",
		})
		return
	}
	category, err := ch.svc.GetCategoryByID(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := ch.svc.DeleteCategory(c, category); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "category deleted successfully",
	})
}

func (ch *CategoryHandler) GetCategoryByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id parameter is required",
		})
		return
	}

	category, err := ch.svc.GetCategoryByID(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, category)
}

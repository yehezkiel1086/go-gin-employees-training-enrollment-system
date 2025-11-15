package handler

import (
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/yehezkiel1086/go-gin-employees-training-enrollment-system/internal/adapter/config"
)

type Router struct {
	r *gin.Engine
}

func InitRouter(
	conf *config.JWT,
	userHandler *UserHandler,
	authHandler *AuthHandler,
	trainingHandler *TrainingHandler,
	enrollmentHandler *EnrollmentHandler,
	statisticsHandler *StatisticsHandler,
	categoryHandler *CategoryHandler,
) *Router {
	r := gin.New()

	// Add CORS middleware
	allowedOrigins := strings.Split(conf.AllowedOrigins, ",")

	r.Use(cors.New(cors.Config{
			AllowOrigins:     allowedOrigins,
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
	}))

	pb := r.Group("/api/v1")
	us := pb.Group("/", AuthMiddleware(conf))
	ad := us.Group("/", AdminAuthMiddleware())

	// public routes
	pb.POST("/register", userHandler.RegisterNewUser)
	pb.POST("/login", authHandler.Login)

	// user only routes
	us.GET("/logout", authHandler.Logout)

	us.GET("/users/:email", CheckEmailParam(), userHandler.GetUserByEmail)

	us.GET("/trainings", trainingHandler.GetTrainings)
	us.GET("/trainings/:id", trainingHandler.GetTrainingByID)
	
	us.GET("/enrollments/:email", CheckEmailParam(), enrollmentHandler.GetEnrollmentsByEmail)
	us.POST("/enrollments", enrollmentHandler.CreateEnrollment)

	us.GET("/statistics/trainings", statisticsHandler.GetTrainingStatistics)

	us.GET("/categories", categoryHandler.GetCategories)
	us.GET("/categories/:id", categoryHandler.GetCategoryByID)

	// admin only routes
	ad.GET("/users", userHandler.GetUsers)

	ad.POST("/trainings", trainingHandler.CreateTraining)
	ad.PUT("/trainings/:id", trainingHandler.UpdateTraining)
	ad.DELETE("/trainings/:id", trainingHandler.DeleteTraining)

	ad.GET("/enrollments", enrollmentHandler.GetEnrollments)
	
	ad.POST("/categories", categoryHandler.CreateCategory)
	ad.DELETE("/categories/:id", categoryHandler.DeleteCategory)

	return &Router{
		r: r,
	}
}

func (r *Router) Serve(conf *config.HTTP) error {
	uri := conf.Host + ":" + conf.Port
	return r.r.Run(uri)
}

package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/yehezkiel1086/go-gin-employees-training-enrollment-system/internal/adapter/config"
)

type Router struct {
	r *gin.Engine
}

func InitRouter(
	userHandler *UserHandler,
	authHandler *AuthHandler,
) *Router {
	r := gin.New()

	pb := r.Group("/api/v1")

	// public routes
	pb.POST("/register", userHandler.RegisterNewUser)
	pb.POST("/login", authHandler.Login)

	// user only routes
	pb.GET("/users", userHandler.GetUsers)
	pb.GET("/users/:email", userHandler.GetUserByEmail)

	return &Router{
		r: r,
	}
}

func (r *Router) Serve(conf *config.HTTP) error {
	uri := conf.Host + ":" + conf.Port
	return r.r.Run(uri)
}

package main

import (
	"context"
	"fmt"

	"github.com/yehezkiel1086/go-gin-employees-training-enrollment-system/internal/adapter/config"
	"github.com/yehezkiel1086/go-gin-employees-training-enrollment-system/internal/adapter/handler"
	"github.com/yehezkiel1086/go-gin-employees-training-enrollment-system/internal/adapter/storage/postgres"
	"github.com/yehezkiel1086/go-gin-employees-training-enrollment-system/internal/adapter/storage/postgres/repository"
	"github.com/yehezkiel1086/go-gin-employees-training-enrollment-system/internal/core/domain"
	"github.com/yehezkiel1086/go-gin-employees-training-enrollment-system/internal/core/service"
)

func main() {
	// init .env config
	conf, err := config.InitConfig()
	if err != nil {
		panic(err)
	}
	fmt.Println(".env configs loaded successfully")

	// create context
	ctx := context.Background()

	// connect to postgres db
	db, err := postgres.InitDB(ctx, conf.DB)
	if err != nil {
		panic(err)
	}
	fmt.Println("database connected successfully")

	// migrate dbs
	if err := db.Migrate(&domain.User{}); err != nil {
		panic(err)
	}
	fmt.Println("database migrated successfully")

	// dependency injections	
	userRepo := repository.InitUserRepository(db)
	userSvc := service.InitUserService(userRepo)
	userHandler := handler.InitUserHandler(userSvc)

	authSvc := service.InitAuthService(userRepo)
	authHandler := handler.InitAuthHandler(conf.JWT, authSvc)

	// init router
	r := handler.InitRouter(
		conf.JWT,
		userHandler,
		authHandler,
	)
	fmt.Println("router initialized successfully")

	// serve api
	if err := r.Serve(conf.HTTP); err != nil {
		panic(err)
	}
}

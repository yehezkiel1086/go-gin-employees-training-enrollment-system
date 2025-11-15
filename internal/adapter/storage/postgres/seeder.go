package postgres

import (
	"fmt"
	"time"

	"github.com/yehezkiel1086/go-gin-employees-training-enrollment-system/internal/core/domain"
	"github.com/yehezkiel1086/go-gin-employees-training-enrollment-system/internal/core/util"
	"gorm.io/gorm"
)

func Seed(db *gorm.DB) error {
	if err := seedUsers(db); err != nil {
		return err
	}
	if err := seedCategories(db); err != nil {
		return err
	}
	if err := seedTrainings(db); err != nil {
		return err
	}
	if err := seedEnrollments(db); err != nil {
		return err
	}

	fmt.Println("database seeded successfully")
	return nil
}

func seedUsers(db *gorm.DB) error {
	var count int64
	db.Model(&domain.User{}).Count(&count)
	if count > 0 {
		return nil // Users table is not empty
	}

	hashedPassword, err := util.HashPassword("password")
	if err != nil {
		return err
	}

	users := []domain.User{
		{Name: "Admin User", Email: "admin@example.com", Password: hashedPassword, Role: domain.ADMIN_ROLE},
		{Name: "Regular User", Email: "user@example.com", Password: hashedPassword, Role: domain.USER_ROLE},
	}

	return db.Create(&users).Error
}

func seedCategories(db *gorm.DB) error {
	var count int64
	db.Model(&domain.Category{}).Count(&count)
	if count > 0 {
		return nil // Categories table is not empty
	}

	categories := []domain.Category{
		{Name: "Software Development"},
		{Name: "Project Management"},
		{Name: "Data Science"},
		{Name: "Human Resources"},
	}

	return db.Create(&categories).Error
}

func seedTrainings(db *gorm.DB) error {
	var count int64
	db.Model(&domain.Training{}).Count(&count)
	if count > 0 {
		return nil // Trainings table is not empty
	}

	var devCategory domain.Category
	if err := db.First(&devCategory, "name = ?", "Software Development").Error; err != nil {
		return err
	}

	var pmCategory domain.Category
	if err := db.First(&pmCategory, "name = ?", "Project Management").Error; err != nil {
		return err
	}

	trainings := []domain.Training{
		{Title: "Go for Beginners", Description: "An introductory course to Go programming language.", Date: time.Now().Add(30 * 24 * time.Hour), Duration: 5, Instructor: "John Doe", CategoryID: devCategory.ID},
		{Title: "Agile and Scrum", Description: "Learn the fundamentals of Agile and Scrum methodologies.", Date: time.Now().Add(45 * 24 * time.Hour), Duration: 3, Instructor: "Jane Smith", CategoryID: pmCategory.ID},
	}

	return db.Create(&trainings).Error
}

func seedEnrollments(db *gorm.DB) error {
	var count int64
	db.Model(&domain.Enrollment{}).Count(&count)
	if count > 0 {
		return nil // Enrollments table is not empty
	}

	var regularUser domain.User
	if err := db.First(&regularUser, "email = ?", "user@example.com").Error; err != nil {
		return err
	}

	var goTraining domain.Training
	if err := db.First(&goTraining, "title = ?", "Go for Beginners").Error; err != nil {
		return err
	}

	enrollment := domain.Enrollment{
		UserID:     regularUser.ID,
		TrainingID: goTraining.ID,
		EnrolledAt: time.Now(),
	}

	return db.Create(&enrollment).Error
}
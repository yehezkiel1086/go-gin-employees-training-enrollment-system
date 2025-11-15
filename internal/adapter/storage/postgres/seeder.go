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
		{Name: "John Doe", Email: "john.doe@example.com", Password: hashedPassword, Role: domain.USER_ROLE},
		{Name: "Jane Smith", Email: "jane.smith@example.com", Password: hashedPassword, Role: domain.USER_ROLE},
		{Name: "Peter Jones", Email: "peter.jones@example.com", Password: hashedPassword, Role: domain.USER_ROLE},
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
		{Name: "Marketing"},
		{Name: "Design"},
		{Name: "Business"},
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

	var dsCategory domain.Category
	if err := db.First(&dsCategory, "name = ?", "Data Science").Error; err != nil {
		return err
	}

	var designCategory domain.Category
	if err := db.First(&designCategory, "name = ?", "Design").Error; err != nil {
		return err
	}

	trainings := []domain.Training{
		{Title: "Go for Beginners", Description: "An introductory course to Go programming language.", Date: time.Now().Add(30 * 24 * time.Hour), Duration: 5, Instructor: "John Doe", CategoryID: devCategory.ID},
		{Title: "Agile and Scrum", Description: "Learn the fundamentals of Agile and Scrum methodologies.", Date: time.Now().Add(45 * 24 * time.Hour), Duration: 3, Instructor: "Jane Smith", CategoryID: pmCategory.ID},
		{Title: "Advanced Go", Description: "Deep dive into advanced Go concepts.", Date: time.Now().Add(60 * 24 * time.Hour), Duration: 7, Instructor: "John Doe", CategoryID: devCategory.ID},
		{Title: "Python for Data Science", Description: "Learn Python for data analysis and visualization.", Date: time.Now().Add(20 * 24 * time.Hour), Duration: 10, Instructor: "Peter Jones", CategoryID: dsCategory.ID},
		{Title: "UI/UX Design Principles", Description: "Fundamentals of User Interface and User Experience design.", Date: time.Now().Add(15 * 24 * time.Hour), Duration: 4, Instructor: "Alice Wonderland", CategoryID: designCategory.ID},
		{Title: "Digital Marketing 101", Description: "Introduction to digital marketing strategies.", Date: time.Now().Add(10 * 24 * time.Hour), Duration: 3, Instructor: "Bob Builder", CategoryID: 5}, // Assuming Marketing has ID 5
		{Title: "React for Web Development", Description: "Build modern web applications with React.", Date: time.Now().Add(25 * 24 * time.Hour), Duration: 8, Instructor: "Chris Evans", CategoryID: devCategory.ID},
		{Title: "Past Training: Intro to SQL", Description: "A course on SQL that has already passed.", Date: time.Now().Add(-10 * 24 * time.Hour), Duration: 5, Instructor: "Peter Jones", CategoryID: dsCategory.ID},
		{Title: "Past Training: SEO Fundamentals", Description: "A course on SEO that has already passed.", Date: time.Now().Add(-20 * 24 * time.Hour), Duration: 2, Instructor: "Bob Builder", CategoryID: 5}, // Assuming Marketing has ID 5
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

	var johnDoe domain.User
	if err := db.First(&johnDoe, "email = ?", "john.doe@example.com").Error; err != nil {
		return err
	}

	var janeSmith domain.User
	if err := db.First(&janeSmith, "email = ?", "jane.smith@example.com").Error; err != nil {
		return err
	}

	var peterJones domain.User
	if err := db.First(&peterJones, "email = ?", "peter.jones@example.com").Error; err != nil {
		return err
	}

	var goTraining domain.Training
	if err := db.First(&goTraining, "title = ?", "Go for Beginners").Error; err != nil {
		return err
	}
	var agileTraining domain.Training
	if err := db.First(&agileTraining, "title = ?", "Agile and Scrum").Error; err != nil {
		return err
	}
	var pythonTraining domain.Training
	if err := db.First(&pythonTraining, "title = ?", "Python for Data Science").Error; err != nil {
		return err
	}
	var sqlTraining domain.Training
	if err := db.First(&sqlTraining, "title = ?", "Past Training: Intro to SQL").Error; err != nil {
		return err
	}
	var seoTraining domain.Training
	if err := db.First(&seoTraining, "title = ?", "Past Training: SEO Fundamentals").Error; err != nil {
		return err
	}

	enrollments := []domain.Enrollment{
		{UserID: regularUser.ID, TrainingID: goTraining.ID, EnrolledAt: time.Now()},
		{UserID: regularUser.ID, TrainingID: agileTraining.ID, EnrolledAt: time.Now().Add(-24 * time.Hour)},
		{UserID: johnDoe.ID, TrainingID: goTraining.ID, EnrolledAt: time.Now().Add(-48 * time.Hour)},
		{UserID: johnDoe.ID, TrainingID: pythonTraining.ID, EnrolledAt: time.Now()},
		{UserID: janeSmith.ID, TrainingID: agileTraining.ID, EnrolledAt: time.Now()},
		{UserID: janeSmith.ID, TrainingID: sqlTraining.ID, EnrolledAt: time.Now().Add(-15 * 24 * time.Hour)}, // Enrolled before the training
		{UserID: peterJones.ID, TrainingID: pythonTraining.ID, EnrolledAt: time.Now()},
		{UserID: peterJones.ID, TrainingID: seoTraining.ID, EnrolledAt: time.Now().Add(-22 * 24 * time.Hour)}, // Enrolled before the training
	}

	return db.Create(&enrollments).Error
}
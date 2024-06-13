package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"simpleGoProject/config"
	"simpleGoProject/controllers"
	"simpleGoProject/repositories"
	"simpleGoProject/routes"
	"simpleGoProject/services"
	"simpleGoProject/validators"
)

func main() {
	config.InitDB()

	// Initialize repositories
	studentRepository := repositories.NewStudentRepository(config.DB)
	courseRepository := repositories.NewCourseRepository(config.DB)

	// Initialize services
	studentService := services.NewStudentService(studentRepository)
	courseService := services.NewCourseService(courseRepository)

	// Initialize controllers
	studentController := controllers.NewStudentController(studentService)
	courseController := controllers.NewCourseController(courseService)

	// Initialize Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Custom Validator
	e.Validator = validators.NewCustomValidator()

	// Register routes
	routes.RegisterStudentRoutes(e, studentController)
	routes.RegisterCourseRoutes(e, courseController)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

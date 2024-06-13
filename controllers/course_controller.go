// Package controllers provides the HTTP request handlers for course operations
package controllers

// Import necessary packages
import (
	"github.com/labstack/echo/v4" // Import the Echo framework
	"net/http"                    // Import the net/http package for HTTP status codes
	"simpleGoProject/models"      // Import the models package containing the Course struct
	"simpleGoProject/services"    // Import the services package containing the CourseService interface
	"strconv"                     // Import the strconv package for string conversion
)

// CourseController defines the controller for handling course-related requests
type CourseController struct {
	courseService services.CourseService // Service to interact with course business logic
}

// NewCourseController creates a new instance of CourseController
func NewCourseController(courseService services.CourseService) *CourseController {
	return &CourseController{courseService}
}

// CreateCourse handle the creation of a new course
func (c *CourseController) CreateCourse(ctx echo.Context) error {
	var course models.Course
	if err := ctx.Bind(&course); err != nil { // Bind the request body to the course struct
		return ctx.JSON(http.StatusBadRequest, err) // Return a 400 status code for bad request
	}
	err := c.courseService.CreateCourse(course) // Call the service to create a course
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err) // Return a 500 status code for internal server error
	}
	return ctx.JSON(http.StatusOK, course) // Return a 200 status code with the created course data
}

// GetCourses handles retrieving all courses
func (c *CourseController) GetCourses(ctx echo.Context) error {
	courses, err := c.courseService.GetCourses() // Call the service to get all courses
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err) // Return a 500 status code for internal server error
	}
	return ctx.JSON(http.StatusOK, courses) // Return a 200 status code with the courses data
}

// GetCourseByID handles retrieving a course by ID
func (c *CourseController) GetCourseByID(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id")) // Convert the URL parameter to an integer
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err) // Return a 400 status code for bad request
	}
	course, err := c.courseService.GetCourseByID(uint(id)) // Cal the service to get a course by ID
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err) // Return a 500 status code for internal server error
	}
	return ctx.JSON(http.StatusOK, course) // Return a 200 status code with the course data
}

// UpdateCourse handles updating an existing course
func (c *CourseController) UpdateCourse(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id")) // Convert the URL parameter to an integer
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err) // Return a 400 status code for bad request
	}
	var course models.Course
	if err := ctx.Bind(&course); err != nil { // Bind the request body to the course struct
		return ctx.JSON(http.StatusBadRequest, err) // Return a 400 status code for bad request
	}
	course.CourseID = uint(id)                                   // Set the CourseID from the URL parameter
	if err := c.courseService.UpdateCourse(course); err != nil { // Call the service to update a course
		return ctx.JSON(http.StatusInternalServerError, err) // Return a 500 status code for internal server error
	}
	return ctx.JSON(http.StatusOK, course) // Return a 200 status code with the updated course data
}

// DeleteCourse handles deleting a course by ID
func (c *CourseController) DeleteCourse(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id")) // Convert the URL parameter to an integer
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err) // Return a 400 status code for bad request
	}
	err = c.courseService.DeleteCourse(uint(id)) // Call the service to delete a course by ID
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err) // Return a 500 status code for internal server error
	}
	return ctx.JSON(http.StatusOK, map[string]string{"message": "Course deleted successfully"}) // Return a 200 status code with a success message
}

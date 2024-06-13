// Package controllers provides the HTTP request handlers for student operations
package controllers

// Import necessary packages.
import (
	"github.com/labstack/echo/v4" // Import the Echo framework
	"net/http"                    // Import the net/http package for HTTP status codes
	"simpleGoProject/models"      // Import the models package containing the Student struct
	"simpleGoProject/services"    // Import the services package containing the StudentService interface
	"strconv"                     // Import the strconv package for string conversion
)

// StudentController defines the controller for handling student-related requests
type StudentController struct {
	studentService services.StudentService // Service to interact with student business logic
}

// NewStudentController creates a new instance of StudentController.
func NewStudentController(studentService services.StudentService) *StudentController {
	return &StudentController{studentService}
}

// CreateStudent handles the creation of a new student.
func (c *StudentController) CreateStudent(ctx echo.Context) error {
	var student models.Student
	if err := ctx.Bind(&student); err != nil { // Bind the request body to the student struct
		return ctx.JSON(http.StatusBadRequest, err) // Return a 400 status code for bad request
	}
	if err := ctx.Validate(&student); err != nil { // Validate the student struct based on validation rules
		return ctx.JSON(http.StatusBadRequest, err.Error()) // Return a 400 status code with the validation error message
	}
	err := c.studentService.CreateStudent(student) // Call the service to create a student
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err) // Return a 500 status code for internal server error
	}
	return ctx.JSON(http.StatusCreated, student) // Return a 201 status code with the created student data
}

// GetStudents handles retrieving all students.
func (c *StudentController) GetStudents(ctx echo.Context) error {
	students, err := c.studentService.GetStudents() // Call the service to get all students
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err) // Return a 500 status code for internal server error
	}
	return ctx.JSON(http.StatusOK, students) // Return a 200 status code with the students data
}

// GetStudentByID handles retrieving a student by ID.
func (c *StudentController) GetStudentByID(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id")) // Convert the URL parameter to an integer (student ID)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err) // Return a 400 status code for bad request
	}
	student, err := c.studentService.GetStudentByID(uint(id)) // Call the service to get a student by ID
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err) // Return a 500 status code for internal server error
	}
	return ctx.JSON(http.StatusOK, student) // Return a 200 status code with the student data
}

// UpdateStudent handles updating an existing student.
func (c *StudentController) UpdateStudent(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id")) // Convert the URL parameter to an integer (student ID)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err) // Return a 400 status code for bad request
	}
	var student models.Student
	if err := ctx.Bind(&student); err != nil { // Bind the request body to the student struct
		return ctx.JSON(http.StatusBadRequest, err) // Return a 400 status code for bad request
	}
	student.StudentID = uint(id)                   // Set the StudentID from the URL parameter
	if err := ctx.Validate(&student); err != nil { // Validate the updated student struct
		return ctx.JSON(http.StatusBadRequest, err.Error()) // Return a 400 status code with the validation error message
	}
	if err := c.studentService.UpdateStudent(student); err != nil { // Call the service to update a student.
		return ctx.JSON(http.StatusInternalServerError, err) // Return a 500 status code for internal server error
	}
	return ctx.JSON(http.StatusOK, student) // Return a 200 status code with the updated student data
}

// DeleteStudent handles deleting a student by ID.
func (c *StudentController) DeleteStudent(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id")) // Convert the URL parameter to an integer (student ID)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err) // Return a 400 status code for bad request
	}
	err = c.studentService.DeleteStudent(uint(id)) // Call the service to delete a student by ID
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err) // Return a 500 status code for internal server error
	}
	return ctx.JSON(http.StatusOK, map[string]string{"message": "Student deleted successfully"}) // Return a 200 status code with a success message
}

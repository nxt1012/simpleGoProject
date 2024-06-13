package routes

import (
	"github.com/labstack/echo/v4"
	"simpleGoProject/controllers"
)

func RegisterStudentRoutes(e *echo.Echo, studentController *controllers.StudentController) {
	e.POST("/students", studentController.CreateStudent)
	e.GET("/students", studentController.GetStudents)
	e.GET("/students/:id", studentController.GetStudentByID)
	e.PUT("/students/:id", studentController.UpdateStudent)
	e.DELETE("/students/:id", studentController.DeleteStudent)
}

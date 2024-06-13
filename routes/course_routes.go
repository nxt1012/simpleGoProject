package routes

import (
	"github.com/labstack/echo/v4"
	"simpleGoProject/controllers"
)

func RegisterCourseRoutes(e *echo.Echo, courseController *controllers.CourseController) {
	e.POST("/courses", courseController.CreateCourse)
	e.GET("/courses", courseController.GetCourses)
	e.GET("/courses/:id", courseController.GetCourseByID)
	e.PUT("/courses/:id", courseController.UpdateCourse)
	e.DELETE("/courses/:id", courseController.DeleteCourse)
}

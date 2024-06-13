// Package services provides the business logic layer for course operations
package services

// Import necessary packages
import (
	"simpleGoProject/models"       // Import the models package containing the Course struct
	"simpleGoProject/repositories" // Import the repositories package containing the CourseRepository interface
)

// CourseService defines the interface for course service operations
type CourseService interface {
	CreateCourse(course models.Course) error            // Creates a new course
	GetCourses() ([]models.Course, error)               // Retrieves all courses
	GetCourseByID(courseID uint) (models.Course, error) // Retrieves a course by ID
	UpdateCourse(course models.Course) error            // Updates an existing course
	DeleteCourse(courseID uint) error                   // Deletes a course by ID
}

// courseServiceImpl implements the CourseService interface
type courseServiceImpl struct {
	courseRepository repositories.CourseRepository
}

// NewCourseService creates a new instance of courseServiceImpl
func NewCourseService(courseRepository repositories.CourseRepository) CourseService {
	return &courseServiceImpl{courseRepository}
}

// CreateCourse creates a new course using the injected repository
func (s *courseServiceImpl) CreateCourse(course models.Course) error {
	return s.courseRepository.CreateCourse(course)
}

// GetCourses retrieves all courses using the injected repository
func (s *courseServiceImpl) GetCourses() ([]models.Course, error) {
	return s.courseRepository.GetCourses()
}

// GetCourseByID retrieves a course by its ID using the injected repository
func (s *courseServiceImpl) GetCourseByID(courseID uint) (models.Course, error) {
	return s.courseRepository.GetCourseByID(courseID)
}

// UpdateCourse updates an existing course using the injected repository
func (s *courseServiceImpl) UpdateCourse(course models.Course) error {
	return s.courseRepository.UpdateCourse(course)
}

// DeleteCourse deletes a course by its ID using the injected repository
func (s *courseServiceImpl) DeleteCourse(courseID uint) error {
	return s.courseRepository.DeleteCourse(courseID)
}

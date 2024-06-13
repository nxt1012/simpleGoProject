// Package repositories provides the data access layer for courses in the microservice
package repositories

// Import necessary packages
import (
	"gorm.io/gorm"           // Import the models package containing the Course struct
	"simpleGoProject/models" // Import the GORM library for ORM operations
)

// CourseRepository defines the interface for the course data operations
type CourseRepository interface {
	CreateCourse(course models.Course) error
	GetCourses() ([]models.Course, error)
	GetCourseByID(courseID uint) (models.Course, error)
	UpdateCourse(course models.Course) error
	DeleteCourse(courseID uint) error
}

// courseRepository is the struct that implements the CourseRepository interface
type courseRepository struct {
	db *gorm.DB
}

// NewCourseRepository create a new instance of courseRepository
func NewCourseRepository(db *gorm.DB) CourseRepository {
	return &courseRepository{db}
}

// CreateCourse inserts a new course record into the database
func (r *courseRepository) CreateCourse(course models.Course) error {
	return r.db.Create(&course).Error
}

// GetCourses retrieves all course records from the database
func (r *courseRepository) GetCourses() ([]models.Course, error) {
	var courses []models.Course
	err := r.db.Find(&courses).Error
	return courses, err
}

// GetCourseByID retrieves a course record by its ID from the database
func (r *courseRepository) GetCourseByID(courseID uint) (models.Course, error) {
	var course models.Course
	err := r.db.First(&course, courseID).Error
	return course, err
}

// UpdateCourse updates an existing course record in the database
func (r *courseRepository) UpdateCourse(course models.Course) error {
	return r.db.Save(&course).Error
}

// DeleteCourse deletes a course record by its ID from the database
func (r *courseRepository) DeleteCourse(courseID uint) error {
	return r.db.Delete(&models.Course{}, courseID).Error
}

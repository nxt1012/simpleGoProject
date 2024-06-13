// Package repositories provides the data access layer for the application
package repositories

// Import the required packages
import (
	"gorm.io/gorm"           // provides the ORM functionality
	"simpleGoProject/models" // imports the models packages containing the Student struct
)

// StudentRepository defines the interface for student data operations
type StudentRepository interface {
	CreateStudent(student models.Student) error
	GetStudents() ([]models.Student, error)
	GetStudentByID(id uint) (models.Student, error)
	UpdateStudent(student models.Student) error
	DeleteStudent(id uint) error
}

// studentRepository is the struct that implements the StudentRepository interface
type studentRepository struct {
	db *gorm.DB // Database connection
}

// NewStudentRepository creates a new instance of studentRepository
func NewStudentRepository(db *gorm.DB) StudentRepository {
	return &studentRepository{db}
}

// CreateStudent ínserts a new student record into the database
func (r *studentRepository) CreateStudent(student models.Student) error {
	return r.db.Create(&student).Error
}

// GetStudents retrieves all student records from the database
func (r *studentRepository) GetStudents() ([]models.Student, error) {
	var students []models.Student
	err := r.db.Find(&students).Error
	return students, err
}

// GetStudentByID retrieves a student record by its ID from the database
func (r *studentRepository) GetStudentByID(id uint) (models.Student, error) {
	var student models.Student
	err := r.db.First(&student, id).Error
	return student, err
}

// UpdateStudent updates an existing student record in the database
func (r *studentRepository) UpdateStudent(student models.Student) error {
	return r.db.Save(&student).Error
}

// DeleteStudent deletes a student record by its ID from the database
func (r *studentRepository) DeleteStudent(id uint) error {
	return r.db.Delete(&models.Student{}, id).Error
}

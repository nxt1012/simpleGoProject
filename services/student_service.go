// Package services provides the business logic layer for student operations
package services

// Import neccessary packages
import (
	"simpleGoProject/models"       // Import the models package containing the Student struct
	"simpleGoProject/repositories" // Import the repositories package containing the StudentRepository interface
)

// StudentService defines the interface for student service operations
type StudentService interface {
	CreateStudent(student models.Student) error     // Creates a new student
	GetStudents() ([]models.Student, error)         // Retrieves all students
	GetStudentByID(id uint) (models.Student, error) // Retrieves a student by ID
	UpdateStudent(student models.Student) error     // Updates an existing student
	DeleteStudent(id uint) error                    // Deletes a student by ID
}

// StudentServiceImpl implements the StudentService interface
type StudentServiceImpl struct {
	studentRepository repositories.StudentRepository
}

// NewStudentService create a new instance of StudentServiceImpl
func NewStudentService(studentRepository repositories.StudentRepository) StudentService {
	return &StudentServiceImpl{studentRepository}
}

// CreateStudent create a new student using the injected repository
func (s *StudentServiceImpl) CreateStudent(student models.Student) error {
	return s.studentRepository.CreateStudent(student)
}

// GetStudents retrieves all students using the injected repository
func (s *StudentServiceImpl) GetStudents() ([]models.Student, error) {
	return s.studentRepository.GetStudents()
}

// GetStudentByID retrieves a student by ID using the injected repository
func (s *StudentServiceImpl) GetStudentByID(id uint) (models.Student, error) {
	return s.studentRepository.GetStudentByID(id)
}

// UpdateStudent updates an existing student using the injected repository
func (s *StudentServiceImpl) UpdateStudent(student models.Student) error {
	return s.studentRepository.UpdateStudent(student)
}

// DeleteStudent deletes a student by ID using the injected repository
func (s *StudentServiceImpl) DeleteStudent(id uint) error {
	return s.studentRepository.DeleteStudent(id)
}

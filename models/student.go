// Package models defines the data structures for the application models
package models

// Student represents the structure of a student record in the database
type Student struct {
	StudentID uint   `gorm:"primary_key;autoIncrement" json:"student_id"` // Primary key for the student
	FirstName string `json:"first_name" validate:"required,max=50"`       // First name of the student
	LastName  string `json:"last_name" validate:"required,max=50"`        // Last name of the student
	Email     string `json:"email" validate:"required,email"`             // Email address of the student
	Age       uint   `json:"age" validate:"required,gt=0,lt=130"`         // Age of the student
	CourseID  uint   `json:"course_id"`                                   // Foreign key referencing the Course
}

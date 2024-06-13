// Package models defines the data structures for the application models
package models

type Course struct {
	CourseID   uint   `gorm:"primaryKey;autoIncrement" json:"course_id"` // Primary key for the course
	CourseName string `gorm:"size:100" json:"course_name"`               // Name of the course
}

package domain

import "github.com/google/uuid"

type CourseRepository interface {
	Course(courseUUID uuid.UUID) (Course, error)
	CourseByID(courseID uint) (Course, error)
	Courses() ([]Course, error)
	CreateCourse(c *Course) error
	UpdateCourseByUUID(c *Course) error
	UpdateCourseByCode(c *Course) error
	DeleteCourse(courseUUID uuid.UUID) error
}

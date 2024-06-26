package database

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/sumelms/microservice-course/internal/course/domain"
	"github.com/sumelms/microservice-course/pkg/errors"
)

func NewCourseRepository(db *sqlx.DB) (CourseRepository, error) { //nolint: revive
	sqlStatements := make(map[string]*sqlx.Stmt)

	for queryName, query := range queriesCourse() {
		stmt, err := db.Preparex(query)
		if err != nil {
			return CourseRepository{}, errors.WrapErrorf(err, errors.ErrCodeUnknown,
				"error preparing statement %s", queryName)
		}
		sqlStatements[queryName] = stmt
	}

	return CourseRepository{
		statements: sqlStatements,
	}, nil
}

type CourseRepository struct {
	statements map[string]*sqlx.Stmt
}

func (r CourseRepository) statement(s string) (*sqlx.Stmt, error) {
	stmt, ok := r.statements[s]
	if !ok {
		return nil, errors.NewErrorf(errors.ErrCodeUnknown, "prepared statement %s not found", s)
	}
	return stmt, nil
}

// Course get the Course by given uuid.
func (r CourseRepository) Course(courseUUID uuid.UUID) (domain.Course, error) {
	stmt, err := r.statement(getCourse)
	if err != nil {
		return domain.Course{}, err
	}

	var course domain.Course
	if err := stmt.Get(&course, courseUUID); err != nil {
		return domain.Course{}, errors.WrapErrorf(err, errors.ErrCodeUnknown, "error getting course")
	}
	return course, nil
}

// Courses list all courses.
func (r CourseRepository) Courses() ([]domain.Course, error) {
	stmt, err := r.statement(listCourses)
	if err != nil {
		return []domain.Course{}, err
	}

	var cc []domain.Course
	if err := stmt.Select(&cc); err != nil {
		return []domain.Course{}, errors.WrapErrorf(err, errors.ErrCodeUnknown, "error getting course")
	}
	return cc, nil
}

// CreateCourse creates a new course.
func (r CourseRepository) CreateCourse(course *domain.Course) error {
	stmt, err := r.statement(createCourse)
	if err != nil {
		return err
	}

	args := []interface{}{
		course.Code,
		course.Name,
		course.Underline,
		course.Image,
		course.ImageCover,
		course.Excerpt,
		course.Description,
	}
	if err := stmt.Get(course, args...); err != nil {
		return errors.WrapErrorf(err, errors.ErrCodeUnknown, "error creating course")
	}
	return nil
}

// UpdateCourse update the given course by ID.
func (r CourseRepository) UpdateCourse(course *domain.Course) error {
	stmt, err := r.statement(updateCourse)
	if err != nil {
		return err
	}

	args := []interface{}{
		// set
		course.Code,
		course.Name,
		course.Underline,
		course.Image,
		course.ImageCover,
		course.Excerpt,
		course.Description,
		// where
		course.UUID,
	}
	if err := stmt.Get(course, args...); err != nil {
		return errors.WrapErrorf(err, errors.ErrCodeUnknown, "error updating course")
	}
	return nil
}

// DeleteCourse soft delete the course by given uuid.
func (r CourseRepository) DeleteCourse(course *domain.DeletedCourse) error {
	stmt, err := r.statement(deleteCourse)
	if err != nil {
		return err
	}

	args := []interface{}{
		course.UUID,
	}
	if err := stmt.Get(course, args...); err != nil {
		return errors.WrapErrorf(err, errors.ErrCodeUnknown, "error deleting course")
	}
	return nil
}

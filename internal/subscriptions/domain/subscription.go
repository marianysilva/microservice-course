package domain

import (
	"time"

	"github.com/google/uuid"
)

type Course struct {
	UUID uuid.UUID `db:"course_uuid" json:"uuid"`
	Name string    `db:"course_name" json:"name"`
}

type Matrix struct {
	UUID uuid.UUID `db:"matrix_uuid" json:"uuid"`
	Name string    `db:"matrix_name" json:"name"`
}

type Subscription struct {
	UUID       uuid.UUID  `json:"uuid"`
	UserUUID   uuid.UUID  `db:"user_uuid"    json:"user_uuid"`
	Course     Course     `db:"-" json:"course"`
	CourseUUID uuid.UUID  `json:"course_uuid"`
	Matrix     *Matrix    `db:"-" json:"matrix"`
	MatrixUUID *uuid.UUID `json:"matrix_uuid"`
	Role       string     `db:"role"       json:"role"`
	ExpiresAt  *time.Time `db:"expires_at" json:"expires_at"`
	CreatedAt  time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt  time.Time  `db:"updated_at" json:"updated_at"`
	DeletedAt  *time.Time `db:"deleted_at" json:"deleted_at"`
}

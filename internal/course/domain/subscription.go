package domain

import (
	"time"

	"github.com/google/uuid"
)

type Subscription struct {
	ID        uint       `json:"id"`
	UUID      uuid.UUID  `json:"uuid"`
	UserUUID  uuid.UUID  `db:"user_uuid"    json:"user_uuid"`
	CourseID  uint       `db:"course_id"  json:"course_id"`
	MatrixID  *uint      `db:"matrix_id"  json:"matrix_id"`
	Role      string     `db:"role"       json:"role"`
	ExpiresAt *time.Time `db:"expires_at" json:"expires_at"`
	CreatedAt time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt time.Time  `db:"updated_at" json:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at" json:"deleted_at"`
}

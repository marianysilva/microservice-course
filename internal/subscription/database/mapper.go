package database

import (
	"github.com/google/uuid"
	"github.com/sumelms/microservice-course/internal/subscription/domain"
)

func toDBModel(entity *domain.Subscription) Subscription {
	s := Subscription{
		UserID:   uuid.MustParse(entity.UserID),
		CourseID: uuid.MustParse(entity.CourseID),
		MatrixID: uuid.MustParse(entity.MatrixID),
	}

	if !entity.ValidUntil.IsZero() {
		s.ValidUntil = entity.ValidUntil
	}

	if entity.ID > 0 {
		// gorm.Model fields
		s.ID = entity.ID
		s.CreatedAt = entity.CreatedAt
		s.UpdatedAt = entity.UpdatedAt

		if !entity.DeletedAt.IsZero() {
			s.DeletedAt = entity.DeletedAt
		}
	}

	return s
}

func toDomainModel(entity *Subscription) domain.Subscription {
	return domain.Subscription{
		ID:         entity.ID,
		UserID:     entity.UserID.String(),
		CourseID:   entity.CourseID.String(),
		MatrixID:   entity.MatrixID.String(),
		ValidUntil: entity.ValidUntil,
		CreatedAt:  entity.CreatedAt,
		UpdatedAt:  entity.UpdatedAt,
		DeletedAt:  entity.DeletedAt,
	}
}
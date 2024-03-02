package domain

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

type SubscriptionFilters struct {
	CourseUUID uuid.UUID `json:"course_uuid,omitempty"`
	UserUUID   uuid.UUID `json:"user_uuid,omitempty"`
}

func (s *Service) Subscription(_ context.Context, subscriptionUUID uuid.UUID) (Subscription, error) {
	sub, err := s.subscriptions.Subscription(subscriptionUUID)
	if err != nil {
		return Subscription{}, fmt.Errorf("service can't find subscription: %w", err)
	}

	return sub, nil
}

func (s *Service) Subscriptions(_ context.Context, filters *SubscriptionFilters) ([]Subscription, error) {
	list, err := func() ([]Subscription, error) {
		if filters != nil {
			if filters.UserUUID != uuid.Nil {
				return s.subscriptions.UserSubscriptions(filters.UserUUID)
			}

			if filters.CourseUUID != uuid.Nil {
				return s.subscriptions.CourseSubscriptions(filters.CourseUUID)
			}
		}

		return s.subscriptions.Subscriptions()
	}()
	if err != nil {
		return []Subscription{}, fmt.Errorf("service didn't found any subscription: %w", err)
	}

	return list, nil
}

func (s *Service) CreateSubscription(_ context.Context, sub *Subscription) error {
	// _, err := s.courses.Course(sub.Course.UUID)
	// if err != nil {
	// 	return fmt.Errorf("error checking if course %d exists: %w", sub.Course.UUID, err)
	// }

	if err := s.subscriptions.CreateSubscription(sub); err != nil {
		return fmt.Errorf("service can't create subscription: %w", err)
	}

	return nil
}

func (s *Service) UpdateSubscription(_ context.Context, sub *Subscription) error {
	if err := s.subscriptions.UpdateSubscription(sub); err != nil {
		return fmt.Errorf("service can't update subscription: %w", err)
	}

	return nil
}

func (s *Service) DeleteSubscription(_ context.Context, subscriptionUUID uuid.UUID) error {
	if err := s.subscriptions.DeleteSubscription(subscriptionUUID); err != nil {
		return fmt.Errorf("service can't delete subscription: %w", err)
	}

	return nil
}

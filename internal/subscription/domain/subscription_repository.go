package domain

import "github.com/google/uuid"

type SubscriptionRepository interface {
	// Create
	CreateSubscription(subscription *Subscription) error
	CreateSubscriptionWithoutMatrix(subscription *Subscription) error
	// Read
	Subscription(subscriptionUUID uuid.UUID) (*Subscription, error)
	Subscriptions() ([]*Subscription, error)
	CourseSubscriptions(courseUUID uuid.UUID) ([]*Subscription, error)
	UserSubscriptions(userUUID uuid.UUID) ([]*Subscription, error)
	// Update
	UpdateSubscription(subscription *Subscription) error
	// Delete
	DeleteSubscription(subscription *DeletedSubscription) error
}

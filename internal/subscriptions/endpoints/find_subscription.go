package endpoints

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/sumelms/microservice-course/internal/subscriptions/domain"
)

type findSubscriptionRequest struct {
	UUID uuid.UUID `json:"uuid"`
}

type findSubscriptionResponse struct {
	UUID      uuid.UUID      `json:"uuid"`
	UserUUID  uuid.UUID      `json:"user_uuid"`
	Course    domain.Course  `json:"course"`
	Matrix    *domain.Matrix `json:"matrix,omitempty"`
	Role      string         `json:"role"`
	ExpiresAt *time.Time     `json:"expires_at,omitempty"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}

func NewFindSubscriptionHandler(s domain.ServiceInterface, opts ...kithttp.ServerOption) *kithttp.Server {
	return kithttp.NewServer(
		makeFindSubscriptionEndpoint(s),
		decodeFindSubscriptionRequest,
		encodeFindSubscriptionResponse,
		opts...,
	)
}

func makeFindSubscriptionEndpoint(s domain.ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(findSubscriptionRequest)
		if !ok {
			return nil, fmt.Errorf("invalid argument")
		}

		sub, err := s.Subscription(ctx, req.UUID)
		if err != nil {
			return nil, err
		}

		return &findSubscriptionResponse{
			UUID:      sub.UUID,
			UserUUID:  sub.UserUUID,
			Course:    sub.Course,
			Matrix:    sub.Matrix,
			Role:      sub.Role,
			ExpiresAt: sub.ExpiresAt,
			CreatedAt: sub.CreatedAt,
			UpdatedAt: sub.UpdatedAt,
		}, nil
	}
}

func decodeFindSubscriptionRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	subscriptionUUID, ok := vars["uuid"]
	if !ok {
		return nil, fmt.Errorf("invalid argument")
	}

	UUID := uuid.MustParse(subscriptionUUID)

	return findSubscriptionRequest{UUID: UUID}, nil
}

func encodeFindSubscriptionResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return kithttp.EncodeJSONResponse(ctx, w, response)
}

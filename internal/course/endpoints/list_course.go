package endpoints

import (
	"context"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/sumelms/microservice-course/internal/course/domain"
)

type listCourseResponse struct {
	Courses []findCourseResponse `json:"courses"`
}

func NewListCourseHandler(s domain.Service, opts ...kithttp.ServerOption) *kithttp.Server {
	return kithttp.NewServer(
		makeListCourseEndpoint(s),
		decodeListCourseRequest,
		encodeListCourseResponse,
		opts...,
	)
}

func makeListCourseEndpoint(s domain.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		courses, err := s.ListCourse(ctx)
		if err != nil {
			return nil, err
		}

		var list []findCourseResponse
		for _, c := range courses {
			list = append(list, findCourseResponse{
				UUID:        c.UUID,
				Title:       c.Title,
				Subtitle:    c.Subtitle,
				Excerpt:     c.Excerpt,
				Description: c.Description,
				CreatedAt:   c.CreatedAt,
				UpdatedAt:   c.UpdatedAt,
			})
		}

		return &listCourseResponse{Courses: list}, nil
	}
}

func decodeListCourseRequest(_ context.Context, _ *http.Request) (interface{}, error) {
	return nil, nil
}

func encodeListCourseResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return kithttp.EncodeJSONResponse(ctx, w, response)
}
package course

import (
	"github.com/go-kit/log"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/sumelms/microservice-course/internal/course/database"
	"github.com/sumelms/microservice-course/internal/course/domain"
	"github.com/sumelms/microservice-course/internal/course/transport/http"
)

func NewService(db *sqlx.DB, logger log.Logger) (*domain.Service, error) {
	course, err := database.NewCourseRepository(db)
	if err != nil {
		return nil, err
	}

	service, err := domain.NewService(
		domain.WithLogger(logger),
		domain.WithCourseRepository(course))
	if err != nil {
		return nil, err
	}
	return service, nil
}

func NewHTTPService(router *mux.Router, service domain.ServiceInterface, logger log.Logger) error {
	http.NewHTTPHandler(router, service, logger)
	return nil
}

package matrix

import (
	"github.com/go-kit/log"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/sumelms/microservice-course/internal/matrix/database"
	"github.com/sumelms/microservice-course/internal/matrix/domain"
	"github.com/sumelms/microservice-course/internal/matrix/transport/http"
	"github.com/sumelms/microservice-course/internal/shared"
)

func NewService(db *sqlx.DB, amqpConnection *shared.AMQPConnection, course domain.CourseClient, logger log.Logger) (*domain.Service, error) {
	matrix, err := database.NewMatrixRepository(db)
	if err != nil {
		return nil, err
	}
	subject, err := database.NewSubjectRepository(db)
	if err != nil {
		return nil, err
	}

	matricesQueue, err := amqpConnection.NewAMQPQueue("MatricesQueue")
	if err != nil {
		return nil, err
	}

	service, err := domain.NewService(
		domain.WithLogger(logger),
		domain.WithQueue(matricesQueue),
		domain.WithMatrixRepository(matrix),
		domain.WithSubjectRepository(subject),
		domain.WithCourseClient(course))
	if err != nil {
		return nil, err
	}
	return service, nil
}

func NewHTTPService(router *mux.Router, service domain.ServiceInterface, logger log.Logger) error {
	http.NewHTTPHandler(router, service, logger)
	return nil
}

package tests

import (
	"fmt"
	"regexp"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/sumelms/microservice-course/tests/database"
)

var (
	Now              = time.Now()
	CourseID         = uint(1)
	CourseUUID       = uuid.MustParse("e8276e31-9a87-4cf1-a16c-080f9c5790d1")
	CourseName       = "CCO"
	SubscriptionID   = uint(1)
	SubscriptionUUID = uuid.MustParse("dd7c915b-849a-4ba4-bc09-aeecd95c40cc")
	UserUUID         = uuid.MustParse("ef2bc01e-be93-4a1f-9e96-c78d3d432088")
	MatrixID         = uint(1)
	MatrixUUID       = uuid.MustParse("825f630d-7604-474d-9b19-5431c09d0c3f")
	MatrixName       = "CCO 2019"
	EmptyRows        = sqlmock.NewRows([]string{})
)

func NewTestDB(queries map[string]string) (*sqlx.DB, sqlmock.Sqlmock, map[string]*sqlmock.ExpectedPrepare) {
	db, mock := database.NewDBMock()

	sqlStatements := make(map[string]*sqlmock.ExpectedPrepare)
	for queryName, query := range queries {
		stmt := mock.ExpectPrepare(fmt.Sprintf("^%s$", regexp.QuoteMeta(query)))
		sqlStatements[queryName] = stmt
	}

	mock.MatchExpectationsInOrder(false)

	return db, mock, sqlStatements
}

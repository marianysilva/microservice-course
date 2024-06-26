package database

import (
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	"github.com/sumelms/microservice-course/internal/matrix/domain"
	utils "github.com/sumelms/microservice-course/tests"
)

var (
	matrix = domain.Matrix{
		UUID:        utils.MatrixUUID,
		CourseUUID:  utils.CourseUUID,
		Code:        "Code",
		Name:        "Name",
		Description: "Description",
		CreatedAt:   utils.Now,
		UpdatedAt:   utils.Now,
	}
)

func newMatrixTestDB() (*sqlx.DB, sqlmock.Sqlmock, map[string]*sqlmock.ExpectedPrepare) {
	return utils.NewTestDB(queriesMatrix())
}

func TestRepository_Matrix(t *testing.T) {
	validRows := sqlmock.NewRows([]string{
		"course_uuid", "uuid", "code", "name", "description",
		"created_at", "updated_at"}).
		AddRow(
			matrix.CourseUUID, matrix.UUID, matrix.Code, matrix.Name, matrix.Description,
			matrix.CreatedAt, matrix.UpdatedAt)

	type args struct {
		UUID uuid.UUID
	}

	tests := []struct {
		name    string
		args    args
		rows    *sqlmock.Rows
		want    domain.Matrix
		wantErr bool
	}{
		{
			name:    "get matrix",
			args:    args{UUID: utils.MatrixUUID},
			rows:    validRows,
			want:    matrix,
			wantErr: false,
		},
		{
			name:    "matrix not found error",
			args:    args{UUID: uuid.MustParse("8281f61e-956e-4f64-ac0e-860c444c5f86")},
			rows:    utils.EmptyRows,
			want:    domain.Matrix{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			db, _, stmts := newMatrixTestDB()
			r, err := NewMatrixRepository(db)
			if err != nil {
				t.Fatalf("an error '%s' was not expected when creating the MatrixRepository", err)
			}
			prep, ok := stmts[getMatrix]
			if !ok {
				t.Fatalf("prepared statement %s not found", getMatrix)
			}

			prep.ExpectQuery().WithArgs(utils.MatrixUUID).WillReturnRows(tt.rows)

			got, err := r.Matrix(tt.args.UUID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Matrix() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Matrix() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_CourseMatrixExists(t *testing.T) {
	type args struct {
		courseUUID uuid.UUID
		matrixUUID uuid.UUID
	}

	tests := []struct {
		name    string
		args    args
		rows    *sqlmock.Rows
		want    bool
		wantErr bool
	}{
		{
			name:    "matrix course exists",
			args:    args{courseUUID: utils.CourseUUID, matrixUUID: utils.MatrixUUID},
			want:    true,
			wantErr: false,
		},
		{
			name:    "matrix course not found error",
			args:    args{courseUUID: utils.CourseUUID, matrixUUID: uuid.MustParse("8281f61e-956e-4f64-ac0e-860c444c5f86")},
			want:    false,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			db, _, stmts := newMatrixTestDB()
			r, err := NewMatrixRepository(db)
			if err != nil {
				t.Fatalf("an error '%s' was not expected when creating the MatrixRepository", err)
			}
			prep, ok := stmts[getCourseMatrixExists]
			if !ok {
				t.Fatalf("prepared statement %s not found", getCourseMatrixExists)
			}

			columns := []string{"exists_relationship"}
			prep.ExpectQuery().WithArgs(tt.args.courseUUID, tt.args.matrixUUID).WillReturnRows(sqlmock.NewRows(columns).AddRow(tt.want))

			got, err := r.CourseMatrixExists(tt.args.courseUUID, tt.args.matrixUUID)
			if (err != nil) != tt.wantErr {
				t.Errorf("CourseMatrixExists() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CourseMatrixExists() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_Matrices(t *testing.T) {
	validRows := sqlmock.NewRows([]string{
		"course_uuid", "uuid", "code", "name", "description",
		"created_at", "updated_at"}).
		AddRow(
			matrix.CourseUUID, matrix.UUID, matrix.Code, matrix.Name, matrix.Description,
			matrix.CreatedAt, matrix.UpdatedAt).
		AddRow(matrix.CourseUUID, uuid.MustParse("e74868b2-72d4-4591-a90d-122a9ac2d945"),
			matrix.Code, matrix.Name, matrix.Description,
			matrix.CreatedAt, matrix.UpdatedAt)

	tests := []struct {
		name    string
		rows    *sqlmock.Rows
		wantLen int
		wantErr bool
	}{
		{
			name:    "get all matrices",
			rows:    validRows,
			wantLen: 2,
			wantErr: false,
		},
		{
			name:    "get no matrices",
			rows:    utils.EmptyRows,
			wantLen: 0,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			db, _, stmts := newMatrixTestDB()
			r, err := NewMatrixRepository(db)
			if err != nil {
				t.Fatalf("an error '%s' was not expected when creating the MatrixRepository", err)
			}
			prep, ok := stmts[listMatrices]
			if !ok {
				t.Fatalf("prepared statement %s not found", listMatrices)
			}

			prep.ExpectQuery().WillReturnRows(tt.rows)

			got, err := r.Matrices()
			if (err != nil) != tt.wantErr {
				t.Errorf("Matrices() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != tt.wantLen {
				t.Errorf("Matrices() got = %v, want %v", got, tt.wantLen)
			}
		})
	}
}

func TestRepository_CreateMatrix(t *testing.T) {
	validRows := sqlmock.NewRows([]string{
		"course_uuid", "uuid", "code", "name", "description",
		"created_at", "updated_at"}).
		AddRow(
			matrix.CourseUUID, matrix.UUID, matrix.Code, matrix.Name, matrix.Description,
			matrix.CreatedAt, matrix.UpdatedAt)

	type args struct {
		m *domain.Matrix
	}

	tests := []struct {
		name    string
		rows    *sqlmock.Rows
		args    args
		want    domain.Matrix
		wantErr bool
	}{
		{
			name:    "create matrix",
			rows:    validRows,
			args:    args{m: &matrix},
			want:    matrix,
			wantErr: false,
		},
		{
			name:    "empty fields",
			rows:    utils.EmptyRows,
			args:    args{m: &matrix},
			want:    domain.Matrix{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			db, _, stmts := newMatrixTestDB()
			r, err := NewMatrixRepository(db)
			if err != nil {
				t.Fatalf("an error '%s' was not expected when creating the repository", err)
			}
			prep, ok := stmts[createMatrix]
			if !ok {
				t.Fatalf("prepared statement %s not found", createMatrix)
			}

			prep.ExpectQuery().WillReturnRows(tt.rows)

			if err := r.CreateMatrix(tt.args.m); (err != nil) != tt.wantErr {
				t.Errorf("CreateMatrix() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr && !reflect.DeepEqual(*tt.args.m, tt.want) {
				t.Errorf("CreateMatrix() got = %v, want %v", *tt.args.m, tt.want)
			}
		})
	}
}

func TestRepository_UpdateMatrix(t *testing.T) {
	validRows := sqlmock.NewRows([]string{
		"course_uuid", "uuid", "code", "name", "description",
		"created_at", "updated_at"}).
		AddRow(
			matrix.CourseUUID, matrix.UUID, matrix.Code, matrix.Name, matrix.Description,
			matrix.CreatedAt, matrix.UpdatedAt)

	type args struct {
		m *domain.Matrix
	}
	tests := []struct {
		name    string
		args    args
		rows    *sqlmock.Rows
		want    domain.Matrix
		wantErr bool
	}{
		{
			name:    "update matrix",
			args:    args{m: &matrix},
			rows:    validRows,
			want:    matrix,
			wantErr: false,
		},
		{
			name:    "empty matrix",
			args:    args{m: &domain.Matrix{}},
			rows:    utils.EmptyRows,
			want:    domain.Matrix{},
			wantErr: true,
		},
	}
	for _, testCase := range tests {
		tt := testCase
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			db, _, stmts := newMatrixTestDB()
			r, err := NewMatrixRepository(db)
			if err != nil {
				t.Fatalf("an error '%s' was not expected creating the MatrixRepository", err)
			}
			prep, ok := stmts[updateMatrix]
			if !ok {
				t.Fatalf("prepared statement %s not found", updateMatrix)
			}
			prep.ExpectExec().WillReturnResult(sqlmock.NewResult(1, 1))

			prep, ok = stmts[getMatrix]
			if !ok {
				t.Fatalf("prepared statement %s not found", getMatrix)
			}
			prep.ExpectQuery().WillReturnRows(tt.rows)

			if err := r.UpdateMatrix(tt.args.m); (err != nil) != tt.wantErr {
				t.Errorf("UpdateMatrix() \nerror = %v, \nwantErr = %v", err, tt.wantErr)
			}

			if !tt.wantErr && !reflect.DeepEqual(*tt.args.m, tt.want) {
				t.Errorf("UpdateMatrix() \ngot = %v, \nwant = %v", *tt.args.m, tt.want)
			}
		})
	}
}

func TestRepository_DeleteMatrix(t *testing.T) {
	validRows := sqlmock.NewRows([]string{"uuid", "deleted_at"}).
		AddRow(matrix.UUID, utils.Now)

	deletedMatrix := domain.DeletedMatrix{
		UUID:      utils.MatrixUUID,
		DeletedAt: utils.Now,
	}

	type args struct {
		m *domain.DeletedMatrix
	}
	tests := []struct {
		name    string
		args    args
		rows    *sqlmock.Rows
		want    domain.DeletedMatrix
		wantErr bool
	}{
		{
			name:    "delete matrix",
			args:    args{m: &deletedMatrix},
			rows:    validRows,
			want:    deletedMatrix,
			wantErr: false,
		},
		{
			name:    "empty matrix",
			args:    args{m: &domain.DeletedMatrix{}},
			rows:    utils.EmptyRows,
			want:    domain.DeletedMatrix{},
			wantErr: true,
		},
	}
	for _, testCase := range tests {
		tt := testCase
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			db, _, stmts := newMatrixTestDB()
			r, err := NewMatrixRepository(db)
			if err != nil {
				t.Fatalf("an error '%s' was not expected creating the MatrixRepository", err)
			}

			prep, ok := stmts[deleteMatrix]
			if !ok {
				t.Fatalf("prepared statement %s not found", deleteMatrix)
			}
			prep.ExpectQuery().WillReturnRows(tt.rows)

			if err := r.DeleteMatrix(tt.args.m); (err != nil) != tt.wantErr {
				t.Errorf("DeleteMatrix() \nerror = %v, \nwantErr = %v", err, tt.wantErr)
			}

			if !tt.wantErr && !reflect.DeepEqual(*tt.args.m, tt.want) {
				t.Errorf("DeleteMatrix() \ngot = %v, \nwant = %v", *tt.args.m, tt.want)
			}
		})
	}
}

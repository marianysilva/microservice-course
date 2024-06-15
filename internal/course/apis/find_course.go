package apis

import (
	"context"
	"fmt"
	"reflect"

	"github.com/google/uuid"
	"github.com/sumelms/microservice-course/internal/course/domain"
)

func isEmptyCourse(c domain.Course) bool {
	val := reflect.ValueOf(c)
	for i := 0; i < val.NumField(); i++ {
		if !reflect.DeepEqual(val.Field(i).Interface(), reflect.Zero(val.Field(i).Type()).Interface()) {
			return false
		}
	}
	return true
}

func IsActiveCourse(s domain.ServiceInterface, courseUUID string) (bool, error) {
	cUUID, err := uuid.Parse(courseUUID)

	if err != nil {
		return false, fmt.Errorf("invalid argument")
	}

	course, _ := s.Course(context.Background(), cUUID)

	if !isEmptyCourse(course) {
		return true, nil
	}

	return false, nil
}

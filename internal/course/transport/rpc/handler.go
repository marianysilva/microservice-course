package rpc

import (
	"context"

	"github.com/sumelms/microservice-course/internal/course/apis"
	"github.com/sumelms/microservice-course/internal/course/domain"
	"github.com/sumelms/microservice-course/internal/course/transport/rpc/pb"
	"github.com/sumelms/microservice-course/internal/shared"
)

type GRPCServer struct {
	service domain.ServiceInterface
	pb.UnimplementedCoursesServer
}

func (s *GRPCServer) IsActiveCourse(ctx context.Context, in *pb.CourseRequest) (*pb.CourseResponse, error) {
	isActive, err := apis.IsActiveCourse(s.service, in.GetCourseUUID())
	if err != nil {
		return nil, err
	}

	return &pb.CourseResponse{
		IsActive: isActive,
	}, nil
}

func NewRPCHandler(server *shared.RPCServer, service domain.ServiceInterface) {
	grpcServer := GRPCServer{
		service: service,
	}
	pb.RegisterCoursesServer(server.Server, &grpcServer)
}

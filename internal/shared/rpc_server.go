package shared

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"

	"github.com/sumelms/microservice-course/pkg/config"
	"github.com/sumelms/microservice-course/pkg/logger"
)

type RPCServer struct {
	Server   *grpc.Server
	Listener *net.Listener
	logger   logger.Logger
	Config   *config.RPCServer
}

func NewRPCServer(cfg *config.RPCServer, logger logger.Logger) (*RPCServer, error) {
	if cfg == nil {
		return nil, fmt.Errorf("invalid server config")
	}

	listener, err := net.Listen("tcp", cfg.Host)
	if err != nil {
		return nil, fmt.Errorf("unnable to create net listener")
	}

	server := grpc.NewServer()

	return &RPCServer{
		Server:   server,
		Listener: &listener,
		logger:   logger,
		Config:   cfg,
	}, nil
}

func (s *RPCServer) Start() error {
	s.logger.Log("msg", "starting RPC Server", "host", s.Config.Host)

	if err := s.Server.Serve(*s.Listener); err != nil {
		s.logger.Log("msg", "unable to start RPC Server", "error", err)
		return err
	}

	return nil
}

func (s *RPCServer) Stop(ctx context.Context) {
	s.logger.Log("msg", "RPC Server started to shutdown")
	s.Server.Stop()
	s.logger.Log("msg", "RPC Server shutdown successfully")
}

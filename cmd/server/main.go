package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/sumelms/microservice-course/internal/course"
	"github.com/sumelms/microservice-course/internal/matrix"
	"github.com/sumelms/microservice-course/internal/shared"
	"github.com/sumelms/microservice-course/internal/shared/clients"
	"github.com/sumelms/microservice-course/internal/subscription"
	"github.com/sumelms/microservice-course/pkg/config"
	database "github.com/sumelms/microservice-course/pkg/database/postgres"
	log "github.com/sumelms/microservice-course/pkg/logger"
	"github.com/sumelms/microservice-course/swagger"
	"golang.org/x/sync/errgroup"
)

var (
	logger     log.Logger
	httpServer *shared.Server
)

//nolint:funlen
func main() {
	// Logger
	logger = log.NewLogger()
	logger.Log("msg", "service started")

	// Configuration
	cfg, err := loadConfig()
	if err != nil {
		logger.Log("msg", "exit", "error", err)
		os.Exit(-1)
	}

	// Database
	db, err := database.Connect(cfg.Database)
	if err != nil {
		logger.Log("msg", "database error", "error", err)
		os.Exit(1)
	}

	// Initialize the domain services
	svcLogger := logger.With("component", "service")

	courseSvc, err := course.NewService(db, svcLogger.Logger())
	if err != nil {
		logger.Log("msg", "unable to start course service", "error", err)
		os.Exit(1)
	}

	matrixSvc, err := matrix.NewService(db, svcLogger.Logger(), clients.NewCourseClient(courseSvc))
	if err != nil {
		logger.Log("msg", "unable to start matrix service", "error", err)
		os.Exit(1)
	}

	subscriptionSvc, err := subscription.NewService(
		db, svcLogger.Logger(),
		clients.NewCourseClient(courseSvc),
		clients.NewMatrixClient(matrixSvc))
	if err != nil {
		logger.Log("msg", "unable to start subscription service", "error", err)
		os.Exit(1)
	}

	// Gracefully shutdown
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(interrupt)

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		// Initialize the router
		router := mux.NewRouter().StrictSlash(true)
		// Global Middlewares
		router.Use(shared.CorsMiddleware)

		// Register Swagger handler
		swagger.Register(router)

		// Initializing the HTTP Services
		httpLogger := logger.With("component", "http")

		if err = course.NewHTTPService(router, courseSvc, httpLogger.Logger()); err != nil {
			logger.Log("msg", "unable to start a service: course", "error", err)
			return err
		}

		if err = matrix.NewHTTPService(router, matrixSvc, httpLogger.Logger()); err != nil {
			logger.Log("msg", "unable to start a service: matrix", "error", err)
			return err
		}

		if err = subscription.NewHTTPService(router, subscriptionSvc, httpLogger.Logger()); err != nil {
			logger.Log("msg", "unable to start a service: subscription", "error", err)
			return err
		}

		// Create the HTTP Server
		httpServer, err = shared.NewServer(cfg.Server.HTTP, router, httpLogger)
		if err != nil {
			return err
		}

		return httpServer.Start()
	})

	select {
	case <-interrupt:
		break
	case <-ctx.Done():
		break
	}

	logger.Log("msg", "received shutdown signal")

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()

	if httpServer != nil {
		httpServer.Stop(shutdownCtx)
	}

	if err := g.Wait(); err != nil {
		logger.Log("msg", "server returning an error", "error", err)
		defer os.Exit(2)
	}

	logger.Log("msg", "service ended")
}

func loadConfig() (*config.Config, error) {
	// Configuration
	configPath := os.Getenv("SUMELMS_CONFIG_PATH")
	if configPath == "" {
		configPath = "./config.yml"
	}

	cfg, err := config.NewConfig(configPath)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

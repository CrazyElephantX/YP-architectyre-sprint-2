package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"cinemaabyss/proxy/internal/controller/handlers"
	"cinemaabyss/proxy/internal/controller/router"
	"cinemaabyss/proxy/internal/controller/server"
	"cinemaabyss/proxy/internal/lib/config"

	"github.com/sirupsen/logrus"
)

var (
	cfg config.Config
)

func init() {

	flag.StringVar(&cfg.Port, "port", "8000", "port")
	flag.StringVar(&cfg.MonolithURL, "monolith-URL", "http://localhost:8080", "monolith URL")
	flag.StringVar(&cfg.MoviesServiceURL, "MOVIES_SERVICE_URL", "http://localhost:8081", "movies service URL")
	flag.StringVar(&cfg.EventsServiceURL, "EVENTS_SERVICE_URL", "http://localhost:8082", "events service URL")
	flag.BoolVar(&cfg.GradualMigration, "GRADUAL_MIGRATION", false, "gradual migration")
	flag.StringVar(&cfg.MoviesMigrationPercent, "MOVIES_MIGRATION_PERCENT", "50", "movies migration percent")

	flag.Parse()

	if envPort := os.Getenv("PORT"); envPort != "" {
		cfg.Port = envPort
	}
	if envMonolithURL := os.Getenv("MONOLITH_URL"); envMonolithURL != "" {
		cfg.MonolithURL = envMonolithURL
	}
	if envMoviesServiceURL := os.Getenv("MOVIES_SERVICE_URL"); envMoviesServiceURL != "" {
		cfg.MoviesServiceURL = envMoviesServiceURL
	}
	if envEventsServiceURL := os.Getenv("EVENTS_SERVICE_URL"); envEventsServiceURL != "" {
		cfg.EventsServiceURL = envEventsServiceURL
	}

	if envMoviesMigrationPercent := os.Getenv("MOVIES_MIGRATION_PERCENT"); envMoviesMigrationPercent != "" {
		cfg.MoviesMigrationPercent = envMoviesMigrationPercent
	}

	if envGradualMigration := os.Getenv("GRADUAL_MIGRATION"); envGradualMigration != "" {
		b := false
		b, err := strconv.ParseBool(envGradualMigration)
		if err != nil {
			logrus.Info("Ошибка при преобразовании GRADUAL_MIGRATION в bool:", err)
		}
		cfg.GradualMigration = b
	}
}

func main() {
	Run()
}

func Run() {

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)
	defer cancel()

	registerHandlers, err := initializeServices()
	if err != nil {
		logrus.Fatalf("Error initializing services: %v", err)
	}

	appRouter := router.NewRouter(registerHandlers, cfg)
	host := fmt.Sprintf(":%s", cfg.Port)
	appServer := server.NewHTTPServer(host, appRouter.Echo)

	logrus.Infof("Starting API Gateway server: %s", host)

	go appServer.Start(ctx)

	<-ctx.Done()
	appServer.Stop(ctx)
}

func initializeServices() ([]handlers.Handler, error) {
	var registerHandlers []handlers.Handler
	return registerHandlers, nil
}

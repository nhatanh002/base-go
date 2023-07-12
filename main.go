package main

import (
	"base-go/adapter/repositories"
	cats_repo "base-go/adapter/repositories/cats-repo"
	"base-go/application"
	"base-go/application/cats"
	"base-go/common/config"
	"base-go/common/logger"
	gw_http "base-go/gateway/http"
	"base-go/migrations"
	"base-go/services"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	mainEcho()
}

func mainEcho() {
	logger.Init()
	cnf := config.Get()

	logger.Info("Initializing...")
	gormdb := repositories.NewGormdb(cnf)
	migrations.Migrate(gormdb)

	logger.Info("Constructing dependencies...")
	catRepo := cats_repo.NewCatsRepo(gormdb)
	catsService := services.NewCatsService(catRepo)
	catsInteractor := cats.NewCatsInteractor(catsService)
	app := application.NewApp(catsInteractor)

	logger.Info("Constructing http server...")
	router := gw_http.EchoRouter(cnf, app)
	server := gw_http.NewHttpServer(cnf, router)
	// Server run context
	serverCtx, serverStopCtx := context.WithCancel(context.Background())

	// Listen for syscall signals for process to interrupt/quit
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	logger.Info("Setup OS signal handler...")
	go func() {
		<-sig

		// Shutdown signal with grace period of 30 seconds
		shutdownCtx, _ := context.WithTimeout(serverCtx, 30*time.Second)

		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				logger.Error("graceful shutdown timed out.. forcing exit, error: %s", shutdownCtx.Err().Error())
			}
		}()

		// Trigger graceful shutdown
		err := server.Shutdown(shutdownCtx)
		if err != nil {
			log.Fatal(err)
		}
		serverStopCtx()
	}()

	// Run the server
	logger.Info("Starting http server at %s:%d", cnf.HttpConfig.Host, cnf.HttpConfig.Port)
	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}

	// Wait for server context to be stopped
	logger.Info("Waiting for http server to exit...")
	<-serverCtx.Done()
	logger.Info("Goodbye.")
}

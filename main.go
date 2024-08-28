package main

import (
	"context"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"to-do-app/config"
	"to-do-app/logger"
	"to-do-app/middleware"
	"to-do-app/repository"
	"to-do-app/route"
)

func init() {
	logger.Log = logger.NewZerolog()
	config.Config = config.NewConfig()
}

func main() {
	repo, err := repository.NewDBConnection(config.Config)
	if err != nil {
		logger.Log.Error(err)
		return
	}

	repository.DB = repo

	r := mux.NewRouter()

	r.Use(middleware.RecoveryMiddleware)
	r.Use(middleware.LoggingMiddleware)
	r.Use(middleware.JSONMiddleware)

	r.HandleFunc("/tasks", route.CreateTasks).Methods("POST")

	r.HandleFunc("/tasks", route.GetTasks).Methods("GET")
	r.HandleFunc("/tasks/{id}", route.GetTask).Methods("GET")

	r.HandleFunc("/tasks/{id}", route.UpdateTask).Methods("PUT")

	r.HandleFunc("/tasks/{id}", route.DeleteTask).Methods("DELETE")

	srv := &http.Server{
		Addr:         "0.0.0.0:8080",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}

	logger.Log.Info("Application starting up...")

	go func() {
		logger.Log.Info("HTTP server starting up...")
		if err := srv.ListenAndServe(); err != nil {
			logger.Log.Error(err)
			os.Exit(0)
		}
	}()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	srv.Shutdown(ctx)

	sql, _ := repository.DB.DB()
	sql.Close()

	logger.Log.Info("Shutting down...")

	os.Exit(0)
}

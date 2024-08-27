package main

import (
	"context"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"to-do-app/logger"
	"to-do-app/route"
)

func init() {
	logger.Log = logger.NewZerolog()
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/tasks", route.CreateTasks).Methods("POST")

	srv := &http.Server{
		Addr:         "0.0.0.0:8080",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			logger.Log.Info(err)
		}
	}()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	srv.Shutdown(ctx)

	logger.Log.Info("Shutting down...")

	os.Exit(0)
}

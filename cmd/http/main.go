package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/dianhadi/golib/log"
	"github.com/dianhadi/search/internal/config"
	"github.com/dianhadi/search/internal/handler/helper"
	handlerPost "github.com/dianhadi/search/internal/handler/http/post"
	repoPost "github.com/dianhadi/search/internal/repo/post"
	usecasePost "github.com/dianhadi/search/internal/usecase/post"
	"github.com/dianhadi/search/pkg/elastic"
	"github.com/go-chi/chi"
	"go.elastic.co/apm/module/apmchi"
)

const (
	serviceName = "search-http-service"
)

func main() {
	log.New(serviceName)

	log.Info("Get Configuration")
	appConfig, err := config.GetConfig("config/main.yaml")
	if err != nil {
		panic(err)
	}

	log.Info("Connect to Elastic")
	elasticModule, err := elastic.New(appConfig.Elastic.Host, appConfig.Elastic.Port)
	if err != nil {
		panic(err)
	}

	log.Info("Init Repo")
	repoPost, err := repoPost.New(elasticModule)
	if err != nil {
		panic(err)
	}

	log.Info("Init Usecase")
	usecasePost, err := usecasePost.New(repoPost)
	if err != nil {
		panic(err)
	}

	log.Info("Init Handler")
	handlerPost, err := handlerPost.New(usecasePost)
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()
	r.Use(apmchi.Middleware())
	r.Use(helper.Common)
	r.Use(helper.Recover)

	log.Info("Register Route")
	r.Get("/v1/search", handlerPost.Search)

	log.Infof("Starting server on port %s...", appConfig.Server.Port)
	startServer(":"+appConfig.Server.Port, r)
}

func startServer(port string, r http.Handler) {
	srv := http.Server{
		Addr:    port,
		Handler: r,
	}

	// Create a channel that listens on incomming interrupt signals
	signalChan := make(chan os.Signal, 1)
	signal.Notify(
		signalChan,
		syscall.SIGHUP,  // kill -SIGHUP XXXX
		syscall.SIGINT,  // kill -SIGINT XXXX or Ctrl+c
		syscall.SIGQUIT, // kill -SIGQUIT XXXX
	)

	// Graceful shutdown
	go func() {
		// Wait for a new signal on channel
		<-signalChan
		// Signal received, shutdown the server
		log.Info("shutting down..")

		// Create context with timeout
		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()
		srv.Shutdown(ctx)

		// Check if context timeouts, in worst case call cancel via defer
		select {
		case <-time.After(21 * time.Second):
			log.Info("not all connections done")
		case <-ctx.Done():
		}
	}()

	err := srv.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}
}

package main

import (
	"github.com/dianhadi/golib/log"
	"github.com/dianhadi/search/internal/config"
	handlerPost "github.com/dianhadi/search/internal/handler/consumer/post"
	repoPost "github.com/dianhadi/search/internal/repo/post"
	usecasePost "github.com/dianhadi/search/internal/usecase/post"
	"github.com/dianhadi/search/pkg/elastic"
	"github.com/dianhadi/search/pkg/mq"
	"go.elastic.co/apm"
)

const (
	serviceName    = "search-consumer-service"
	serviceVersion = "0.0.1"
)

func main() {
	log.New(serviceName)

	log.Info("Get Configuration")
	appConfig, err := config.GetConfig("config/main.yaml")
	if err != nil {
		panic(err)
	}

	log.Info("Init Tracer")
	apm.DefaultTracer, _ = apm.NewTracer(serviceName, serviceVersion)

	log.Info("Connect to Consumer")
	consumer, err := mq.New(appConfig.Consumer.Host, appConfig.Consumer.Port, appConfig.Consumer.Username, appConfig.Consumer.Password)
	if err != nil {
		panic(err)
	}
	defer consumer.Close()

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

	consumer.Consume("post-created", handlerPost.PostCreatedConsumer())
}

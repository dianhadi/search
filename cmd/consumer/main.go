package main

import (
	"github.com/dianhadi/golib/log"
	"github.com/dianhadi/search/internal/config"
	handlerPost "github.com/dianhadi/search/internal/handler/consumer/post"
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

	apm.DefaultTracer, _ = apm.NewTracer(serviceName, serviceVersion)

	consumer, err := mq.New(appConfig.Consumer.Host, appConfig.Consumer.Port, appConfig.Consumer.Username, appConfig.Consumer.Password)
	if err != nil {
		panic(err)
	}
	defer consumer.Close()

	consumer.Consume("post-created", handlerPost.PostCreatedConsumer())
}

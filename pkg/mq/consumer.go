package mq

import (
	"context"

	"go.elastic.co/apm"
)

func (m *MQ) Consume(name string, handlerConsumer HandlerConsumer) error {
	ch, err := m.channel.QueueDeclare(
		name,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	msgs, err := m.channel.Consume(
		ch.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			tx := apm.DefaultTracer.StartTransaction("/consumer/"+name, "consumer")
			ctx := apm.ContextWithTransaction(context.Background(), tx)
			handlerConsumer(ctx, d.Body)
			tx.End()
		}
	}()

	<-forever
	return nil
}

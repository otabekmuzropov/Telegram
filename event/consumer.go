package main

import "github.com/streadway/amqp"

type Consumer struct {
	conn *amqp.Connection
	channel *amqp.Channel
	exchangeName  string
	queueName   string
	routingKey  string
}

func NewConsumer(conn *amqp.Connection, exchangeName, queueName, routingKey string) (*Consumer, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	err = declareExchange(ch, exchangeName)
	if err != nil {
		return nil, err
	}

	q, err := declareQueue(ch, queueName)
	if err != nil {
		return nil, err
	}

	err = ch.QueueBind(
		q.Name,
		routingKey,
		exchangeName,
		false,
		nil,
		)

	if err != nil {
		return nil, err
	}

	return &Consumer{
		conn:         conn,
		channel:      ch,
		exchangeName: exchangeName,
		queueName:    queueName,
		routingKey:   routingKey,
	}, nil
}


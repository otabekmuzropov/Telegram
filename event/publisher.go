package main

import "github.com/streadway/amqp"

type Publisher struct {
	conn *amqp.Connection
	channel *amqp.Channel
	exchangeName string
}

func NewPublisher(conn *amqp.Connection, exchangeName string) (*Publisher, error) {
	ch, err := conn.Channel()

	if err != nil {
		return nil, err
	}

	err = declareExchange(ch, exchangeName)

	if err != nil {
		return nil, err
	}

	return &Publisher{
		conn:         conn,
		channel:      ch,
		exchangeName: exchangeName,
	}, nil
}

func (p *Publisher) Push(routingKey string, msg amqp.Publishing) error {
	err := p.channel.Publish(
		p.exchangeName,
		routingKey,
		false,
		false,
		msg,
		)

	return err
}





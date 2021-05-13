package main

import "github.com/streadway/amqp"

func declareQueue(ch *amqp.Channel, name string) (amqp.Queue, error){
	return ch.QueueDeclare(
		name,
		true,
		false,
		false,
		false,
		nil,
		)
}

func declareExchange(ch *amqp.Channel, name string) error {
	return ch.ExchangeDeclare(
		name,
		amqp.ExchangeTopic,
		true,
		false,
		false,
		false,
		nil,
		)
}




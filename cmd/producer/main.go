package main

import (
	"fmt"

	"github.com/henrique77/rabbitmq_producer_consumer/pkg/rabbitmq"
)

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	for i := 0; i < 10; i++ {
		if (i % 2) == 0 {
			rabbitmq.Publish(ch, fmt.Sprint(i), "amq.direct", "even")
		} else {
			rabbitmq.Publish(ch, fmt.Sprint(i), "amq.direct", "odd")
		}
	}
}

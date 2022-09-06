package workers

import (
	"fmt"
	"test_tech/common/parsers/order"
	dbOrder "test_tech/common/persistors/order"
)

type consumer struct {
	id           int
	consumerPool chan *consumer
	messages     chan []byte
	Parser       order.Parser
	Persistor    dbOrder.Persistor
}

func (c *consumer) start() {
	fmt.Printf("Starting worker %d\n", c.id)
	for {
		select {
		case message := <-c.messages:
			c.process(message)
			consumerPool <- c
		}
	}
}

func (c *consumer) process(message []byte) {
	fmt.Printf("Consumer %d, processing message: %s\n", c.id, string(message))

	order, err := c.Parser.Parse(message)
	if err != nil {
		c.fallback(message, err)
		return
	}

	if err = c.Persistor.Create(order); err != nil {
		c.fallback(message, err)
		return
	}

	fmt.Printf("Consumer %d, end processing message\n", c.id)
}

func (c *consumer) fallback(message []byte, err error) {
	fmt.Printf("fallback for message %s\n", string(message))
	fmt.Printf("reasons: %s", err.Error())

	//TODO implement fallback for invalid messages
	return
}

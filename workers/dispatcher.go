package workers

import (
	"test_tech/common/parsers/order"
	dbOrder "test_tech/common/persistors/order"
)

var consumerPool chan *consumer
var InputChannel chan []byte

func Init() {
	consumerPool = make(chan *consumer, 10)
	InputChannel = make(chan []byte)

	for i := 0; i < 10; i++ {
		c := &consumer{
			id:           i,
			consumerPool: consumerPool,
			messages:     make(chan []byte),
			Parser:       order.GetParser(),
			Persistor:    dbOrder.GetPersistor(),
		}

		go c.start()
		consumerPool <- c
	}
}

func Dispatch() {
	for {
		consumer := <-consumerPool
		select {
		case message := <-InputChannel:
			consumer.messages <- message
		}
	}
}
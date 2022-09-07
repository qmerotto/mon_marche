package workers

import (
	"test_tech/common/parsers/order"
	dbOrder "test_tech/common/persistors/order"
)

var consumerPool chan *consumer
var InputChannel chan []byte

func Init() {
	consumerPool = make(chan *consumer, 10)
	InputChannel = make(chan []byte, 10000)

	for i := 0; i < 10; i++ {
		c := &consumer{
			id:           i,
			consumerPool: consumerPool,
			messages:     make(chan []byte, 1),
			parser:       order.GetParser(),
			persistor:    dbOrder.GetPersistor(),
		}

		go c.start()
		consumerPool <- c
	}
}

func Dispatch() {
	defer func() {
		close(InputChannel)
		close(consumerPool)
	}()

	for {
		consumer := <-consumerPool
		select {
		case message := <-InputChannel:
			consumer.messages <- message
		}
	}
}

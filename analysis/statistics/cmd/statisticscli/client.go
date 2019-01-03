package statisticscli

import (
	"fmt"
	"github.com/Shopify/sarama"
)

type Statistics interface {
	Close()

	Event(who, where, function string, time int64)
}

type istatistics struct {
	producer sarama.AsyncProducer
}

func (s *istatistics) Close() {
	s.producer.Close()
}

func (s *istatistics) Event(who, where, function string, time int64) {
	s.Kafka(&sarama.ProducerMessage{
		Topic: "foundation_statistics_streaming",
		Key:   sarama.StringEncoder("statistics"),
		Value: sarama.ByteEncoder(fmt.Sprintf("%s;%s;%s;%d", function, who, where, time)),
	})
}

func (s *istatistics) Kafka(msg *sarama.ProducerMessage) {
	s.producer.Input() <- msg
	select {
	case success := <-s.producer.Successes():
		fmt.Printf("offset: %d,  timestamp: %s", success.Offset, success.Timestamp.String())
	case err := <-s.producer.Errors():
		fmt.Printf("err: %s\n", err.Err.Error())
	}
}

func NewConnect(addr string) (Statistics, error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	config.Version = sarama.V0_10_0_1

	producer, err := sarama.NewAsyncProducer([]string{addr}, config)
	if err != nil {
		fmt.Printf("kafka create producer error :%s\n", err.Error())
		return nil, err
	}
	return &istatistics{producer: producer}, nil
}

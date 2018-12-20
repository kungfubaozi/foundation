package streaming

type Streaming interface {
	//生产消息
	Producer(function, action string, kvs ...string)

	Close()
}

type steamingcli struct {
}

//1：发送kafka消息
//2：发送rabbitmq消息
func (cli *steamingcli) Producer(method, action string, kvs ...string) {

}

func (cli *steamingcli) Close() {

}

func NewClient(svc, sparkKafkaAddress, loggerMQAddress string) (Streaming, error) {
	return &steamingcli{}, nil
}

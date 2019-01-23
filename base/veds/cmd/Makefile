build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo
	docker build -t fs-base-veds .

run:
	docker run -p 58099:58099 -e CONSUL_ADDR=http://192.168.80.67:8500 \
    -e MONGODB_ADDR=192.168.2.60:27017 \
    -e MICRO_PORT=58099 \
    -e SECRET_KEY=29f15ad625b75cbed9b464e7ad63f644\
    -e ZIPKIN_ADDR=http://192.168.2.60:9411/api/v2/spans \
    -e STATISTICS_KAFKA_ADDR=192.168.2.60:9092 \
    -e REPORTER_AMQP_ADDR=amqp://root:123456@192.168.2.60:5672/ \
    -e MESSAGE_AMQP_ADDR=amqp://root:123456@192.168.2.60:5672/ \
    -e HOST_ADDR=192.168.80.67 fs-base-veds


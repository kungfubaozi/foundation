build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo
	docker build -t fs-entry-register .

run:
	docker run -p 58097:58097 -e CONSUL_ADDR=http://192.168.80.67:8500 \
    -e MONGODB_ADDR=192.168.2.60:27017 \
    -e MICRO_PORT=58097 \
    -e REDIS_ADDR=192.168.2.60:6379 \
    -e ZIPKIN_ADDR=http://192.168.2.60:9411/api/v2/spans \
    -e STATISTICS_KAFKA_ADDR=192.168.2.60:9092 \
    -e REPORTER_AMQP_ADDR=amqp://root:123456@192.168.2.60:5672/ \
    -e HOST_ADDR=192.168.80.67 fs-entry-register


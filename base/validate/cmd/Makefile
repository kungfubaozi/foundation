build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo
	docker build -t fs-base-validate .

run:
	docker run -p 58095:58095 -e CONSUL_ADDR=http://192.168.80.67:8500 \
    -e MONGODB_ADDR=192.168.2.60:27017 \
    -e MICRO_PORT=58095 \
    -e ZIPKIN_ADDR=http://192.168.2.60:9411/api/v2/spans \
    -e VALIDATE_TEMPLATE=【嘉实集团】你的验证码为%s。请在%d分钟内输入 \
    -e STATISTICS_KAFKA_ADDR=192.168.2.60:9092 \
    -e REPORTER_AMQP_ADDR=amqp://root:123456@192.168.2.60:5672/ \
    -e MESSAGE_AMQP_ADDR=amqp://root:123456@192.168.2.60:5672/ \
    -e HOST_ADDR=192.168.80.67 fs-base-validate


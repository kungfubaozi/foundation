build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo
	docker build -t fs-base-invite .

run:
	docker run -p 58119:58119 -e CONSUL_ADDR=http://192.168.80.67:8500 \
    -e MONGODB_ADDR=192.168.2.60:27017 \
    -e MICRO_PORT=58119 \
    -e ZIPKIN_ADDR=http://192.168.2.60:9411/api/v2/spans \
    -e INVITE_MESSAGE=【嘉实集团】邀请您加入我们！请使用邀请码%s在产品官网或手机APP使用邀请码进行登录后设置密码。感谢您的支持 \
    -e STATISTICS_KAFKA_ADDR=192.168.2.60:9092 \
    -e REPORTER_AMQP_ADDR=amqp://root:123456@192.168.2.60:5672/ \
    -e MESSAGE_AMQP_ADDR=amqp://root:123456@192.168.2.60:5672/ \
    -e HOST_ADDR=192.168.80.67 fs-base-invite


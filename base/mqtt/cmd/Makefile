build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo
	docker build -t fs-base-mqtt .

run:
	docker run -p 58088:58081 -e CONSUL_ADDR=http://192.168.80.67:8500 \
    -e MONGODB_ADDR=192.168.2.60:27017 \
    -e MICRO_PORT=58088 \
    -e MQTT_BROKER=tcp://foundation.mqtt.iot.gz.baidubce.com:1883 \
    -e MQTT_USERNAME=foundation/user \
    -e MQTT_PASSWORD=3zyh6KBHtD/1TojGv0tKOVzUPegPlvQkKfQuowHCuUQ= \
    -e STATISTICS_KAFKA_ADDR=192.168.2.60:9092 \
    -e REPORTER_AMQP_ADDR=amqp://root:123456@192.168.2.60:5672/ \
    -e HOST_ADDR=192.168.80.67 fs-base-mqtt


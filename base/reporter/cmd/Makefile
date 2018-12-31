build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo
	docker build -t fs-base-reporter .

run:
	docker run -p 58087:58081 -e CONSUL_ADDR=http://192.168.80.67:8500 \
    -e MONGODB_ADDR=192.168.2.60:27017 \
    -e MICRO_PORT=58087 \
    -e STATISTICS_KAFKA_ADDR=192.168.2.60:9092 \
    -e HOST_ADDR=192.168.80.67 fs-base-reporter


package osenv

import (
	"os"
	"strconv"
)

func GetConsulAddr() string {
	s := os.Getenv("CONSUL_ADDR")
	if len(s) == 0 {
		return "http://192.168.80.67:8500"
	}
	return s
}

func GetAMQPAddr() string {
	return os.Getenv("AMQP_ADDR")
}

func GetReporterAMQPAddr() string {
	return os.Getenv("REPORTER_AMQP_ADDR")
}

func GetReporterKafkaAddr() string {
	return os.Getenv("STATISTICS_KAFKA_ADDR")
}

func GetFaceCompareScore() float64 {
	f, e := strconv.ParseFloat(os.Getenv("FACE_COMPARE_SCORE"), 64)
	if e != nil {
		return 80.0
	}
	return f
}

func GetMongoDBAddr() string {
	return os.Getenv("MONGODB_ADDR")
}

func GetHostIp() string {
	return os.Getenv("HOST_ADDR")
}

func GetMicroPortString() string {
	return ":" + strconv.FormatInt(GetMicroPort(), 10)
}

func GetMicroPort() int64 {
	i, e := strconv.ParseInt(os.Getenv("MICRO_PORT"), 10, 64)
	if e != nil {
		return 0
	}
	return i
}

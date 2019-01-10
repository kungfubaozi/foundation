package osenv

import (
	"os"
	"strconv"
)

func GetTokenKey() string {
	token := os.Getenv("TOKEN_KEY")
	if len(token) <= 16 {
		return "e48df34a-0f32-11e9-ab14-d663bd873d93"
	}
	return token
}

func GetValidateTemplate() string {
	return os.Getenv("VALIDATE_TEMPLATE")
}

func GetNodeNumber() int64 {
	i, e := strconv.ParseInt(os.Getenv("NODE"), 10, 64)
	if e != nil {
		return 1
	}
	return i
}

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

func GetMessageAMQPAddr() string {
	return os.Getenv("MESSAGE_AMQP_ADDR")
}

func GetRedisAddr() string {
	return os.Getenv("REDIS_ADDR")
}

func GetReporterAMQPAddr() string {
	return os.Getenv("REPORTER_AMQP_ADDR")
}

func GetReporterKafkaAddr() string {
	return os.Getenv("REPORTER_KAFKA_ADDR")
}

func GetStatisticsKafkaAddr() string {
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

func GetZipkinAddr() string {
	return os.Getenv("ZIPKIN_ADDR")
}

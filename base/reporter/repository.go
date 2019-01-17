package reporter

import (
	"gopkg.in/mgo.v2"
)

type repository interface {
	Close()

	Write(logger *logger) error
}

type logger struct {
	Func      string `bson:"func"`
	Who       string `bson:"who"`
	Where     string `bson:"where"`
	Timestamp int64  `bson:"timestamp"`
	Date      string `bson:"date"`
	Status    int64  `bson:"Status"`
}

type reporterRepository struct {
	session *mgo.Session
}

func (repo *reporterRepository) Close() {
	repo.session.Close()
}

func (repo *reporterRepository) Write(logger *logger) error {
	return repo.session.DB("fds_logger").C(logger.Func + "_" + logger.Date).Insert(logger)
}

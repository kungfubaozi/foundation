package reporter

import (
	"gopkg.in/mgo.v2"
	"zskparker.com/foundation/pkg/constants"
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
	return repo.session.DB(fs_constants.DB_LOGGER).C(logger.Func + "_" + logger.Date).Insert(logger)
}

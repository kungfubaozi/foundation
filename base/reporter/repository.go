package reporter

import (
	"gopkg.in/mgo.v2"
	"strconv"
)

type repository interface {
	Close()

	Write(logger *logger) error
}

type logger struct {
	do        int64
	who       string
	where     string
	timestamp string
	date      string
}

type reporterRepository struct {
	session *mgo.Session
}

func (repo *reporterRepository) Close() {
	repo.session.Close()
}

func (repo *reporterRepository) Write(logger *logger) error {
	return repo.session.DB("foundation_logger").C(strconv.FormatInt(logger.do, 10) + "_" + logger.date).Insert(logger)
}

package user

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type repository interface {
	Add()

	Get()

	Close()

	UpdatePassword(userId, password string) error

	UpdatePhone(userId, phone string) error

	UpdateEnterprise(userId, enterprise string) error

	UpdateEmail(userId, email string) error
}

type userRepository struct {
	session *mgo.Session
}

func (repo *userRepository) Close() {
	repo.session.Close()
}

func (repo *userRepository) Add() {
	panic("implement me")
}

func (repo *userRepository) Get() {
	panic("implement me")
}

func (repo *userRepository) UpdatePassword(userId, password string) error {
	_, e := repo.collection().Upsert(bson.M{"user_id": userId}, bson.M{"$set": bson.M{"password": password}})
	return e
}

func (repo *userRepository) UpdatePhone(userId, phone string) error {
	_, e := repo.collection().Upsert(bson.M{"user_id": userId}, bson.M{"$set": bson.M{"phone": phone}})
	return e
}

func (repo *userRepository) UpdateEnterprise(userId, enterprise string) error {
	_, e := repo.collection().Upsert(bson.M{"user_id": userId}, bson.M{"$set": bson.M{"enterprise": enterprise}})
	return e
}

func (repo *userRepository) UpdateEmail(userId, email string) error {
	_, e := repo.collection().Upsert(bson.M{"user_id": userId}, bson.M{"$set": bson.M{"email": email}})
	return e
}

func (repo *userRepository) collection() *mgo.Collection {
	return repo.session.DB("foundation").C("user")
}

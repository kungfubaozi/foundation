package validate

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type repository interface {
	Create(v *verification) error

	Get(id string) (*verification, error)

	Update(id string, state int64) error

	Close()
}

type verification struct {
	VerId    bson.ObjectId `bson:"_id"`
	Code     string        `bson:"code"`
	CreateAt int64         `bson:"create_at"`
	To       string        `bson:"to"`
	Func     string        `bson:"func"`
	Voucher  string        `bson:"voucher"`
	State    int64         `bson:"state"`
}

type validateRepository struct {
	session *mgo.Session
}

func (svc *validateRepository) Close() {
	svc.session.Close()
}

func (svc *validateRepository) Create(v *verification) error {
	return svc.collection(v.VerId.Hex()).Insert(v)
}

func (svc *validateRepository) Update(id string, state int64) error {
	err := svc.collection(id).Update(bson.M{"_id": bson.ObjectIdHex(id)}, bson.M{"$set": bson.M{"state": state}})
	return err
}

//func (svc *validateRepository) FindLast(voucher string) (*verification, error) {
//	ver := &verification{}
//	err := svc.collection().Find(bson.M{"voucher": voucher}).Sort("-create_at").One(ver)
//	return ver, err
//}

func (svc *validateRepository) Get(id string) (*verification, error) {
	ver := &verification{}
	err := svc.collection(id).Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(ver)
	return ver, err
}

func (svc *validateRepository) collection(id string) *mgo.Collection {
	return svc.session.DB("fds_validate").C("v_" + id[len(id)-1:])
}

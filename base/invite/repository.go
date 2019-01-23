package invite

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
	"zskparker.com/foundation/pkg/constants"
)

type repository interface {
	Add(m *model) error

	Close()

	FindInvite(code string) (*model, error)

	FindInviteByAccount(account string, phone bool) (*model, error)

	//GetInviteByUserId(userId string, pageIndex, pageSize int64) ([]*model, error)

	Update(account, id string) error
}

type inviteRepository struct {
	session *mgo.Session
}

func (repo *inviteRepository) FindInvite(code string) (*model, error) {
	m := &model{}
	err := repo.collection().Find(bson.M{"code": code}).One(m)
	return m, err
}

//func (repo *inviteRepository) GetInviteByUserId(userId string, pageIndex, pageSize int64) ([]*model, error) {
//	m := &model{}
//	err := repo.collection().Find(bson.M{"operateUserId": userId}).Sort("-create_at").
//	return m, err
//}

func (repo *inviteRepository) FindInviteByAccount(account string, phone bool) (*model, error) {
	m := &model{}
	var b bson.M
	if phone {
		b = bson.M{"phone": account}
	} else {
		b = bson.M{"email": account}
	}
	err := repo.collection().Find(b).One(m)
	return m, err
}

func (repo *inviteRepository) Add(m *model) error {
	_, e := repo.collection().Upsert(bson.M{"_id": m.InviteId}, m)
	if e != nil {
		return e
	}
	return e
}

func (repo *inviteRepository) Update(account, id string) error {
	e := repo.collection().Update(bson.M{"_id": bson.ObjectIdHex(id)}, bson.M{"$set": bson.M{"ok": true, "ok_time": time.Now().UnixNano()}})
	return e
}

func (repo *inviteRepository) collection() *mgo.Collection {
	return repo.session.DB(fs_constants.DB_USER).C("invite")
}

func (repo *inviteRepository) Close() {
	repo.session.Close()
}

type model struct {
	Phone         string        `bson:"phone"`
	Email         string        `bson:"email"`
	Enterprise    string        `bson:"enterprise"`
	Username      string        `bson:"username"`
	RealName      string        `bson:"real_name"`
	Level         int64         `bson:"level"`
	ExpireAt      int64         `bson:"expire_at"` //过期时间
	CreateAt      int64         `bson:"create_at"`
	Code          string        `bson:"code"`
	InviteId      bson.ObjectId `bson:"_id"`
	FromProject   string        `bson:"from_project"`
	FromClientId  string        `bson:"from_client_id"`
	Ok            bool          `bson:"ok"`
	OkTime        int64         `bson:"ok_time"`
	OperateUserId string        `bson:"operate_user_id"`
	Scope         int64         `bson:"scope"`
}

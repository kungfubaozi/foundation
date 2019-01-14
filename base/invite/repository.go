package invite

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type repository interface {
	Add(inviteCode, fix string, m *model) error

	Close()

	FindInvite(inviteCode string, code string) (*model, error)

	Update(inviteCode, fix, account, id string) error

	FindHistory(fix, account string) (*history, error)
}

type inviteRepository struct {
	session *mgo.Session
}

func (repo *inviteRepository) FindInvite(inviteCode string, code string) (*model, error) {
	m := &model{}
	err := repo.collection(inviteCode).Find(bson.M{"code": code}).One(m)
	return m, err
}

func (repo *inviteRepository) Add(inviteCode, fix string, m *model) error {
	_, e := repo.collection(inviteCode).Upsert(bson.M{"_id": m.InviteId}, m)
	if e != nil {
		return e
	}
	tag := fmt.Sprintf("%s%s", m.Phone, m.Email)
	_, e = repo.inviteHistoryCollection(fix).Upsert(bson.M{"tag": tag}, &history{
		Tag:      tag,
		At:       inviteCode[:1],
		ExpireAt: m.ExpireAt,
	})
	return e
}

func (repo *inviteRepository) Update(inviteCode, fix, account, id string) error {
	e := repo.collection(inviteCode).Update(bson.M{"_id": bson.ObjectIdHex(id)}, bson.M{"$set": bson.M{"ok": true, "ok_time": time.Now().UnixNano()}})
	e = repo.inviteHistoryCollection(fix).Update(bson.M{"tag": account}, bson.M{"$set": bson.M{"ok": true}})
	return e
}

func (repo *inviteRepository) FindHistory(fix, account string) (*history, error) {
	h := &history{}
	e := repo.inviteHistoryCollection(fix).Find(bson.M{"tag": account}).One(h)
	return h, e
}

func (repo *inviteRepository) collection(code string) *mgo.Collection {
	return repo.session.DB("foundation").C("invite_" + code[0:1])
}

func (repo *inviteRepository) inviteHistoryCollection(fix string) *mgo.Collection {
	return repo.session.DB("foundation").C("invite_history_" + fix)
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
}

type history struct {
	Tag      string `bson:"tag"`
	At       string `bson:"at"`
	ExpireAt int64  `bson:"expire_at"` //过期时间
	Ok       bool   `bson:"ok"`
}

package face

import (
	"crypto/tls"
	"encoding/json"
	"github.com/garyburd/redigo/redis"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"net/http"
)

type repository interface {
	Get(userId string) (*faceset, error)

	Upsert(faceset *faceset) error

	Delete(userId string) error

	GetToken() (string, error)

	SetToken(token string) error

	Close()
}

type faceRepository struct {
	session *mgo.Session
	conn    redis.Conn
}

func (repo *faceRepository) GetToken() (string, error) {
	t, err := redis.String(repo.conn.Do("get", "baidu_access_token"))
	if err != nil && err == redis.ErrNil {
		token, err := resetAccessToken()
		if err != nil {
			return "", err
		}
		err = repo.SetToken(token)
		if err == nil {
			return token, nil
		}
	}
	if err != nil {
		return "", err
	}
	return t, nil
}

func (repo *faceRepository) SetToken(token string) error {
	_, err := repo.conn.Do("set", "baidu_access_token", token)
	return err
}

//获取accessToken
func resetAccessToken() (string, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get("https://aip.baidubce.com/oauth/2.0/token?grant_type=client_credentials" +
		"&client_id=5QQkC038Xnrp3m3A5rPg0TjV&client_secret=34IgrK7WUIMpZgXRCLCusC3xQ5Nt2gBb")
	if err != nil {
		return "", err
	}
	if resp.StatusCode == 200 {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}
		at := &accessToken{}
		err = json.Unmarshal(body, at)
		if err != nil {
			return "", err
		}
		if len(at.Error) > 0 {
			return "", RequestErr
		}
		return at.AccessToken, nil
	}
	return "", RequestErr
}

func (repo *faceRepository) Get(userId string) (*faceset, error) {
	coll := repo.collection(userId)
	fs := &faceset{}
	err := coll.Find(bson.M{"user_id": userId}).One(fs)
	return fs, err
}

func (repo *faceRepository) Upsert(faceset *faceset) error {
	coll := repo.collection(faceset.UserId)
	_, err := coll.Upsert(bson.M{"user_id": faceset.UserId}, faceset)
	return err
}

func (repo *faceRepository) Delete(userId string) error {
	coll := repo.collection(userId)

	return coll.Remove(bson.M{"user_id": userId})
}

func (repo *faceRepository) Close() {
	repo.session.Close()
}

func (repo *faceRepository) collection(userId string) *mgo.Collection {
	return repo.session.DB("foundation").C("faceset_" + userId[0:5])
}

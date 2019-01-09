package facecli

import (
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"zskparker.com/foundation/base/face"
)

func TestNewClient(t *testing.T) {

	//zipkinTracer, reporter := serv.NewZipkin("http://192.168.2.60:9411/api/v2/spans", names.F_SVC_FACE, "58085")
	//defer reporter.Close()
	//
	//c := NewClient(zipkinTracer)

	b, err := ioutil.ReadFile("system_admin.jpg")
	if err != nil {
		panic(err)
	}

	str := base64.StdEncoding.EncodeToString(b)
	fmt.Println(str)

	//resp, err := c.Upsert(context.Background(), &fs_base_face.UpsertRequest{
	//	UserId:     "5c31d107133cf440c4a55209",
	//	Base64Face: str,
	//})

	//resp, err := c.Compare(context.Background(), &fs_base_face.CompareRequest{
	//	UserId:     "5c31d107133cf440c4a55209",
	//	Base64Face: str,
	//})

	//resp, err := c.Search(context.Background(), &fs_base_face.SearchRequest{
	//	Base64Face: str,
	//})
	//
	//if err != nil {
	//	panic(err)
	//}
	//
	//fmt.Println(resp.State)
	//fmt.Println(resp.UserId)

	//str, err := resetAccessToken()
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(str)

}

//获取accessToken
func resetAccessToken() (string, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get("https://aip.baidubce.com/oauth/2.0/token?grant_type=client_credentials" +
		"&client_id=dsU7P5P3lRT9wR8pQDLlOyBX&client_secret=wXOnBUX87GDh14rPMdWZe31WrhxhX1ZM")
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
			return "", face.RequestErr
		}
		return at.AccessToken, nil
	}
	return "", face.RequestErr
}

//errCode:100/110/111
type accessToken struct {
	AccessToken   string `json:"access_token"`
	SessionKey    string `json:"session_key"`
	Scope         string `json:"scope"`
	RefreshToken  string `json:"refresh_token"`
	SessionSecret string `json:"session_secret"`
	ExpiresIn     int64  `json:"expires_in"`
	Error         string `json:"error"`
}

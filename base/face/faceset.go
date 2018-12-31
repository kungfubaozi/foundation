package face

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

var RequestErr = errors.New("faceset request error")

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

type UserFaceToken string

var currentToken *accessToken

//注册人脸(image,image_type,group_id,user_id)
func RegisterFace(base64Face, userId, groupId string) (map[string]interface{}, error) {
	return faceAI("https://aip.baidubce.com/rest/2.0/face/v3/faceset/user/add", map[string]interface{}{
		"group_id":   groupId,
		"user_id":    userId,
		"image":      base64Face,
		"image_type": "BASE64",
	})
}

//更新人脸
func UpdateFace(base64Face, userId, groupId string) (map[string]interface{}, error) {
	return faceAI("https://aip.baidubce.com/rest/2.0/face/v3/faceset/user/update", map[string]interface{}{
		"group_id":   groupId,
		"user_id":    userId,
		"image":      base64Face,
		"image_type": "BASE64",
	})
}

//删除人脸
func DeleteFace(userId, groupId string) (map[string]interface{}, error) {
	return faceAI("https://aip.baidubce.com/rest/2.0/face/v3/faceset/user/delete", map[string]interface{}{
		"group_id": groupId,
		"user_id":  userId,
	})
}

//创建用户组
func createFaceset(groupId string) (map[string]interface{}, error) {
	return faceAI("https://aip.baidubce.com/rest/2.0/face/v3/faceset/group/add", map[string]interface{}{
		"group_id": groupId,
	})
}

//对比人脸
func faceCompare(needMatchBase64Face, targetFaceToken string) (map[string]interface{}, error) {
	return faceAI("https://aip.baidubce.com/rest/2.0/face/v3/match", []map[string]interface{}{
		{
			"image":      needMatchBase64Face,
			"image_type": "BASE64",
		},
		{
			"image":      targetFaceToken,
			"image_type": "FACE_TOKEN",
		},
	})
}

//人脸搜索
func faceSearch(base64Face, groupId string) (map[string]interface{}, error) {
	return faceAI("https://aip.baidubce.com/rest/2.0/face/v3/search", map[string]interface{}{
		"image":         base64Face,
		"image_type":    "BASE64",
		"group_id_list": groupId,
		"max_user_num":  1,
	})
}

//统一请求
func faceAI(url string, values interface{}) (map[string]interface{}, error) {
	if currentToken == nil {
		err := resetAccessToken()
		if err != nil {
			return nil, RequestErr
		}
	}
	b, err := json.Marshal(values)
	if err != nil {
		return nil, err
	}
	resp, err := http.Post(url+"?access_token="+currentToken.AccessToken, "application/json",
		bytes.NewReader(b))
	fmt.Println("url", resp.Request.URL)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == 200 {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		fmt.Println(string(body))
		v := make(map[string]interface{})
		err = json.Unmarshal(body, &v)
		if err != nil {
			fmt.Println(err)
		}
		code := v["error_code"].(float64)
		fmt.Println("error_code", code)
		if code == 100 || code == 110 || code == 111 {
			err = resetAccessToken()
			if err != nil {
				return nil, RequestErr
			}
			return faceAI(url, values)
		}
		return v, RequestErr
	}
	return nil, RequestErr
}

//获取accessToken
func resetAccessToken() error {
	resp, err := http.Get("https://aip.baidubce.com/oauth/2.0/token?grant_type=client_credentials&client_id=dsU7P5P3lRT9wR8pQDLlOyBX&client_secret=wXOnBUX87GDh14rPMdWZe31WrhxhX1ZM")
	if err != nil {
		return err
	}
	if resp.StatusCode == 200 {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		at := &accessToken{}
		err = json.Unmarshal(body, at)
		if err != nil {
			return err
		}
		if len(at.Error) > 0 {
			return RequestErr
		}
		currentToken = at
		return nil
	}
	return RequestErr
}

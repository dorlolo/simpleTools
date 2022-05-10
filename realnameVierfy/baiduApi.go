package realnameVierfy

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

//  百度云身份证、姓名比对
// api文档：https://ai.baidu.com/ai-doc/FACE/Tkqahnjtk

type BaiduIdcardVerify struct {
	apiKey       string
	secretKey    string
	accessToken  string
	expiressTime int64
}

func New_BaiduIdcardVerify(accessKey, secretKey string) IdcardVierfy {
	reqobj := BaiduIdcardVerify{
		apiKey:    accessKey,
		secretKey: secretKey,
	}
	var err error
	reqobj.accessToken, reqobj.expiressTime, err = getToken(reqobj.apiKey, reqobj.secretKey)
	if err != nil {
		panic("获取百度云token失败")
	}
	return &reqobj
}

func (s *BaiduIdcardVerify) Verify(realName, idcard string) (bool, error) {
	if s.accessToken == "" {
		var err error
		s.accessToken, s.expiressTime, err = getToken(s.apiKey, s.secretKey)
		if err != nil {
			//global.GVA_LOG.Error("获取百度云token失败")
			return false, errors.New("服务错误")
		}
	}
	var host = "https://aip.baidubce.com/rest/2.0/face/v3/person/idmatch"
	uri, err := url.Parse(host)
	if err != nil {
		fmt.Println(err)
	}
	query := uri.Query()
	query.Set("access_token", s.accessToken)
	uri.RawQuery = query.Encode()

	var params = map[string]string{}
	params["id_card_number"] = idcard
	params["name"] = realName
	sendBody, err := json.Marshal(params)
	if err != nil {
		return false, err
	}
	sendData := string(sendBody)
	client := &http.Client{}
	request, err := http.NewRequest("POST", uri.String(), strings.NewReader(sendData))
	if err != nil {
		return false, err
	}
	request.Header.Set("Content-Type", "application/json")
	response, err := client.Do(request)
	defer response.Body.Close()
	result, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return false, err
	}
	res := verifyResult{}
	if err = json.Unmarshal(result, &res); err != nil {
		return false, err
	}
	switch res.ErrorCode {
	case 0:
		return true, nil
	default:
		return false, errors.New(err.Error())
	}
}

type verifyResult struct {
	ErrorCode int    `json:"error_code"`
	ErrorMsg  string `json:"error_msg"`
}

func getToken(apiKey, secretKey string) (token string, expriessTime int64, err error) {
	var host = "https://aip.baidubce.com/oauth/2.0/token"
	var param = map[string]string{
		"grant_type":    "client_credentials",
		"client_id":     apiKey,
		"client_secret": secretKey,
	}

	uri, err := url.Parse(host)
	if err != nil {
		fmt.Println(err)
	}
	query := uri.Query()
	for k, v := range param {
		query.Set(k, v)
	}
	uri.RawQuery = query.Encode()

	response, err := http.Get(uri.String())
	if err != nil {
		fmt.Println(err)
	}
	result, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}
	res := tokenRes{}
	if err = json.Unmarshal(result, &res); err != nil {
		return
	}
	return res.AccessToken, time.Now().Unix() + res.ExpiresIn, nil

}

//
//type errResonse struct {
//	ErrorDescription string `json:"error_description"`
//	Error            string `json:"error"`
//}
//token返回结果
type tokenRes struct {
	RefreshToken  string `json:"refresh_token"`
	ExpiresIn     int64  `json:"expires_in"`
	Scope         string `json:"scope"`
	SessionKey    string `json:"session_key"`
	AccessToken   string `json:"access_token"`
	SessionSecret string `json:"session_secret"`
}

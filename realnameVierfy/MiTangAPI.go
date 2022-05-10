/*
 * @FileName:   idCardVerifyOnline.go
 * @Author:		JuneXu
 * @CreateTime:	2022/4/5 下午1:31
 * @Description:
 */

package realnameVierfy

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dorlolo/simpleRequest"
	uuid "github.com/satori/go.uuid"
	"io/ioutil"
	"net/http"
	"strings"
)

type resonse struct {
	Code    string `json:"code"`
	ReqNo   string `json:"reqNo"`
	Message string `json:"message"`
}

// 实名认证二要素验证接口
type IdCardOnlineVerifyer interface {
	Verify(idCard, name string) (bool, error)
}

//蜜堂有信实名认证接口（百度云第三方接口）
type MiTangIdCardVierfy struct {
	UserLicenseNo string
	AliAppcode    string
	BaiduAppcode  string
	ReqNo         string //请求编号,当前用不到
}

func New_MitangcardVierfy_UseBaiduApi(userLicenseNo, appcode string) *MiTangIdCardVierfy {
	return &MiTangIdCardVierfy{
		UserLicenseNo: userLicenseNo,
		BaiduAppcode:  appcode,
	}
}

func New_MitangcardVierfy_UseAliApi(userLicenseNo, appcode string) *MiTangIdCardVierfy {
	return &MiTangIdCardVierfy{
		UserLicenseNo: userLicenseNo,
		AliAppcode:    appcode,
	}
}
func (s *MiTangIdCardVierfy) Verify(idCard, name string) (bool, error) {
	if s.BaiduAppcode != "" {
		//获取身份证中的姓名
		url := "https://miitang.api.bdymkt.com/v1/tools/person/realnameVierfy"
		body := fmt.Sprintf("reqNo=1641803077694&name=%v&idCardNo=%v&userLicenseNo=%v", idCard, name, s.UserLicenseNo)
		req, err := http.NewRequest("POST", url, strings.NewReader(body))
		if err != nil {
			panic(err)
		}
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		if s.BaiduAppcode != "" {
			req.Header.Add("X-Bce-Signature", fmt.Sprintf("AppCode/%v", s.BaiduAppcode))
		} else if s.AliAppcode != "" {
			req.Header.Set("Authorization", "APPCODE "+s.AliAppcode)
		}
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		if resp.StatusCode == 200 {
			respBody, _ := ioutil.ReadAll(resp.Body)
			var res = resonse{}
			err = json.Unmarshal(respBody, &res)
			if err != nil {
				return false, err
			}
			switch res.Code {
			case "FP00000":
				return true, nil
			default:
				return false, errors.New(res.Message)
			}
		} else {
			return false, errors.New("实名认证请求失败")
		}
	} else if s.AliAppcode != "" {
		var req = simpleRequest.NewRequest()
		req.Headers().ConentType_formUrlencoded()
		req.Headers().SetAuthorization("APPCODE " + s.AliAppcode)
		req.Headers().Set("X-Ca-Nonce", uuid.NewV4().String())
		req.Body().Set("idCardNo", idCard).Set("name", name)
		req.Body().Set("userLicenseNo", s.UserLicenseNo)
		req.Body().Set("reqNo", uuid.NewV4().String())
		respBody, err := req.Post("https://miitangs01.market.alicloudapi.com/v1/tools/person/realnameVierfy")
		if err != nil {
			return false, errors.New("请求出错")
		}
		var res = resonse{}
		err = json.Unmarshal(respBody, &res)
		if err != nil {
			return false, errors.New("请求错误")
		}
		switch res.Code {
		case "FP00000":
			return true, nil
		default:
			return false, errors.New(res.Message)
		}
	}

	return true, nil
}

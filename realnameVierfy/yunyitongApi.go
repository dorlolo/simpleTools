package realnameVierfy

import (
	"encoding/json"
	"errors"
	"github.com/dorlolo/simpleRequest"
)

//=============================================================================
//
//	苏州云亿互通信息服务有限公司 二要素身份证识别api接口
//	IDCard SDK
//
// 官方api链接：https://market.aliyun.com/products/57126001/cmapi025518.html?#sku=yuncode1951800000
//=============================================================================

var codeToString = map[string]string{
	"0000": "身份证识别成功",
	"0001": "开户名不能为空",
	"0002": "开户名不能包含特殊字符",
	"0003": "身份证号不能为空",
	"0004": "身份证号格式错误",
	"0007": "无此身份证号码",
	"0008": "身份证信息不匹配",
	"0010": "系统维护，请稍后再试",
}

type response struct {
	Name        string `json:"name"`
	IdNo        string `json:"idNo"`
	RespMessage string `json:"respMessage"`
	RespCode    string `json:"respCode"`
	Province    string `json:"province"`
	City        string `json:"city"`
	County      string `json:"county"`
	Birthday    string `json:"birthday"`
	Sex         string `json:"sex"`
	Age         string `json:"age"`
}
type YunyitongcardVierfy struct {
	AppCode string
}

func New_YunyitongcardVierfyForTest(appcode string) IdcardVierfy {
	return &YunyitongcardVierfy{appcode}
}

func (s *YunyitongcardVierfy) Verify(realName, idcard string) (bool, error) {
	var req = simpleRequest.NewRequest()
	req.Headers().SetAuthorization("APPCODE " + s.AppCode)
	req.Headers().ConentType_formUrlencoded()
	req.Body().Set("idNo", idcard).Set("name", realName)
	jsonRes, err := req.Post("https://idenauthen.market.alicloudapi.com/idenAuthentication")
	if err != nil {
		return false, err
	}
	res := &response{}
	if err = json.Unmarshal(jsonRes, &res); err != nil {
		return false, errors.New("请求错误，请稍后再试")
	}
	if res.RespCode == "0000" {
		return true, nil
	} else {
		return false, errors.New(codeToString[res.RespCode])
	}
}

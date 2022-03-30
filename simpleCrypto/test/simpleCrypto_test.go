/*
 * @FileName:   simpleCrypto_test.go
 * @Author:		JuneXu
 * @CreateTime:	2022/3/10 上午12:37
 * @Description:
 */

package test

import (
	"fmt"
	"github.com/dorlolo/simpleTools/simpleCrypto"
	"testing"
	"time"
)

var (
	text = "你好，我是dorlolo"
)

//TestMd5 md5加密
func TestMd5(t *testing.T) {
	res1 := simpleCrypto.Md5Enscrypto(text)
	t.Log("md5加密-1:", res1)
}

var (
	appSecret = "41e2188c94ca4ee0"
	aesKey    = "KaYhCwtWQ4KEifkL"
)

//TestECB ECB加密
func TestECB(t *testing.T) {
	//1.1.2 组合明文
	comboStr := fmt.Sprintf("%v:%v:%v", appSecret, "1921681111", time.Now().Unix())
	//ecb加密dim
	aesStr := simpleCrypto.SimpleEncryptAesECB([]byte(comboStr), []byte(aesKey))
	t.Log("aesStr", aesStr)
}

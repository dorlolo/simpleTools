/*
 * @FileName:   idCardVerifyOnline_test.go
 * @Author:		JuneXu
 * @CreateTime:	2022/4/5 下午1:41
 * @Description:
 */

package realnameVierfy

import "testing"

func TestVierfy(t *testing.T) {
	var test = MiTangIdCardVierfy{
		UserLicenseNo: "1641803077694",
		AliAppcode:    "a38dcb97330f464d9f5f28913a2c643b",
	}
	t.Log(test.Verify("32058119910306341X", "徐俊杰"))
}

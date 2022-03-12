/*
 * @FileName:   addr_test.go
 * @Author:		JuneXu
 * @CreateTime:	2022/3/10 下午1:48
 * @Description:
 */

package utils

import "testing"

func TestGetAdr(t *testing.T) {
	port, err := GetFreePort()
	if err != nil {
		t.Log(err)
	}
	t.Log(port)

}

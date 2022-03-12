/*
 * @FileName:   random_test.go
 * @Author:		JuneXu
 * @CreateTime:	2022/2/28 下午1:49
 * @Description:
 */

package random

import (
	"testing"
)

func TestRandom(t *testing.T) {
	for i := 0; i < 3; i++ {
		a := RangeInt64(1000, 9999)
		t.Log(a)

	}

}

/*
 * @FileName:   time_test.go
 * @Author:		JuneXu
 * @CreateTime:	2022/2/25 下午2:37
 * @Description:
 */

package simpleTime

import (
	"fmt"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	result := NowTimeToTime(Format.Normal_YMDhms)
	fmt.Println(result)
}

func TestFormat(t *testing.T) {
	nowtime := time.Now()
	trans1 := nowtime.Format(Format.Normal_YMDhms)
	t.Log(trans1)
	trans2 := nowtime.Format(Format.CN_YMDhms)
	t.Log(trans2)
}

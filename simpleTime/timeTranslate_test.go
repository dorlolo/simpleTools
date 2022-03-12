/*
 * @FileName:   time_test.go
 * @Author:		JuneXu
 * @CreateTime:	2022/2/25 下午2:37
 * @Description:
 */

package timeUtil

import (
	"fmt"
	"testing"
)

func TestTime(t *testing.T) {
	result := NowTimeToTime(TimeFormat.Normal_YMDhms)
	fmt.Println(result)
}

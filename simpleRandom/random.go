/*
 * @FileName:   random.go
 * @Author:		JuneXu
 * @CreateTime:	2022/2/28 下午1:46
 * @Description:
 */

package random

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

//
//  RangeInt64
//  @Description: 随机范围
//  @param min 最小值  不能为0
//  @param max 最大值 不能为0
//  @return int64
//
func RangeInt64(min, max int64) int64 {
	if min >= max || min == 0 || max == 0 {
		return max
	}
	//time.Sleep(time.Millisecond * 850)
	return rand.Int63n(max - min)
}

//
//  RangeList
//  @Description: 从列表中随机取值
//  @param datalist
//  @return interface{}
//
func RangeList(datalist []interface{}) interface{} {
	index := rand.Intn(len(datalist))
	return datalist[index]
}

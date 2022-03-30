/*
 * @FileName:   arrary.go
 * @Author:		JuneXu
 * @CreateTime:	2022/3/30 下午10:05
 * @Description:
 */

package utils

import (
	"fmt"
	"strings"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: ArrayToString
//@description: 将数组格式化为字符串
//@param: array []interface{}
//@return: string

func ArrayToString(array []interface{}) string {
	return strings.Replace(strings.Trim(fmt.Sprint(array), "[]"), " ", ",", -1)
}

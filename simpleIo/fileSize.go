/*
 * @FileName:   fileSize.go
 * @Author:		JuneXu
 * @CreateTime:	2022/3/12 下午5:23
 * @Description:
 */

package utils

import "strconv"

//
//  StrDataTransToKB
//  @Description: 从人类可阅读的文件尺寸转化为int64类型的kb
//  @param originData
//  @return result
//  @return ok
//
func StrDataTransToKB(sizeStr string) (result int64, ok bool) {
	spaceUnit := sizeStr[len(sizeStr)-1:]
	spaceData := sizeStr[:len(sizeStr)-1]
	switch spaceUnit {
	case "T":
		data64, _ := strconv.ParseInt(spaceData, 10, 64)
		return data64 * 1024 * 1024 * 1024, true
	case "G":
		data64, _ := strconv.ParseInt(spaceData, 10, 64)
		return data64 * 1024 * 1024, true
	case "M":
		data64, _ := strconv.ParseInt(spaceData, 10, 64)
		return data64 * 1024, true
	default:
		return 0, false
	}
}

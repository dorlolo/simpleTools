/*
 * @FileName:   fileSize.go
 * @Author:		JuneXu
 * @CreateTime:	2022/3/12 下午5:23
 * @Description:
 */

package simpleIo

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

//
//  ReadFileSizeReadable
//  @Description:  获取文件大小，转化为利于阅读的格式
//  @param path
//  @return string
//  @return error
//
func ReadFileSizeReadable(path string) (string, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return "0KB", err
	}
	intSize := float64(len(content))
	kbSize := fmt.Sprintf("%.2f", intSize/KB)
	tmp := strings.Split(kbSize, ".")[0]
	if len(tmp) > 4 {
		mbSize := fmt.Sprintf("%.2f", intSize/MB)
		tmp = strings.Split(mbSize, ".")[0]
		if len(tmp) > 4 {
			tbSize := fmt.Sprintf("%.2f", intSize/GB)
			return fmt.Sprintf("%vGB", tbSize), nil
		}
		return fmt.Sprintf("%vMB", mbSize), nil
	}
	return fmt.Sprintf("%vKB", kbSize), nil
}

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

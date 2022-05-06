/*
 * @FileName:
 * @Author:		xjj
 * @CreateTime:	下午5:34
 * @Description:
 */
package simpleIo

import (
	"encoding/base64"
	"io/ioutil"
	"os"
	"strings"
)

//检查文件是否存在
func FileExist(filePath string) bool {
	var exist = true
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

//
//  Base64DataSaveToFIle
//  @Description: base64数据存储到文件
//  @param bs64Data string base64格式字符串
//  @param filePath string 文件路径
//  @param isCover bool 文件已存在时是否覆盖
//  @return err
//
func Base64DataSaveToFIle(bs64Data string, filePath string, isCover bool) (err error) {
	isExists := CheckFileIsExist(filePath)
	if isExists && isCover == false {
		return nil
	}
	//目录检测
	tempPathSplitList := strings.SplitAfter(filePath, "/")
	dirPath := ""
	for _, v := range tempPathSplitList[:len(tempPathSplitList)-1] {
		dirPath = dirPath + v + "/"
	}
	if err := os.MkdirAll(dirPath, 0777); err != nil {
		return err
	}
	//数据校准
	//if len(bs64Data)%2!=0{
	//	bs64Data=bs64Data+"="
	//}
	//存储
	binData, _ := base64.StdEncoding.DecodeString(bs64Data) //成图片文件并把文件写入到buffer
	//存储后需要延时或进行校验
	if err = ioutil.WriteFile(filePath, binData, 0666); err != nil {
		return err
	}

	return nil
}

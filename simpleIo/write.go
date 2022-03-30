/*
 * @FileName:   write.go
 * @Author:		JuneXu
 * @CreateTime:	2022/3/30 下午10:08
 * @Description:
 */

package simpleIo

import (
	"io/ioutil"
	"os"
)

//
//  WirteByteToFile
//  @Description: byte格式数据写入到文件 。文件不存在则自动创建，内容会自动追到加值末尾。
//  @param filePath
//  @param data []byte
//  @return err
//
func WirteByteToFile(filePath string, data []byte) (err error) {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	defer file.Close()
	err = ioutil.WriteFile(filePath, data, os.ModePerm)
	return nil
}

//
//  WirteStrToFile
//  @Description: String格式数据写入到文件 。文件不存在则自动创建，内容会自动追到加值末尾。
//  @param filePath
//  @param data
//  @return err
//
func WirteStrToFile(filePath string, data string) (err error) {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err = file.WriteString(data); err != nil {
		return err
	}
	return nil
}

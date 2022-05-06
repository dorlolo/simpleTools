/*
 * @FileName:   file.go
 * @Author:		JuneXu
 * @CreateTime:	2022/3/12 下午5:35
 * @Description:
 */

package simpleWeb

import (
	"errors"
	"github.com/dorlolo/simpleTools/simpleCrypto"
	"io"
	"mime/multipart"
	"os"
	"path"
	"strings"
	"time"
)

//
//  SaveFile
//  @Description: 数据存储到文件
//  @param file： http请求中，form表单携带的文件数据
//  @param relativePath： 文件基于项目的相对路径
//  @return filePath : 文件的相对路径
//  @return code ： 文件名和时间戳拼接的md5编码
//  @return size ： 文件大小  单位kb
//  @return err
//
func SaveFile(file *multipart.FileHeader, saveDir string) (filePath string, code string, size float32, err error) {
	// 读取文件后缀
	ext := path.Ext(file.Filename)
	// 读取文件名并加密
	name := strings.TrimSuffix(file.Filename, ext)
	name = simpleCrypto.Md5Enscrypto(name)
	// 拼接新文件名
	filename := name + "_" + time.Now().Format("20060102150405") + ext
	// 尝试创建此路径
	mkdirErr := os.MkdirAll(saveDir, os.ModePerm)
	if mkdirErr != nil {
		err = errors.New("function os.MkdirAll() Filed, err:" + mkdirErr.Error())
		return
	}
	// 拼接路径和文件名
	var p string
	p = saveDir + "/" + filename
	f, openError := file.Open() // 读取文件
	if openError != nil {
		err = errors.New("function file.Open() Filed, err:" + openError.Error())
		return
	}
	defer f.Close() // 创建文件 defer 关闭

	out, createErr := os.Create(p)
	if createErr != nil {
		err = errors.New("function os.Create() Filed, err:" + createErr.Error())
		return
	}
	defer out.Close() // 创建文件 defer 关闭

	_, copyErr := io.Copy(out, f) // 传输（拷贝）文件
	//获取文件大小
	var tmpData []byte
	osize, _ := f.Read(tmpData)
	size = float32(osize) / 1024
	if copyErr != nil {
		return "", "", 0, errors.New("function io.Copy() Filed, err:" + copyErr.Error())
	}
	return p, filename, size, nil
}

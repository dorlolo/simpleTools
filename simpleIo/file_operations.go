package simpleIo

import (
	"encoding/base64"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"strings"
)

//@author: [songzhibin97](https://github.com/songzhibin97)
//@function: FileMove
//@description: 文件移动供外部调用
//@param: src string, dst string(src: 源位置,绝对路径or相对路径, dst: 目标位置,绝对路径or相对路径,必须为文件夹)
//@return: err error

func FileMove(src string, dst string) (err error) {
	if dst == "" {
		return nil
	}
	src, err = filepath.Abs(src)
	if err != nil {
		return err
	}
	dst, err = filepath.Abs(dst)
	if err != nil {
		return err
	}
	revoke := false
	dir := filepath.Dir(dst)
Redirect:
	_, err = os.Stat(dir)
	if err != nil {
		err = os.MkdirAll(dir, 0o755)
		if err != nil {
			return err
		}
		if !revoke {
			revoke = true
			goto Redirect
		}
	}
	return os.Rename(src, dst)
}

func DeLFile(filePath string) error {
	return os.RemoveAll(filePath)
}

//@author: [songzhibin97](https://github.com/songzhibin97)
//@function: TrimSpace
//@description: 去除结构体空格
//@param: target interface (target: 目标结构体,传入必须是指针类型)
//@return: null
func TrimSpace(target interface{}) {
	t := reflect.TypeOf(target)
	if t.Kind() != reflect.Ptr {
		return
	}
	t = t.Elem()
	v := reflect.ValueOf(target).Elem()
	for i := 0; i < t.NumField(); i++ {
		switch v.Field(i).Kind() {
		case reflect.String:
			v.Field(i).SetString(strings.TrimSpace(v.Field(i).String()))
		}
	}
}

// FileExist 判断文件是否存在
func FileExist(path string) bool {
	fi, err := os.Lstat(path)
	if err == nil {
		return !fi.IsDir()
	}
	return !os.IsNotExist(err)
}

//  Base64DataSaveToFIle
//  @Description: base64数据存储到文件
//  @param bs64Data string base64格式字符串
//  @param filePath string 文件路径
//  @param isCover bool 文件已存在时是否覆盖
//  @return err
//
func Base64DataSaveToFIle(bs64Data string, filePath string, isCover bool) (err error) {
	isExists := FileExist(filePath)
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

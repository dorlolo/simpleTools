/*
 * @FileName:   path.go
 * @Author:		JuneXu
 * @CreateTime:	2022/3/10 上午8:59
 * @Description:
 */

package simpleIo

import "os"

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

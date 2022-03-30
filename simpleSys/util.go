/*
 * @FileName:   util.go
 * @Author:		JuneXu
 * @CreateTime:	2022/3/30 下午8:28
 * @Description:
 */

package log

import (
	"os"
	"path/filepath"
	"strings"
)

//IsRelease 当前执行的是否是已编译文件
func IsRelease() bool {
	arg1 := strings.ToLower(os.Args[0])
	name := filepath.Base(arg1)

	return strings.Index(name, "__") != 0 && strings.Index(arg1, "go-build") < 0
}

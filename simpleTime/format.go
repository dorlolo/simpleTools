/*
 * @FileName:   format.go
 * @Author:		JuneXu
 * @CreateTime:	2022/2/25 下午2:30
 * @Description:
 */

package simpleTime

import "time"

var MonthMap = make(map[string]string)

type DefineTimeFormat struct {
	//常规时间格式(日期带横杠)
	Normal_YMDhms string
	Normal_YMDhm  string
	Normal_YMD    string
	Normal_hms    string
	//带斜杠的时间格式
	Slash_YMDhms string
	Slash_YMD    string
	//无间隔符
	NoSpacer_YMDhms string
	NoSpacer_YMD    string
	CN_YMDhms       string
	CN_YMDhm        string
	CN_YMD          string
	CN_hms          string
}

var Format DefineTimeFormat
var Loc *time.Location

func init() {
	MonthMap[""] = "00"
	MonthMap["January"] = "01"
	MonthMap["February"] = "02"
	MonthMap["March"] = "03"
	MonthMap["April"] = "04"
	MonthMap["May"] = "05"
	MonthMap["June"] = "06"
	MonthMap["July"] = "07"
	MonthMap["August"] = "08"
	MonthMap["September"] = "09"
	MonthMap["October"] = "10"
	MonthMap["November"] = "11"
	MonthMap["December"] = "12"

	Format = DefineTimeFormat{
		Normal_YMDhms:   "2006-01-02 15:04:05",
		Normal_YMDhm:    "2006-01-02 15:04",
		Normal_YMD:      "2006-01-02",
		Normal_hms:      "15:04:05",
		Slash_YMDhms:    "2006/01/02 15:04:05",
		Slash_YMD:       "2006/01/02",
		NoSpacer_YMDhms: "20060102150405",
		NoSpacer_YMD:    "20060102",

		CN_YMDhms: "2006年01月02日 15时04分05秒",
		CN_YMDhm:  "2006年01月02日 15时04分",
		CN_YMD:    "2006年01月02日",
		CN_hms:    "15时04分05秒",
	}
	Loc, _ = time.LoadLocation("Asia/Shanghai")
}

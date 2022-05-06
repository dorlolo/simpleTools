/*
 * @FileName:   getTime.go
 * @Author:		JuneXu
 * @CreateTime:	2022/3/1 下午6:35
 * @Description:
 */

package simpleTime

import (
	"fmt"
	"time"
)

//ThisMormingTime 今天凌晨
func ThisMorning(format string) (strTime string) {
	thisTime := time.Now()
	year := thisTime.Year()
	month := MonthMap[thisTime.Month().String()]
	day := fmt.Sprintf("%02d", thisTime.Day())
	strTime = fmt.Sprintf("%v-%v-%v 00:00:00", year, month, day)
	if format != Format.Normal_YMDhms {
		t1, _ := time.ParseInLocation(Format.Normal_YMDhms, strTime, Loc)
		strTime = t1.Format(format)
	}
	return strTime
}

func TomorrowMorning(baseTime time.Time) *time.Time {
	year := baseTime.Year()
	month := MonthMap[baseTime.Month().String()]
	day := fmt.Sprintf("%02d", baseTime.Day()+1)
	strTime := fmt.Sprintf("%v-%v-%v 00:00:00", year, month, day)
	res, _ := StringToTime(strTime)
	return res
}

//ThisMorningUnix 获取当日凌晨的时间戳
func ThisMorningToUnix() int64 {
	thist := time.Now()
	zero_tm := time.Date(thist.Year(), thist.Month(), thist.Day(), 0, 0, 0, 0, thist.Location()).Unix()
	return zero_tm
}

//ThisTimeUnix 获取当前时间的时间戳
func CurrentimeToUnix() int64 {
	return time.Now().Unix()
}

//Currentime 获取当前时间
func Currentime(format string) (strTime string) {
	strTime = time.Now().Format(format)
	return
}

//HoursAgo 若干小时之前的时间
func HoursAgo(hours time.Duration, format string) (lastTimeStr string) {
	lastStamp := time.Now().Unix() - int64((time.Hour * hours).Seconds())
	lastTime := time.Unix(lastStamp, 0).In(Loc)
	lastTimeStr = lastTime.Format(format)
	return
}

//MinutesAgo 若干分钟之前的时间
func MinutesAgo(minute time.Duration, format string) (lastTimeStr string) {
	lastStamp := time.Now().Unix() - int64((time.Minute * minute).Seconds())
	lastTime := time.Unix(lastStamp, 0).In(Loc)
	lastTimeStr = lastTime.Format(format)
	return
}

//SecondsAgo 若干秒之前的时间
func SecondsAgo(second time.Duration, format string) (lastTimeStr string) {
	lastStamp := time.Now().Unix() - int64((time.Minute * second).Seconds())
	lastTime := time.Unix(lastStamp, 0).In(Loc)
	lastTimeStr = lastTime.Format(format)
	return
}

//HoursAfter 若干小时之后的时间
func HoursAfter(hours time.Duration, format string) (lastTimeStr string) {
	lastStamp := time.Now().Unix() + int64((time.Hour * hours).Seconds())
	lastTime := time.Unix(lastStamp, 0).In(Loc)
	lastTimeStr = lastTime.Format(format)
	return
}

//MinutesAfter 若干分钟之后的时间
func MinutesAfter(minute time.Duration, format string) (lastTimeStr string) {
	lastStamp := time.Now().Unix() + int64((time.Minute * minute).Seconds())
	lastTime := time.Unix(lastStamp, 0).In(Loc)
	lastTimeStr = lastTime.Format(format)
	return
}

//SecondsAfter 若干秒之后的时间
func SecondsAfter(second time.Duration, format string) (lastTimeStr string) {
	lastStamp := time.Now().Unix() + int64((time.Minute * second).Seconds())
	lastTime := time.Unix(lastStamp, 0).In(Loc)
	lastTimeStr = lastTime.Format(format)
	return
}

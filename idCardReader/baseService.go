/*
 * @FileName:   idCardSlipt.go
 * @Author:		xjj
 * @CreateTime:	2021/12/14 上午10:05
 * @Description: 从证件号码中获取相关信息
 */
package idCardReader

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type IdCardService struct {
	Error error
}

/*
 * @Function   : GetBirthday
 * @Description: 获取出生日期
 * @Params     : id string 证件号码
 * @Return     :  string
 */
func (i *IdCardService) GetBirthday(id string) string {
	if id != "" {
		return fmt.Sprintf("%s-%s-%s", id[6:10], id[10:12], id[12:14])
	}
	return "1991-01-01"
}

/*
 * @Function   : GetAge
 * @Description: 获取出生日期
 * @Params     : :
 * @Return     :  string
 */
func (i *IdCardService) GetAge(id string) (age uint) {
	age = 25
	if id != "" {
		nowYear := time.Now().Year()
		thisYear, _ := strconv.Atoi(id[6:10])
		age = uint(nowYear - thisYear + 1)
		return age
	}
	return age
}

//
//  GetAge
//  @Description: 获取
//  @receiver i
//  @param id
//  @return age
//
func (i *IdCardService) GetSex(id string) (sex string) {
	//偶数为女性，奇书为男性
	mark, _ := strconv.Atoi(id[16:17])
	if mark%2 == 0 {
		return "女"
	}
	return "男"
}

/*
 * @Function   : GetArea
 * @Description: 获取省市区
 * @Params     :
 * @Return     : province string,city string,county string
 */
func (i *IdCardService) GetArea(id string) (province, city, county string) {
	i.Error = nil
	if id != "" {
		code := fmt.Sprintf("%s000000", id[:6]) //获取城市编码
		AreaDb, err := gorm.Open(sqlite.Open(DBPATH), &gorm.Config{})
		if err != nil {
			i.Error = err
			return
		}
		var countyData = ResCounty{}
		AreaDb.Debug().Where("code = ?", code).First(&countyData)
		//fmt.Println(countyData)
		var cityData = ResCity{}
		AreaDb.Where("id = ?", countyData.CityId).First(&cityData)
		var provinceData = ResProvince{}
		AreaDb.Where("id = ?", cityData.ProvinceId).First(&provinceData)
		province, city, county = provinceData.Name, cityData.Name, countyData.Name
	} else {
		province, city, county = "无信息", "无信息", "无信息"
		return province, city, county
	}
	return province, city, county
}

func (i *IdCardService) Check(id string) bool {
	//var sum = []int{}
	//for i := 0; i < 17; i++ {
	//	base, _ := strconv.Atoi(string(id[i]))
	//	sum := base * CHECK[i]
	//
	//}

	//for _, v := range id {
	//	fmt.Println(fmt.Printf("%q", v))
	//
	//}

	return false
}

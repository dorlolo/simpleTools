/*
 * @FileName:   srevice.go
 * @Author:		xjj
 * @CreateTime:	2021/12/22
 * @Description: 将省市区数据加载map中以供快速查询
 */
package idCard

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"strconv"
)

var (
	//IsInit 用于确认数据是否初始化到了下面三个map中
	IsInit        bool = false
	Res_Provice        = make(map[uint]ResProvince) //这里的uint是 在数据表中的id
	Res_CityMap        = make(map[uint]ResCity)     //这里的uint是 在数据表中的id
	Res_CountyMap      = make(map[uint]ResCounty)   //这里的uint是 区县编码里的内容！！！
)

func Init_areaInfoToMap() {
	AreaDb, err := gorm.Open(sqlite.Open("./utils/idCard/area.sqlite"), &gorm.Config{})
	if err != nil {
		fmt.Println("连接area数据库失败")
	}
	var countyData = []ResCounty{}
	AreaDb.Select("*").Find(&countyData)
	var cityData = []ResCity{}
	AreaDb.Select("*").Find(&cityData)
	var provinceData = []ResProvince{}
	AreaDb.Select("*").Find(&provinceData)
	for _, county := range countyData {
		Res_CountyMap[county.Code] = county
	}
	for _, city := range cityData {
		Res_CityMap[city.ID] = city
	}
	for _, province := range provinceData {
		Res_Provice[province.ID] = province
	}
	IsInit = true

}

/**
 * @Function   : GetAreaFromMap
 * @Description: 获取省市区
 * @Params     :
 * @Return     : province string,city string,county string
 */
func (i *IdCardService) GetAreaFromMap(id string) (province string, city string, county string) {
	if IsInit == false {
		fmt.Println("初始化数据表")
		Init_areaInfoToMap()
	}
	if id != "" {
		coutyCode, _ := strconv.Atoi(fmt.Sprintf("%s000000", id[:6])) //获取城市编码
		countyData := Res_CountyMap[uint(coutyCode)]
		if countyData.Name == "" {
			province, city, county = "无信息", "无信息", "无信息"
			return province, city, county
		}
		//fmt.Println(countyData)
		cityData := Res_CityMap[countyData.CityId]
		provinceData := Res_Provice[cityData.ProvinceId]
		province, city, county = provinceData.Name, cityData.Name, countyData.Name
	} else {
		province, city, county = "无信息", "无信息", "无信息"
		return province, city, county
	}
	return province, city, county
}

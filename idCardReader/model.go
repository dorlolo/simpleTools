/*
 * @FileName:   res_area.go
 * @Author:		xjj
 * @CreateTime:	2021/12/14 上午11:23
 * @Description: 此数据结构用于配合utils/icCard进行省市区的识别
 */
package idCardReader

//区县表
type ResCounty struct {
	ID       uint   `gorm:"primarykey"`
	Name     string `json:"name"`
	Code     uint   `json:"code"`
	CityId   uint   `json:"city_id"`
	AZ       string `json:"a_z"`
	Spelling string `json:"spelling"`
}

func (r ResCounty) TableName() string {
	return "res_county"
}

//市表
type ResCity struct {
	ID         uint   `gorm:"primarykey"`
	Name       string `json:"name"`
	Code       uint   `json:"code"`
	ProvinceId uint   `json:"province_id"`
	AZ         string `json:"a_z"`
	Spelling   string `json:"spelling"`
}

func (r ResCity) TableName() string {
	return "res_city"
}

//省表
type ResProvince struct {
	ID        uint   `gorm:"primarykey"`
	Name      string `json:"name"`
	Code      uint   `json:"code"`
	CountryId uint   `json:"country_id"`
	AZ        string `json:"a_z"`
	Spelling  string `json:"spelling"`
}

func (r ResProvince) TableName() string {
	return "res_Province"
}

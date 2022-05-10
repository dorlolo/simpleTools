/*
 * @FileName:   idcard_test.go
 * @Author:		JuneXu
 * @CreateTime:	2022/3/27 下午9:43
 * @Description:
 */

package test

import (
	"myExcample/idCard"
	"testing"
)

var (
	myidcard = "32058119910306341x"
)

//获取性别
func TestSex(t *testing.T) {
	var id = idCard.IdCardService{}
	sex := id.GetSex(myidcard)
	t.Log(sex)
}

//获取出生日期
func TestBirthday(t *testing.T) {
	var id = idCard.IdCardService{}
	sex := id.GetBirthday(myidcard)
	t.Log(sex)
}

//获取省市区
func TestArea(t *testing.T) {
	var id = idCard.IdCardService{}
	province, city, county := id.GetArea(myidcard)
	t.Log(province, city, county)
}

//从map中获取省市区信息，首次读取会把数据加载到缓存中，后续操作直接读取缓存
func TestAreaFromMap(t *testing.T) {
	var id = idCard.IdCardService{}
	province, city, county := id.GetAreaFromMap(myidcard)
	t.Log(province, city, county)
}

//获取年龄
func TestGet(t *testing.T) {
	var id = idCard.IdCardService{}
	age := id.GetAge(myidcard)
	t.Log(age)
}

func TestCheck(t *testing.T) {
	var id = idCard.IdCardService{}
	isok := id.Check(myidcard)
	t.Log(isok)

}

func TestSQLtest(t *testing.T) {
	laborsql := `select a.id ,wn.cn ,a.cc from ( `
	laborsql = laborsql + `SELECT work_name_id as id,COUNT(work_name_id) as cc `
	laborsql = laborsql + `FROM dtsite_iot.project_workers `
	laborsql = laborsql + `where work_name_id <> 0 `
	laborsql = laborsql + `group by work_name_id ) a left JOIN work_names wn on  wn.id  = a.id`

	t.Log(laborsql)
}

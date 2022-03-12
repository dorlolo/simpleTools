/*
 * @FileName:   dataHandler.go
 * @Author:		xjj
 * @CreateTime:	2021/12/20 下午4:28
 * @Description: 处理数据常用方法封装
 */
package utils

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"reflect"
)

/*
@Function:Isdatainlist
@Description:判断数据是否在切片/列表中
@Param: datalist 类型：interface{}，传入切片值
@Param: data 类型：interface{}，传入要查询的数据,一般为与datalist中元素一直的结构体
@Param: fieldList []string.需要对比的字段,注意这是一个切片. 传nil为完全对比
@Return: index int 数据在列表中的索引
@Return: result bool 数据是否在列表中
@author:JunjieXu
@Time:2021/12/22
*/
func Isdatainlist(datalist interface{}, data interface{}, fieldList []string) (index int, result bool) {
	interDatalist := reflect.ValueOf(datalist)
	compareValue := reflect.ValueOf(data)
	if interDatalist.Len() == 0 {
		return -1, false
	}
	//完全匹配
	if fieldList == nil {
		for i := 0; i < interDatalist.Len(); i++ {
			if interDatalist.Index(i).Interface() == compareValue.Interface() {
				return i, true
			}
		}
	} else {
		//只匹配字段值
		for i := 0; i < interDatalist.Len(); i++ {
			fieldCheck := true
			for _, field := range fieldList {
				//有一个字段值不匹配就失败
				if interDatalist.Index(i).FieldByName(field).Interface() != compareValue.FieldByName(field).Interface() {
					fieldCheck = false
				}
			}
			if fieldCheck == true {
				return i, true
			}
		}
	}
	return -1, false
}

/*
@Function:FindIndexInDataList
@Description:查询数据在切片/列表中的索引，就是Isdatainlist函数起得别名
@Param: datalist 类型：interface{}，传入切片值
@Param: data 类型：interface{}，传入要查询的数据,一般为与datalist中元素一直的结构体
@Param: fieldList []string.需要对比的字段,注意这是一个切片. 传nil为完全对比
@Return: index int 数据在列表中的索引
@Return: result bool 数据是否在列表中
@author:JunjieXu
@Time:2021/12/22
*/
func FindIndexInDataList(datalist interface{}, data interface{}, fieldList []string) (index int, result bool) {
	return Isdatainlist(datalist, data, fieldList)
}

/**
@Function:StructToMap
@Description:结构体转map
@Param:obj interface{} 结构体对象,!!不支持带有私有属性的结构体(字段首字母需要大写)!!
@Return:map[string]interface{}
@author:JunjieXu
@Time:2021/12/23
*/
func StructToMap(obj interface{}) (result map[string]interface{}) {
	obj1 := reflect.TypeOf(obj)
	obj2 := reflect.ValueOf(obj)
	result = make(map[string]interface{})
	for i := 0; i < obj1.NumField(); i++ {
		if obj2.Field(i).CanInterface() {
			result[obj1.Field(i).Name] = obj2.Field(i).Interface()
		}

	}
	return
}

/**
@Function:StructToMapUseTag
@Description:使用指定tag标签将结构体转map
@Param:obj interface{} 结构体对象，!!不支持带有私有属性的结构体(字段首字母需要大写)!!
@Param:tag string 标签
@Return:map[string]interface{}
@author:JunjieXu
@Time:2021/3/2
*/
func StructToMapUseTag(data interface{}, tag string) (result map[string]interface{}) {
	switch tag {
	case "json":
		if jsonData, err := json.Marshal(&data); err != nil {
			return
		} else {
			err = json.Unmarshal(jsonData, &result)
		}
	default:
		obj1 := reflect.TypeOf(data)
		obj2 := reflect.ValueOf(data)
		result = make(map[string]interface{})
		for i := 0; i < obj1.NumField(); i++ {
			if obj2.Field(i).CanInterface() {
				fmt.Println(obj1.Field(i).Tag.Get(tag))
				result[obj1.Field(i).Tag.Get(tag)] = obj2.Field(i).Interface()
			}
		}
	}

	return
}

//使用json标签将结构体转换为http请求的prams字符串结构
func StructToHttpParamsWithJson(data interface{}) (result string) {
	t := reflect.TypeOf(data)
	v := reflect.ValueOf(data)
	for i := 0; i < t.NumField(); i++ {
		jsonField := t.Field(i).Tag.Get("json")
		result = fmt.Sprintf("%v&%v=%v", result, jsonField, v.Field(i))
	}
	result = result[1:len(result)]
	return
}

/**
@Function:StructListToMap
@Description: 结构体列表转字典
@Param: obj （结构体对象）
@Param: fields （string/list/nil） 以哪几个字段的值组成key，传nil为所有
@Return: result map数据
@Return: ok  false/ture  失败/成功
@author:JunjieXu
@Time:2021/12/24
*/
func StructListToMap(obj interface{}, fields interface{}) (map[string]interface{}, bool) {
	var result = make(map[string]interface{})
	interVlaueObj := reflect.ValueOf(obj)
	switch fields.(type) {
	case string:
		//遍历切片获取结构体
		for i := 0; i < interVlaueObj.Len(); i++ {
			item := interVlaueObj.Index(i)
			//获取结构体字段值作为key
			key := item.FieldByName(fields.(string))
			result[key.String()] = item.Interface()
		}
	case []string:
		//遍历切片获取结构体
		for i := 0; i < interVlaueObj.Len(); i++ {
			item := interVlaueObj.Index(i)
			var fillKey string
			//取字段值组合为key
			for _, field := range fields.([]string) {
				value := item.FieldByName(field)
				fillKey = fmt.Sprintf("%v-%v", fillKey, value)
			}
			result[fillKey[1:]] = item.Interface()
		}
	case nil:
		//遍历切片获取结构体
		for i := 0; i < interVlaueObj.Len(); i++ {
			item := interVlaueObj.Index(i)
			var fillKey string
			//取所有的字段值组合为key
			for ii := 0; ii < item.NumField(); ii++ {
				fillKey = fmt.Sprintf("%v-%v", fillKey, item.Field(ii).Interface())
			}
			result[fillKey[1:]] = item.Interface()
		}
	}
	return result, true
}

/**
@Function:DiffList
@Description: 与另一个结构体数组/切片对比，获取相同值和不同值。
@Param: currentList 当前的结构体数组/切片
@Param: beforeList 需要对比的结构体数字/切片
@Param: fields (string/list/slice/nil)对比哪几个字段，传nil为对比所有数据
@Return: diffList 不同部分
@Return: sameList 相同部分
@Return: ok true/false 成功/失败
@author:JunjieXu
@Time:2021/12/24
*/
func DiffList(currentList interface{}, beforeList interface{}, fields interface{}) (diffList []interface{}, sameList []interface{}, ok bool) {
	m1, _ := StructListToMap(currentList, fields)
	m2, _ := StructListToMap(beforeList, fields)
	for key, _ := range m1 {
		if value, ok := m2[key]; ok == false {
			diffList = append(diffList, m1[key])
		} else {
			print(value)
			sameList = append(sameList, value)
		}
	}
	return diffList, sameList, true
}

//base64数据转字符串
func Base64DataToString(base64Data string) (result []byte, err error) {
	sEnc := base64.StdEncoding.EncodeToString([]byte(base64Data))
	buff, err := base64.StdEncoding.DecodeString(sEnc)
	if err != nil {
		return nil, err
	}
	return buff, nil
}
func StringToBase64(strData string) (result []byte) {
	var b bytes.Buffer
	w := base64.NewEncoder(base64.URLEncoding, &b)
	w.Write([]byte(strData))
	w.Close()
	result = b.Bytes()
	return
}

//列表/切片的排序
func Sort(dataList interface{}, column interface{}, desc bool) {
	panic("暂未开发")
}

//列表合并
func Merge(...interface{}) {
	panic("暂未开发")
}

//列表反转
func Reverse(dataList interface{}) {
	panic("暂未开发")

}

//返回末尾的n个元素
func Tail(dataList interface{}, n int) {
	panic("暂未开发")

}

//返回开头的n个元素
func Head(dataList interface{}, n int) {
	panic("暂未开发")

}

func If(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}

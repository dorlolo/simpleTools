package utils

import (
	"fmt"
	"testing"
)

func TestStructToHttpParamsWithJson(t *testing.T) {
	newstr := struct {
		name string `json:"name"`
		age  int    `json:"age"`
	}{
		"张三",
		16,
	}

	res := StructToHttpParamsWithJson(newstr)
	fmt.Println(res)
}

func TestStructToMap(t *testing.T) {
	//不支持的方法，里面的字段需要首字母大写
	var a = struct {
		User string
		name string
	}{
		"张三", "王德发",
	}
	data := StructToMap(a)
	fmt.Println(data)
}

//cannot return value obtained from unexported field or method [recovered]
func TestStructToMapUseTag(t *testing.T) {
	var a = struct {
		User string `json:"user"`
		Name string `json:"name,omitempty"`
	}{
		"张三", "",
	}
	data := StructToMapUseTag(a, "json")
	fmt.Println(data)
}

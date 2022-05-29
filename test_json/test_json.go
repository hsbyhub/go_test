package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name string `json:"StuName"`
	Age  int    `json:"age"`
	Sex  int    `json:"sex"`
}

func testPointStruct() {
	//指针变量
	stu := &Student{
		Name: "mike",
		Age:  19,
		Sex:  0,
	}
	//Marshal失败时err!=nil
	jsonStu, err := json.Marshal(*stu)
	if err != nil {
		fmt.Println("生成json字符串错误")
	}

	//jsonStu是[]byte类型，转化成string类型便于查看
	fmt.Println(string(jsonStu))

}

func testMapToJson() {
	paramsJsonStr := `{
			"domain": "video01.dgsltx.com",
			"ts": 1652842800,
			"regions": null,
			"isps": null,
			"provs": null,
			"nodes": null
		}`
	var params map[string]interface{}
	err := json.Unmarshal([]byte(paramsJsonStr), &params)
	if err != nil {
		fmt.Println("json.Unmarshal error [%v]", err)
		return
	}
	fmt.Println("json.Unmarshal result [%v]", params)

	var resultJson []byte
	resultJson, err = json.Marshal(params)
	if err != nil {
		fmt.Println("json.Marshal error [%v]", err)
		return
	}
	fmt.Println("json.Marshal return [%v]", string(resultJson))
}

func main() {
	testMapToJson()
}

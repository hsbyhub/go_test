package main

import (
    "encoding/json"
    "fmt"
)

type Student struct {
    Name    string  `json:"StuName"`
    Age     int
    Sex     string
    Class   *Class  `json:"StuClass"`
}

type Class struct {
    Name    string
    Grade   int
}

func TestJson() {
    //实例化一个数据结构，用于生成json字符串
    stu := Student{
        Name: "张三",
        Age:  18,
        Sex:  "男",
    }

    //指针变量
    cla := new(Class)
    cla.Name = "1班"
    cla.Grade = 3
    stu.Class=cla

    //Marshal失败时err!=nil
    jsonStu, err := json.Marshal(stu)
    if err != nil {
        fmt.Println("生成json字符串错误")
    }

    //jsonStu是[]byte类型，转化成string类型便于查看
    fmt.Println(string(jsonStu))
}

func TestMapToJson() {
    var mp = map[string]map[string]interface{}{
        "A" : map[string]interface{}{"a" : 1},
        "B" : map[string]interface{}{"a" : 1},
        "C" : map[string]interface{}{"a" : 1},
    }
    mp_json, err := json.Marshal(mp)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(string(mp_json))
}

func main() {
    TestMapToJson()
}

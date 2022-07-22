package test

import "fmt"

func main() {
	mp := make(map[int]interface{})
	mp[1] = "str"
	v, ok := mp[1].(int64)
	fmt.Println(v, ok)
}

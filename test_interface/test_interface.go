package main

func main() {
	mp map[string]string
	mpJson, _ := json.Marshal(mp)
	fmt.Println(mpJson)
}

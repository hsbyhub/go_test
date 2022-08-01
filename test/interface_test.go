package test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestInterface(t *testing.T) {
	var mp map[string]string
	mpJson, _ := json.Marshal(mp)
	fmt.Println(mpJson)
}

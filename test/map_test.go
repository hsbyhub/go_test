package test

import (
	"log"
	"testing"
)

func TestMap(t *testing.T) {
	mp := make(map[int]int)
	mp[100] += 999
	log.Println(mp[100])
}

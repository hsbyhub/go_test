package test

import (
	"fmt"
	"testing"
)

func TestFmt(t *testing.T) {
	v := []string{
		"abc", "efg", "hij",
	}
	fmt.Printf("%q \n", v)
}

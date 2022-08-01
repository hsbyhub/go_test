package test

import (
	"fmt"
	"os"
	"testing"
)

func TestOsArgs(t *testing.T) {
	for _, v := range os.Args {
		fmt.Println(v)
	}
}

package test

import (
	"fmt"
	"github.com/hsbyhub/gotools"
	"testing"
	"time"
)

func TestSleep(t *testing.T) {
	gotools.SetTimestamp()
	time.Sleep(5 * time.Second)
	fmt.Println(gotools.GetDuration())
}

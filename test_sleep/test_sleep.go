package test_sleep

import (
	"fmt"
	"github.com/hsbyhub/gotools"
	"time"
)

func main() {
	gotools.SetTimestamp()
	time.Sleep(5 * time.Second)
	fmt.Println(gotools.GetDuration())
}

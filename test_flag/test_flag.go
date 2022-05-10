/**
 * @File: test_flag.go
 * @Author: 数据体系-xiesenxin
 * @Date：2022/4/29
 */
package main

import (
	"flag"
	"os"
)
import "fmt"

var (
	h bool
	v bool
	d bool
	c string
)

func init() {
	flag.BoolVar(&h, "h", false, "this help")
	flag.BoolVar(&v, "v", false, "show version and exit")
	flag.BoolVar(&d, "d", false, "run tlog process by daemon mode")
	flag.StringVar(&c, "c", "/opt/bdtlog/tlog/conf/tlog.yaml", "set configuration `file`")
	flag.Usage = usage
}

func usage() {
	helpInfo := fmt.Sprintf("tlog version: tlog\nUsage: tlog [-hvd] [-c filename]\nOptions:")
	fmt.Fprintf(os.Stderr, helpInfo)
	flag.PrintDefaults()
}

func main() {
	flag.Parse()
	fmt.Println(h)
	fmt.Println(v)
	fmt.Println(d)
	fmt.Println(c)
}
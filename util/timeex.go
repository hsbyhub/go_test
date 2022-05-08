package util

import "time"

var timestamp int64

func Duration() int64 {
	tmp := timestamp
	timestamp = time.Now().Unix()
	return timestamp - tmp
}

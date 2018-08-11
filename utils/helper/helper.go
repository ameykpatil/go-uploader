package helper

import "time"

//GetCurrentTime return current time in millis
func GetCurrentTime() int64 {
	return time.Now().UnixNano() * int64(time.Nanosecond) / int64(time.Millisecond)
}

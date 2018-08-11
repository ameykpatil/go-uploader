package helper

import (
	"fmt"
	. "github.com/franela/goblin"
	. "github.com/onsi/gomega"
	"testing"
	"time"
)

func Test(t *testing.T) {

	g := Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("GetCurrentTime", func() {

		g.It("Should return current epoch time", func() {
			preEpoch := time.Now().UnixNano() * int64(time.Nanosecond) / int64(time.Millisecond)
			epoch := GetCurrentTime()
			postEpoch := time.Now().UnixNano() * int64(time.Nanosecond) / int64(time.Millisecond)

			time1 := time.Unix(0, preEpoch*int64(time.Millisecond))
			currentTime := time.Unix(0, epoch*int64(time.Millisecond))
			time2 := time.Unix(0, postEpoch*int64(time.Millisecond))

			fmt.Println(time1, currentTime, time2)
			Expect(currentTime.Before(time1)).ToNot(BeTrue())
			Expect(currentTime.After(time2)).ToNot(BeTrue())
		})
	})
}

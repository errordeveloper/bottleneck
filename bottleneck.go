package bottleneck

import (
	"math/rand"
	"time"
)

const (
	ShortSleepMin = 500
	ShortSleepMax = 1000
	LongSleepMin  = 3500
	LongSleepMax  = 5000
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func Sleep(min, max int) {
	time.Sleep(time.Duration(rand.Intn(max)+min) * time.Millisecond)
}

func ShortSleep() {
	Sleep(ShortSleepMin, ShortSleepMax)
}

func LongSleep() {
	Sleep(LongSleepMin, LongSleepMax)
}

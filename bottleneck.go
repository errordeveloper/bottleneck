package bottleneck

import (
	"math/rand"
	"time"
)

func Sleep(min, max int) {
	rand.Seed(time.Now().UnixNano())
	time.Sleep(time.Duration(rand.Intn(max)+min) * time.Millisecond)
}

func ShortSleep() {
	Sleep(500, 1000)
}

func LongSleep() {
	Sleep(3500, 5000)
}

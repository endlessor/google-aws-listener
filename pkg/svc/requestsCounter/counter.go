package requestsCounter

import "sync/atomic"

var requestsCount int64

func Increment() int64 {
	return atomic.AddInt64(&requestsCount, 1)
}

func GetNumber() int64 {
	return atomic.LoadInt64(&requestsCount)
}

func RestartCounter() {
	requestsCount = 0
}

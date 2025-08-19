package model

import "time"

type TimeSpan struct {
	Time100ms time.Duration
	Time500ms time.Duration
	Time1s    time.Duration
	Time2s    time.Duration
	Time3s    time.Duration
}

// Span contains timespans useful for sleep methods and can be mocked.
var Span = InitTimeSpan()

func InitTimeSpan() TimeSpan {
	return TimeSpan{
		Time100ms: 100 * time.Millisecond,
		Time500ms: 500 * time.Millisecond,
		Time1s:    1 * time.Second,
		Time2s:    2 * time.Second,
		Time3s:    3 * time.Second,
	}
}

package ptr

import "time"

func Bool(val bool) *bool {
	return &val
}

func Int(val int) *int {
	return &val
}

func String(val string) *string {
	return &val
}

func Time(val time.Time) *time.Time {
	return &val
}

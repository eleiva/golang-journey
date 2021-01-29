package examples

import "time"

type Time struct {}

func (t * Time) GetCurrentTime() time.Time {
	return time.Now()
}


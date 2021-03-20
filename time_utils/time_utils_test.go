package time_utils

import (
	"testing"
	"time"
)

func TestTimes(t *testing.T) {
	now := time.Now()
	date := time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local)
	
	t.Log(date.Before(now))
	
	t.Log(int64(now.Unix() - date.Unix()) / int64(24 * 3600))
	
	sub := date.Sub(now)
	t.Log(sub.Hours())
	t.Log(sub.String())
	
}

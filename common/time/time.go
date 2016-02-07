package time

import "time"

func Now() time.Time {
	return time.Unix(time.Now().Unix(), 0)
}

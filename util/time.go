package util

import "time"

func ParseDateTime(src string) time.Time {
	t, err := time.Parse(time.RFC3339, src)
	if err != nil {
		panic(err)
	}
	return t
}

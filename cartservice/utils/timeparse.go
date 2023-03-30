package utils

import "time"

func Parse_timestr_to_datetime(timeStr string) time.Time {
	t, error3 := time.Parse("2006-01-02", timeStr)
	if error3 != nil {
		panic(error3)
	}
	return t
}

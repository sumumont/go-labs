package utils

import (
	"fmt"
	"time"
)

func ParseTime(mils int64) time.Time {
	if mils < 0 {
		return time.Now()
	}

	return time.Unix(mils/1e3, mils%1e3*1e6)
}

func GetNowTime() int64 {
	return time.Now().UnixNano() / 1e6
}

var second = int64(1)
var min = 60 * second
var hour = 60 * min
var day = 24 * hour

func GetRunTime(timeLong int64) string {

	if timeLong < 1000 {
		return fmt.Sprintf("%+vms", timeLong)
	}
	timeLong = timeLong / 1000

	{
		t := timeLong / day
		if t > 0 {
			return fmt.Sprintf("%+vd", t)
		}
	}
	{
		t := timeLong / hour
		if t > 0 {
			return fmt.Sprintf("%+vh", t)
		}
	}
	{
		t := timeLong / min
		if t > 0 {
			return fmt.Sprintf("%+vm", t)
		}
	}
	return fmt.Sprintf("%+vs", timeLong)
}

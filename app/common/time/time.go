package time

import time2 "time"

type Time time2.Time

func (t Time) YmdHis() string {
	ti := time2.Time(t)
	return ti.Format("2006-01-02 15:04:05")
}

func Now() Time {
	return Time(time2.Now())
}

func NowPointer() *Time {
	t := Time(time2.Now())
	return &t
}

func (t Time) Ymd() string {
	ti := time2.Time(t)
	return ti.Format("2006-01-02")
}

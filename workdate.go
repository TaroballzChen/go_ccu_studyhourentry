package main

import (
	"strconv"
	"time"
)

type WorkTime struct {
	Year int
	Month int
	WorkHour int
}

func daysInMonth(year,month int) int {
	switch time.Month(month) {
	case time.April, time.June, time.September, time.November:
		return 30
	case time.February:
		if year%4 == 0 && (year%100 != 0 || year%400 == 0) { // leap year
			return 29
		}
		return 28
	default:
		return 31
	}
}

func IsWorkDay(year,month,day int)bool{
	t := time.Date(year+1911, time.Month(month),day,0,0,0,0,time.UTC).Weekday()
	if weekday := int(t); weekday == 0 || weekday == 6 {
		return false
	}
	return true
}

func AppendWorkday(year,month int)chan string{
	buffer := daysInMonth(year,month)
	workdaylist := make(chan string,60)
	for i:=1;i<=buffer;i++{
		if ok := IsWorkDay(year,month,i);ok {
			day := strconv.Itoa(i)
			workdaylist <- day
			workdaylist <- day
		}
	}
	close(workdaylist)
	return workdaylist
}
package other

import (
	"strconv"
	"time"
)

func dayOfTheWeek(day int, month int, year int) string {
	str := strconv.Itoa(year)
	if month < 10 {
		str += "0" + strconv.Itoa(month)
	} else {
		str += strconv.Itoa(month)
	}
	if day < 10 {
		str += "0" + strconv.Itoa(day)
	} else {
		str += strconv.Itoa(day)
	}
	t, _ := time.Parse("20060102", str)
	wd := t.Weekday()
	return wd.String()
}

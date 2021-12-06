package beijing_time

import "time"

var CST = time.FixedZone("CST", 8*60*60)

// 返回北京所在时区的当前时间
func Now() time.Time {
	return time.Now().In(CST)
}

// 获取指定时间对应日期的零点时间
func StartTimeOfDay(t time.Time) time.Time {
	beijingTime := t.In(CST)
	return time.Date(beijingTime.Year(), beijingTime.Month(), beijingTime.Day(), 0, 0, 0, 0, CST)
}

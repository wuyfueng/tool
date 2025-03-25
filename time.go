package tool

import "time"

var (
	LocalZone = time.Now().Location()
)

// NowDiffDays 距今多少天, 包含今天
func NowDiffDays(oldTime int64) (days int) {
	return DiffDays(time.Now(), time.Unix(oldTime, 0)) + 1
}

// DiffDays 相差天数 t1-t2
func DiffDays(t1, t2 time.Time) (days int) {
	t1 = time.Date(t1.Year(), t1.Month(), t1.Day(), 0, 0, 0, 0, LocalZone)
	t2 = time.Date(t2.Year(), t2.Month(), t2.Day(), 0, 0, 0, 0, LocalZone)

	return int(t1.Sub(t2).Hours() / 24)
}

// GetFiveTimestamp 时间戳转为当天5点
func GetFiveTimestamp(timestamp int64) int64 {
	datetime := time.Unix(timestamp, 0).Format("2006-01-02")
	t, _ := time.ParseInLocation("2006-01-02", datetime, LocalZone)
	return t.Unix() + 5*3600
}

// ThisMondayZeroTimestamp 本周一零点
func ThisMondayZeroTimestamp() int64 {
	now := time.Now()
	var deductDay int
	if now.Weekday() == time.Sunday {
		deductDay = 6
	} else {
		deductDay = int(now.Weekday()) - 1
	}
	return time.Date(now.Year(), now.Month(), now.Day()-deductDay, 0, 0, 0, 0, now.Location()).Unix()
}

// NextMondayZeroTimestamp 下周一零点
func NextMondayZeroTimestamp() int64 {
	now := time.Now()
	var addDay int
	if now.Weekday() == time.Sunday {
		addDay = 1
	} else {
		addDay = 8 - int(now.Weekday())
	}
	return time.Date(now.Year(), now.Month(), now.Day()+addDay, 0, 0, 0, 0, now.Location()).Unix()
}

// AppointTodayEndTimestamp 获取指定天的23:59:59秒时间戳
func AppointTodayEndTimestamp(day int64) int64 {
	datetime := time.Now().Format("2006-01-02")
	res, _ := time.ParseInLocation("2006-01-02", datetime, time.Local)
	return res.Unix() + day*86400
}

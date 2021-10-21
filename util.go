package main

import (
	"github.com/tidwall/gjson"
	"time"
)

func GetWeatherData(weatherData, name string) string {
	str := "daily.0." + name
	return gjson.Get(weatherData, str).String()
}

// 获取两个时间相差的天数，0表同一天，正数表End>start
func GetDiffDays(startTime, EndTime time.Time) int {
	return int(EndTime.Sub(startTime).Hours() / 24)
}

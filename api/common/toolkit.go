package common

import (
	"time"
	"strconv"
	"fmt"
	"os"
	"encoding/json"
)

//计算两个日期相差多少个月
func DiffMonth(from string,to string) int{
	year_from ,_:= strconv.Atoi(from[0:4])
	year_to,_:=strconv.Atoi(to[0:4])
	month1 ,_:= strconv.Atoi(from[5:7])
	month2,_:=strconv.Atoi(to[5:7])
	month_diff := (year_to-year_from-1)*12+12+month2-month1+1
	return month_diff
}
//计算两个日期之间有多少天
func DiffDay(from string,to string) int{
	t1,_:= time.Parse("2006-01-02",from)
	t2,_:= time.Parse("2006-01-02",to)
	timestamp1 := t1.Unix()
	timestamp2 := t2.Unix()
	day_inner :=  int((timestamp2-timestamp1)/(60*60*24))+1
	return day_inner
}


//判断一个月有几天
func CountDaysInEveryMonth(year int , month int) (days int) {
	year_string:= strconv.Itoa(year)
	month_string := strconv.Itoa(month)
	if year_string+"-"+ month_string == time.Now().Format("2006-01"){
		days = time.Now().Day()
		return
	}
	if month != 2 {
		if month == 4 || month == 6 || month == 9 || month == 11 {
			days = 30

		} else {
			days = 31
			fmt.Fprintln(os.Stdout, "The month has 31 days");
		}
	} else {
		if (((year % 4) == 0 && (year % 100) != 0) || (year % 400) == 0) {
			days = 29
		} else {
			days = 28
		}
	}
	return
}
//计算两个日期相差多少秒
func DiffSecond(from string , to string) (second int64){
	t1,_:= time.Parse("2006-01-02 15:04:05",from)
	t2,_:= time.Parse("2006-01-02 15:04:05",to)
	timestamp1 := t1.Unix()
	timestamp2 := t2.Unix()
	second =  timestamp2-timestamp1
	return
}


//判断一年有几天
func countDaysInEveryYear(year int) (days int) {
	//通常说平年365天,闰年366天
	if (((year % 4) == 0 && (year % 100) != 0) || (year % 400) == 0) {
		days = 366
	} else {
		days = 365
	}
	return
}

//把2018-05-28(utc)转换成2018-05-28T16:00:00Z这种形式
func TranslateTime1(t string) string{
	time_now_0,_:=time.ParseInLocation("2006-01-02 15:04:05",t+" 00:00:00",time.Local)
	time_now_0_utc := time_now_0.UTC().Format("2006-01-02T15:04:05Z")
	return  time_now_0_utc
}
//把2018-05-28 08(utc)转换成2018-05-28T16:00:00Z这种形式
func TranslateTime2(t string) string{
	time_now_0,_:=time.ParseInLocation("2006-01-02 15:04:05",t+":00:00",time.Local)
	time_now_0_utc := time_now_0.UTC().Format("2006-01-02T15:04:05Z")
	return  time_now_0_utc
}

//将接口类型 断言是json.Number -> float64
func TransforType (n interface{}) float64{
	a,_:=strconv.ParseFloat(string(n.(json.Number)), 64)
	return a
}
//将2006-01-02T15:04:05Z->2018-06-10 23
func TransforTimeByHour (t string)  string{
	s,_:=time.Parse("2006-01-02T15:04:05Z", t)
	s=s.In(time.Local)
	t1:=s.Format("2006-01-02 15")
	return t1
}
//将2006-01-02T15:04:05Z->2018-06-10
func TransforTimeByDay (t string)  string{
	s,_:=time.Parse("2006-01-02T15:04:05Z", t)
	s=s.In(time.Local)
	t1:=s.Format("2006-01-02")
	return t1
}
//专门为to准备的函数
func TransforTimeByDay_1 (t string)  string{
	s,_:=time.Parse("2006-01-02T15:04:05Z", t)
	s=s.In(time.Local).AddDate(0,0,-1)
	t1:=s.Format("2006-01-02")
	return t1
}
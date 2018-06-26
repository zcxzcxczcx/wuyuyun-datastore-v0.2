package common
//
//import (
//	"strconv"
//	"wuyuyun.com/data_store/api/datastructs"
//	"fmt"
//
//	"time"
//)
//
//
////计算两个日期之间有多少天
//func  CalDays1(startTime string,endTime string) int64{
//	fmt.Println(startTime)
//	fmt.Println(endTime)
//	fmt.Println(1)
//	var days int64
//	t1, err := time.ParseInLocation("2006-01-02", startTime, time.Local)
//	fmt.Println(2)
//	fmt.Println(t1)
//	t2, err := time.ParseInLocation("2006-01-02", endTime, time.Local)
//	fmt.Println(3)
//	fmt.Println(t2)
//	if err == nil && t1.Before(t2) {
//		diff := t2.Unix() - t1.Unix()
//		days = diff / 86400   //计算多少天
//		return days
//	} else {
//		return days
//		return days
//	}
//}
////计算两个日期之间有多少天
//func  CalDays(startTime string,endTime string) int64{
//	fmt.Println(startTime)
//	fmt.Println(endTime)
//	var days int64
//	t1, err := time.ParseInLocation("2006-01-02 15:04:05", startTime, time.Local)
//	t2, err := time.ParseInLocation("2006-01-02 15:04:05", endTime, time.Local)
//	if err == nil && t2.Before(t1) {
//		diff := t1.Unix() - t2.Unix()
//		days = diff / 86400   //计算多少天
//		return days
//	} else {
//		return days
//		return days
//	}
//}
////计算两个日期之间有多少年
//func  CalYears(startTime string,endTime string) int{
//	s1,_:=strconv.Atoi(endTime)
//	s2,_:=strconv.Atoi(startTime)
//	years:=s1 - s2
//	return years
//}
//
////计算两个日期之间有多少月份
//func  CalMonths(startTime string,endTime string) int{
//	s1:=HandleTime(startTime)
//	s2:=HandleTime(endTime)
//	months:=s2[1]+(s2[0]-s1[0]-1)*12+12-s1[1]
//	return months
//}
////把2018-02换成数组[2018,2]功能的函数
//func HandleTime(needHandleTime string) []int{
//	needReturnData :=make([]int,0,2)
//	rs := []rune(needHandleTime)
//	s1,_:=strconv.Atoi(string(rs[0:4]))
//	s2,_:=strconv.Atoi(string(rs[5:]))
//	needReturnData=append(needReturnData, s1,s2)
//	return needReturnData
//}
//
///*获取今天的功能项的数据用到的方法*/
////截取字符串 得到年月日
//func Substr(str string) string {
//	rs := []rune(str)
//	return string(rs[0:10])
//}
////截取字符串 得到时间点
//func Substr2(str string) int {
//	rs := []rune(str)
//	i,_:=strconv.Atoi(string(rs[11:]))
//	return i
//}
//////对获取到的数据进行排序
////func SortData(data []datastructs.DataAgr) []int{
////	index := make([]string, 0, 24)
////	for _,v:=range data{
////		index=append(index, v.TimePoint)
////	}
////	sort.Ints(index)
////	return index
////}
////整理每天获取到的数据
//func TodaySliceAfter(data []datastructs.DataAgr) []datastructs.DataAgr{
//	lenlen:=len(data)
//	todaySliceAfter := make([]datastructs.DataAgr, data[lenlen-1].TimePoint+1, data[lenlen-1].TimePoint+1)//用来存储今天的数据(整理后)
//	for i:=0;i<len(todaySliceAfter);i++{
//		todaySliceAfter[i].TimePoint= i
//	}
//	for j:=0;j<len(data);j++{
//		todaySliceAfter[data[j].TimePoint] = data[j]
//	}
//	return todaySliceAfter[0:len(todaySliceAfter)-1]
//}
////整理昨天获取到的数据
//func YesterdaySliceAfter(data []datastructs.DataAgr) []datastructs.DataAgr{
//	lenlen:=len(data)
//	yesterdaySliceAfter := make([]datastructs.DataAgr, data[lenlen-1].TimePoint+1, data[lenlen-1].TimePoint+1)//用来存储今天的数据(整理后)
//	for i:=0;i<len(yesterdaySliceAfter);i++{
//		yesterdaySliceAfter[i].TimePoint= i
//	}
//	for j:=0;j<len(data);j++{
//		yesterdaySliceAfter[data[j].TimePoint] = data[j]
//	}
//	return yesterdaySliceAfter
//}
/////*获取昨天的功能项的数据用到的方法*/
////func YesterdaySliceAfter(data []datastructs.DataAgr) []datastructs.DataAgr{
////	arraydata:=SortData(data)
////	fmt.Println(arraydata)
////	todaySliceAfter := make([]datastructs.DataAgr, 24, 24)//用来存储今天的数据
////	for i:=0;i<cap(todaySliceAfter);i++{
////		todaySliceAfter[i].TimePoint=i+1
////	}
////	for j:=0;j<len(arraydata);j++{
////		todaySliceAfter[arraydata[j]-1]=data[j]
////	}
////	return todaySliceAfter
////}
//
///*整理前七天的数据*/
//func SevenDays(sevDays []datastructs.DataAgrdays) []datastructs.DataAgrdays{
//	sevDays := make([]datastructs.DataAgrdays,168,168)
//	//
//
//}
//
///*整理前一个月的数据*/
//func OneMonthDays(oneMonthDays []datastructs.DataAgrday,days int64) []datastructs.DataAgrday{
//	oneMonthDay:= make([]datastructs.DataAgrday, days, days)//用来存储30天的数据
//	nTime := time.Now()
//	for i:=0;i<int(days);i++{
//		sevTime := nTime.AddDate(0,0,-i-1)
//		sevDay := sevTime.Format("2006-01-02")
//		oneMonthDay[i].TimePoint=sevDay
//	}
//	for i,v:=range oneMonthDay{
//		for i1,v1:=range oneMonthDays{
//			if v1.TimePoint==v.TimePoint{
//				oneMonthDay[i]=oneMonthDays[i1]
//			}
//		}
//	}
//	return oneMonthDay
//}
//
///*自定义按年取数据*/
////按年
//func DataByYearCustom(oneMonthDays []datastructs.CustomData,years int) []datastructs.CustomData{
//	oneMonthDay:= make([]datastructs.CustomData, years, years)//用来存储30天的数据
//	nTime := time.Now()
//	for i:=0;i<years;i++{
//		sevTime := nTime.AddDate(-i-1,0,0)
//		sevDay := sevTime.Format("2006")
//		oneMonthDay[i].TimePoint=sevDay
//	}
//	for i,v:=range oneMonthDay{
//		for i1,v1:=range oneMonthDays{
//			if v1.TimePoint==v.TimePoint{
//				oneMonthDay[i]=oneMonthDays[i1]
//			}
//		}
//	}
//	return oneMonthDay
//}
////按月
//func DataByMonthCustom(oneMonthDays []datastructs.CustomData,months int,endTime time.Time) []datastructs.CustomData{
//	oneMonthDay:= make([]datastructs.CustomData, months, months)//用来存储30天的数据
//
//	for i:=0;i<months;i++{
//		sevTime := endTime.AddDate(0,-i-1,0)
//		sevDay := sevTime.Format("2006-01")
//		oneMonthDay[i].TimePoint=sevDay
//	}
//	for i,v:=range oneMonthDay{
//		for i1,v1:=range oneMonthDays{
//			if v1.TimePoint==v.TimePoint{
//				oneMonthDay[i]=oneMonthDays[i1]
//			}
//		}
//	}
//	return oneMonthDay
//}
////按月
//func DataByDayCustom(oneMonthDays []datastructs.CustomData,days int64) []datastructs.CustomData{
//	oneMonthDay:= make([]datastructs.CustomData, days, days)//用来存储30天的数据
//	nTime := time.Now()
//	for i:=0;i<int(days);i++{
//		sevTime := nTime.AddDate(0,0,-i-1)
//		sevDay := sevTime.Format("2006-01-02")
//		oneMonthDay[i].TimePoint=sevDay
//	}
//	for i,v:=range oneMonthDay{
//		for i1,v1:=range oneMonthDays{
//			if v1.TimePoint==v.TimePoint{
//				oneMonthDay[i]=oneMonthDays[i1]
//			}
//		}
//	}
//	return oneMonthDay
//}
//

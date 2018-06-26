package restructure
import (
	"wuyuyun-datastore-v0.2/api/handler/protocol/db"
	"wuyuyun-datastore-v0.2/api/handler/protocol/datastruct"
	"wuyuyun-datastore-v0.2/api/common"
	"time"
	"fmt"
	"strconv"
	"sort"
	"math"
)
/**************************mysql **************************/
//重构获取到的今天的数据
func ListToday(devid string,protocolid int) datastruct.Todays{
	//定义重构数据的结构体
	var todaysDataRestructure datastruct.Todays
	//从数据库中获取到的数据
	todayDataFromDb := db.ListTodayDataFromDb(devid,protocolid)
	todaysDataRestructure.Date = time.Now().Format("2006-01-02")

	//如果todayDataFromDb的长度为0，则返回[]
	if len(todayDataFromDb)==0{
		todaysDataRestructure.Hours = todayDataFromDb
		return todaysDataRestructure
	}
	//如果todayDataFromDb的长度大于0,则初始化重构数据(todaysDataRestructure)的Min，Max，AvgMax，AvgMin
	if len(todayDataFromDb) > 0{
		todaysDataRestructure.Min = todayDataFromDb[0].Min
		todaysDataRestructure.Max =todayDataFromDb[0].Max
		todaysDataRestructure.AvgMax = todayDataFromDb[0].Avg
		todaysDataRestructure.AvgMin =todayDataFromDb[0].Avg
	}
	avg := 0.000
	for i:=0;i<len(todayDataFromDb);i++{
		//求平均数的和
		avg+=todayDataFromDb[i].Avg
		//如果从数据库中获取到的值有比todaysDataRestructure.Max 大，则更新todaysDataRestructure.Max
		if todayDataFromDb[i].Max > todaysDataRestructure.Max {
			todaysDataRestructure.Max =todayDataFromDb[i].Max
		}
		//如果从数据库中获取到的值有比todaysDataRestructure.Min 小，则更新todaysDataRestructure.Min
		if todayDataFromDb[i].Min < todaysDataRestructure.Min {
			todaysDataRestructure.Min =todayDataFromDb[i].Min
		}
		//如果从数据库中获取到的值有比todaysDataRestructure.AvgMax 大，则更新todaysDataRestructure.AvgMax
		if todayDataFromDb[i].Avg > todaysDataRestructure.AvgMax {
			todaysDataRestructure.AvgMax =todayDataFromDb[i].Avg
		}
		//如果从数据库中获取到的值有比todaysDataRestructure.AvgMin 小，则更新todaysDataRestructure.AvgMin
		if todayDataFromDb[i].Avg < todaysDataRestructure.AvgMin{
			todaysDataRestructure.AvgMin =todayDataFromDb[i].Avg
		}
	}
	//求todaysDataRestructure.Avg，数据库中所有avg的和除以len(todayDataFromDb)
	todaysDataRestructure.Avg = avg/float64(len(todayDataFromDb))
	hour:=time.Now().Hour()
	hoursRestructure:=make([]datastruct.Hours,hour,hour)
	for j:=0;j<len(hoursRestructure); j++ {
		hoursRestructure[j].Hour = j
	}
	fmt.Println(todayDataFromDb)
	for z:=0;z<len(todayDataFromDb); z++{
		hoursRestructure[todayDataFromDb[z].Hour]=todayDataFromDb[z]

	}
	todaysDataRestructure.Hours=hoursRestructure
	return todaysDataRestructure
}
//重构获取到的某一天的数据
func ListOneday(devid string,protocolid int,date string) datastruct.Todays{
	fmt.Println("2")
	var onetodaysDataRestructure datastruct.Todays
	todayDataFromDb := db.ListOneDataFromDb(devid,protocolid,date)
	onetodaysDataRestructure.Date=date
	if len(todayDataFromDb)==0{
		onetodaysDataRestructure.Hours = todayDataFromDb
		return onetodaysDataRestructure
	}
	if len(todayDataFromDb) > 0{
		onetodaysDataRestructure.Min = todayDataFromDb[0].Min
		onetodaysDataRestructure.Max =todayDataFromDb[0].Max
		onetodaysDataRestructure.AvgMax = todayDataFromDb[0].Avg
		onetodaysDataRestructure.AvgMin =todayDataFromDb[0].Avg
	}
	avg := 0.000
	for i:=0;i<len(todayDataFromDb);i++{
		avg+=todayDataFromDb[i].Avg
		if todayDataFromDb[i].Max > onetodaysDataRestructure.Max {
			onetodaysDataRestructure.Max =todayDataFromDb[i].Max
		}
		if todayDataFromDb[i].Min < onetodaysDataRestructure.Min {
			onetodaysDataRestructure.Min =todayDataFromDb[i].Min
		}
		if todayDataFromDb[i].Avg > onetodaysDataRestructure.AvgMax {
			onetodaysDataRestructure.AvgMax =todayDataFromDb[i].Avg
		}
		if todayDataFromDb[i].Avg < onetodaysDataRestructure.AvgMin {
			onetodaysDataRestructure.AvgMin =todayDataFromDb[i].Avg
		}
	}
	onetodaysDataRestructure.Avg = avg/float64(len(todayDataFromDb))

	hoursRestructure:=make([]datastruct.Hours,24,24)
	for j:=0;j<len(hoursRestructure); j++ {
		hoursRestructure[j].Hour = j
	}
	fmt.Println(todayDataFromDb)
	for z:=0;z<len(todayDataFromDb); z++{
		hoursRestructure[todayDataFromDb[z].Hour]=todayDataFromDb[z]

	}
	onetodaysDataRestructure.Hours=hoursRestructure
	return onetodaysDataRestructure
}

//重构获取到的七天的数据
func ListByhourRestructure(prodid string,protocolid int ,from string , to string) interface{}{
	var sureday datastruct.Sevdays
   //获取从数据库得到的还未整理过的数据
	sevendaysDataFromDb:=db.ListSevaldaysAgoFromDb( prodid, protocolid ,from ,to)
	sureday.From = sevendaysDataFromDb.From
	sureday.To = sevendaysDataFromDb.To
	sureday.Max= sevendaysDataFromDb.Max
	sureday.Min= sevendaysDataFromDb.Min
	sureday.Avg = sevendaysDataFromDb.Avg
	sureday.AvgMax = sevendaysDataFromDb.Dates[0].ValueAvg
	sureday.AvgMin = sevendaysDataFromDb.Dates[0].ValueAvg
	if len(sevendaysDataFromDb.Dates) ==  0{
		sureday.Dates =make([]datastruct.Datesbyhour,0)
		return sureday
	}
	var dates = make([]datastruct.Datesbyhour,0,5)
	day_inner :=common.DiffDay(from,to)//计算两个日期相差多少天
	for i:=0;i<day_inner;i++{
        var date datastruct.Datesbyhour
        date.Date = sevendaysDataFromDb.Dates[24*i].Date[0:10]
		var avg float64
		var num int
		var hours = make([]datastruct.Hours,0,5)
		avg=0
		num=0
		date.Min= sevendaysDataFromDb.Dates[24*i].ValueMin
		date.Max = sevendaysDataFromDb.Dates[24*i].ValueMax
		date.AvgMax= sevendaysDataFromDb.Dates[24*i].ValueAvg
		date.AvgMin = sevendaysDataFromDb.Dates[24*i].ValueAvg
		for j:=0;j<24;j++{
			var hour datastruct.Hours
			hour_int,_:=  strconv.Atoi(sevendaysDataFromDb.Dates[24*i+j].Date[11:])
			hour.Hour = hour_int
			hour.Max = sevendaysDataFromDb.Dates[24*i+j].ValueMax
			hour.Min= sevendaysDataFromDb.Dates[24*i+j].ValueMin
			hour.Avg= sevendaysDataFromDb.Dates[24*i+j].ValueAvg
			hours=append(hours,hour)
			if hour.Avg != 0{
				avg += hour.Avg
				num += 1
			}
			if date.AvgMin > hour.Avg {
				date.AvgMin = hour.Avg
			}
			if date.AvgMax < hour.Avg {
				date.AvgMax = hour.Avg
			}
			if date.Min > hour.Min {
				date.Min = hour.Min
			}
			if date.Max < hour.Max {
				date.Max = hour.Max
			}
		}
		if avg != 0{
			date.Avg = avg/float64(num)
		}else{
			date.Avg = avg
		}
		date.Hours =hours
		dates=append(dates,date)
		if sureday.AvgMax < date.Avg  {
			sureday.AvgMax = date.Avg
		}
		if sureday.AvgMin > date.Avg{
			sureday.AvgMin = date.Avg
		}
	}
	sureday.Dates= dates
	return sureday

}

//重构获取到的30天的数据
func ListBydayRestructure(prodid string,protocolid int ,from string , to string) interface{}{
	var sureday datastruct.Monthdays
	//获取从数据库得到的还未整理过的数据
	sevendaysDataFromDb:=db.ListSevaldaysAgoFromDb( prodid, protocolid ,from ,to)
	sureday.From = sevendaysDataFromDb.From
	sureday.To = sevendaysDataFromDb.To
	sureday.Max= sevendaysDataFromDb.Max//原始值中的Max
	sureday.Min= sevendaysDataFromDb.Min//原始值中的Min
	sureday.Avg = sevendaysDataFromDb.Avg//原始值中的Avg
	sureday.AvgMax = sevendaysDataFromDb.Dates[0].ValueAvg//平均值当中的AvgMax
	sureday.AvgMin = sevendaysDataFromDb.Dates[0].ValueAvg//平均值当中的AvgMin
	if len(sevendaysDataFromDb.Dates) ==  0{
		sureday.Dates =make([]datastruct.Dates,0)
		return sureday
	}
	day_inner :=common.DiffDay(from,to)//计算两个日期相差多少天
	var dates = make([]datastruct.Dates,0,5)
	for i:=0;i<day_inner;i++{
		var date datastruct.Dates
		date.Date = sevendaysDataFromDb.Dates[24*i].Date[0:10]
		var avg float64
		var num int
		avg=0
		num=0
		date.Min= sevendaysDataFromDb.Dates[24*i].ValueMin
		date.Max = sevendaysDataFromDb.Dates[24*i].ValueMax
		date.AvgMin= sevendaysDataFromDb.Dates[24*i].ValueAvg
		date.AvgMax= sevendaysDataFromDb.Dates[24*i].ValueAvg
		for j:=0;j<24;j++{
			var hour datastruct.Hours
			hour_int , _:=  strconv.Atoi(sevendaysDataFromDb.Dates[24*i+j].Date[11:])
			hour.Hour = hour_int
			hour.Max = sevendaysDataFromDb.Dates[24*i+j].ValueMax
			hour.Min= sevendaysDataFromDb.Dates[24*i+j].ValueMin
			hour.Avg= sevendaysDataFromDb.Dates[24*i+j].ValueAvg
			if hour.Avg != 0{
				avg += hour.Avg
				num += 1
			}
			if date.Min > hour.Min {
				date.Min = hour.Min
			}
			if date.Max < hour.Max {
				date.Max = hour.Max
			}
			if date.AvgMin > hour.Avg {
				date.AvgMin = hour.Avg
			}
			if date.AvgMax < hour.Avg {
				date.AvgMax = hour.Avg
			}

		}
		if avg != 0{
			date.Avg = avg/float64(num)
		}else{
			date.Avg = avg
		}
		dates=append(dates,date)
		if sureday.AvgMin > date.Avg{
			sureday.AvgMin = date.Avg
		}
		if sureday.AvgMax< date.Avg {
			sureday.AvgMax =date.Avg
		}
	}
	sureday.Dates = dates
	return sureday

}

//重构 按月返回 一般是from是2018-01或者2018，下面的是2018-01
func ListBySecondRestructure(devid string ,protocolid int , from string, to string) interface{}{
	var sureday datastruct.SureData
	//获取从数据库得到的还未整理过的数据
	dataFromDb:=db.GetDataFromDbBySeconds( devid, protocolid ,from ,to)
	sureday.From = dataFromDb.From
	sureday.To = dataFromDb.To
	sureday.Max= dataFromDb.Max
	sureday.Min= dataFromDb.Min
	sureday.Avg = dataFromDb.Avg
	if len(dataFromDb.Dates) ==  0{
		return dataFromDb
	}
	month_inner:= common.DiffMonth(from,to)//计算两个日期相差多少月
	fmt.Println("month_inner",month_inner)
	/************************** 初始化一个年的list **************************/
	year_flag ,_:= strconv.Atoi(dataFromDb.Dates[0].Date[0:4])
	month_flag ,_:= strconv.Atoi(dataFromDb.Dates[0].Date[5:7])
	var monthlist = make([]datastruct.Month,0,5)
	totolday:=0
	for a:=0;a<month_inner;a++{
		//计算这个月有几天
		thismonthday:=common.CountDaysInEveryMonth(year_flag,month_flag)
        var month_one datastruct.Month
		month_one.Month= dataFromDb.Dates[a*thismonthday+3].Date[0:7]
		num:=0
		avg := 0.000
        for b:=0;b<thismonthday;b++{
        	flag :=0
			var day_one datastruct.DateAgr
			day_one.Date=dataFromDb.Dates[totolday+b].Date
			day_one.ValueMax = dataFromDb.Dates[totolday+b].ValueMax
			day_one.ValueMin =  dataFromDb.Dates[totolday+b].ValueMin
			day_one.ValueAvg =  dataFromDb.Dates[totolday+b].ValueAvg
			if day_one.ValueMax != 0 && flag == 0{
				month_one.Max =day_one.ValueMax
				month_one.Min =day_one.ValueMin
				flag = 1
			}
			if flag == 1{
				if month_one.Max < day_one.ValueMax {
					month_one.Max = day_one.ValueAvg
				}
				if month_one.Min != 0 {
					if month_one.Min > day_one.ValueMin {
						month_one.Min = day_one.ValueMin
					}
				}
			}
			if day_one.ValueAvg != 0{
				avg += day_one.ValueAvg
				num += 1
			}

		}
		if avg != 0{
			month_one.Avg = avg/float64(num)
		}else{
			month_one.Avg = avg
		}
		if month_flag == 12{
			month_flag = 0
			year_flag  += 1
		}else{
			month_flag +=1
		}
		monthlist=append(monthlist,month_one )
		totolday+=thismonthday
	}

	sureday.Month = monthlist

	return sureday

}


/************************** influxdb **************************/
func ListFromInfluxdbByHour(from string,to string,devid string,protocolid string,n int) interface{}{
	fmt.Println(2)
   data:=db.GetDataFromInfluxdbByHour(from,to,devid,protocolid)
   if len(data) == 0{
   	return data
   }
	fmt.Println(3)
   var restructdata  datastruct.DataFromInfluxdb
	restructdata.From = common.TransforTimeByDay(from)
   if n == 1{
	   restructdata.To = common.TransforTimeByDay(to)
   }else{
	   restructdata.To = common.TransforTimeByDay_1(to)
   }


	fmt.Println(4)

	restructdata.Max=common.TransforType(data[0].Max)
	restructdata.Min=common.TransforType(data[0].Min)

	restructdata.AvgMax= common.TransforType(data[0].Mean)
	restructdata.AvgMin= common.TransforType(data[0].Mean)
	restructdata.MedianMax =common.TransforType(data[0].Median)
	restructdata.MedianMin = common.TransforType(data[0].Median)
	var index int
	var avg float64
	var median_arr []float64
	for i:=0;i<len(data);i++{

		data[i].Time=common.TransforTimeByHour(data[i].Time)
		//typessss:=reflect.TypeOf(data[i].Time)
		fmt.Println(5)
		avg+=common.TransforType(data[i].Mean)
		index++
		median_arr=append(median_arr,common.TransforType(data[i].Median) )

		if common.TransforType(data[i].Max)> restructdata.Max{
			restructdata.Max=common.TransforType(data[i].Max)
		}
		if common.TransforType(data[i].Min)< restructdata.Min{
			restructdata.Min=common.TransforType(data[i].Min)
		}
		if common.TransforType(data[i].Mean)> restructdata.AvgMax{
			restructdata.AvgMax=common.TransforType(data[i].Mean)
		}
		if common.TransforType(data[i].Mean)< restructdata.AvgMin{
			restructdata.AvgMin=common.TransforType(data[i].Mean)
		}

		if common.TransforType(data[i].Median)> restructdata.MedianMax{
			restructdata.MedianMax=common.TransforType(data[i].Median)
		}
		if common.TransforType(data[i].Median)< restructdata.MedianMin{
			restructdata.MedianMin=common.TransforType(data[i].Median)
		}

	}
	restructdata.Datas = data
	fmt.Println(6)
	restructdata.Avg = avg/float64(index)
	fmt.Println(7)
	sort.Float64s(median_arr) //升序排序

	restructdata.Median = median_arr[int(math.Floor ( float64(len(median_arr)/2) ))]
   return restructdata
}


func ListFromInfluxdbByDay(from string,to string,devid string,protocolid string) interface{}{
	fmt.Println(2)
	data:=db.GetDataFromInfluxdbByDay(from,to,devid,protocolid)
	if len(data) == 0{
		return data
	}
	var restructdata  datastruct.DataFromInfluxdb
	restructdata.From = common.TransforTimeByDay(from)
	restructdata.To = common.TransforTimeByDay_1( to)
	fmt.Println(3)
	//data的时间间隔是8小时，重新整理成时间间隔为一天
	dataRestucture :=make([]datastruct.Datas,len(data)/3)
	for j:=0;j<len(data)/3;j++{
		time:=common.TransforTimeByHour(data[j*3].Time)[0:10]
		mean:=(common.TransforType(data[j*3].Mean)+common.TransforType(data[j*3+1].Mean)+common.TransforType(data[j*3+2].Mean))/3
		median := common.TransforType(data[j*3+1].Median)
		max:=math.Max(math.Max(common.TransforType(data[j*3].Max),common.TransforType(data[j*3+1].Max)),common.TransforType(data[j*3+2].Max))
		min:=math.Min(math.Max(common.TransforType(data[j*3].Min),common.TransforType(data[j*3+1].Min)),common.TransforType(data[j*3+2].Min))
		dataRestucture[j]=datastruct.Datas{time,mean,median,max,min}
	}

	fmt.Println(4)

	restructdata.Max=dataRestucture[0].Max.(float64)
	restructdata.Min=dataRestucture[0].Min.(float64)

	restructdata.AvgMax= dataRestucture[0].Mean.(float64)
	restructdata.AvgMin= dataRestucture[0].Mean.(float64)
	restructdata.MedianMax =dataRestucture[0].Median.(float64)
	restructdata.MedianMin = dataRestucture[0].Median.(float64)
	var index int
	var avg float64
	var median_arr []float64
	for i:=0;i<len(dataRestucture);i++{

		//typessss:=reflect.TypeOf(data[i].Time)

		avg+=dataRestucture[i].Mean.(float64)
		index++
		median_arr=append(median_arr,dataRestucture[i].Median.(float64) )

		if dataRestucture[i].Max.(float64)> restructdata.Max{
			restructdata.Max=dataRestucture[i].Max.(float64)
		}
		if dataRestucture[i].Min.(float64)< restructdata.Min{
			restructdata.Min=dataRestucture[i].Min.(float64)
		}
		if dataRestucture[i].Mean.(float64)> restructdata.AvgMax{
			restructdata.AvgMax=dataRestucture[i].Mean.(float64)
		}
		if dataRestucture[i].Mean.(float64)< restructdata.AvgMin{
			restructdata.AvgMin=dataRestucture[i].Mean.(float64)
		}

		if dataRestucture[i].Median.(float64)> restructdata.MedianMax{
			restructdata.MedianMax=dataRestucture[i].Median.(float64)
		}
		if dataRestucture[i].Median.(float64)< restructdata.MedianMin{
			restructdata.MedianMin=dataRestucture[i].Median.(float64)
		}

	}
	restructdata.Datas = dataRestucture
	fmt.Println(6)
	restructdata.Avg = avg/float64(index)
	fmt.Println(7)
	sort.Float64s(median_arr) //升序排序

	restructdata.Median = median_arr[int(math.Floor ( float64(len(median_arr)/2) ))]
	return restructdata
}
package restructure
import(
	"wuyuyun-datastore-v0.2/api/handler/onlineRate/db"
	"wuyuyun-datastore-v0.2/api/handler/onlineRate/datastruct"
	"wuyuyun-datastore-v0.2/api/common"
	"time"
	"fmt"
	"strconv"
)

//重构所有项目的在线率数据
// 粒度是天
func RestructureOnlinerateDataProjAllByDay(from string,to string) interface{}{
	daydiff:= common.DiffDay(from,to)
	//var RateData_arr []datastruct.RateData
	rateData_arr:=db.ProjAllOnlineRateDataFromDbByDay(from,daydiff)
	return rateData_arr
}
//重构所有项目的在线率数据
// 粒度是小时
func  RestructureOnlinerateDataProjAllByHour(from string,to string) interface{}{
	var rateData datastruct.RateDataByHour
	daydiff:= common.DiffDay(from,to)
	RateData_arr :=make([]datastruct.RateDataByHour,0)
	from_later ,_:=time.ParseInLocation("2006-01-02",from ,time.Local)
	for a:=0;a<daydiff;a++{
		from_todb:=from_later.AddDate(0,0,a).Format("2006-01-02")
		var dataFromDb []datastruct.RateData
		dataFromDb =db.ProjAllOnlineRateDataFromDbByHour(from_todb)
		rateData.TimePoint = from_todb
		rateData.OnlineRateArr=dataFromDb
		RateData_arr=append(RateData_arr,rateData)
	}
	return RateData_arr
}

//重构某个项目的在线率数据 粒度是天
func RestructureOnlinerateDataProjByDay(from string,to string,projectid int) interface{}{
	daydiff:= common.DiffDay(from,to)
	rateData_arr:=db.ProjOnlineRateDataFromDbByDay(projectid,from,daydiff)
	return rateData_arr
}
//重构某个项目的在线率数据 粒度是小时
func RestructureOnlinerateDataProjByHour(from string,to string,projectid int) interface{}{
	fmt.Println("3")
	var rateData datastruct.RateDataByHour
	daydiff:= common.DiffDay(from,to)
	RateData_arr :=make([]datastruct.RateDataByHour,0)
	from_later ,_:=time.ParseInLocation("2006-01-02",from ,time.Local)
	for a:=0;a<daydiff;a++{
		from_todb:=from_later.AddDate(0,0,a).Format("2006-01-02")
		var dataFromDb []datastruct.RateData
		fmt.Println("4")
			dataFromDb =db.ProjOnlineRateDataFromDbByHour(projectid,from_todb)
			rateData.TimePoint = from_todb
			rateData.OnlineRateArr=dataFromDb
			RateData_arr=append(RateData_arr,rateData)
	}
	return RateData_arr
}


//重构某个产品的在线率数据
// 粒度是天
func RestructureOnlinerateDataProdByDay(from string,to string,productid int) interface{}{
	daydiff:= common.DiffDay(from,to)
	rateData_arr:=db.ProdOnlineRateDataFromDbByDay(productid,from,daydiff)
	return rateData_arr
}
func RestructureOnlinerateDataProdByHour(from string,to string,productid int) interface{}{
	var rateData datastruct.RateDataByHour
	daydiff:= common.DiffDay(from,to)
	RateData_arr :=make([]datastruct.RateDataByHour,0)
	from_later ,_:=time.ParseInLocation("2006-01-02",from ,time.Local)
	for a:=0;a<daydiff;a++{
		from_todb:=from_later.AddDate(0,0,a).Format("2006-01-02")
		var dataFromDb []datastruct.RateData
		dataFromDb =db.ProjOnlineRateDataFromDbByHour(productid,from_todb)
		rateData.TimePoint = from_todb
		rateData.OnlineRateArr=dataFromDb
		RateData_arr=append(RateData_arr,rateData)
	}
	return RateData_arr
}



//重构某个设备的在线率数据 from to 形式 粒度是天,可以到今天
func RestructureOnlinerateDataDevFromToByDay(deviceid string,from string,to string) interface{}{
	var rateData datastruct.RateData
	from_later,_:=time.Parse("2006-01-02",from)
	//to_later,_:=time.Parse("2006-01-02",to)
	daydiff:= common.DiffDay(from,to)
	RateData_arr :=make([]datastruct.RateData,0)
	for a:=0;a<daydiff;a++{
		from_todb:=from_later.AddDate(0,0,a).Format("2006-01-02")
		var dataFromDb int64
		if from_todb == time.Now().Format("2006-01-02"){
			dataFromDb=db.DeviceOnlineRateDataFromDb(deviceid,from_todb+" 00:00:00",time.Now().Format("2006-01-02 15:04:05"))
		}else{
			dataFromDb=db.DeviceOnlineRateDataFromDb(deviceid,from_todb+" 00:00:00",from_todb+ " 23:59:59")
		}
		fmt.Println("dataFromDb",dataFromDb)
		rateData.TimePoint = from_todb
		rateData.OnlineRate = float64(dataFromDb)*100/(24*60*60)
		if rateData.OnlineRate < 0{
			rateData.OnlineRate = 0
		}
		if rateData.OnlineRate > 100{
			rateData.OnlineRate = 100
		}
		fmt.Println("rateData.OnlineRate",float64(dataFromDb/(24*60*60)))
		RateData_arr=append(RateData_arr,rateData )
	}

	return RateData_arr
}
//重构某个设备的在线率数据 from to 形式 粒度是小时
func RestructureOnlinerateDataDevFromToByHour(deviceid string,from string,to string) interface{}{
	var rateData datastruct.RateData
	from_later,_:=time.Parse("2006-01-02",from)
	//to_later,_:=time.Parse("2006-01-02",to)
	daydiff:= common.DiffDay(from,to)
	RateData_arr :=make([]datastruct.RateData,0)
	for a:=0;a<daydiff;a++{
		from_todb_date:=from_later.AddDate(0,0,a).Format("2006-01-02")
		var dataFromDb int64
		var len int
		if from_todb_date == time.Now().Format("2006-01-02"){
			len= time.Now().Hour()
		}else{
			len =24
		}
		for b:=0;b<len;b++{
			var from_todb string
			var to_todb string
			if b<10{
				from_todb = from_todb_date + " 0" + strconv.Itoa(b) +":00:00"
				to_todb = from_todb_date + " 0" + strconv.Itoa(b) +":59:59"
			}
			if b >= 10{
				from_todb = from_todb_date + " " + strconv.Itoa(b) +":00:00"
				to_todb = from_todb_date + " " + strconv.Itoa(b) +":59:59"
			}
			dataFromDb=db.DeviceOnlineRateDataFromDb(deviceid,from_todb,to_todb)
			fmt.Println("dataFromDb",dataFromDb)
			rateData.TimePoint = from_todb[0:13]
			rateData.OnlineRate = float64(dataFromDb)*100/(60*60)
			if rateData.OnlineRate < 0{
				rateData.OnlineRate = 0
			}
			if rateData.OnlineRate > 100{
				rateData.OnlineRate = 100
			}
			fmt.Println("rateData.OnlineRate",float64(dataFromDb)*100/(60*60))
			RateData_arr=append(RateData_arr,rateData )
		}

	}

	return RateData_arr
}

package protocol

import (
	"net/http"
	"wuyuyun-datastore-v0.2/api/handler/protocol/restructure"
	"wuyuyun-datastore-v0.2/api/common"
	"github.com/gorilla/mux"
	"strconv"
	"fmt"
	"log"
	"encoding/json"
	"time"
)
/**************************mysql **************************/
//获取今天功能项的数据的handler
func ListToday(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	devid := vars["devid"]
	protocolid := vars["protocolid"]
	protocolid_atoi,err:=strconv.Atoi(protocolid)
	if err != nil {
		fmt.Println("strconv.Atoi(protocolid)出现错误!")
		//log.Fatal(err)
		common.JsonResponse(err.Error(),500,w)
	}
	data:=restructure.ListToday(devid,protocolid_atoi)

	//data_jsonMarshal,err_marshal:=json.Marshal(data)
	//if err_marshal != nil {
	//	fmt.Println("json.Marshal(data)出现错误!")
	//	server.JsonResponse(err_marshal.Error(),500,w)
	//}
	// 发送一个状态代码的HTTP响应头。
	common.JsonResponse(data,200,w)

}

//获取昨天功能项的数据的handler
func ListYesterday(w http.ResponseWriter, r *http.Request){
	fmt.Println("1")
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	vars := mux.Vars(r)
	devid := vars["devid"]
	protocolid := vars["protocolid"]
	protocolid_atoi,err:=strconv.Atoi(protocolid)
	if err != nil {
		common.JsonResponse(err.Error(),500,w)
	}
	nTime := time.Now()
	fmt.Println(nTime)
	yesTime := nTime.AddDate(0,0,-1).Format("2006-01-02")
	fmt.Println(yesTime)
	data:=restructure.ListOneday(devid,protocolid_atoi,yesTime)

	common.JsonResponse(data,200,w)

}


//获取前七天的功能项的数据的handler
func ListSevendaysAgo(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	vars := mux.Vars(r)
	devid := vars["devid"]
	protocolid := vars["protocolid"]
	//from := r.FormValue("from")
	//to := r.FormValue("to")
	//fmt.Println("from",from)
	//fmt.Println("to",to)
	from:=time.Now().AddDate(0,0,-7).Format("2006-01-02")
	to:=time.Now().AddDate(0,0,-1).Format("2006-01-02")
	day_inner :=common.DiffDay(from,to)//计算两个日期相差多少天
	fmt.Println("day_inner", day_inner)
	if day_inner!=7 {
		common.JsonResponse("传入的时间间隔必须为七天",500,w)

	}

	protocolid_atoi,err:=strconv.Atoi(protocolid)
	if err != nil {
		fmt.Println("strconv.Atoi(protocolid)出现错误!")
		common.JsonResponse("strconv.Atoi(protocolid)出现错误!",500,w)
	}
	data:=restructure.ListByhourRestructure(devid,protocolid_atoi,from,to)
	common.JsonResponse(data,200,w)
}

//获取前三十天的功能项的数据的handler
func ListOneMonthAgo(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	vars := mux.Vars(r)
	devid := vars["devid"]
	protocolid := vars["protocolid"]
	//from := r.FormValue("from")
	//to := r.FormValue("to")
	from:=time.Now().AddDate(0,0,-30).Format("2006-01-02")
	to:=time.Now().AddDate(0,0,-1).Format("2006-01-02")
	fmt.Println("from",from)
	fmt.Println("to",to)
	day_inner :=common.DiffDay(from,to)//计算两个日期相差多少天
	fmt.Println("day_inner", day_inner)
	if day_inner!=30 {
		common.JsonResponse("传入的时间间隔必须为30天",500,w)

	}
	protocolid_atoi,err:=strconv.Atoi(protocolid)
	if err != nil {
		fmt.Println("strconv.Atoi(protocolid)出现错误!")
		common.JsonResponse("strconv.Atoi(protocolid)出现错误!",500,w)

	}
	data:=restructure.ListBydayRestructure(devid,protocolid_atoi,from,to)
	common.JsonResponse(data,200,w)
}

//按月份返回，测试接口
func MonthReturnTest(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	vars := mux.Vars(r)
	devid := vars["devid"]
	protocolid := vars["protocolid"]
	from := r.FormValue("from")
	to := r.FormValue("to")
	fmt.Println("from",from)
	fmt.Println("to",to)
    //不能是当月
	fmt.Println("2016-01",time.Now().Format("2006-01"))
	if to==time.Now().Format("2006-01") {
		msg,_:=json.Marshal("不能选择当月")
		w.Write(msg)
		return

	}
	protocolid_atoi,err:=strconv.Atoi(protocolid)
	if err != nil {
		fmt.Println("strconv.Atoi(protocolid)出现错误!")
		log.Fatal(err)
		return
	}
	data:=restructure.ListBySecondRestructure(devid,protocolid_atoi,from,to)
	data_jsonMarshal,err_marshal:=json.Marshal(data)
	if err_marshal != nil {
		fmt.Println("json.Marshal(data)出现错误!")
		fmt.Println("err_marshal",err_marshal)
	}
	w.Write(data_jsonMarshal)
	return
}


/************************** influxdb **************************/


func ListToday2(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	devid := vars["devid"]
	protocolid := vars["protocolid"]
	// 时间 -> 2018-05-27T16:00:00Z
	t1:=common.TranslateTime1(time.Now().Format("2006-01-02"))
	t2:=common.TranslateTime2(time.Now().Format("2006-01-02 15"))
	fmt.Println(t1)
	fmt.Println(t2)
	data:=restructure.ListFromInfluxdbByHour(t1, t2 ,devid,protocolid ,1)
	common.JsonResponse(data,200,w)
}
func ListYesterday2(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	devid := vars["devid"]
	protocolid := vars["protocolid"]
	// 时间 -> 2018-05-27T16:00:00Z
	t1:=common.TranslateTime1(time.Now().AddDate(0,0,-1).Format("2006-01-02"))
	t2:=common.TranslateTime1(time.Now().Format("2006-01-02"))
	fmt.Println(t1)
	fmt.Println(t2)
	data:=restructure.ListFromInfluxdbByHour(t1, t2 ,devid,protocolid,2)
	common.JsonResponse(data,200,w)
}
func ListSevendaysAgo2(w http.ResponseWriter, r *http.Request){
	fmt.Println(1)
	vars := mux.Vars(r)
	devid := vars["devid"]
	protocolid := vars["protocolid"]
	// 时间 -> 2018-05-27T16:00:00Z
	t1:=common.TranslateTime1(time.Now().AddDate(0,0,-7).Format("2006-01-02"))
	t2:=common.TranslateTime1(time.Now().Format("2006-01-02"))
	fmt.Println(t1)
	fmt.Println(t2)
	data:=restructure.ListFromInfluxdbByHour(t1, t2 ,devid,protocolid,2)
	//fmt.Println("最后的数据是",data)
	common.JsonResponse(data,200,w)
}
func ListOneMonthAgo2(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	devid := vars["devid"]
	protocolid := vars["protocolid"]
	// 时间 -> 2018-05-27T16:00:00Z
	t1:=common.TranslateTime1(time.Now().AddDate(0,0,-30).Format("2006-01-02"))
	t2:=common.TranslateTime1(time.Now().Format("2006-01-02"))
	fmt.Println(t1)
	fmt.Println(t2)
	data:=restructure.ListFromInfluxdbByDay(t1, t2 ,devid,protocolid)
	common.JsonResponse(data,200,w)
}

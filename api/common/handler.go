package common
//
//import (
//	"github.com/gorilla/mux"
//	"net/http"
//	"wuyuyun.com/data_store/store"
//	"log"
//	"fmt"
//	"encoding/json"
//	"wuyuyun.com/data_store/api/common"
//	"wuyuyun.com/data_store/api/datastructs"
//	"io/ioutil"
//	"time"
//)
//
//
//////获取今天的功能项的数据的handler
////func ListTest(w http.ResponseWriter, r *http.Request){
////	db , err := store.ConectDb2()
////	defer db.Close()
////	if err != nil {
////		fmt.Println("连接数据库出现错误!")
////		return
////	}
////	rows, err := db.Query("SELECT FROM_UNIXTIME(create_time/1000000000,'%Y-%m-%d %H'),AVG(value) as value  FROM product_data WHERE id = 7084 AND agr_id = 43 GROUP BY FROM_UNIXTIME(create_time/1000000000,'%Y-%m-%d %H')")
////	if err != nil {
////		log.Fatal(err)
////		return
////	}
////	defer rows.Close()
////
////	var value string//时间点
////	var value2 float64//采集的平均值
////
////	for rows.Next(){
////		err := rows.Scan(&value,&value2)
////
////		//要今天的数据
////		fmt.Println(value,value2)
////		if err != nil {
////			log.Fatal(err)
////			return
////		}
////	}
////	err = rows.Err()
////	if err != nil {
////		log.Fatal(err)
////	}
////
////}
//
////获取今天的功能项的数据的handler
//func ListToday(w http.ResponseWriter, r *http.Request){
//	todaySliceBefore := make([]datastructs.DataAgr, 0, 5)//用来存储今天的数据
//	vars := mux.Vars(r)
//	prodid := vars["prodid"]//产品id
//	protocolid := vars["protocolid"]//产品协议项id
//	db , err1 := store.ConectDb2()
//	defer db.Close()
//	if err1 != nil {
//		fmt.Println("连接数据库出现错误!")
//		return
//	}
//
//	rows, err2 := db.Query("SELECT FROM_UNIXTIME(create_time/1000000000,'%H'), MAX(value) as valuemax,MIN(value) as valuemin,AVG(value) as valueavg  FROM product_data WHERE pro_id = ? AND agr_id = ? AND FROM_UNIXTIME(create_time/1000000000,'%Y-%m-%d') = DATE_FORMAT(NOW(),'%Y-%m-%d') GROUP BY FROM_UNIXTIME(create_time/1000000000,'%H')",prodid,protocolid)
//	if err2 != nil {
//		fmt.Println("查询出现错误!")
//		log.Fatal(err2)
//		return
//	}
//	defer rows.Close()
//
//	var hour int//时间点
//	var valuemax float64//采集的平均值
//	var valuemin float64//采集的平均值
//	var valueavg float64//采集的平均值
//
//	for rows.Next(){
//		err3 := rows.Scan(&hour,&valuemax,&valuemin,&valueavg)
//		if err3 != nil {
//			fmt.Println("rows.Scan出现错误!")
//			log.Fatal(err3)
//			return
//		}
//		//要今天的数据
//		todaySliceBefore=append(todaySliceBefore,datastructs.DataAgr{hour,valuemax,valuemin,valueavg})
//
//	}
//	err4 := rows.Err()
//	if err4 != nil {
//		fmt.Println("遍历出现错误!")
//		log.Fatal(err4)
//	}
//	if len(todaySliceBefore) == 0{
//		noData,_:=json.Marshal("今天没有采集到数据哦")
//		w.Write(noData)
//	}else{
//		//整理成前端需要的数据格式
//		todaySliceAfter:=common.TodaySliceAfter(todaySliceBefore)
//		dd,_:=json.Marshal(todaySliceAfter)
//		w.Write(dd)
//	}
//}
//
////获取昨天的功能项的数据的handler
//func ListYesterday(w http.ResponseWriter, r *http.Request){
//	yesterdaySliceBefore := make([]datastructs.DataAgr, 0, 5)//用来存储昨天的数据
//	vars := mux.Vars(r)
//	prodid := vars["prodid"]//产品id
//	protocolid := vars["protocolid"]//产品协议项id
//	db , err := store.ConectDb2()
//	defer db.Close()
//	if err != nil {
//		fmt.Println("连接数据库出现错误!")
//		return
//	}
//	rows, err1 := db.Query("SELECT FROM_UNIXTIME(create_time/1000000000,'%H'), MAX(value) as valuemax,MIN(value) as valuemin,AVG(value) as valueavg  FROM product_data WHERE pro_id = ? AND agr_id = ? AND FROM_UNIXTIME(create_time/1000000000,'%Y-%m-%d') = DATE_SUB(curdate(),INTERVAL 1 DAY) GROUP BY FROM_UNIXTIME(create_time/1000000000,'%H')",prodid,protocolid)
//	if err1 != nil {
//		fmt.Println("查询出现错误!")
//		log.Fatal(err)
//		return
//	}
//	defer rows.Close()
//	var hour int//时间点
//	var valuemax float64//采集的平均值
//	var valuemin float64//采集的平均值
//	var valueavg float64//采集的平均值
//
//	for rows.Next(){
//		err3 := rows.Scan(&hour,&valuemax,&valuemin,&valueavg)
//		if err3 != nil {
//			fmt.Println("rows.Scan出现错误!")
//			log.Fatal(err3)
//			return
//		}
//		//要今天的数据
//		yesterdaySliceBefore=append(yesterdaySliceBefore,datastructs.DataAgr{hour,valuemax,valuemin,valueavg})
//
//	}
//	err4 := rows.Err()
//	if err4 != nil {
//		fmt.Println("遍历出现错误!")
//		log.Fatal(err4)
//	}
//	if len(yesterdaySliceBefore) == 0{
//		noData,_:=json.Marshal("今天没有采集到数据哦")
//		w.Write(noData)
//	}else{
//		//整理成前端需要的数据格式
//		yesterdaySliceAfter:=common.YesterdaySliceAfter(yesterdaySliceBefore)
//		dd,_:=json.Marshal(yesterdaySliceAfter)
//		w.Write(dd)
//	}
//
//
//}
//
////获取前七天的功能项的数据的handler
//func ListSevendaysAgo(w http.ResponseWriter, r *http.Request){
//	sevDays := make([]datastructs.DataAgrday, 0, 5)//用来存储昨天的数据
//	var parameter datastructs.Parameter
//	body, _ := ioutil.ReadAll(r.Body)
//	json.Unmarshal(body, &parameter)
//	db , err := store.ConectDb2()
//	defer db.Close()
//	if err != nil {
//		fmt.Println("连接数据库出现错误!")
//		return
//	}
//	//得到今天的日期
//	nTime := time.Now()
//	yesterDay := nTime.Format("2006-01-02")
//	t1, err := time.Parse("2006-01-02",yesterDay)
//	fmt.Println(t1)
//
//	//得到第八天前的日期
//	sevTime := nTime.AddDate(0,0,-8)
//	sevDay := sevTime.Format("2006-01-02")
//	t2, err := time.Parse("2006-01-02",sevDay)
//	//t12:=t2.UnixNano()
//	fmt.Println(t2)
//	rows, err := db.Query("SELECT FROM_UNIXTIME(create_time/1000000000,'%Y-%m-%d %H'), MAX(value) as valuemax,MIN(value) as valuemin,AVG(value) as valueavg  FROM product_data WHERE  pro_id = ? AND agr_id = ? AND FROM_UNIXTIME(create_time/1000000000,'%Y-%m-%d') BETWEEN DATE_SUB(curdate(),INTERVAL 7 DAY) AND DATE_SUB(curdate(),INTERVAL 1 DAY) GROUP BY FROM_UNIXTIME(create_time/1000000000,'%Y-%m-%d %H')",prodid,protocolid)
//	if err != nil {
//		log.Fatal(err)
//		return
//	}
//	defer rows.Close()
//	var value string//时间点
//	var value2 float64//采集的平均值
//
//	for rows.Next(){
//		err := rows.Scan(&value,&value2)
//
//		t, err := time.Parse("2006-01-02",value)
//		fmt.Println(t)
//		fmt.Println(value,value2)
//	    if t.Before(t1)&& t.After(t2){
//			sevDays=append(sevDays, datastructs.DataAgrday{value,value2})
//		}
//		if err != nil {
//			log.Fatal(err)
//			return
//		}
//	}
//	fmt.Println(sevDays)
//	err = rows.Err()
//	if err != nil {
//		log.Fatal(err)
//	}
//	if len(sevDays) == 0{
//		noData,_:=json.Marshal("这七天没有采集到数据哦")
//		w.Write(noData)
//	}else {
//		ss := common.SevenDays(sevDays)
//		dd, _ := json.Marshal(ss)
//		w.Write(dd)
//	}
//}
//
//////获取前一个月的功能项的数据的handler
////func ListOneMonthAgo(w http.ResponseWriter, r *http.Request){
////	oneMonthDays := make([]datastructs.DataAgrday, 0, 5)//用来存储昨天的数据
////	var parameter datastructs.Parameter
////	body, _ := ioutil.ReadAll(r.Body)
////	json.Unmarshal(body, &parameter)
////	db , err := store.ConectDb2()
////	defer db.Close()
////	if err != nil {
////		fmt.Println("连接数据库出现错误!")
////		return
////	}
////	//得到今天的日期
////	nTime := time.Now()
////	yesterDay := nTime.Format("2006-01-02")
////	t1, err := time.Parse("2006-01-02",yesterDay)
////
////	//得到第一个月前的日期
////	oneMonthTime := nTime.AddDate(0,-1,0)
////	oneMonthDay := oneMonthTime.Format("2006-01-02")
////	t2, err := time.Parse("2006-01-02",oneMonthDay)
////	rows, err := db.Query("SELECT FROM_UNIXTIME(create_time/1000000000,'%Y-%m-%d'),AVG(value) as value  FROM product_data WHERE pro_id = ? AND agr_id = ? GROUP BY FROM_UNIXTIME(create_time/1000000000,'%Y-%m-%d')",parameter.ProId,parameter.AgrId)
////	if err != nil {
////		log.Fatal(err)
////		return
////	}
////	defer rows.Close()
////	var value string//时间点
////	var value2 float64//采集的平均值
////
////	for rows.Next(){
////		err := rows.Scan(&value,&value2)
////		t, err := time.Parse("2006-01-02",value)
////		if t.Before(t1)&& t.After(t2){
////			oneMonthDays=append(oneMonthDays, datastructs.DataAgrday{value,value2})
////		}
////		if err != nil {
////			log.Fatal(err)
////			return
////		}
////	}
////	fmt.Println(oneMonthDays)
////	err = rows.Err()
////	if err != nil {
////		log.Fatal(err)
////	}
////	if len(oneMonthDays) == 0{
////		noData,_:=json.Marshal("前一个月没有采集到数据哦")
////		w.Write(noData)
////	}else {
////		fmt.Println(t1)
////		fmt.Println(t2)
////		days := common.CalDays(t1.Format("2006-01-02 15:04:05"),t2.Format("2006-01-02 15:04:05"))
////		fmt.Println(days)
////		ss := common.OneMonthDays(oneMonthDays, days)
////		dd, _ := json.Marshal(ss)
////		w.Write(dd)
////	}
////
////}
////
//////获取前三个月的功能项的数据的handler
////func ListThreeMonthAgo(w http.ResponseWriter, r *http.Request){
////	oneMonthDays := make([]datastructs.DataAgrday, 0, 5)//用来存储昨天的数据
////	var parameter datastructs.Parameter
////	body, _ := ioutil.ReadAll(r.Body)
////	json.Unmarshal(body, &parameter)
////	db , err := store.ConectDb2()
////	defer db.Close()
////	if err != nil {
////		fmt.Println("连接数据库出现错误!")
////		return
////	}
////	//得到今天的日期
////	nTime := time.Now()
////	yesterDay := nTime.Format("2006-01-02")
////	t1, err := time.Parse("2006-01-02",yesterDay)
////	fmt.Println(t1)
////
////	//得到第三个月前的日期
////	oneMonthTime := nTime.AddDate(0,-3,0)
////	oneMonthDay := oneMonthTime.Format("2006-01-02")
////	t2, err := time.Parse("2006-01-02",oneMonthDay)
////	//t12:=t2.UnixNano()
////	fmt.Println(t2)
////	rows, err := db.Query("SELECTFROM_UNIXTIME(create_time/1000000000,'%Y-%m-%d'),AVG(value) as value  FROM product_data WHERE pro_id = ? AND agr_id = ? GROUP BY FROM_UNIXTIME(create_time/1000000000,'%Y-%m-%d')",parameter.ProId,parameter.AgrId)
////	if err != nil {
////		log.Fatal(err)
////		return
////	}
////	defer rows.Close()
////	var value string//时间点
////	var value2 float64//采集的平均值
////
////	for rows.Next(){
////		err := rows.Scan(&value,&value2)
////
////		t, err := time.Parse("2006-01-02",value)
////		fmt.Println(t)
////		fmt.Println(value,value2)
////		if t.Before(t1)&& t.After(t2){
////			oneMonthDays=append(oneMonthDays, datastructs.DataAgrday{value,value2})
////		}
////		if err != nil {
////			log.Fatal(err)
////			return
////		}
////	}
////	fmt.Println(oneMonthDays)
////	err = rows.Err()
////	if err != nil {
////		log.Fatal(err)
////	}
////	if len(oneMonthDays) == 0{
////		noData,_:=json.Marshal("前三个月没有采集到数据哦")
////		w.Write(noData)
////	}else {
////		days := common.CalDays( t1.Format("2006-01-02 15:04:05"),t2.Format("2006-01-02 15:04:05"))
////		fmt.Println(days)
////		ss := common.OneMonthDays(oneMonthDays, days)
////		dd, _ := json.Marshal(ss)
////		w.Write(dd)
////	}
////}
//////获取自定义时间的功能项的数据的handler
////func ListCustomAgo(w http.ResponseWriter, r *http.Request){
////	customData := make([]datastructs.CustomData, 0, 5)//用来存储数据
////	var customtime datastructs.CustomTimeRes
////	body, _ := ioutil.ReadAll(r.Body)
////	json.Unmarshal(body, &customtime)
////	fmt.Fprintf(w, "获取json中的ProId: %s\n", customtime.ProId)
////	fmt.Fprintf(w, "获取json中的AgrId: %s\n", customtime.AgrId)
////	fmt.Fprintf(w, "获取json中的Type: %s\n", customtime.Type)
////	fmt.Fprintf(w, "获取json中的StartTime: %s\n", customtime.StartTime)
////	fmt.Fprintf(w, "获取json中的EndTime: %s\n", customtime.EndTime)
////
////	db , err := store.ConectDb2()
////	defer db.Close()
////	if err != nil {
////		fmt.Println("连接数据库出现错误!")
////		return
////	}
////	if customtime.Type == 1{
////		startTime,_:=time.Parse("2006",customtime.StartTime)
////		endTime,_:=time.Parse("2006",customtime.EndTime)
////		years:=common.CalYears(customtime.StartTime,customtime.EndTime)
////		rows, err := db.Query("FROM_UNIXTIME(create_time/1000000000,'%Y'),AVG(value) as value  FROM product_data WHERE pro_id = ? AND agr_id = ? GROUP BY FROM_UNIXTIME(create_time/1000000000,'%Y')",customtime.ProId, customtime.AgrId)
////		if err != nil {
////			log.Fatal(err)
////			return
////		}
////		defer rows.Close()
////		var value string//时间点
////		var value2 float64//采集的平均值
////
////		for rows.Next(){
////			err := rows.Scan(&value,&value2)
////			t, err := time.Parse("2006",value)
////			fmt.Println(t)
////			fmt.Println(value,value2)
////			if t.Before(endTime)&& t.After(startTime){
////				customData=append(customData, datastructs.CustomData{value,value2})
////			}
////			if err != nil {
////				log.Fatal(err)
////				return
////			}
////		}
////		fmt.Println(customData)
////		err = rows.Err()
////		if err != nil {
////			log.Fatal(err)
////		}
////		ss:=common.DataByYearCustom(customData,years)
////		dd,err:=json.Marshal(ss)
////		w.Write(dd)
////	}else if customtime.Type == 2{
////		startTime,_:=time.Parse("2006-01",customtime.StartTime)
////		endTime,_:=time.Parse("2006-01",customtime.EndTime)
////        months:=common.CalMonths(customtime.StartTime,customtime.EndTime)
////		rows, err := db.Query("SELECT FROM_UNIXTIME(create_time/1000000000,'%Y-%m'),AVG(value) as value  FROM product_data WHERE pro_id = ? AND agr_id = ? GROUP BY FROM_UNIXTIME(create_time/1000000000,'%Y-%m')",customtime.ProId, customtime.AgrId)
////		if err != nil {
////			log.Fatal(err)
////			return
////		}
////		defer rows.Close()
////		var value string//时间点
////		var value2 float64//采集的平均值
////
////		for rows.Next(){
////			err := rows.Scan(&value,&value2)
////			t, err := time.Parse("2006-01",value)
////			fmt.Println(t)
////			fmt.Println(value,value2)
////			if t.Before(endTime)&& t.After(startTime){
////				customData=append(customData, datastructs.CustomData{value,value2})
////			}
////			if err != nil {
////				log.Fatal(err)
////				return
////			}
////		}
////		fmt.Println(customData)
////		err = rows.Err()
////		if err != nil {
////			log.Fatal(err)
////		}
////		s,_:=time.Parse("2006-01",customtime.EndTime)
////		ss:=common.DataByMonthCustom(customData,months,s)
////		dd,err:=json.Marshal(ss)
////		w.Write(dd)
////
////	}else{
////		startTime,_:=time.Parse("2006-01-02",customtime.StartTime)
////		endTime,_:=time.Parse("2006-01-02",customtime.EndTime)
////		days:=common.CalDays(customtime.StartTime,customtime.EndTime)
////		fmt.Println(days)
////		rows, err := db.Query("SELECT FROM_UNIXTIME(create_time/1000000000,'%Y-%m-%d'),AVG(value) as value  FROM product_data WHERE pro_id = ? AND agr_id = ? GROUP BY FROM_UNIXTIME(create_time/1000000000,'%Y-%m-%d')",customtime.ProId, customtime.AgrId)
////		if err != nil {
////			log.Fatal(err)
////			return
////		}
////		defer rows.Close()
////		var value string//时间点
////		var value2 float64//采集的平均值
////
////		for rows.Next(){
////			err := rows.Scan(&value,&value2)
////			t, err := time.Parse("2006-01-02",value)
////			fmt.Println(t)
////			fmt.Println(value,value2)
////			if t.Before(endTime)&& t.After(startTime){
////				customData=append(customData, datastructs.CustomData{value,value2})
////			}
////			if err != nil {
////				log.Fatal(err)
////				return
////			}
////		}
////		fmt.Println(customData)
////		err = rows.Err()
////		if err != nil {
////			log.Fatal(err)
////		}
////		ss:=common.DataByDayCustom(customData,days)
////		dd,err:=json.Marshal(ss)
////		w.Write(dd)
////	}
////
////}
//
//

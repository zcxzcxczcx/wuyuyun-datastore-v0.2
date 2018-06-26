package db

import (
	"fmt"
	"log"
	"wuyuyun-datastore-v0.2/api/handler/onlineRate/datastruct"
	"database/sql"
	"time"
	"wuyuyun-datastore-v0.2/api/common"
	"wuyuyun-datastore-v0.2/conn"
)
//从数据库获取所有项目
//粒度是天
func ProjAllOnlineRateDataFromDbByDay(from string ,day_inner int) []datastruct.RateData{
	//所有项目有多少个设备
	devNum :=GetAllDevInProj(-1)
	rateData:= make([]datastruct.RateData,0)
	if devNum == 0{
		return rateData
	}
	db , err := conn.ConectDb5()
	defer db.Close()
	if err != nil {
		fmt.Println("连接数据库出现错误!")
	}

	rows, err2 := db.Query(`CALL online_rate_by_day(?,?)`,from,day_inner)

	if err2 != nil {
		fmt.Println("查询出现错误!")
		log.Fatal(err2)
	}
	defer rows.Close()

	var TimePoint string//时间点
	var number int//在线的设备
	for rows.Next(){
		err3 := rows.Scan(&TimePoint,&number)
		if err3 != nil {
			fmt.Println("rows.Scan出现错误!")
			log.Fatal(err3)
		}
		rate := float64(number)/float64(devNum)
		//要某一天的数据
		rateData=append(rateData,datastruct.RateData{TimePoint,rate})

	}
	err4 := rows.Err()
	if err4 != nil {
		fmt.Println("遍历出现错误!")
		log.Fatal(err4)
	}
	return rateData
}
//粒度是时
func ProjAllOnlineRateDataFromDbByHour(date string) []datastruct.RateData{
	//所有项目有多少个设备
	devNum :=GetAllDevInProj(-1)
	rateData:= make([]datastruct.RateData,0)
	if devNum == 0{
		return rateData
	}
	db , err := conn.ConectDb5()
	defer db.Close()
	if err != nil {
		fmt.Println("连接数据库出现错误!")
	}
	var num int
	if date ==time.Now().Format("2006-01-02"){
		num = time.Now().Hour()
	} else{
		num = 24
	}
	fmt.Println("num",num)
	fmt.Println("date",date)
	rows, err2 := db.Query(`CALL online_rate_by_hour(?,?)`,num ,date)

	if err2 != nil {
		fmt.Println("查询出现错误!")
		log.Fatal(err2)
	}
	defer rows.Close()
	var TimePoint string//时间点
	var number int//上线的时长
	for rows.Next(){

		err3 := rows.Scan(&TimePoint,&number)

		if err3 != nil {
			fmt.Println("rows.Scan出现错误!")
			log.Fatal(err3)
		}
		fmt.Println("number",number)
		fmt.Println("devNum",devNum)
		onlineRate := float64(number)/float64(devNum)
		fmt.Println("onlineRate",onlineRate)
		rateData = append(rateData,datastruct.RateData{TimePoint,onlineRate} )
	}
	err4 := rows.Err()
	if err4 != nil {
		fmt.Println("遍历出现错误!")
		log.Fatal(err4)
	}

	return rateData
}



//从数据库获取单个项目
//粒度是天
func ProjOnlineRateDataFromDbByDay(projectid int,from string ,day_inner int) []datastruct.RateData {
	//单个项目有多少个设备
	rateData:= make([]datastruct.RateData,0)
	devNum :=GetAllDevInProj(projectid)
	if devNum == 0{
		return rateData
	}
	db , err := conn.ConectDb5()
	defer db.Close()
	if err != nil {
		fmt.Println("连接数据库出现错误!")
	}

		rows, err2 := db.Query(`CALL online_rate_by_day_aproj(?,?,?)`,from,day_inner,projectid)

	if err2 != nil {
		fmt.Println("查询出现错误!")
		log.Fatal(err2)
	}
	defer rows.Close()

	var TimePoint string//时间点
	var number int//在线的设备
	for rows.Next(){
		err3 := rows.Scan(&TimePoint,&number)
		if err3 != nil {
			fmt.Println("rows.Scan出现错误!")
			log.Fatal(err3)
		}
		rate := float64(number)/float64(devNum)
		//要某一天的数据
		rateData=append(rateData,datastruct.RateData{TimePoint,rate})

	}
	err4 := rows.Err()
	if err4 != nil {
		fmt.Println("遍历出现错误!")
		log.Fatal(err4)
	}
	return rateData
}
//从数据库获取单个项目
//粒度是小时
func ProjOnlineRateDataFromDbByHour(projectid int,date string) []datastruct.RateData{
	fmt.Println("5")
	//某个项目有多少个设备,
	rateData:= make([]datastruct.RateData,0)
	devNum :=GetAllDevInProj(projectid)
	if devNum == 0{
		fmt.Println("6")
		return rateData
	}
	db , err := conn.ConectDb5()
	defer db.Close()
	if err != nil {
		fmt.Println("连接数据库出现错误!")
	}
	

	var num int
	if date ==time.Now().Format("2006-01-02"){
		num = time.Now().Hour()
	} else{
		num = 24
	}
	fmt.Println("dddd")
	fmt.Println("num",num)
	fmt.Println("date",date)
	rows, err2 := db.Query(`CALL online_rate_by_hour_aproj(?,?,?)`,num ,date,projectid)
	if err2 != nil {
		fmt.Println("查询出现错误!")
		log.Fatal(err2)
	}
	defer rows.Close()
	var TimePoint string//时间点
	var number int//这个时间点在线的设备的数量
	for rows.Next(){

		err3 := rows.Scan(&TimePoint,&number)

		if err3 != nil {
			fmt.Println("rows.Scan出现错误!")
			log.Fatal(err3)
		}
		fmt.Println("number",number)
		fmt.Println("devNum",devNum)
		onlineRate := float64(number)/float64(devNum)
		fmt.Println("onlineRate",onlineRate)
		rateData = append(rateData,datastruct.RateData{TimePoint,onlineRate} )
	}
	err4 := rows.Err()
	if err4 != nil {
		fmt.Println("遍历出现错误!")
		log.Fatal(err4)
	}

	return rateData
}

//更新一下所有设备最后一次上线的时长
func UpdateLastOnlineDuration()  {

}

//从数据库获取单个产品
//粒度是天
func ProdOnlineRateDataFromDbByDay(productid int,from string ,day_inner int) []datastruct.RateData {
	//单个产品有多少个设备
	devNum :=GetAllDevInProj(productid)
	rateData:= make([]datastruct.RateData,0)
	if devNum == 0{
		return rateData
	}
	db , err := conn.ConectDb5()
	defer db.Close()
	if err != nil {
		fmt.Println("连接数据库出现错误!")
	}

	rows, err2 := db.Query(`CALL online_rate_by_day_aprod(?,?,?)`,from,day_inner,productid)

	if err2 != nil {
		fmt.Println("查询出现错误!")
		log.Fatal(err2)
	}
	defer rows.Close()

	var TimePoint string//时间点
	var number int//在线的设备
	for rows.Next(){
		err3 := rows.Scan(&TimePoint,&number)
		if err3 != nil {
			fmt.Println("rows.Scan出现错误!")
			log.Fatal(err3)
		}
		rate := float64(number)/float64(devNum)
		//要某一天的数据
		rateData=append(rateData,datastruct.RateData{TimePoint,rate})

	}
	err4 := rows.Err()
	if err4 != nil {
		fmt.Println("遍历出现错误!")
		log.Fatal(err4)
	}
	return rateData
}
//从数据库获取单个产品
//粒度是小时
func ProdOnlineRateDataFromDbByHour(productid int,date string) []datastruct.RateData{
	//某个项目有多少个设备,
	devNum :=GetAllDevInProj(productid)
	rateData:= make([]datastruct.RateData,0)
	if devNum == 0{
		return rateData
	}
	db , err := conn.ConectDb5()
	defer db.Close()
	if err != nil {
		fmt.Println("连接数据库出现错误!")
	}
	var num int
	if date ==time.Now().Format("2006-01-02"){
		num = time.Now().Hour()
	} else{
		num = 24
	}
	fmt.Println("num",num)
	fmt.Println("date",date)
	rows, err2 := db.Query(`CALL online_rate_by_hour_aprod(?,?,?)`,num ,date,productid)

	if err2 != nil {
		fmt.Println("查询出现错误!")
		log.Fatal(err2)
	}
	defer rows.Close()
	var TimePoint string//时间点
	var number int//上线的时长
	for rows.Next(){

		err3 := rows.Scan(&TimePoint,&number)

		if err3 != nil {
			fmt.Println("rows.Scan出现错误!")
			log.Fatal(err3)
		}
		fmt.Println("number",number)
		fmt.Println("devNum",devNum)
		onlineRate := float64(number)/float64(devNum)
		fmt.Println("onlineRate",onlineRate)
		rateData = append(rateData,datastruct.RateData{TimePoint,onlineRate} )
	}
	err4 := rows.Err()
	if err4 != nil {
		fmt.Println("遍历出现错误!")
		log.Fatal(err4)
	}

	return rateData
}



//从数据库获取单个设备
func DeviceOnlineRateDataFromDb(deviceid string,from string , to string) (t int64) {
	fmt.Println("from",from)
	fmt.Println("to",to)
	db , err := conn.ConectDb5()
	defer db.Close()
	if err != nil {
		fmt.Println("连接数据库出现错误!")
	}
	var first_time string
	//如果to的时间比第一条记录的时间小,则它的在线率为0
	rows0, err0 := db.Query(`SELECT pr.operation_time
                            FROM iot_device_process_record pr
                            WHERE pr.device_id = ?
                            AND pr.state = 1
                             ORDER BY pr.id 
                             LIMIT 1`,deviceid)

	if err0 != nil {
		fmt.Println("rows0查询出现错误!")
		log.Fatal(err)
	}
	for rows0.Next(){
		err3 := rows0.Scan(&first_time)
		if err3 != nil {
			fmt.Println("rows0.Scan出现错误!")
			log.Fatal(err3)
		}

	}
	err = rows0.Err()
	if err != nil {
		fmt.Println("rows0遍历出现错误!")
		log.Fatal(err)
	}
	time_to,_:=time.ParseInLocation("2006-01-02 15:04:05",to,time.Local)
	time_first,_:=time.ParseInLocation("2006-01-02 15:04:05",first_time,time.Local)

	if time_to.Unix()<time_first.Unix(){
		return 0
	}


	////如果from的时间比最后一条记录的时间大,如果最后一条记录的state为不在线,则它的在线率为0,如果最后一条记录的state为在线,则它的在线率为100%
	//var last_time string
	//var last_state int
	//rows_last, err_last := db.Query(`SELECT pr.operation_time,pr.state
     //                       FROM iot_device_process_record pr
     //                       WHERE pr.device_id = ?
     //                       AND pr.state = 1
     //                        ORDER BY pr.id DESC
     //                        LIMIT 1`,deviceid)
	//if err_last != nil {
	//	fmt.Println("rows_last查询出现错误!")
	//	log.Fatal(err)
	//}
	//for rows_last.Next(){
	//	err3 := rows_last.Scan(&last_time,&last_state)
	//	if err3 != nil {
	//		fmt.Println("rows0.Scan出现错误!")
	//		log.Fatal(err3)
	//	}
	//
	//}
	//err = rows_last.Err()
	//if err != nil {
	//	fmt.Println("rows0遍历出现错误!")
	//	log.Fatal(err)
	//}
	//time_from,_:=time.ParseInLocation("2006-01-02 15:04:05",from,time.Local)
	//time_last,_:=time.ParseInLocation("2006-01-02 15:04:05",last_time,time.Local)
	//fmt.Println("time_lasttime_lasttime_lasttime_lasttime_last",time_last)
	//if time_from.Unix()>time_last.Unix(){
	//	if last_state == 1{
	//		return 24*3600
	//	}else{
	//		return 0
	//	}
	//}


	var from_later string
	var from_later_duration int64
	var from_later_state int64
	//查找from落在哪个operation_time
	rows, err := db.Query(`SELECT pr.operation_time,pr.duration ,pr.state
                            from iot_device_process_record pr
                            WHERE pr.device_id = ?
                            AND UNIX_TIMESTAMP(pr.operation_time)<= UNIX_TIMESTAMP(?)
                            and UNIX_TIMESTAMP(?)< UNIX_TIMESTAMP(pr.operation_time)+pr.duration`,deviceid,from,from)
	if err != nil {
		fmt.Println("rows查询出现错误!")
		log.Fatal(err)
	}
	for rows.Next(){
		err := rows.Scan(&from_later,&from_later_duration,&from_later_state)
		if err != nil {
			fmt.Println("rows.Scan出现错误!")
			log.Fatal(err)
		}

	}
	err = rows.Err()
	if err != nil {
		fmt.Println("遍历出现错误!")
		log.Fatal(err)
	}
	if  from_later == ""{
		fmt.Println("from_later进来了")
		rows11, err11 := db.Query(`SELECT pr.operation_time, pr.duration ,pr.state
                            FROM iot_device_process_record pr
                            WHERE pr.device_id = ?
                            AND pr.state = 1
                             ORDER BY pr.id DESC 
                             LIMIT 1`,deviceid)

		if err11 != nil {
			fmt.Println("rows11查询出现错误!")
			log.Fatal(err)
		}
		for rows11.Next(){
			err3 := rows11.Scan(&from_later,&from_later_duration,&from_later_state)
			if err3 != nil {
				fmt.Println("rows11.Scan出现错误!")
				log.Fatal(err3)
			}

		}
		err = rows11.Err()
		if err != nil {
			fmt.Println("rows11遍历出现错误!")
			log.Fatal(err)
		}

	}
	fmt.Println("from_later",from_later)
	fmt.Println("from_later_duration",from_later_duration)
	fmt.Println("from_later_state",from_later_state)


	//查找to落在哪个operation_time
	var to_later string
	var to_later_duration int64
	var to_later_state int64
	rows2, err := db.Query(`SELECT pr.operation_time,pr.duration ,pr.state
                            from iot_device_process_record pr
                            WHERE pr.device_id = ?
                            AND UNIX_TIMESTAMP(pr.operation_time)<= UNIX_TIMESTAMP(?)
                            AND UNIX_TIMESTAMP(?)<UNIX_TIMESTAMP(pr.operation_time)+pr.duration`,deviceid,to,to)
	if err != nil {
		fmt.Println("rows2查询出现错误!")
		log.Fatal(err)
	}
	for rows2.Next(){
		err3 := rows2.Scan(&to_later,&to_later_duration,&to_later_state)
		if err3 != nil {
			fmt.Println("rows2.Scan出现错误!")
			log.Fatal(err3)
		}

	}
	err = rows2.Err()
	if err != nil {
		fmt.Println("rows2遍历出现错误!")
		log.Fatal(err)
	}

	if  to_later == ""{
		fmt.Println("to_later进来了")
		rows22, err22 := db.Query(`SELECT pr.operation_time,pr.duration ,pr.state
                            FROM iot_device_process_record pr
                            WHERE pr.device_id = ?
                            AND pr.state = 1
                             ORDER BY pr.id DESC 
                             LIMIT 1`,deviceid)

		if err22 != nil {
			fmt.Println("rows3查询出现错误!")
			log.Fatal(err)
		}
		for rows22.Next(){
			err3 := rows22.Scan(&to_later,&to_later_duration,&to_later_state)
			if err3 != nil {
				fmt.Println("rows2.Scan出现错误!")
				log.Fatal(err3)
			}

		}
		err = rows22.Err()
		if err != nil {
			fmt.Println("rows2遍历出现错误!")
			log.Fatal(err)
		}

	}
	fmt.Println("to_later",to_later)
	fmt.Println("to_later_state",to_later_state)
	fmt.Println("to_later_duration",to_later_duration)
	var rows3 *sql.Rows

     var duration int64
	//查找在这个区间的所有数据
	rows3, err = db.Query(`SELECT pr.duration
                            FROM iot_device_process_record pr
                            WHERE pr.device_id = ?
                            AND UNIX_TIMESTAMP(?)<= UNIX_TIMESTAMP(pr.operation_time)
                            AND UNIX_TIMESTAMP(pr.operation_time)<=UNIX_TIMESTAMP(?)
                            AND pr.state = 1`,deviceid,from_later,to_later)

	if err != nil {
		fmt.Println("rows3查询出现错误!")
		log.Fatal(err)
	}

	//1.第一个数据是上线,最后一个数据是上线
	if from_later_state == 1 && to_later_state == 1{
		fmt.Println("情况1")
        t1:=common.DiffSecond(from_later,from)
		fmt.Println("t1",t1)
		fmt.Println("from_later",from_later)
		fmt.Println("from",from)
		var t2 int64
		for rows3.Next() {
			err = rows3.Scan(&duration)
			t2 += duration
			if err != nil {
				fmt.Println("rows3.1.Scan出现错误!")
				log.Fatal(err)
			}
		}
		fmt.Println("t2",t2)
		t3:=common.DiffSecond(to_later,to)
		fmt.Println("t3",t3)
		fmt.Println("情况1duration",duration)
		t= t2-t1-duration+t3
	}
	//2.第一个数据是下线,最后一个数据是上线
	if from_later_state != 1 && to_later_state == 1{
		fmt.Println("情况2")
		var t2 int64
		for rows3.Next() {
			err = rows3.Scan(&duration)
			t2 += duration
			if err != nil {
				fmt.Println("rows3.2.Scan出现错误!")
				log.Fatal(err)
			}
		}
		t3:=common.DiffSecond(to_later,to)
		t= t2-duration+t3
	}
	//3.第一个数据是上线,最后一个数据是下线
	if from_later_state == 1 && to_later_state != 1{
		fmt.Println("情况3")
		fmt.Println("from",from)
		fmt.Println("from_later",from_later)
		t1:=common.DiffSecond(from_later,from)
		var t2 int64
		for rows3.Next() {
			err = rows3.Scan(&duration)
			t2 += duration
			if err != nil {
				fmt.Println("rows3.3.Scan出现错误!")
				log.Fatal(err)
			}
		}
		fmt.Println("duration",duration)
		fmt.Println(t2)
		fmt.Println(t1)
		fmt.Println("t",t2-t1)
		t= t2-t1
	}
	//4.第一个数据是下线,最后一个数据是下线
	if from_later_state != 1 && to_later_state != 1{
		fmt.Println("情况4")
		var t2 int64
		for rows3.Next() {
			err = rows3.Scan(&duration)
			t2 += duration
			if err != nil {
				fmt.Println("rows3.4.Scan出现错误!")
				log.Fatal(err)
			}
		}
		t= t2
	}
	err = rows3.Err()
	if err != nil {
		fmt.Println("rows3遍历出现错误!")
		log.Fatal(err)
	}
	fmt.Println("t",t)
   return
}

//获取所有项目下的所有设备 projectid =-1
//获取具体的项目下有多少设备则接收具体的项目
func GetAllDevInProj(projectid int) int{
	db , err := conn.ConectDb5()
	defer db.Close()
	if err != nil {
		fmt.Println("连接数据库出现错误!")
	}
	var rows *sql.Rows
	var err2 error
	if projectid == -1{
		rows, err2 = db.Query(`SELECT count(*) as num
                            FROM prj_region_device prd
							LEFT JOIN prj_region pr ON pr.id = prd.region_id`)
	}else {
		fmt.Println("dddfsdfds")
		rows, err2 = db.Query(`SELECT count(*) as num
                            FROM prj_region_device prd
							LEFT JOIN prj_region pr ON pr.id = prd.region_id
							LEFT JOIN prj_project pp ON pp.id = pr.project_id
							 WHERE pp.id = ?
							`,projectid)
	}

	if err2 != nil {
		fmt.Println("查询出现错误!")
		log.Fatal(err2)
	}
	defer rows.Close()
	var num int
	for rows.Next(){
		fmt.Println(22)
		err3 := rows.Scan(&num)
		if err3 != nil {
			fmt.Println("rows.Scan出现错误!")
			log.Fatal(err3)
		}
	}
	fmt.Println("dss",num)
	return num
}

//获取某个产品下的设备的总数
func GetAllDevInProd(productid int) (num int){
	db , err := conn.ConectDb5()
	defer db.Close()
	if err != nil {
		fmt.Println("连接数据库出现错误!")
	}

	rows, err2 := db.Query(`SELECT count(*) as num
                            FROM iot_device  d
							INNER JOIN prj_region_device rd ON rd.device_id = d.id
							INNER JOIN prj_region r ON r.id = rd.region_id
							INNER JOIN prj_project p ON p.id = r.project_id
							WHERE d.product_id = ?
							`,productid)

	if err2 != nil {
		fmt.Println("查询出现错误!")
		log.Fatal(err2)
	}
	defer rows.Close()
	for rows.Next(){
		err3 := rows.Scan(&num)
		if err3 != nil {
			fmt.Println("rows.Scan出现错误!")
			log.Fatal(err3)
		}
	}
	return
}



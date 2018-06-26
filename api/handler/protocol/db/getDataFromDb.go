package db

import (
	"fmt"
	"wuyuyun-datastore-v0.2/api/handler/protocol/datastruct"
	"wuyuyun-datastore-v0.2/api/common"
	"log"
	"time"
	"strconv"
	"github.com/influxdata/influxdb/client/v2"
	"wuyuyun-datastore-v0.2/conn"
	"wuyuyun-datastore-v0.2/config"
)
/************************** mysql **************************/

//从数据库中查到今天的数据
func ListTodayDataFromDb(devid string,protocolid int ) []datastruct.Hours{
	todayData:=make([]datastruct.Hours,0,5)
	db , err := conn.ConectDb5()
	defer db.Close()
	if err != nil {
		fmt.Println("连接数据库出现错误!")
	}
	rows, err2 := db.Query(`SELECT 
							FROM_UNIXTIME(insert_time/1000000000,'%H'), 
							MAX(value) as valuemax,
							MIN(value) as valuemin,
							AVG(value) as valueavg
							FROM iot_device_data
							WHERE device_id = ?
                         AND proto_id =?
							AND FROM_UNIXTIME(insert_time/1000000000,'%Y-%m-%d') = DATE_FORMAT(now(),'%Y-%m-%d')
						    AND UNIX_TIMESTAMP(FROM_UNIXTIME(insert_time/1000000000,'%Y-%m-%d %H')) <= UNIX_TIMESTAMP(DATE_FORMAT(date_sub(now(), interval 1 hour),'%Y-%m-%d %H'))  
							GROUP BY FROM_UNIXTIME(insert_time/1000000000,'%H')`,devid,protocolid)
	if err2 != nil {
		fmt.Println("查询出现错误!")
		log.Fatal(err2)
	}
	defer rows.Close()

	var hour int//时间点
	var valuemax float64//采集的平均值
	var valuemin float64//采集的平均值
	var valueavg float64//采集的平均值
	for rows.Next(){
		err3 := rows.Scan(&hour,&valuemax,&valuemin,&valueavg)
		if err3 != nil {
			fmt.Println("rows.Scan出现错误!")
			log.Fatal(err3)
		}
		//要今天的数据
		todayData=append(todayData,datastruct.Hours{hour,valuemax,valuemin,valueavg})

	}
	err4 := rows.Err()
	if err4 != nil {
		fmt.Println("遍历出现错误!")
		log.Fatal(err4)
	}
	return todayData
}
//从数据库中查某一天的数据
func ListOneDataFromDb(devid string,protocolid int,date string) []datastruct.Hours{
	fmt.Println("dddddddddddddddddd")
	onedayData:=make([]datastruct.Hours,0,5)
	db , err := conn.ConectDb5()
	defer db.Close()
	if err != nil {
		fmt.Println("连接数据库出现错误!")
	}
	rows, err2 := db.Query(`SELECT 
							FROM_UNIXTIME(insert_time/1000000000,'%H'), 
							MAX(value) as valuemax,
							MIN(value) as valuemin,
							AVG(value) as valueavg
							FROM iot_device_data
							WHERE device_id = ?
                         AND proto_id =?
							AND FROM_UNIXTIME(insert_time/1000000000,'%Y-%m-%d') = ?
							GROUP BY FROM_UNIXTIME(insert_time/1000000000,'%H')`,devid,protocolid,date)
	if err2 != nil {
		fmt.Println("查询出现错误!")
		log.Fatal(err2)
	}
	defer rows.Close()

	var hour int//时间点
	var valuemax float64//采集的平均值
	var valuemin float64//采集的平均值
	var valueavg float64//采集的平均值
	for rows.Next(){

		err3 := rows.Scan(&hour,&valuemax,&valuemin,&valueavg)
		if err3 != nil {
			fmt.Println("rows.Scan出现错误!")
			log.Fatal(err3)
		}
		//要某一天的数据
		onedayData=append(onedayData,datastruct.Hours{hour,valuemax,valuemin,valueavg})

	}
	err4 := rows.Err()
	if err4 != nil {
		fmt.Println("遍历出现错误!")
		log.Fatal(err4)
	}
	return onedayData
}
//从数据库查到给定日期中的七天的数据
func ListSevaldaysAgoFromDb(devid string,protocolid int ,from string , to string) datastruct.Predays{
	//原始的
	initdays:=make([]datastruct.DateAgr,0,5)
	//预重构
	var predays datastruct.Predays
	predays.From = from
	predays.To = to
	//把string的from解析成time
	t1,_:= time.Parse("2006-01-02",from)
	db , err := conn.ConectDb5()
	day_inner :=common.DiffDay(from,to)//计算两个日期相差多少天
	//初始化这样一个数组[{2018-05-01 00 0 0 0},{2018-05-01 01 0 0 0},...]
	for i:=0;i<day_inner;i++{
		t3:=t1.AddDate(0,0,i)
		for j:=0;j<24;j++{
			s := strconv.Itoa(j)
			h, _ := time.ParseDuration(s+"h")
			date:=t3.Add(h).Format("2006-01-02 15")
			initdays=append(initdays,datastruct.DateAgr{date,0,0,0} )
		}
	}

	defer db.Close()
	if err != nil {
		fmt.Println("连接数据库出现错误!")
		log.Fatal(err)
	}
	rows, err1 := db.Query(`SELECT FROM_UNIXTIME(insert_time/1000000000,'%Y-%m-%d %H') as time,
                      MAX(value) as valuemax,
						 MIN(value) as valuemin,
						 AVG(value) as valueavg
						 FROM iot_device_data
						 WHERE device_id = ?
                      AND proto_id = ? 
						 AND UNIX_TIMESTAMP(?)<=UNIX_TIMESTAMP(FROM_UNIXTIME(insert_time/1000000000,'%Y-%m-%d'))
						 AND UNIX_TIMESTAMP(FROM_UNIXTIME(insert_time/1000000000,'%Y-%m-%d'))<=UNIX_TIMESTAMP(?) 
						 GROUP BY FROM_UNIXTIME(insert_time/1000000000,'%Y-%m-%d %H')`,devid,protocolid,from,to)
	if err1 != nil {
		fmt.Println("查询表出现错误!")
		log.Fatal(err1)
	}
	defer rows.Close()
	var date string//日期 %Y-%m-%d %H
	var valuemax float64//采集的平均值
	var valuemin float64//采集的平均值
	var valueavg float64//采集的平均值
	//循环initdays数组
	var i int
	if ok:=rows.Next();ok{
		err2 := rows.Scan(&date,&valuemax,&valuemin,&valueavg)
		if err2 != nil {
			fmt.Println("Scan出现错误!")
			log.Fatal(err2)
		}
		predays.Min = valuemin
		predays.Max = valuemax
		avg := valueavg
		i=1
		for z:=0;z<len(initdays);z++ {
			if initdays[z].Date == date {
				initdays[z] = datastruct.DateAgr{date,valueavg,valuemax,valuemin}
				if ok:=rows.Next();ok {
					err2 := rows.Scan(&date, &valuemax, &valuemin, &valueavg)
					if err2 != nil {
						fmt.Println("Scan出现错误!")
						log.Fatal(err2)
					}
					if predays.Min > valuemin{
						predays.Min = valuemin
					}
					if predays.Max < valuemax{
						predays.Max= valuemax
					}
					i+=1
					avg += valueavg
				  }
				}
		}
		predays.Avg = avg/float64(i)
		predays.Dates = initdays
	}

	err3:= rows.Err()
	if err3 != nil {
		fmt.Println("rows.Err出现错误!")
		log.Fatal(err3)
	}
	 return predays
}

//从数据库中查找数据,格式化到月份
func GetDataFromDbBySeconds (devid string ,protocolid int , from string, to string) datastruct.Predays{
	year_to ,_:= strconv.Atoi(to[0:4])
	month_to ,_:= strconv.Atoi(to[5:7])
	day_int:=common.CountDaysInEveryMonth(year_to,month_to)
    from_day:=from + "-01"
	to_day:=to +"-"+ strconv.Itoa(day_int)
	//原始的
	initdays:=make([]datastruct.DateAgr,0,5)
	//预重构
	var predays datastruct.Predays
	predays.From = from
	predays.To = to
	//把string的from解析成time
	t1,_:= time.Parse("2006-01-02",from_day)
	day_inner :=common.DiffDay(from_day,to_day)//计算两个日期相差多少天
	//初始化这样一个数组[{2018-05-01 00-00 0 0 0},{2018-05-01 01-00 0 0 0},...]
	//for i:=0;i<day_inner;i++{
	//	t3:=t1.AddDate(0,0,i)
	//	for j:=0;j<24;j++{
	//		s := strconv.Itoa(j)
	//		h, _ := time.ParseDuration(s+"h")
	//		date_h:=t3.Add(h)
	//		for z:=0;z<60;z++{
	//			s1 := strconv.Itoa(z)
	//			m, _ := time.ParseDuration(s1+"m")
	//			date_m:=date_h.Add(m)
	//			date:=date_m.Format("2006-01-02 15:04")
	//			initdays=append(initdays,datastruct.DateAgr{date,0,0,0} )
	//		}
	//	}
	//}
	for i:=0;i<day_inner;i++{
		t3:=t1.AddDate(0,0,i)
		date:=t3.Format("2006-01-02")
		initdays=append(initdays,datastruct.DateAgr{date,0,0,0} )
	}
	db , err := conn.ConectDb5()
	defer db.Close()
	if err != nil {
		fmt.Println("连接数据库出现错误!")
		log.Fatal(err)
	}
	rows, err1 := db.Query(`SELECT FROM_UNIXTIME(insert_time/1000000000,'%Y-%m-%d') as time,
                         MAX(value) as valuemax,
						    MIN(value) as valuemin,
						    AVG(value) as valueavg
						    FROM iot_device_data
						    WHERE device_id = ?
                          AND proto_id = ? 
						    AND UNIX_TIMESTAMP(?)<=UNIX_TIMESTAMP(FROM_UNIXTIME(insert_time/1000000000,'%Y-%m-%d'))
						    AND UNIX_TIMESTAMP(FROM_UNIXTIME(insert_time/1000000000,'%Y-%m-%d'))<=UNIX_TIMESTAMP(?) 
						    GROUP BY FROM_UNIXTIME(insert_time/1000000000,'%Y-%m-%d')`,devid,protocolid,from_day,to_day)
	if err1 != nil {
		fmt.Println("查询表出现错误!")
		log.Fatal(err1)
	}
	defer rows.Close()
	var date string//日期 %Y-%m-%d %H-%i-%s
	var valuemax float64//采集的平均值
	var valuemin float64//采集的平均值
	var valueavg float64//采集的平均值
	//循环initdays数组
	var i int
	if ok:=rows.Next();ok{
		err2 := rows.Scan(&date,&valuemax,&valuemin,&valueavg)
		if err2 != nil {
			fmt.Println("Scan出现错误!")
			log.Fatal(err2)
		}
		predays.Min = valuemin
		predays.Max = valuemax
		avg := valueavg
		i=1
		for z:=0;z<len(initdays);z++ {
			if initdays[z].Date == date {
				initdays[z] = datastruct.DateAgr{date,valueavg,valuemax,valuemin}
				if ok:=rows.Next();ok {
					err2 := rows.Scan(&date, &valuemax, &valuemin, &valueavg)
					if err2 != nil {
						fmt.Println("Scan出现错误!")
						log.Fatal(err2)
					}
					if predays.Min > valuemin{
						predays.Min = valuemin
					}
					if predays.Max < valuemax{
						predays.Max= valuemax
					}
					i+=1
					avg += valueavg
				}
			}
		}
		predays.Avg = avg/float64(i)
		predays.Dates = initdays
	}

	//for rows.Next(){
	//	err2 := rows.Scan(&date,&valuemax,&valuemin,&valueavg)
	//	if err2 != nil {
	//		fmt.Println("Scan出现错误!")
	//		log.Fatal(err2)
	//	}
	//	sevendays=append(sevendays, datastruct.DateAgr{date,valuemax,valuemin,valueavg})
	//}
	err3:= rows.Err()
	if err3 != nil {
		fmt.Println("rows.Err出现错误!")
		log.Fatal(err3)
	}
	return predays
}


/************************** influxdb **************************/
// 获取今天的数据,粒度是小时
func GetDataFromInfluxdbByHour(from string,to string ,devid string,protocolid string )[]datastruct.Datas{
	fmt.Println("dddddddd",from)
	fmt.Println("dddddddd",to)
	db , err := conn.ConnInfluxdb_xiayan()
	defer db.Close()
	if err != nil {
		fmt.Println("连接数据库出现错误!")
	}
	qs := fmt.Sprintf("select  mean(value),median(value), max(value), min(value)  from iot_device_data  where time >= '%s' and time < '%s'  and  device_id = '%s' and  proto_id= '%s'  group by time(1h) fill(0)", from,to,devid,protocolid )
	res, err := QueryDB(db,qs )
	if err != nil {
		log.Fatal(err)
	}
	datas := make([]datastruct.Datas, 0)

    if res[0].Series != nil {

		for i := 0; i < len(res[0].Series[0].Values); i++ {
			datas = append(datas, datastruct.Datas{res[0].Series[0].Values[i][0].(string), res[0].Series[0].Values[i][1], res[0].Series[0].Values[i][2], res[0].Series[0].Values[i][3], res[0].Series[0].Values[i][4]})
		}
		return datas
	}
	//fmt.Println("获取今天的数据,粒度是小时",res)

	//}
	return datas


}
// 获取今天的数据,粒度是天
func GetDataFromInfluxdbByDay(from string,to string ,devid string,protocolid string ) []datastruct.Datas{
	db , err := conn.ConnInfluxdb_xiayan()
	fmt.Println("dddddddd",from)
	fmt.Println("dddddddd",to)
	defer db.Close()
	if err != nil {
		fmt.Println("连接数据库出现错误!")
	}
	//t1:=time.Now().Format("2006-01-02 15:04:05")
	qs := fmt.Sprintf("select  mean(value),median(value), max(value), min(value)  from iot_device_data  where time >= '%s' and time < '%s'  and  device_id = '%s' and  proto_id= '%s'  group by time(8h) fill(0)", from,to,devid,protocolid )
	//qs := fmt.Sprintf("select * from iot_device_data" )
	fmt.Println("qs数据",qs)
	res, err := QueryDB(db,qs )
	if err != nil {
		log.Fatal(err)
	}
	datas := make([]datastruct.Datas, 0)
	if res[0].Series != nil {

		for i := 0; i < len(res[0].Series[0].Values); i++ {
			datas = append(datas, datastruct.Datas{res[0].Series[0].Values[i][0].(string), res[0].Series[0].Values[i][1], res[0].Series[0].Values[i][2], res[0].Series[0].Values[i][3], res[0].Series[0].Values[i][4]})
		}
		fmt.Println("datas数据",datas)

		return datas
	}
	//t2:=time.Now().Format("2006-01-02 15:04:05")
	//t:=common.DiffSecond(t1,t2)
	//fmt.Println("查询所有的数据需要多少时间",t)
	//fmt.Println("获取今天的数据,粒度是小时",res)
	return datas


}

//query
func QueryDB(cli client.Client, cmd string) (res []client.Result, err error) {
	q := client.Query{
		Command:  cmd,
		Database:config.MyDB_official ,
	}
	if response, err := cli.Query(q); err == nil {
		if response.Error() != nil {
			return res, response.Error()
		}
		res = response.Results
	} else {
		return res, err
	}
	return res, nil
}

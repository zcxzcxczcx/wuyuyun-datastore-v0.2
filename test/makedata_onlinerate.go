package test

import (
	"time"
	"fmt"
	"strconv"
	"wuyuyun-datastore-v0.2/conn"
)

//往iot_device_process_record 表造数据
func MakeData(){
	//连接数据库
	db,err := conn.ConectDb4()
	defer db.Close()
	if err!=nil{
		fmt.Println("连接数据库发生错误")
		return
	}
    start_time_0:=time.Now().AddDate(0,0,-30).Format("2006-01-02")
	fmt.Println("start_time_0",start_time_0)

	start_time ,_:=	time.ParseInLocation("2006-01-02",start_time_0,time.Local)
	fmt.Println("start_time",start_time)
	var b int
	for i:=0;i<29;i++{
		start_time_1:=start_time.AddDate(0,0,i)
		var len int
		if start_time_1.Format("2006-01-02") == time.Now().Format("2006-01-02"){
			len= time.Now().Hour()
		}else{
			len =24
		}
    	for j:=0;j<len;j++{

    		if j%2 ==0{
				b =1
			}else{
				b =0
			}
			j_string:=strconv.Itoa(j)
			add,_:=time.ParseDuration(j_string+"h")
			fmt.Println("add",add)
			start_time_2:=start_time_1.Add(add).Format("2006-01-02 15:04:05")
			fmt.Println("start_time_2",start_time_2)
			db.Exec("INSERT INTO iot_device_process_record (operation_time,device_id,state,duration ) values (?,?,?,?)",start_time_2,"68F0E2D73C9E0FF14D3C14FA242A806C",b,3600)
		}
	}
}

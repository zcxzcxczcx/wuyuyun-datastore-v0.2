package sub

import (
	"database/sql"
	"fmt"
	"encoding/json"
	"time"
	"wuyuyun-datastore-v0.2/mqtt/structs"
	MQTT "github.com/eclipse/paho.mqtt.golang"//eclipse组织
	"wuyuyun-datastore-v0.2/api/common"
	"github.com/influxdata/influxdb/client/v2"
	"wuyuyun-datastore-v0.2/config"
	"strconv"
	"log"
	"wuyuyun-datastore-v0.2/conn"
	"math/rand"

)


/************************** 订阅网关传感器的数据 存入mysql数据库 **************************/

func MqttSubStoreInMysql(topicData string,topicStatus string,c MQTT.Client)  {
	//连接mysql数据库
	dbmysql,err := conn.ConectDb5()
	defer dbmysql.Close()

	//处理器"+/data"
	var subDataToMysql MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
		//client是我自己的
        fmt.Println("这是主题'+/data'的处理器")
		fmt.Printf("TOPIC: %s\n", msg.Topic())
		fmt.Printf("没有Unmarshal的Message: %s\n", msg.Payload())

		//创建并初始化一个管道
		out := make(chan int)

		//从 msg.Topic()中拿到devid
		topic:=fmt.Sprint(msg.Topic())
		devid:=GetProidFromTopid(topic)
		fmt.Println(devid)


		if err!=nil{
			fmt.Println("连接数据库发生错误")
			return
		}

		var s []structs.Productdata
		json.Unmarshal(msg.Payload(),&s)
		fmt.Printf("Unmarshal的Message: %v\n", s)
		//
		////1：先存入缓存,所以在缓存中没有存入数据库的时间CreateTime
		//redisresult:=cache.WriteToRedis(proid,s)
		//fmt.Println(redisresult)
		////测试有没有存入缓存
		//productdataGet:=cache.GetFromRedis(proid)
		//fmt.Println(productdataGet)

		////往管道里面存传感器的数据
	     //StoreToChan(out,s)
		//2：把拿到的数据并发的存入设备数据表里
		for i:=0; i<len(s);i++{
			go StoreInDb(s[i],dbmysql,devid,out)
		}
		for j:=0; j<len(s);j++{
		    <-out
		}
		close(out)

	}
	//处理器"+/status"
	var subStatus MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
		fmt.Println("这是主题'+/status'的处理器")
		fmt.Printf("TOPIC: %s\n", msg.Topic())
		fmt.Printf("没有Unmarshal的Message: %s\n", msg.Payload())
		//从TOPIC拿到设备ID
		deviceID:=GetDeviceIdFromTopic(msg.Topic())
		fmt.Println("拿到的设备ID",deviceID)
		//拿到payload为status_payload
		var status_payload structs.StatusPayload
		json.Unmarshal(msg.Payload(),&status_payload)
		fmt.Println("status的数据:",status_payload )
		//status_payload.Type=1:上线;status_payload.Type=2:正常下线;status_payload.Type=3:异常下线;
		//state=0正常下线，state=1正上线，state=2异常下线常
		if status_payload.Type == 1{
			fmt.Println("设备上线了")
			//有这个设备，则更新设备表最近上线时间，状态为在线1。
			stmt1, err1 := dbmysql.Prepare(`UPDATE iot_device SET last_online=?,online=? WHERE id=?`)

			//查找这个设备最近一条记录，获取最近一条记录的下线时间
			var id int
			var operation_time string
			rows3, err3 :=dbmysql.Query("select id,operation_time from iot_device_process_record  where device_id=? order by id desc limit 1",deviceID);
			defer rows3.Close()
			if err3 != nil {
				fmt.Printf("查找这个设备最近一条记录: %v\n", err3)
				return
			}
			for rows3.Next() {
				rows3.Scan(&id, &operation_time)
			}

			err33 := rows3.Err()
			if err33 != nil {
				fmt.Printf("遍历这个设备最近一条记录: %v\n", err33)
			}
			//更新这条记录的时长
			tdiff :=common.DiffSecond(operation_time,time.Now().Format("2006-01-02 15:04:05"))
			fmt.Println("tdiff",tdiff)
			dbmysql.Exec("UPdate iot_device_process_record SET duration =? where id=? ",tdiff,id)

			//往设备运行记录表里插入一条上线记录
			stmt2, err2:= dbmysql.Prepare(`INSERT INTO iot_device_process_record (operation_time,device_id,state,duration ) values (?,?,?,?)`)
			if err1!=nil{
				fmt.Println("Prepare stmt1 数据不成功")
				fmt.Println(err1)
				return
			}
			if err2!=nil{
				fmt.Println("Prepare stmt2 数据不成功")
				fmt.Println(err2)
				return
			}
			result1, err11 := stmt1.Exec(time.Now().Format("2006-01-02 15:04:05"),1,deviceID )
			//duration填一个很大的数方便统计
			result2, err22 := stmt2.Exec(time.Now().Format("2006-01-02 15:04:05"),deviceID,1,31536000)
			fmt.Println("往iot_device表更新数据结果",result1)
			fmt.Println("往iot_device_process_record表插入数据结果",result2)
			if err11!=nil{
				fmt.Println(err11)
				fmt.Println("往iot_device表插入数据不成功")
				return
			}
			if err22!=nil{
				fmt.Println(err22)
				fmt.Println("往iot_device_process_record表更新数据不成功")
				return
			}
		}else if status_payload.Type == 2{
			//设备正常下线，更新iot_device表里这个设备的下线时间并且状态改为下线0.往iot_device_process_record表插入一条这个设备下线的记录
			fmt.Println("设备正常下线了")
			stmt1, err1 := dbmysql.Prepare(`UPDATE iot_device SET  last_offline=?,online=? WHERE id=?`)

			//查找这个设备最近一条记录，获取最近一条记录的上线时间
			var id int
			var operation_time string
			rows3, err3 :=dbmysql.Query("select id,operation_time from iot_device_process_record  where device_id=? order by id desc limit 1",deviceID);
			defer rows3.Close()
			if err3 != nil {
				fmt.Printf("查找这个设备最近一条记录: %v\n", err3)
				return
			}
			for rows3.Next() {
				rows3.Scan(&id, &operation_time)
			}

			err33 := rows3.Err()
			if err33 != nil {
				fmt.Printf("遍历这个设备最近一条记录: %v\n", err33)
			}
			//更新这条记录的时长
			tdiff :=common.DiffSecond(operation_time,time.Now().Format("2006-01-02 15:04:05"))
			fmt.Println("tdiff",tdiff)
			dbmysql.Exec("UPdate iot_device_process_record SET duration =? where id=? ",tdiff,id)
			stmt2, err2:= dbmysql.Prepare(`INSERT INTO iot_device_process_record ( operation_time,device_id,state) values (?,?,?)`)
			if err1!=nil{
				fmt.Println("Prepare stmt1 数据不成功")
				fmt.Println(err1)
				return
			}
			if err2!=nil{
				fmt.Println("Prepare stmt2 数据不成功")
				fmt.Println(err2)
				return
			}
			result1, err11 := stmt1.Exec(time.Now().Format("2006-01-02 15-04-05"),0,deviceID )
			result2, err22 := stmt2.Exec(time.Now().Format("2006-01-02 15-04-05"),deviceID,0)
			fmt.Println("往iot_device表更新数据结果",result1)
			fmt.Println("往iot_device_process_record表插入数据结果",result2)
			if err11!=nil{
				fmt.Println(err11)
				fmt.Println("往iot_device表插入数据不成功")
				return
			}
			if err22!=nil{
				fmt.Println(err22)
				fmt.Println("往iot_device_process_record表更新数据不成功")
				return
			}
		}else if status_payload.Type == 3{
			//设备异常下线，更新iot_device表里这个设备的下线时间并且状态改为下线0.往iot_device_process_record表插入一条这个设备下线的记录并且标注设备是异常下线
			fmt.Println("设备异常下线")
			stmt1, err1 := dbmysql.Prepare(`UPDATE iot_device SET  last_offline=?,online=? WHERE id=?`)



			//查找这个设备最近一条记录，获取最近一条记录的上线时间
			var id int
			var operation_time string
			rows3, err3 :=dbmysql.Query("select id,operation_time from iot_device_process_record  where device_id=? order by id desc limit 1",deviceID);
			defer rows3.Close()
			if err3 != nil {
				fmt.Printf("查找这个设备最近一条记录: %v\n", err3)
				return
			}
			for rows3.Next() {
				rows3.Scan(&id, &operation_time)
			}

			err33 := rows3.Err()
			if err33 != nil {
				fmt.Printf("遍历这个设备最近一条记录: %v\n", err33)
			}
			//更新这条记录的时长
			tdiff :=common.DiffSecond(operation_time,time.Now().Format("2006-01-02 15:04:05"))
			fmt.Println("tdiff",tdiff)
			dbmysql.Exec("UPdate iot_device_process_record SET duration =? where id=? ",tdiff,id)

			stmt2, err2:= dbmysql.Prepare(`INSERT INTO iot_device_process_record (operation_time,device_id,state) values (?,?,?)`)
			if err1!=nil{
				fmt.Println("Prepare stmt1 数据不成功")
				fmt.Println(err1)
				return
			}
			if err2!=nil{
				fmt.Println("Prepare stmt2 数据不成功")
				fmt.Println(err2)
				return
			}
			result1, err11 := stmt1.Exec(time.Now().Format("2006-01-02 15-04-05"),0,deviceID )
			result2, err22 := stmt2.Exec(time.Now().Format("2006-01-02 15-04-05"),deviceID,3)
			fmt.Println("往iot_device表更新数据结果",result1)
			fmt.Println("往iot_device_process_record表插入数据结果",result2)
			if err11!=nil{
				fmt.Println(err11)
				fmt.Println("往iot_device表插入数据不成功")
				return
			}
			if err22!=nil{
				fmt.Println(err22)
				fmt.Println("往iot_device_process_record表更新数据不成功")
				return
			}
		}
	}
	/************************** sub data 存入mysql **************************/
	tokenDataToMysql := c.Subscribe(topicData, 0, subDataToMysql)
	if tokenDataToMysql.Wait() && tokenDataToMysql.Error() != nil {
		fmt.Printf("%v\n", tokenDataToMysql.Error())
	} else {
		fmt.Printf("SubData success\n")
	}


	/************************** sub status 存入mysql **************************/
	tokenStatus := c.Subscribe(topicStatus, 0, subStatus)
	if tokenStatus.Wait();tokenStatus.Error() != nil {
		fmt.Printf("%v\n", tokenStatus.Error())
	} else {
		fmt.Printf("SubStatus success\n")
	}

	//不能让这个程序退出
	for{
		time.Sleep(1*time.Second)
	}
	//c.Disconnect(250)
}
//往mysql数据库里存入传感器的的数据
func  StoreInDb(chs structs.Productdata , db *sql.DB,devid string,out chan int){
		//fmt.Println(chs)
	sampleTimeLocal :=chs.SampleTime.In(time.Local)
	_, err := db.Exec(
		"INSERT INTO iot_device_data ( device_id,proto_id,insert_time,value,sample_time) values (?,?,?,?,?)",
		devid,
		chs.AgrId,
		time.Now().UnixNano(),
		chs.Value,
		sampleTimeLocal.UnixNano(),
	)
		//stmt, err := db.Prepare(`INSERT wuyuyun0.2.iot_device_data ( device_id,proto_id,insert_time,value,sample_time) values (?,?,?,?,?)`)
		//if err != nil {
		//	fmt.Println("Prepare数据不成功")
		//	return
		//}
		//time.Now()输出默认CST时区时间。在网关里面默认输出UTC时区时间
		//time.Parse()默认输出UTC时区时间。
		//Parse()函数解析的时候，会默为UTC时间，获取的Time对象转换为Unix()对象后，会比当前时间多8小时。
		//把UTC时间转换成本地时间
		//fmt.Println(time.Now().UTC())
		//sampleTimeLocal:=time.Now().UTC().In(time.Local)
		//fmt.Println(sampleTimeLocal)

		//sampleTimeLocal :=chs.SampleTime.In(time.Local)
		//_, err = stmt.Exec(devid, chs.AgrId, time.Now().UnixNano(), chs.Value, sampleTimeLocal.UnixNano())
         //defer stmt.Close() //关闭之
		//获取最后一条插入的id
		//result,err:=stmt.Exec(proid,ch.AgrId,time.Now().Format("2006-01-02 15:04:05"),ch.Value,ch.SampleTime)
		//id,err:= result.LastInsertId()
		//fmt.Println(id)
	if err != nil {
		fmt.Println("插入数据不成功")
		panic(err)
		return
	}
	 out<-1

}
//拿到proid
func GetProidFromTopid(topic string) string{
	s:=Substr(topic,0)
	return s
}
//截取字符串 start 起点下标, length -5:需要截取的长度
func Substr(str string, start int) string {
	rs := []rune(str)
	rl := len(rs)
	return string(rs[start:rl-5])
}
//从TOPIC拿到设备ID的函数
func GetDeviceIdFromTopic(topic string) string{
	index :=len(topic)
	infos :=make([]string,0,10)
	for i:=len(topic)-1; i >=0 ; i--{
		if topic[i] == '/'{
			infos = append(infos,topic[i+1:index] )
			index=i
		}
	}
	infos = append(infos,topic[0:index] )
	return infos[1]
}


/************************** 订阅网关传感器的数据 存入influxdb **************************/

func MqttSubStoreToInfluxdb(c MQTT.Client,topic string){
	httpclient,err := conn.ConnInfluxdbTest()
	if err!=nil{
		fmt.Println("连接influxdb出现错误",err)
	}
	defer httpclient.Close()
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  config.MyDB_official,
		Precision: "s",
	})
	if err != nil {
		log.Fatal(err)
	}
	var sub_data_cache11 [100000]structs.Productdata//缓冲管道1
	var sub_data_cache1_copy [100000]structs.Productdata//备份数组
	var index int
	//处理器
	var callback MQTT.MessageHandler = func(c MQTT.Client, msg MQTT.Message) {
		//client是我自己的
		//fmt.Printf("TOPIC: %s\n", msg.Topic())
		//fmt.Printf("没有Unmarshal的Message: %s\n", msg.Payload())
		//从 msg.Topic()中拿到devid
		devid:=GetProidFromTopid(msg.Topic())

		//fmt.Println(devid)
		var sub_data []structs.Productdata
		json.Unmarshal(msg.Payload(),&sub_data)
		//fmt.Printf("Unmarshal的Message: %v\n", sub_data)
		//tags := map[string]string{
		//}
		//fmt.Println(1)
		for i:=0; i<len(sub_data);i++ {
			//fmt.Println(2)
			sub_data_cache11[index] = sub_data[i]
			if index == 99999 {
				fmt.Println(3)
				sub_data_cache1_copy = sub_data_cache11//备份
				go func() {
					fmt.Println(4)
					writesPoints(httpclient, bp, sub_data_cache1_copy, devid)
				}()
				index =-1
			}
			//fmt.Println(7)
			index++

		}
	}


	token := c.Subscribe(topic , 1, callback)

	if  token.Wait()&&token.Error() != nil {
		fmt.Printf("%v\n", token.Error())
	} else {
		fmt.Printf("Sub success\n")
	}
}


//往influxdb存入传感器的数据
//我暂时认为Write()就是一次http写入
func writesPoints(conn client.Client,bp client.BatchPoints,sub_data_cache1 [100000]structs.Productdata,devid string) {
	for i:=0; i<len(sub_data_cache1);i++{
		//fmt.Println(4)

		tags := map[string]string{
			"proto_id":  strconv.Itoa(int(sub_data_cache1[i].AgrId)),
			"device_id":   devid+strconv.Itoa(rand.Intn(100000000000)),
		}
		fields := map[string]interface{}{
			"value": float64(rand.Intn(100000000000))+float64(rand.Intn(100000000000)),
			"sample_time": sub_data_cache1[i].SampleTime,
		}
		pt, err := client.NewPoint(
			"iot_device_data_test",
			tags,
			fields,
			time.Now(),
		)
		if err != nil{
			log.Fatal(err)
		}
		bp.AddPoint(pt)
	}
	t1:=time.Now()
	if err := conn.Write(bp); err != nil {
		fmt.Println("dddd")
		log.Fatal(err)
	} else{
		fmt.Println("写入influxdb成功")
		//out <- 1
		t2:=time.Now()
		t:=t2.UnixNano()-t1.UnixNano()
		fmt.Println(t)
		fmt.Println("dsddddddddd")

	}
}
package test
import (
	"wuyuyun-datastore-v0.2/conn"
	"fmt"
	"log"
	"strconv"
	"time"
	"github.com/influxdata/influxdb/client/v2"
	"wuyuyun-datastore-v0.2/mqtt/structs"
	"wuyuyun-datastore-v0.2/config"
)
//从夏衍服务器mysql数据库的wuyuyun0.2 的iot_device_data表拿到数据
//存入夏衍服务器influxdb数据库wuyuyun-sub的iot_device_data表
func MakedatToInfluxdab(){
	db , err := conn.ConectDb5()
	defer db.Close()
	if err != nil {
		fmt.Println("连接数据库出现错误!")
	}

	httpc,errhttpc := conn.ConnInfluxdb_xiayan()
	if errhttpc!=nil{
		fmt.Println("连接httpclient出现错误",errhttpc)
	}
	defer httpc.Close()

	rows,err:=db.Query("select device_id,proto_id,value,sample_time from iot_device_data")

	sub_data_cache:=make([]structs.DataToInfluxdb,0)//缓冲管道

	var deviceid string
	var protocoid int32
    var value float64
    var sampletime string
	for rows.Next(){
		err3 := rows.Scan(&deviceid,&protocoid,&value,&sampletime)
		if err3 != nil {
			fmt.Println("rows.Scan出现错误!")
			log.Fatal(err3)
		}
		fmt.Println("2")
		sub_data_cache = append(sub_data_cache, structs.DataToInfluxdb{protocoid,deviceid,value,sampletime})

	}
	err4 := rows.Err()
	if err4 != nil {
		fmt.Println("遍历出现错误!")
		log.Fatal(err4)
	}
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  config.MyDB_official,
		Precision: "s",
	})
	if err != nil {
		log.Fatal(err)
	}
  for i:=0;i<1968;i++{
	  fmt.Println("3")
	  writesPoints2(httpc,bp,sub_data_cache)
  }

}

func writesPoints2(conn client.Client,bp client.BatchPoints,sub_data_cache []structs.DataToInfluxdb) {
	for i:=0; i<1000;i++{
		fmt.Println(4)

		tags := map[string]string{
			"proto_id":  strconv.Itoa(int(sub_data_cache[i].AgrId)),
			"device_id":   sub_data_cache[i].DeviceId,
		}
		fields := map[string]interface{}{
			"value": float64(sub_data_cache[i].Value),
			"sample_time": sub_data_cache[i].SampleTime,
		}
		pt, err := client.NewPoint(
			"iot_device_data",
			tags,
			fields,
			time.Now(),
		)
		if err != nil{
			log.Fatal(err)
		}
		bp.AddPoint(pt)
	}
	if err := conn.Write(bp); err != nil {
		fmt.Println("dddd")
		log.Fatal(err)
	} else{
		fmt.Println("写入influxdb成功")
	}
}
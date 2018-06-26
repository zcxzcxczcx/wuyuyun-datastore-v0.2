package test
import (
	"wuyuyun-datastore-v0.2/conn"
	"fmt"
	"log"
	"time"
	"github.com/influxdata/influxdb/client/v2"
	"wuyuyun-datastore-v0.2/config"
	//"math/rand"
	"strconv"
)
func StoreIntoInfluxdb(n int){
	db,err := conn.ConnInfluxdb_121()
	if err!=nil {
		fmt.Println("连接influxdb出现错误",err)
	}
	defer db.Close()


	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  config.MyDB_official,
		Precision: "s",
	})
	if err != nil {
		log.Fatal(err)
	}


	for i:=0;i<n;i++{
		fmt.Println(i)
		tags := map[string]string{
			"proto_id":  "11s3"+strconv.Itoa(i),
			"device_id":   "dddd",
		}
		fields := map[string]interface{}{
			"value": float64(i),
			"sample_time": "dddddd",
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
	writesPointsPure(db,bp)
	//for i:=0;i<n;i++{
		//<-out
	//}
    //close(out)
}

//往influxdb存入造的传感器的数据
func writesPointsPure(conn client.Client,bp client.BatchPoints ) {
	t1:=time.Now()
	if err := conn.Write(bp); err != nil {
		log.Fatal(err)
	} else{
		fmt.Println("写入成功")
		//out <- 1
		t2:=time.Now()
		t:=t2.UnixNano()-t1.UnixNano()
		fmt.Println("dsddddddddd")
		fmt.Println(t/1000000000)
	}


}
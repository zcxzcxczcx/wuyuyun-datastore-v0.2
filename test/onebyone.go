package test

import (
	"wuyuyun-datastore-v0.2/conn"
	"fmt"
	"encoding/json"
	"strconv"
	"wuyuyun-datastore-v0.2/mqtt"
	MQTT "github.com/eclipse/paho.mqtt.golang"//eclipse组织
	"log"
	"wuyuyun-datastore-v0.2/mqtt/structs"
	"github.com/influxdata/influxdb/client/v2"
	"wuyuyun-datastore-v0.2/config"
	"math/rand"

	"time"
)
var n int
func MqttSubStoreInInfluxdbTestOneByOnE(topic string){
	//连接Influxdb
	c:=mqtt.BrokerSub()
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	db,err := conn.ConnInfluxdbTest()
	if err!=nil {
		fmt.Println("连接influxdb出现错误",err)
	}
	defer db.Close()
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  config.MyDB_official,
	})
	if err != nil {
		log.Fatal(err)
	}

	//处理器
	var callback MQTT.MessageHandler = func(c MQTT.Client, msg MQTT.Message) {
		//client是我自己的
		fmt.Printf("TOPIC: %s\n", msg.Topic())
		//fmt.Printf("没有Unmarshal的Message: %s\n", msg.Payload())
		//从 msg.Topic()中拿到devid
		devid:=GetProidFromTopid(msg.Topic())

		//fmt.Println(devid)
		var sub_data []structs.Productdata
		json.Unmarshal(msg.Payload(),&sub_data)
		//fmt.Printf("Unmarshal的Message: %v\n", sub_data)
		//tags := map[string]string{
		//}
		fmt.Println(1)
		for i:=0; i<len(sub_data);i++ {
			n++
			fmt.Println(n)
			tags := map[string]string{
				"proto_id":  "11s3"+strconv.Itoa(n),
				"device_id":   devid,
			}
			fields := map[string]interface{}{
				"value": float64(sub_data[i].Value)+float64(rand.Intn(100000000000)),
				"sample_time": sub_data[i].SampleTime,
			}
			pt, err := client.NewPoint(
				"iot_device_data_test",
				tags,
				fields,
				sub_data[i].SampleTime,
			)
			if err != nil{
				log.Fatal(err)
			}
			bp.AddPoint(pt)

			writesPointsOneByOne(db,bp)
		}
	}
	token := c.Subscribe(topic, 1, callback)

	if  token.Wait()&&token.Error() != nil {
		fmt.Printf("%v\n", token.Error())
	} else {
		fmt.Printf("Sub success\n")
	}
	for {
		time.Sleep(1*time.Second)
	}

}

//往influxdb存入传感器的数据
//我暂时认为Write()就是一次http写入
func writesPointsOneByOne(conn client.Client,bp client.BatchPoints) {
	fmt.Println(5)
	if err := conn.Write(bp); err != nil {
		log.Fatal(err)
	} else{
		fmt.Println("写入influxdb成功")
	}
}

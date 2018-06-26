package test

import (
	"wuyuyun-datastore-v0.2/conn"
	"fmt"
	"encoding/json"
	"strconv"
	"time"
	"wuyuyun-datastore-v0.2/mqtt"
	MQTT "github.com/eclipse/paho.mqtt.golang"//eclipse组织
	"log"
	"wuyuyun-datastore-v0.2/mqtt/structs"
	"github.com/influxdata/influxdb/client/v2"
	"wuyuyun-datastore-v0.2/config"
	"math/rand"
)


func MqttSubStoreInInfluxdbTest(topic string){
	//连接Influxdb
	c:=mqtt.BrokerTest()
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
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

	var sub_data_cache11 [10000]structs.Productdata//缓冲管道1
	var sub_data_cache1_copy [10000]structs.Productdata//备份数组
	var index int
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
				fmt.Println(2)
				sub_data_cache11[index] = sub_data[i]
				if index == 9999 {
					fmt.Println(3)
					sub_data_cache1_copy = sub_data_cache11//备份
					go func() {
						fmt.Println(4)
						writesPoints(db, bp, sub_data_cache1_copy, devid)
					}()
					index =-1
				}
				fmt.Println(7)
				index++
			}
	}
	token := c.Subscribe(topic, 1, callback)

	if  token.Wait()&&token.Error() != nil {
		fmt.Printf("%v\n", token.Error())
	} else {
		fmt.Printf("Sub success\n")
	}

	for i:=0;i<5;i++{
		time.Sleep(12*time.Second)
	}

}

//往influxdb存入传感器的数据
//我暂时认为Write()就是一次http写入
func writesPoints(conn client.Client,bp client.BatchPoints,sub_data_cache1 [10000]structs.Productdata,devid string) {
	fmt.Println(5)
	//time.Sleep(10*time.Second) //用来测试，防止他太快
	for i:=0; i<len(sub_data_cache1);i++{
		tags := map[string]string{
			"proto_id":  strconv.Itoa(int(sub_data_cache1[i].AgrId)),
			"device_id":   devid,
		}
		fields := map[string]interface{}{
			"value": float64(sub_data_cache1[i].Value)+float64(rand.Intn(100000000000)),
			"sample_time": sub_data_cache1[i].SampleTime,
		}
		pt, err := client.NewPoint(
			"iot_device_data_test",
			tags,
			fields,
			sub_data_cache1[i].SampleTime,
		)
		if err != nil{
			log.Fatal(err)
		}
		bp.AddPoint(pt)
	}
	fmt.Println(6)
	if err := conn.Write(bp); err != nil {
		log.Fatal(err)
	} else{
		fmt.Println("写入influxdb成功")
	}

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

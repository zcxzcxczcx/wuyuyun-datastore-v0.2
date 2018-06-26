package mqtt
import (
	"github.com/eclipse/paho.mqtt.golang/packets"
	MQTT "github.com/eclipse/paho.mqtt.golang"//eclipse组织
	"fmt"
	"time"
	"wuyuyun-datastore-v0.2/mqtt/sub"
)

type NoOpStore struct {
	// Contain nothing
}
func (store *NoOpStore) Open() {
	// Do nothing
}
func (store *NoOpStore) Put(string, packets.ControlPacket) {
	// Do nothing
}
func (store *NoOpStore) Get(string) packets.ControlPacket {
	// Do nothing
	return nil
}
func (store *NoOpStore) Del(string) {
	// Do nothing
}
func (store *NoOpStore) All() []string {
	return nil
}
func (store *NoOpStore) Close() {
	// Do Nothing
}
func (store *NoOpStore) Reset() {
	// Do Nothing
}

func OnLostBySub(c MQTT.Client,err error)   {
	fmt.Println("连接丢失")
	//EOF是文件结束符（End Of File）的缩写
	fmt.Println(err)
}

/************************** mysql**************************/

func OnConnectHandlerBySub(c MQTT.Client)   {
	fmt.Println("连接成功")
	sub.MqttSubStoreInMysql("+/data","+/status",c)
}

func OnConnectHandlerBySubByInfluxdb(c MQTT.Client){
	fmt.Println("连接成功")
	sub.MqttSubStoreToInfluxdb(c,"+/data")
}
func BrokerSub() MQTT.Client{
	//willPayload:=structs.StatusPayload{3}
	//willPayloadJ,_:=json.Marshal(willPayload)
	myNoOpStore := &NoOpStore{}
	opts := MQTT.NewClientOptions()
	opts.AddBroker("192.168.128.129:1883")
	//opts.AddBroker("localhost:1883")
	opts.SetClientID("this_is_sub_8")
	opts.SetUsername("this_is_sub_8")
	opts.SetPassword("112233")
	opts.SetAutoReconnect(true)
	opts.SetConnectionLostHandler(OnLostBySub)
	//opts.SetOnConnectHandler(OnConnectHandlerBySubByInfluxdb)
	opts.SetProtocolVersion(3)
	opts.SetMaxReconnectInterval( 1 * time.Second)
	opts.SetStore(myNoOpStore)//retain 客户端存储 测试存储 服务质量等级重发
	//opts.SetWill("1/status",string(willPayloadJ),1,true)
	c := MQTT.NewClient(opts)
	return c
}


/************************** influxdb 121.199.18.4**************************/


func BrokerSubToInfluxdb() MQTT.Client{
	//willPayload:=structs.StatusPayload{3}
	//willPayloadJ,_:=json.Marshal(willPayload)
	myNoOpStore := &NoOpStore{}
	opts := MQTT.NewClientOptions()
	opts.AddBroker("121.199.18.4:1883")
	//opts.AddBroker("localhost:1883")
	opts.SetClientID("this_is_sub_4")
	opts.SetUsername("this_is_sub_4")
	opts.SetPassword("112233")
	opts.SetAutoReconnect(true)
	opts.SetConnectionLostHandler(OnLostBySub)
	opts.SetOnConnectHandler(OnConnectHandlerBySub)
	opts.SetProtocolVersion(3)
	opts.SetMaxReconnectInterval( 1 * time.Second)
	opts.SetStore(myNoOpStore)//retain 客户端存储 测试存储 服务质量等级重发
	//opts.SetWill("1/status",string(willPayloadJ),1,true)
	c := MQTT.NewClient(opts)
	return c
}
/************************** influxdb xiayan**************************/

func BrokerSubToInfluxdbXiayan() MQTT.Client{
	//willPayload:=structs.StatusPayload{3}
	//willPayloadJ,_:=json.Marshal(willPayload)
	myNoOpStore := &NoOpStore{}
	opts := MQTT.NewClientOptions()
	opts.AddBroker("116.62.48.231:1883")
	//opts.AddBroker("localhost:1883")
	opts.SetClientID("this_is_sub_xiayan_3")
	opts.SetUsername("this_is_sub_xiayan_3")
	opts.SetPassword("112233")
	opts.SetAutoReconnect(true)
	opts.SetConnectionLostHandler(OnLostBySub)
	opts.SetOnConnectHandler(OnConnectHandlerBySub)
	opts.SetProtocolVersion(3)
	opts.SetMaxReconnectInterval( 1 * time.Second)
	opts.SetStore(myNoOpStore)//retain 客户端存储 测试存储 服务质量等级重发
	//opts.SetWill("1/status",string(willPayloadJ),1,true)
	c := MQTT.NewClient(opts)
	return c
}

/************************** other**************************/
func Broker() MQTT.Client{
	//willPayload:=structs.StatusPayload{3}
	//willPayloadJ,_:=json.Marshal(willPayload)
	myNoOpStore := &NoOpStore{}
	opts := MQTT.NewClientOptions()
	opts.AddBroker("121.199.18.4:1883")
	//opts.AddBroker("localhost:1883")
	opts.SetClientID("this_is_zcx")
	opts.SetUsername("this_is_zcx")
	opts.SetPassword("112233")
	opts.SetStore(myNoOpStore)//retain 客户端存储 测试存储 服务质量等级重发
	//opts.SetWill("1/status",string(willPayloadJ),1,true)
	c := MQTT.NewClient(opts)
	return c
}
func BrokerTest() MQTT.Client{
	//willPayload:=structs.StatusPayload{3}
	//willPayloadJ,_:=json.Marshal(willPayload)
	myNoOpStore := &NoOpStore{}
	opts := MQTT.NewClientOptions()
	//opts.AddBroker("121.199.18.4:1883")
	opts.AddBroker("121.199.18.4:1883")
	opts.SetClientID("2")
	opts.SetUsername("zcxesssxzcxzcxzczxcxzcxzczxcxccczccczczxczcc")
	opts.SetPassword("112233333")
	opts.SetStore(myNoOpStore)//retain 客户端存储 测试存储 服务质量等级重发
	//opts.SetWill("2/status",string(willPayloadJ),1,true)
	c := MQTT.NewClient(opts)
	return c
}
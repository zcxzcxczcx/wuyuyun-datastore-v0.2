package main

import (
	//"wuyuyun-datastore-v0.2/mqtt/sub"
	_ "net/http/pprof"
	//"wuyuyun-datastore-v0.2/mqtt/sub/status"
	//"time"
	//"wuyuyun-datastore-v0.2/test"
	//"time"
	//"wuyuyun-datastore-v0.2/mqtt"
	//"time"
	//"wuyuyun-datastore-v0.2/mqtt/sub/brokerInfo"
	//"time"
	"wuyuyun-datastore-v0.2/test"
	//"log"
	//"net/http"
)

func main() {
	/**************************  只要main函数不退出,那些goroutine就会执行完自己的生命周期  **************************/
	//go func() {
	//	log.Println(http.ListenAndServe("localhost:6060", nil))
	//}()
	//go status.MqttSubClientStatus("+/status")
	//go data.MqttSubStoreInMysql("+/data")
	//go brokerInfo.MqttSubBrokerInfo("$SYS/broker/connection/#")
	//test.MakeData()

	//test.MqttSubStoreInInfluxdbTest("+/data")
	//test.StoreIntoInfluxdb(1500000)
    //test.StoreIntoInfluxdb(350000)

  test.MqttSubStoreInInfluxdbTestOneByOnE("+/data")


	//c:=mqtt.BrokerSub()
	//if token := c.Connect(); token.Wait() && token.Error() != nil {
	//	panic(token.Error())
	//}
   //for{
	//	time.Sleep(1*time.Second)
	//}

	//test.StoreIntoInfluxdb(15000)
	//test.MakedatToInfluxdab()


}
package conn
import (
	"wuyuyun-datastore-v0.2/config"
	"github.com/influxdata/influxdb/client/v2"
)
// 连接influxdb数据库 我的虚拟机
func ConnInfluxdbTest() (client.Client,error) {
	conn,err := client.NewHTTPClient(client.HTTPConfig{
			Addr:     "http://localhost:8086",
		Username: config.Username,
		Password: config.Password,
	})
	if err != nil {
		return nil,err
	}
	return conn,nil
}

// 连接influxdb数据库
func ConnInfluxdb_121() (client.Client,error) {
	conn,err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     "http://localhost:8086",
		Username: config.Username,
		Password: config.Password,
	})
	if err != nil {
		return nil,err
	}
	return conn,nil
}
// 连接influxdb数据库
func ConnInfluxdb_xiayan() (client.Client,error) {
	conn,err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     "http://localhost:8086",
		Username: config.Username,
		Password: config.Password,
	})
	if err != nil {
		return nil,err
	}
	return conn,nil
}


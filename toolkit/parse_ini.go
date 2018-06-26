package toolkit

import (
	"github.com/vaughan0/go-ini"
	"strconv"
)
type Config_info struct{
	Adapter string
	Host string
	Port int
	Username string
	Password string
	Dbname string
	Charset string
}
//解析ini文件
func ParseInifile(filepath string) (config_info Config_info){
	file, _ := ini.LoadFile(filepath)
	config_info.Adapter, _= file.Get("database","adapter")
	config_info.Host, _ = file.Get("database", "host")
	port, _ :=file.Get("database","port")
	config_info.Port, _ =	strconv.Atoi(port)
	config_info.Username, _ = file.Get("database", "username")
	config_info.Password, _ = file.Get("database","password")
	config_info.Dbname, _ = file.Get("database", "dbname")
	config_info.Charset, _ = file.Get("database","charset")
	return
}
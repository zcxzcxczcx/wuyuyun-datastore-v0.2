package config

import (
	"strconv"
	"flag"
	"wuyuyun-datastore-v0.2/toolkit"
)
type DataBaseStruct struct {
	Adapter   string
	Host      string
	Port   	  int
	Username  string
	Password  string
	Dbname    string
	Charset   string
	Dsn string
}

var DataBase1 DataBaseStruct
var DataBase2 DataBaseStruct
var DataBase3 DataBaseStruct
var DataBase4 DataBaseStruct
var DataBase5 DataBaseStruct

var config_file *string = flag.String("mysql_config", "./file/config.ini", "Use -mysql_congif <filesource>")

func init()  {
	flag.Parse()
	config_info:=toolkit.ParseInifile(*config_file)
	//DataBase1
	DataBase1.Adapter = config_info.Adapter
	DataBase1.Host = config_info.Host
	DataBase1.Port = config_info.Port
	DataBase1.Username = config_info.Username
	DataBase1.Password = config_info.Password
	DataBase1.Dbname = config_info.Dbname
	DataBase1.Charset = config_info.Charset
	DataBase1.Dsn = DataBase1.Username + ":" + DataBase1.Password + "@tcp(" + DataBase1.Host + ":" + strconv.Itoa(DataBase1.Port) + ")/" + DataBase1.Dbname

	//DataBase2
	DataBase2.Adapter = "mysql"
	DataBase2.Host = "influxdb.go"
	DataBase2.Port = 3306
	DataBase2.Username = "root"
	DataBase2.Password = ""
	DataBase2.Dbname = ""
	DataBase2.Charset = "utf8"
	DataBase2.Dsn = DataBase2.Username + ":" + DataBase2.Password + "@tcp(" + DataBase2.Host + ":" + strconv.Itoa(DataBase2.Port) + ")/" + DataBase2.Dbname

}

package handler

import (
	"github.com/gorilla/mux"
	"wuyuyun-datastore-v0.2/api/handler/protocol"
	"wuyuyun-datastore-v0.2/api/handler/onlineRate"
	)

func CustomHandler(r *mux.Router){
	//data数据统计
	// 路径前缀
	//粒度是月份的必须是2018或者2018-05；粒度是天，时，分是以from:2018-05-02,to:2018-05-08这种形式；单一的某个日期如2018-05-03则以粒度是时，分。

		dataStatisticsRouter := r.PathPrefix("/data").Subrouter()
	// 子路由

	/************************** mysql **************************/

	//获取今天的功能项的数据的route
	//设备id:devid;产品协议项id:protocolid
	dataStatisticsRouter.HandleFunc("/{devid:[0-9A-Za-z]+}/{protocolid:[0-9]+}/today", protocol.ListToday).Methods("GET")
	//获取昨天的功能项的数据的route
	dataStatisticsRouter.HandleFunc("/{devid:[0-9A-Za-z]+}/{protocolid:[0-9]+}/yesterday", protocol.ListYesterday).Methods("GET")
	//获取前七天的功能项的数据的route
	dataStatisticsRouter.HandleFunc("/{devid:[0-9A-Za-z]+}/{protocolid:[0-9]+}/sevendaysago", protocol.ListSevendaysAgo).Methods("GET")
	//获取前一个月的功能项的数据的route
	dataStatisticsRouter.HandleFunc("/{devid:[0-9A-Za-z]+}/{protocolid:[0-9]+}/onemonthago", protocol.ListOneMonthAgo).Methods("GET")
	//以月份为粒度返回，from，to为2019-04种格式，不能有当月
	dataStatisticsRouter.HandleFunc("/{devid:[0-9A-Za-z]+}/{protocolid:[0-9]+}/month", protocol.MonthReturnTest).Methods("GET")

	/************************** influxdb **************************/
	//获取今天的功能项的数据的route
	//产品id:prodid;产品协议项id:protocolid
	dataStatisticsRouter.HandleFunc("/{devid:[0-9A-Za-z]+}/{protocolid:[0-9]+}/today2", protocol.ListToday2).Methods("GET")
	//获取昨天的功能项的数据的route
	dataStatisticsRouter.HandleFunc("/{devid:[0-9A-Za-z]+}/{protocolid:[0-9]+}/yesterday2", protocol.ListYesterday2).Methods("GET")
	//获取前七天的功能项的数据的route
	dataStatisticsRouter.HandleFunc("/{devid:[0-9A-Za-z]+}/{protocolid:[0-9]+}/sevendaysago2", protocol.ListSevendaysAgo2).Methods("GET")
	//获取前一个月的功能项的数据的route
	dataStatisticsRouter.HandleFunc("/{devid:[0-9A-Za-z]+}/{protocolid:[0-9]+}/onemonthago2", protocol.ListOneMonthAgo2).Methods("GET")





	//status统计
	// 路径前缀
	statusStatisticsRouter := r.PathPrefix("/status").Subrouter()

	//获取所有项目的设备在线率
	statusStatisticsRouter.HandleFunc("/online_rate_project_all", onlineRate.ProjAllOnlineRate).Methods("GET")
	//获取某个项目的设备的在线率
	statusStatisticsRouter.HandleFunc("/{projectid:[0-9]+}/online_rate_project", onlineRate.ProjOnlineRate).Methods("GET")
	//获取某个产品的设备的在线率
	statusStatisticsRouter.HandleFunc("/{productid:[0-9]+}/online_rate_product", onlineRate.ProdOnlineRate).Methods("GET")
	//获取某个设备的在线率//from,to形式
	// 力度intensity =1是天,力度intensity =2是 时
	statusStatisticsRouter.HandleFunc("/{deviceid:[0-9A-Za-z]+}/online_rate_dev", onlineRate.DevOnlineRate).Methods("GET")

}
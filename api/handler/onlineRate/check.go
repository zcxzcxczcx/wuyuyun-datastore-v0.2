package onlineRate

import (
	"net/http"
	"github.com/gorilla/mux"
	"wuyuyun-datastore-v0.2/api/handler/onlineRate/restructure"
	"wuyuyun-datastore-v0.2/api/common"
	"strconv"
	"fmt"
)
//采样计算
//所有项目的在线率//intensity_num = 1按天返回,每天零点去统计一次
// intensity_num = 2按小时返回 每小时的零分去统计一次
func ProjAllOnlineRate(w http.ResponseWriter, r *http.Request)  {
	from := r.FormValue("from")
	to := r.FormValue("to")
	intensity := r.FormValue("intensity")
	if from == ""{
		common.JsonResponse("from不能为空",500,w)
		return
	}
	if to == ""{
		common.JsonResponse("to不能为空",500,w)
		return
	}
	if intensity == ""{
		common.JsonResponse("力度不能为空",500,w)
		return
	}
	intensity_num,_:=strconv.Atoi(intensity)
	if intensity_num == 1{
		data:=restructure.RestructureOnlinerateDataProjAllByDay(from,to)
		common.JsonResponse(data,200,w)
		return
	}
	if intensity_num == 2{
		data:=restructure.RestructureOnlinerateDataProjAllByHour(from,to)
		common.JsonResponse(data,200,w)
		return
	}
}
//某个项目的在线率
func ProjOnlineRate(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("1")
	from := r.FormValue("from")
	to := r.FormValue("to")
	intensity := r.FormValue("intensity")
	vars := mux.Vars(r)

	projectid := vars["projectid"]
	if from == ""{
		common.JsonResponse("from不能为空",500,w)
		return
	}
	if to == ""{
		common.JsonResponse("to不能为空",500,w)
		return
	}
	if intensity == ""{
		common.JsonResponse("力度不能为空",500,w)
		return
	}
	intensity_num,_:=strconv.Atoi(intensity)
	projectid_num,_:=strconv.Atoi(projectid)
	if intensity_num == 1{
		data:=restructure.RestructureOnlinerateDataProjByDay(from,to,projectid_num)
		common.JsonResponse(data,200,w)
		return
	}
	if intensity_num == 2{
		fmt.Println("2")
		data:=restructure.RestructureOnlinerateDataProjByHour(from,to,projectid_num)
		common.JsonResponse(data,200,w)
		return
	}


}
//某个产品的在线率
func ProdOnlineRate(w http.ResponseWriter, r *http.Request)  {
	from := r.FormValue("from")
	to := r.FormValue("to")
	intensity := r.FormValue("intensity")
	vars := mux.Vars(r)

	productid := vars["productid"]
	if from == ""{
		common.JsonResponse("from不能为空",500,w)
		return
	}
	if to == ""{
		common.JsonResponse("to不能为空",500,w)
		return
	}
	if intensity == ""{
		common.JsonResponse("力度不能为空",500,w)
		return
	}
	intensity_num,_:=strconv.Atoi(intensity)
	productid_num,_:=strconv.Atoi(productid)
	if intensity_num == 1{
		data:=restructure.RestructureOnlinerateDataProdByDay(from,to,productid_num)
		common.JsonResponse(data,200,w)
		return
	}
	if intensity_num == 2{
		data:=restructure.RestructureOnlinerateDataProdByHour(from,to,productid_num)
		common.JsonResponse(data,200,w)
		return
	}


}
//某个设备的在线率 力度intensity =1是天,力度intensity =2是 时
func DevOnlineRate(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	deviceid := vars["deviceid"]
	from := r.FormValue("from")
	to := r.FormValue("to")
	intensity := r.FormValue("intensity")
	if from == ""{
		common.JsonResponse("from不能为空",500,w)
		return
	}
	if to == ""{
		common.JsonResponse("to不能为空",500,w)
		return
	}
	if intensity == ""{
		common.JsonResponse("力度不能为空",500,w)
		return
	}
	var data interface{}
	if  intensity_num,_:=strconv.Atoi(intensity);intensity_num == 1{
		data =restructure.RestructureOnlinerateDataDevFromToByDay(deviceid,from,to)
	}
	if  intensity_num,_:=strconv.Atoi(intensity);intensity_num == 2{
		data =restructure.RestructureOnlinerateDataDevFromToByHour(deviceid,from,to)
	}
	common.JsonResponse(data,200,w)
}


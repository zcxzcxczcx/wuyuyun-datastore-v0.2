package datastructs

//golang json omitempty是什么意思:为空则不输出
type DataAgr struct {
	TimePoint       int  	`json:"time_point"`
	ValueAvg        float64    	`json:"value_avg"`
	ValueMax      float64 `json:"value_max"`
	ValueMin     float64 `json:"value_min"`
}
type DataAgrdays struct {
	TimePoint       string   	`json:"time_point"`
	ValueAvg        float64    	`json:"value_avg"`
	ValueMax      float64 `json:"value_max"`
	ValueMin     float64 `json:"value_min"`
}
type CustomTimeRes struct {
	ProId         string         `json:"pro_id"`
	AgrId          int64         `json:"agr_id"`
	Type           int            `json:"type"`//1:按年；2：按月；3：按日
	StartTime      string   	`json:"start_time"`
	EndTime       string    	`json:"end_time"`
}
type CustomData struct {
	TimePoint       string   	`json:"time_point"`
	ValueAvg        float64    	`json:"value_avg"`
}
//参数结构体
//type Parameter struct {
//	ProId     int   	`json:"pro_id"`
//	AgrId       int    	`json:"agr_id"`
//}
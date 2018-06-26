package structs

import "time"

//产品数据的结构体
type Productdata struct {
	AgrId int32
	CreateTime time.Time
	Value float32
	SampleTime time.Time
}

//控制继电器
//继电器返回的结果的结构体
type DirectiveData struct {
	Relay uint16
	Directive int
	Value bool
	PubTime time.Time
	ResTime time.Time
}
//type:1,正常上线;2,正常下线;3,异常下线
type StatusPayload struct {
	Type  int  `json:"type"`
}



type DataToInfluxdb struct {
	AgrId int32
    DeviceId string
	Value float64
	SampleTime string
}
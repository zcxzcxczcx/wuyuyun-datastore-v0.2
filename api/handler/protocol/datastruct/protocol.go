package datastruct


//按日期查找的数据结构（未重构之前的）
type DateAgr struct {
	Date       string `json:"date"`
	ValueAvg        float64    	`json:"value_avg"`
	ValueMax      float64 `json:"value_max"`
	ValueMin     float64 `json:"value_min"`
}
type Predays struct {
	From string `json:"from"`
	To string `json:"to"`
	Min float64 `json:"min"`
	Max float64 `json:"max"`
	Avg float64 `json:"avg"`
	Dates []DateAgr `json:"dates"`
}
//整理得到的几天的数据，返回给前端需要的数据结构
type Sevdays struct {
	From string `json:"from"`
	To string `json:"to"`
	Min float64 `json:"min"`
	Max float64 `json:"max"`
	Avg float64 `json:"avg"`
	AvgMax float64 `json:"avg_max"`
	AvgMin float64 `json:"avg_min"`
	Dates []Datesbyhour `json:"dates"`
}
type Monthdays struct {
	From string `json:"from"`
	To string `json:"to"`
	Min float64 `json:"min"`
	Max float64 `json:"max"`
	Avg float64 `json:"avg"`
	AvgMax float64 `json:"avg_max"`
	AvgMin float64 `json:"avg_min"`
	Dates []Dates `json:"dates"`
}
type Dates struct {
	Date string `json:"date"`
	Min float64 `json:"min"`
	Max float64 `json:"max"`
	AvgMax float64 `json:"avg_max"`
	AvgMin float64 `json:"avg_min"`
	Avg float64 `json:"avg"`
}
type Datesbyhour struct {
	Date string `json:"date"`
	Min float64 `json:"min"`
	Max float64 `json:"max"`
	Avg float64 `json:"avg"`
	AvgMax float64 `json:"avg_max"`
	AvgMin float64 `json:"avg_min"`
	Hours []Hours `json:"hours"`
}
type Hours struct {
	Hour int `json:"hour"`
	Max float64 `json:"max"`
	Min float64 `json:"min"`
	Avg float64 `json:"avg"`
}


//重构的数据结构
type Todays struct {
	Date string `json:"date"`
	Min float64 `json:"min"`
	Max float64 `json:"max"`
	Avg float64 `json:"avg"`
	AvgMax float64 `json:"avg_max"`
	AvgMin float64 `json:"avg_min"`
	Hours []Hours `json:"hours"`
}
/************************** 以下的粒度是从年到分**************************/
type SureData struct {
	From string `json:"from"`
	To string `json:"to"`
	Min float64 `json:"min"`
	Max float64 `json:"max"`
	Avg float64 `json:"avg"`
	Month []Month `json:"month"`
}

type Month struct {
	Month string `json:"month"`
	Min float64 `json:"min"`
	Max float64 `json:"max"`
	Avg float64 `json:"avg"`
}

//type Day struct {
//	Day string `json:"day"`
//	Min float64 `json:"min"`
//	Max float64 `json:"max"`
//	Avg float64 `json:"avg"`
//}

//
//type Hour struct {
//	Hour string `json:"hour"`
//	Min float64 `json:"min"`
//	Max float64 `json:"max"`
//	Avg float64 `json:"avg"`
//	Minute []Minute `json:"day"`
//}
//
//type Minute struct {
//	Minute string `json:"minute"`
//	Min float64 `json:"min"`
//	Max float64 `json:"max"`
//	Avg float64 `json:"avg"`
//}
//

type DataFromInfluxdb struct {
	From string `json:"from"`
	To string `json:"to"`
	Max float64 `json:"max"`
	Min float64 `json:"min"`
	Avg float64 `json:"avg"`
	Median float64 `json:"median"`
	AvgMax float64 `json:"avg_max"`
	AvgMin float64 `json:"avg_min"`
	MedianMax float64 `json:"median_max"`
	MedianMin float64 `json:"median_min"`
	Datas []Datas `json:"datas"`
}
type Datas struct {
	Time string `json:"time"`
	Mean interface{} `json:"mean"`
	Median interface{} `json:"median"`
	Max interface{} `json:"max"`
	Min interface{} `json:"min"`
}



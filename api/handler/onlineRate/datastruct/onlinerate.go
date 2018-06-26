package datastruct


type RateData struct {
	TimePoint string `json:"time_point"`
	OnlineRate float64 `json:"online_rate"`
}

type InitData struct {
	TimePoint int `json:"time_point"`
	OnlineRate float64 `json:"online_rate"`
}

type RateData2 struct {
	TimePoint string `json:"time_point"`
	number int `json:"number"`
}

type RateDataByHour struct {
	TimePoint string `json:"time_point"`
	OnlineRateArr []RateData `json:"online_rate"`
}



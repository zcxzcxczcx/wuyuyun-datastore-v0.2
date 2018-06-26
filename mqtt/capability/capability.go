package capability

import (
	"time"
	"fmt"
)


// 计算发送指令给继电器到继电器响应的时间差
func TimeDifference(sendouttime time.Time,relayrestime time.Time)  int64{
	timedef := relayrestime.UTC().UnixNano() - sendouttime.UTC().UnixNano()
	d := time.Duration(timedef)
	fmt.Printf("'String: %v', 'Nanoseconds: %v', 'Seconds: %v', 'Minutes: %v', 'Hours: %v'\n",
		d.String(), d.Nanoseconds(), d.Seconds(), d.Minutes(), d.Hours())
	return timedef
}
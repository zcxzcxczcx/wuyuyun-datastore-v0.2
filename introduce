
2.go time
time.Millisecond    表示1毫秒
time.Microsecond    表示1微妙
time.Nanosecond    表示1纳秒

time.Now().Unix() // 获取当前时间戳
time.Now().UnixNano()   // 精确到纳秒，通过纳秒就可以计算出毫秒和微妙

time.Now().Format("2006-01-02 15:04:05")   // 获取当前时间，进行格式化 //output: 2016-07-27 08:57:46
time.Unix(1469579899, 0).Format("2006-01-02 15:04:05") // 指定的时间进行格式化  // output: 2016-07-27 08:38:19

// 获取指定时间戳的年月日，小时分钟秒
 // output: 2016-7-27 8:38:19
t := time.Unix(1469579899, 0)
fmt.Printf("%d-%d-%d %d:%d:%d\n", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())

 // 先用time.Parse对时间字符串进行分析，如果正确会得到一个time.Time对象
t2, err := time.Parse("2006-01-02 15:04:05", "2016-07-27 08:46:15")
// 后面就可以用time.Time对象的函数Unix进行获取
// output:
//     2016-07-27 08:46:15 +0000 UTC
//     1469609175
fmt.Println(t2)
fmt.Println(t2.Unix())


//mosquitto
//前台启用服务器：mosquitto -c /etc/mosquitto/mosquitto.conf -v

//将打包好的文件上传到服务器
sudo scp ./data_store  www@121.199.18.4:/home/www
服务器/usr/sbin目录下 执行sudo scp /home/www/data_store  .



/************************** 测试influxdb **************************/
//{"error":"engine: error syncing wal"}-》并发太高
//connection reset by peer -》并发太高
//[shard 2] open /var/lib/influxdb/wal/wuyuyun_sub/autogen/2/_00026.wal: permission denied
/{"error":"partial write: max-values-per-tag limit exceeded (100001/100000): measurement=\"iot_device_data_test\" tag=\"proto_id\" value=\"10000011\" dropped=50000"}
//2018/06/01 01:06:25 Post http://121.199.18.4:8086/write?consistency=&db=wuyuyun_sub&precision=s&rp=: EOF
package conn

import (
	"database/sql"
	"wuyuyun-datastore-v0.2/config"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

//连接数据库store_data
func ConectDb1()  (*sql.DB, error ){
	db, err := sql.Open(config.DataBase1.Adapter,
		config.DataBase1.Dsn)
	if err != nil {
		log.Fatal(err)
	}
	return db, err
}

//连接数据库wuyuyun_0_1
func ConectDb2()  (*sql.DB, error ){
	db, err := sql.Open(config.DataBase2.Adapter,
		config.DataBase2.Dsn)
	if err != nil {
		log.Fatal(err)
	}
	return db, err
}


//连接数据库wuyuyun01
func ConectDb3()  (*sql.DB, error ){
	db, err := sql.Open(config.DataBase3.Adapter,
		config.DataBase3.Dsn)
	if err != nil {
		log.Fatal(err)
	}
	return db, err
}

//连接夏衍服务器数据库xyfc
func ConectDb4()  (*sql.DB, error ){
	db, err := sql.Open(config.DataBase4.Adapter,
		config.DataBase4.Dsn)
	if err != nil {
		log.Fatal(err)
	}
	return db, err
}
//连接夏衍服务器数据库
func ConectDb5()  (*sql.DB, error ){
	db, err := sql.Open(config.DataBase5.Adapter,
		config.DataBase5.Dsn)
	if err != nil {
		log.Fatal(err)
	}
	return db, err
}
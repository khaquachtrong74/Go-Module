package mysql

import (
	"database/sql"
	"log"
	"os"
	"sync"
	"time"
	"github.com/go-sql-driver/mysql"
)
var(
	one sync.Once
)

var	DBManager *sql.DB
var Cfg mysql.Config = mysql.Config{
	User: os.Getenv("DBUsername"),
	Passwd: os.Getenv("DBPassword"),
	Net: "tcp",
	Addr: "127.0.0.1:3306",
	DBName: "test",
	ParseTime: true,
	AllowNativePasswords: true,
}
func InitDB()(*sql.DB, error){
	var err error
	one.Do(func(){
		print(1)
		DBManager, err = sql.Open("mysql", Cfg.FormatDSN());
		if err != nil{
			log.Panic("Cann't connect DB: ", err)
		}
		DBManager.SetConnMaxLifetime(time.Minute*3)
		DBManager.SetMaxOpenConns(10)
		DBManager.SetMaxIdleConns(10)
	})
	return DBManager, err
}
func CloseDB()error {
	if DBManager == nil{
		return nil}
	print("Close here")
	return DBManager.Close()
}


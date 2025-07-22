package sql

import (
	"database/sql"
	"os"
	"sync"
	"time"
	"github.com/go-sql-driver/mysql"
	"github.com/khaquachtrong74/Go-Module/utils"
)

var(
	initOnce sync.Once
	DBManager *sql.DB
	err error
)

var Cfg mysql.Config = mysql.Config{
	User: os.Getenv("DBUsername"),
	Passwd: os.Getenv("DBPassword"),
	Net: "tcp",
	Addr: "127.0.0.1:3306",
	DBName: "test",
	ParseTime: true,
	AllowNativePasswords: true,
}
// Init connect one time - singleton
func InitDB()(*sql.DB, error){
	initOnce.Do(func(){
		DBManager, err = sql.Open("mysql", Cfg.FormatDSN());
		if err != nil{
			err = utils.ErrorWrap(err, "Failed to connect DB")
			return 
		}
		DBManager.SetConnMaxLifetime(time.Minute*3)
		DBManager.SetMaxOpenConns(10)
		DBManager.SetMaxIdleConns(10)
	})
	return DBManager, err
}

func CloseDB() error {
	if DBManager == nil{
		return nil
	}
	return DBManager.Close()
}


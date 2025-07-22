package sql

import (
	"database/sql"
	"fmt"
	"github.com/khaquachtrong74/Go-Module/utils"
)
func ExecAttribute(db *sql.DB, tableName string)error{
	var task []byte
	fmt.Println("Add content right now! ( ｡ •`ᴖ´• ｡)")
	fmt.Scanln(&task)
	_, err := db.Exec("INSERT INTO ?(content) VALUE(?);",tableName, task)
	if err != nil{
		return utils.ErrorWrap(err, "write query failed!")
	}
	return nil
}

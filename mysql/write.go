package mysql
import (
	"database/sql"
	"fmt"
)
func Write(DBM *sql.DB){
	var task []byte
	fmt.Println("Add task right now! ( ｡ •`ᴖ´• ｡)")
	fmt.Scanln(&task)
	_, err := DBM.Exec("INSERT INTO work(task) VALUE(?);", task)
	if err != nil{
		fmt.Errorf("Cook! here")
	}
	fmt.Print("(  ≧ᗜ≦)")
}

package mysql	
import (
	"database/sql"
	"log"
)

func ReadDB(DBM *sql.DB)[]Task{
	var list []Task
	if DBM != nil{
		queryResult, err := DBM.Query("SELECT id, task FROM work;")
		if err != nil{
			log.Panic("Something's wrong here")
			return nil
		}else{
			for queryResult.Next(){
				var t Task 
				if err := queryResult.Scan(&t.Id, &t.Task); err != nil{
					return nil
				}
				list = append(list, t)
			}
		}
	}
	return list
}


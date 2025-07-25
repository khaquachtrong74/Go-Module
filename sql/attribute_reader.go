package sql

import (
	"database/sql"
	"github.com/khaquachtrong74/Go-Module/utils"
)

func QueryAttributes(db *sql.DB, tableName string, query string)([]Attribute, error){
	rows, err := db.Query(query) 
	if err != nil{
		return nil, utils.ErrorWrap(err, "query failed!")
	}
	defer rows.Close()
	var result []Attribute
	for rows.Next(){
		var attr Attribute
		if err := rows.Scan(&attr.Id, &attr.Content); err != nil{
			return nil, utils.ErrorWrap(err, "scan failed!")
		}
		result = append(result, attr)
	}
	return result, nil
}
func GetContent(attr Attribute)string{
	return attr.Content
}

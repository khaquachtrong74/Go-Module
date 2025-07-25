package sql
type Task struct{
	Id int 
	Task string
}

type Table struct{
	tableName string
}
type Attribute struct{
	Id int
	Content string
}
func GetContent(attr Attribute)string{
	return attr.Content
}

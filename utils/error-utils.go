package utils
import "fmt"
func ErrorWrap(err error, context string) error {
	if err != nil{
		return fmt.Errorf("%s : %w",context, err) 
	}

	return nil
}

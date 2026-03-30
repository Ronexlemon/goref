package json

import (
	"encoding/json"
	"fmt"
)


func UnmarshialToken(b []byte)*Token{
	var t Token
	err:=json.Unmarshal(b,&t)
	if err !=nil{
		fmt.Println("Failed to unmarshal")
		return nil
	}
	return &t

}
package main

import (
	"fmt"

	myJson "github.com/ronexlemon/dsa/serialization/Json"
)


func main(){
	token := myJson.Token{Name: "USD circle",Symbol: "USDC",Decimal: 6}
	//t:=myJson.NewToken(token)
	
	marshal:=myJson.MarshialToken(token)
	fmt.Println("Marshal token",marshal)
	unmarshal := myJson.UnmarshialToken(marshal)
	fmt.Println("Tokens",unmarshal)
	
}
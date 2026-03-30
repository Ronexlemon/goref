package json

import (
	"encoding/json"
	"fmt"
)


type Token struct{
	Name string
	Symbol string
	Decimal int64
}

func NewToken(t Token)*Token{
	return &Token{Name: t.Name,Symbol: t.Symbol,Decimal: t.Decimal}
}

func (t *Token) TokenName()string{
	return t.Name
}

func (t *Token) TokenSymbol()string{
	return t.Symbol
}

func (t *Token) TokenDecimal()int64{
	return t.Decimal
}

func MarshialToken(t Token)[]byte{
	token := NewToken(t)
	b,err := json.Marshal(token)
	if err !=nil{
		fmt.Println("Failed to marshial")
		return  nil
	}
	return b
}
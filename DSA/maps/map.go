package main

import "fmt"

/****
-> used to store data values in key:value pairs
-> A map is an unordered and changeable collection that does not allow duplicates
Ways of creating a map
1. using var
	var a =map[keyType]ValueType{key1:value1,...}
2. usin make
	var a = make(map[keyType]ValueType)
3. an empty map
	var a map[keytype]valuetype
**/


func UsingVar(){
	var a =map[string]int{"one":1,"two":2,"three":3}
	 b :=map[string]int{"four":4,"five":5,"six":6}
	fmt.Printf("The length  a%d\n", len(a))
	fmt.Printf("The content a %v\n",a)

	fmt.Printf("The length b %d\n", len(b))
	fmt.Printf("The content b %v\n",b)
}

func UsingMake(){
	var a = make(map[string]int)
	a["year"]=2026
	a["age"]=24
	a["days"]=355

	b:=make(map[string]int)
	b["sal"]=170000
	b["ken"]=50000

	fmt.Printf("The length a %d\n", len(a))
	fmt.Printf("The content a %v\n",a)

	fmt.Printf("The length b %d\n", len(b))
	fmt.Printf("The content b %v\n",b)
}

func EmptyMap(){
	var a map[string]int
	var b =make(map[string]int)

	b["sal"]=170000
	b["ken"]=50000
	fmt.Printf("The length a %d\n", len(a))
	fmt.Printf("The content a %v\n",a)

	fmt.Printf("The length b %d\n", len(b))
	fmt.Printf("The content b %v\n",b)

}

func Delete(){
	var a = make(map[string]int)
	a["one"]=1
	a["two"]=2
	a["three"]=3
		fmt.Printf("The content a %v\n",a)
		delete(a,"one") //dlete(map_name,key)
		fmt.Printf("The content a %v\n",a)
	
}


/****
Maps Are reference to hash tables
NB: if two map variables refer to the same hash table, changing the content of one variable affect the content of the other
****/

func Reference(){
	var a =map[string]int{"one":1,"two":2,"three":3,"four":4,"five":5,"six":6,"seven":7}
	b:=a

	fmt.Printf("The content a %v\n",a)
	fmt.Printf("The content b %v\n",b)

	b["six"]=12
	fmt.Println("After changed b")
	fmt.Printf("The content a %v\n",a)
}

func Iterate(){
	var a =map[string]int{"one":1,"two":2,"three":3,"four":4,"five":5,"six":6,"seven":7}
	for k,v:=range a{
		fmt.Printf("%v: %v ",k,v)
	}

}
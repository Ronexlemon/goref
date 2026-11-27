package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main(){
	//os_app:=InitiateOS()
	bufio_app := InitiateBUFIO()
	//os_app.Opening_file("test.txt")
	//os_app.Create_file("text.txt")
	//os_app.Read_file("text.txt")
	//os_app.Bufio_file("yollow.txt")
	bufio_app.Read_input()

}

//opening a file

type OS struct{

}
type BUFIO struct{

}

func InitiateOS()(*OS){
	return &OS{}
}
func InitiateBUFIO()(*BUFIO){
	return &BUFIO{}
}

func (o *OS) Opening_file(fileName string){
	file,err := os.Open(fileName)
	if err !=nil{
		log.Fatal("Error opening a file",err)
		
	}
	fmt.Println(file.Name())
}

func (o *OS) Create_file(file_name string){
	file,err:=os.Create(file_name)
	if err !=nil{
		log.Fatal("Error Creating a file",err)
		
	}
	defer file.Close()
	data := []byte("hellow")
	file.Write(data)
}
func (o *OS) ReadFile_file(file_name string){
	data,err :=os.ReadFile(file_name)
	if err !=nil{
		log.Fatal("Error Reading a file",err)
		
	}
	fmt.Println(string(data))
}


func (o *OS) Read_file(file_name string){
	file,err :=os.Open(file_name)

	if err !=nil{
		log.Fatal(err)
	}
	defer file.Close()

	buffer:=make([]byte,100)

	for {
		n,err :=file.Read(buffer)
		if err == io.EOF{
			break
		}
		if err !=nil{
			log.Fatal(err)
		}
		fmt.Printf(string(buffer[:n]))
	}
	
}

func (o *OS) Bufio_file(file_name string){
	file,err :=os.Open(file_name)

	if err !=nil{
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}

	if err = scanner.Err();err !=nil{
		fmt.Println("The scanner got a  cut:", err)
	}
}

//bufio

func (b *BUFIO) Read_input(){
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Whats your name")

	for scanner.Scan(){
		fmt.Println("Your name is:",scanner.Text())
	}
	 if err := scanner.Err(); err != nil {
        fmt.Println("Error reading input:", err)
    }
}
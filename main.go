package main

import "fmt"

func Hello(name string) string{
	message := fmt.Sprintf("Ola %v. oia", name)
	return message
}

func main(){
	msg := Hello("Fernanda")
	fmt.Println(msg)
}
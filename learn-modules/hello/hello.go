package main


import (
 "fmt"

 "learning/greetings"
)


func main() {
	msg := greetings.Hello("D")
	fmt.Println(msg)
}

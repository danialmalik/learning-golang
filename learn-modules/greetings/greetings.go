package greetings

import "fmt"


func Hello(name string) string {
	msg := fmt.Sprintf("Hi, %v. Howdy!!!", name)
	return msg
}

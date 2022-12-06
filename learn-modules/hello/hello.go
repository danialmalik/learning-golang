package main


import (
 "fmt"
 "log"
 "learning/greetings"
)


func main() {
	// Set properties of the predefined Logger, including
    // the log entry prefix and a flag to disable printing
    // the time, source file, and line number.
    log.SetPrefix("greetings: ")
    log.SetFlags(0)


    // take user input
    var input string

    fmt.Print("Enter a name to be greeted: ")
    fmt.Scanln(&input)

	msg, error := greetings.Hello(input)
	if error != nil {
		log.Fatal(error)
	}
	fmt.Println(msg)
}

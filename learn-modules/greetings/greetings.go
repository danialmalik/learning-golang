package greetings

import (
	"fmt"
	"errors"
)


func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("No name is provided.")
	}

	msg := fmt.Sprintf("Hi, %v. Howdy!!!", name)
	return msg, nil
}

package greetings

import (
	"fmt"
	"errors"
	"math/rand"
    "time"
)


func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("No name is provided.")
	}

	msg := fmt.Sprintf(randomFormat(), name)
	return msg, nil
}


func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		message, error := Hello(name)

		if error != nil {
			return nil, error
		}

		messages[name] = message
	}

	return messages, nil
}


func init() {
	rand.Seed(time.Now().UnixNano())
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.
func randomFormat() string {
    // A slice of message formats.
    formats := []string{
        "Hi, %v. Welcome!",
        "Great to see you, %v!",
        "Hail, %v! Well met!",
    }

    // Return a randomly selected message format by specifying
    // a random index for the slice of formats.
    return formats[rand.Intn(len(formats))]
}

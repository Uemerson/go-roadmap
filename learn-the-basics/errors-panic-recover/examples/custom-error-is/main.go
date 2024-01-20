package main

import (
	"errors"
	"fmt"
)

type Error struct {
	Path string
	User string
}

func (e *Error) Error() string {
	return "error"
}

func (e *Error) Is(target error) bool {
	t, ok := target.(*Error)
	if !ok {
		return false
	}
	return (e.Path == t.Path || t.Path == "") &&
		(e.User == t.User || t.User == "")
}

func main() {
	err := &Error{Path: "", User: "someuser"}
	if errors.Is(err, &Error{User: "someuser"}) {
		fmt.Println("err's User field is 'someuser'")
	}
}

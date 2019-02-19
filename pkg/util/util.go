package util

import (
	"errors"
	"fmt"
)

//Object ...
type Object interface {
	DoSomething(string) (string, error)
}

//Pod ...
type Pod struct {
	name string
}

//DoSomething ...
func (p Pod) DoSomething(s string) (string, error) {
	fmt.Print("DoSomething ...." + s)
	return "", errors.New("problem")
}

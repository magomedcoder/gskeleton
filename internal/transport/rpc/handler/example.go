package handler

import "fmt"

type Example struct{}

func NewExampleHandler() *Example {
	return &Example{}
}

func (e *Example) Get() {
	fmt.Println("example.get")
}

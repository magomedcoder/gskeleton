package handler

import "context"

type Example struct{}

func NewExampleHandler() *Example {
	return &Example{}
}

type ExampleResponse struct {
	Text string `json:"text"`
}

func (e *Example) Get(ctx context.Context) (ExampleResponse, error) {
	return ExampleResponse{Text: "example.get"}, nil
}

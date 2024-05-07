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

type Request struct {
	Text string `json:"text"`
}

func (e *Example) Set(ctx context.Context, request *Request) (ExampleResponse, error) {
	return ExampleResponse{Text: "example.set"}, nil
}

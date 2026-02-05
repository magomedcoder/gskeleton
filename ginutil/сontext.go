package ginutil

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

type Context struct {
	Context *gin.Context
}

func NewContext(ctx *gin.Context) *Context {
	return &Context{ctx}
}

func (c *Context) Ctx() context.Context {
	return c.Context.Request.Context()
}

func (c *Context) Success(code int, data any, message ...string) error {
	resp := &Response{
		Message: "success",
		Data:    data,
	}
	if len(message) > 0 {
		resp.Message = message[0]
	}

	if value, ok := data.(proto.Message); ok {
		bt, _ := MarshalOptions.Marshal(value)
		var body map[string]any
		_ = json.Unmarshal(bt, &body)
		resp.Data = body
	}

	c.Context.AbortWithStatusJSON(code, resp)

	return nil
}

func (c *Context) InvalidParams(code int, message any) error {
	resp := &Response{
		Message: "Invalid parameters",
	}
	switch msg := message.(type) {
	case error:
		resp.Message = Translate(msg)
	case string:
		resp.Message = msg
	default:
		resp.Message = fmt.Sprintf("%v", msg)
	}

	c.Context.AbortWithStatusJSON(code, resp)

	return nil
}

func (c *Context) Error(code int, message any) error {
	resp := &Response{
		Message: "error",
	}
	switch msg := message.(type) {
	case error:
		resp.Message = msg.Error()
	case string:
		resp.Message = msg
	default:
		resp.Message = fmt.Sprintf("%v", msg)
	}

	c.Context.AbortWithStatusJSON(code, resp)

	return nil
}

var trans ut.Translator

var MarshalOptions = protojson.MarshalOptions{
	UseProtoNames:   true,
	EmitUnpopulated: true,
}

func Translate(err error) string {
	if errs, ok := err.(validator.ValidationErrors); ok {
		for _, err := range errs {
			return err.Translate(trans)
		}
	}

	return err.Error()
}

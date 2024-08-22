package http_server

import (
	"fmt"
	"github.com/magomedcoder/gskeleton/pkg/jsonutil"
	"net/http"

	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

var trans ut.Translator

var MarshalOptions = protojson.MarshalOptions{
	UseProtoNames:   true,
	EmitUnpopulated: true,
}

type Context struct {
	Context *gin.Context
}

func New(ctx *gin.Context) *Context {
	return &Context{ctx}
}

func (c *Context) Error(error string) error {
	c.Context.AbortWithStatusJSON(http.StatusInternalServerError, &Response{
		Message: error,
	})

	return nil
}

func (c *Context) Success(data any) error {
	resp := data

	if value, ok := data.(proto.Message); ok {
		bt, _ := MarshalOptions.Marshal(value)
		var data any
		if err := jsonutil.Decode(string(bt), &data); err != nil {
			return c.Error(err.Error())
		}
		resp = data
	}

	c.Context.AbortWithStatusJSON(http.StatusOK, resp)

	return nil
}

func Translate(err error) string {
	if errs, ok := err.(validator.ValidationErrors); ok {
		for _, err := range errs {
			return err.Translate(trans)
		}
	}

	return err.Error()
}

func (c *Context) InvalidParams(message any) error {
	resp := &Response{Message: "Недопустимые параметры"}

	switch msg := message.(type) {
	case error:
		resp.Message = Translate(msg)
	case string:
		resp.Message = msg
	default:
		resp.Message = fmt.Sprintf("%v", msg)
	}

	c.Context.AbortWithStatusJSON(http.StatusOK, resp)

	return nil
}

func (c *Context) ErrorBusiness(message any) error {
	resp := &Response{Message: "Недопустимые параметры"}

	switch msg := message.(type) {
	case error:
		resp.Message = msg.Error()
	case string:
		resp.Message = msg
	default:
		resp.Message = fmt.Sprintf("%v", msg)
	}

	c.Context.AbortWithStatusJSON(http.StatusBadRequest, resp)

	return nil
}

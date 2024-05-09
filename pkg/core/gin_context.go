package core

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

type GinContext struct {
	Context *gin.Context
}

func NewGinContext(ctx *gin.Context) *GinContext {
	return &GinContext{ctx}
}

func (g *GinContext) Error(error string) error {
	g.Context.AbortWithStatusJSON(http.StatusInternalServerError, &Response{
		Message: error,
	})

	return nil
}

func (g *GinContext) Success(data any) error {
	resp := data

	if value, ok := data.(proto.Message); ok {
		bt, _ := MarshalOptions.Marshal(value)
		var data any
		if err := jsonutil.Decode(string(bt), &data); err != nil {
			return g.Error(err.Error())
		}
		resp = data
	}

	g.Context.AbortWithStatusJSON(http.StatusOK, resp)

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

func (g *GinContext) InvalidParams(message any) error {
	resp := &Response{Message: "Недопустимые параметры"}

	switch msg := message.(type) {
	case error:
		resp.Message = Translate(msg)
	case string:
		resp.Message = msg
	default:
		resp.Message = fmt.Sprintf("%v", msg)
	}

	g.Context.AbortWithStatusJSON(http.StatusOK, resp)

	return nil
}

func (g *GinContext) ErrorBusiness(message any) error {
	resp := &Response{Message: "Недопустимые параметры"}

	switch msg := message.(type) {
	case error:
		resp.Message = msg.Error()
	case string:
		resp.Message = msg
	default:
		resp.Message = fmt.Sprintf("%v", msg)
	}

	g.Context.AbortWithStatusJSON(http.StatusBadRequest, resp)

	return nil
}

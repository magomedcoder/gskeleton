package jsonutil

import (
	"errors"

	"github.com/bytedance/sonic"
	jsoniter "github.com/json-iterator/go"
)

func Encode(value any) string {
	data, _ := sonic.MarshalString(value)
	return data
}

func Decode(data any, resp any) error {
	switch data.(type) {
	case string:
		return jsoniter.UnmarshalFromString(data.(string), resp)
	case []byte:
		return jsoniter.Unmarshal(data.([]byte), resp)
	default:
		return errors.New("неизвестный тип")
	}
}

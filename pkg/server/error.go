package server

import "fmt"

const (
	ErrCodeParseError     = -32700
	ErrCodeInvalidRequest = -32600
	ErrCodeMethodNotFound = -32601
	ErrCodeInvalidParams  = -32602
	ErrCodeInternalError  = -32603
	ErrUser               = -32000
)

var errors = map[int]string{
	ErrCodeParseError:     "Ошибка синтаксического анализа",
	ErrCodeInvalidRequest: "Недопустимый запрос",
	ErrCodeMethodNotFound: "Метод не найден",
	ErrCodeInvalidParams:  "Недопустимые параметры",
	ErrCodeInternalError:  "Внутренняя ошибка",
	ErrUser:               "Другая ошибка",
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e Error) Error() string {
	return fmt.Sprintf("Ошибка json-rpc: код: %d сообщение: %s", e.Code, e.Message)
}

func ErrorFromCode(code int) Error {
	if _, ok := errors[code]; ok {
		return Error{Code: code, Message: errors[code]}
	}

	return Error{Code: code}
}

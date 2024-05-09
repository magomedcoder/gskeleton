package v1

import (
	"github.com/magomedcoder/gskeleton/internal/usecase"
	"github.com/magomedcoder/gskeleton/pkg/http-server"
	"strconv"
)

type User struct {
	UserUseCase usecase.IUserUseCase
}

func NewUserHandler(
	userUseCase usecase.IUserUseCase,
) *User {
	return &User{
		UserUseCase: userUseCase,
	}
}

func (u *User) List(ctx *http_server.Context) error {

	// TODO

	return nil
}

type IGet struct {
	Id int `json:"id"`
}

func (u *User) Get(ctx *http_server.Context) error {
	id, err := strconv.Atoi(ctx.Context.Param("id"))
	if err != nil {
		return ctx.ErrorBusiness("Неверный id")
	}

	// TODO

	return ctx.Success(IGet{
		Id: id,
	})
}

package v1

import (
	"github.com/magomedcoder/gskeleton/internal/service"
	"github.com/magomedcoder/gskeleton/pkg/http-server"
	"strconv"
)

type User struct {
	UserService service.IUserService
}

func NewPostHandler(
	userService service.IUserService,
) *User {
	return &User{
		UserService: userService,
	}
}

func (u *User) List(ctx *http_server.Context) error {

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

	return ctx.Success(IGet{
		Id: id,
	})
}

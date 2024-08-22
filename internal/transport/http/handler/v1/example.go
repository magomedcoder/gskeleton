package v1

import (
	"github.com/magomedcoder/gskeleton/internal/service"
	"github.com/magomedcoder/gskeleton/pkg/http-server"
	"strconv"
)

type Example struct {
	UserService service.IUserService
}

func NewPostHandler(
	userService service.IUserService,
) *Example {
	return &Example{
		UserService: userService,
	}
}

type IGet struct {
	Id int `json:"id"`
}

func (p *Example) Get(ctx *http_server.Context) error {
	id, err := strconv.Atoi(ctx.Context.Param("id"))
	if err != nil {
		return ctx.ErrorBusiness("Неверный id")
	}

	return ctx.Success(IGet{
		Id: id,
	})
}

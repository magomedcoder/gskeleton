package v2

import (
	"github.com/magomedcoder/gskeleton/internal/usecase"
	"github.com/magomedcoder/gskeleton/pkg/core"
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

func (u *User) List(ctx *core.GinContext) error {

	return nil
}

type Get struct {
	Id int `json:"id"`
}

func (u *User) Get(ctx *core.GinContext) error {
	id, err := strconv.Atoi(ctx.Context.Param("id"))
	if err != nil {
		return ctx.ErrorBusiness("Неверный id")
	}

	user, _ := u.UserUseCase.GetUserById(ctx.Context, id)
	if user.Id == 0 {
		return ctx.ErrorBusiness("Пользователь не найден")
	}

	return ctx.Success(Get{
		Id: user.Id,
	})
}

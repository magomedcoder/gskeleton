package v1

import (
	"fmt"
	"github.com/magomedcoder/gskeleton/internal/domain/entity"
	"github.com/magomedcoder/gskeleton/internal/usecase"
	"github.com/magomedcoder/gskeleton/pkg/core"
	"github.com/magomedcoder/gskeleton/pkg/db/gormrepo"
	"gorm.io/gorm"
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

func (u *User) Create(ctx *core.GinContext) error {

	return ctx.Success(Get{})
}

type ListResponse struct {
	Total int64          `json:"total"`
	Items []*entity.User `json:"items"`
}

func (u *User) List(ctx *core.GinContext) error {
	var params entity.Pagination
	if err := ctx.Context.ShouldBindQuery(&params); err != nil {
		fmt.Println(err)
	}

	pagination := &gormrepo.Pagination{}
	pagination.SetPage(params.Page)
	pagination.SetPageSize(params.Limit)

	var count int64
	users, err := u.UserUseCase.GetUsers(ctx.Ctx(), func(db *gorm.DB) {
		db.Scopes(gormrepo.Paginate(pagination)).Count(&count)
	})
	if err != nil {
		return ctx.ErrorBusiness("Пользователи не найдены")
	}

	items := make([]*entity.User, 0)
	for _, item := range users {
		items = append(items, &entity.User{
			Id:   item.Id,
			Name: item.Name,
		})
	}

	return ctx.Success(ListResponse{
		Total: count,
		Items: items,
	})
}

type Get struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
}

func (u *User) Get(ctx *core.GinContext) error {
	id, err := strconv.ParseInt(ctx.Context.Param("id"), 10, 64)
	if err != nil {
		return ctx.ErrorBusiness("Неверный id")
	}

	user, _ := u.UserUseCase.GetUserById(ctx.Ctx(), id)
	if user.Id == 0 {
		return ctx.ErrorBusiness("Пользователь не найден")
	}

	return ctx.Success(Get{
		Id:       user.Id,
		Username: user.Username,
	})
}

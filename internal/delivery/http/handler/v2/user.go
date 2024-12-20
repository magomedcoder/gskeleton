package v2

import (
	"fmt"
	"github.com/magomedcoder/gskeleton/internal/domain/entity"
	"github.com/magomedcoder/gskeleton/internal/usecase"
	"github.com/magomedcoder/gskeleton/pkg/ginutil"
	"github.com/magomedcoder/gskeleton/pkg/gormutil"
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

func (u *User) Create(ctx *ginutil.Context) error {
	return ctx.Success(Get{})
}

type ListResponse struct {
	Total int64          `json:"total"`
	Items []*entity.User `json:"items"`
}

func (u *User) List(ctx *ginutil.Context) error {
	var params entity.Pagination
	if err := ctx.Context.ShouldBindQuery(&params); err != nil {
		fmt.Println(err)
	}

	var count int64
	users, err := u.UserUseCase.GetUsers(ctx.Ctx(), func(db *gorm.DB) {
		db.Scopes(gormutil.Paginate(&gormutil.Pagination{
			Page:     params.Page,
			PageSize: params.Limit,
		})).Count(&count)
	})
	if err != nil {
		return ctx.Error("Пользователи не найдены")
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
	Id int64 `json:"id"`
}

func (u *User) Get(ctx *ginutil.Context) error {
	id, err := strconv.ParseInt(ctx.Context.Param("id"), 10, 64)
	if err != nil {
		return ctx.Error("Неверный id")
	}

	user, _ := u.UserUseCase.GetUserById(ctx.Ctx(), id)
	if user.Id == 0 {
		return ctx.Error("Пользователь не найден")
	}

	return ctx.Success(Get{
		Id: user.Id,
	})
}

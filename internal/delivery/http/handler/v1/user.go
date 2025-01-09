package v1

import (
	v1Pb "github.com/magomedcoder/gskeleton/api/http/pb/v1"
	postgresModel "github.com/magomedcoder/gskeleton/internal/infrastructure/postgres/model"
	"github.com/magomedcoder/gskeleton/internal/usecase"
	"github.com/magomedcoder/gskeleton/pkg/ginutil"
	"github.com/magomedcoder/gskeleton/pkg/gormutil"
	"gorm.io/gorm"
	"strconv"
	"time"
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
	params := &v1Pb.CreateUserRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(err)
	}

	passwordHash, err := u.UserUseCase.HashPassword(params.Password)
	if err != nil {
		return ctx.Error("Не удалось хешировать пароль")
	}

	user := postgresModel.User{
		Username:  params.Username,
		Password:  passwordHash,
		CreatedAt: time.Now(),
	}

	if _, err = u.UserUseCase.Create(ctx.Ctx(), &user); err != nil {
		return ctx.Error(err.Error())
	}

	return ctx.Success(&v1Pb.GetUserResponse{
		Id:       user.Id,
		Username: user.Username,
		Name:     user.Name,
	})
}

func (u *User) List(ctx *ginutil.Context) error {
	page, err := strconv.Atoi(ctx.Context.DefaultQuery("page", "1"))
	if err != nil {
		return ctx.Error("Неверный номер страницы")
	}

	pageSize, err := strconv.Atoi(ctx.Context.DefaultQuery("pageSize", "15"))
	if err != nil {
		return ctx.Error("Неверный размер страницы")
	}

	pagination := &gormutil.Pagination{}
	pagination.SetPage(page)
	pagination.SetPageSize(pageSize)

	var count int64
	users, err := u.UserUseCase.GetUsers(ctx.Ctx(), func(db *gorm.DB) {
		db.Scopes(gormutil.Paginate(pagination)).Count(&count)
	})
	if err != nil {
		return ctx.Error("Пользователи не найдены")
	}

	items := make([]*v1Pb.UserItem, 0)
	for _, item := range users {
		items = append(items, &v1Pb.UserItem{
			Id:       item.Id,
			Username: item.Username,
			Name:     item.Name,
		})
	}

	return ctx.Success(&v1Pb.GetUsersResponse{
		Total: count,
		Items: items,
	})
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

	return ctx.Success(&v1Pb.GetUserResponse{
		Id:       user.Id,
		Username: user.Username,
		Name:     user.Name,
	})
}

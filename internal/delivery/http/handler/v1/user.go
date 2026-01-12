package v1

import (
	v1Pb "github.com/magomedcoder/gskeleton/api/http/pb/v1"
	"github.com/magomedcoder/gskeleton/internal/domain/entity"
	"github.com/magomedcoder/gskeleton/internal/usecase"
	"github.com/magomedcoder/gskeleton/pkg/ginutil"
	"github.com/magomedcoder/gskeleton/pkg/gormutil"
	"gorm.io/gorm"
	"log"
	"net/http"
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

func (u *User) CreateUser(ctx *ginutil.Context) error {
	params := &v1Pb.CreateUserRequest{}
	if err := ctx.Context.ShouldBind(params); err != nil {
		return ctx.InvalidParams(http.StatusBadRequest, err)
	}

	user := &entity.UserOpt{
		Username: params.Username,
		Password: params.Password,
	}

	userModel, err := u.UserUseCase.Create(ctx.Ctx(), user)
	if err != nil {
		log.Printf("Ошибка создания пользователя: %v", err)
		return ctx.Error(http.StatusBadRequest, "Ошибка создания пользователя")
	}

	return ctx.Success(http.StatusOK, &v1Pb.CreateUserResponse{
		Id: userModel.Id,
	})
}

func (u *User) GetUsers(ctx *ginutil.Context) error {
	page, err := strconv.Atoi(ctx.Context.DefaultQuery("page", "1"))
	if err != nil {
		return ctx.Error(http.StatusBadRequest, "Неверный номер страницы")
	}

	pageSize, err := strconv.Atoi(ctx.Context.DefaultQuery("pageSize", "15"))
	if err != nil {
		return ctx.Error(http.StatusBadRequest, "Неверный размер страницы")
	}

	pagination := &gormutil.Pagination{}
	pagination.SetPage(page)
	pagination.SetPageSize(pageSize)

	var count int64
	users, err := u.UserUseCase.GetUsers(ctx.Ctx(), func(db *gorm.DB) {
		db.Scopes(gormutil.Paginate(pagination)).Count(&count)
	})
	if err != nil {
		return ctx.Error(http.StatusBadRequest, "Пользователи не найдены")
	}

	items := make([]*v1Pb.User, 0)
	for _, item := range users {
		items = append(items, &v1Pb.User{
			Id:       item.Id,
			Username: item.Username,
			Name:     item.Name,
		})
	}

	return ctx.Success(http.StatusOK, &v1Pb.GetUsersResponse{
		Total: count,
		Items: items,
	})
}

func (u *User) GetUser(ctx *ginutil.Context) error {
	id, err := strconv.ParseInt(ctx.Context.Param("id"), 10, 64)
	if err != nil {
		return ctx.Error(http.StatusBadRequest, "Неверный id")
	}

	user, _ := u.UserUseCase.GetUserById(ctx.Ctx(), id)
	if user.Id == 0 {
		return ctx.Error(http.StatusBadRequest, "Пользователь не найден")
	}

	return ctx.Success(http.StatusOK, &v1Pb.GetUserResponse{
		Id:       user.Id,
		Username: user.Username,
		Name:     user.Name,
	})
}

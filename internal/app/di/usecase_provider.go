package di

import "github.com/magomedcoder/gskeleton/internal/usecase"

type UseCaseProvider struct {
	UserUseCase *usecase.UserUseCase
}

func NewUseCaseProvider(infra *InfrastructureProvider) *UseCaseProvider {
	userUseCase := &usecase.UserUseCase{
		UserRepo:            infra.UserRepository,
		UserCacheRepository: infra.UserCacheRepository,
		UserLogRepository:   infra.UserLogRepository,
	}

	return &UseCaseProvider{
		UserUseCase: userUseCase,
	}
}

package di

import (
	"github.com/magomedcoder/gskeleton/internal/cli/handler"
	"github.com/magomedcoder/gskeleton/internal/config"
)

type CLIProvider struct {
	Conf    *config.Config
	Migrate *handler.Migrate
}

func NewCliInjector(conf *config.Config) *CLIProvider {
	provider := NewProvider(conf)
	infra := NewInfrastructureProvider(provider)
	useCases := NewUseCaseProvider(infra)

	migrate := &handler.Migrate{
		Conf:        conf,
		UserUseCase: useCases.UserUseCase,
	}

	return &CLIProvider{
		Conf:    conf,
		Migrate: migrate,
	}
}

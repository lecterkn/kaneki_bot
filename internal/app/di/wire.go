//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/lecterkn/kaneki_bot/internal/app/client"
	"github.com/lecterkn/kaneki_bot/internal/app/handler"
	"github.com/lecterkn/kaneki_bot/internal/app/repository"
	"github.com/lecterkn/kaneki_bot/internal/app/usecase"
)

var clientSet = wire.NewSet(
	client.GetGeminiClient,
)

var repositorySet = wire.NewSet(
	repository.NewGenerateRepositoryImpl,
)

var usecaseSet = wire.NewSet(
	usecase.NewGenerateUsecase,
)

var handlerSet = wire.NewSet(
	handler.NewMessageHandler,
)

type HandlerSets struct {
	MessageHandler *handler.MessageHandler
}

func InitializeHandlers() *HandlerSets {
	wire.Build(
		clientSet,
		repositorySet,
		usecaseSet,
		handlerSet,
		wire.Struct(new(HandlerSets), "*"),
	)
	return nil
}

//go:build wireinject
// +build wireinject

package di

import (
	"Bot-or-Not/internal/app/handler"
	"Bot-or-Not/internal/app/router"
	"Bot-or-Not/internal/app/service"
	"Bot-or-Not/internal/infra/database"
	"Bot-or-Not/internal/infra/repository"

	"github.com/google/wire"
)

func New() *router.Root {
	wire.Build(
		database.New,
		repository.NewGameRepository,
		service.NewGameService,
		handler.NewGameHandler,
		router.New,
	)
	return &router.Root{}
}

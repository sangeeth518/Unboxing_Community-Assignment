//go:build wireinject
// +build wireinject

package di

import (
	http "Assignment/pkg/api"
	"Assignment/pkg/api/handler"
	"Assignment/pkg/config"
	"Assignment/pkg/db"
	"Assignment/pkg/repository"
	"Assignment/pkg/usecase"

	"github.com/google/wire"
)

func InitializeAPI(cfg config.Config) (*http.ServerHttp, error) {
	wire.Build(db.ConnectDB, repository.NewAdminRepository, usecase.NewAdminUsecase, handler.NewAdminHandler, http.NewServerHttp)
	return &http.ServerHttp{}, nil
}

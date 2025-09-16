//go:build wireinject
// +build wireinject

package clean_architecture

import (
	"clean_architecture/internal/config"
	"clean_architecture/internal/handler"
	"clean_architecture/internal/repository"
	"clean_architecture/internal/service"
	"clean_architecture/pkg/database"

	"github.com/google/wire"
	"gorm.io/gorm"
)

// Provider Sets
var ConfigSet = wire.NewSet(
	config.Load,
)

var DatabaseSet = wire.NewSet(
	database.NewPostgresDB,
)

var RepositorySet = wire.NewSet(
	repository.NewRepositories,
)

var ServiceSet = wire.NewSet(
	service.NewServices,
)

var HandlerSet = wire.NewSet(
	handler.NewHandlers,
)

// App structure
type App struct {
	DB       *gorm.DB
	Handlers *handler.Handlers
	Config   *config.Config
}

// Wire injector
func InitializeApp() (*App, func(), error) {
	wire.Build(
		ConfigSet,
		DatabaseSet,
		RepositorySet,
		ServiceSet,
		HandlerSet,
		wire.Struct(new(App), "*"),
	)
	return nil, nil, nil
}

// Cleanup provider
func ProvideCleanup(db *gorm.DB) func() {
	return func() {
		sqlDB, err := db.DB()
		if err == nil {
			sqlDB.Close()
		}
	}
}

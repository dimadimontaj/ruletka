package cmd

import (
	"context"
	"log"

	"database/sql"
	"cases/migrations"
	"cases/pkg/logger"
)

// контейнер внешних зависимостей приложения
// тут мы инициализируем все инфраструктурные зависимости
type Container struct {
	gCtx             context.Context
	configuration    *configuration
	db               *sql.DB
	migrator         *migrations.Migrator
	logger           *logger.Logger
}

func NewContainer() *Container {
	return &Container{
		configuration: newFromEnv(),
	}
}

func (e *Container) GetConfiguration() *configuration {
	return e.configuration
}

func (e *Container) GetGlobalContext() context.Context {
	if e.gCtx == nil {
		e.gCtx = context.Background()
	}

	return e.gCtx
}

func (e *Container) GetPostgres() *sql.DB {
	if e.db == nil {
		var err error
		e.db, err = NewSqlConn(e.configuration.GetPostgresConfiguration())
		if err != nil {
			log.Fatal(err)
		}
	}

	return e.db
}

func (e *Container) GetLogger() *logger.Logger {
	if e.logger == nil {
		e.logger = logger.New()
	}

	return e.logger
}

func (e *Container) GetMigrator() *migrations.Migrator {
	if e.migrator == nil {
		e.migrator = migrations.NewMigrator(
			e.configuration.
				GetPostgresConfiguration().
				GetMigrateConnectionString(),
		)
	}

	return e.migrator
}
package cmd

import (
	"github.com/timurzdev/mentorship-test-task/internal/handler/server"
	"github.com/timurzdev/mentorship-test-task/internal/repository"
)

// контейнер внутренних зависимостей
type Internal struct {
	//external
	*Container

	repository *repository.Repository

	server *server.Server

	//handlers

	//usecases

	//middlewares
}

func NewInternal(container *Container) *Internal {
	return &Internal{Container: container}
}

func (i *Internal) GetRepository() *repository.Repository {
	if i.repository == nil {
		i.repository = repository.NewRepository(i.GetPostgres())
	}

	return i.repository
}

func (i *Internal) GetServer() *server.Server {
	if i.server == nil {
		i.server = server.NewServer(
			i.GetLogger(),
			i.configuration.GetServerConfiguration().GetAddress(),
		)
	}

	return i.server
}
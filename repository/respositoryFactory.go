package repository

import (
	"glints-part2/config"
)

type RepositoryFactory struct {
	Repository *Repository
}

func (factory *RepositoryFactory) GetRepository() *Repository {
	return factory.Repository
}

func NewRepositoryFactory() *RepositoryFactory {
	conf := config.GetDBConfig()
	db := newDBOP(conf.UserName, conf.Password, conf.Host, conf.Port, conf.Name)

	return &RepositoryFactory{
		Repository: newRepository(db),
	}
}

package repository

import (
	"glints-part2/config"

	log "github.com/sirupsen/logrus"
)

type RepositoryFactory struct {
	Repository *Repository
}

func (factory *RepositoryFactory) GetRepository() *Repository {
	return factory.Repository
}

func NewRepositoryFactory() *RepositoryFactory {
	conf := config.GetDBConfig()
	log.Debug("Config:", conf.UserName, " ", conf.Password, " ", conf.Host, " ", conf.Port, " ", conf.Name)
	db := newDBOP(conf.UserName, conf.Password, conf.Host, conf.Port, conf.Name)

	return &RepositoryFactory{
		Repository: newRepository(db),
	}
}

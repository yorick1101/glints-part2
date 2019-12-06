package etl

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"

	"glints-part2/model"
	"glints-part2/repository"
)

type ETLService interface {
	Process(filePath string)
}

type ETLServiceImpl struct {
	repo *repository.Repository
}

func (srv ETLServiceImpl) Process(filePath string) {
	rows := extract(filePath)
	srv.repo.ImportComanies(rows)

}

func extract(filePath string) []model.Company {
	absPath, _ := filepath.Abs(filePath)
	log.Info(absPath)
	file, err := os.Open(absPath)
	if err != nil {
		log.Panic("cannot open josn file", err)
	}
	defer file.Close()

	byteValue, _ := ioutil.ReadAll(file)

	var data []model.Company
	err = json.Unmarshal(byteValue, &data)
	if err != nil {
		log.Panic("cannot marshal from file", err)
	}

	return data
}

func NewService() *ETLServiceImpl {

	factory := repository.NewRepositoryFactory()

	return &ETLServiceImpl{
		repo: factory.GetRepository(),
	}
}

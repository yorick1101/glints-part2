package main

import (
	"encoding/json"
	"fmt"
	"glints-part2/config"
	"glints-part2/model"
	"glints-part2/repository"
	"time"

	log "github.com/sirupsen/logrus"
)

func main() {
	config.Init()
	repo := repository.NewRepositoryFactory().GetRepository()
	//testFindComany(repo)
	//testAmountFundingRound(repo)
	//testFundingRate(repo)
	//testFindByPersonId(repo)
	testFindByCompanyId(repo)
}

func testFundingRate(repo *repository.Repository) {
	var filters []model.DateFilter
	var startDat = time.Date(2007, time.Month(1), 1, 0, 0, 0, 0, time.UTC)
	//var startDat = time.Date(2020, time.Month(1), 1, 0, 0, 0, 0, time.UTC)
	filters = append(filters, model.DateFilter{"gt", startDat})
	//filters = append(filters, model.DateFilter{"lt", 20})
	results := repo.FilterByFundingDate(filters)

	for _, result := range results {
		fmt.Println(result)
	}
}

func testFindComany(repo *repository.Repository) {

	companies := repo.FindCompanyByIds("52cdef7d4bab8bd675299fe9", "52cdef7f4bab8bd67529bfee")
	for _, company := range companies {

		data, err := json.Marshal(company)
		if err != nil {
			log.Panic(err)
		} else {
			log.Info(string(data))
		}

	}

}

func testCountFundingRound(repo *repository.Repository) {
	var filters []model.IntFilter
	filters = append(filters, model.IntFilter{"gt", 11})
	filters = append(filters, model.IntFilter{"lt", 20})
	results := repo.FilterByNumberOfFundingRounds(filters)

	for _, result := range results {
		fmt.Println(result)
	}
}

func testAmountFundingRound(repo *repository.Repository) {
	var filters []model.IntFilter
	filters = append(filters, model.IntFilter{"eq", 121000000})
	//filters = append(filters, model.IntFilter{"lt", 20000})
	results := repo.FilterByAmountOfFundingRounds(filters)

	for _, result := range results {
		fmt.Println(result)
	}
}

func testFindByPersonId(repo *repository.Repository) {
	//for relationship
	//result := repo.FindPersonOnRelationship(true, "5dec8fe306c3720e7fe72e09")

	//for fundingrouonds
	result := repo.FindPersonOnFundingRounds("5dec8fe306c3720e7fe7435b")
	data, err := json.Marshal(result)
	if err != nil {
		log.Panic(err)
	} else {
		log.Info(string(data))
	}
}

func testFindByCompanyId(repo *repository.Repository) {
	result := repo.FindCompanyOnAcquisitions("52cdef7e4bab8bd67529b5b9")
	data, err := json.Marshal(result)
	if err != nil {
		log.Panic(err)
	} else {
		log.Info(string(data))
	}
}

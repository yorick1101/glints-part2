package main

import (
	"fmt"
	"glints-part2/config"
	"glints-part2/model"
	"glints-part2/repository"
)

func main() {
	config.Init()
	repo := repository.NewRepositoryFactory().GetRepository()
	testFindComany(repo)
}

func testFindComany(repo *repository.Repository) {

	repo.FindCompanyById("52cdef7d4bab8bd675299fe9", "52cdef7f4bab8bd67529bfee")
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

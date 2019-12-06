package repository

import (
	"fmt"
	"glints-part2/model"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	db DBOP
}

func (repo *Repository) ImportComanies(companies []model.Company) {

	companyCollection := repo.db.GetCollection("companies")
	personCollection := repo.db.GetCollection("persons")

	personMap := make(map[string]interface{})
	personIdMap := make(map[string]string)

	var schemas []*model.BsonCompany
	companyPermalinkToId := make(map[string]string)
	for _, srcCompany := range companies {
		companyPermalinkToId[srcCompany.Permalink] = srcCompany.Id.Id
		company := srcCompany.ToSchema()
		schemas = append(schemas, company)
		for _, relationship := range company.Relationships {
			person := relationship.Person
			personMap[person.Permalink] = person
		}
		for _, funding := range company.FundingRounds {
			for _, investment := range funding.Investments {
				if investment.Person != nil {
					person := investment.Person
					personMap[person.Permalink] = person
				}
			}
		}
	}

	var persons []interface{}
	var personPermaLinks []string

	for personPermaLink, person := range personMap {
		personPermaLinks = append(personPermaLinks, personPermaLink)
		persons = append(persons, person)
	}

	personids := personCollection.InsertMany(persons)
	for index, personId := range personids {
		personIdMap[personPermaLinks[index]] = personId.(primitive.ObjectID).Hex()
	}

	var icompanies []interface{}
	for _, company := range schemas {

		for _, relationship := range company.Relationships {
			person := relationship.Person
			relationship.PersonId = personIdMap[person.Permalink]
			//log.Info("relationship.PersonId: ", relationship.PersonId)
		}
		for _, funding := range company.FundingRounds {
			for _, investment := range funding.Investments {
				if investment.Person != nil {
					if value, ok := personIdMap[investment.Person.Permalink]; ok {
						investment.PersonId = value
					} else {
						//log.Warn("No person id found for ", investment.Person.Permalink)
					}
				}
				if investment.Company != nil {
					if value, ok := companyPermalinkToId[investment.Company.Permalink]; ok {
						investment.CompanyId = value
					} else {
						//log.Warn("No company id found for ", investment.Company.Permalink)
					}
				}
			}
		}
		{
			value := companyPermalinkToId[company.Acquisition.AcquiringCompany.Permalink]
			company.Acquisition.AcquiringCompanyId = value
		}

		for _, acqusition := range company.Acquisitions {
			if value, ok := companyPermalinkToId[acqusition.AcquiringCompany.Permalink]; ok {
				acqusition.AcquiringCompanyId = value
			}
		}

		for _, competition := range company.Competitions {
			if value, ok := companyPermalinkToId[competition.Competitor.Permalink]; ok {
				competition.CompetitorId = value
			}
		}

		icompanies = append(icompanies, company)
	}

	ids := companyCollection.InsertMany(icompanies)
	log.Info("inserted company:", len(ids))
}

func (repo *Repository) FilterByNumberOfFundingRounds(criterias []model.IntFilter) []string {
	companyCollection := repo.db.GetCollection("companies")

	var criteriaD bson.D
	for _, criteria := range criterias {
		var op = "$" + criteria.Operation
		criteriaD = append(criteriaD, bson.E{op, criteria.Value})
	}

	pipeline := mongo.Pipeline{
		{{"$project", bson.D{
			{"_id", 1},
			{"count", bson.D{{"$size", bson.D{{"$ifNull", []interface{}{"$fundingRounds", []interface{}{}}}}}}},
		}}},

		{{"$match", bson.D{{"count", criteriaD}}}},
	}
	return companyCollection.FindIdByAggregate(pipeline)
}

func (repo *Repository) FindCompanyById(ids ...string) []interface{} {
	companyCollection := repo.db.GetCollection("companies")
	var criteriaD bson.D
	criteriaD = append(criteriaD, bson.E{"_id", bson.D{{"$in", ids}}})
	fmt.Println(criteriaD)
	return companyCollection.Find(criteriaD)
}

func newRepository(db DBOP) *Repository {
	return &Repository{
		db: db,
	}
}

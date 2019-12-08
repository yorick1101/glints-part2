package repository

import (
	"glints-part2/model"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct {
	db DBOP
}

type IdContainer struct {
	Id string `bson:"_id"`
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

func (repo *Repository) FilterByNumberOfFundingRounds(criterias []model.IntFilter) ([]string, error) {
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

func (repo *Repository) FilterByAmountOfFundingRounds(criterias []model.IntFilter) ([]string, error) {
	companyCollection := repo.db.GetCollection("companies")

	var criteriaD bson.D
	for _, criteria := range criterias {
		var op = "$" + criteria.Operation
		criteriaD = append(criteriaD, bson.E{op, criteria.Value})
	}

	pipeline := mongo.Pipeline{
		{{"$project", bson.D{
			{"_id", 1},
			{"count", bson.D{{"$sum", "$fundingRounds.raisedAmount"}}},
		}}},

		{{"$match", bson.D{{"count", criteriaD}}}},
	}
	return companyCollection.FindIdByAggregate(pipeline)
}

func filterByDate(repo *Repository, field string, criterias []model.DateFilter) []string {
	companyCollection := repo.db.GetCollection("companies")

	var criteriaD bson.D
	for _, criteria := range criterias {
		var op = "$" + criteria.Operation
		criteriaD = append(criteriaD, bson.E{op, criteria.Value})
	}

	criteria := bson.D{{"foundedDate", criteriaD}}

	projection := bson.D{{"_id", 1}}
	opt := options.Find().SetProjection(projection)

	var companyIds []IdContainer
	var results []string
	companyCollection.FindWithOptions(criteria, &companyIds, opt)
	for _, companyId := range companyIds {
		results = append(results, companyId.Id)
	}
	return results
}

func (repo *Repository) FilterByFundingDate(criterias []model.DateFilter) []string {
	return filterByDate(repo, "foundedDate", criterias)
}

func (repo *Repository) FilterByDeadedpoolDate(criterias []model.DateFilter) []string {
	return filterByDate(repo, "deadpooledDate", criterias)
}

func (repo *Repository) FindPersonOnRelationship(ispast bool, personId string) []string {
	collection := repo.db.GetCollection("companies")

	criteria := bson.D{{"relationships", bson.D{{"$elemMatch", bson.D{{"personId", personId}, {"isPast", ispast}}}}}}
	var projection bson.D
	projection = append(projection, bson.E{"_id", 1})
	//projection = append(projection, bson.E{"relationships.$", 1})
	opt := options.Find().SetProjection(projection)

	var containers []IdContainer
	collection.FindWithOptions(criteria, &containers, opt)

	var ids []string
	for _, result := range containers {
		ids = append(ids, result.Id)
	}

	return ids
}

func (repo *Repository) FindPersonOnFundingRounds(personId string) []string {
	collection := repo.db.GetCollection("companies")
	criteria := bson.D{{"fundingRounds.investments.personId", personId}}

	projection := bson.D{{"_id", 1}, {"fundingRounds.$", 1}}
	opt := options.Find().SetProjection(projection)

	var containers []IdContainer
	collection.FindWithOptions(criteria, &containers, opt)

	var ids []string
	for _, result := range containers {
		ids = append(ids, result.Id)
	}
	return ids
}

func (repo *Repository) FindCompanyOnFundingRounds(compaynId string) []string {
	collection := repo.db.GetCollection("companies")
	criteria := bson.D{{"fundingRounds.investments.companyId", compaynId}}

	//projection := bson.D{{"_id", 1}, {"fundingRounds.$", 1}}
	projection := bson.D{{"_id", 1}}
	opt := options.Find().SetProjection(projection)

	var containers []IdContainer
	collection.FindWithOptions(criteria, &containers, opt)

	var ids []string
	for _, result := range containers {
		ids = append(ids, result.Id)
	}
	return ids
}

func (repo *Repository) FindCompanyOnAcquisitions(compaynId string) []string {
	collection := repo.db.GetCollection("companies")
	criteria := bson.D{{"$or", []interface{}{
		bson.D{{"acquisitions.acquiringCompanyId", compaynId}},
		bson.D{{"acquisition.acquiringCompanyId", compaynId}},
	}}}

	projection := bson.D{{"_id", 1}}
	opt := options.Find().SetProjection(projection)

	var containers []IdContainer
	collection.FindWithOptions(criteria, &containers, opt)

	var ids []string
	for _, result := range containers {
		ids = append(ids, result.Id)
	}
	return ids
}

func (repo *Repository) FindCompanyByIds(ids []string) ([]model.BsonCompany, error) {
	companyCollection := repo.db.GetCollection("companies")
	var criteriaD bson.D
	criteriaD = append(criteriaD, bson.E{"_id", bson.D{{"$in", ids}}})

	var companies []model.BsonCompany
	err := companyCollection.Find(criteriaD, &companies)
	if err != nil {
		return companies, err
	}
	return companies, nil
}

func (repo *Repository) ReplaceCompany(company *model.BsonCompany) (int64, error) {
	companyCollection := repo.db.GetCollection("companies")
	return companyCollection.Replace(company.Id, company)
}

func newRepository(db DBOP) *Repository {
	return &Repository{
		db: db,
	}
}

package model

import (
	"time"
)

type BsonCompany struct {
	Id                string              `bson:"_id"`
	Name              string              `bson:"name"`
	Permalink         string              `bson:"permalink"`
	CrunchbaseUrl     string              `bson:"crunchbaseUrl"`
	HomepageUrl       string              `bson:"homepageUrl"`
	CategoryCode      string              `bson:"categoryCode"`
	NumberOfEmployees int                 `bson:"numberOfEmployees"`
	FoundedDate       *time.Time          `bson:"foundedDate"`
	DeadpooledDate    *time.Time          `bson:"deadpooledDate"`
	TagList           string              `bson:"tagList"`
	EmailAddress      string              `bson:"emailAddress"`
	Overview          string              `bson:"overview"`
	TotalMoneyRaised  string              `bson:"totalMoneyRaised"`
	Acquisition       *BsonAcquisition    `bson:"acquisition"`
	Acquisitions      []BsonAcquisition   `bson:"acquisitions"`
	Relationships     []*BsonRelationship `bson:"relationships"`
	FundingRounds     []*BsonFundingRound `bson:"fundingRounds"`
	Competitions      []*BsonCompetitor   `bson:"competitions"`
}

type BsonAcquisition struct {
	PriceAmount        float64       `bson:"priceAmount"`
	PriceCurrencyCode  string        `bson:"priceCurrencyCode"`
	TermCode           string        `bson:"termCode"`
	SourceUrl          string        `bson:"sourceUrl"`
	SourceDescription  string        `bson:"sourceDescription"`
	AcquiredDate       time.Time     `bson:"acquiredDate"`
	AcquiringCompany   PermanentLink `bson:"_"`
	AcquiringCompanyId string        `bson:"acquiringCompanyId"`
}

type BsonFundingRound struct {
	RoundCode          string            `bson:"roundCode"`
	SourceUrl          string            `bson:"sourceUrl"`
	SourceDescription  string            `bson:"sourceDescription"`
	RaisedAmount       float64           `bson:"raisedAmount"`
	RaisedCurrencyCode string            `bson:"raisedCurrencyCode"`
	FundedDate         *time.Time        `bson:"fundedDate"`
	Investments        []*BsonInvestment `bson:"investments"`
}

type BsonRelationship struct {
	IsPast   bool       `bson:"isPast"`
	Title    string     `bson:"title"`
	Person   BsonPerson `bson:"-"`
	PersonId string     `bson:"personId"`
}

type BsonPerson struct {
	FirstName string `bson:"firstName"`
	LastName  string `bson:"lastName"`
	Permalink string `bson:"permalink"`
}

type BsonInvestment struct {
	Company      *PermanentLink `bson:"-"`
	CompanyId    string         `bson:"companyId"`
	FinancialOrg *PermanentLink `bson:"financialOrg"`
	Person       *BsonPerson    `bson:"_"`
	PersonId     string         `bson:"personId"`
}

type BsonCompetitor struct {
	Competitor   PermanentLink `bson:"-"`
	CompetitorId string        `bson:"companyId"`
}

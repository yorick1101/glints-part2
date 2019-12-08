package model

import (
	"time"
)

type BsonCompany struct {
	Id                string              `bson:"_id" json:"id"`
	Name              string              `bson:"name" json:"name"`
	Permalink         string              `bson:"permalink" json:"permalink"`
	CrunchbaseUrl     string              `bson:"crunchbaseUrl" json:"crunchbase_url"`
	HomepageUrl       string              `bson:"homepageUrl" json:"homepage_url"`
	CategoryCode      string              `bson:"categoryCode" json:"category_code"`
	NumberOfEmployees int                 `bson:"numberOfEmployees" json:"number_of_employees"`
	FoundedDate       *time.Time          `bson:"foundedDate" json:"-"`
	DeadpooledDate    *time.Time          `bson:"deadpooledDate" json:"-"`
	TagList           string              `bson:"tagList" json:"tag_list"`
	EmailAddress      string              `bson:"emailAddress" json:"email_address"`
	Overview          string              `bson:"overview" json:"overview"`
	TotalMoneyRaised  string              `bson:"totalMoneyRaised" json:"total_money_raised"`
	Acquisition       *BsonAcquisition    `bson:"acquisition" json:"acquisition"`
	Acquisitions      []BsonAcquisition   `bson:"acquisitions" json:"acquisitions"`
	Relationships     []*BsonRelationship `bson:"relationships" json:"relationships"`
	FundingRounds     []*BsonFundingRound `bson:"fundingRounds" json:"funding_rounds"`
	Competitions      []*BsonCompetitor   `bson:"competitions" json:"competitions"`
}

type BsonAcquisition struct {
	PriceAmount        float64       `bson:"priceAmount" json:"price_amount"`
	PriceCurrencyCode  string        `bson:"priceCurrencyCode" json:"price_currency_code"`
	TermCode           string        `bson:"termCode" json:"term_code"`
	SourceUrl          string        `bson:"sourceUrl" json:"source_url"`
	SourceDescription  string        `bson:"sourceDescription" json:"source_description"`
	AcquiredDate       *time.Time    `bson:"acquiredDate" json:"-"`
	AcquiringCompany   PermanentLink `bson:"-" json:"-"`
	AcquiringCompanyId string        `bson:"acquiringCompanyId" json:"acquiring_company_id"`
}

type BsonFundingRound struct {
	RoundCode          string            `bson:"roundCode" json:"round_code"`
	SourceUrl          string            `bson:"sourceUrl" json:"source_url"`
	SourceDescription  string            `bson:"sourceDescription" json:"source_description"`
	RaisedAmount       float64           `bson:"raisedAmount" json:"raised_amount"`
	RaisedCurrencyCode string            `bson:"raisedCurrencyCode" json:"raised_currency_code"`
	FundedDate         *time.Time        `bson:"fundedDate" json:"-"`
	Investments        []*BsonInvestment `bson:"investments" json:"investments"`
}

type BsonRelationship struct {
	IsPast   bool       `bson:"isPast" json:"is_past"`
	Title    string     `bson:"title" json:"title"`
	Person   BsonPerson `bson:"-" json:"-"`
	PersonId string     `bson:"personId" json:"person_id"`
}

type BsonPerson struct {
	FirstName string `bson:"firstName" json:"first_name"`
	LastName  string `bson:"lastName" json:"last_name"`
	Permalink string `bson:"permalink" json:"permalink"`
}

type BsonInvestment struct {
	Company      *PermanentLink `bson:"-" json:"-"`
	CompanyId    string         `bson:"companyId" json:"company_id"`
	FinancialOrg *PermanentLink `bson:"financialOrg" json:"financial_org"`
	Person       *BsonPerson    `bson:"-" json:"-"`
	PersonId     string         `bson:"personId" json:"persion_id"`
}

type BsonCompetitor struct {
	Competitor   PermanentLink `bson:"-" json:"-"`
	CompetitorId string        `bson:"companyId" json:"company_id"`
}

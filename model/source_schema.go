package model

import "time"

type OID struct {
	Id string `json:"$oid"`
}

type Company struct {
	Id                OID            `json:"_id"`
	Name              string         `json:"name"`
	Permalink         string         `json:"permalink"`
	CrunchbaseUrl     string         `json:"crunchbase_url"`
	HomepageUrl       string         `json:"homepage_url"`
	CategoryCode      string         `json:"category_code"`
	NumberOfEmployees int            `json:"number_of_employees"`
	FoundedYear       int            `json:"founded_year"`
	FoundedMonth      int            `json:"founded_month"`
	FoundedDay        int            `json:"founded_day"`
	DeadpooledYear    int            `json:"deadpooled_year"`
	DeadpooledMonth   int            `json:"deadpooled_month"`
	DeadpooledDay     int            `json:"deadpooled_day"`
	TagList           string         `json:"tag_list"`
	EmailAddress      string         `json:"email_address"`
	Overview          string         `json:"overview"`
	TotalMoneyRaised  string         `json:"total_money_raised"`
	Acquisition       Acquisition    `json:"acquisition"`
	Relationships     []Relationship `json:"relationships"`
	Acquisitions      []Acquisition2 `json:"acquisitions"`
	FundingRounds     []FundingRound `json:"funding_rounds"`
	Competitions      []Competitor   `json:"competitions"`
}

type Competitor struct {
	Competitor PermanentLink `json:"competitor"`
}

type Acquisition struct {
	PriceAmount       float64       `json:"price_amount"`
	PriceCurrencyCode string        `json:"price_currency_code"`
	TermCode          string        `json:"term_code"`
	SourceUrl         string        `json:"source_url"`
	SourceDescription string        `json:"source_description"`
	AcquiredYear      int           `json:"acquired_year"`
	AcquiredMonth     int           `json:"acquired_month"`
	AcquiredDay       int           `json:"acquired_day"`
	AcquiringCompany  PermanentLink `json:"acquiring_company"`
}

type Acquisition2 struct {
	PriceAmount       float64       `json:"price_amount"`
	PriceCurrencyCode string        `json:"price_currency_code"`
	TermCode          string        `json:"term_code"`
	SourceUrl         string        `json:"source_url"`
	SourceDescription string        `json:"source_description"`
	AcquiredYear      int           `json:"acquired_year"`
	AcquiredMonth     int           `json:"acquired_month"`
	AcquiredDay       int           `json:"acquired_day"`
	AcquiringCompany  PermanentLink `json:"company"`
}

type Relationship struct {
	IsPast bool   `json:"is_past"`
	Title  string `json:"title"`
	Person Person `json:"person"`
}

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Permalink string `json:"permalink"`
}

type FundingRound struct {
	RoundCode          string       `json:"round_code"`
	SourceUrl          string       `json:"source_url"`
	SourceDescription  string       `json:"source_description"`
	RaisedAmount       float64      `json:"raised_amount"`
	RaisedCurrencyCode string       `json:"raised_currency_code"`
	FundedYear         int          `json:"funded_year"`
	FundedMonth        int          `json:"funded_month"`
	FundedDay          int          `json:"funded_day"`
	Investments        []Investment `json:"investments"`
}

type Investment struct {
	Company      *PermanentLink `json:"company"`
	FinancialOrg *PermanentLink `json:"financial_org"`
	Person       *Person        `json:"person"`
}

type PermanentLink struct {
	Name      string `json:"name"`
	Permalink string `json:"permalink"`
}

func (source *Company) ToSchema() *BsonCompany {

	company := &BsonCompany{
		Id:                source.Id.Id,
		Name:              source.Name,
		Permalink:         source.Permalink,
		CrunchbaseUrl:     source.CrunchbaseUrl,
		HomepageUrl:       source.HomepageUrl,
		CategoryCode:      source.CategoryCode,
		NumberOfEmployees: source.NumberOfEmployees,
		FoundedDate:       toTime(source.FoundedYear, source.FoundedMonth, source.FoundedDay),
		DeadpooledDate:    toTime(source.DeadpooledYear, source.DeadpooledMonth, source.DeadpooledDay),
		TagList:           source.TagList,
		EmailAddress:      source.EmailAddress,
		Overview:          source.Overview,
		TotalMoneyRaised:  source.TotalMoneyRaised,
		Acquisition:       source.Acquisition.toSchema(),
		Acquisitions:      toAcquisitionSchemas(source.Acquisitions),
		Relationships:     toRelationshipSchemas(source.Relationships),
		FundingRounds:     toFundingRoundSchemas(source.FundingRounds),
		Competitions:      toCompetitorSchemas(source.Competitions),
	}

	return company

}

func toTime(year int, month int, day int) *time.Time {
	if year == 0 || month == 0 || day == 0 {
		return nil
	}
	var t = time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	return &t
}

func (source *Relationship) toSchema() BsonRelationship {
	return BsonRelationship{
		IsPast: source.IsPast,
		Title:  source.Title,
		Person: source.Person.toSchema(),
	}
}

func (source *Person) toSchema() BsonPerson {
	return BsonPerson{
		FirstName: source.FirstName,
		LastName:  source.LastName,
		Permalink: source.Permalink,
	}
}

func (source *Acquisition) toSchema() *BsonAcquisition {
	return &BsonAcquisition{
		PriceAmount:       source.PriceAmount,
		PriceCurrencyCode: source.PriceCurrencyCode,
		TermCode:          source.TermCode,
		SourceUrl:         source.SourceUrl,
		SourceDescription: source.SourceDescription,
		AcquiredDate:      time.Date(source.AcquiredYear, time.Month(source.AcquiredMonth), source.AcquiredDay, 0, 0, 0, 0, time.UTC),
		AcquiringCompany:  source.AcquiringCompany,
	}

}

func (source *Acquisition2) toSchema() BsonAcquisition {
	return BsonAcquisition{
		PriceAmount:       source.PriceAmount,
		PriceCurrencyCode: source.PriceCurrencyCode,
		TermCode:          source.TermCode,
		SourceUrl:         source.SourceUrl,
		SourceDescription: source.SourceDescription,
		AcquiredDate:      time.Date(source.AcquiredYear, time.Month(source.AcquiredMonth), source.AcquiredDay, 0, 0, 0, 0, time.UTC),
		AcquiringCompany:  source.AcquiringCompany,
	}

}

func (source FundingRound) toSchema() BsonFundingRound {
	return BsonFundingRound{
		RoundCode:          source.RoundCode,
		SourceUrl:          source.SourceUrl,
		SourceDescription:  source.SourceDescription,
		RaisedAmount:       source.RaisedAmount,
		RaisedCurrencyCode: source.RaisedCurrencyCode,
		FundedDate:         toTime(source.FundedYear, source.FundedMonth, source.FundedDay),
		Investments:        toInvestmentSchemas(source.Investments),
	}
}

func (source Investment) toSchema() BsonInvestment {
	var bsonPerson *BsonPerson
	if source.Person != nil {
		schema := source.Person.toSchema()
		bsonPerson = &schema
	}

	return BsonInvestment{
		Company:      source.Company,
		FinancialOrg: source.FinancialOrg,
		Person:       bsonPerson,
	}
}

func (source Competitor) toSchema() BsonCompetitor {
	return BsonCompetitor{
		Competitor: source.Competitor,
	}
}

func toCompetitorSchemas(sources []Competitor) []*BsonCompetitor {
	var bs []*BsonCompetitor
	for _, source := range sources {
		b := source.toSchema()
		bs = append(bs, &b)
	}

	return bs
}

func toInvestmentSchemas(sources []Investment) []*BsonInvestment {
	var bs []*BsonInvestment
	for _, source := range sources {
		b := source.toSchema()
		bs = append(bs, &b)
	}

	return bs
}

func toFundingRoundSchemas(sources []FundingRound) []*BsonFundingRound {
	var bs []*BsonFundingRound
	for _, source := range sources {
		b := source.toSchema()
		bs = append(bs, &b)
	}

	return bs
}

func toAcquisitionSchemas(sources []Acquisition2) []BsonAcquisition {
	var bs []BsonAcquisition
	for _, source := range sources {
		b := source.toSchema()
		bs = append(bs, b)
	}

	return bs
}

func toRelationshipSchemas(sources []Relationship) []*BsonRelationship {
	var bs []*BsonRelationship
	for _, source := range sources {
		b := source.toSchema()
		bs = append(bs, &b)
	}
	return bs

}

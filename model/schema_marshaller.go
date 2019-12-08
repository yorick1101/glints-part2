package model

import (
	"encoding/json"
	"time"
)

var DATE_FORMAT string = "2006-01-02"

type AliasBsonCompany BsonCompany
type JSONBsonCompany struct {
	AliasBsonCompany
	FoundedDate    string `json:"foundedDate"`
	DeadpooledDate string `json:"deadpooledDate"`
}

func parseDateStr(str string) (*time.Time, error) {
	var date *time.Time = nil
	if len(str) != 0 {
		t, err := time.Parse(DATE_FORMAT, str)
		if err != nil {
			return date, err
		}
		date = &t
	}
	return date, nil
}

func (this BsonCompany) MArshalJSON() ([]byte, error) {

	var fundedDate = ""
	if this.FoundedDate != nil {
		fundedDate = this.FoundedDate.Format(DATE_FORMAT)
	}
	var deadpooledDate = ""
	if this.DeadpooledDate != nil {
		deadpooledDate = this.DeadpooledDate.Format(DATE_FORMAT)
	}

	tmpCompany := JSONBsonCompany{
		AliasBsonCompany: (AliasBsonCompany)(this),
		FoundedDate:      fundedDate,
		DeadpooledDate:   deadpooledDate,
	}
	return json.Marshal(tmpCompany)
}

func (this *BsonCompany) UnmarshalJSON(b []byte) error {
	var jcompany JSONBsonCompany
	err := json.Unmarshal(b, &jcompany)
	if err != nil {
		return err
	}

	fundedDate, err := parseDateStr(jcompany.FoundedDate)
	if err != nil {
		return err
	}

	deadpooledDate, err := parseDateStr(jcompany.DeadpooledDate)
	if err != nil {
		return err
	}

	company := BsonCompany(jcompany.AliasBsonCompany)
	company.FoundedDate = fundedDate
	company.DeadpooledDate = deadpooledDate

	this = &company
	return nil
}

type AliasBsonAcquisition BsonAcquisition
type JSONAcquisition struct {
	AliasBsonAcquisition
	AcquiredDate string `json:"acquiredDate"`
}

func (this BsonAcquisition) MArshalJSON() ([]byte, error) {

	var acquiredDate = ""
	if this.AcquiredDate != nil {
		acquiredDate = this.AcquiredDate.Format(DATE_FORMAT)
	}

	tmpAcquisition := JSONAcquisition{
		AliasBsonAcquisition: (AliasBsonAcquisition)(this),
		AcquiredDate:         acquiredDate,
	}
	return json.Marshal(tmpAcquisition)
}

func (this *BsonAcquisition) UnmarshalJSON(b []byte) error {
	var jacquisition JSONAcquisition
	err := json.Unmarshal(b, &jacquisition)
	if err != nil {
		return err
	}

	acquiredDate, err := parseDateStr(jacquisition.AcquiredDate)
	if err != nil {
		return err
	}

	acquisition := BsonAcquisition(jacquisition.AliasBsonAcquisition)
	acquisition.AcquiredDate = acquiredDate

	this = &acquisition
	return nil
}

type AliasBsonFundingRound BsonFundingRound
type JSONFundingRound struct {
	AliasBsonFundingRound
	FundedDate string `json:"fundedDate"`
}

func (this BsonFundingRound) MArshalJSON() ([]byte, error) {

	var fundedDate = ""
	if this.FundedDate != nil {
		fundedDate = this.FundedDate.Format(DATE_FORMAT)
	}

	tmpAcquisition := JSONFundingRound{
		AliasBsonFundingRound: (AliasBsonFundingRound)(this),
		FundedDate:            fundedDate,
	}
	return json.Marshal(tmpAcquisition)
}

func (this *BsonFundingRound) UnmarshalJSON(b []byte) error {
	var jfundingRound JSONFundingRound
	err := json.Unmarshal(b, &jfundingRound)
	if err != nil {
		return err
	}

	fundedDate, err := parseDateStr(jfundingRound.FundedDate)
	if err != nil {
		return err
	}

	fundingRound := BsonFundingRound(jfundingRound.AliasBsonFundingRound)
	fundingRound.FundedDate = fundedDate

	this = &fundingRound
	return nil
}

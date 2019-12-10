package web

/*
GET /company/funding/rounds? lt, gt, eq
GET /company/funding/amount? lt, gt, eq
GET /company/funding/date? lt, gt, eq
GET /company/deadpool/date? lt, gt, eq
GET /company/person/invested/{personId}
GET /company/person/employed/{personId}
GET /company/other/acquisition/{companyId}
GET /company/other/invested/{companyId}
GET /company/{companyId}
POST /company/{companyId}

GET /person/permalink/{permalink}
GET /person/{personId}
*/

import (
	"github.com/gin-gonic/gin"
)

func InitRouters() *gin.Engine {
	router := gin.Default()
	hadler := NewHandler()

	router.GET("/company", hadler.getCompanyHandler)
	router.POST("/company", hadler.updateCompanyHandler)
	router.DELETE("/company", hadler.deleteCompanyHandler)

	companyFunding := router.Group("/company/search/funding")
	{
		companyFunding.GET("/rounds", hadler.fundingRoundsHandler)
		companyFunding.GET("/amount", hadler.fundingAmountHandler)
		companyFunding.GET("/date", hadler.fundingDateHandler)
	}

	router.GET("/company/search/deadpool/date", hadler.deadpoolHandler)

	companyPerson := router.Group("/company/search/person")
	{
		companyPerson.GET("/invested/:personId", hadler.personInvestedHandler)
		companyPerson.GET("/employed/:personId", hadler.personEmployedtHandler)
	}

	companyOther := router.Group("/company/search/other")
	{
		companyOther.GET("/acquisition/:companyId", hadler.otherAcquisitiondHandler)
		companyOther.GET("/invested/:companyId", hadler.otherInvesteddHandler)
	}

	person := router.Group("/person")
	{
		person.GET("/permalink/:permalink", hadler.getPersonByPermalinkHandler)
		person.GET("/id/:personId", hadler.getPersonHandler)
	}

	return router
}

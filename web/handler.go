package web

import (
	"errors"
	"glints-part2/model"
	"glints-part2/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func NewHandler() *RequestHandler {
	factory := repository.NewRepositoryFactory()

	return &RequestHandler{
		factory.GetRepository(),
	}
}

type RequestHandler struct {
	repo *repository.Repository
}

func response(c *gin.Context, httpStatus int, message interface{}) {
	c.JSON(httpStatus, gin.H{
		"code":    httpStatus,
		"message": message,
	})
}

//getComapnyHandler
func (handler *RequestHandler) getCompanyHandler(c *gin.Context) {
	companyIds := c.QueryArray("id")
	if len(companyIds) == 0 {
		response(c, http.StatusBadRequest, "id cannot be empty")
		return
	}

	companies, err := handler.repo.FindCompanyByIds(companyIds)
	if err != nil {
		response(c, http.StatusInternalServerError, err.Error())
	} else if companies == nil {
		response(c, http.StatusNotFound, "no companies found")
	} else {
		response(c, http.StatusOK, companies)
	}
}

//updateCompanyHandler
func (handler *RequestHandler) updateCompanyHandler(c *gin.Context) {
	var company model.BsonCompany
	c.BindJSON(&company)
	count, err := handler.repo.ReplaceCompany(&company)
	if err != nil {
		response(c, http.StatusInternalServerError, err.Error())
		return
	}

	if count == 1 {
		response(c, http.StatusOK, "document updated")
	} else {
		response(c, http.StatusNotFound, "document not found")
	}
}

func checkIntFilter(c *gin.Context) (error, []model.IntFilter) {
	lt := c.Query("lt")
	gt := c.Query("gt")
	eq := c.Query("eq")

	var filters []model.IntFilter
	if eq == "" && lt == "" && gt == "" {
		err := errors.New("Must have query parameter lt,gt or eq")
		return err, filters
	}
	if eq != "" && (lt != "" || gt != "") {
		err := errors.New("Ambigious request query, eq should not appear with lt or gt")
		return err, filters
	}
	if eq != "" {
		value, err := strconv.Atoi(eq)
		if err != nil {
			return err, filters
		}
		filters = append(filters, model.IntFilter{Operation: "eq", Value: value})
		return nil, filters
	}
	if lt != "" {
		value, err := strconv.Atoi(lt)
		if err != nil {
			return err, filters
		}
		filters = append(filters, model.IntFilter{Operation: "lt", Value: value})
	}

	if gt != "" {
		value, err := strconv.Atoi(gt)
		if err != nil {
			return err, filters
		}
		filters = append(filters, model.IntFilter{Operation: "gt", Value: value})
	}

	return nil, filters
}

//fundingRoundsHandler
func (handler *RequestHandler) fundingRoundsHandler(c *gin.Context) {
	err, filters := checkIntFilter(c)
	if err != nil {
		response(c, http.StatusBadRequest, err.Error())
		return
	}
	ids, err := handler.repo.FilterByNumberOfFundingRounds(filters)
	if err != nil {
		response(c, http.StatusInternalServerError, err.Error())
		return
	}
	response(c, http.StatusOK, ids)
}

//fundingAmountHandler
func (handler *RequestHandler) fundingAmountHandler(c *gin.Context) {
	err, filters := checkIntFilter(c)
	if err != nil {
		response(c, http.StatusBadRequest, err.Error())
		return
	}
	ids, err := handler.repo.FilterByAmountOfFundingRounds(filters)
	if err != nil {
		response(c, http.StatusInternalServerError, err.Error())
		return
	}
	response(c, http.StatusOK, ids)
}

func checkDateFilter(c *gin.Context) (error, []model.DateFilter) {
	lt := c.Query("lt")
	gt := c.Query("gt")
	eq := c.Query("eq")

	var filters []model.DateFilter
	if eq == "" && lt == "" && gt == "" {
		err := errors.New("Must have query parameter lt,gt or eq")
		return err, filters
	}

	if eq != "" && (lt != "" || gt != "") {
		err := errors.New("Ambigious request query, eq should not appear with lt or gt")
		return err, filters
	}
	if eq != "" {
		value, err := model.ParseDateStr(eq)
		if err != nil {
			return err, filters
		}
		filters = append(filters, model.DateFilter{Operation: "eq", Value: value})
		return nil, filters
	}
	if lt != "" {
		value, err := model.ParseDateStr(lt)
		if err != nil {
			return err, filters
		}
		filters = append(filters, model.DateFilter{Operation: "lt", Value: value})
	}

	if gt != "" {
		value, err := model.ParseDateStr(gt)
		if err != nil {
			return err, filters
		}
		filters = append(filters, model.DateFilter{Operation: "gt", Value: value})
	}

	return nil, filters
}

//fundingDateHandler
func (handler *RequestHandler) fundingDateHandler(c *gin.Context) {
	err, filters := checkDateFilter(c)
	if err != nil {
		response(c, http.StatusBadRequest, err.Error())
		return
	}
	ids, err := handler.repo.FilterByFundingDate(filters)
	if err != nil {
		response(c, http.StatusInternalServerError, err.Error())
		return
	}
	response(c, http.StatusOK, ids)
}

//deadpoolHandler
func (handler *RequestHandler) deadpoolHandler(c *gin.Context) {
	err, filters := checkDateFilter(c)
	if err != nil {
		response(c, http.StatusBadRequest, err)
		return
	}
	ids, err := handler.repo.FilterByDeadedpoolDate(filters)
	if err != nil {
		response(c, http.StatusInternalServerError, err.Error())
		return
	}
	response(c, http.StatusOK, ids)
}

//personInvestedHandler
func (handler *RequestHandler) personInvestedHandler(c *gin.Context) {
	personId := c.Param("personId")
	if personId == "" {
		response(c, http.StatusBadRequest, "cannot find personId")
	}

	ids, err := handler.repo.FindPersonOnFundingRounds(personId)
	if err != nil {
		response(c, http.StatusInternalServerError, err.Error())
		return
	}

	response(c, http.StatusOK, ids)
}

//personEmployedtHandler
func (handler *RequestHandler) personEmployedtHandler(c *gin.Context) {
	personId := c.Param("personId")
	if personId == "" {
		response(c, http.StatusBadRequest, "cannot find personId")
	}

	ids, err := handler.repo.FindPersonOnRelationship(personId)
	if err != nil {
		response(c, http.StatusInternalServerError, err.Error())
		return
	}

	response(c, http.StatusOK, ids)
}

//otherAcquisitiondHandler
func (handler *RequestHandler) otherAcquisitiondHandler(c *gin.Context) {
	companyId := c.Param("companyId")
	if companyId == "" {
		response(c, http.StatusBadRequest, "cannot find companyId")
	}

	ids, err := handler.repo.FindCompanyOnAcquisitions(companyId)
	if err != nil {
		response(c, http.StatusInternalServerError, err.Error())
		return
	}

	response(c, http.StatusOK, ids)
}

//otherInvesteddHandler
func (handler *RequestHandler) otherInvesteddHandler(c *gin.Context) {
	companyId := c.Param("companyId")
	if companyId == "" {
		response(c, http.StatusBadRequest, "cannot find companyId")
	}

	ids, err := handler.repo.FindCompanyOnFundingRounds(companyId)
	if err != nil {
		response(c, http.StatusInternalServerError, err.Error())
		return
	}

	response(c, http.StatusOK, ids)
}

//getPersonByPermalinkHandler
func (handler *RequestHandler) getPersonByPermalinkHandler(c *gin.Context) {
	permalink := c.Param("permalink")
	if permalink == "" {
		response(c, http.StatusBadRequest, "cannot find permalink")
	}
	person, err := handler.repo.FindPersonByPermalink(permalink)
	if err != nil {
		response(c, http.StatusInternalServerError, err.Error())
		return
	}
	response(c, http.StatusOK, person)
}

//getPersonHandler
func (handler *RequestHandler) getPersonHandler(c *gin.Context) {
	personId := c.Param("personId")
	if personId == "" {
		response(c, http.StatusBadRequest, "cannot find personId")
		return
	}
	person, err := handler.repo.FindPersonById(personId)
	if err != nil {
		response(c, http.StatusInternalServerError, err.Error())
		return
	}
	response(c, http.StatusOK, person)

}

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
		response(c, http.StatusInternalServerError, err)
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
		response(c, http.StatusInternalServerError, err)
		return
	}

	if count == 1 {
		response(c, http.StatusOK, "document updated")
	} else {
		response(c, http.StatusNotFound, "document not found")
	}
}

func checkRangeFilter(c *gin.Context) (error, []model.IntFilter) {
	lt := c.Query("lt")
	gt := c.Query("gt")
	eq := c.Query("eq")

	var filters []model.IntFilter
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
	err, filters := checkRangeFilter(c)
	if err != nil {
		response(c, http.StatusBadRequest, err)
		return
	}
	ids, err := handler.repo.FilterByNumberOfFundingRounds(filters)
	if err != nil {
		response(c, http.StatusInternalServerError, err)
		return
	}
	response(c, http.StatusOK, ids)
}

//fundingAmountHandler
func (handler *RequestHandler) fundingAmountHandler(c *gin.Context) {
	err, filters := checkRangeFilter(c)
	if err != nil {
		response(c, http.StatusBadRequest, err)
		return
	}
	ids, err := handler.repo.FilterByAmountOfFundingRounds(filters)
	if err != nil {
		response(c, http.StatusInternalServerError, err)
		return
	}
	response(c, http.StatusOK, ids)
}

//fundingDateHandler
func (handler *RequestHandler) fundingDateHandler(c *gin.Context) {
	response(c, http.StatusOK, "not implemented yet")
}

//deadpoolHandler
func (handler *RequestHandler) deadpoolHandler(c *gin.Context) {
	response(c, http.StatusOK, "not implemented yet")
}

//personInvestedHandler
func (handler *RequestHandler) personInvestedHandler(c *gin.Context) {
	response(c, http.StatusOK, "not implemented yet")
}

//personEmployedtHandler
func (handler *RequestHandler) personEmployedtHandler(c *gin.Context) {
	response(c, http.StatusOK, "not implemented yet")
}

//otherAcquisitiondHandler
func (handler *RequestHandler) otherAcquisitiondHandler(c *gin.Context) {
	response(c, http.StatusOK, "not implemented yet")
}

//otherInvesteddHandler
func (handler *RequestHandler) otherInvesteddHandler(c *gin.Context) {
	response(c, http.StatusOK, "not implemented yet")
}

//getPersonByPermalinkHandler
func (handler *RequestHandler) getPersonByPermalinkHandler(c *gin.Context) {
	response(c, http.StatusOK, "not implemented yet")
}

//getPersonHandler
func (handler *RequestHandler) getPersonHandler(c *gin.Context) {
	response(c, http.StatusOK, "not implemented yet")

}

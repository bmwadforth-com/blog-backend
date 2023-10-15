package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web-template/database"
	"web-template/models"
	"web-template/util"
)

// CreatePerson example godoc
// @Summary Create Person
// @Schemes
// @Description Create Person
// @Tags Create Person
// @Param Create Person body models.PersonCreateRequest true "Create Person model object"
// @Accept json
// @Produce json
// @Success 200 {object}  util.ApiResponse
// @Router /person [post]
func CreatePerson(c *gin.Context) {
	var request models.PersonCreateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		statusCode, response := util.NewResponse(http.StatusBadRequest, "Failed to parse JSON", nil, err)
		c.JSON(statusCode, response)
		return
	}

	r := database.CreatePerson(request)
	if r.GetError() != nil {
		statusCode, response := util.NewResponse(http.StatusInternalServerError, "An error has occurred", nil, r.GetError())
		c.JSON(statusCode, response)
		return
	}

	statusCode, response := util.NewResponse(http.StatusOK, r.Message, r.Data, nil)
	c.JSON(statusCode, response)
}

// GetPeople example godoc
// @Summary Create Person
// @Schemes
// @Description Get people
// @Tags Get people
// @Accept json
// @Produce json
// @Success 200 {object}  util.ApiResponse
// @Router /person [get]
func GetPeople(c *gin.Context) {
	r := database.GetPeople()
	if r.GetError() != nil {
		statusCode, response := util.NewResponse(http.StatusInternalServerError, "An error has occurred", nil, r.GetError())
		c.JSON(statusCode, response)
		return
	}

	statusCode, response := util.NewResponse(http.StatusOK, r.Message, r.Data, nil)
	c.JSON(statusCode, response)
}

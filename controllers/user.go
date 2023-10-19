package controllers

import (
	"blog-backend/database"
	"blog-backend/models"
	"blog-backend/service"
	"blog-backend/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateUser example godoc
// @Summary Create user
// @Schemes
// @Description Create user
// @Tags Create user
// @Param Create user body models.CreateUserRequest true "Create user model object"
// @Accept json
// @Produce json
// @Success 200 {object}  util.ApiResponse
// @Router /user [post]
func CreateUser(c *gin.Context) {
	var request models.CreateUserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		response := util.NewResponse(http.StatusBadRequest, "Failed to parse JSON", "", err)
		c.JSON(response.GetStatusCode(), response)
		return
	}

	r := database.CreateUser(request)
	if r.GetError() != nil {
		response := util.NewResponse(http.StatusInternalServerError, "An error has occurred", "", r.GetError())
		c.JSON(response.GetStatusCode(), response)
		return
	}

	response := util.NewResponse(http.StatusOK, r.Message, r.Data, nil)
	c.JSON(response.GetStatusCode(), response)
}

// LoginUser example godoc
// @Summary Login user
// @Schemes
// @Description Login user
// @Tags Login user
// @Param Create user body models.LoginUserRequest true "Login user model object"
// @Accept json
// @Produce json
// @Success 200 {object}  util.ApiResponse
// @Router /login [post]
func LoginUser(c *gin.Context) {
	var request models.LoginUserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		response := util.NewResponse(http.StatusBadRequest, "Failed to parse JSON", "", err)
		c.JSON(response.GetStatusCode(), response)
		return
	}

	valid, token, err := service.LoginUser(request)
	if valid == false {
		response := util.NewResponse(http.StatusUnauthorized, "Unsuccessful", "", err)
		c.JSON(response.GetStatusCode(), response)
		return
	}

	response := util.NewResponse(http.StatusOK, "Successful", token, nil)
	c.JSON(response.GetStatusCode(), response)
}

// GetSessions example godoc
// @Summary Get user sessions
// @Schemes
// @Description Get user sessions
// @Tags Get user sessions
// @Accept json
// @Produce json
// @Success 200 {object}  util.ApiResponse
// @Router /sessions [get]
func GetSessions(c *gin.Context) {
	sessions := service.Sessions
	response := util.NewResponse(http.StatusOK, "Successful", sessions, nil)
	c.JSON(response.GetStatusCode(), response)
}

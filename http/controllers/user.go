package controllers

import (
	"blog-backend/data/database"
	"blog-backend/data/models"
	"blog-backend/service"
	"blog-backend/util"
	armorUtil "github.com/bmwadforth-com/armor-go/src/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateUser example godoc
// @Summary Create user
// @Schemes
// @Description Create user
// @Tags Create user
// @Param user body models.CreateUserRequest true "Create user model object"
// @Accept json
// @Produce json
// @Success 200 {object}  util.ApiResponse[string]
// @Router /user [post]
func CreateUser(c *gin.Context) {
	var request models.CreateUserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		response := util.NewResponse(http.StatusBadRequest, "Failed to parse JSON", "", err)
		c.JSON(response.GetStatusCode(), response)
		return
	}

	r := database.CreateUser(request, c.Request.Context())
	if r.GetError() != nil {
		response := util.NewResponse(http.StatusInternalServerError, "An error has occurred", "", r.GetError())
		c.JSON(response.GetStatusCode(), response)
		return
	}

	armorUtil.LogInfo("new request to create user with username: %s", request.Username)
	response := util.NewResponse(http.StatusOK, r.Message, r.Data, nil)
	c.JSON(response.GetStatusCode(), response)
}

// LoginUser example godoc
// @Summary Login user
// @Schemes
// @Description Login user
// @Tags Login user
// @Param User body models.LoginUserRequest true "Login user model object"
// @Accept json
// @Produce json
// @Success 200 {object}  util.ApiResponse[string]
// @Router /login [post]
func LoginUser(c *gin.Context) {
	var request models.LoginUserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		response := util.NewResponse(http.StatusBadRequest, "Failed to parse JSON", "", err)
		c.JSON(response.GetStatusCode(), response)
		return
	}

	valid, token, err := service.LoginUser(request, c.Request.Context())
	if valid == false {
		response := util.NewResponse(http.StatusUnauthorized, "Unsuccessful", "", err)
		c.JSON(response.GetStatusCode(), response)
		return
	}

	armorUtil.LogInfo("username: %s has logged in", request.Username)
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
// @Success 200 {object}  util.ApiResponse[models.UserSessionModel]
// @Router /sessions [get]
func GetSessions(c *gin.Context) {
	sessions := service.Sessions
	response := util.NewResponse(http.StatusOK, "Successful", sessions, nil)
	c.JSON(response.GetStatusCode(), response)
}

// GetStatus example godoc
// @Summary Get user status
// @Schemes
// @Description Get user status
// @Tags Get user status
// @Accept json
// @Produce json
// @Success 200 {object}  util.ApiResponse[models.UserStatusModel]
// @Router /status [get]
func GetStatus(c *gin.Context) {
	bearer, err := util.GetBearerToken(c.Request)
	if err != nil {
		response := util.NewResponse(http.StatusUnauthorized, "invalid authentication", "", nil)
		c.AbortWithStatusJSON(response.GetStatusCode(), response)
		return
	}

	claims := service.GetTokenClaims(bearer)
	session := service.Sessions[claims["sub"].(string)]
	r := models.UserStatusModel{
		Username:      session.Username,
		Active:        session.Active,
		LoggedInSince: session.LoggedIn.String(),
	}

	response := util.NewResponse(http.StatusOK, "Successful", r, nil)
	c.JSON(response.GetStatusCode(), response)
}

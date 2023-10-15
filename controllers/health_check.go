package controllers

import (
	"blog-backend/database"
	"blog-backend/util"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HealthCheckResponse struct {
	Database bool `json:"database"`
}

// HealthCheck Generic healthcheck endpoint godoc
// @Summary HealthCheck
// @Schemes
// @Description HealthCheck
// @Tags HealthCheck
// @Accept json
// @Produce json
// @Success 200 {object}  util.ApiResponse
// @Router /healthz [get]
func HealthCheck(c *gin.Context) {
	healthCheck := HealthCheckResponse{Database: true}
	response := util.NewResponse(http.StatusOK, "healthz", healthCheck, nil)
	dbErr := database.HealthCheck()
	if dbErr != nil {
		healthCheck.Database = false
		response.SetData(healthCheck)
		response.SetError(errors.New("the api is not healthy"))
		response.SetStatusCode(http.StatusBadGateway)
	}

	c.JSON(response.GetStatusCode(), response)
}

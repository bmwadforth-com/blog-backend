package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"blog-backend/util"
)

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
	statusCode, response := util.NewResponse(http.StatusOK, "healthz", nil, nil)
	// TODO write logic in here to ping upstream services, if any are down return bad healthchecl
	c.JSON(statusCode, response)
}

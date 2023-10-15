package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Ping Generic ping endpoint godoc
// @Summary Ping
// @Schemes
// @Description Ping
// @Tags Ping
// @Accept json
// @Produce json
// @Success 200 {object}  util.ApiResponse
// @Router /ping [get]
func Ping(c *gin.Context) {
	c.Status(http.StatusOK)
}

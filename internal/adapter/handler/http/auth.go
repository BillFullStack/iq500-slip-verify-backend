package http

import (
	"main/internal/core/port"
	"main/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthenticationHandler struct {
	svc port.AuthenticationService
}

func NewAuthenticationHandler(svc port.AuthenticationService) *AuthenticationHandler {
	return &AuthenticationHandler{
		svc,
	}
}

func (h *AuthenticationHandler) Login(c *gin.Context) {
	utils.Response(c, http.StatusOK, 200, "success", "ok", nil)
}

func (h *AuthenticationHandler) Register(c *gin.Context) {
	utils.Response(c, http.StatusOK, 200, "success", "ok", nil)

}

func (h *AuthenticationHandler) ResetPassword(c *gin.Context) {
	utils.Response(c, http.StatusOK, 200, "success", "ok", nil)

}

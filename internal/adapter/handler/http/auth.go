package http

import (
	"fmt"
	"main/internal/core/domain"
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
	fmt.Println("Login")

	var payload domain.PayloadUser

	if err := c.ShouldBind(&payload); err != nil {
		fmt.Println("error bind", err)
		utils.Response(c, http.StatusBadRequest, 400, "ข้อมูลไม่ถูกต้อง กรุณาลองใหม่อีกครั้ง", err.Error(), nil)
		return
	}

	token, err := h.svc.Login(c, payload)
	if err != nil {
		return
	}

	utils.Response(c, http.StatusOK, 200, "success", "ok", token)
}

func (h *AuthenticationHandler) Register(c *gin.Context) {
	fmt.Println("Register")

	var payload domain.PayloadUser
	if err := c.ShouldBind(&payload); err != nil {
		fmt.Println("error bind", err)
		utils.Response(c, http.StatusBadRequest, 400, "ข้อมูลไม่ถูกต้อง กรุณาลองใหม่อีกครั้ง", err.Error(), nil)
		return
	}

	if err := h.svc.Register(c, payload); err != nil {
		return
	}

	utils.Response(c, http.StatusOK, 200, "success", "ok", nil)
}

func (h *AuthenticationHandler) ResetPassword(c *gin.Context) {
	fmt.Println("ResetPassword")

	var payload domain.PayloadResetPassword
	if err := c.ShouldBind(&payload); err != nil {
		fmt.Println("error bind", err)
		utils.Response(c, http.StatusBadRequest, 400, "ข้อมูลไม่ถูกต้อง กรุณาลองใหม่อีกครั้ง", err.Error(), nil)
		return
	}

	if err := h.svc.ResetPassword(c, payload); err != nil {
		return
	}

	utils.Response(c, http.StatusOK, 200, "success", "ok", nil)

}

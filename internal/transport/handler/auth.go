package handler

import (
	"net/http"
	"net/mail"

	"github.com/Korisss/concet-backend/internal/domain"
	"github.com/gin-gonic/gin"
)

type loginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type requestWithAuth struct {
	Id int `uri:"id" binding:"required"`
}

func checkUserAccess(ctx *gin.Context) (int, bool) {
	ctxId, _ := ctx.Get("id")
	id := ctxId.(int)

	var req requestWithAuth

	if err := ctx.ShouldBindUri(&req); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return id, false
	}

	if id != req.Id {
		newErrorResponse(ctx, http.StatusForbidden, "no access")
		return id, false
	}

	return id, true
}

func (h *Handler) signIn(ctx *gin.Context) {
	var req loginRequest

	if err := ctx.BindJSON(&req); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id, token, err := h.services.GenerateToken(req.Email, req.Password)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id":    id,
		"token": token,
	})
}

func (h *Handler) signUp(ctx *gin.Context) {
	var req domain.User

	if err := ctx.BindJSON(&req); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if _, err := mail.ParseAddress(req.Email); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if len(req.Password) < 6 {
		newErrorResponse(ctx, http.StatusBadRequest, "password must contain at least 6 characters")
		return
	}

	id, err := h.services.CreateUser(req)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

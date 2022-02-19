package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "id"
)

func (h *Handler) userIdentity(ctx *gin.Context) {
	header := ctx.GetHeader(authorizationHeader)
	if header == "" {
		newErrorResponse(ctx, http.StatusUnauthorized, "empty auth header")
		return
	}

	authHeader := strings.Split(header, " ")
	if len(authHeader) != 2 {
		newErrorResponse(ctx, http.StatusUnauthorized, "invalid auth header")
		return
	}

	userId, err := h.services.Authorization.ParseToken(authHeader[1])
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	ctx.Set(userCtx, userId)
}

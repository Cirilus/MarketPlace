package http

import (
	"CrowdProject/internal/auth"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Handler struct {
	useCase auth.UseCase
}

func NewHandler(useCase auth.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

type signInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *Handler) SignUp(c *gin.Context) {
	inp := new(signInput)

	if err := c.BindJSON(inp); err != nil {
		logrus.Error("Bad request, err =", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if err := h.useCase.SignUp(c.Request.Context(), inp.Username, inp.Password); err != nil {
		logrus.Error("Error in handler err= ", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusCreated)
}

type signInResponse struct {
	Token string `json:"token"`
}

func (h *Handler) SignIn(c *gin.Context) {
	inp := new(signInput)

	if err := c.BindJSON(inp); err != nil {
		logrus.Error("Bad request, err =", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	token, err := h.useCase.SignIn(c.Request.Context(), inp.Username, inp.Password)
	if err != nil {
		if err == auth.ErrUserNotFound {
			logrus.Error(err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		logrus.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, signInResponse{Token: token})
}

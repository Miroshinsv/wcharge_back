package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Miroshinsv/wcharge_back/internal/entity"
	"github.com/Miroshinsv/wcharge_back/internal/usecase"
	"github.com/Miroshinsv/wcharge_back/pkg/logger"
)

type userRoutes struct {
	t usecase.User
	l logger.Interface
}

func newUserRoutes(handler *gin.RouterGroup, u usecase.User, l logger.Interface) {
	r := &userRoutes{u, l}

	h := handler.Group("/user")
	{
		h.GET("/all", r.GetAllUsers)
		h.POST("/update/{id}", r.UpdateUser)
	}
}

type historyResponse struct {
	History []entity.User `json:"history"`
}

// @Summary     Get all users
// @Accept      json
// @Produce     json
// @Success     200 {object} usersResponse
// @Failure     500 {object} response
// @Router      /user/all [get]
func (r *userRoutes) GetAllUsers(c *gin.Context) {
	translations, err := r.t.History(c.Request.Context())
	if err != nil {
		r.l.Error(err, "http - v1 - users")
		errorResponse(c, http.StatusInternalServerError, "database problems")

		return
	}

	c.JSON(http.StatusOK, historyResponse{translations})
}

type UpdateRequest struct {
	ID       int    `json:"id"       binding:"required"`
	Username string `json:"username" binding:"required"`
	Email    string `json:"email"    binding:"required"`
}

func (r *userRoutes) UpdateUser(c *gin.Context) {
	var request UpdateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		r.l.Error(err, "http - v1 - doTranslate")
		errorResponse(c, http.StatusBadRequest, "invalid request body")

		return
	}

	translation, err := r.t.Translate(
		c.Request.Context(),
		entity.User{
			ID:       request.ID,
			Username: request.Username,
			Email:    request.Email,
		},
	)
	if err != nil {
		r.l.Error(err, "http - v1 - doTranslate")
		errorResponse(c, http.StatusInternalServerError, "translation service problems")

		return
	}

	c.JSON(http.StatusOK, translation)
}

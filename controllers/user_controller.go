package controllers

import (
	"autumn/models"
	"autumn/pkg/autumncore/generic"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController[T generic.Model[T]] struct {
	Val T
	*generic.ControllerImpl[T]
}

func NewUserController() *UserController[models.User] {

	return &UserController[models.User]{}
}

func (u *UserController[T]) Get(g *gin.Context) {
	data, err := generic.NewModelImpl[models.Block]().Get()
	if err != nil {
		g.AbortWithError(http.StatusBadRequest, err)
		return
	}

	g.JSON(http.StatusAccepted, data)
}

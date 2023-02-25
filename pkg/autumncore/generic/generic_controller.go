package generic

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller[T Model[T]] interface {
	Get(*gin.Context)
	Insert(*gin.Context)
	Update(*gin.Context)
	Delete(*gin.Context)
}

type ControllerImpl[T Model[T]] struct {
	Val T
}

func NewControllerImpl[T Model[T]]() Controller[T] {
	return &ControllerImpl[T]{}
}

func (c *ControllerImpl[T]) Get(g *gin.Context) {
	genericModel := NewModelImpl[T]()

	data, err := genericModel.Get()
	if err != nil {
		g.AbortWithError(http.StatusBadRequest, err)
		return
	}

	g.JSON(http.StatusAccepted, data)
}

func (c *ControllerImpl[T]) Insert(g *gin.Context) {
	genericModel := NewModelImpl[T]()

	var requestPayload T

	err := g.BindJSON(&requestPayload)
	if err != nil {
		g.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = genericModel.Insert(requestPayload)
	if err != nil {
		g.AbortWithError(http.StatusBadRequest, err)
		return
	}

	g.JSON(http.StatusAccepted, "Success")
}

func (c *ControllerImpl[T]) Update(g *gin.Context) {
	genericModel := NewModelImpl[T]()

	var requestPayload T

	err := g.BindJSON(&requestPayload)
	if err != nil {
		g.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = genericModel.Update(requestPayload)
	if err != nil {
		g.AbortWithError(http.StatusBadRequest, err)
		return
	}

	g.JSON(http.StatusAccepted, "Success")
}

func (c *ControllerImpl[T]) Delete(g *gin.Context) {
	genericModel := NewModelImpl[T]()

	id := g.Query("id")

	err := genericModel.Delete(id)
	if err != nil {
		g.AbortWithError(http.StatusBadRequest, err)
		return
	}

	g.JSON(http.StatusAccepted, "Success")
}

func Resources[T Model[T]](groupName string, router *gin.RouterGroup, c Controller[T]) {
	genericGroup := router.Group(groupName)
	{
		genericGroup.GET("get", c.Get)
		genericGroup.POST("create", c.Insert)
		genericGroup.POST("update", c.Update)
		genericGroup.DELETE("delete", c.Delete)
	}
}

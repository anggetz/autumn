package generic

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller[T Model[T]] struct {
	Val T
}

func NewController[T Model[T]]() *Controller[T] {
	return &Controller[T]{}
}

func (c *Controller[T]) Get(g *gin.Context) {
	data, err := c.Val.Get()
	if err != nil {
		g.AbortWithError(http.StatusBadRequest, err)
		return
	}

	g.JSON(http.StatusAccepted, data)
}

func (c *Controller[T]) Insert(g *gin.Context) {

	var requestPayload T

	err := g.BindJSON(&requestPayload)
	if err != nil {
		g.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = c.Val.Insert(requestPayload)
	if err != nil {
		g.AbortWithError(http.StatusBadRequest, err)
		return
	}

	g.JSON(http.StatusAccepted, "Success")
}

func (c *Controller[T]) Update(g *gin.Context) {

	var requestPayload T

	err := g.BindJSON(&requestPayload)
	if err != nil {
		g.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = c.Val.Update(requestPayload)
	if err != nil {
		g.AbortWithError(http.StatusBadRequest, err)
		return
	}

	g.JSON(http.StatusAccepted, "Success")
}

func (c *Controller[T]) Delete(g *gin.Context) {

	id := g.Query("id")

	err := c.Val.Delete(id)
	if err != nil {
		g.AbortWithError(http.StatusBadRequest, err)
		return
	}

	g.JSON(http.StatusAccepted, "Success")
}

func (c *Controller[T]) Resources(groupName string, router *gin.RouterGroup) {
	genericGroup := router.Group(groupName)
	{
		genericGroup.GET("get", c.Get)
		genericGroup.POST("create", c.Insert)
		genericGroup.POST("update", c.Update)
		genericGroup.DELETE("delete", c.Delete)
	}
}

package controllers

import (
	"base-go/application/cats"
	"base-go/common/logger"
	"net/http"

	"github.com/labstack/echo"
)

type CatsController struct {
	catsInteractor cats.CatsInteractor
}

func NewCatsController(catsInteractor cats.CatsInteractor) *CatsController {
	return &CatsController{catsInteractor}
}

func (controller *CatsController) Mount(e *echo.Echo) {
	g := e.Group("/cats")
	g.GET("/ping", controller.Ping)
	g.GET("/:id", controller.GetCat)
	g.POST("", controller.AddCat)
}

func (controller *CatsController) Ping(c echo.Context) error {
	c.JSON(200, "pong")
	return nil
}

func (controller *CatsController) GetCat(c echo.Context) error {
	id := c.Param("id")
	// should validate ipt here
	logger.Info("GetCat input: id=%s", id)
	cat, err := controller.catsInteractor.GetCat(c.Request().Context(), id)
	if err != nil {
		// should delegate to echo's error handler instead, but for now it hasn't been setup yet
		c.JSON(http.StatusBadGateway, err)
		return nil
	}
	// might need a presenter layer and a response model
	c.JSON(http.StatusOK, cat)
	return nil
}

func (controller *CatsController) AddCat(c echo.Context) error {
	catIpt := cats.AddCatIpt{}
	// should validate ipt here
	logger.Info("AddCat input: %+v", catIpt)
	err := c.Bind(&catIpt)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return nil
	}

	newCat, err := controller.catsInteractor.AddCat(c.Request().Context(), catIpt)
	if err != nil {
		c.JSON(http.StatusBadGateway, err)
		return nil
	}
	c.JSON(http.StatusOK, newCat)
	return nil
}

package handlers

import (
	"fmt"
	"github.com/anggakes/gosvelte/src/backend/pkg/models"
	"github.com/anggakes/gosvelte/src/backend/pkg/usecases/todo"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type TodoHandler struct {
	TodoUseCase todo.ITodo
}

func (c *TodoHandler) Get(e echo.Context) error {

	ctx := e.Request().Context()
	id := e.Param("id")

	td, err := c.TodoUseCase.Get(ctx, &models.TodoGetQuery{
		ID: id,
	})
	if err != nil {
		fmt.Println(err)
		return err
	}

	return e.JSON(http.StatusOK, td)
}

func (c *TodoHandler) Create(e echo.Context) error {

	ctx := e.Request().Context()

	req := &models.TodoCreateCommand{}

	if err := e.Bind(req); err != nil {
		return echo.ErrBadRequest
	}

	td, err := c.TodoUseCase.Create(ctx, req)
	if err != nil {
		return err
	}

	return e.JSON(http.StatusOK, td)
}

func (c *TodoHandler) List(e echo.Context) error {

	ctx := e.Request().Context()

	page, _ := strconv.Atoi(e.QueryParam("page"))
	size, _ := strconv.Atoi(e.QueryParam("size"))

	if page == 0 {
		page = 1
	}

	if size == 0 {
		size = 5
	}

	td, err := c.TodoUseCase.List(ctx, &models.TodoListQuery{
		Page: page,
		Size: size,
	})
	if err != nil {
		return err
	}

	return e.JSON(http.StatusOK, td)

}

func (c *TodoHandler) Update(e echo.Context) error {

	ctx := e.Request().Context()

	req := &models.TodoUpdateCommand{}
	if err := e.Bind(req); err != nil {
		return echo.ErrBadRequest
	}

	req.ID, _ = strconv.Atoi(e.Param("id"))

	td, err := c.TodoUseCase.Update(ctx, req)
	if err != nil {
		return err
	}

	return e.JSON(http.StatusOK, td)
}

func (c *TodoHandler) Delete(e echo.Context) error {

	ctx := e.Request().Context()

	id := e.Param("id")

	td, err := c.TodoUseCase.Delete(ctx, &models.TodoGetQuery{
		ID: id,
	})
	if err != nil {
		return err
	}

	return e.JSON(http.StatusOK, td)
}

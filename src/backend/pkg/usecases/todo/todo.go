package todo

import (
	"context"
	"database/sql"
	"github.com/anggakes/gosvelte/src/backend/dbmodels"
	"github.com/anggakes/gosvelte/src/backend/pkg/events"
	"github.com/anggakes/gosvelte/src/backend/pkg/models"
	"github.com/pkg/errors"
	"github.com/ulule/deepcopier"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

var (
	ErrExist = errors.New("Todo already exist")
)

type ITodo interface {
	Get(ctx context.Context, query *models.TodoGetQuery) (*models.Todo, error)
	Create(ctx context.Context, req *models.TodoCreateCommand) (*models.Todo, error)
	Update(ctx context.Context, req *models.TodoUpdateCommand) (*models.Todo, error)
	Delete(ctx context.Context, req *models.TodoGetQuery) (*models.Todo, error)
	List(ctx context.Context, req *models.TodoListQuery) (*models.TodoListResults, error)
}

func New(db *sql.DB, ps events.IEvent) ITodo {

	// register event

	return &service{
		DB:     db,
		Events: ps,
	}
}

type service struct {
	DB     *sql.DB
	Events events.IEvent
}

func (c *service) Update(ctx context.Context, req *models.TodoUpdateCommand) (*models.Todo, error) {

	tMod, err := dbmodels.Todos(qm.Where("id=?", req.ID)).One(ctx, c.DB)
	if err != nil {
		return nil, err
	}

	tMod.Description = req.Description
	tMod.Title = req.Title

	_, err = tMod.Update(ctx, c.DB, boil.Infer());
	if err != nil {
		return nil, err
	}

	md := &models.Todo{}

	if err := deepcopier.Copy(tMod).To(md); err != nil {
		return nil, err
	}

	return md, nil

}

func (c *service) Delete(ctx context.Context, req *models.TodoGetQuery) (*models.Todo, error) {

	tMod, err := dbmodels.Todos(qm.Where("id=?", req.ID)).One(ctx, c.DB)
	if err != nil {
		return nil, err
	}

	md := &models.Todo{}
	if err := deepcopier.Copy(tMod).To(md); err != nil {
		return nil, err
	}

	_, err = tMod.Delete(ctx, c.DB)
	if err != nil {
		return nil, err
	}

	return md, nil

}

func (c *service) List(ctx context.Context, req *models.TodoListQuery) (*models.TodoListResults, error) {

	skip := req.Size * (req.Page - 1)

	tdr := &models.TodoListResults{
		CurrentPage: req.Page,
		Results:     []models.Todo{},
	}

	tdRes, err := dbmodels.Todos(qm.Limit(req.Size), qm.Offset(skip), qm.OrderBy("id DESC")).All(ctx, c.DB)
	if err != nil {
		return nil, err
	}

	for _, v := range tdRes {
		md := models.Todo{}

		if err := deepcopier.Copy(v).To(&md); err != nil {
			return nil, err
		}

		tdr.Results = append(tdr.Results, md)
	}

	return tdr, nil

}

func (c *service) Create(ctx context.Context, req *models.TodoCreateCommand) (*models.Todo, error) {

	tMod := dbmodels.Todo{
		Title:       req.Title,
		Description: req.Description,
	}

	if err := tMod.Insert(ctx, c.DB, boil.Infer()); err != nil {
		return nil, err
	}

	md := &models.Todo{}

	if err := deepcopier.Copy(tMod).To(md); err != nil {
		return nil, err
	}

	// synchronous publish event
	if err := c.Events.Publish(ctx, "todo.create", md); err != nil {
		//return
	}

	return md, nil
}

func (c *service) Get(ctx context.Context, query *models.TodoGetQuery) (*models.Todo, error) {

	tMod, err := dbmodels.Todos(qm.Where("id=?", query.ID)).One(ctx, c.DB)
	if err != nil {
		return nil, err
	}

	md := &models.Todo{}

	if err := deepcopier.Copy(tMod).To(md); err != nil {
		return nil, err
	}

	return md, nil
}

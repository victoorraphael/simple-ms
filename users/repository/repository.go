package repository

import (
	"context"
	"github.com/gocraft/dbr/v2"
	"github.com/victoorraphael/simple-ms/adapters/database"
	"github.com/victoorraphael/simple-ms/users/entities"
	"github.com/victoorraphael/simple-ms/utils"
	"time"
)

type UserRepo interface {
	List(ctx context.Context) ([]entities.User, error)
	Get(ctx context.Context, id int64) (entities.User, error)
	Create(ctx context.Context, user entities.User) (int64, error)
	Update(ctx context.Context, user entities.User) error
	Delete(ctx context.Context, id int64) error
	Search(ctx context.Context, filter entities.User) (entities.User, error)
}

func New(provider database.Provider[*dbr.Session]) UserRepo {
	return &Provider{
		conn:  provider,
		table: "users",
	}
}

type Provider struct {
	conn  database.Provider[*dbr.Session]
	table string
}

func (p *Provider) Search(ctx context.Context, filter entities.User) (entities.User, error) {
	query, values, err := utils.BuildSearchQuery(filter)
	if err != nil {
		return filter, err
	}

	var resp entities.User
	errQuery := p.conn.Exec().
		Select("*").
		From(p.table).
		Where(query, values...).
		LoadOneContext(ctx, &resp)
	return resp, errQuery
}

func (p *Provider) List(ctx context.Context) ([]entities.User, error) {
	res := make([]entities.User, 0)
	_, err := p.conn.Exec().
		Select("*").
		From(p.table).
		LoadContext(ctx, &res)
	return res, err
}

func (p *Provider) Get(ctx context.Context, id int64) (entities.User, error) {
	var res entities.User
	err := p.conn.Exec().
		Select("*").
		From(p.table).
		Where("id = ?", id).
		LoadOneContext(ctx, &res)
	return res, err
}

func (p *Provider) Create(ctx context.Context, user entities.User) (int64, error) {
	err := p.conn.Exec().
		InsertInto(p.table).
		Pair("name", user.Name).
		Pair("email", user.Email).
		Pair("phone_number", user.Phone).
		Pair("created_at", time.Now()).
		Pair("updated_at", time.Now()).
		Returning("id").
		LoadContext(ctx, &user.ID)
	return user.ID, err
}

func (p *Provider) Update(ctx context.Context, user entities.User) error {
	_, err := p.conn.Exec().
		Update(p.table).
		Set("name", user.Name).
		Set("email", user.Email).
		Set("phone_number", user.Phone).
		Set("updated_at", time.Now()).
		Where("id = ?", user.ID).
		ExecContext(ctx)
	return err
}

func (p *Provider) Delete(ctx context.Context, id int64) error {
	_, err := p.conn.Exec().
		DeleteFrom(p.table).
		Where("id = ?", id).
		ExecContext(ctx)
	return err
}

package service

import (
	"context"
	"errors"
	"github.com/victoorraphael/simple-ms/users/entities"
	"github.com/victoorraphael/simple-ms/users/repository"
)

type UserService interface {
	List(ctx context.Context) ([]entities.User, error)
	Get(ctx context.Context, id int64) (entities.User, error)
	Create(ctx context.Context, user entities.User) (int64, error)
	Update(ctx context.Context, user entities.User) error
	Delete(ctx context.Context, id int64) error
}

func New(repo repository.UserRepo) UserService {
	return &Provider{
		repo: repo,
	}
}

type Provider struct {
	repo repository.UserRepo
}

func (p *Provider) List(ctx context.Context) ([]entities.User, error) {
	return p.repo.List(ctx)
}

func (p *Provider) Get(ctx context.Context, id int64) (entities.User, error) {
	return p.repo.Get(ctx, id)
}

func (p *Provider) Create(ctx context.Context, user entities.User) (int64, error) {
	_, err := p.repo.Search(ctx, entities.User{Email: user.Email})
	if err == nil {
		return 0, errors.New("email already registered")
	}

	return p.repo.Create(ctx, user)
}

func (p *Provider) Update(ctx context.Context, user entities.User) error {
	_, err := p.repo.Get(ctx, user.ID)
	if err != nil {
		return errors.New("user not found")
	}

	result, err := p.repo.Search(ctx, entities.User{Email: user.Email})
	if err != nil {
		return errors.New("failed to update user")
	}

	if result.ID != user.ID {
		return errors.New("email already registered")
	}

	return p.repo.Update(ctx, user)
}

func (p *Provider) Delete(ctx context.Context, id int64) error {
	_, err := p.repo.Get(ctx, id)
	if err != nil {
		return errors.New("user not found")
	}

	return p.repo.Delete(ctx, id)
}

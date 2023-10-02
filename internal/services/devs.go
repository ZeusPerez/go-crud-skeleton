package services

import (
	"context"

	"github.com/ZeusPerez/go-crud-skeleton/internal/adapters/storage"
	"github.com/ZeusPerez/go-crud-skeleton/internal/models"
)

//go:generate mockery --case underscore --inpackage --name Devs
type Devs interface {
	Get(ctx context.Context, email string) (models.Dev, error)
	Create(ctx context.Context, dev models.Dev) error
	Update(ctx context.Context, dev models.Dev) (models.Dev, error)
	Delete(ctx context.Context, email string) error
}

func NewDevs(storage storage.MySQL) Devs {
	return devs{storage: storage}
}

type devs struct {
	storage storage.MySQL
}

func (d devs) Get(ctx context.Context, email string) (models.Dev, error) {
	return d.storage.Get(ctx, email)
}

func (d devs) Create(ctx context.Context, dev models.Dev) error {
	return d.storage.Create(ctx, dev)
}

func (d devs) Update(ctx context.Context, dev models.Dev) (models.Dev, error) {
	return d.storage.Update(ctx, dev)
}

func (d devs) Delete(ctx context.Context, email string) error {
	return d.storage.Delete(ctx, email)
}

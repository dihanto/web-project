package usecase

import (
	"context"

	"github.com/dihanto/crud-web/entity"
)

type ProductUsecase interface {
	Create(ctx context.Context, product *entity.Product) (err error)
	GetAll(ctx context.Context) (products []entity.Product, err error)
	FindById(ctx context.Context, id int) (products []entity.Product, err error)
	Update(ctx context.Context, product *entity.Product) (err error)
	Delete(ctx context.Context, id int) (err error)
}

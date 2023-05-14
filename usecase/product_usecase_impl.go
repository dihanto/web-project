package usecase

import (
	"context"
	"time"

	"github.com/dihanto/crud-web/entity"
	"github.com/dihanto/crud-web/repository"
)

type productUsecaseImpl struct {
	productRepo    repository.ProductRepository
	contextTimeout time.Duration
}

func NewProductUsecase(pr repository.ProductRepository, timeout time.Duration) ProductUsecase {
	return &productUsecaseImpl{
		productRepo:    pr,
		contextTimeout: timeout,
	}
}

func (pu *productUsecaseImpl) Create(ctx context.Context, product *entity.Product) (err error) {
	ctx, cancel := context.WithTimeout(ctx, pu.contextTimeout)
	defer cancel()
	if err = pu.productRepo.Create(ctx, product); err != nil {
		return
	}
	return
}

func (pu *productUsecaseImpl) GetAll(ctx context.Context) (products []entity.Product, err error) {
	ctx, cancel := context.WithTimeout(ctx, pu.contextTimeout)
	defer cancel()
	products, err = pu.productRepo.GetAll(ctx)
	if err != nil {
		return
	}
	return
}

func (pu *productUsecaseImpl) FindById(ctx context.Context, id int) (products []entity.Product, err error) {
	ctx, cancel := context.WithTimeout(ctx, pu.contextTimeout)
	defer cancel()
	products, err = pu.productRepo.FindById(ctx, id)
	if err != nil {
		return
	}
	return
}

func (pu *productUsecaseImpl) Update(ctx context.Context, product *entity.Product) (err error) {
	ctx, cancel := context.WithTimeout(ctx, pu.contextTimeout)
	defer cancel()
	if err = pu.productRepo.Update(ctx, product); err != nil {
		return
	}
	return
}

func (pu *productUsecaseImpl) Delete(ctx context.Context, id int) (err error) {
	ctx, cancel := context.WithTimeout(ctx, pu.contextTimeout)
	defer cancel()
	if err = pu.productRepo.Delete(ctx, id); err != nil {
		return
	}
	return
}

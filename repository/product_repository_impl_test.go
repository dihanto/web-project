package repository

import (
	"context"
	"fmt"
	"testing"

	"github.com/dihanto/crud-web/app"
	"github.com/dihanto/crud-web/entity"
)

func TestProductCreate(t *testing.T) {

	productRepository := NewProductRepository(app.GetConnection())

	ctx := context.Background()
	product := entity.Product{
		Name:     "Kabel panjanggggg",
		Price:    20000,
		Quantity: 12,
	}
	err := productRepository.Create(ctx, &product)
	if err != nil {
		t.Errorf("failed to create product: %v", err)
	}
}
func TestProductGet(t *testing.T) {
	productRepository := NewProductRepository(app.GetConnection())

	ctx := context.Background()

	result, _ := productRepository.GetAll(ctx)
	fmt.Println(result)
}

func TestFindById(t *testing.T) {
	productRepository := NewProductRepository(app.GetConnection())

	ctx := context.Background()

	result, _ := productRepository.FindById(ctx, 1)
	for _, p := range result {
		fmt.Printf("ID: %d, Name: %s, Price: %0.1f, Quantity: %d\n", p.ID, p.Name, p.Price, p.Quantity)

	}
}
func TestProductUpdate(t *testing.T) {

	productRepository := NewProductRepository(app.GetConnection())

	ctx := context.Background()
	product := entity.Product{
		Name:     "Jauhdfdfhhh",
		Price:    120000,
		Quantity: 21,
		ID:       1,
	}
	err := productRepository.Update(ctx, &product)
	if err != nil {
		t.Errorf("failed to create product: %v", err)
	}
}

func TestProductDelete(t *testing.T) {

	productRepository := NewProductRepository(app.GetConnection())

	ctx := context.Background()
	id := 1

	err := productRepository.Delete(ctx, id)
	if err != nil {
		t.Errorf("failed to delete product: %v", err)
	}
}

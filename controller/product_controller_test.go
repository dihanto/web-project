package controller

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/dihanto/crud-web/entity"
)

type mockProductUsecase struct {
	// mockCreate   func(ctx context.Context, product *entity.Product) error
	// mockGetAll   func(ctx context.Context) ([]entity.Product, error)
	// mockFindById func(ctx context.Context, id int) ([]entity.Product, error)
	// mockUpdate   func(ctx context.Context, product *entity.Product) error
	// mockDelete   func(ctx context.Context, id int) error
}

func (m *mockProductUsecase) Create(ctx context.Context, product *entity.Product) error {
	return nil
}
func (m *mockProductUsecase) GetAll(ctx context.Context) (products []entity.Product, err error) {
	return
}
func (m *mockProductUsecase) FindById(ctx context.Context, id int) (products []entity.Product, err error) {
	return
}
func (m *mockProductUsecase) Update(ctx context.Context, product *entity.Product) error {
	return nil
}
func (m *mockProductUsecase) Delete(ctx context.Context, id int) error {
	return nil
}

func TestProductController_Create(t *testing.T) {
	pc := &ProductController{
		ProductUsecase: &mockProductUsecase{},
	}

	body := strings.NewReader(`{"name": "baju", "price": 1000}`)
	req, err := http.NewRequest("POST", "/products", body)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(pc.Create)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}
	expected := ""
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

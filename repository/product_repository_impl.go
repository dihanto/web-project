package repository

import (
	"context"
	"database/sql"

	"github.com/dihanto/crud-web/entity"
)

type productRepositoryImpl struct {
	Conn *sql.DB
}

func NewProductRepository(conn *sql.DB) ProductRepository {
	return &productRepositoryImpl{conn}
}
func (pri *productRepositoryImpl) Create(ctx context.Context, product *entity.Product) (err error) {

	script := "INSERT INTO products(name, price, quantity) VALUES(?,?,?)"
	stmt, err := pri.Conn.PrepareContext(ctx, script)
	if err != nil {
		return
	}

	result, err := stmt.ExecContext(ctx, product.Name, product.Price, product.Quantity)
	if err != nil {
		return
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return
	}
	product.ID = int64(lastId)
	return
}

func (pri *productRepositoryImpl) GetAll(ctx context.Context) (products []entity.Product, err error) {
	script := `SELECT id, name, price FROM products`
	rows, err := pri.Conn.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		product := entity.Product{}
		err = rows.Scan(
			&product.ID,
			&product.Name,
			&product.Price,
		)
		if err != nil {
			panic(err)
		}
		products = append(products, product)
	}
	return
}

func (pri *productRepositoryImpl) FindById(ctx context.Context, id int) (products []entity.Product, err error) {
	script := `SELECT id, name, price, quantity FROM products WHERE id=?`
	rows, err := pri.Conn.QueryContext(ctx, script, id)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		p := entity.Product{}
		err = rows.Scan(&p.ID, &p.Name, &p.Price, &p.Quantity)
		if err != nil {
			return
		}
		products = append(products, p)
	}
	return
}

func (pri *productRepositoryImpl) Update(ctx context.Context, product *entity.Product) (err error) {
	script := `UPDATE products SET name=?, price=?, quantity=? WHERE id=?`
	stmt, err := pri.Conn.PrepareContext(ctx, script)
	if err != nil {
		return
	}

	_, err = stmt.ExecContext(ctx, product.Name, product.Price, product.Quantity, product.ID)
	if err != nil {
		return
	}
	return
}

func (pri *productRepositoryImpl) Delete(ctx context.Context, id int) (err error) {
	script := `DELETE FROM products WHERE id=?`
	stmt, err := pri.Conn.PrepareContext(ctx, script)
	if err != nil {
		return
	}

	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		return
	}
	return
}

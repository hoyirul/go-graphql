package resolver

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/hoyirul/go-graphql/models"
	_ "github.com/mattn/go-sqlite3"
)

type Resolver struct {
  DB *sql.DB
}

func (r *Resolver) Query_products(ctx context.Context) ([]*models.Product, error) {
	rows, err := r.DB.Query("SELECT id, name, price FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*models.Product
	for rows.Next() {
		var id int
		var name string
		var price float64
		err := rows.Scan(&id, &name, &price)
		if err != nil {
			return nil, err
		}
		products = append(products, &models.Product{
			ID:    fmt.Sprintf("%d", id),
			Name:  name,
			Price: price,
		})
	}
	return products, nil
}

func (r *Resolver) Mutation_createProduct(ctx context.Context, name string, price float64) (*models.Product, error) {
	res, err := r.DB.Exec("INSERT INTO products(name, price) VALUES(?, ?)", name, price)
	if err != nil {
		return nil, err
	}
	id, _ := res.LastInsertId()
	return &models.Product{ID: fmt.Sprintf("%d", id), Name: name, Price: price}, nil
}
package repositories

import (
	"api/src/models"
	"database/sql"
)

type Product struct {
	db *sql.DB
}

func NewRepositoryProduct(db *sql.DB) *Product {
	return &Product{db: db}
}

func (p *Product) InsertNewProduct(product models.Product) error {
	statement, err := p.db.Prepare("insert into product (name, description, price, stock) values(?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(&product.Name, &product.Description, &product.Price, &product.Stock); err != nil {
		return err
	}
	return nil
}

func (p *Product) SearchProductByID(productID uint64) (models.Product, error) {
	row, err := p.db.Query("select * from product where id = ?", productID)
	if err != nil {
		return models.Product{}, err
	}
	defer row.Close()

	var product models.Product
	if row.Next() {
		if err = row.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Stock, &product.AddedIn); err != nil {
			return models.Product{}, err
		}
	}
	return product, nil
}

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

func (p *Product) InsertImage(image models.Images) (int64, error) {
	statement, err := p.db.Prepare("insert into images (filename, path) values(?, ?)")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(&image.Filename, &image.Path)
	if err != nil {
		return 0, err
	}

	imageID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return imageID, nil
}

func (p *Product) InsertNewProduct(product models.Product) error {
	statement, err := p.db.Prepare("insert into product (name, description, price, stock, image) values(?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(&product.Name, &product.Description, &product.Price, &product.Stock, &product.Image.ImageID); err != nil {
		return err
	}
	return nil
}

func (p *Product) SearchProductByID(productID uint64) (models.Product, error) {
	row, err := p.db.Query("select p.*, i.filename, i.path from product p inner join images i on i.image_id = p.product_image_id where p.id = ?", productID)
	if err != nil {
		return models.Product{}, err
	}
	defer row.Close()

	var product models.Product
	if row.Next() {
		if err = row.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Stock, &product.Image.ImageID, &product.AddedIn, &product.Image.Filename, &product.Image.Path); err != nil {
			return models.Product{}, err
		}
	}
	return product, nil
}

func (p *Product) ChangeProductPrice(productID uint64, price float64) error {
	statement, err := p.db.Prepare("update product set price = ? where id = ?")
	if err != nil {
		return nil
	}
	defer statement.Close()

	if _, err := statement.Exec(&productID, &price); err != nil {
		return err
	}

	return nil
}

func (p *Product) Delete(productID uint64) error {
	statement, err := p.db.Prepare("delete from product where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(&productID); err != nil {
		return err
	}
	return nil
}

func (p *Product) UpdateImage(image models.Images) error {
	statement, err := p.db.Prepare("update images set filename = ? and path = ? where image_id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(&image.Filename, &image.Path, &image.ImageID); err != nil {
		return err
	}

	return nil
}

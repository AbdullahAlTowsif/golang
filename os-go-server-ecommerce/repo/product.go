package repo

import (
	"database/sql"
	"ecommerce/domain"
	"ecommerce/product"

	"github.com/jmoiron/sqlx"
)

type ProductRepo interface {
	product.ProductRepo
}

type productRepo struct {
	// productList []*Product
	dbCon *sqlx.DB
}

func NewProductRepo(dbCon *sqlx.DB) ProductRepo {
	return &productRepo{
		dbCon: dbCon,
	}
}

func (r *productRepo) Create(p domain.Product) (*domain.Product, error) {
	query := `
		INSERT INTO products (title, description, price, image_url)
		VALUES ($1, $2, $3, $4)
		RETURNING id;
	`

	row := r.dbCon.QueryRow(query, p.Title, p.Description, p.Price, p.ImgUrl)
	err := row.Scan(&p.ID)

	if err != nil {
		return nil, err
	}

	return &p, nil

}

func (r *productRepo) Get(productId int) (*domain.Product, error) {
	var prd domain.Product
	query := `
	SELECT id, title, description, price, image_url FROM products
	WHERE id = $1
	`

	err := r.dbCon.Get(&prd, query, productId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, nil
	}
	return &prd, nil
}

func (r *productRepo) List(page, limit int64) ([]*domain.Product, error) {
	offset := ((page - 1) * limit) +1
	var prdList []*domain.Product
	query := `
	SELECT id, title, description, price, image_url FROM products LIMIT $1 OFFSET $2
	`

	err := r.dbCon.Select(&prdList, query, limit, offset)
	if err != nil {
		return nil, nil
	}
	return prdList, nil
}

func (r *productRepo) Count() (int64, error) {
	query := `
	SELECT COUNT(*) FROM products
	`
	var count int64
	err := r.dbCon.QueryRow(query).Scan(&count)
	if err != nil {
		return 0, err
	}
	return int64(count), nil
}

func (r *productRepo) Delete(productId int) error {
	query := `
	DELETE FROM products WHERE id = $1
	`
	_, err := r.dbCon.Exec(query, productId)
	if err != nil {
		return err
	}
	return nil
}

func (r *productRepo) Update(product domain.Product) (*domain.Product, error) {
	query := `
	UPDATE products
	SET title=$1, description=$2, price=$3, image_url=$4
	WHERE id=$5
	RETURNING id = $5
	`

	row := r.dbCon.QueryRow(query, product.Title, product.Description, product.Price, product.ImgUrl, product.ID)
	err := row.Err()

	if err != nil {
		return nil, err
	}
	return &product, nil
}

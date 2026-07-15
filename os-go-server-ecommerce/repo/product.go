package repo

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type Product struct {
	ID          int     `json:"id" db:"id"`
	Title       string  `json:"title" db:"title"`
	Description string  `json:"description" db:"description"`
	Price       float64 `json:"price" db:"price"`
	ImgUrl      string  `json:"imageUrl" db:"image_url"`
}

type ProductRepo interface {
	Create(p Product) (*Product, error)
	Get(productId int) (*Product, error)
	List() ([]*Product, error)
	Delete(productId int) error
	Update(p Product) (*Product, error)
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

func (r *productRepo) Create(p Product) (*Product, error) {
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

func (r *productRepo) Get(productId int) (*Product, error) {
	var prd Product
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

func (r *productRepo) List() ([]*Product, error) {
	var prdList []*Product
	query := `
	SELECT id, title, description, price, image_url FROM products
	`

	err := r.dbCon.Select(&prdList, query)
	if err != nil {
		return nil, nil
	}
	return prdList, nil
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

func (r *productRepo) Update(product Product) (*Product, error) {
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

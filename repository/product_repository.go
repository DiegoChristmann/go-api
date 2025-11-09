package repository

import (
	"database/sql"
	"fmt"
	"go-api/model"
)

type ProductRepository struct {
	connection *sql.DB
}

func NewProductRepository(connection *sql.DB) ProductRepository {
	return ProductRepository{
		connection: connection,
	}
}

func (pr *ProductRepository) GetProducts() ([]model.Product, error) {
	query := "SELECT id, product_name, price FROM product"

	fmt.Printf("Executando query: %s\n", query)
	fmt.Printf("Verificando conexão: %v\n", pr.connection != nil)

	rows, err := pr.connection.Query(query)
	if err != nil {
		fmt.Printf("Erro ao executar query GetProducts: %v\n", err)
		return []model.Product{}, fmt.Errorf("erro ao buscar produtos: %w", err)
	}
	defer rows.Close()

	var productList []model.Product
	var count int

	for rows.Next() {
		count++
		var productObj model.Product
		err = rows.Scan(
			&productObj.ID,
			&productObj.Name,
			&productObj.Price)

		if err != nil {
			fmt.Printf("Erro ao escanear produto (linha %d): %v\n", count, err)
			return []model.Product{}, fmt.Errorf("erro ao processar produto: %w", err)
		}

		fmt.Printf("Produto encontrado: ID=%d, Nome=%s, Preço=%.2f\n",
			productObj.ID, productObj.Name, productObj.Price)
		productList = append(productList, productObj)
	}

	// Verificar erros do loop (importante!)
	if err = rows.Err(); err != nil {
		fmt.Printf("Erro durante iteração: %v\n", err)
		return []model.Product{}, fmt.Errorf("erro ao iterar produtos: %w", err)
	}

	// Se não houver produtos, retorna array vazio (não nil)
	if productList == nil {
		productList = []model.Product{}
	}

	fmt.Printf("GetProducts: %d produtos encontrados\n", len(productList))
	return productList, nil
}

func (pr *ProductRepository) CreateProduct(product model.Product) (int, error) {
	var id int
	query, err := pr.connection.Prepare("INSERT INTO product" +
		"(product_name, price) VALUES ($1, $2) RETURNING id")

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	err = query.QueryRow(product.Name, product.Price).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	query.Close()

	return id, nil
}

func (pr *ProductRepository) GetProductById(id_product int) (*model.Product, error) {

	query, err := pr.connection.Prepare("SELECT id, product_name, price FROM product WHERE id = $1")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var produto model.Product

	err = query.QueryRow(id_product).Scan(
		&produto.ID,
		&produto.Name,
		&produto.Price,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	query.Close()

	return &produto, nil

}

func (pr *ProductRepository) DeleteProduct(id_product int) error {
	query, err := pr.connection.Prepare("DELETE FROM product WHERE id = $1")
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer query.Close()

	result, err := query.Exec(id_product)
	if err != nil {
		fmt.Println(err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("produto com ID %d não encontrado", id_product)
	}

	return nil
}

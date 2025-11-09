package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func ConnectDB() (*sql.DB, error) {
	// Usa variáveis de ambiente ou valores padrão
	// No Docker: DB_HOST=go_db (nome do serviço)
	// Localmente: DB_HOST=localhost
	host := getEnv("DB_HOST", "go_db")
	port := getEnv("DB_PORT", "5432")
	user := getEnv("DB_USER", "postgres")
	password := getEnv("DB_PASSWORD", "1234")
	dbname := getEnv("DB_NAME", "postgres")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping database at %s:%s: %w", host, port, err)
	}

	fmt.Printf("Connected to database %s at %s:%s\n", dbname, host, port)

	// Criar tabela se não existir
	err = Migrate(db)
	if err != nil {
		return nil, fmt.Errorf("failed to migrate database: %w", err)
	}

	return db, nil
}

// Migrate cria as tabelas necessárias no banco de dados
func Migrate(db *sql.DB) error {
	query := `
		CREATE TABLE IF NOT EXISTS product (
			id SERIAL PRIMARY KEY,
			product_name VARCHAR(255) NOT NULL,
			price DECIMAL(10, 2) NOT NULL
		)
	`

	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to create product table: %w", err)
	}

	fmt.Println("Product table migrated successfully")
	return nil
}

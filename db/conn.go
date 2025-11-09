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

	// LOG DE DIAGNÓSTICO
	fmt.Printf("🔌 Tentando conectar ao banco:\n")
	fmt.Printf("   Host: %s\n", host)
	fmt.Printf("   Port: %s\n", port)
	fmt.Printf("   User: %s\n", user)
	fmt.Printf("   Database: %s\n", dbname)
	fmt.Printf("   Connection String: host=%s port=%s user=%s dbname=%s\n", host, port, user, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, fmt.Errorf("erro ao abrir conexão: %w", err)
	}

	// Configurar pool de conexões
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping database at %s:%s: %w", host, port, err)
	}

	fmt.Printf("Connected to database %s at %s:%s\n", dbname, host, port)

	// VERIFICAR QUAL BANCO ESTAMOS USANDO
	var currentDb string
	err = db.QueryRow("SELECT current_database()").Scan(&currentDb)
	if err == nil {
		fmt.Printf("Banco atual: %s\n", currentDb)
	}

	var currentUser string
	err = db.QueryRow("SELECT current_user").Scan(&currentUser)
	if err == nil {
		fmt.Printf("Usuário atual: %s\n", currentUser)
	}

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

	// Verificar quantos produtos existem
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM product").Scan(&count)
	if err != nil {
		fmt.Printf("Aviso: Não foi possível contar produtos: %v\n", err)
	} else {
		fmt.Printf("Total de produtos no banco: %d\n", count)
		if count > 0 {
			// Listar os primeiros produtos
			rows, err := db.Query("SELECT id, product_name, price FROM product LIMIT 5")
			if err == nil {
				defer rows.Close()
				fmt.Println("Produtos encontrados:")
				for rows.Next() {
					var id int
					var name string
					var price float64
					if err := rows.Scan(&id, &name, &price); err == nil {
						fmt.Printf("   - ID: %d, Nome: %s, Preço: %.2f\n", id, name, price)
					}
				}
			}
		}
	}

	return nil
}

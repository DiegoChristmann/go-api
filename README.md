# API Go - Gerenciamento de Produtos

API REST desenvolvida em Go (Gin) para gerenciamento de produtos com arquitetura em camadas (Controller, UseCase, Repository).

## 🚀 Tecnologias

- **Go 1.23**
- **Gin Web Framework**
- **PostgreSQL 12**
- **Docker & Docker Compose**

## 📋 Pré-requisitos

- Go 1.23 ou superior
- Docker e Docker Compose
- PostgreSQL (se rodar localmente)

## 🏗️ Arquitetura

```
├── cmd/           # Ponto de entrada da aplicação
├── controller/    # Camada de controle (handlers HTTP)
├── usecase/       # Camada de casos de uso (lógica de negócio)
├── repository/    # Camada de repositório (acesso a dados)
├── model/         # Modelos de dados
└── db/            # Configuração e migração do banco de dados
```

## 🚀 Como executar

### Usando Docker Compose (Recomendado)

1. Clone o repositório:
```bash
git clone <seu-repositorio>
cd API1
```

2. Execute com Docker Compose:
```bash
docker compose up -d
```

3. A API estará disponível em: `http://localhost:8000`

### Executando localmente

1. Certifique-se de que o PostgreSQL está rodando

2. Configure as variáveis de ambiente:
```bash
export DB_HOST=localhost
export DB_PORT=5432
export DB_USER=postgres
export DB_PASSWORD=1234
export DB_NAME=postgres
```

3. Execute a aplicação:
```bash
go run cmd/main.go
```

## 📡 Endpoints

### Health Check
- **GET** `/ping` - Verifica se a API está funcionando

### Produtos
- **GET** `/products` - Lista todos os produtos
- **GET** `/product/:productId` - Busca um produto por ID
- **POST** `/product` - Cria um novo produto

#### Exemplo de criação de produto:
```json
POST /product
Content-Type: application/json

{
  "name": "Produto Exemplo",
  "price": 29.90
}
```

## 🗄️ Banco de Dados

A tabela `product` é criada automaticamente quando a aplicação inicia através de migrations.

### Estrutura da tabela:
```sql
CREATE TABLE product (
    id SERIAL PRIMARY KEY,
    product_name VARCHAR(255) NOT NULL,
    price DECIMAL(10, 2) NOT NULL
);
```

## 🔧 Variáveis de Ambiente

| Variável | Descrição | Valor Padrão |
|----------|-----------|--------------|
| `DB_HOST` | Host do PostgreSQL | `go_db` (Docker) / `localhost` (local) |
| `DB_PORT` | Porta do PostgreSQL | `5432` |
| `DB_USER` | Usuário do PostgreSQL | `postgres` |
| `DB_PASSWORD` | Senha do PostgreSQL | `1234` |
| `DB_NAME` | Nome do banco de dados | `postgres` |

## 📦 Estrutura do Projeto

```
.
├── cmd/
│   └── main.go              # Ponto de entrada
├── controller/
│   └── product_controller.go # Handlers HTTP
├── usecase/
│   └── product_usecase.go   # Lógica de negócio
├── repository/
│   └── product_repository.go # Acesso aos dados
├── model/
│   ├── product.go           # Modelo Product
│   └── response.go          # Modelo Response
├── db/
│   └── conn.go              # Conexão e migrations
├── Dockerfile
├── docker-compose.yaml
├── go.mod
└── README.md
```

## 🧪 Testando a API

### Exemplo com curl:

```bash
# Health check
curl http://localhost:8000/ping

# Listar produtos
curl http://localhost:8000/products

# Criar produto
curl -X POST http://localhost:8000/product \
  -H "Content-Type: application/json" \
  -d '{"name":"Produto Teste","price":19.90}'

# Buscar produto por ID
curl http://localhost:8000/product/1
```

## 🛠️ Desenvolvimento

### Rodar em modo desenvolvimento:
```bash
go run cmd/main.go
```

### Build da aplicação:
```bash
go build -o bin/api cmd/main.go
```

### Build Docker:
```bash
docker build -t go-api .
```

## 📝 Licença

Este projeto está sob a licença MIT.

## 👤 Autor

Seu Nome

---

Desenvolvido com ❤️ usando Go


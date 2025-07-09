# 🎮 Game API - Projeto de Estudos em Go

Este projeto é uma **API REST desenvolvida em Go utilizando o framework Gin**.  
Ela foi criada com fins educativos como **meu primeiro contato prático com desenvolvimento de APIs em Go (Golang)**.

## 🚀 Sobre o projeto

A API permite realizar operações CRUD sobre uma lista de jogos (games), utilizando:

- 📦 Go 1.21+
- 🧩 Gin Web Framework
- 🐬 MySQL como banco de dados
- 🔍 SQL puro via `database/sql`
- 📁 Estrutura modular com `handler`, `repository`, `model`, e `request`

## 🔧 Rotas da API

| Método | Rota               | Descrição                     |
|--------|--------------------|-------------------------------|
| GET    | `/api/v1/games`    | Lista todos os jogos          |
| POST   | `/api/v1/games`    | Cria um novo jogo             |
| PUT    | `/api/v1/games/:id`| Atualiza um jogo existente    |
| DELETE | `/api/v1/games/:id`| Remove um jogo                |

## 📋 Exemplo de JSON

```json
{
  "id": 1,
  "name": "The Legend of Zelda",
  "publish_date": "2023-01-01T00:00:00Z",
  "beated": true,
  "beated_at": "2023-03-01T00:00:00Z",
  "rating": 9
}
```
## ⚙️ Configuração e execução

1. Clone o repositório:
```
git clone https://github.com/feribeirods/api-go-gin.git
cd game-api-go
```
2. Crie um arquivo .env com as variáveis de conexão:
```
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASS=sua_senha
DB_NAME=nome_do_banco
```

3. Instale as dependências:
```
go mod tidy
```
4. Execute a aplicação:
```
go run main.go
```
## 📌 Observações
Este projeto é puramente educativo, sem foco em segurança, escalabilidade ou performance.

Foi desenvolvido com o objetivo de aprender Go, Gin e práticas RESTful.

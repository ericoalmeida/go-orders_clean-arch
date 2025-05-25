# 🧾 Order API - Microserviço em Go

Este projeto é um **microserviço para gerenciamento de pedidos (orders)**, desenvolvido em **Go** com a arquitetura **Clean Architecture**. Ele expõe um endpoint `/orders` para listagem de pedidos e utiliza banco de dados PostgreSQL com **migrations e seeds automáticos**.

---

## 🚀 Objetivo

O objetivo desta aplicação é fornecer uma base sólida para construção de microsserviços em Go, aplicando boas práticas como:

- Clean Architecture
- Injeção de dependência
- Separação clara de responsabilidades (domain/usecase/interface)
- Provisionamento completo com Docker

---

## 🛠 Tecnologias Utilizadas

- [Go](https://golang.org/)
- [PostgreSQL](https://www.postgresql.org/)
- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)
- [golang-migrate](https://github.com/golang-migrate/migrate) — Migrations
- [gofakeit](https://github.com/brianvoe/gofakeit) — Seed com dados aleatórios


---

## 📦 Subindo com Docker Compose

### Pré-requisitos

- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

### Passos

1. Clone o repositório:
```bash
   git clone https://github.com/seu-usuario/order-api.git
   cd order-api
```

2. Crie um arquivo .env com as variáveis de ambiente necessárias:
```bash
DATABASE_URL=
PORT=8080
```

3. Suba os containers:
```bash
docker-compose up --build
```

4. Acesse a API:
```bash
curl -X GET http://localhost:8080/orders
```

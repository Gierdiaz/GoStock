# GoStock

GoStock is an API developed in Go using the MVC architecture, applying SOLID principles, and using Docker for environment management.

## Technologies Used

- [Gin](https://github.com/gin-gonic/gin) - Web Framework
- [GORM](https://gorm.io/) - ORM
- [Docker](https://www.docker.com/) - Container Management
- [PostgreSQL](https://www.postgresql.org/) - Relational Database
- [Faker](https://github.com/jaswdr/faker) - Fake Data Generation

#### Rotas

## Routes

### Users

- `POST /api/register`: Register a new user
- `POST /api/login`: User login

### Products

- `GET /api/products`: List all products (authentication required)
- `POST /api/products`: Create a new product (authentication required)

## Seeds

To populate the database with initial data, run:

```sh
docker-compose run api go run internal/database/seeds/seed.go

# External Consultation REST API

This project aims to develop a RESTful API for external consultations and return only expected values.


# How to run
## Requirements
- Docker
- Golang
- [Migrate](https://pkg.go.dev/github.com/golang-migrate/migrate/v4)

### Step 1
Run the docker compose to have a local postgres
```bash
docker-compose up -d
```

### Step 2
Run migrations to have a functional DB
```
POSTGRESQL_URL=postgres://fr-user:123456@localhost:5432/fr?sslmode=disable make migrate_up
```

> If you have different db please replace credentials above

### Step 3
You need to have valid environments
The project has the file `.env.example` you can use it to create the .env with valid environments

### Step 4
Now you just have to run the project
```bash
go run main.go
```

## Testing
The file insomnia.json are in the project, is just to import and use environment dev. 


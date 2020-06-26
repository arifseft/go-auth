<h1 align="center">Golang AUTH</h1>

## How to run

There are two ways to run this application, with docker or without docker

```bash
# running with docker

# copy .env
cp .env.example .env

# running in development mode, you can use live-reload when safe file
make run-local

# remove production container
make down-local


# running in production image
make run-production
docker logs --tail=100 -f golang_example_production # monitoring production container
docker exec -it golang_example_production sh # access bash on production container

# remove production container
make down-production
```

```bash
# running in local/without docker

# copy .env
cp .env.example .env
make install
make run
```

## Run tests

```bash
make test
```

## Run lint

```bash
make lint
```

## Run migration

```bash
make migrate
```

## Run seeder

```bash
# running migration required
make seed
```

## Endpoints

```bash
# Get All User - GET http://127.0.0.1:8080/users

# Get User - GET http://127.0.0.1:8080/user/1

# Create User - POST http://127.0.0.1:8080/user
Request :
{
    "name": "Arif",
    "age": 27,
    "email": "arif@sefrianto.com",
    "address": "Jl. kikuk",
    "password": "1234"
}

# Login - POST http://127.0.0.1:8080/login
Request :
{
	"username": "arif@sefrianto.com",
	"password": "1234"
}

# Logout - GET http://127.0.0.1:8080/logout
Header :
Authorization : Bearer <access_token>

# Update User - PATCH http://127.0.0.1:8080/user/1
Request :
{
	"name": "MArifS",
	"age": 27,
	"email": "arif@sefrianto.com",
	"address": "Jl. kikuk",
	"password": "1234"
}

# Delete User - DELETE http://127.0.0.1:8080/user/21

```

## Postman Collection

```
https://www.getpostman.com/collections/b8c9e65d02c9d4210c4d
```

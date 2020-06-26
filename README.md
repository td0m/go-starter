# Go-Starter

## Starting the application

### Development

First, you'll need to start the postgres database:
```bash
sudo docker-compose -f docker-compose-dev.yml up
```

To run the server, simply `cd` into the server directory and run:
```bash
go run cmd/server/main.go
```

The server will run on port `8080`, and postgres database should be at `5432`.

You can also connect to it with `psql`: 

```bash
sudo docker exec -it go-starter_dev_db_1 psql -U admin
```

#### Cleaning the database

If you need to reinit the database with `init-db.sql`, you'll need to clean it:
```
sudo docker-compose -f docker-compose-dev.yml down --volumes
```

### Production

In development, the `Dockerfile` builds the server and starts it on port `80`. To start both the server and postgres database, run:
```bash
sudo docker-compose up
```


## Structure

The structure of this projects strictly follows Clean Architecture rules.

The entry point of the server is located in:
```
cmd
 |-server
    |-main.go
```

This is where you should connect to the database (and remember to close it). It's also the place where you should get all your environmental variables for use later.

## Domain

Domain is where you should define your `struct`s and `interface`s.

pkg/domain/user.go:
```go
package domain

type User struct {
	Email        string `json:"email"`
	PasswordHash string `json:"-"`
}

type UserRepo interface {
	GetUserByEmail(email string) (*User, error)
	CreateUser(email, passwordHash string) error
	UserExists(email string) bool
}

type UserService interface {
	GetUser(email string) (*User, error)
	Register(email string, password string) (*User, error)
	Login(email string, password string) (*User, error)
}

```

## Repository

A repository is something that saves/gathers data, such as a database. It uses models from the `domain` such as `User`.

```go
func (r *repo) GetUserByEmail(email string) (*domain.User, error) {
	u := domain.User{}
	err := r.Conn.
		QueryRow("SELECT email,password_hash FROM users WHERE email=$1", email).
		Scan(&u.Email, &u.PasswordHash)
	if err != nil {
		return nil, domain.ErrNotFound
	}
	return &u, nil
}
```

## Service

A service can use the repositories and models. It performs more real life use case operations such as `Login` that might require some logic.

```go
func (s *service) Login(email string, password string) (*domain.User, error) {
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	if !matchesHash(password, user.PasswordHash) {
		return nil, domain.ErrCredentials
	}
	return user, nil
}
```

## Delivery

Delivery is a vague name I decided to use for "endpoints" such as REST API or gRPC. It's a set of functions that use `services` to interface with the endpoints. Here is an example:

```go
func (h *handler) getUser(w http.ResponseWriter, r *http.Request) {
	email := mux.Vars(r)["email"]
	user, err := h.userService.GetUser(email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
```
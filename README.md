# Go RESTful API

This is a simple RESTful API written in go.

## Dependencies

- [gin](https://github.com/gin-gonic/gin)
- [gorm](https://pkg.go.dev/gorm.io/gorm)
- [gorm/postgres](https://pkg.go.dev/gorm.io/driver/postgres)
- [godotenv](https://pkg.go.dev/github.com/joho/godotenv)
- [golang.org/x/crypto/bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt)

## Directories

- `/api/`: API endpoints
- `/helpers/`: Helper functions
- `/models/`: DB models

## Example endpoint

`/user`

With this endpoint is possible to:

- **List all users**: `GET /user`
- **List a certain user**: `GET /user/:id`
- **Create a user**: `POST /user`
- **Update a user**: `PUT /user/:id`
- **Delete a user**: `DELETE /user/:id`

package usersApi

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/iafoosball/users-service/restapi/operations"
	"github.com/iafoosball/users-service/users/usersImpl"
)

func CreateUser() func(params operations.PostUsersParams) middleware.Responder {
	return func(params operations.PostUsersParams) middleware.Responder {
		ops, _ := usersImpl.CreateUser(*params.Body)
		return ops
	}
}

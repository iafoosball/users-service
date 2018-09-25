package usersImpl

import (
	"github.com/arangodb/go-driver"
	"github.com/go-openapi/runtime/middleware"
	"github.com/iafoosball/users-service/models"
	"github.com/iafoosball/users-service/restapi/operations"
)

func GetUserByID() func(params operations.GetUsersUserIDParams) middleware.Responder {
	return func(params operations.GetUsersUserIDParams) middleware.Responder {
		//Log the user
		var u = models.User{}
		_, _ = usersCol.ReadDocument(nil, params.UserID, &u)
		return operations.NewGetUsersUserIDOK().WithPayload(&u)
	}
}

func CreateUser(u models.User) (*operations.PostUsersOK, driver.DocumentMeta) {
	meta, _ := usersCol.CreateDocument(nil, &u)
	return operations.NewPostUsersOK(), meta
}

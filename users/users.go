package users

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/iafoosball/users-service/models"
	"github.com/iafoosball/users-service/restapi/operations"
	"log"
)

func GetUserByID() func(params operations.GetUsersUserIDParams) middleware.Responder {
	return func(params operations.GetUsersUserIDParams) middleware.Responder {
		//Log the user
		var u = models.User{}
		_, _ = usersCol.ReadDocument(nil, params.UserID, &u)
		log.Println(u)
		log.Println("users/" + params.UserID)
		_, _ = usersCol.ReadDocument(nil, "users/"+params.UserID, &u)
		log.Println(u)
		return operations.NewGetUsersUserIDOK().WithPayload(&u)
	}
}

func CreateUser() func(params operations.PostUsersParams) middleware.Responder {
	return func(params operations.PostUsersParams) middleware.Responder {
		_, _ = usersCol.CreateDocument(nil, &params.Body)
		return operations.NewPostUsersOK()
	}
}

package users

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/iafoosball/users-service/models"
	"github.com/iafoosball/users-service/restapi/operations"
	"log"
)

// GetUserByID returns user document identified with userID in path parameter
func GetUserByID() func(params operations.GetUsersUserIDParams) middleware.Responder {
	return func(params operations.GetUsersUserIDParams) middleware.Responder {
		//Log the user
		var u = models.User{}
		_, _ = usersCol.ReadDocument(nil, params.UserID, &u)
		_, _ = usersCol.ReadDocument(nil, "users/"+params.UserID, &u)
		return operations.NewGetUsersUserIDOK().WithPayload(&u)
	}
}

// CreateUser creates User document in ArangoDB and registers user in Kong
func CreateUser() func(params operations.PostUsersParams) middleware.Responder {
	return func(params operations.PostUsersParams) middleware.Responder {
		_, err := usersCol.CreateDocument(nil, &params.Body)
		if err != nil {
			log.Fatal(err)
		}
		return operations.NewPostUsersOK()
	}
}

// UpdateUserByID updates user data with content's of payload parameter entries
func UpdateUserByID() func(params operations.PutUsersUserIDParams) middleware.Responder {
	return func(params operations.PutUsersUserIDParams) middleware.Responder {
		_, err := usersCol.UpdateDocument(nil, params.Body.UserID, &params.Body)
		if err != nil {
			log.Fatal(err)
		}
		return operations.NewPutUsersUserIDOK()
	}
}

// DeleteUserByID deletes user of specified identifier. Operation is irreversible
func DeleteUserByID() func(params operations.DeleteUsersUserIDParams) middleware.Responder {
	return func(params operations.DeleteUsersUserIDParams) middleware.Responder {
		_, err := usersCol.RemoveDocument(nil, params.UserID)

		if err != nil {
			log.Fatal(err)
		}
		return operations.NewDeleteFriendsFriendshipIDOK()
	}
}

package users

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/iafoosball/users-service/models"
	"github.com/iafoosball/users-service/restapi/operations"
	"github.com/kevholditch/gokong"
	"log"
)

func kongClient() *gokong.KongAdminClient {
	config := &gokong.Config{
		HostAddress: "http://kong:8001",
		Username:    "admin",
		Password:    "adminadminadmin",
	}
	return gokong.NewClient(config)
}

// GetUserByID returns user document identified with userID in path parameter
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

// CreateUser creates User document in ArangoDB and registers user in Kong
func CreateUser() func(params operations.PostUsersParams) middleware.Responder {
	return func(params operations.PostUsersParams) middleware.Responder {
		meta, err := usersCol.CreateDocument(nil, &params.Body)
		if err != nil {
			log.Fatal(err)
		}
		kongClient().Consumers().Create(
			&gokong.ConsumerRequest{
				Username: meta.UserID,
			},
		)
		return operations.NewPostUsersOK()
	}
}

// UpdateUserByID updates user data with content's of payload parameter entries
func UpdateUserByID() func(params operations.PutUsersUserID) middleware.Responder {
	return func(params operations.PutUsersUserIDParams) middleware.Responder {
		_, err := usersCol.UpdateDocument(nil, &params.UserID, &params.Body)
		if err != nil {
			log.Fatal(err)
		}
		return operations.NewPutUsersUserIDOK()
	}
}

// DeleteUserByID deletes user of specified identifier. Operation is irreversible
func DeleteUserByID() func(params operations.DeleteUsersUserIDParams) middleware.Responder {
	return func(params operations.DeleteUsersUserIDParams) middleware.Responder {
		_, err := usersCol.DeleteUserByID(nil, &params.UserID)
		if err != nil {
			log.Fatal(err)
		}
		kongClient().Consumers().DeleteById(params.UserID)
		return operations.NewDeleteFriendsFriendshipIDOK()
	}
}

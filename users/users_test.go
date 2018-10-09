package users

import (
	"github.com/IAFoosball/public-server/iaf/users"
	"github.com/iafoosball/users-service/models"
	"github.com/iafoosball/users-service/restapi/operations"
	"testing"
)

func TestCreateUser(t *testing.T) {
	var user operations.PostUsersHandlerFunc
	user(users.CreateUser())
}

package users

import (
	"testing"
	"github.com/IAFoosball/public-server/iaf/users"
	"github.com/iafoosball/users-service/models"
	"github.com/iafoosball/users-service/restapi/operations"
)

func TestCreateUser(t *testing.T) {
	var user operations.PostUsersHandlerFunc
	user(users.CreateUser())
}

package users

import (
	"bytes"
	"encoding/json"
	"github.com/iafoosball/users-service/models"
	"log"
	"net/http"
	"testing"
)

func TestCreateUser(t *testing.T) {
	u := models.User{
		Username: "Username",
		Lastname: "Lastname",
	}
	json, _ := json.Marshal(&u)
	if resp, err := http.Post("http://localhost:9001/users/createUser", "application/json", bytes.NewReader(json)); err != nil || http.StatusOK != resp.StatusCode {
		log.Fatal("User could not be created")
	}

}

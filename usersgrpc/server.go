package usersgrpc

import (
	"context"
	"encoding/json"
	"github.com/arangodb/go-driver"
)

type usersServiceServer struct {
	// database struct
	usersCol      driver.Collection
}

// Code 0 is http 200, 13 is http 500
func (s *usersServiceServer) CreateUser(ctx context.Context, body *User) (*UsersReply, error) {
	m, err := json.Marshal(body)
	if err != nil {
		return &UsersReply{Code:13}, err
	}
	_, err = s.usersCol.CreateDocument(nil, &m)
	if err != nil {
		return &UsersReply{Code:13}, err
	}
	return &UsersReply{Code:0}, nil
}

func (s *usersServiceServer) ReadUser(ctx context.Context, req *UsersRequest) (*User, error) {
	var u User
	_, err := s.usersCol.ReadDocument(nil, req.UserID, &u)
	if err != nil {
		return nil, err
	}
	_, err = s.usersCol.ReadDocument(nil, "users/"+req.UserID, &u)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (s *usersServiceServer) UpdateUser(ctx context.Context, body *User) (*UsersReply, error) {
	m, err := json.Marshal(body)
	if err != nil {
		return &UsersReply{Code:13}, err
	}
	_, err = s.usersCol.UpdateDocument(nil, body.UserID, &m)
	if err != nil {
		return &UsersReply{Code:13}, err
	}
	return &UsersReply{Code:0}, nil
}

func (s *usersServiceServer) DeleteUser(ctx context.Context, req *UsersRequest) (*UsersReply, error) {
	_, err := s.usersCol.RemoveDocument(nil, req.UserID)
	if err != nil {
		return &UsersReply{Code:13}, err
	}
	return &UsersReply{Code:0}, nil
}

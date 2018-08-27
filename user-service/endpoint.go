package user_service

import "github.com/go-kit/kit/endpoint"
import (
	"context"
	"errors"
)

type NewUserRequest struct {
	Name     string
	Email    string
	Password string
}

type NewUserResponse struct {
	Id  string
	Err string
}

type GetUserByEmailRequest struct {
	Email string
}

type GetUserByEmailResponse struct {
	User User
	Err  string
}

func MakeNewUserEndpoint(u Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(NewUserRequest)
		id, err := u.NewUser(ctx, req.Name, req.Email, req.Password)
		return NewUserResponse{Id: id}, nil
	}
}

func MakeGetUserByEmailEndpoint(u Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetUserByEmailRequest)
		user, err := u.GetUserByEmail(ctx, req.Email)
		return GetUserByEmailResponse{User: user}, nil
	}
}

type Endpoints struct {
	NewUserEndpoint        endpoint.Endpoint
	GetUserByEmailEndpoint endpoint.Endpoint
}

func (e Endpoints) NewUser(ctx context.Context, name string, email string, password string) (string, error) {
	req := NewUserRequest{Name: name, Email: email, Password: password}
	resp, err := e.NewUserEndpoint(ctx, req)
	if err != nil {
		return "", errors.New("create user error")
	}
	userResp := resp.(NewUserResponse)
	return userResp.Id, nil
}

func (e Endpoints) GetUserByEmail(ctx context.Context, email string) (User, error) {
	req := GetUserByEmailRequest{Email: email}
	resp, err := e.GetUserByEmailEndpoint(ctx, req)
	if err != nil {
		return User{}, err
	}
	userResp := resp.(GetUserByEmailResponse)
	return userResp.User, nil
}

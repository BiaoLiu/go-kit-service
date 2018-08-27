package user_service

import "context"

type Service interface {
	NewUser(ctx context.Context, name string, email string, password string) (string, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
}

type UserService struct {
}

func (UserService) NewUser(ctx context.Context, name string, email string, password string) (string, error) {
	return "1", nil
}

func (UserService) GetUserByEmail(ctx context.Context, email string) (User, error) {
	return User{Id: "1", Email: "4513@qq.com"}, nil
}

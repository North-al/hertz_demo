package service

import (
	"context"
	user "github.com/North-al/hertz_demo/mall/app/user/kitex_gen/user"
)

type RegisterService struct {
	ctx context.Context
} // NewRegisterService new RegisterService
func NewRegisterService(ctx context.Context) *RegisterService {
	return &RegisterService{ctx: ctx}
}

// Run create note info
func (s *RegisterService) Run(req *user.RegisterRequest) (resp *user.RegisterResponse, err error) {
	// Finish your business logic.

	return
}

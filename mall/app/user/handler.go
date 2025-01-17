package main

import (
	"context"
	user "github.com/North-al/hertz_demo/mall/app/user/kitex_gen/user"
	"github.com/North-al/hertz_demo/mall/app/user/biz/service"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *user.RegisterRequest) (resp *user.RegisterResponse, err error) {
	resp, err = service.NewRegisterService(ctx).Run(req)

	return resp, err
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *user.LoginRequest) (resp *user.LoginResponse, err error) {
	resp, err = service.NewLoginService(ctx).Run(req)

	return resp, err
}

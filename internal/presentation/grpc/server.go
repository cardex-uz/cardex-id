package grpc

import (
	ssov1 "cardex-id/pkg/grpc/sso"
	"context"
	"google.golang.org/grpc"
)

type serverAPI struct {
	ssov1.UnimplementedSSOServer
	auth Auth
}

type Auth interface {
	ValidateUsername(grpc.BidiStreamingServer[ssov1.ValidateUsernameReq, ssov1.ValidateUsernameRes]) error
	ValidatePhone(ctx context.Context, req *ssov1.ValidatePhoneReq) (*ssov1.ValidatePhoneRes, error)
	ConfirmSubmit(ctx context.Context, req *ssov1.ConfirmSubmitReq) (*ssov1.ConfirmSubmitRes, error)
	Confirm(ctx context.Context, req *ssov1.ConfirmReq) (*ssov1.ConfirmRes, error)
	FindAccounts(ctx context.Context, req *ssov1.FindAccountsReq) (*ssov1.FindAccountsRes, error)
	SignIn(ctx context.Context, req *ssov1.SignInReq) (*ssov1.SignInRes, error)
	SignUp(ctx context.Context, req *ssov1.SignUpReq) (*ssov1.SignUpRes, error)
	SignOut(ctx context.Context, req *ssov1.SignOutReq) (*ssov1.SignOutRes, error)
	GetMe(ctx context.Context, req *ssov1.GetMeReq) (*ssov1.GetMeRes, error)
}

func Register(gRPCServer *grpc.Server, auth Auth) {
	ssov1.RegisterSSOServer(gRPCServer, &serverAPI{auth: auth})
}

func (s *serverAPI) ValidateUsername(grpc.BidiStreamingServer[ssov1.ValidateUsernameReq, ssov1.ValidateUsernameRes]) error {
	return nil
}

func (s *serverAPI) ValidatePhone(ctx context.Context, req *ssov1.ValidatePhoneReq) (*ssov1.ValidatePhoneRes, error) {
	return &ssov1.ValidatePhoneRes{}, nil
}

func (s *serverAPI) ConfirmSubmit(ctx context.Context, req *ssov1.ConfirmSubmitReq) (*ssov1.ConfirmSubmitRes, error) {
	return &ssov1.ConfirmSubmitRes{}, nil
}

func (s *serverAPI) Confirm(ctx context.Context, req *ssov1.ConfirmReq) (*ssov1.ConfirmRes, error) {
	return &ssov1.ConfirmRes{}, nil
}

func (s *serverAPI) FindAccounts(ctx context.Context, req *ssov1.FindAccountsReq) (*ssov1.FindAccountsRes, error) {
	return &ssov1.FindAccountsRes{}, nil
}

func (s *serverAPI) SignIn(ctx context.Context, req *ssov1.SignInReq) (*ssov1.SignInRes, error) {
	return &ssov1.SignInRes{}, nil
}

func (s *serverAPI) SignUp(ctx context.Context, req *ssov1.SignUpReq) (*ssov1.SignUpRes, error) {
	return &ssov1.SignUpRes{}, nil
}

func (s *serverAPI) SignOut(ctx context.Context, req *ssov1.SignOutReq) (*ssov1.SignOutRes, error) {
	return &ssov1.SignOutRes{}, nil
}

func (s *serverAPI) GetMe(ctx context.Context, req *ssov1.GetMeReq) (*ssov1.GetMeRes, error) {
	return &ssov1.GetMeRes{}, nil
}

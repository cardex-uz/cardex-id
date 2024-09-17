package usecase

import (
	ssov1 "cardex-id/pkg/grpc/sso"
	"google.golang.org/grpc"
)

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

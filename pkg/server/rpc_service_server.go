package server

import (
	"context"
	"github.com/glide-im/auth/pkg/proto"
)

type RpcServer interface {
	Auth(ctx context.Context, request *proto.AuthRequest, response *proto.AuthResponse) error

	RemoveToken(ctx context.Context, request *proto.RemoveTokenRequest, response *proto.RemoveTokenResponse) error

	GetToken(ctx context.Context, request *proto.GenTokenRequest, response *proto.GetTokenResponse) error
}

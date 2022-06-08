package auth_service

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/glide-im/auth/internal/rpc"
	"github.com/glide-im/auth/pkg/proto"
	"github.com/glide-im/glide/pkg/auth"
	"github.com/glide-im/glide/pkg/auth/jwt_auth"
)

type Options struct {
	Name      string
	Network   string
	Addr      string
	Port      int
	JwtSecret string
}

var ja auth.Authorize

type Server struct{}

func Run(opts *Options) error {
	ja = jwt_auth.NewAuthorizeImpl(opts.JwtSecret)
	options := rpc.ServerOptions{
		Name:    opts.Name,
		Network: "tcp",
		Addr:    opts.Addr,
		Port:    opts.Port,
	}
	srv := rpc.NewBaseServer(&options)
	srv.Register(opts.Name, &Server{})
	return srv.Run()
}

func (*Server) Auth(ctx context.Context, request *proto.AuthRequest, response *proto.AuthResponse) error {
	info := request.GetAuthInfo()
	jwtAuthInfo := jwt_auth.JwtAuthInfo{}
	err := json.Unmarshal(info, &jwtAuthInfo)
	if err != nil {
		return errors.New("auth info is invalid, err: " + err.Error())
	}
	token := auth.Token{
		Token: request.GetToken().GetToken(),
	}
	result, err := ja.Auth(&jwtAuthInfo, &token)
	if err != nil {
		return err
	}
	response.Success = result.Success
	response.Response, err = json.Marshal(result.Response)
	return err
}

func (*Server) RemoveToken(ctx context.Context, request *proto.RemoveTokenRequest, response *proto.RemoveTokenResponse) error {
	token := auth.Token{
		Token: request.GetToken().GetToken(),
	}
	return ja.RemoveToken(&token)
}

func (*Server) GetToken(ctx context.Context, request *proto.GenTokenRequest, response *proto.GetTokenResponse) error {

	info := request.GetAuthInfo()

	jwtAuthInfo := jwt_auth.JwtAuthInfo{}
	err := json.Unmarshal(info, &jwtAuthInfo)
	if err != nil {
		return errors.New("auth info is invalid, err: " + err.Error())
	}

	token, err := ja.GetToken(&jwtAuthInfo)
	if err != nil {
		return err
	}
	response.Token = &proto.Token{
		Token: token.Token,
	}
	return nil
}

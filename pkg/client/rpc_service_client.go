package client

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/glide-im/auth/internal/rpc"
	"github.com/glide-im/auth/pkg/proto"
	"github.com/glide-im/glide/pkg/auth"
	"github.com/glide-im/glide/pkg/auth/jwt_auth"
)

type CommonAuthClient struct {
	rpc *rpc.BaseClient
}

func NewAuthServiceClient(name, addr string, port int) (*CommonAuthClient, error) {
	options := rpc.ClientOptions{
		Addr: addr,
		Port: port,
		Name: name,
	}
	cli, err := rpc.NewBaseClient(&options)
	if err != nil {
		return nil, err
	}
	return &CommonAuthClient{
		rpc: cli,
	}, nil
}

func (s *CommonAuthClient) Auth(c auth.Info, t *auth.Token) (*auth.Result, error) {
	if c == nil || t == nil {
		return nil, errors.New("auth info or token is nil")
	}

	autInfo, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}
	token := proto.Token{
		Token: t.Token,
	}
	request := proto.AuthRequest{
		AuthInfo: autInfo,
		Token:    &token,
	}
	response := proto.AuthResponse{}
	err = s.rpc.Call(context.TODO(), "Auth", &request, &response)
	if err != nil {
		return nil, err
	}

	jr := jwt_auth.Response{}

	err = json.Unmarshal(response.GetResponse(), &jr)
	if err != nil {
		return nil, err
	}

	result := auth.Result{
		Success:  response.GetSuccess(),
		Response: jr,
	}
	return &result, nil
}

func (s *CommonAuthClient) RemoveToken(t *auth.Token) error {
	if t == nil {
		return errors.New("token is nil")
	}
	token := &proto.Token{
		Token: t.Token,
	}
	request := proto.RemoveTokenRequest{
		Token: token,
	}
	response := proto.RemoveTokenResponse{}
	err := s.rpc.Call(context.TODO(), "RemoveToken", &request, &response)
	return err
}

func (s *CommonAuthClient) GetToken(c auth.Info) (*auth.Token, error) {
	if c == nil {
		return nil, errors.New("auth info is nil")
	}
	bytes, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}
	request := proto.GenTokenRequest{
		AuthInfo: bytes,
	}
	response := proto.GetTokenResponse{}
	err = s.rpc.Call(context.TODO(), "GetToken", &request, &response)
	if err != nil {
		return nil, err
	}
	token := auth.Token{
		Token: response.Token.GetToken(),
	}
	return &token, nil
}

func (s *CommonAuthClient) Close() error {
	return s.rpc.Close()
}

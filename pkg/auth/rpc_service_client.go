package auth

import (
	"github.com/glide-im/glide/pkg/auth"
)

type CommonAuthService struct {
}

func NewAuthServiceClient() *CommonAuthService {
	return &CommonAuthService{}
}

func (s *CommonAuthService) Auth(c auth.Info, t *auth.Token) (*auth.Result, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CommonAuthService) RemoveToken(t *auth.Token) error {
	//TODO implement me
	panic("implement me")
}

func (s *CommonAuthService) GetToken(c auth.Info) (*auth.Token, error) {
	//TODO implement me
	panic("implement me")
}

package auth_service

import (
	"github.com/glide-im/glide/pkg/auth"
)

type Server struct {
}

func (s *Server) Auth(c auth.Info, t *auth.Token) (*auth.Result, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) RemoveToken(t *auth.Token) error {
	//TODO implement me
	panic("implement me")
}

func (s *Server) GetToken(c auth.Info) (*auth.Token, error) {
	//TODO implement me
	panic("implement me")
}

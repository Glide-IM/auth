package main

import (
	"fmt"
	"github.com/glide-im/auth/pkg/client"
	"github.com/glide-im/glide/pkg/auth"
	"github.com/glide-im/glide/pkg/auth/jwt_auth"
)

var authService auth.Authorize
var authInfo = jwt_auth.JwtAuthInfo{
	UID:    "8",
	Device: "1",
}
var token string

func main() {
	cli, err := client.NewAuthServiceClient("auth", "127.0.0.1", 8093)
	defer cli.Close()
	if err != nil {
		panic(err)
	}
	authService = cli

	ExampleAuthServiceGetAuth()
	ExampleAuthServiceAuth()
}

func ExampleAuthServiceGetAuth() {
	t, err := authService.GetToken(authInfo)
	if err != nil {
		panic(err)
	}
	token = t.Token
	fmt.Println("GetAuth:", token)
}

func ExampleAuthServiceAuth() {
	result, err := authService.Auth(jwt_auth.JwtAuthInfo{
		UID:    "1",
		Device: "0",
	}, &auth.Token{Token: token})
	if err != nil {
		panic(err)
	}

	fmt.Println("Auth:", result.Success, result.Response)
}

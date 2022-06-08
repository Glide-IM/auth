package main

import (
	"fmt"
	"github.com/glide-im/auth/internal/auth_service"
	"github.com/glide-im/auth/internal/config"
)

func main() {
	config.MustLoad()

	options := auth_service.Options{
		Name:      config.AuthServer.Name,
		Addr:      config.AuthServer.Addr,
		Port:      config.AuthServer.Port,
		JwtSecret: config.AuthServer.JwtSecret,
	}
	fmt.Println("[auth] server start at:", options.Addr, options.Port)
	err := auth_service.Run(&options)
	if err != nil {
		panic(err)
	}
}

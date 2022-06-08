package config

import "github.com/spf13/viper"

var (
	Redis      *RedisConf
	AuthServer *AuthRpcServerConf
)

type WsServerConf struct {
	ID        string
	Addr      string
	Port      int
	JwtSecret string
}

type ApiHttpConf struct {
	Addr string
	Port int
}

type AuthRpcServerConf struct {
	Addr      string
	Port      int
	Network   string
	Name      string
	JwtSecret string
}

type MySqlConf struct {
	Host     string
	Port     int
	Username string
	Password string
	Db       string
	Charset  string
}

type RedisConf struct {
	Host     string
	Port     int
	Password string
	Db       int
}

func MustLoad() {

	viper.SetConfigName("config.toml")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")
	viper.AddConfigPath("/etc/")
	viper.AddConfigPath("$HOME/.config/")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	c := struct {
		Redis      *RedisConf
		AuthServer *AuthRpcServerConf
	}{}

	err = viper.Unmarshal(&c)
	if err != nil {
		panic(err)
	}
	Redis = c.Redis
	AuthServer = c.AuthServer

	if c.Redis == nil {
		panic("redis config is nil")
	}
	if c.AuthServer == nil {
		panic("auth server config is nil")
	}
}

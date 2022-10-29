package app

import (
	"fmt"

	"github.com/cristalhq/aconfig"
	"github.com/cristalhq/aconfig/aconfigdotenv"
)

type Config struct {
	Log  LogConfig  `env:"IFI_SPACESHIP_LOG"`
	GRPC GRPCConfig `env:"IFI_SPACESHIP_GRPC"`
	DB   DBConfig   `env:"IFI_SPACESHIP_MYSQL"`
}

type LogConfig struct {
	Format string `env:"FORMAT" default:"console"`
	Debug  bool   `env:"DEBUG" default:"false"`
}

type GRPCConfig struct {
	Port int `env:"PORT" default:"50061"`
}
type DBConfig struct {
	Username string `env:"USERNAME" default:"root"`
	Password string `env:"PASSWORD" default:"root"`
	Host     string `env:"HOST" default:"mysql"`
	Port     string `env:"PORT" default:"3306"`
	Database string `env:"DATABASE" default:"inventory"`
}

func (a *Application) initConfig() error {
	var cfg Config
	loader := aconfig.LoaderFor(&cfg, aconfig.Config{
		AllowUnknownFields: true,
		AllowUnknownEnvs:   true,
		FailOnFileNotFound: false,
		Files:              []string{".env"},
		FileDecoders:       map[string]aconfig.FileDecoder{".env": aconfigdotenv.New()},
	})
	if err := loader.Load(); err != nil {
		fmt.Println("error:", err)
		return err
	}

	a.config = cfg

	return nil
}

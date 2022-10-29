package app

import (
	"fmt"

	"github.com/cristalhq/aconfig"
	"github.com/cristalhq/aconfig/aconfigdotenv"
)

type Config struct {
	Log  LogConfig  `env:"IFI_HTTP_LOG"`
	HTTP HTTPConfig `env:"IFI_HTTP_SERVICE"`
}

type LogConfig struct {
	Format string `env:"FORMAT" default:"console"`
	Debug  bool   `env:"DEBUG" default:"false"`
}

type HTTPConfig struct {
	Port int `env:"PORT" default:"8080"`
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

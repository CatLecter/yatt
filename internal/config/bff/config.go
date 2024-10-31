package bffconfig

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"time"
)

type AppConfig struct {
	Host         string        `yaml:"host" env:"APP_HOST" env-default:"127.0.0.1"`
	Port         int           `yaml:"port" env:"APP_PORT" env-default:"8000"`
	ReadTimeout  time.Duration `yaml:"read_timeout" env:"APP_READ_TIMEOUT" env-default:"300ms"`
	WriteTimeout time.Duration `yaml:"write_timeout" env:"APP_WRITE_TIMEOUT" env-default:"300ms"`
	AllowOrigins string        `yaml:"allow_origins" env:"APP_ALLOW_ORIGIN" env-default:"localhost;127.0.0.1;0.0.0.0"`
	StopTimeout  time.Duration `yaml:"stop_timeout" env:"APP_STOP_TIMEOUT" env-default:"3s"`
}

type IaaConfig struct {
	URI          string        `yaml:"uri" env:"IAA_URI" env-default:"127.0.0.1:5080"`
	Retries      uint          `yaml:"retries" env:"IAA_RETRIES" env-default:"3"`
	RetryTimeout time.Duration `yaml:"retry_timeout" env:"IAA_RETRY_TIMEOUT" env-default:"300ms"`
}

type Config struct {
	App AppConfig `yaml:"app"`
	Iaa IaaConfig `yaml:"iaa"`
}

func MustNew() *Config {
	var configPath string

	flag.StringVar(&configPath, "config", "", "Path to YAML config file")
	flag.Parse()

	if configPath == "" {
		configPath = os.Getenv("CONFIG_PATH")
	}

	if configPath == "" {
		panic("Specify the path to the configuration file using the --config parameter or in the CONFIG_PATH environment variable")
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic("Cannot read config: " + err.Error())
	}

	return &cfg
}

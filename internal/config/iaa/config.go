package iaaconfig

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"time"
)

type AppConfig struct {
	Host        string        `yaml:"host" env:"APP_HOST" env-default:"127.0.0.1"`
	Port        int           `yaml:"port" env:"APP_PORT" env-default:"5090"`
	StopTimeout time.Duration `yaml:"stop_timeout" env:"APP_STOP_TIMEOUT" env-default:"3s"`
}

type StorageConfig struct {
	URI             string        `yaml:"uri" env:"STORAGE_URI" env-default:"postgresql://admin:admin@127.0.0.1:5432/db?sslmode=disable"`
	MaxConnections  int32         `yaml:"max_connections" env:"STORAGE_MAX_CONNECTIONS" env-default:"150"`
	MinConnections  int32         `yaml:"min_connections" env:"STORAGE_MIN_CONNECTIONS" env-default:"5"`
	MaxConnLifetime time.Duration `yaml:"max_conn_life_time" env:"STORAGE_MAX_CONN_LIFE_TIME" env-default:"300ms"`
	MaxConnIdleTime time.Duration `yaml:"max_conn_idle_time" env:"STORAGE_MAX_CONN_IDLE_TIME" env-default:"100ms"`
}

type AuthConfig struct {
	JWTExp    time.Duration `yaml:"jwt_exp" env:"AUTH_JWT_EXP" env-default:"24h"`
	JWTSecret string        `yaml:"jwt_secret" env:"AUTH_JWT_SECRET" env-default:"secret"`
}

type Config struct {
	App     AppConfig     `yaml:"app"`
	Storage StorageConfig `yaml:"storage"`
	Auth    AuthConfig    `yaml:"auth"`
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

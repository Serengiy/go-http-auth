package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"time"
)

type Config struct {
	Env            string        `yaml:"env" env-default:"dev"`
	MigrationsPath string        `yaml:"migrations_path"`
	TokenTTL       time.Duration `yaml:"token_ttl" env-default:"1h"`
	Port           string        `yaml:"port" env-default:"8081"`
	AppSecret      string        `yaml:"app_secret" env-default:"secret"`
	Database       Database
}
type Database struct {
	Driver   string `yaml:"driver" env-default:"postgres"`
	Username string `yaml:"username" env-default:"postgres"`
	Database string `yaml:"database" env-default:"postgres"`
	Password string `yaml:"password" env-default:"postgres"`
	Port     int    `yaml:"port" env-default:"5432"`
	Sslmode  string `yaml:"sslmode" env-default:"sslmode"`
	Host     string `yaml:"host" env-default:"http://localhost"`
}

type AppSecret struct {
	Secret string `yaml:"app_secret" env-default:"secret"`
}

func MustLoad() *Config {
	path := getFilePath()

	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("config file not found: " + path)
	}
	return MustLoadFromPath(path)
}

func MustLoadFromPath(path string) *Config {
	var cfg Config
	err := cleanenv.ReadConfig(path, &cfg)
	if err != nil {
		panic(err)
	}
	return &cfg
}

func GetSecretKey() string {
	secretKey := os.Getenv("SECRET_KEY")
	if secretKey == "" {
		var secret AppSecret
		path := getFilePath()
		if _, err := os.Stat(path); os.IsNotExist(err) {
			panic("config file not found: " + path)
		}
		err := cleanenv.ReadConfig(path, &secret)
		if err != nil {
			panic(err)
		}
		secretKey = secret.Secret
	}
	return secretKey
}

func getFilePath() string {
	var filePath string

	if _, err := os.Stat("../../config/config.yaml"); err != nil {
		filePath = os.Getenv("CONFIG_PATH")
	} else {
		filePath = "../../config/config.yaml"
	}
	return filePath
}

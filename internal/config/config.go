package config

import "time"

type Config struct {
	HTTP ServerConfig `yaml:"http"`
	DB   DBConfig     `yaml:"db"`
}

type ServerConfig struct {
	Port            string        `yaml:"port"`
	Timeout         time.Duration `yaml:"timeout"`
	ShutDownTimeout time.Duration `yaml:"shutdown_timeout"`
	ReadTimeout     time.Duration `yaml:"read_timeout"`
	WriteTimeout    time.Duration `yaml:"write_timeout"`
}

type DBConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	DBName   string `yaml:"db_name"`
	Username string `yaml:"username"`
	Password string `env:"DB_PASSWORD" env-required:"true"`
}

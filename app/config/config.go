package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/goccy/go-yaml"
)

type Config struct {
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
}

type ServerConfig struct {
	Port string `yaml:"port"`
}

type DatabaseConfig struct {
	Host         string `yaml:"host"`
	Port         int    `yaml:"port"`
	Name         string `yaml:"name"`
	User         string `yaml:"user"`
	Password     string `yaml:"password"`
	MaxOpenConns int    `yaml:"maxOpenConns"`
	MaxIdleConns int    `yaml:"maxIdleConns"`
}

func Load(path string) (Config, error) {
	cfg := defaultConfig()

	if path == "" {
		path = "config.yaml"
	}

	if data, err := os.ReadFile(path); err == nil {
		if err := yaml.Unmarshal(data, &cfg); err != nil {
			return Config{}, fmt.Errorf("parse config file: %w", err)
		}
	} else if !os.IsNotExist(err) {
		return Config{}, fmt.Errorf("read config file: %w", err)
	}

	overrideFromEnv(&cfg)

	if cfg.Database.Host == "" || cfg.Database.Name == "" || cfg.Database.User == "" || cfg.Database.Password == "" {
		return Config{}, fmt.Errorf("database configuration is incomplete")
	}

	return cfg, nil
}

func defaultConfig() Config {
	return Config{
		Server: ServerConfig{
			Port: "8080",
		},
		Database: DatabaseConfig{
			Host:         "localhost",
			Port:         3306,
			Name:         "coffee_db",
			User:         "root",
			MaxOpenConns: 25,
			MaxIdleConns: 5,
		},
	}
}

func overrideFromEnv(cfg *Config) {
	if value := os.Getenv("SERVER_PORT"); value != "" {
		cfg.Server.Port = value
	}
	if value := os.Getenv("MYSQL_HOST"); value != "" {
		cfg.Database.Host = value
	}
	if value := os.Getenv("MYSQL_PORT"); value != "" {
		if port, err := strconv.Atoi(value); err == nil {
			cfg.Database.Port = port
		}
	}
	if value := os.Getenv("MYSQL_DATABASE"); value != "" {
		cfg.Database.Name = value
	}
	if value := os.Getenv("MYSQL_USER"); value != "" {
		cfg.Database.User = value
	}
	if value := os.Getenv("MYSQL_PASSWORD"); value != "" {
		cfg.Database.Password = value
	}
	if value := os.Getenv("MYSQL_MAX_OPEN_CONNS"); value != "" {
		if maxOpen, err := strconv.Atoi(value); err == nil {
			cfg.Database.MaxOpenConns = maxOpen
		}
	}
	if value := os.Getenv("MYSQL_MAX_IDLE_CONNS"); value != "" {
		if maxIdle, err := strconv.Atoi(value); err == nil {
			cfg.Database.MaxIdleConns = maxIdle
		}
	}
}

package configs

import (
	"fmt"
	"os"
	"strconv"
	"testing"

	"encoding/json"
)

type envConfig struct {
	DBDriver      string
	DBHost        string
	DBPort        string
	DBUser        string
	DBPassword    string
	DBName        string
	DBOptions     string
	WebServerPort string
	JWTSecret     string
	JWTExpiresIn  int
}

func LoadConfig() *envConfig {
	jwtExpiresIn, err := strconv.Atoi(os.Getenv("JWT_EXPIRES_IN"))

	if err != nil {
		panic(err)
	}

	mapCfg := map[string]interface{}{
		"DBDriver":      os.Getenv("DB_DRIVER"),
		"DBHost":        os.Getenv("DB_HOST"),
		"DBPort":        os.Getenv("DB_PORT"),
		"DBUser":        os.Getenv("DB_USER"),
		"DBPassword":    os.Getenv("DB_PASSWORD"),
		"DBName":        os.Getenv("DB_NAME"),
		"DBOptions":     os.Getenv("DB_OPTIONS"),
		"WebServerPort": os.Getenv("WEB_SERVER_PORT"),
		"JWTSecret":     os.Getenv("JWT_SECRET"),
		"JWTExpiresIn":  jwtExpiresIn,
	}

	for key, value := range mapCfg {

		if value == "" {
			panic(fmt.Sprintf("%s is missing on environment loading", key))
		}
	}

	jsonCfg, err := json.Marshal(mapCfg)

	if err != nil {
		panic(err)
	}

	var cfg envConfig

	err = json.Unmarshal(jsonCfg, &cfg)

	if err != nil {
		panic(err)
	}

	return &cfg
}

func LoadTestEnv(t *testing.T) *envConfig {
	t.Setenv("DB_DRIVER", "pgx")
	t.Setenv("DB_HOST", "localhost")
	t.Setenv("DB_PORT", "5432")
	t.Setenv("DB_USER", "test")
	t.Setenv("DB_PASSWORD", "test")
	t.Setenv("DB_NAME", "inventory_test")
	t.Setenv("DB_OPTIONS", "sslmode=disable")
	t.Setenv("WEB_SERVER_PORT", "8000")
	t.Setenv("JWT_SECRET", "secret")
	t.Setenv("JWT_EXPIRES_IN", "500")

	jwtExpiresIn, err := strconv.Atoi(os.Getenv("JWT_EXPIRES_IN"))

	if err != nil {
		panic(err)
	}

	return &envConfig{
		DBDriver:      os.Getenv("DB_DRIVER"),
		DBHost:        os.Getenv("DB_HOST"),
		DBPort:        os.Getenv("DB_PORT"),
		DBUser:        os.Getenv("DB_USER"),
		DBPassword:    os.Getenv("DB_PASSWORD"),
		DBName:        os.Getenv("DB_NAME"),
		DBOptions:     os.Getenv("DB_OPTIONS"),
		WebServerPort: os.Getenv("WEB_SERVER_PORT"),
		JWTSecret:     os.Getenv("JWT_SECRET"),
		JWTExpiresIn:  jwtExpiresIn,
	}
}

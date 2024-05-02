package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	App struct {
		Host       string `envconfig:"APP_HOST" default:"localhost"`
		Port       string `envconfig:"APP_PORT" default:"8080"`
		ApiVersion string `envconfig:"API_VERSION" default:"v0"`
		AppVersion string `envconfig:"APP_VERSION" default:"v0.0.1"`
	}
	Database struct {
		MongoUri      string `envconfig:"MONGODB_URI" default:"mongodb://localhost:27017"`
		AccessTimeout int    `envconfig:"MONGODB_ACCESS_TIMEOUT" default:"5"`
	}
	Auth struct {
		JWTSecret        string `envconfig:"JWT_SECRET" default:"token-secret"`
		JWTExpireInHours int    `envconfig:"JWT_EXPIRE" default:"24"`
		TokenExpire      int    `envconfig:"TOKEN_EXPIRE" default:"60"`
		ShortTokenExpire int    `envconfig:"SHORT_TOKEN_EXPIRE" default:"15"`
		JWTIssuer        string `envconfig:"JWT_ISSUER" default:"student-halls.com"`
	}
	Cors struct {
		AllowOrigins     []string `envconfig:"CORS_ALLOW_ORIGINS" default:"*"`
		AllowMethods     []string `envconfig:"CORS_ALLOW_METHODS" default:"GET, POST, PUT, DELETE, OPTIONS"`
		AllowHeaders     []string `envconfig:"CORS_ALLOW_HEADERS" default:"Origin, Content-Length, Content-Type, Authorization, Tenant"`
		AllowCredentials bool     `envconfig:"CORS_ALLOW_CREDENTIALS" default:"true"`
	}
}

// It is initialized once when the application starts.
var appConfig = &Config{}

// This function provides access to the appConfig variable.
func AppConfig() *Config {
	return appConfig
}

// LoadConfig loads environment variables and populates appConfig.
// It first attempts to load variables from a .env file and then
// processes environment variables according to the Config struct tags.
func LoadConfig() error {
	godotenv.Load()
	if err := envconfig.Process("", appConfig); err != nil {
		return err
	}

	return nil
}

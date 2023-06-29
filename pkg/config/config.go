package config

import (
	"os"
	"regexp"
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

type Configuration struct {
	Http     HttpConfiguration
	Logger   LoggerConfiguration
	Database DatabaseConfiguration
}

type HttpConfiguration struct {
	Port string
}

type LoggerConfiguration struct {
	Path  string
	Level string
}

type DatabaseConfiguration struct {
	Name     string
	Host     string
	Port     string
	Username string
	Password string
	Ssl      string
}

func Load() (*Configuration, error) {
	// load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	v := viper.New()

	// set config path as the root of project
	v.AddConfigPath(".")

	// set the config file name
	v.SetConfigName("config")

	// set config type to `yml`` file
	v.SetConfigType("yml")

	// load environment variables
	v.AutomaticEnv()

	// Viper reads all the variables from env file and log error if any found
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	config := &Configuration{}

	// Viper unmarshals the loaded env varialbes into the struct
	if err := v.Unmarshal(&config); err != nil {
		return nil, err
	}

	// Define a regular expression to match placeholders of the form "${NAME:DEFAULT_VALUE}"
	placeholderRegex := regexp.MustCompile(`\${([^}]+)}`)

	// Replace placeholders in the Config struct with corresponding environment variables or default values
	yamlString, err := yaml.Marshal(config)
	if err != nil {
		return nil, err
	}

	replacedString := placeholderRegex.ReplaceAllStringFunc(string(yamlString), func(match string) string {
		parts := strings.Split(strings.Trim(match, "${}"), ":")
		name := parts[0]
		defaultValue := ""

		if len(parts) > 1 {
			defaultValue = parts[1]
		}

		value := os.Getenv(name)

		if value == "" {
			// Placeholder was not replaced with environment variable value, try default value
			if defaultValue == "" {
				// No default value specified, keep the placeholder as is
				return match
			}

			// Use the default value
			value = defaultValue
		}

		// Placeholder was replaced with the environment variable or default value
		return value
	})

	// Unmarshal the final YAML string into the Config struct
	err = yaml.Unmarshal([]byte(replacedString), config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

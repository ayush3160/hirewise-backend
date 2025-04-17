package config

import (
	"io/fs"
	"os"
	"strings"
	"sync"

	"github.com/spf13/viper"
	yamlLib "gopkg.in/yaml.v3"
)

type (
	Config struct {
		Server *Server
		Db     *Db
	}

	Server struct {
		Port int
	}

	Db struct {
		Host     string
		Port     int
		User     string
		Password string
		DBName   string
	}
)

var (
	once           sync.Once
	configInstance *Config
)

func GetConfig() *Config {
	once.Do(func() {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("./")
		viper.AutomaticEnv()
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

		if err := viper.ReadInConfig(); err != nil {
			panic(err)
		}

		if err := viper.Unmarshal(&configInstance); err != nil {
			panic(err)
		}
	})

	return configInstance
}

func CreateDefaultConfig() error {
	defaultConfig := `
server:
  port: 8080
db:
  host: localhost
  port: 5432
  user: postgres
  password: 123456
  dbname: hirewise-db
`
	var node yamlLib.Node

	if err := yamlLib.Unmarshal([]byte(defaultConfig), &node); err != nil {
		return err
	}

	results, err := yamlLib.Marshal(node.Content[0])
	if err != nil {
		return err
	}

	err = os.WriteFile("config.yaml", results, fs.ModePerm)
	if err != nil {
		return err
	}

	err = os.Chmod("config.yaml", 0777) // Set permissions to 777
	if err != nil {
		return err
	}

	return nil
}

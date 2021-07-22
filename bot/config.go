package bot

import (
	"bot/pkg/alicloud"
	"strings"

	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var configFile = pflag.StringP("config", "c", "config.yaml", "config file")

// Config ...
type Config struct {
	APIToken       string             `mapstructure:"api-token"`
	MasterUsername string             `mapstructure:"master-username"`
	AliCloud       alicloud.AccessKey `mapstructure:"alicloud"`
}

// ParseConfig ...
func ParseConfig() (*Config, error) {
	godotenv.Load()
	pflag.Parse()

	viper.AddConfigPath(".")
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_", ".", "_"))
	viper.AutomaticEnv()

	viper.SetConfigFile(*configFile)
	viper.SetConfigType("yaml")

	config := new(Config)
	if err := viper.Unmarshal(config); err != nil {
		return nil, err
	}

	if config.APIToken == "" {
		return nil, errors.New("API Token cloud not be empty")
	}

	return config, nil
}

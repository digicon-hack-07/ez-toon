package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	MariaDBHostname string `mapstructure:"mariadb_hostname" json:"mariadb_hostname,omitempty"`
	MariaDBPort     int    `mapstructure:"mariadb_port" json:"mariadb_port,omitempty"`
	MariaDBDatabase string `mapstructure:"mariadb_database" json:"mariadb_database,omitempty"`
	MariaDBUsername string `mapstructure:"mariadb_username" json:"mariadb_username,omitempty"`
	MariaDBPassword string `mapstructure:"mariadb_password" json:"mariadb_password,omitempty"`
}

func GetConfig() (*Config, error) {
	viper.SetDefault("MariaDB_Hostname", "localhost")
	viper.SetDefault("MariaDB_Port", 3306)
	viper.SetDefault("MariaDB_Database", "ez-toon")
	viper.SetDefault("MariaDB_Username", "root")
	viper.SetDefault("MariaDB_Password", "password")

	viper.AutomaticEnv()

	var c Config
	err := viper.Unmarshal(&c)
	if err != nil {
		return nil, fmt.Errorf("Error: failed to parse configs - %s ", err)
	}

	return &c, nil
}

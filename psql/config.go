/**
 * Copyright 2019 Whiteblock Inc. All rights reserved.
 * Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

package psql

import (
	"github.com/spf13/viper"
)

// Config is the configuration of the Postgresql database
type Config struct {
	DBHost           string `mapstructure:"dbhost"`
	DBPort           int    `mapstructure:"dbport"`
	DBUser           string `mapstructure:"dbuser"`
	DBPassword       string `mapstructure:"dbpassword"`
	DBName           string `mapstructure:"dbname"`
	MaxOpenConns     int    `mapstructure:"maxopenconns"`
	MaxIdleConns     int    `mapstructure:"maxidleconns"`
	MaxLifetimeConns int    `mapstructure:"maxlifetimeconns"`
}

// NewConfig gets the PSQLConfig from viper
func NewConfig(v *viper.Viper) (out Config, _ error) {
	return out, v.Unmarshal(&out)
}

// SetConfig the env vars and defaults with viper
func SetConfig(v *viper.Viper) {
	v.BindEnv("dbhost", "DBHOST")
	v.BindEnv("dbport", "DBPORT")
	v.BindEnv("dbuser", "DBUSER")
	v.BindEnv("dbpassword", "DBPASSWORD")
	v.BindEnv("dbname", "DBNAME")
	v.BindEnv("maxopenconns", "DB_MAX_OPEN")
	v.BindEnv("maxidleconns", "DB_MAX_IDLE")
	v.BindEnv("maxlifetimeconns", "DB_LIFETIME")

	v.SetDefault("dbhost", "127.0.0.1")
	v.SetDefault("dbport", 5432)
	v.SetDefault("dbuser", "genesis")
	v.SetDefault("dbpassword", "genesis")
	v.SetDefault("dbname", "genesis")
	v.SetDefault("maxopenconns", 5)
	v.SetDefault("maxidleconns", 2)
	v.SetDefault("maxlifetimeconns", 60)
}

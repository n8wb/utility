/**
 * Copyright 2019 Whiteblock Inc. All rights reserved.
 * Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

package metrics

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Prefix            string        `mapstructure:"metricsPrefix"`
	ProjectID         string        `mapstructure:"projectID"`
	ReportingInterval time.Duration `mapstructure:"metricsReportingInterval"`
}

// NewConfig gets the Config from viper
func NewConfig(v *viper.Viper) (out Config, _ error) {
	return out, v.Unmarshal(&out)
}

// SetConfig the env vars and defaults with viper
func SetConfig(v *viper.Viper) {
	v.BindEnv("metricsReportingInterval", "METRIC_REPORTING_INTERVAL")
	v.BindEnv("metricsPrefix", "METRICS_PREFIX")
	v.BindEnv("projectID", "PROJECT_ID")

	v.SetDefault("metricsReportingInterval", 60*time.Second)
}

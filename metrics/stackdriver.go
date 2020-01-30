/**
 * Copyright 2019 Whiteblock Inc. All rights reserved.
 * Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

package metrics

import (
	"contrib.go.opencensus.io/exporter/stackdriver"
	"contrib.go.opencensus.io/exporter/stackdriver/monitoredresource"
)

// StackdriverExporter returns a configured stackdriver exporter
func StackdriverExporter(conf Config) (*stackdriver.Exporter, error) {
	return stackdriver.NewExporter(stackdriver.Options{
		ProjectID:         conf.ProjectID,
		MetricPrefix:      conf.Prefix,
		ReportingInterval: conf.ReportingInterval,
		MonitoredResource: monitoredresource.Autodetect(),
	})
}

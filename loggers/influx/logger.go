// influx is the package the implements logger interface. This package reports the logs to influxdb directly for presisting http access logs.
package influx

import (
	"context"
	"fmt"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/pbivrell/httplog"
)

const (
	// Org is the default influx organization. This can be configured by providing WithOrg to NewLogger
	Org = "test"
	// Bucket is the default influx bucket. This can be configured by providing WithBucket to NewLogger
	Bucket = "webmetrics"
)

// Logger is the influx implementation of httplog.Logger
type Logger struct {
	// Client is the influx client to communicate to influx with
	Client influxdb2.Client
	// Org is the influx organization to write the logs to. https://docs.influxdata.com/influxdb/v2.0/organizations/
	Org string
	// Bucket is the influx bucket to write the logs to. https://docs.influxdata.com/influxdb/v2.0/organizations/buckets/
	Bucket string
}

// WithClient allows configuration of the influ client to write to.
func WithClient(client influxdb2.Client) loggerOpt {
	return func(l *Logger) {
		l.Client = client
	}
}

// WithOrg allows configuration of the influx organization to write to.
func WithOrg(org string) loggerOpt {
	return func(l *Logger) {
		l.Org = org
	}
}

// WithBucket allows for configuration of the influx bucket to write to.
func WithBucket(bucket string) loggerOpt {
	return func(l *Logger) {
		l.Bucket = bucket
	}
}

type loggerOpt func(*Logger)

// NewLogger returns a httplog.Logger implementation. Storing logs using the default influx client. You may also provide 0 or more optional configuration functions for over riding defaults.
func NewLogger(opts ...loggerOpt) *Logger {

	l := &Logger{
		Client: influxdb2.NewClient("http://localhost:8086", ""),
		Org:    Org,
		Bucket: Bucket,
	}

	// Loop over optional configuration functions applying them to our logger before we return
	for _, opt := range opts {
		opt(l)
	}

	return l
}

// Write implements the httplog.Logger interface writing the provided httlog.Data to influx.
func (l *Logger) Write(data httplog.Data) error {

	writeAPI := l.Client.WriteAPIBlocking(l.Org, l.Bucket)
	p := influxdb2.NewPoint(
		"request",
		map[string]string{
			"method":      data.Method,
			"endpoint":    data.Endpoint,
			"code":        fmt.Sprintf("%d", data.Code),
			"ip":          data.IP,
			"os":          data.OS,
			"decive_type": data.DeviceType,
		},
		map[string]interface{}{
			"os_version":        data.OSVersion,
			"device":            data.Device,
			"useragent":         data.UserAgent,
			"useragent_version": data.UserAgentVersion,
			"duration":          data.Duration,
			"referer":           data.Referer,
		},
		data.Time,
	)
	return writeAPI.WritePoint(context.Background(), p)
}

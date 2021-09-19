// httplog is a package for logging http access requests. It defines an interface for Writing log data as well as the structure of http access log data one must be able to write.
package httplog

import (
	"time"
)

// Logger is the interface for writing the http access data
type Logger interface {
	Write(Data) error
}

// Data defines the data members the http access logger must be willing to write
type Data struct {
	// Method HTTP Method
	Method string
	// Endpoint of the request
	Endpoint string
	// Code the integer status code from the completed request
	Code int
	// Referer the value of the 'Referer' HTTP header
	Referer string
	// Duration of the request
	Duration int64
	// IP of the requestor
	IP string
	// UserAgent in the uas string
	UserAgent string
	// UserAgentVersion in the uas string
	UserAgentVersion string
	// OS in the uas string
	OS string
	// OSVersion in the uas string
	OSVersion string
	// Device in the uas string
	Device string
	// DeviceType in the uas string
	DeviceType string
	// Time when the request was completed
	Time time.Time
}

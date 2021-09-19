package httplog

import (
	"net/http"
	"strings"
	"time"

	"github.com/felixge/httpsnoop"

	ua "github.com/mileusna/useragent"
)

// Middlewear is an implementation of a logging http middlewear that uses felixge/httpsnoop to collect the relavent
// request information and mileusna/useragent to collect information from the 'User-Agent' header, then output the logs to the configured logger. This middlewear runs after the wrapped request has completed.
func Middlewear(logger Logger, handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Preform the request and capture metrics about the result
		m := httpsnoop.CaptureMetrics(handler, w, r)

		t := time.Now()

		// Parse useragent
		ua := ua.Parse(r.Header.Get("User-Agent"))

		deviceType := "n/a"
		if ua.Bot {
			deviceType = "bot"
		} else if ua.Mobile {
			deviceType = "mobile"
		} else if ua.Tablet {
			deviceType = "tablet"
		} else if ua.Desktop {
			deviceType = "desktop"
		}

		logger.Write(Data{
			Method:           r.Method,
			Endpoint:         r.URL.String(),
			Referer:          r.Header.Get("Referer"),
			Code:             m.Code,
			Duration:         m.Duration.Milliseconds(),
			IP:               r.RemoteAddr[0:strings.LastIndex(r.RemoteAddr, ":")],
			UserAgent:        ua.Name,
			UserAgentVersion: ua.Version,
			OS:               ua.OS,
			OSVersion:        ua.OSVersion,
			Device:           ua.Device,
			DeviceType:       deviceType,
			Time:             t,
		})
	}
}

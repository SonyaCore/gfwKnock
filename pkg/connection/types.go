package connection

import (
	"time"
)

// Configuration holds the configuration settings for the server.
type Configuration struct {
	SocketTimeout   time.Duration `json:"socket_timeout"`
	AcceptTimeSleep time.Duration `json:"accept_time_sleep"`
	FragmentSleep   time.Duration `json:"fragment_sleep"`
	LFragment       int           `json:"l_fragment"`
}

// Server represents the server configuration.
type Server struct {
	ListenPort     int    `json:"listen_port"`
	CloudFlareIP   string `json:"cloud_flare_ip"`
	CloudFlarePort int    `json:"cloud_flare_port"`
	Configuration
}

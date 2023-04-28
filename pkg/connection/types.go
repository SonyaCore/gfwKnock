package connection

import (
	"time"
)

type Configuration struct {
	SocketTimeout   time.Duration `json:"socket_timeout"`
	AcceptTimeSleep time.Duration `json:"accept_time_sleep"`
	FragmentSleep   time.Duration `json:"fragment_sleep"`
	FirstTimeSleep  time.Duration `json:"first_time_sleep"`
	LFragment       int           `json:"l_fragment"`
}

type Server struct {
	ListenPort     int    `json:"listen_port"`
	CloudFlareIP   string `json:"cloud_flare_ip"`
	CloudFlarePort int    `json:"cloud_flare_port"`
	Configuration
}

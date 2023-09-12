package connection

import (
	"errors"
	"github.com/spf13/viper"
	"log"
	"time"
)

// InitialConfiguration initializes the server configuration by reading values from Viper and setting defaults.
func (C Server) InitialConfiguration() Server {
	C.ListenPort = viper.GetInt("listen_port")
	C.CloudFlareIP = viper.GetString("cloud_flare_ip")
	C.CloudFlarePort = viper.GetInt("cloud_flare_port")
	C.SocketTimeout = viper.GetDuration("socket_timeout") * time.Second
	C.FragmentSleep = viper.GetDuration("fragment_sleep") * time.Millisecond
	C.LFragment = viper.GetInt("l_fragment")

	err := configErrorHandler(&C)
	if err != nil {
		log.Fatalf("error : %s", err.Error())
	}

	// fill the values if they are not being set in json conf
	if C.SocketTimeout == 0 {
		C.SocketTimeout = 50
	}

	if C.AcceptTimeSleep == 0 {
		C.AcceptTimeSleep = 10
	}

	if C.LFragment == 0 {
		C.LFragment = 80
	}

	return C
}

func configErrorHandler(configuration *Server) error {
	if configuration.CloudFlareIP == "" {
		return errors.New("cloudflare ip are empty")
	}

	if configuration.CloudFlarePort == 0 {
		return errors.New("cloudflare port are empty")
	}

	if configuration.ListenPort == 0 {
		return errors.New("listen port are empty")
	}

	return nil
}

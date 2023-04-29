# gfwKnock
![go]

knock up GFW detection with packet fragmentation method PoC

## Installation

### Getting the Latest Release
You can get the release of gfwKnock from [Releases](https://github.com/SonyaCore/gfwKnock/releases) section

Make sure to properly set configuration file.


### Building from Source
Clone the Repository this by running this command :

```bash
 git clone https://github.com/SonyaCore/gfwKnock.git
```

Now Cd into project directory and use the following command to compile the program and then run the compiled binary :

```bash
cd gfwKnock
go build -o gfwKnock -ldflags "-s -w -buildid=" .
 ./gfwKnock
```

## Configuration 
gfwKnock Uses default a configuration file with the following specifications :

* `listen_port` : port for running gfwKnock client.
* `cloud_flare_ip` : cloudflare ip.
* `cloud_flare_port` : cloudflare port.
* `socket_timeout`: time in seconds for closing the server socket if the socket doesn't respond.
* `fragment_sleep` : sleep in milliseconds between each fragment. 
* `l_fragment` : length of fragments of Client hello packets.

example of config.json :

```bash
{
  // key : value : type
  "listen_port": 8080, // Integer
  "cloud_flare_ip": "45.85.118.88", // String
  "cloud_flare_port": 443, // Integer
  "socket_timeout": "60", // int Seconds
  "fragment_sleep": "100" // int Miliseconds,
  "l_fragment": 77 // int Byte

}
```

## License

Licensed under the [GPL-3][license] license.


[license]: LICENSE
[go]: https://img.shields.io/badge/Go-cyan?logo=go



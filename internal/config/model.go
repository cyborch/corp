package config

type Configurations struct {
	Server       ServerConfigurations
	VirtualHosts []VirtualHostConfigurations
}

type ServerConfigurations struct {
	Port int
}

type VirtualHostConfigurations struct {
	Hostname    string   // fqdn of the virtual host
	Scheme      string   // request scheme
	Origin      string   // origin which the request should be forwarded to
	EnableCors  bool     // add CORS headers if true
	SkipHeaders []string // array of headers which should be removed from origin response
}

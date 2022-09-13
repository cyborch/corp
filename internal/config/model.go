package config

type Configurations struct {
	Server       ServerConfigurations
	VirtualHosts []VirtualHostConfigurations
}

type ServerConfigurations struct {
	Port int
}

type VirtualHostConfigurations struct {
	Hostname    string // fqdn of the virtual host
	Scheme      string
	Origin      string
	EnableCors  bool
	SkipHeaders []string
}

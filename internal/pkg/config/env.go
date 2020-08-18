package config

// ENV ...
type ENV struct {
	// App
	IsDev                  bool
	ZookeeperURI           string
	ZookeeperGRPCPrefix    string
	ZookeeperCommonPrefix  string
	ZookeeperServicePrefix string

	// Reading from Zookeeper
	// App port
	AppPort   string
	AdminPort string

	// gRPC addresses
	GRPCAddresses struct {
		User        string
		Transaction string
	}

	// Database
	DatabaseURI  string
	DatabaseName string

	// RabbitMQ
	RabbitMQURI string

	// Auth
	AuthSecret string

	// File
	FileHost string
}

var env ENV

// GetEnv return .env data
func GetEnv() *ENV {
	return &env
}

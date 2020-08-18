package config

type Zookeeper struct {
	GRPCAddressUser               string
	GRPCAddressTransaction        string
	GRPCAddressTransactionOffline string
	GRPCAddressMerchant           string
	GRPCAddressWithdraw           string
	GRPCAddressReferral           string
	GRPCAddressVoucher            string
	DatabaseURL                   string
	RabbitMQURI                   string
	FileHost                      string
	AdminPort                     string
	AppPort                       string
}

var (
	ZookeeperValue *Zookeeper
)

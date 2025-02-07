package etherman

import "github.com/0xPolygon/zkevm-ethtx-manager/etherman/etherscan"

// Config represents the configuration of the etherman
type Config struct {
	// URL is the URL of the Ethereum node for L1
	URL string `mapstructure:"URL"`

	// allow that L1 gas price calculation use multiples sources
	MultiGasProvider bool `mapstructure:"MultiGasProvider"`
	// Configuration for use Etherscan as used as gas provider, basically it needs the API-KEY
	Etherscan etherscan.Config
	// L1ChainID specifies the chain ID of the network to which transactions will be sent
	// If set to 0, the chain ID will be fetched from the RPC endpoint
	L1ChainID uint64 `mapstructure:"L1ChainID"`
	// HTTPHeaders are the headers to be used in the HTTP requests
	HTTPHeaders map[string]string `mapstructure:"HTTPHeaders"`
}

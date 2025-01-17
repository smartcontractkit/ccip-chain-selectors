// Code generated by go generate please DO NOT EDIT
package chain_selectors

type TronChain struct {
	ChainID  uint64
	Selector uint64
	Name     string
	VarName  string
}

var (
	TRON_MAINNET        = TronChain{ChainID: 728126428, Selector: 1546563616611573946, Name: "tron-mainnet"}
	TRON_TESTNET_NILE   = TronChain{ChainID: 3448148188, Selector: 2052925811360307749, Name: "tron-testnet-nile"}
	TRON_TESTNET_SHASTA = TronChain{ChainID: 2494104990, Selector: 13231703482326770598, Name: "tron-testnet-shasta"}
)

var TronALL = []TronChain{
	TRON_MAINNET,
	TRON_TESTNET_NILE,
	TRON_TESTNET_SHASTA,
}
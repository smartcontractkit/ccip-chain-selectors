// Code generated by go generate please DO NOT EDIT
package chain_selectors

type AptosChain struct {
	ChainID  uint64
	Selector uint64
	Name     string
	VarName  string
}

var (
	APTOS_MAINNET = AptosChain{ChainID: 1, Selector: 4741433654826277614, Name: "aptos-mainnet"}
	APTOS_TESTNET = AptosChain{ChainID: 2, Selector: 743186221051783445, Name: "aptos-testnet"}
)

var AptosALL = []AptosChain{
	APTOS_MAINNET,
	APTOS_TESTNET,
}

package selectors

import "fmt"

var EvmChainIdToChainSelector = map[uint64]uint64{
	// Testnets
	97:       13264668187771770619, // BSC Testnet
	420:      2664363617261496610,  // Optimism Goerli
	1337:     3379446385462418246,  // Dev network
	43113:    14767482510784806043, // Avalanche Fuji
	84531:    5790810961207155433,  // BASE Goerli
	80001:    12532609583862916517, // Polygon Mumbai
	421613:   6101244977088475029,  // Arbitrum Goerli
	11155111: 16015286601757825753, // Sepolia

	// Mainnets
	1:     5009297550715157269,  // Ethereum
	10:    3734403246176062136,  // Optimism
	56:    11344663589394136015, // BSC
	137:   4051577828743386545,  // Polygon
	8453:  15971525489660198786, // BASE
	42161: 4949039107694359620,  // Arbitrum
	43114: 6433500567565415381,  // Avalanche
}

func ChainIdFromSelector(chainSelectorId uint64) (uint64, error) {
	for k, v := range EvmChainIdToChainSelector {
		if v == chainSelectorId {
			return k, nil
		}
	}
	return 0, fmt.Errorf("chain not found for chain selector %d", chainSelectorId)
}

func SelectorFromChainId(chainId uint64) (uint64, error) {
	if chainSelectorId, exist := EvmChainIdToChainSelector[chainId]; exist {
		return chainSelectorId, nil
	}
	return 0, fmt.Errorf("chain selector not found for chain %d", chainId)
}

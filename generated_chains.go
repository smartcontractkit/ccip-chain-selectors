// Code generated by go generate please DO NOT EDIT
package chain_selectors

type Chain struct {
	EvmChainID uint64
	Selector   uint64
	Name       string
	VarName    string
}

var (
	AVALANCHE_MAINNET                        = Chain{EvmChainID: 43114, Selector: 6433500567565415381, Name: "avalanche-mainnet"}
	AVALANCHE_TESTNET_FUJI                   = Chain{EvmChainID: 43113, Selector: 14767482510784806043, Name: "avalanche-testnet-fuji"}
	AVALANCHE_TESTNET_NEXON                  = Chain{EvmChainID: 595581, Selector: 7837562506228496256, Name: "avalanche-testnet-nexon"}
	BERACHAIN_TESTNET_ARTIO                  = Chain{EvmChainID: 80085, Selector: 12336603543561911511, Name: "berachain-testnet-artio"}
	BINANCE_SMART_CHAIN_MAINNET              = Chain{EvmChainID: 56, Selector: 11344663589394136015, Name: "binance_smart_chain-mainnet"}
	BINANCE_SMART_CHAIN_TESTNET              = Chain{EvmChainID: 97, Selector: 13264668187771770619, Name: "binance_smart_chain-testnet"}
	BITCOIN_TESTNET_ROOTSTOCK                = Chain{EvmChainID: 31, Selector: 8953668971247136127, Name: "bitcoin-testnet-rootstock"}
	BITTORRENT_CHAIN_MAINNET                 = Chain{EvmChainID: 199, Selector: 3776006016387883143, Name: "bittorrent_chain-mainnet"}
	BITTORRENT_CHAIN_TESTNET                 = Chain{EvmChainID: 1029, Selector: 4459371029167934217, Name: "bittorrent_chain-testnet"}
	CELO_TESTNET_ALFAJORES                   = Chain{EvmChainID: 44787, Selector: 3552045678561919002, Name: "celo-testnet-alfajores"}
	CRONOS_TESTNET                           = Chain{EvmChainID: 338, Selector: 2995292832068775165, Name: "cronos-testnet"}
	ETHEREUM_MAINNET                         = Chain{EvmChainID: 1, Selector: 5009297550715157269, Name: "ethereum-mainnet"}
	ETHEREUM_MAINNET_ARBITRUM_1              = Chain{EvmChainID: 42161, Selector: 4949039107694359620, Name: "ethereum-mainnet-arbitrum-1"}
	ETHEREUM_MAINNET_BASE_1                  = Chain{EvmChainID: 8453, Selector: 15971525489660198786, Name: "ethereum-mainnet-base-1"}
	ETHEREUM_MAINNET_KROMA_1                 = Chain{EvmChainID: 255, Selector: 3719320017875267166, Name: "ethereum-mainnet-kroma-1"}
	ETHEREUM_MAINNET_LINEA_1                 = Chain{EvmChainID: 59144, Selector: 4627098889531055414, Name: "ethereum-mainnet-linea-1"}
	ETHEREUM_MAINNET_MANTLE_1                = Chain{EvmChainID: 5000, Selector: 1556008542357238666, Name: "ethereum-mainnet-mantle-1"}
	ETHEREUM_MAINNET_OPTIMISM_1              = Chain{EvmChainID: 10, Selector: 3734403246176062136, Name: "ethereum-mainnet-optimism-1"}
	ETHEREUM_MAINNET_POLYGON_ZKEVM_1         = Chain{EvmChainID: 1101, Selector: 4348158687435793198, Name: "ethereum-mainnet-polygon-zkevm-1"}
	ETHEREUM_MAINNET_SCROLL_1                = Chain{EvmChainID: 534352, Selector: 13204309965629103672, Name: "ethereum-mainnet-scroll-1"}
	ETHEREUM_TESTNET_GOERLI_ARBITRUM_1       = Chain{EvmChainID: 421613, Selector: 6101244977088475029, Name: "ethereum-testnet-goerli-arbitrum-1"}
	ETHEREUM_TESTNET_GOERLI_BASE_1           = Chain{EvmChainID: 84531, Selector: 5790810961207155433, Name: "ethereum-testnet-goerli-base-1"}
	ETHEREUM_TESTNET_GOERLI_LINEA_1          = Chain{EvmChainID: 59140, Selector: 1355246678561316402, Name: "ethereum-testnet-goerli-linea-1"}
	ETHEREUM_TESTNET_GOERLI_MANTLE_1         = Chain{EvmChainID: 5001, Selector: 4168263376276232250, Name: "ethereum-testnet-goerli-mantle-1"}
	ETHEREUM_TESTNET_GOERLI_OPTIMISM_1       = Chain{EvmChainID: 420, Selector: 2664363617261496610, Name: "ethereum-testnet-goerli-optimism-1"}
	ETHEREUM_TESTNET_GOERLI_POLYGON_ZKEVM_1  = Chain{EvmChainID: 1442, Selector: 11059667695644972511, Name: "ethereum-testnet-goerli-polygon-zkevm-1"}
	ETHEREUM_TESTNET_GOERLI_ZKSYNC_1         = Chain{EvmChainID: 280, Selector: 6802309497652714138, Name: "ethereum-testnet-goerli-zksync-1"}
	ETHEREUM_TESTNET_HOLESKY_FRAXTAL_1       = Chain{EvmChainID: 2522, Selector: 8901520481741771655, Name: "ethereum-testnet-holesky-fraxtal-1"}
	ETHEREUM_TESTNET_SEPOLIA                 = Chain{EvmChainID: 11155111, Selector: 16015286601757825753, Name: "ethereum-testnet-sepolia"}
	ETHEREUM_TESTNET_SEPOLIA_ARBITRUM_1      = Chain{EvmChainID: 421614, Selector: 3478487238524512106, Name: "ethereum-testnet-sepolia-arbitrum-1"}
	ETHEREUM_TESTNET_SEPOLIA_BASE_1          = Chain{EvmChainID: 84532, Selector: 10344971235874465080, Name: "ethereum-testnet-sepolia-base-1"}
	ETHEREUM_TESTNET_SEPOLIA_BLAST_1         = Chain{EvmChainID: 168587773, Selector: 2027362563942762617, Name: "ethereum-testnet-sepolia-blast-1"}
	ETHEREUM_TESTNET_SEPOLIA_KROMA_1         = Chain{EvmChainID: 2358, Selector: 5990477251245693094, Name: "ethereum-testnet-sepolia-kroma-1"}
	ETHEREUM_TESTNET_SEPOLIA_LINEA_1         = Chain{EvmChainID: 59141, Selector: 5719461335882077547, Name: "ethereum-testnet-sepolia-linea-1"}
	ETHEREUM_TESTNET_SEPOLIA_LISK_1          = Chain{EvmChainID: 4202, Selector: 5298399861320400553, Name: "ethereum-testnet-sepolia-lisk-1"}
	ETHEREUM_TESTNET_SEPOLIA_MANTLE_1        = Chain{EvmChainID: 5003, Selector: 8236463271206331221, Name: "ethereum-testnet-sepolia-mantle-1"}
	ETHEREUM_TESTNET_SEPOLIA_OPTIMISM_1      = Chain{EvmChainID: 11155420, Selector: 5224473277236331295, Name: "ethereum-testnet-sepolia-optimism-1"}
	ETHEREUM_TESTNET_SEPOLIA_POLYGON_ZKEVM_1 = Chain{EvmChainID: 2442, Selector: 1654667687261492630, Name: "ethereum-testnet-sepolia-polygon-zkevm-1"}
	ETHEREUM_TESTNET_SEPOLIA_SCROLL_1        = Chain{EvmChainID: 534351, Selector: 2279865765895943307, Name: "ethereum-testnet-sepolia-scroll-1"}
	ETHEREUM_TESTNET_SEPOLIA_ZKSYNC_1        = Chain{EvmChainID: 300, Selector: 6898391096552792247, Name: "ethereum-testnet-sepolia-zksync-1"}
	FANTOM_TESTNET                           = Chain{EvmChainID: 4002, Selector: 4905564228793744293, Name: "fantom-testnet"}
	GETH_TESTNET                             = Chain{EvmChainID: 1337, Selector: 3379446385462418246, Name: "geth-testnet"}
	GNOSIS_CHAIN_MAINNET                     = Chain{EvmChainID: 100, Selector: 465200170687744372, Name: "gnosis_chain-mainnet"}
	GNOSIS_CHAIN_TESTNET_CHIADO              = Chain{EvmChainID: 10200, Selector: 8871595565390010547, Name: "gnosis_chain-testnet-chiado"}
	HEDERA_TESTNET                           = Chain{EvmChainID: 296, Selector: 222782988166878823, Name: "hedera-testnet"}
	KAVA_MAINNET                             = Chain{EvmChainID: 2222, Selector: 7550000543357438061, Name: "kava-mainnet"}
	KAVA_TESTNET                             = Chain{EvmChainID: 2221, Selector: 2110537777356199208, Name: "kava-testnet"}
	POLYGON_MAINNET                          = Chain{EvmChainID: 137, Selector: 4051577828743386545, Name: "polygon-mainnet"}
	POLYGON_TESTNET_AMOY                     = Chain{EvmChainID: 80002, Selector: 16281711391670634445, Name: "polygon-testnet-amoy"}
	POLYGON_TESTNET_MUMBAI                   = Chain{EvmChainID: 80001, Selector: 12532609583862916517, Name: "polygon-testnet-mumbai"}
	TEST_1000                                = Chain{EvmChainID: 1000, Selector: 11787463284727550157, Name: "1000"}
	TEST_2337                                = Chain{EvmChainID: 2337, Selector: 12922642891491394802, Name: "2337"}
	TEST_76578                               = Chain{EvmChainID: 76578, Selector: 781901677223027175, Name: "76578"}
	TEST_90000001                            = Chain{EvmChainID: 90000001, Selector: 909606746561742123, Name: "90000001"}
	TEST_90000002                            = Chain{EvmChainID: 90000002, Selector: 5548718428018410741, Name: "90000002"}
	TEST_90000003                            = Chain{EvmChainID: 90000003, Selector: 789068866484373046, Name: "90000003"}
	TEST_90000004                            = Chain{EvmChainID: 90000004, Selector: 5721565186521185178, Name: "90000004"}
	TEST_90000005                            = Chain{EvmChainID: 90000005, Selector: 964127714438319834, Name: "90000005"}
	TEST_90000006                            = Chain{EvmChainID: 90000006, Selector: 8966794841936584464, Name: "90000006"}
	TEST_90000007                            = Chain{EvmChainID: 90000007, Selector: 8412806778050735057, Name: "90000007"}
	TEST_90000008                            = Chain{EvmChainID: 90000008, Selector: 4066443121807923198, Name: "90000008"}
	TEST_90000009                            = Chain{EvmChainID: 90000009, Selector: 6747736380229414777, Name: "90000009"}
	TEST_90000010                            = Chain{EvmChainID: 90000010, Selector: 8694984074292254623, Name: "90000010"}
	TEST_90000011                            = Chain{EvmChainID: 90000011, Selector: 328334718812072308, Name: "90000011"}
	TEST_90000012                            = Chain{EvmChainID: 90000012, Selector: 7715160997071429212, Name: "90000012"}
	TEST_90000013                            = Chain{EvmChainID: 90000013, Selector: 3574539439524578558, Name: "90000013"}
	TEST_90000014                            = Chain{EvmChainID: 90000014, Selector: 4543928599863227519, Name: "90000014"}
	TEST_90000015                            = Chain{EvmChainID: 90000015, Selector: 6443235356619661032, Name: "90000015"}
	TEST_90000016                            = Chain{EvmChainID: 90000016, Selector: 13087962012083037329, Name: "90000016"}
	TEST_90000017                            = Chain{EvmChainID: 90000017, Selector: 11985232338641871056, Name: "90000017"}
	TEST_90000018                            = Chain{EvmChainID: 90000018, Selector: 7777066535355430289, Name: "90000018"}
	TEST_90000019                            = Chain{EvmChainID: 90000019, Selector: 1273605685587320666, Name: "90000019"}
	TEST_90000020                            = Chain{EvmChainID: 90000020, Selector: 17810359353458878177, Name: "90000020"}
	WEMIX_MAINNET                            = Chain{EvmChainID: 1111, Selector: 5142893604156789321, Name: "wemix-mainnet"}
	WEMIX_TESTNET                            = Chain{EvmChainID: 1112, Selector: 9284632837123596123, Name: "wemix-testnet"}
)

var ALL = []Chain{
	AVALANCHE_MAINNET,
	AVALANCHE_TESTNET_FUJI,
	AVALANCHE_TESTNET_NEXON,
	BERACHAIN_TESTNET_ARTIO,
	BINANCE_SMART_CHAIN_MAINNET,
	BINANCE_SMART_CHAIN_TESTNET,
	BITCOIN_TESTNET_ROOTSTOCK,
	BITTORRENT_CHAIN_MAINNET,
	BITTORRENT_CHAIN_TESTNET,
	CELO_TESTNET_ALFAJORES,
	CRONOS_TESTNET,
	ETHEREUM_MAINNET,
	ETHEREUM_MAINNET_ARBITRUM_1,
	ETHEREUM_MAINNET_BASE_1,
	ETHEREUM_MAINNET_KROMA_1,
	ETHEREUM_MAINNET_LINEA_1,
	ETHEREUM_MAINNET_MANTLE_1,
	ETHEREUM_MAINNET_OPTIMISM_1,
	ETHEREUM_MAINNET_POLYGON_ZKEVM_1,
	ETHEREUM_MAINNET_SCROLL_1,
	ETHEREUM_TESTNET_GOERLI_ARBITRUM_1,
	ETHEREUM_TESTNET_GOERLI_BASE_1,
	ETHEREUM_TESTNET_GOERLI_LINEA_1,
	ETHEREUM_TESTNET_GOERLI_MANTLE_1,
	ETHEREUM_TESTNET_GOERLI_OPTIMISM_1,
	ETHEREUM_TESTNET_GOERLI_POLYGON_ZKEVM_1,
	ETHEREUM_TESTNET_GOERLI_ZKSYNC_1,
	ETHEREUM_TESTNET_HOLESKY_FRAXTAL_1,
	ETHEREUM_TESTNET_SEPOLIA,
	ETHEREUM_TESTNET_SEPOLIA_ARBITRUM_1,
	ETHEREUM_TESTNET_SEPOLIA_BASE_1,
	ETHEREUM_TESTNET_SEPOLIA_BLAST_1,
	ETHEREUM_TESTNET_SEPOLIA_KROMA_1,
	ETHEREUM_TESTNET_SEPOLIA_LINEA_1,
	ETHEREUM_TESTNET_SEPOLIA_LISK_1,
	ETHEREUM_TESTNET_SEPOLIA_MANTLE_1,
	ETHEREUM_TESTNET_SEPOLIA_OPTIMISM_1,
	ETHEREUM_TESTNET_SEPOLIA_POLYGON_ZKEVM_1,
	ETHEREUM_TESTNET_SEPOLIA_SCROLL_1,
	ETHEREUM_TESTNET_SEPOLIA_ZKSYNC_1,
	FANTOM_TESTNET,
	GETH_TESTNET,
	GNOSIS_CHAIN_MAINNET,
	GNOSIS_CHAIN_TESTNET_CHIADO,
	HEDERA_TESTNET,
	KAVA_MAINNET,
	KAVA_TESTNET,
	POLYGON_MAINNET,
	POLYGON_TESTNET_AMOY,
	POLYGON_TESTNET_MUMBAI,
	TEST_1000,
	TEST_2337,
	TEST_76578,
	TEST_90000001,
	TEST_90000002,
	TEST_90000003,
	TEST_90000004,
	TEST_90000005,
	TEST_90000006,
	TEST_90000007,
	TEST_90000008,
	TEST_90000009,
	TEST_90000010,
	TEST_90000011,
	TEST_90000012,
	TEST_90000013,
	TEST_90000014,
	TEST_90000015,
	TEST_90000016,
	TEST_90000017,
	TEST_90000018,
	TEST_90000019,
	TEST_90000020,
	WEMIX_MAINNET,
	WEMIX_TESTNET,
}

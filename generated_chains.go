// Code generated by go generate please DO NOT EDIT
package chain_selectors

type Chain struct {
	Selector uint64
	Name     string
	ChainID  string
	Family   string
	VarName  string
}

var (
	AREON_MAINNET                                  = Chain{ChainID: "463", Selector: 1939936305787790600, Family: "evm", Name: "areon-mainnet"}
	AREON_TESTNET                                  = Chain{ChainID: "462", Selector: 7317911323415911000, Family: "evm", Name: "areon-testnet"}
	AVALANCHE_MAINNET                              = Chain{ChainID: "43114", Selector: 6433500567565415381, Family: "evm", Name: "avalanche-mainnet"}
	AVALANCHE_SUBNET_DEXALOT_MAINNET               = Chain{ChainID: "432204", Selector: 5463201557265485081, Family: "evm", Name: "avalanche-subnet-dexalot-mainnet"}
	AVALANCHE_SUBNET_DEXALOT_TESTNET               = Chain{ChainID: "432201", Selector: 1458281248224512906, Family: "evm", Name: "avalanche-subnet-dexalot-testnet"}
	AVALANCHE_TESTNET_FUJI                         = Chain{ChainID: "43113", Selector: 14767482510784806043, Family: "evm", Name: "avalanche-testnet-fuji"}
	AVALANCHE_TESTNET_NEXON                        = Chain{ChainID: "595581", Selector: 7837562506228496256, Family: "evm", Name: "avalanche-testnet-nexon"}
	BERACHAIN_TESTNET_ARTIO                        = Chain{ChainID: "80085", Selector: 12336603543561911511, Family: "evm", Name: "berachain-testnet-artio"}
	BERACHAIN_TESTNET_BARTIO                       = Chain{ChainID: "80084", Selector: 8999465244383784164, Family: "evm", Name: "berachain-testnet-bartio"}
	BINANCE_SMART_CHAIN_MAINNET                    = Chain{ChainID: "56", Selector: 11344663589394136015, Family: "evm", Name: "binance_smart_chain-mainnet"}
	BINANCE_SMART_CHAIN_MAINNET_OPBNB_1            = Chain{ChainID: "204", Selector: 465944652040885897, Family: "evm", Name: "binance_smart_chain-mainnet-opbnb-1"}
	BINANCE_SMART_CHAIN_TESTNET                    = Chain{ChainID: "97", Selector: 13264668187771770619, Family: "evm", Name: "binance_smart_chain-testnet"}
	BINANCE_SMART_CHAIN_TESTNET_OPBNB_1            = Chain{ChainID: "5611", Selector: 13274425992935471758, Family: "evm", Name: "binance_smart_chain-testnet-opbnb-1"}
	BITCICHAIN_MAINNET                             = Chain{ChainID: "1907", Selector: 4874388048629246000, Family: "evm", Name: "bitcichain-mainnet"}
	BITCICHAIN_TESTNET                             = Chain{ChainID: "1908", Selector: 4888058894222120000, Family: "evm", Name: "bitcichain-testnet"}
	BITCOIN_MAINNET_BITLAYER_1                     = Chain{ChainID: "200901", Selector: 7937294810946806131, Family: "evm", Name: "bitcoin-mainnet-bitlayer-1"}
	BITCOIN_MAINNET_BOB_1                          = Chain{ChainID: "60808", Selector: 3849287863852499584, Family: "evm", Name: "bitcoin-mainnet-bob-1"}
	BITCOIN_MAINNET_BOTANIX                        = Chain{ChainID: "3637", Selector: 4560701533377838164, Family: "evm", Name: "bitcoin-mainnet-botanix"}
	BITCOIN_MAINNET_BSQUARED_1                     = Chain{ChainID: "223", Selector: 5406759801798337480, Family: "evm", Name: "bitcoin-mainnet-bsquared-1"}
	BITCOIN_MERLIN_MAINNET                         = Chain{ChainID: "4200", Selector: 241851231317828981, Family: "evm", Name: "bitcoin-merlin-mainnet"}
	BITCOIN_TESTNET_BITLAYER_1                     = Chain{ChainID: "200810", Selector: 3789623672476206327, Family: "evm", Name: "bitcoin-testnet-bitlayer-1"}
	BITCOIN_TESTNET_BOTANIX                        = Chain{ChainID: "3636", Selector: 1467223411771711614, Family: "evm", Name: "bitcoin-testnet-botanix"}
	BITCOIN_TESTNET_BSQUARED_1                     = Chain{ChainID: "1123", Selector: 1948510578179542068, Family: "evm", Name: "bitcoin-testnet-bsquared-1"}
	BITCOIN_TESTNET_MERLIN                         = Chain{ChainID: "686868", Selector: 5269261765892944301, Family: "evm", Name: "bitcoin-testnet-merlin"}
	BITCOIN_TESTNET_ROOTSTOCK                      = Chain{ChainID: "31", Selector: 8953668971247136127, Family: "evm", Name: "bitcoin-testnet-rootstock"}
	BITCOIN_TESTNET_SEPOLIA_BOB_1                  = Chain{ChainID: "808813", Selector: 5535534526963509396, Family: "evm", Name: "bitcoin-testnet-sepolia-bob-1"}
	BITTORRENT_CHAIN_MAINNET                       = Chain{ChainID: "199", Selector: 3776006016387883143, Family: "evm", Name: "bittorrent_chain-mainnet"}
	BITTORRENT_CHAIN_TESTNET                       = Chain{ChainID: "1029", Selector: 4459371029167934217, Family: "evm", Name: "bittorrent_chain-testnet"}
	CELO_MAINNET                                   = Chain{ChainID: "42220", Selector: 1346049177634351622, Family: "evm", Name: "celo-mainnet"}
	CELO_TESTNET_ALFAJORES                         = Chain{ChainID: "44787", Selector: 3552045678561919002, Family: "evm", Name: "celo-testnet-alfajores"}
	COINEX_SMART_CHAIN_MAINNET                     = Chain{ChainID: "52", Selector: 1761333065194157300, Family: "evm", Name: "coinex_smart_chain-mainnet"}
	COINEX_SMART_CHAIN_TESTNET                     = Chain{ChainID: "53", Selector: 8955032871639343000, Family: "evm", Name: "coinex_smart_chain-testnet"}
	CRONOS_MAINNET                                 = Chain{ChainID: "25", Selector: 1456215246176062136, Family: "evm", Name: "cronos-mainnet"}
	CRONOS_TESTNET                                 = Chain{ChainID: "338", Selector: 2995292832068775165, Family: "evm", Name: "cronos-testnet"}
	CRONOS_TESTNET_ZKEVM_1                         = Chain{ChainID: "282", Selector: 3842103497652714138, Family: "evm", Name: "cronos-testnet-zkevm-1"}
	ETHEREUM_MAINNET                               = Chain{ChainID: "1", Selector: 5009297550715157269, Family: "evm", Name: "ethereum-mainnet"}
	ETHEREUM_MAINNET_ARBITRUM_1                    = Chain{ChainID: "42161", Selector: 4949039107694359620, Family: "evm", Name: "ethereum-mainnet-arbitrum-1"}
	ETHEREUM_MAINNET_ARBITRUM_1_L3X_1              = Chain{ChainID: "12324", Selector: 3162193654116181371, Family: "evm", Name: "ethereum-mainnet-arbitrum-1-l3x-1"}
	ETHEREUM_MAINNET_ARBITRUM_1_TREASURE_1         = Chain{ChainID: "978670", Selector: 1010349088906777999, Family: "evm", Name: "ethereum-mainnet-arbitrum-1-treasure-1"}
	ETHEREUM_MAINNET_ASTAR_ZKEVM_1                 = Chain{ChainID: "3776", Selector: 1540201334317828111, Family: "evm", Name: "ethereum-mainnet-astar-zkevm-1"}
	ETHEREUM_MAINNET_BASE_1                        = Chain{ChainID: "8453", Selector: 15971525489660198786, Family: "evm", Name: "ethereum-mainnet-base-1"}
	ETHEREUM_MAINNET_BLAST_1                       = Chain{ChainID: "81457", Selector: 4411394078118774322, Family: "evm", Name: "ethereum-mainnet-blast-1"}
	ETHEREUM_MAINNET_IMMUTABLE_ZKEVM_1             = Chain{ChainID: "13371", Selector: 1237925231416731909, Family: "evm", Name: "ethereum-mainnet-immutable-zkevm-1"}
	ETHEREUM_MAINNET_KROMA_1                       = Chain{ChainID: "255", Selector: 3719320017875267166, Family: "evm", Name: "ethereum-mainnet-kroma-1"}
	ETHEREUM_MAINNET_LINEA_1                       = Chain{ChainID: "59144", Selector: 4627098889531055414, Family: "evm", Name: "ethereum-mainnet-linea-1"}
	ETHEREUM_MAINNET_MANTLE_1                      = Chain{ChainID: "5000", Selector: 1556008542357238666, Family: "evm", Name: "ethereum-mainnet-mantle-1"}
	ETHEREUM_MAINNET_METIS_1                       = Chain{ChainID: "1088", Selector: 8805746078405598895, Family: "evm", Name: "ethereum-mainnet-metis-1"}
	ETHEREUM_MAINNET_MODE_1                        = Chain{ChainID: "34443", Selector: 7264351850409363825, Family: "evm", Name: "ethereum-mainnet-mode-1"}
	ETHEREUM_MAINNET_OPTIMISM_1                    = Chain{ChainID: "10", Selector: 3734403246176062136, Family: "evm", Name: "ethereum-mainnet-optimism-1"}
	ETHEREUM_MAINNET_POLYGON_ZKEVM_1               = Chain{ChainID: "1101", Selector: 4348158687435793198, Family: "evm", Name: "ethereum-mainnet-polygon-zkevm-1"}
	ETHEREUM_MAINNET_SCROLL_1                      = Chain{ChainID: "534352", Selector: 13204309965629103672, Family: "evm", Name: "ethereum-mainnet-scroll-1"}
	ETHEREUM_MAINNET_TAIKO_1                       = Chain{ChainID: "16700", Selector: 16468599424800719238, Family: "evm", Name: "ethereum-mainnet-taiko-1"}
	ETHEREUM_MAINNET_WORLDCHAIN_1                  = Chain{ChainID: "480", Selector: 2049429975587534727, Family: "evm", Name: "ethereum-mainnet-worldchain-1"}
	ETHEREUM_MAINNET_XLAYER_1                      = Chain{ChainID: "196", Selector: 3016212468291539606, Family: "evm", Name: "ethereum-mainnet-xlayer-1"}
	ETHEREUM_MAINNET_ZKSYNC_1                      = Chain{ChainID: "324", Selector: 1562403441176082196, Family: "evm", Name: "ethereum-mainnet-zksync-1"}
	ETHEREUM_TESTNET_GOERLI_ARBITRUM_1             = Chain{ChainID: "421613", Selector: 6101244977088475029, Family: "evm", Name: "ethereum-testnet-goerli-arbitrum-1"}
	ETHEREUM_TESTNET_GOERLI_BASE_1                 = Chain{ChainID: "84531", Selector: 5790810961207155433, Family: "evm", Name: "ethereum-testnet-goerli-base-1"}
	ETHEREUM_TESTNET_GOERLI_LINEA_1                = Chain{ChainID: "59140", Selector: 1355246678561316402, Family: "evm", Name: "ethereum-testnet-goerli-linea-1"}
	ETHEREUM_TESTNET_GOERLI_MANTLE_1               = Chain{ChainID: "5001", Selector: 4168263376276232250, Family: "evm", Name: "ethereum-testnet-goerli-mantle-1"}
	ETHEREUM_TESTNET_GOERLI_OPTIMISM_1             = Chain{ChainID: "420", Selector: 2664363617261496610, Family: "evm", Name: "ethereum-testnet-goerli-optimism-1"}
	ETHEREUM_TESTNET_GOERLI_POLYGON_ZKEVM_1        = Chain{ChainID: "1442", Selector: 11059667695644972511, Family: "evm", Name: "ethereum-testnet-goerli-polygon-zkevm-1"}
	ETHEREUM_TESTNET_GOERLI_ZKSYNC_1               = Chain{ChainID: "280", Selector: 6802309497652714138, Family: "evm", Name: "ethereum-testnet-goerli-zksync-1"}
	ETHEREUM_TESTNET_HOLESKY                       = Chain{ChainID: "17000", Selector: 7717148896336251131, Family: "evm", Name: "ethereum-testnet-holesky"}
	ETHEREUM_TESTNET_HOLESKY_FRAXTAL_1             = Chain{ChainID: "2522", Selector: 8901520481741771655, Family: "evm", Name: "ethereum-testnet-holesky-fraxtal-1"}
	ETHEREUM_TESTNET_HOLESKY_MORPH_1               = Chain{ChainID: "2810", Selector: 8304510386741731151, Family: "evm", Name: "ethereum-testnet-holesky-morph-1"}
	ETHEREUM_TESTNET_HOLESKY_TAIKO_1               = Chain{ChainID: "167009", Selector: 7248756420937879088, Family: "evm", Name: "ethereum-testnet-holesky-taiko-1"}
	ETHEREUM_TESTNET_SEPOLIA                       = Chain{ChainID: "11155111", Selector: 16015286601757825753, Family: "evm", Name: "ethereum-testnet-sepolia"}
	ETHEREUM_TESTNET_SEPOLIA_ARBITRUM_1            = Chain{ChainID: "421614", Selector: 3478487238524512106, Family: "evm", Name: "ethereum-testnet-sepolia-arbitrum-1"}
	ETHEREUM_TESTNET_SEPOLIA_ARBITRUM_1_L3X_1      = Chain{ChainID: "12325", Selector: 3486622437121596122, Family: "evm", Name: "ethereum-testnet-sepolia-arbitrum-1-l3x-1"}
	ETHEREUM_TESTNET_SEPOLIA_ARBITRUM_1_TREASURE_1 = Chain{ChainID: "978657", Selector: 10443705513486043421, Family: "evm", Name: "ethereum-testnet-sepolia-arbitrum-1-treasure-1"}
	ETHEREUM_TESTNET_SEPOLIA_BASE_1                = Chain{ChainID: "84532", Selector: 10344971235874465080, Family: "evm", Name: "ethereum-testnet-sepolia-base-1"}
	ETHEREUM_TESTNET_SEPOLIA_BLAST_1               = Chain{ChainID: "168587773", Selector: 2027362563942762617, Family: "evm", Name: "ethereum-testnet-sepolia-blast-1"}
	ETHEREUM_TESTNET_SEPOLIA_CORN_1                = Chain{ChainID: "21000000", Selector: 1467427327723633929, Family: "evm", Name: "ethereum-testnet-sepolia-corn-1"}
	ETHEREUM_TESTNET_SEPOLIA_IMMUTABLE_ZKEVM_1     = Chain{ChainID: "13473", Selector: 4526165231216331901, Family: "evm", Name: "ethereum-testnet-sepolia-immutable-zkevm-1"}
	ETHEREUM_TESTNET_SEPOLIA_KROMA_1               = Chain{ChainID: "2358", Selector: 5990477251245693094, Family: "evm", Name: "ethereum-testnet-sepolia-kroma-1"}
	ETHEREUM_TESTNET_SEPOLIA_LENS_1                = Chain{ChainID: "37111", Selector: 6827576821754315911, Family: "evm", Name: "ethereum-testnet-sepolia-lens-1"}
	ETHEREUM_TESTNET_SEPOLIA_LINEA_1               = Chain{ChainID: "59141", Selector: 5719461335882077547, Family: "evm", Name: "ethereum-testnet-sepolia-linea-1"}
	ETHEREUM_TESTNET_SEPOLIA_LISK_1                = Chain{ChainID: "4202", Selector: 5298399861320400553, Family: "evm", Name: "ethereum-testnet-sepolia-lisk-1"}
	ETHEREUM_TESTNET_SEPOLIA_MANTLE_1              = Chain{ChainID: "5003", Selector: 8236463271206331221, Family: "evm", Name: "ethereum-testnet-sepolia-mantle-1"}
	ETHEREUM_TESTNET_SEPOLIA_METIS_1               = Chain{ChainID: "59902", Selector: 3777822886988675105, Family: "evm", Name: "ethereum-testnet-sepolia-metis-1"}
	ETHEREUM_TESTNET_SEPOLIA_MODE_1                = Chain{ChainID: "919", Selector: 829525985033418733, Family: "evm", Name: "ethereum-testnet-sepolia-mode-1"}
	ETHEREUM_TESTNET_SEPOLIA_OPTIMISM_1            = Chain{ChainID: "11155420", Selector: 5224473277236331295, Family: "evm", Name: "ethereum-testnet-sepolia-optimism-1"}
	ETHEREUM_TESTNET_SEPOLIA_POLYGON_VALIDIUM_1    = Chain{ChainID: "717160", Selector: 4418231248214522936, Family: "evm", Name: "ethereum-testnet-sepolia-polygon-validium-1"}
	ETHEREUM_TESTNET_SEPOLIA_POLYGON_ZKEVM_1       = Chain{ChainID: "2442", Selector: 1654667687261492630, Family: "evm", Name: "ethereum-testnet-sepolia-polygon-zkevm-1"}
	ETHEREUM_TESTNET_SEPOLIA_SCROLL_1              = Chain{ChainID: "534351", Selector: 2279865765895943307, Family: "evm", Name: "ethereum-testnet-sepolia-scroll-1"}
	ETHEREUM_TESTNET_SEPOLIA_SONEIUM_1             = Chain{ChainID: "1946", Selector: 686603546605904534, Family: "evm", Name: "ethereum-testnet-sepolia-soneium-1"}
	ETHEREUM_TESTNET_SEPOLIA_UNICHAIN_1            = Chain{ChainID: "1301", Selector: 14135854469784514356, Family: "evm", Name: "ethereum-testnet-sepolia-unichain-1"}
	ETHEREUM_TESTNET_SEPOLIA_WORLDCHAIN_1          = Chain{ChainID: "4801", Selector: 5299555114858065850, Family: "evm", Name: "ethereum-testnet-sepolia-worldchain-1"}
	ETHEREUM_TESTNET_SEPOLIA_XLAYER_1              = Chain{ChainID: "195", Selector: 2066098519157881736, Family: "evm", Name: "ethereum-testnet-sepolia-xlayer-1"}
	ETHEREUM_TESTNET_SEPOLIA_ZIRCUIT_1             = Chain{ChainID: "48899", Selector: 4562743618362911021, Family: "evm", Name: "ethereum-testnet-sepolia-zircuit-1"}
	ETHEREUM_TESTNET_SEPOLIA_ZKSYNC_1              = Chain{ChainID: "300", Selector: 6898391096552792247, Family: "evm", Name: "ethereum-testnet-sepolia-zksync-1"}
	FANTOM_MAINNET                                 = Chain{ChainID: "250", Selector: 3768048213127883732, Family: "evm", Name: "fantom-mainnet"}
	FANTOM_TESTNET                                 = Chain{ChainID: "4002", Selector: 4905564228793744293, Family: "evm", Name: "fantom-testnet"}
	FILECOIN_MAINNET                               = Chain{ChainID: "314", Selector: 4561443241176882990, Family: "evm", Name: "filecoin-mainnet"}
	FILECOIN_TESTNET                               = Chain{ChainID: "31415926", Selector: 7060342227814389000, Family: "evm", Name: "filecoin-testnet"}
	FRAXTAL_MAINNET                                = Chain{ChainID: "252", Selector: 1462016016387883143, Family: "evm", Name: "fraxtal-mainnet"}
	GETH_TESTNET                                   = Chain{ChainID: "1337", Selector: 3379446385462418246, Family: "evm", Name: "geth-testnet"}
	GNOSIS_CHAIN_MAINNET                           = Chain{ChainID: "100", Selector: 465200170687744372, Family: "evm", Name: "gnosis_chain-mainnet"}
	GNOSIS_CHAIN_TESTNET_CHIADO                    = Chain{ChainID: "10200", Selector: 8871595565390010547, Family: "evm", Name: "gnosis_chain-testnet-chiado"}
	HEDERA_MAINNET                                 = Chain{ChainID: "295", Selector: 3229138320728879060, Family: "evm", Name: "hedera-mainnet"}
	HEDERA_TESTNET                                 = Chain{ChainID: "296", Selector: 222782988166878823, Family: "evm", Name: "hedera-testnet"}
	HYPERLIQUID_TESTNET                            = Chain{ChainID: "998", Selector: 4286062357653186312, Family: "evm", Name: "hyperliquid-testnet"}
	KAVA_MAINNET                                   = Chain{ChainID: "2222", Selector: 7550000543357438061, Family: "evm", Name: "kava-mainnet"}
	KAVA_TESTNET                                   = Chain{ChainID: "2221", Selector: 2110537777356199208, Family: "evm", Name: "kava-testnet"}
	KUSAMA_MAINNET_MOONRIVER                       = Chain{ChainID: "1285", Selector: 1355020143337428062, Family: "evm", Name: "kusama-mainnet-moonriver"}
	NEAR_MAINNET                                   = Chain{ChainID: "397", Selector: 2039744413822257700, Family: "evm", Name: "near-mainnet"}
	NEAR_TESTNET                                   = Chain{ChainID: "398", Selector: 5061593697262339000, Family: "evm", Name: "near-testnet"}
	NEONLINK_MAINNET                               = Chain{ChainID: "259", Selector: 8239338020728974000, Family: "evm", Name: "neonlink-mainnet"}
	NEONLINK_TESTNET                               = Chain{ChainID: "9559", Selector: 1113014352258747600, Family: "evm", Name: "neonlink-testnet"}
	PLUME_TESTNET                                  = Chain{ChainID: "161221135", Selector: 14684575664602284776, Family: "evm", Name: "plume-testnet"}
	POLKADOT_MAINNET_ASTAR                         = Chain{ChainID: "592", Selector: 6422105447186081193, Family: "evm", Name: "polkadot-mainnet-astar"}
	POLKADOT_MAINNET_CENTRIFUGE                    = Chain{ChainID: "2031", Selector: 8175830712062617656, Family: "evm", Name: "polkadot-mainnet-centrifuge"}
	POLKADOT_MAINNET_DARWINIA                      = Chain{ChainID: "46", Selector: 8866418665544333000, Family: "evm", Name: "polkadot-mainnet-darwinia"}
	POLKADOT_MAINNET_MOONBEAM                      = Chain{ChainID: "1284", Selector: 1252863800116739621, Family: "evm", Name: "polkadot-mainnet-moonbeam"}
	POLKADOT_TESTNET_ASTAR_SHIBUYA                 = Chain{ChainID: "81", Selector: 6955638871347136141, Family: "evm", Name: "polkadot-testnet-astar-shibuya"}
	POLKADOT_TESTNET_CENTRIFUGE_ALTAIR             = Chain{ChainID: "2088", Selector: 2333097300889804761, Family: "evm", Name: "polkadot-testnet-centrifuge-altair"}
	POLKADOT_TESTNET_DARWINIA_PANGORO              = Chain{ChainID: "45", Selector: 4340886533089894000, Family: "evm", Name: "polkadot-testnet-darwinia-pangoro"}
	POLKADOT_TESTNET_MOONBEAM_MOONBASE             = Chain{ChainID: "1287", Selector: 5361632739113536121, Family: "evm", Name: "polkadot-testnet-moonbeam-moonbase"}
	POLYGON_MAINNET                                = Chain{ChainID: "137", Selector: 4051577828743386545, Family: "evm", Name: "polygon-mainnet"}
	POLYGON_TESTNET_AMOY                           = Chain{ChainID: "80002", Selector: 16281711391670634445, Family: "evm", Name: "polygon-testnet-amoy"}
	POLYGON_TESTNET_MUMBAI                         = Chain{ChainID: "80001", Selector: 12532609583862916517, Family: "evm", Name: "polygon-testnet-mumbai"}
	PRIVATE_TESTNET_MICA                           = Chain{ChainID: "424242", Selector: 4489326297382772450, Family: "evm", Name: "private-testnet-mica"}
	PRIVATE_TESTNET_OPALA                          = Chain{ChainID: "45439", Selector: 8446413392851542429, Family: "evm", Name: "private-testnet-opala"}
	RONIN_MAINNET                                  = Chain{ChainID: "2020", Selector: 6916147374840168594, Family: "evm", Name: "ronin-mainnet"}
	RONIN_TESTNET_SAIGON                           = Chain{ChainID: "2021", Selector: 13116810400804392105, Family: "evm", Name: "ronin-testnet-saigon"}
	SEI_MAINNET                                    = Chain{ChainID: "1329", Selector: 9027416829622342829, Family: "evm", Name: "sei-mainnet"}
	SEI_TESTNET_ATLANTIC                           = Chain{ChainID: "1328", Selector: 1216300075444106652, Family: "evm", Name: "sei-testnet-atlantic"}
	SOLANA_DEVNET                                  = Chain{ChainID: "EtWTRABZaYq6iMfeYKouRu166VU2xqa1wcaWoxPkrZBG", Selector: 16423721717087811551, Family: "solana", Name: "solana-devnet"}
	SOLANA_MAINNET                                 = Chain{ChainID: "5eykt4UsFv8P8NJdTREpY1vzqKqZKvdpKuc147dw2N9d", Selector: 124615329519749607, Family: "solana", Name: "solana-mainnet"}
	SOLANA_TESTNET                                 = Chain{ChainID: "4uhcVJyU9pJkvQyS88uRDiswHXSCkY3zQawwpjk2NsNY", Selector: 6302590918974934319, Family: "solana", Name: "solana-testnet"}
	SONIC_TESTNET                                  = Chain{ChainID: "64165", Selector: 3676871237479449268, Family: "evm", Name: "sonic-testnet"}
	STORY_TESTNET                                  = Chain{ChainID: "1513", Selector: 4237030917318060427, Family: "evm", Name: "story-testnet"}
	TELOS_EVM_MAINNET                              = Chain{ChainID: "40", Selector: 1477345371608778000, Family: "evm", Name: "telos-evm-mainnet"}
	TELOS_EVM_TESTNET                              = Chain{ChainID: "41", Selector: 729797994450396300, Family: "evm", Name: "telos-evm-testnet"}
	TEST_781901677223027175                        = Chain{ChainID: "", Selector: 781901677223027175, Family: "evm", Name: ""}
	VELAS_MAINNET                                  = Chain{ChainID: "106", Selector: 374210358663784372, Family: "evm", Name: "velas-mainnet"}
	VELAS_TESTNET                                  = Chain{ChainID: "111", Selector: 572210378683744374, Family: "evm", Name: "velas-testnet"}
	WEMIX_MAINNET                                  = Chain{ChainID: "1111", Selector: 5142893604156789321, Family: "evm", Name: "wemix-mainnet"}
	WEMIX_TESTNET                                  = Chain{ChainID: "1112", Selector: 9284632837123596123, Family: "evm", Name: "wemix-testnet"}
	ZIRCUIT_MAINNET                                = Chain{ChainID: "48900", Selector: 17198166215261833993, Family: "evm", Name: "zircuit-mainnet"}
	ZKLINK_NOVA_MAINNET                            = Chain{ChainID: "810180", Selector: 4350319965322101699, Family: "evm", Name: "zklink_nova-mainnet"}
	ZKLINK_NOVA_TESTNET                            = Chain{ChainID: "810181", Selector: 5837261596322416298, Family: "evm", Name: "zklink_nova-testnet"}
)

var ALL = []Chain{
	AREON_MAINNET,
	AREON_TESTNET,
	AVALANCHE_MAINNET,
	AVALANCHE_SUBNET_DEXALOT_MAINNET,
	AVALANCHE_SUBNET_DEXALOT_TESTNET,
	AVALANCHE_TESTNET_FUJI,
	AVALANCHE_TESTNET_NEXON,
	BERACHAIN_TESTNET_ARTIO,
	BERACHAIN_TESTNET_BARTIO,
	BINANCE_SMART_CHAIN_MAINNET,
	BINANCE_SMART_CHAIN_MAINNET_OPBNB_1,
	BINANCE_SMART_CHAIN_TESTNET,
	BINANCE_SMART_CHAIN_TESTNET_OPBNB_1,
	BITCICHAIN_MAINNET,
	BITCICHAIN_TESTNET,
	BITCOIN_MAINNET_BITLAYER_1,
	BITCOIN_MAINNET_BOB_1,
	BITCOIN_MAINNET_BOTANIX,
	BITCOIN_MAINNET_BSQUARED_1,
	BITCOIN_MERLIN_MAINNET,
	BITCOIN_TESTNET_BITLAYER_1,
	BITCOIN_TESTNET_BOTANIX,
	BITCOIN_TESTNET_BSQUARED_1,
	BITCOIN_TESTNET_MERLIN,
	BITCOIN_TESTNET_ROOTSTOCK,
	BITCOIN_TESTNET_SEPOLIA_BOB_1,
	BITTORRENT_CHAIN_MAINNET,
	BITTORRENT_CHAIN_TESTNET,
	CELO_MAINNET,
	CELO_TESTNET_ALFAJORES,
	COINEX_SMART_CHAIN_MAINNET,
	COINEX_SMART_CHAIN_TESTNET,
	CRONOS_MAINNET,
	CRONOS_TESTNET,
	CRONOS_TESTNET_ZKEVM_1,
	ETHEREUM_MAINNET,
	ETHEREUM_MAINNET_ARBITRUM_1,
	ETHEREUM_MAINNET_ARBITRUM_1_L3X_1,
	ETHEREUM_MAINNET_ARBITRUM_1_TREASURE_1,
	ETHEREUM_MAINNET_ASTAR_ZKEVM_1,
	ETHEREUM_MAINNET_BASE_1,
	ETHEREUM_MAINNET_BLAST_1,
	ETHEREUM_MAINNET_IMMUTABLE_ZKEVM_1,
	ETHEREUM_MAINNET_KROMA_1,
	ETHEREUM_MAINNET_LINEA_1,
	ETHEREUM_MAINNET_MANTLE_1,
	ETHEREUM_MAINNET_METIS_1,
	ETHEREUM_MAINNET_MODE_1,
	ETHEREUM_MAINNET_OPTIMISM_1,
	ETHEREUM_MAINNET_POLYGON_ZKEVM_1,
	ETHEREUM_MAINNET_SCROLL_1,
	ETHEREUM_MAINNET_TAIKO_1,
	ETHEREUM_MAINNET_WORLDCHAIN_1,
	ETHEREUM_MAINNET_XLAYER_1,
	ETHEREUM_MAINNET_ZKSYNC_1,
	ETHEREUM_TESTNET_GOERLI_ARBITRUM_1,
	ETHEREUM_TESTNET_GOERLI_BASE_1,
	ETHEREUM_TESTNET_GOERLI_LINEA_1,
	ETHEREUM_TESTNET_GOERLI_MANTLE_1,
	ETHEREUM_TESTNET_GOERLI_OPTIMISM_1,
	ETHEREUM_TESTNET_GOERLI_POLYGON_ZKEVM_1,
	ETHEREUM_TESTNET_GOERLI_ZKSYNC_1,
	ETHEREUM_TESTNET_HOLESKY,
	ETHEREUM_TESTNET_HOLESKY_FRAXTAL_1,
	ETHEREUM_TESTNET_HOLESKY_MORPH_1,
	ETHEREUM_TESTNET_HOLESKY_TAIKO_1,
	ETHEREUM_TESTNET_SEPOLIA,
	ETHEREUM_TESTNET_SEPOLIA_ARBITRUM_1,
	ETHEREUM_TESTNET_SEPOLIA_ARBITRUM_1_L3X_1,
	ETHEREUM_TESTNET_SEPOLIA_ARBITRUM_1_TREASURE_1,
	ETHEREUM_TESTNET_SEPOLIA_BASE_1,
	ETHEREUM_TESTNET_SEPOLIA_BLAST_1,
	ETHEREUM_TESTNET_SEPOLIA_CORN_1,
	ETHEREUM_TESTNET_SEPOLIA_IMMUTABLE_ZKEVM_1,
	ETHEREUM_TESTNET_SEPOLIA_KROMA_1,
	ETHEREUM_TESTNET_SEPOLIA_LENS_1,
	ETHEREUM_TESTNET_SEPOLIA_LINEA_1,
	ETHEREUM_TESTNET_SEPOLIA_LISK_1,
	ETHEREUM_TESTNET_SEPOLIA_MANTLE_1,
	ETHEREUM_TESTNET_SEPOLIA_METIS_1,
	ETHEREUM_TESTNET_SEPOLIA_MODE_1,
	ETHEREUM_TESTNET_SEPOLIA_OPTIMISM_1,
	ETHEREUM_TESTNET_SEPOLIA_POLYGON_VALIDIUM_1,
	ETHEREUM_TESTNET_SEPOLIA_POLYGON_ZKEVM_1,
	ETHEREUM_TESTNET_SEPOLIA_SCROLL_1,
	ETHEREUM_TESTNET_SEPOLIA_SONEIUM_1,
	ETHEREUM_TESTNET_SEPOLIA_UNICHAIN_1,
	ETHEREUM_TESTNET_SEPOLIA_WORLDCHAIN_1,
	ETHEREUM_TESTNET_SEPOLIA_XLAYER_1,
	ETHEREUM_TESTNET_SEPOLIA_ZIRCUIT_1,
	ETHEREUM_TESTNET_SEPOLIA_ZKSYNC_1,
	FANTOM_MAINNET,
	FANTOM_TESTNET,
	FILECOIN_MAINNET,
	FILECOIN_TESTNET,
	FRAXTAL_MAINNET,
	GETH_TESTNET,
	GNOSIS_CHAIN_MAINNET,
	GNOSIS_CHAIN_TESTNET_CHIADO,
	HEDERA_MAINNET,
	HEDERA_TESTNET,
	HYPERLIQUID_TESTNET,
	KAVA_MAINNET,
	KAVA_TESTNET,
	KUSAMA_MAINNET_MOONRIVER,
	NEAR_MAINNET,
	NEAR_TESTNET,
	NEONLINK_MAINNET,
	NEONLINK_TESTNET,
	PLUME_TESTNET,
	POLKADOT_MAINNET_ASTAR,
	POLKADOT_MAINNET_CENTRIFUGE,
	POLKADOT_MAINNET_DARWINIA,
	POLKADOT_MAINNET_MOONBEAM,
	POLKADOT_TESTNET_ASTAR_SHIBUYA,
	POLKADOT_TESTNET_CENTRIFUGE_ALTAIR,
	POLKADOT_TESTNET_DARWINIA_PANGORO,
	POLKADOT_TESTNET_MOONBEAM_MOONBASE,
	POLYGON_MAINNET,
	POLYGON_TESTNET_AMOY,
	POLYGON_TESTNET_MUMBAI,
	PRIVATE_TESTNET_MICA,
	PRIVATE_TESTNET_OPALA,
	RONIN_MAINNET,
	RONIN_TESTNET_SAIGON,
	SEI_MAINNET,
	SEI_TESTNET_ATLANTIC,
	SOLANA_DEVNET,
	SOLANA_MAINNET,
	SOLANA_TESTNET,
	SONIC_TESTNET,
	STORY_TESTNET,
	TELOS_EVM_MAINNET,
	TELOS_EVM_TESTNET,
	TEST_781901677223027175,
	VELAS_MAINNET,
	VELAS_TESTNET,
	WEMIX_MAINNET,
	WEMIX_TESTNET,
	ZIRCUIT_MAINNET,
	ZKLINK_NOVA_MAINNET,
	ZKLINK_NOVA_TESTNET,
}

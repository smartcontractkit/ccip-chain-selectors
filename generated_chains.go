// Code generated by go generate please DO NOT EDIT
package chain_selectors

type Chain struct {
	EvmChainID uint64
	Selector   uint64
	Name       string
	VarName    string
}

var (
	AREON_MAINNET                              = Chain{EvmChainID: 463, Selector: 1939936305787790600, Name: "areon-mainnet"}
	AREON_TESTNET                              = Chain{EvmChainID: 462, Selector: 7317911323415911000, Name: "areon-testnet"}
	AVALANCHE_MAINNET                          = Chain{EvmChainID: 43114, Selector: 6433500567565415381, Name: "avalanche-mainnet"}
	AVALANCHE_SUBNET_DEXALOT_MAINNET           = Chain{EvmChainID: 432204, Selector: 5463201557265485081, Name: "avalanche-subnet-dexalot-mainnet"}
	AVALANCHE_SUBNET_DEXALOT_TESTNET           = Chain{EvmChainID: 432201, Selector: 1458281248224512906, Name: "avalanche-subnet-dexalot-testnet"}
	AVALANCHE_TESTNET_FUJI                     = Chain{EvmChainID: 43113, Selector: 14767482510784806043, Name: "avalanche-testnet-fuji"}
	AVALANCHE_TESTNET_NEXON                    = Chain{EvmChainID: 595581, Selector: 7837562506228496256, Name: "avalanche-testnet-nexon"}
	BERACHAIN_TESTNET_ARTIO                    = Chain{EvmChainID: 80085, Selector: 12336603543561911511, Name: "berachain-testnet-artio"}
	BINANCE_SMART_CHAIN_MAINNET                = Chain{EvmChainID: 56, Selector: 11344663589394136015, Name: "binance_smart_chain-mainnet"}
	BINANCE_SMART_CHAIN_TESTNET                = Chain{EvmChainID: 97, Selector: 13264668187771770619, Name: "binance_smart_chain-testnet"}
	BITCICHAIN_MAINNET                         = Chain{EvmChainID: 1907, Selector: 4874388048629246000, Name: "bitcichain-mainnet"}
	BITCICHAIN_TESTNET                         = Chain{EvmChainID: 1908, Selector: 4888058894222120000, Name: "bitcichain-testnet"}
	BITCOIN_MAINNET_BOTANIX                    = Chain{EvmChainID: 3637, Selector: 4560701533377838164, Name: "bitcoin-mainnet-botanix"}
	BITCOIN_MERLIN_MAINNET                     = Chain{EvmChainID: 4200, Selector: 241851231317828981, Name: "bitcoin-merlin-mainnet"}
	BITCOIN_TESTNET_BOTANIX                    = Chain{EvmChainID: 3636, Selector: 1467223411771711614, Name: "bitcoin-testnet-botanix"}
	BITCOIN_TESTNET_MERLIN                     = Chain{EvmChainID: 686868, Selector: 5269261765892944301, Name: "bitcoin-testnet-merlin"}
	BITCOIN_TESTNET_ROOTSTOCK                  = Chain{EvmChainID: 31, Selector: 8953668971247136127, Name: "bitcoin-testnet-rootstock"}
	BITTORRENT_CHAIN_MAINNET                   = Chain{EvmChainID: 199, Selector: 3776006016387883143, Name: "bittorrent_chain-mainnet"}
	BITTORRENT_CHAIN_TESTNET                   = Chain{EvmChainID: 1029, Selector: 4459371029167934217, Name: "bittorrent_chain-testnet"}
	CELO_MAINNET                               = Chain{EvmChainID: 42220, Selector: 1346049177634351622, Name: "celo-mainnet"}
	CELO_TESTNET_ALFAJORES                     = Chain{EvmChainID: 44787, Selector: 3552045678561919002, Name: "celo-testnet-alfajores"}
	COINEX_SMART_CHAIN_MAINNET                 = Chain{EvmChainID: 52, Selector: 1761333065194157300, Name: "coinex_smart_chain-mainnet"}
	COINEX_SMART_CHAIN_TESTNET                 = Chain{EvmChainID: 53, Selector: 8955032871639343000, Name: "coinex_smart_chain-testnet"}
	CRONOS_MAINNET                             = Chain{EvmChainID: 25, Selector: 1456215246176062136, Name: "cronos-mainnet"}
	CRONOS_TESTNET                             = Chain{EvmChainID: 338, Selector: 2995292832068775165, Name: "cronos-testnet"}
	CRONOS_TESTNET_ZKEVM_1                     = Chain{EvmChainID: 282, Selector: 3842103497652714138, Name: "cronos-testnet-zkevm-1"}
	ETHEREUM_MAINNET                           = Chain{EvmChainID: 1, Selector: 5009297550715157269, Name: "ethereum-mainnet"}
	ETHEREUM_MAINNET_ARBITRUM_1                = Chain{EvmChainID: 42161, Selector: 4949039107694359620, Name: "ethereum-mainnet-arbitrum-1"}
	ETHEREUM_MAINNET_ASTAR_ZKEVM_1             = Chain{EvmChainID: 3776, Selector: 1540201334317828111, Name: "ethereum-mainnet-astar-zkevm-1"}
	ETHEREUM_MAINNET_BASE_1                    = Chain{EvmChainID: 8453, Selector: 15971525489660198786, Name: "ethereum-mainnet-base-1"}
	ETHEREUM_MAINNET_BLAST_1                   = Chain{EvmChainID: 81457, Selector: 4411394078118774322, Name: "ethereum-mainnet-blast-1"}
	ETHEREUM_MAINNET_IMMUTABLE_ZKEVM_1         = Chain{EvmChainID: 13371, Selector: 1237925231416731909, Name: "ethereum-mainnet-immutable-zkevm-1"}
	ETHEREUM_MAINNET_KROMA_1                   = Chain{EvmChainID: 255, Selector: 3719320017875267166, Name: "ethereum-mainnet-kroma-1"}
	ETHEREUM_MAINNET_LINEA_1                   = Chain{EvmChainID: 59144, Selector: 4627098889531055414, Name: "ethereum-mainnet-linea-1"}
	ETHEREUM_MAINNET_MANTLE_1                  = Chain{EvmChainID: 5000, Selector: 1556008542357238666, Name: "ethereum-mainnet-mantle-1"}
	ETHEREUM_MAINNET_METIS_1                   = Chain{EvmChainID: 1088, Selector: 8805746078405598895, Name: "ethereum-mainnet-metis-1"}
	ETHEREUM_MAINNET_MODE_1                    = Chain{EvmChainID: 34443, Selector: 7264351850409363825, Name: "ethereum-mainnet-mode-1"}
	ETHEREUM_MAINNET_OPTIMISM_1                = Chain{EvmChainID: 10, Selector: 3734403246176062136, Name: "ethereum-mainnet-optimism-1"}
	ETHEREUM_MAINNET_POLYGON_ZKEVM_1           = Chain{EvmChainID: 1101, Selector: 4348158687435793198, Name: "ethereum-mainnet-polygon-zkevm-1"}
	ETHEREUM_MAINNET_SCROLL_1                  = Chain{EvmChainID: 534352, Selector: 13204309965629103672, Name: "ethereum-mainnet-scroll-1"}
	ETHEREUM_MAINNET_XLAYER_1                  = Chain{EvmChainID: 196, Selector: 3016212468291539606, Name: "ethereum-mainnet-xlayer-1"}
	ETHEREUM_MAINNET_ZKSYNC_1                  = Chain{EvmChainID: 324, Selector: 1562403441176082196, Name: "ethereum-mainnet-zksync-1"}
	ETHEREUM_TESTNET_GOERLI_ARBITRUM_1         = Chain{EvmChainID: 421613, Selector: 6101244977088475029, Name: "ethereum-testnet-goerli-arbitrum-1"}
	ETHEREUM_TESTNET_GOERLI_BASE_1             = Chain{EvmChainID: 84531, Selector: 5790810961207155433, Name: "ethereum-testnet-goerli-base-1"}
	ETHEREUM_TESTNET_GOERLI_LINEA_1            = Chain{EvmChainID: 59140, Selector: 1355246678561316402, Name: "ethereum-testnet-goerli-linea-1"}
	ETHEREUM_TESTNET_GOERLI_MANTLE_1           = Chain{EvmChainID: 5001, Selector: 4168263376276232250, Name: "ethereum-testnet-goerli-mantle-1"}
	ETHEREUM_TESTNET_GOERLI_OPTIMISM_1         = Chain{EvmChainID: 420, Selector: 2664363617261496610, Name: "ethereum-testnet-goerli-optimism-1"}
	ETHEREUM_TESTNET_GOERLI_POLYGON_ZKEVM_1    = Chain{EvmChainID: 1442, Selector: 11059667695644972511, Name: "ethereum-testnet-goerli-polygon-zkevm-1"}
	ETHEREUM_TESTNET_GOERLI_ZKSYNC_1           = Chain{EvmChainID: 280, Selector: 6802309497652714138, Name: "ethereum-testnet-goerli-zksync-1"}
	ETHEREUM_TESTNET_HOLESKY_FRAXTAL_1         = Chain{EvmChainID: 2522, Selector: 8901520481741771655, Name: "ethereum-testnet-holesky-fraxtal-1"}
	ETHEREUM_TESTNET_HOLESKY_MORPH_1           = Chain{EvmChainID: 2810, Selector: 8304510386741731151, Name: "ethereum-testnet-holesky-morph-1"}
	ETHEREUM_TESTNET_SEPOLIA                   = Chain{EvmChainID: 11155111, Selector: 16015286601757825753, Name: "ethereum-testnet-sepolia"}
	ETHEREUM_TESTNET_SEPOLIA_ARBITRUM_1        = Chain{EvmChainID: 421614, Selector: 3478487238524512106, Name: "ethereum-testnet-sepolia-arbitrum-1"}
	ETHEREUM_TESTNET_SEPOLIA_BASE_1            = Chain{EvmChainID: 84532, Selector: 10344971235874465080, Name: "ethereum-testnet-sepolia-base-1"}
	ETHEREUM_TESTNET_SEPOLIA_BLAST_1           = Chain{EvmChainID: 168587773, Selector: 2027362563942762617, Name: "ethereum-testnet-sepolia-blast-1"}
	ETHEREUM_TESTNET_SEPOLIA_CORN_1            = Chain{EvmChainID: 21000000, Selector: 1467427327723633929, Name: "ethereum-testnet-sepolia-corn-1"}
	ETHEREUM_TESTNET_SEPOLIA_IMMUTABLE_ZKEVM_1 = Chain{EvmChainID: 13473, Selector: 4526165231216331901, Name: "ethereum-testnet-sepolia-immutable-zkevm-1"}
	ETHEREUM_TESTNET_SEPOLIA_KROMA_1           = Chain{EvmChainID: 2358, Selector: 5990477251245693094, Name: "ethereum-testnet-sepolia-kroma-1"}
	ETHEREUM_TESTNET_SEPOLIA_LINEA_1           = Chain{EvmChainID: 59141, Selector: 5719461335882077547, Name: "ethereum-testnet-sepolia-linea-1"}
	ETHEREUM_TESTNET_SEPOLIA_LISK_1            = Chain{EvmChainID: 4202, Selector: 5298399861320400553, Name: "ethereum-testnet-sepolia-lisk-1"}
	ETHEREUM_TESTNET_SEPOLIA_MANTLE_1          = Chain{EvmChainID: 5003, Selector: 8236463271206331221, Name: "ethereum-testnet-sepolia-mantle-1"}
	ETHEREUM_TESTNET_SEPOLIA_METIS_1           = Chain{EvmChainID: 59902, Selector: 3777822886988675105, Name: "ethereum-testnet-sepolia-metis-1"}
	ETHEREUM_TESTNET_SEPOLIA_MODE_1            = Chain{EvmChainID: 919, Selector: 829525985033418733, Name: "ethereum-testnet-sepolia-mode-1"}
	ETHEREUM_TESTNET_SEPOLIA_OPTIMISM_1        = Chain{EvmChainID: 11155420, Selector: 5224473277236331295, Name: "ethereum-testnet-sepolia-optimism-1"}
	ETHEREUM_TESTNET_SEPOLIA_POLYGON_ZKEVM_1   = Chain{EvmChainID: 2442, Selector: 1654667687261492630, Name: "ethereum-testnet-sepolia-polygon-zkevm-1"}
	ETHEREUM_TESTNET_SEPOLIA_SCROLL_1          = Chain{EvmChainID: 534351, Selector: 2279865765895943307, Name: "ethereum-testnet-sepolia-scroll-1"}
	ETHEREUM_TESTNET_SEPOLIA_XLAYER_1          = Chain{EvmChainID: 195, Selector: 2066098519157881736, Name: "ethereum-testnet-sepolia-xlayer-1"}
	ETHEREUM_TESTNET_SEPOLIA_ZIRCUIT           = Chain{EvmChainID: 48899, Selector: 4562743618362911021, Name: "ethereum-testnet-sepolia-zircuit"}
	ETHEREUM_TESTNET_SEPOLIA_ZKSYNC_1          = Chain{EvmChainID: 300, Selector: 6898391096552792247, Name: "ethereum-testnet-sepolia-zksync-1"}
	FANTOM_TESTNET                             = Chain{EvmChainID: 4002, Selector: 4905564228793744293, Name: "fantom-testnet"}
	FANTOM_TESTNET_OPERA                       = Chain{EvmChainID: 250, Selector: 3768048213127883732, Name: "fantom-testnet-opera"}
	FRAXTAL_MAINNET                            = Chain{EvmChainID: 252, Selector: 1462016016387883143, Name: "fraxtal-mainnet"}
	GETH_DEVNET_2                              = Chain{EvmChainID: 2337, Selector: 12922642891491394802, Name: "geth-devnet-2"}
	GETH_DEVNET_3                              = Chain{EvmChainID: 3337, Selector: 4793464827907405086, Name: "geth-devnet-3"}
	GETH_TESTNET                               = Chain{EvmChainID: 1337, Selector: 3379446385462418246, Name: "geth-testnet"}
	GNOSIS_CHAIN_MAINNET                       = Chain{EvmChainID: 100, Selector: 465200170687744372, Name: "gnosis_chain-mainnet"}
	GNOSIS_CHAIN_TESTNET_CHIADO                = Chain{EvmChainID: 10200, Selector: 8871595565390010547, Name: "gnosis_chain-testnet-chiado"}
	HEDERA_TESTNET                             = Chain{EvmChainID: 296, Selector: 222782988166878823, Name: "hedera-testnet"}
	KAVA_MAINNET                               = Chain{EvmChainID: 2222, Selector: 7550000543357438061, Name: "kava-mainnet"}
	KAVA_TESTNET                               = Chain{EvmChainID: 2221, Selector: 2110537777356199208, Name: "kava-testnet"}
	KUSAMA_MAINNET_MOONRIVER                   = Chain{EvmChainID: 1285, Selector: 1355020143337428062, Name: "kusama-mainnet-moonriver"}
	NEAR_MAINNET                               = Chain{EvmChainID: 397, Selector: 2039744413822257700, Name: "near-mainnet"}
	NEAR_TESTNET                               = Chain{EvmChainID: 398, Selector: 5061593697262339000, Name: "near-testnet"}
	NEONLINK_MAINNET                           = Chain{EvmChainID: 259, Selector: 8239338020728974000, Name: "neonlink-mainnet"}
	NEONLINK_TESTNET                           = Chain{EvmChainID: 9559, Selector: 1113014352258747600, Name: "neonlink-testnet"}
	POLKADOT_MAINNET_ASTAR                     = Chain{EvmChainID: 592, Selector: 6422105447186081193, Name: "polkadot-mainnet-astar"}
	POLKADOT_MAINNET_DARWINIA                  = Chain{EvmChainID: 46, Selector: 8866418665544333000, Name: "polkadot-mainnet-darwinia"}
	POLKADOT_MAINNET_MOONBEAM                  = Chain{EvmChainID: 1284, Selector: 1252863800116739621, Name: "polkadot-mainnet-moonbeam"}
	POLKADOT_TESTNET_ASTAR_SHIBUYA             = Chain{EvmChainID: 81, Selector: 6955638871347136141, Name: "polkadot-testnet-astar-shibuya"}
	POLKADOT_TESTNET_DARWINIA_PANGORO          = Chain{EvmChainID: 45, Selector: 4340886533089894000, Name: "polkadot-testnet-darwinia-pangoro"}
	POLKADOT_TESTNET_MOONBEAM_MOONBASE         = Chain{EvmChainID: 1287, Selector: 5361632739113536121, Name: "polkadot-testnet-moonbeam-moonbase"}
	POLYGON_MAINNET                            = Chain{EvmChainID: 137, Selector: 4051577828743386545, Name: "polygon-mainnet"}
	POLYGON_TESTNET_AMOY                       = Chain{EvmChainID: 80002, Selector: 16281711391670634445, Name: "polygon-testnet-amoy"}
	POLYGON_TESTNET_MUMBAI                     = Chain{EvmChainID: 80001, Selector: 12532609583862916517, Name: "polygon-testnet-mumbai"}
	TELOS_EVM_MAINNET                          = Chain{EvmChainID: 40, Selector: 1477345371608778000, Name: "telos-evm-mainnet"}
	TELOS_EVM_TESTNET                          = Chain{EvmChainID: 41, Selector: 729797994450396300, Name: "telos-evm-testnet"}
	TEST_1000                                  = Chain{EvmChainID: 1000, Selector: 11787463284727550157, Name: "1000"}
	TEST_76578                                 = Chain{EvmChainID: 76578, Selector: 781901677223027175, Name: "76578"}
	TEST_90000001                              = Chain{EvmChainID: 90000001, Selector: 909606746561742123, Name: "90000001"}
	TEST_90000002                              = Chain{EvmChainID: 90000002, Selector: 5548718428018410741, Name: "90000002"}
	TEST_90000003                              = Chain{EvmChainID: 90000003, Selector: 789068866484373046, Name: "90000003"}
	TEST_90000004                              = Chain{EvmChainID: 90000004, Selector: 5721565186521185178, Name: "90000004"}
	TEST_90000005                              = Chain{EvmChainID: 90000005, Selector: 964127714438319834, Name: "90000005"}
	TEST_90000006                              = Chain{EvmChainID: 90000006, Selector: 8966794841936584464, Name: "90000006"}
	TEST_90000007                              = Chain{EvmChainID: 90000007, Selector: 8412806778050735057, Name: "90000007"}
	TEST_90000008                              = Chain{EvmChainID: 90000008, Selector: 4066443121807923198, Name: "90000008"}
	TEST_90000009                              = Chain{EvmChainID: 90000009, Selector: 6747736380229414777, Name: "90000009"}
	TEST_90000010                              = Chain{EvmChainID: 90000010, Selector: 8694984074292254623, Name: "90000010"}
	TEST_90000011                              = Chain{EvmChainID: 90000011, Selector: 328334718812072308, Name: "90000011"}
	TEST_90000012                              = Chain{EvmChainID: 90000012, Selector: 7715160997071429212, Name: "90000012"}
	TEST_90000013                              = Chain{EvmChainID: 90000013, Selector: 3574539439524578558, Name: "90000013"}
	TEST_90000014                              = Chain{EvmChainID: 90000014, Selector: 4543928599863227519, Name: "90000014"}
	TEST_90000015                              = Chain{EvmChainID: 90000015, Selector: 6443235356619661032, Name: "90000015"}
	TEST_90000016                              = Chain{EvmChainID: 90000016, Selector: 13087962012083037329, Name: "90000016"}
	TEST_90000017                              = Chain{EvmChainID: 90000017, Selector: 11985232338641871056, Name: "90000017"}
	TEST_90000018                              = Chain{EvmChainID: 90000018, Selector: 7777066535355430289, Name: "90000018"}
	TEST_90000019                              = Chain{EvmChainID: 90000019, Selector: 1273605685587320666, Name: "90000019"}
	TEST_90000020                              = Chain{EvmChainID: 90000020, Selector: 17810359353458878177, Name: "90000020"}
	TEST_90000021                              = Chain{EvmChainID: 90000021, Selector: 13648736134397881410, Name: "90000021"}
	TEST_90000022                              = Chain{EvmChainID: 90000022, Selector: 6742472197519042017, Name: "90000022"}
	TEST_90000023                              = Chain{EvmChainID: 90000023, Selector: 16702426279731183946, Name: "90000023"}
	TEST_90000024                              = Chain{EvmChainID: 90000024, Selector: 16449698933146693970, Name: "90000024"}
	TEST_90000025                              = Chain{EvmChainID: 90000025, Selector: 5614341928911841614, Name: "90000025"}
	TEST_90000026                              = Chain{EvmChainID: 90000026, Selector: 9932483170498916221, Name: "90000026"}
	TEST_90000027                              = Chain{EvmChainID: 90000027, Selector: 9248511054298050610, Name: "90000027"}
	TEST_90000028                              = Chain{EvmChainID: 90000028, Selector: 15733873364998401606, Name: "90000028"}
	TEST_90000029                              = Chain{EvmChainID: 90000029, Selector: 10199579733509604193, Name: "90000029"}
	TEST_90000030                              = Chain{EvmChainID: 90000030, Selector: 11754399446572002459, Name: "90000030"}
	TEST_90000031                              = Chain{EvmChainID: 90000031, Selector: 15804983202763665802, Name: "90000031"}
	TEST_90000032                              = Chain{EvmChainID: 90000032, Selector: 8794884152664322911, Name: "90000032"}
	TEST_90000033                              = Chain{EvmChainID: 90000033, Selector: 7005880874640146484, Name: "90000033"}
	TEST_90000034                              = Chain{EvmChainID: 90000034, Selector: 15998314635132476942, Name: "90000034"}
	TEST_90000035                              = Chain{EvmChainID: 90000035, Selector: 6676710761873615962, Name: "90000035"}
	TEST_90000036                              = Chain{EvmChainID: 90000036, Selector: 13973515790491921010, Name: "90000036"}
	TEST_90000037                              = Chain{EvmChainID: 90000037, Selector: 12226902941055802385, Name: "90000037"}
	TEST_90000038                              = Chain{EvmChainID: 90000038, Selector: 10547673735879567911, Name: "90000038"}
	TEST_90000039                              = Chain{EvmChainID: 90000039, Selector: 2953028829530698683, Name: "90000039"}
	TEST_90000040                              = Chain{EvmChainID: 90000040, Selector: 3740583887329090549, Name: "90000040"}
	TEST_90000041                              = Chain{EvmChainID: 90000041, Selector: 4716670523656754658, Name: "90000041"}
	TEST_90000042                              = Chain{EvmChainID: 90000042, Selector: 12965905455277595820, Name: "90000042"}
	TEST_90000043                              = Chain{EvmChainID: 90000043, Selector: 6448403805635971860, Name: "90000043"}
	TEST_90000044                              = Chain{EvmChainID: 90000044, Selector: 176199025415897437, Name: "90000044"}
	TEST_90000045                              = Chain{EvmChainID: 90000045, Selector: 17251043223284625647, Name: "90000045"}
	TEST_90000046                              = Chain{EvmChainID: 90000046, Selector: 14943531413383612703, Name: "90000046"}
	TEST_90000047                              = Chain{EvmChainID: 90000047, Selector: 8015762103567576333, Name: "90000047"}
	TEST_90000048                              = Chain{EvmChainID: 90000048, Selector: 2783890746839497525, Name: "90000048"}
	TEST_90000049                              = Chain{EvmChainID: 90000049, Selector: 16591966440843528322, Name: "90000049"}
	TEST_90000050                              = Chain{EvmChainID: 90000050, Selector: 9156614022853705708, Name: "90000050"}
	TEST_90000051                              = Chain{EvmChainID: 90000051, Selector: 10089241509396411113, Name: "90000051"}
	TEST_90000052                              = Chain{EvmChainID: 90000052, Selector: 7585715102059681757, Name: "90000052"}
	TEST_90000053                              = Chain{EvmChainID: 90000053, Selector: 9574369650680012313, Name: "90000053"}
	TEST_90000054                              = Chain{EvmChainID: 90000054, Selector: 15767478222558315144, Name: "90000054"}
	TEST_90000055                              = Chain{EvmChainID: 90000055, Selector: 928756709184343973, Name: "90000055"}
	TEST_90000056                              = Chain{EvmChainID: 90000056, Selector: 13936493323944617843, Name: "90000056"}
	TEST_90000057                              = Chain{EvmChainID: 90000057, Selector: 9264503539336248559, Name: "90000057"}
	TEST_90000058                              = Chain{EvmChainID: 90000058, Selector: 7032045258883126022, Name: "90000058"}
	TEST_90000059                              = Chain{EvmChainID: 90000059, Selector: 13781595843667691007, Name: "90000059"}
	TEST_90000060                              = Chain{EvmChainID: 90000060, Selector: 6751512843227450641, Name: "90000060"}
	TEST_90000061                              = Chain{EvmChainID: 90000061, Selector: 12027427861168955422, Name: "90000061"}
	TEST_90000062                              = Chain{EvmChainID: 90000062, Selector: 6690738652320128159, Name: "90000062"}
	TEST_90000063                              = Chain{EvmChainID: 90000063, Selector: 12513826466599144030, Name: "90000063"}
	TEST_90000064                              = Chain{EvmChainID: 90000064, Selector: 7823363553221722351, Name: "90000064"}
	TEST_90000065                              = Chain{EvmChainID: 90000065, Selector: 17759418850483131633, Name: "90000065"}
	TEST_90000066                              = Chain{EvmChainID: 90000066, Selector: 1488785539820432596, Name: "90000066"}
	TEST_90000067                              = Chain{EvmChainID: 90000067, Selector: 12470167056735102403, Name: "90000067"}
	TEST_90000068                              = Chain{EvmChainID: 90000068, Selector: 6059917085984771915, Name: "90000068"}
	TEST_90000069                              = Chain{EvmChainID: 90000069, Selector: 8698844633699288298, Name: "90000069"}
	TEST_90000070                              = Chain{EvmChainID: 90000070, Selector: 11335955773964346155, Name: "90000070"}
	TEST_90000071                              = Chain{EvmChainID: 90000071, Selector: 15210860601736105873, Name: "90000071"}
	TEST_90000072                              = Chain{EvmChainID: 90000072, Selector: 15447447865219782832, Name: "90000072"}
	TEST_90000073                              = Chain{EvmChainID: 90000073, Selector: 7404045285477377670, Name: "90000073"}
	TEST_90000074                              = Chain{EvmChainID: 90000074, Selector: 14506622911400094011, Name: "90000074"}
	TEST_90000075                              = Chain{EvmChainID: 90000075, Selector: 18316006852148771137, Name: "90000075"}
	TEST_90000076                              = Chain{EvmChainID: 90000076, Selector: 7961714422080771198, Name: "90000076"}
	TEST_90000077                              = Chain{EvmChainID: 90000077, Selector: 15168140751097121912, Name: "90000077"}
	TEST_90000078                              = Chain{EvmChainID: 90000078, Selector: 8354317460459584308, Name: "90000078"}
	TEST_90000079                              = Chain{EvmChainID: 90000079, Selector: 1974710175227680991, Name: "90000079"}
	TEST_90000080                              = Chain{EvmChainID: 90000080, Selector: 15896959195233368219, Name: "90000080"}
	TEST_90000081                              = Chain{EvmChainID: 90000081, Selector: 13819071330241498802, Name: "90000081"}
	TEST_90000082                              = Chain{EvmChainID: 90000082, Selector: 3632230855428784129, Name: "90000082"}
	TEST_90000083                              = Chain{EvmChainID: 90000083, Selector: 3330151784927722907, Name: "90000083"}
	TEST_90000084                              = Chain{EvmChainID: 90000084, Selector: 973671184102733124, Name: "90000084"}
	TEST_90000085                              = Chain{EvmChainID: 90000085, Selector: 7353384334508842175, Name: "90000085"}
	TEST_90000086                              = Chain{EvmChainID: 90000086, Selector: 4174149892778961910, Name: "90000086"}
	TEST_90000087                              = Chain{EvmChainID: 90000087, Selector: 10497629267361915835, Name: "90000087"}
	TEST_90000088                              = Chain{EvmChainID: 90000088, Selector: 10537986502862404866, Name: "90000088"}
	TEST_90000089                              = Chain{EvmChainID: 90000089, Selector: 10106333385848939617, Name: "90000089"}
	TEST_90000090                              = Chain{EvmChainID: 90000090, Selector: 2509173735760116798, Name: "90000090"}
	TEST_90000091                              = Chain{EvmChainID: 90000091, Selector: 12499149790922928210, Name: "90000091"}
	TEST_90000092                              = Chain{EvmChainID: 90000092, Selector: 665284410079532457, Name: "90000092"}
	TEST_90000093                              = Chain{EvmChainID: 90000093, Selector: 17514102371649734225, Name: "90000093"}
	TEST_90000094                              = Chain{EvmChainID: 90000094, Selector: 8211981504472319767, Name: "90000094"}
	TEST_90000095                              = Chain{EvmChainID: 90000095, Selector: 15945074456050759193, Name: "90000095"}
	TEST_90000096                              = Chain{EvmChainID: 90000096, Selector: 17580537314894454709, Name: "90000096"}
	TEST_90000097                              = Chain{EvmChainID: 90000097, Selector: 13443138560923813712, Name: "90000097"}
	TEST_90000098                              = Chain{EvmChainID: 90000098, Selector: 9675086780529785020, Name: "90000098"}
	TEST_90000099                              = Chain{EvmChainID: 90000099, Selector: 7431973150957944526, Name: "90000099"}
	TEST_90000100                              = Chain{EvmChainID: 90000100, Selector: 6875898693582952601, Name: "90000100"}
	VELAS_MAINNET                              = Chain{EvmChainID: 106, Selector: 374210358663784372, Name: "velas-mainnet"}
	VELAS_TESTNET                              = Chain{EvmChainID: 111, Selector: 572210378683744374, Name: "velas-testnet"}
	WEMIX_MAINNET                              = Chain{EvmChainID: 1111, Selector: 5142893604156789321, Name: "wemix-mainnet"}
	WEMIX_TESTNET                              = Chain{EvmChainID: 1112, Selector: 9284632837123596123, Name: "wemix-testnet"}
	ZKLINK_NOVA_MAINNET                        = Chain{EvmChainID: 810180, Selector: 4350319965322101699, Name: "zklink_nova-mainnet"}
	ZKLINK_NOVA_TESTNET                        = Chain{EvmChainID: 810181, Selector: 5837261596322416298, Name: "zklink_nova-testnet"}
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
	BINANCE_SMART_CHAIN_MAINNET,
	BINANCE_SMART_CHAIN_TESTNET,
	BITCICHAIN_MAINNET,
	BITCICHAIN_TESTNET,
	BITCOIN_MAINNET_BOTANIX,
	BITCOIN_MERLIN_MAINNET,
	BITCOIN_TESTNET_BOTANIX,
	BITCOIN_TESTNET_MERLIN,
	BITCOIN_TESTNET_ROOTSTOCK,
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
	ETHEREUM_MAINNET_XLAYER_1,
	ETHEREUM_MAINNET_ZKSYNC_1,
	ETHEREUM_TESTNET_GOERLI_ARBITRUM_1,
	ETHEREUM_TESTNET_GOERLI_BASE_1,
	ETHEREUM_TESTNET_GOERLI_LINEA_1,
	ETHEREUM_TESTNET_GOERLI_MANTLE_1,
	ETHEREUM_TESTNET_GOERLI_OPTIMISM_1,
	ETHEREUM_TESTNET_GOERLI_POLYGON_ZKEVM_1,
	ETHEREUM_TESTNET_GOERLI_ZKSYNC_1,
	ETHEREUM_TESTNET_HOLESKY_FRAXTAL_1,
	ETHEREUM_TESTNET_HOLESKY_MORPH_1,
	ETHEREUM_TESTNET_SEPOLIA,
	ETHEREUM_TESTNET_SEPOLIA_ARBITRUM_1,
	ETHEREUM_TESTNET_SEPOLIA_BASE_1,
	ETHEREUM_TESTNET_SEPOLIA_BLAST_1,
	ETHEREUM_TESTNET_SEPOLIA_CORN_1,
	ETHEREUM_TESTNET_SEPOLIA_IMMUTABLE_ZKEVM_1,
	ETHEREUM_TESTNET_SEPOLIA_KROMA_1,
	ETHEREUM_TESTNET_SEPOLIA_LINEA_1,
	ETHEREUM_TESTNET_SEPOLIA_LISK_1,
	ETHEREUM_TESTNET_SEPOLIA_MANTLE_1,
	ETHEREUM_TESTNET_SEPOLIA_METIS_1,
	ETHEREUM_TESTNET_SEPOLIA_MODE_1,
	ETHEREUM_TESTNET_SEPOLIA_OPTIMISM_1,
	ETHEREUM_TESTNET_SEPOLIA_POLYGON_ZKEVM_1,
	ETHEREUM_TESTNET_SEPOLIA_SCROLL_1,
	ETHEREUM_TESTNET_SEPOLIA_XLAYER_1,
	ETHEREUM_TESTNET_SEPOLIA_ZIRCUIT,
	ETHEREUM_TESTNET_SEPOLIA_ZKSYNC_1,
	FANTOM_TESTNET,
	FANTOM_TESTNET_OPERA,
	FRAXTAL_MAINNET,
	GETH_DEVNET_2,
	GETH_DEVNET_3,
	GETH_TESTNET,
	GNOSIS_CHAIN_MAINNET,
	GNOSIS_CHAIN_TESTNET_CHIADO,
	HEDERA_TESTNET,
	KAVA_MAINNET,
	KAVA_TESTNET,
	KUSAMA_MAINNET_MOONRIVER,
	NEAR_MAINNET,
	NEAR_TESTNET,
	NEONLINK_MAINNET,
	NEONLINK_TESTNET,
	POLKADOT_MAINNET_ASTAR,
	POLKADOT_MAINNET_DARWINIA,
	POLKADOT_MAINNET_MOONBEAM,
	POLKADOT_TESTNET_ASTAR_SHIBUYA,
	POLKADOT_TESTNET_DARWINIA_PANGORO,
	POLKADOT_TESTNET_MOONBEAM_MOONBASE,
	POLYGON_MAINNET,
	POLYGON_TESTNET_AMOY,
	POLYGON_TESTNET_MUMBAI,
	TELOS_EVM_MAINNET,
	TELOS_EVM_TESTNET,
	TEST_1000,
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
	TEST_90000021,
	TEST_90000022,
	TEST_90000023,
	TEST_90000024,
	TEST_90000025,
	TEST_90000026,
	TEST_90000027,
	TEST_90000028,
	TEST_90000029,
	TEST_90000030,
	TEST_90000031,
	TEST_90000032,
	TEST_90000033,
	TEST_90000034,
	TEST_90000035,
	TEST_90000036,
	TEST_90000037,
	TEST_90000038,
	TEST_90000039,
	TEST_90000040,
	TEST_90000041,
	TEST_90000042,
	TEST_90000043,
	TEST_90000044,
	TEST_90000045,
	TEST_90000046,
	TEST_90000047,
	TEST_90000048,
	TEST_90000049,
	TEST_90000050,
	TEST_90000051,
	TEST_90000052,
	TEST_90000053,
	TEST_90000054,
	TEST_90000055,
	TEST_90000056,
	TEST_90000057,
	TEST_90000058,
	TEST_90000059,
	TEST_90000060,
	TEST_90000061,
	TEST_90000062,
	TEST_90000063,
	TEST_90000064,
	TEST_90000065,
	TEST_90000066,
	TEST_90000067,
	TEST_90000068,
	TEST_90000069,
	TEST_90000070,
	TEST_90000071,
	TEST_90000072,
	TEST_90000073,
	TEST_90000074,
	TEST_90000075,
	TEST_90000076,
	TEST_90000077,
	TEST_90000078,
	TEST_90000079,
	TEST_90000080,
	TEST_90000081,
	TEST_90000082,
	TEST_90000083,
	TEST_90000084,
	TEST_90000085,
	TEST_90000086,
	TEST_90000087,
	TEST_90000088,
	TEST_90000089,
	TEST_90000090,
	TEST_90000091,
	TEST_90000092,
	TEST_90000093,
	TEST_90000094,
	TEST_90000095,
	TEST_90000096,
	TEST_90000097,
	TEST_90000098,
	TEST_90000099,
	TEST_90000100,
	VELAS_MAINNET,
	VELAS_TESTNET,
	WEMIX_MAINNET,
	WEMIX_TESTNET,
	ZKLINK_NOVA_MAINNET,
	ZKLINK_NOVA_TESTNET,
}

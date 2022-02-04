# EthLog
Simple ethereum logs parsing/dumping tool that combines input/output/events/receipt of all txs for range of blocks
, also decoding values according to a list of contracts ABI provided

### Development
```
make install-deps
make lint
make build - builds test contracts
make test - run simulated backend tests
make clean - clean generated abi/bindings and dump files
```
After running tests you can watch example report from tests with simulated backend in `ethlog/`

### Example usage as a library
```go
package main

import (
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/skudasov/ethlog/erc20"
	"github.com/skudasov/ethlog/ethlog"
)

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout}).Level(zerolog.DebugLevel)
}

func main() {
	client, err := ethclient.Dial("${ETH_URL}")
	if err != nil {
		log.Fatal().Err(err).Send()
	}
	e := ethlog.NewEthLog(client, 100)
	_, err = e.History(&ethlog.HistoryConfig{
		FromBlock:  big.NewInt(14142000),
		ToBlock:    big.NewInt(14142010),
		Format:     ethlog.FormatJSON,
		OutputFile: "mainnet_ERC20_blocks",
		ContractsData: []ethlog.ContractData{
			{
				Name:    "ERC20",
				ABI:     erc20.Erc20ABI,
				Address: common.HexToAddress("0x..."),
			},
		},
	})
	if err != nil {
		log.Fatal().Err(err).Send()
	}
}
```

#### TODO
- [x] Basic tests
- [x] Parallel processing 
- [x] ERC20 on different networks (mainnet/testnet/evm-compatible) (partially)
- [ ] More tests and live validation on different contracts


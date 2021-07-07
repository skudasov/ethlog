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
	"ethlog/ethlog"
	"ethlog/events_test_contract"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

func main() {
	client, err := ethclient.Dial("ws://localhost:8545")
	if err != nil {
		log.Fatal(err)
	}
	e := ethlog.NewEthLog(client)
	cfg := &ethlog.BlockHistoryConfig{
		FromBlock: big.NewInt(100),
		ToBlock:   big.NewInt(200),
		Format:    ethlog.FormatJSON,
		Rewrite:   true,
		ContractsData: []ethlog.ContractData{
			{
				Name:    "eventTestContract",
				ABI:     events_test_contract.EventsTestContractABI,
				Address: common.HexToAddress("0x0"), // contract address
			},
		},
	}
	// dumps all the logs from 100 to 200 block with decoded data
	historyMap, _ := e.RequestBlocksHistory(cfg)
	_ = e.DumpBlockHistory("all_blocks", cfg, historyMap)
	// dumps events of every contract in a file
	eventsMap, _ := e.RequestEventsHistory(cfg)
	err = e.DumpEventsByFile(cfg, eventsMap)
}
```

#### TODO
- [x] Basic tests
- [ ] More tests
- [ ] CLI
- [ ] Fields filtering config


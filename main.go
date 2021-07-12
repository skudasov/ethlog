package main

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rs/zerolog"
	"github.com/skudasov/ethlog/ethlog"
	"github.com/skudasov/ethlog/events_test_contract"
	"log"
)

func main() {
	client, err := ethclient.Dial("ws://localhost:8545")
	if err != nil {
		log.Fatal(err)
	}
	e := ethlog.NewEthLog(client, zerolog.DebugLevel)
	cfg := &ethlog.BlockHistoryConfig{
		Format:  ethlog.FormatYAML,
		Rewrite: true,
		ContractsData: []ethlog.ContractData{
			{
				Name:    "eventTestContract",
				ABI:     events_test_contract.EventsTestContractABI,
				Address: common.HexToAddress("0xa513E6E4b8f2a923D98304ec87F64353C4D5C853"), // contract address
			},
		},
	}
	historyMap, _ := e.RequestBlocksHistory(cfg)
	_ = e.DumpBlockHistory("all_blocks", cfg, historyMap)
	eventsMap, _ := e.RequestEventsHistory(cfg)
	err = e.DumpEventsByFile(cfg, eventsMap)
}

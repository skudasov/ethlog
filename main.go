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
	historyMap, _ := e.RequestBlocksHistory(cfg)
	_ = e.DumpBlockHistory("all_blocks", cfg, historyMap)
	eventsMap, _ := e.RequestEventsHistory(cfg)
	err = e.DumpEventsByFile(cfg, eventsMap)
}

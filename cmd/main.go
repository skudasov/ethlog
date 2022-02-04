package main

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rs/zerolog/log"
	"github.com/skudasov/ethlog/ethlog"
	"github.com/skudasov/ethlog/events_test_contract"
)

func main() {
	client, err := ethclient.Dial("ws://localhost:8545")
	if err != nil {
		log.Fatal().Err(err).Send()
	}
	e := ethlog.NewEthLog(client)
	cfg := &ethlog.HistoryConfig{
		FromBlock: big.NewInt(0),
		ToBlock:   big.NewInt(30),
		Format:    ethlog.FormatJSON,
		Rewrite:   true,
		ContractsData: []ethlog.ContractData{
			{
				Name:    "link",
				ABI:     events_test_contract.EventsTestContractABI,
				Address: common.HexToAddress("0x..."),
			},
		},
	}
	// dumps all the logs from 100 to 200 block with decoded data
	historyMap, _ := e.History(cfg)
	_ = e.DumpBlockHistory("all_blocks", cfg, historyMap)
	// dumps events of every contract in a file
	eventsMap, _ := e.RequestEventsHistory(cfg)
	_ = e.DumpEventsByFile(cfg, eventsMap)
}

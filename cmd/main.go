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
	//p := profile.Start(profile.MemProfile)
	//defer p.Stop()
	url := os.Getenv("URL")
	if url == "" {
		log.Fatal().Msg("Must provide URL=${ETH_URL}")
	}
	client, err := ethclient.Dial(url)
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
				Address: common.HexToAddress("0x949bEd886c739f1A3273629b3320db0C5024c719"),
			},
		},
	})
	if err != nil {
		log.Fatal().Err(err).Send()
	}
}

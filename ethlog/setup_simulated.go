package ethlog

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/rs/zerolog/log"
	"math/big"
)

type SimulatedBackendSuite struct {
	RootPrivateKey *ecdsa.PrivateKey
	Client         EthClient
	Auth           *bind.TransactOpts
}

type SimulatedBackendExt struct {
	*backends.SimulatedBackend
}

// BlockNumber returns the most recent block number
func (sbe *SimulatedBackendExt) BlockNumber(_ context.Context) (uint64, error) {
	bn := sbe.SimulatedBackend.Blockchain().CurrentBlock().Number().Uint64()
	return bn, nil
}

func NewSimulatedBackend() *SimulatedBackendSuite {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal().Err(err)
	}
	// nolint
	auth := bind.NewKeyedTransactor(privateKey)
	balance := new(big.Int)
	balance.SetString("100000000000000000000", 10) // 10 eth in wei

	address := auth.From
	genesisAlloc := map[common.Address]core.GenesisAccount{
		address: {
			Balance: balance,
		},
	}
	blockGasLimit := uint64(4712388)
	return &SimulatedBackendSuite{
		RootPrivateKey: privateKey,
		Client:         &SimulatedBackendExt{backends.NewSimulatedBackend(genesisAlloc, blockGasLimit)},
		Auth:           auth,
	}
}

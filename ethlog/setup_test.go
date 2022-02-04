package ethlog

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/skudasov/ethlog/another_events_test_contract"
	"github.com/skudasov/ethlog/events_test_contract"
)

func txOpts(client EthClient, fromAddress common.Address, privateKey *ecdsa.PrivateKey) *bind.TransactOpts {
	log.Debug().Str("addr", fromAddress.Hex()).Send()
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal().Err(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal().Err(err)
	}

	// nolint
	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = gasPrice
	return auth
}

func FireMultipleEvents(s *TestContracts) {
	tx, err := s.eventsContractInstance.FireMultipleEvents(txOpts(s.Client, s.RootAddr, s.RootPrivateKey))
	if err != nil {
		log.Fatal().Err(err).Send()
	}
	log.Debug().Str("tx_hash", tx.Hash().Hex()).Msg("tx sent")
}

func FireSetDataWithEvent(s *TestContracts) {
	tx, err := s.eventsContractInstance.Set(txOpts(s.Client, s.RootAddr, s.RootPrivateKey), big.NewInt(100500))
	if err != nil {
		log.Fatal().Err(err).Send()
	}
	log.Debug().Str("tx_hash", tx.Hash().Hex()).Msg("tx sent")
}

func FireGetData(s *TestContracts) {
	res, err := s.eventsContractInstance.Get(nil)
	if err != nil {
		log.Fatal().Err(err)
	}
	log.Debug().Int64("getter result", res.Int64()).Msg("getter result")
}

func FireNoIndexEvent(s *TestContracts) {
	tx, err := s.eventsContractInstance.FireNoIndexEvent(txOpts(s.Client, s.RootAddr, s.RootPrivateKey))
	if err != nil {
		log.Fatal().Err(err)
	}
	log.Debug().Str("tx_hash", tx.Hash().Hex()).Msg("tx sent")
}

func FireOneIndexEvent(s *TestContracts) {
	tx, err := s.eventsContractInstance.FireOneIndexEvent(txOpts(s.Client, s.RootAddr, s.RootPrivateKey))
	if err != nil {
		log.Fatal().Err(err)
	}
	log.Debug().Str("tx_hash", tx.Hash().Hex()).Msg("tx sent")
}

func FireNoIndexStructEvent(s *TestContracts) {
	tx, err := s.eventsContractInstance.FireNoIndexStructEvent(txOpts(s.Client, s.RootAddr, s.RootPrivateKey))
	if err != nil {
		log.Fatal().Err(err)
	}
	log.Debug().Str("tx_hash", tx.Hash().Hex()).Msg("tx sent")
}

func FireTwoIndexEventNext(s *TestContracts) {
	tx, err := s.anotherEventsContractInstance.FireTwoIndexEvent(txOpts(s.Client, s.RootAddr, s.RootPrivateKey))
	if err != nil {
		log.Fatal().Err(err)
	}
	log.Debug().Str("tx_hash", tx.Hash().Hex()).Msg("tx sent")
}

func FireTwoIndexEvent(s *TestContracts) {
	tx, err := s.eventsContractInstance.FireTwoIndexEvent(txOpts(s.Client, s.RootAddr, s.RootPrivateKey))
	if err != nil {
		log.Fatal().Err(err)
	}
	log.Debug().Str("tx_hash", tx.Hash().Hex()).Msg("tx sent")
}

func FireThreeIndexEvent(s *TestContracts) {
	tx, err := s.eventsContractInstance.FireThreeIndexEvent(txOpts(s.Client, s.RootAddr, s.RootPrivateKey))
	if err != nil {
		log.Fatal().Err(err)
	}
	log.Debug().Str("tx_hash", tx.Hash().Hex()).Msg("tx sent")
}

func FireThreeIndexEventNext(s *TestContracts) {
	tx, err := s.anotherEventsContractInstance.FireThreeIndexEvent(txOpts(s.Client, s.RootAddr, s.RootPrivateKey))
	if err != nil {
		log.Fatal().Err(err)
	}
	log.Debug().Str("tx_hash", tx.Hash().Hex()).Msg("tx sent")
}

type TestContracts struct {
	RootAddr                      common.Address
	RootPrivateKey                *ecdsa.PrivateKey
	Client                        EthClient
	AddrStore                     common.Address
	AddrStoreNext                 common.Address
	eventsContractInstance        *events_test_contract.EventsTestContract
	anotherEventsContractInstance *another_events_test_contract.AnotherEventsTestContract
}

func DeployTestContractsLocal(ds *SimulatedBackendSuite) *TestContracts {
	publicKey := ds.RootPrivateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal().Err(errors.New("not a public key"))
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	addrStore, _, instance, err := events_test_contract.DeployEventsTestContract(txOpts(ds.Client, fromAddress, ds.RootPrivateKey), ds.Client)
	if err != nil {
		log.Fatal().Err(err)
	}
	addrStoreNext, _, instanceNext, err := another_events_test_contract.DeployAnotherEventsTestContract(txOpts(ds.Client, fromAddress, ds.RootPrivateKey), ds.Client)
	if err != nil {
		log.Fatal().Err(err)
	}
	return &TestContracts{
		RootAddr:                      fromAddress,
		RootPrivateKey:                ds.RootPrivateKey,
		Client:                        ds.Client,
		AddrStore:                     addrStore,
		AddrStoreNext:                 addrStoreNext,
		eventsContractInstance:        instance,
		anotherEventsContractInstance: instanceNext,
	}
}

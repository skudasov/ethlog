package ethlog

import (
	"math/big"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/skudasov/ethlog/events_test_contract"
	"github.com/stretchr/testify/require"
)

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}

type SimulatedTestSuite struct {
	ss *SimulatedBackendSuite
	e  *EthLog
	c  *TestContracts
}

func NewSimulatedTestSuite() *SimulatedTestSuite {
	ss := NewSimulatedBackend()
	e := NewEthLog(ss.Client, 1000)
	c := DeployTestContractsLocal(ss)
	sts := &SimulatedTestSuite{
		ss: ss,
		e:  e,
		c:  c,
	}
	sts.commit()
	return sts
}

func (ss *SimulatedTestSuite) commit() {
	ss.c.Client.(*SimulatedBackendExt).SimulatedBackend.Commit()
}

func defaultCfg(contractAddr common.Address) *HistoryConfig {
	return &HistoryConfig{
		Format: FormatJSON,
		ContractsData: []ContractData{
			{
				Name:    "eventTestContract",
				ABI:     events_test_contract.EventsTestContractABI,
				Address: contractAddr,
			},
		},
	}
}

func TestEthlog(t *testing.T) {
	t.Run("can parse initial deployment blocks", func(t *testing.T) {
		ss := NewSimulatedTestSuite()
		cfg := defaultCfg(ss.c.AddrStore)
		historyResult, err := ss.e.History(cfg)
		require.NoError(t, err)
		require.Len(t, historyResult.Blocks, 2)
		require.Equal(t, int64(0), historyResult.Blocks[0].Number.Int64())
		require.Equal(t, 0, historyResult.Blocks[0].TransactionsCount)
		require.Equal(t, int64(1), historyResult.Blocks[1].Number.Int64())
		require.Equal(t, 2, historyResult.Blocks[1].TransactionsCount)
	})

	t.Run("can parse event data", func(t *testing.T) {
		ss := NewSimulatedTestSuite()
		FireSetDataWithEvent(ss.c)
		ss.commit()
		cfg := defaultCfg(ss.c.AddrStore)
		historyResult, err := ss.e.History(cfg)
		require.NoError(t, err)
		require.Len(t, historyResult.Blocks, 3)
		require.Equal(t, int64(2), historyResult.Blocks[2].Number.Int64())
		require.Equal(t, int64(100500),
			historyResult.Blocks[2].Transactions[0].Events[0].EventData["value"].(*big.Int).Int64())
	})

	t.Run("can parse tx input", func(t *testing.T) {
		ss := NewSimulatedTestSuite()
		FireSetDataWithEvent(ss.c)
		ss.commit()
		cfg := defaultCfg(ss.c.AddrStore)
		historyResult, err := ss.e.History(cfg)
		require.NoError(t, err)
		require.Len(t, historyResult.Blocks, 3)
		require.Equal(t, int64(2), historyResult.Blocks[2].Number.Int64())
		require.Equal(t, int64(100500),
			historyResult.Blocks[2].Transactions[0].Input["x"].(*big.Int).Int64())
	})

	t.Run("can parse tx output", func(t *testing.T) {
		ss := NewSimulatedTestSuite()
		FireSetDataWithEvent(ss.c)
		ss.commit()
		cfg := defaultCfg(ss.c.AddrStore)
		historyResult, err := ss.e.History(cfg)
		require.NoError(t, err)
		require.Len(t, historyResult.Blocks, 3)
		require.Equal(t, int64(2), historyResult.Blocks[2].Number.Int64())
		require.Equal(t, int64(100500),
			historyResult.Blocks[2].Transactions[0].Output[""].(*big.Int).Int64())
	})

	t.Run("can parse multiple events", func(t *testing.T) {
		ss := NewSimulatedTestSuite()
		FireMultipleEvents(ss.c)
		ss.commit()
		cfg := defaultCfg(ss.c.AddrStore)
		historyResult, err := ss.e.History(cfg)
		require.NoError(t, err)
		require.Len(t, historyResult.Blocks, 3)
		require.Equal(t, int64(2), historyResult.Blocks[2].Number.Int64())
		require.Equal(t, int64(3),
			historyResult.Blocks[2].Transactions[0].Events[0].EventData["startedAt"].(*big.Int).Int64())
		require.Equal(t, int64(3),
			historyResult.Blocks[2].Transactions[0].Events[1].EventData["startedAt"].(*big.Int).Int64())
	})

	t.Run("test all and dump", func(t *testing.T) {
		ss := NewSimulatedTestSuite()
		// calls
		FireSetDataWithEvent(ss.c)
		ss.commit()
		FireGetData(ss.c)
		ss.commit()
		// events
		FireNoIndexEvent(ss.c)
		ss.commit()
		FireNoIndexStructEvent(ss.c)
		ss.commit()
		FireOneIndexEvent(ss.c)
		ss.commit()
		FireTwoIndexEvent(ss.c)
		ss.commit()
		FireThreeIndexEvent(ss.c)
		ss.commit()

		cfg := defaultCfg(ss.c.AddrStore)
		cfg.OutputFile = "tmp_history"
		_, err := ss.e.History(cfg)
		require.NoError(t, err)
	})
}

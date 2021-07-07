package ethlog

import (
	"ethlog/events_test_contract"
	"github.com/ethereum/go-ethereum/common"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"
	"math/big"
	"os"
	"testing"
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
	e := NewEthLog(ss.Client)
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

func defaultCfg(contractAddr common.Address) *BlockHistoryConfig {
	return &BlockHistoryConfig{
		Format:  FormatJSON,
		Rewrite: true,
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
		historyResult, err := ss.e.RequestBlocksHistory(cfg)
		require.NoError(t, err)
		require.Len(t, historyResult.History, 2)
		require.Equal(t, int64(0), historyResult.History[0][KeyBlockNumber].(*big.Int).Int64())
		require.Equal(t, 0, historyResult.History[0][KeyBlockTransactionsCount])
		require.Equal(t, int64(1), historyResult.History[1][KeyBlockNumber].(*big.Int).Int64())
		require.Equal(t, 2, historyResult.History[1][KeyBlockTransactionsCount])
	})

	t.Run("can parse event data", func(t *testing.T) {
		ss := NewSimulatedTestSuite()
		FireSetDataWithEvent(ss.c)
		ss.commit()
		cfg := defaultCfg(ss.c.AddrStore)
		historyResult, err := ss.e.RequestBlocksHistory(cfg)
		require.NoError(t, err)
		require.Len(t, historyResult.History, 3)
		require.Equal(t, int64(2), historyResult.History[2][KeyBlockNumber].(*big.Int).Int64())
		require.Equal(t, int64(100500),
			historyResult.History[2][KeyTransactions].([]ParsedTx)[0]["events"].([]ParsedEvent)[0][KeyEventData].(RawParsedEventData)["value"].(*big.Int).Int64())
	})

	t.Run("can parse tx input", func(t *testing.T) {
		ss := NewSimulatedTestSuite()
		FireSetDataWithEvent(ss.c)
		ss.commit()
		cfg := defaultCfg(ss.c.AddrStore)
		historyResult, err := ss.e.RequestBlocksHistory(cfg)
		require.NoError(t, err)
		require.Len(t, historyResult.History, 3)
		require.Equal(t, int64(2), historyResult.History[2][KeyBlockNumber].(*big.Int).Int64())
		require.Equal(t, int64(100500),
			historyResult.History[2][KeyTransactions].([]ParsedTx)[0]["input"].(ParsedInput)["x"].(*big.Int).Int64())
	})

	t.Run("can parse tx output", func(t *testing.T) {
		ss := NewSimulatedTestSuite()
		FireSetDataWithEvent(ss.c)
		ss.commit()
		cfg := defaultCfg(ss.c.AddrStore)
		historyResult, err := ss.e.RequestBlocksHistory(cfg)
		require.NoError(t, err)
		require.Len(t, historyResult.History, 3)
		require.Equal(t, int64(2), historyResult.History[2][KeyBlockNumber].(*big.Int).Int64())
		require.Equal(t, int64(100500),
			historyResult.History[2][KeyTransactions].([]ParsedTx)[0]["output"].(ParsedOutput)[""].(*big.Int).Int64())
	})

	t.Run("can parse multiple events", func(t *testing.T) {
		ss := NewSimulatedTestSuite()
		FireMultipleEvents(ss.c)
		ss.commit()
		cfg := defaultCfg(ss.c.AddrStore)
		historyResult, err := ss.e.RequestBlocksHistory(cfg)
		require.NoError(t, err)
		require.Len(t, historyResult.History, 3)
		require.Equal(t, int64(2), historyResult.History[2][KeyBlockNumber].(*big.Int).Int64())
		require.Equal(t, int64(3),
			historyResult.History[2][KeyTransactions].([]ParsedTx)[0]["events"].([]ParsedEvent)[0][KeyEventData].(RawParsedEventData)["startedAt"].(*big.Int).Int64())
		require.Equal(t, int64(3),
			historyResult.History[2][KeyTransactions].([]ParsedTx)[0]["events"].([]ParsedEvent)[1][KeyEventData].(RawParsedEventData)["startedAt"].(*big.Int).Int64())
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
		historyResult, err := ss.e.RequestBlocksHistory(cfg)
		require.NoError(t, err)
		err = ss.e.DumpBlockHistory("example_history_simulated", cfg, historyResult)
		require.NoError(t, err)
	})
}

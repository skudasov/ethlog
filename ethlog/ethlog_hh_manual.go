package ethlog

//// TODO: see if simulated is OK for now
//
//import (
//	"ethlog/another_events_test_contract"
//	"ethlog/events_test_contract"
//	"github.com/stretchr/testify/require"
//	"testing"
//)
//
//func TestHardhat(t *testing.T) {
//	t.Run("get blocks history", func(t *testing.T) {
//		ss := NewLocalHardhatDeploySuite()
//		s := DeployTestContractsLocal(ss) // block 2
//		e := NewEthLog(s.Client)
//		// calls
//		FireSetDataWithEvent(s)
//		FireGetData(s)
//		// events
//		FireNoIndexEvent(s)
//		FireNoIndexStructEvent(s)
//		FireOneIndexEvent(s)
//		FireTwoIndexEvent(s)
//		FireThreeIndexEvent(s)
//
//		bhConfig := &BlockHistoryConfig{
//			//FromBlock: big.NewInt(0),
//			//ToBlock: big.NewInt(4),
//			Format:  FormatJSON,
//			Rewrite: true,
//			ContractsData: []ContractData{
//				{
//					Name:    "eventTestContract",
//					ABI:     events_test_contract.EventsTestContractABI,
//					Address: s.AddrStore,
//				},
//			},
//		}
//		historyResult, err := e.RequestBlocksHistory(bhConfig)
//		require.NoError(t, err)
//		err = e.DumpBlockHistory("example_history", bhConfig, historyResult)
//		require.NoError(t, err)
//	})
//
//	t.Run("get events dumped to files by contract", func(t *testing.T) {
//		ss := NewLocalHardhatDeploySuite()
//		s := DeployTestContractsLocal(ss) // block 2
//		e := NewEthLog(s.Client)
//		FireSetDataWithEvent(s)
//		FireGetData(s)
//		// events
//		FireNoIndexEvent(s)
//		FireNoIndexStructEvent(s)
//		FireOneIndexEvent(s)
//		FireTwoIndexEvent(s)
//		FireThreeIndexEvent(s)
//		FireThreeIndexEventNext(s)
//		// block 8
//
//		bhConfig := &BlockHistoryConfig{
//			Format:  FormatJSON,
//			Rewrite: true,
//			ContractsData: []ContractData{
//				{
//					Name:    "eventTestContract",
//					ABI:     events_test_contract.EventsTestContractABI,
//					Address: s.AddrStore,
//				},
//				{
//					Name:    "anotherEventTestContract",
//					ABI:     another_events_test_contract.AnotherEventsTestContractABI,
//					Address: s.AddrStoreNext,
//				},
//			},
//		}
//		events, err := e.RequestEventsHistory(bhConfig)
//		require.NoError(t, err)
//		err = e.DumpEventsByFile(bhConfig, events)
//		require.NoError(t, err)
//	})
//}
//

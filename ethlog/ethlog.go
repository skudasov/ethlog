package ethlog

import (
	"context"
	"encoding/json"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/fatih/structs"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v2"
	"math/big"
	"os"
	"strings"
)

type (
	ParsedBlock        map[string]interface{}
	ParsedTx           map[string]interface{}
	ParsedInput        map[string]interface{}
	ParsedOutput       map[string]interface{}
	ParsedReceipt      map[string]interface{}
	ParsedEvent        map[string]interface{}
	RawParsedEventData map[string]interface{}
)

type HistoryResult struct {
	History []ParsedBlock `json:"history" yaml:"history"`
}

type HistoryFormat int

const (
	// default keys
	KeyTransactions = "transactions"

	// blocks
	KeyBlockNumber            = "block_number"
	KeyBlockHash              = "block_hash"
	KeyBlockTransactionsCount = "block_tx_count"
	KeyBlockGasLimit          = "block_gas_limit"
	KeyBlockGasUsed           = "block_gas_used"
	KeyBlockTime              = "block_time"
	KeyBlockUncles            = "uncles"
	KeyBlockReceivedFrom      = "block_received_from"
	KeyBlockReceivedAt        = "block_received_at"

	// events
	KeyEventData            = "event_data"
	KeyEventMetaAddress     = "event_address"
	KeyEventMetaTopics      = "event_topics"
	KeyEventMetaBlockNumber = "event_block_number"
	KeyEventMetaIndex       = "index"
	KeyEventMetaTxIndex     = "tx_index"
	KeyEventMetaTxHash      = "tx_hash"
	KeyEventMetaTxRemoved   = "tx_removed"

	KeyEventFileTag = "file_tag"

	// enum
	FormatJSON HistoryFormat = iota
	FormatYAML
)

func mergeEventTopicsMap(eventsMap ParsedEvent, topicsMap ParsedEvent) ParsedEvent {
	m := make(ParsedEvent)
	m[KeyEventData] = make(RawParsedEventData)
	for k, v := range eventsMap {
		m[KeyEventData].(RawParsedEventData)[k] = v
	}
	for k, v := range topicsMap {
		m[KeyEventData].(RawParsedEventData)[k] = v
	}
	return m
}

type EthLog struct {
	client EthClient
	ABIS   []abi.ABI
}

type BlockHistoryConfig struct {
	FromBlock     *big.Int
	ToBlock       *big.Int
	Format        HistoryFormat
	Rewrite       bool
	ContractsData []ContractData
}

type ContractData struct {
	Name        string
	ABI         string
	ABIInstance abi.ABI
	Address     common.Address
}

type LatestBlockNumber interface {
	BlockNumber(context context.Context) (uint64, error)
}

type EthClient interface {
	ethereum.ChainReader
	ethereum.ChainStateReader
	ethereum.TransactionReader
	bind.ContractFilterer
	bind.ContractTransactor
	bind.ContractCaller
	LatestBlockNumber
}

func NewEthLog(client EthClient) *EthLog {
	return &EthLog{
		client: client,
		ABIS:   make([]abi.ABI, 0),
	}
}

// mergeLogMeta add metadata from log
func (e *EthLog) mergeLogMeta(m ParsedEvent, l types.Log) {
	m[KeyEventMetaAddress] = l.Address
	m[KeyEventMetaTopics] = l.Topics
	m[KeyEventMetaBlockNumber] = l.BlockNumber
	m[KeyEventMetaIndex] = l.Index
	m[KeyEventMetaTxHash] = l.TxHash
	m[KeyEventMetaTxIndex] = l.TxIndex
	m[KeyEventMetaTxRemoved] = l.Removed
}

// parseTxInputs parses tx inputs
func (e *EthLog) parseTxInputs(tx *types.Transaction, a abi.ABI) (map[string]interface{}, error) {
	inputMap := make(map[string]interface{})
	input := tx.Data()
	if len(input) == 0 {
		return nil, nil
	}
	// recover Method from signature and ABI
	method, err := a.MethodById(input[:4])
	if err != nil {
		return nil, err
	}
	// unpack method inputs
	err = method.Inputs.UnpackIntoMap(inputMap, input[4:])
	if err != nil {
		return nil, err
	}
	log.Debug().Interface("map", inputMap).Msg("parsed input")
	return inputMap, nil
}

// parseTxOutputs parses tx outputs
func (e *EthLog) parseTxOutputs(tx *types.Transaction, a abi.ABI) (map[string]interface{}, error) {
	outputMap := make(map[string]interface{})
	input := tx.Data()
	if len(input) == 0 {
		return nil, nil
	}
	// recover Method from signature and ABI
	method, err := a.MethodById(input[:4])
	if err != nil {
		return nil, err
	}
	// unpack method outputs
	err = method.Outputs.UnpackIntoMap(outputMap, input[4:])
	if err != nil {
		return nil, err
	}
	log.Debug().Interface("map", outputMap).Msg("parsed output")
	return outputMap, nil
}

// SetABI parse all ABIs from JSON and set to struct
func (e *EthLog) SetABI(bh *BlockHistoryConfig) error {
	for i := 0; i < len(bh.ContractsData); i++ {
		a, err := abi.JSON(strings.NewReader(bh.ContractsData[i].ABI))
		if err != nil {
			return err
		}
		bh.ContractsData[i].ABIInstance = a
	}
	return nil
}

func (e *EthLog) blockToMap(b *types.Block) ParsedBlock {
	return ParsedBlock{
		KeyBlockNumber:            b.Number(),
		KeyBlockHash:              b.Hash(),
		KeyBlockTransactionsCount: b.Transactions().Len(),
		KeyBlockGasLimit:          b.GasLimit(),
		KeyBlockGasUsed:           b.GasUsed(),
		KeyBlockTime:              b.Time(),
		KeyBlockUncles:            b.Uncles(),
		KeyBlockReceivedFrom:      b.ReceivedFrom,
		KeyBlockReceivedAt:        b.ReceivedAt,
	}
}

func (e *EthLog) txToMap(tx *types.Transaction) ParsedTx {
	m := make(ParsedTx)
	m["protected"] = tx.Protected()
	m["tx_hash"] = tx.Hash()
	return m
}

// RequestBlocksHistory parses all blocks with inputs/outputs, receipts and events for all contracts
func (e *EthLog) RequestBlocksHistory(bh *BlockHistoryConfig) (*HistoryResult, error) {
	if err := e.prepare(bh); err != nil {
		return nil, err
	}
	historyResult := &HistoryResult{}
	for i := bh.FromBlock.Int64(); i < bh.ToBlock.Int64(); i++ {
		b, err := e.client.BlockByNumber(context.Background(), big.NewInt(i))
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				log.Err(err).Msg("missing block")
			}
			return nil, err
		}
		parsedBlockMap := e.blockToMap(b)
		log.Debug().Str("block", b.Hash().Hex()).Int64("bn", b.Number().Int64()).Msg("block")
		log.Debug().Interface("block_map", parsedBlockMap).Msg("parsed block map")
		parsedTxs := make([]ParsedTx, 0)
		for _, tx := range b.Transactions() {
			var txInput ParsedInput
			var txOutput ParsedOutput
			var txReceipt ParsedReceipt
			var txEvents []ParsedEvent
			for _, cd := range bh.ContractsData {
				txInput, err = e.parseTxInputs(tx, cd.ABIInstance)
				if err != nil {
					if strings.Contains(err.Error(), "no method with id") {
						continue
					}
					return nil, err
				}
				txOutput, err = e.parseTxOutputs(tx, cd.ABIInstance)
				if err != nil {
					if strings.Contains(err.Error(), "no method with id") {
						continue
					}
					return nil, err
				}
				receipt, err := e.client.TransactionReceipt(context.Background(), tx.Hash())
				if err != nil {
					log.Err(err)
				}
				txReceipt = structs.Map(receipt)
				log.Debug().Interface("receipt", receipt).Msg("tx receipt")
				var logsValues []types.Log
				for _, l := range receipt.Logs {
					logsValues = append(logsValues, *l)
				}
				txEvents, err = e.parseContractEvents(logsValues, cd.ABIInstance, bh)
				if err != nil {
					return nil, err
				}
				log.Debug().Interface("events", txEvents).Msg("parsed events")
			}
			log.Debug().Interface("all_inputs", txInput).Msg("inputs collected")
			log.Debug().Interface("all_outputs", txOutput).Msg("inputs collected")
			parsedTx := e.txToMap(tx)
			parsedTx["input"] = txInput
			parsedTx["output"] = txOutput
			parsedTx["receipt"] = txReceipt
			parsedTx["events"] = txEvents
			parsedTxs = append(parsedTxs, parsedTx)
		}
		parsedBlockMap["transactions"] = parsedTxs
		historyResult.History = append(historyResult.History, parsedBlockMap)
	}
	log.Info().Interface("history", historyResult).Send()
	return historyResult, nil
}

func (e *EthLog) NewFile(name string) *os.File {
	f, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		log.Fatal().Err(err)
	}
	return f
}

func (e *EthLog) formatBlob(hr interface{}, bh *BlockHistoryConfig) ([]byte, string, error) {
	var fileSuffix string
	var blob []byte
	var err error
	switch bh.Format {
	case FormatJSON:
		fileSuffix = ".json"
		blob, err = e.formatPrettyJSON(hr)
	case FormatYAML:
		fileSuffix = ".yaml"
		blob, err = e.formatYAML(hr)
	default:
		log.Fatal().Err(errors.New("unknown output format")).Send()
	}
	if err != nil {
		return []byte{}, "", err
	}
	return blob, fileSuffix, nil
}

func (e *EthLog) DumpBlockHistory(name string, bh *BlockHistoryConfig, hr *HistoryResult) error {
	blob, fileSuffix, err := e.formatBlob(hr, bh)
	if err != nil {
		return err
	}
	f := e.NewFile(name + fileSuffix)
	_, err = f.Write(blob)
	if err != nil {
		return err
	}
	_ = f.Close()
	log.Debug().Bytes("blob", blob).Send()
	return nil
}

func (e *EthLog) DumpEventsByFile(bh *BlockHistoryConfig, eventsByContract [][]ParsedEvent) error {
	for _, events := range eventsByContract {
		if len(events) == 0 {
			log.Info().Msg("no events to dump")
			continue
		}
		blob, fileSuffix, err := e.formatBlob(events, bh)
		if err != nil {
			return err
		}
		f := e.NewFile(events[0][KeyEventFileTag].(string) + "_events" + fileSuffix)
		if _, err := f.Write(blob); err != nil {
			return err
		}
		_ = f.Close()
		log.Debug().Bytes("blob", blob).Send()
	}
	return nil
}

func (e *EthLog) ParseParams(bh *BlockHistoryConfig) {
	if bh.FromBlock == nil {
		bh.FromBlock = big.NewInt(0)
	}
	bn, err := e.client.BlockNumber(context.Background())
	if err != nil {
		log.Fatal().Err(err)
	}
	bn = bn + 1 // ToBlock is non inclusive on upper bound
	if bh.ToBlock == nil {
		bh.ToBlock = big.NewInt(int64(bn))
	}
}

func (e *EthLog) prepare(bh *BlockHistoryConfig) error {
	if err := e.SetABI(bh); err != nil {
		return err
	}
	e.ParseParams(bh)
	return nil
}

// RequestEventsHistory parse all events for all contracts
func (e *EthLog) RequestEventsHistory(bh *BlockHistoryConfig) ([][]ParsedEvent, error) {
	var allEvents [][]ParsedEvent
	if err := e.prepare(bh); err != nil {
		return nil, err
	}
	for _, d := range bh.ContractsData {
		logs, err := e.client.FilterLogs(context.Background(), ethereum.FilterQuery{
			FromBlock: bh.FromBlock,
			ToBlock:   bh.ToBlock,
			Addresses: []common.Address{
				d.Address,
			},
		})
		if err != nil {
			return allEvents, err
		}
		events, err := e.parseContractEvents(logs, d.ABIInstance, bh)
		if err != nil {
			return allEvents, err
		}
		if len(events) > 0 {
			events[0][KeyEventFileTag] = d.Name
		}
		allEvents = append(allEvents, events)
	}
	return allEvents, nil
}

// parseEventFromLog parses event body and additional index fields
func (e *EthLog) parseEventFromLog(a abi.ABI, eventABISpec abi.Event, l types.Log) (ParsedEvent, ParsedEvent, error) {
	eventsMap := make(map[string]interface{})
	topicsMap := make(map[string]interface{})
	// if no data event has only indexed fields
	if len(l.Data) != 0 {
		err := a.UnpackIntoMap(eventsMap, eventABISpec.RawName, l.Data)
		if err != nil {
			log.Err(err).Msg("error parsing non indexed data")
			return nil, nil, err
		}
	}
	// might have up to 3 additional index fields
	if len(l.Topics) > 1 {
		var indexed []abi.Argument
		for _, arg := range eventABISpec.Inputs {
			if arg.Indexed {
				indexed = append(indexed, arg)
			}
		}
		err := abi.ParseTopicsIntoMap(topicsMap, indexed, l.Topics[1:])
		if err != nil {
			return eventsMap, topicsMap, err
		}
	}
	return eventsMap, topicsMap, nil
}

func (e *EthLog) formatPrettyJSON(history interface{}) ([]byte, error) {
	d, err := json.MarshalIndent(history, "", "\t")
	if err != nil {
		return nil, err
	}
	return d, nil
}

func (e *EthLog) formatYAML(history interface{}) ([]byte, error) {
	d, err := yaml.Marshal(history)
	if err != nil {
		return nil, err
	}
	return d, nil
}

// parseContractEvents parses all events related to contract addr
func (e *EthLog) parseContractEvents(logs []types.Log, a abi.ABI, bh *BlockHistoryConfig) ([]ParsedEvent, error) {
	log.Info().
		Int64("from_block", bh.FromBlock.Int64()).
		Int64("to_block", bh.ToBlock.Int64()).
		Msg("parsing events")
	var eventsParsed []ParsedEvent
	for _, l := range logs {
		// TODO: iterate over possible events and lose speed, or waste some memory on map and lose ordering?
		for _, evSpec := range a.Events {
			sig := crypto.Keccak256Hash([]byte(evSpec.Sig)).Hex()
			// first topic is always 4 bytes keccak256 signature
			if sig == l.Topics[0].Hex() {
				log.Debug().Str("event_name", evSpec.RawName).Msg("unpacking event")
				log.Debug().Str("event_sig", evSpec.Sig).Msg("event signature")
				eventsMap, topicsMap, err := e.parseEventFromLog(a, evSpec, l)
				if err != nil {
					log.Err(err)
				}
				eventParsedMap := mergeEventTopicsMap(eventsMap, topicsMap)
				e.mergeLogMeta(eventParsedMap, l)
				eventsParsed = append(eventsParsed, eventParsedMap)
				log.Debug().Interface("event_map", eventParsedMap).Msg("event")
			}
		}
	}
	return eventsParsed, nil
}

package ethlog

import (
	"context"
	"encoding/json"
	"math/big"
	"net/http"
	"os"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go.uber.org/ratelimit"
	"golang.org/x/sync/errgroup"
	"gopkg.in/yaml.v2"
)

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	http.Handle("/metrics", promhttp.Handler())
	// nolint
	go http.ListenAndServe(":2222", nil)
}

var (
	getReceiptHist = promauto.NewHistogram(prometheus.HistogramOpts{
		Name: "get_receipt",
		Help: "TX Receipt response time",
	})
	getBlockHist = promauto.NewHistogram(prometheus.HistogramOpts{
		Name: "get_block",
		Help: "Get Block response time",
	})
)

type ParsedTX struct {
	ParseError string                 `json:"parse_error,omitempty"`
	Index      uint                   `json:"index"`
	Hash       string                 `json:"hash,omitempty"`
	Protected  bool                   `json:"protected,omitempty"`
	Input      map[string]interface{} `json:"input,omitempty"`
	Output     map[string]interface{} `json:"output,omitempty"`
	Receipt    *types.Receipt         `json:"receipt,omitempty"`
	Events     []ParsedEvent          `json:"events,omitempty"`
}

type ParsedBlock struct {
	Number            *big.Int   `json:"number"`
	Hash              string     `json:"hash"`
	TransactionsCount int        `json:"transactions_count"`
	Transactions      []ParsedTX `json:"transactions"`
	GasLimit          uint64     `json:"gas_limit"`
	GasUsed           uint64     `json:"gas_used"`
	Time              uint64     `json:"time"`
	Uncles            uint64     `json:"uncles"`
	ReceivedAt        time.Time  `json:"received_at"`
}

type History struct {
	Blocks []ParsedBlock `json:"history" yaml:"history"`
}

type HistoryFormat int

type ParsedEvent struct {
	Signature   string                 `json:"signature"`
	Address     common.Address         `json:"address"`
	EventData   map[string]interface{} `json:"event_data"`
	Topics      []string               `json:"topics,omitempty"`
	BlockNumber uint64                 `json:"block_number"`
	Index       uint                   `json:"index"`
	TXHash      string                 `json:"hash"`
	TXIndex     uint                   `json:"tx_index"`
	Removed     bool                   `json:"removed"`
	FileTag     string                 `json:"file_tag,omitempty"`
}

const (
	FormatJSON HistoryFormat = iota
	FormatYAML
)

func mergeEventTopicsMap(eventsMap map[string]interface{}, topicsMap map[string]interface{}) ParsedEvent {
	m := ParsedEvent{}
	m.EventData = map[string]interface{}{}
	for k, v := range eventsMap {
		m.EventData[k] = v
	}
	for k, v := range topicsMap {
		m.EventData[k] = v
	}
	return m
}

type EthLog struct {
	client EthClient
	rl     ratelimit.Limiter
	ABIS   []abi.ABI
}

type HistoryConfig struct {
	FromBlock     *big.Int
	ToBlock       *big.Int
	Format        HistoryFormat
	OutputFile    string
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

func NewEthLog(client EthClient, rpsLimit int) *EthLog {
	return &EthLog{
		client: client,
		ABIS:   make([]abi.ABI, 0),
		rl:     ratelimit.New(rpsLimit),
	}
}

// mergeLogMeta add metadata from log
func (e *EthLog) mergeLogMeta(m *ParsedEvent, l types.Log) {
	m.Address = l.Address
	m.Topics = make([]string, 0)
	for _, t := range l.Topics {
		m.Topics = append(m.Topics, t.String())
	}
	m.BlockNumber = l.BlockNumber
	m.Index = l.Index
	m.TXHash = l.TxHash.String()
	m.TXIndex = l.TxIndex
	m.Removed = l.Removed
}

// parseTxInputs parses tx inputs
func (e *EthLog) parseTxInputs(txData []byte, method *abi.Method) (map[string]interface{}, error) {
	log.Debug().Msg("Parsing tx inputs")
	inputMap := make(map[string]interface{})
	err := method.Inputs.UnpackIntoMap(inputMap, txData[4:])
	if err != nil {
		log.Error().Err(err).Send()
	}
	return inputMap, nil
}

// parseTxOutputs parses tx outputs
func (e *EthLog) parseTxOutputs(txData []byte, method *abi.Method) (map[string]interface{}, error) {
	log.Debug().Msg("Parsing tx outputs")
	outputMap := make(map[string]interface{})
	// unpack method outputs
	err := method.Outputs.UnpackIntoMap(outputMap, txData[4:])
	if err != nil {
		log.Error().Err(err).Send()
	}
	return outputMap, nil
}

// SetABI parse all ABIs from JSON and set to struct
func (e *EthLog) SetABI(bh *HistoryConfig) error {
	for i := 0; i < len(bh.ContractsData); i++ {
		a, err := abi.JSON(strings.NewReader(bh.ContractsData[i].ABI))
		if err != nil {
			return err
		}
		bh.ContractsData[i].ABIInstance = a
	}
	return nil
}

func (e *EthLog) parseTX(tx *types.Transaction, bh *HistoryConfig) (ParsedTX, error) {
	var txInput map[string]interface{}
	var txOutput map[string]interface{}
	var txEvents []ParsedEvent
	var err error
	e.rl.Take()
	before := time.Now()
	receipt, err := e.client.TransactionReceipt(context.Background(), tx.Hash())
	if err != nil {
		return ParsedTX{}, err
	}
	getReceiptHist.Observe(float64(time.Since(before).Milliseconds()))
	txData := tx.Data()
	// if there is no input data hence no events
	if len(txData) == 0 {
		return ParsedTX{
			Index:   receipt.TransactionIndex,
			Receipt: receipt,
		}, nil
	}
	// parse inputs and events
	for _, cd := range bh.ContractsData {
		// check if we have that method in that ABI
		method, err := cd.ABIInstance.MethodById(txData[:4])
		if err != nil {
			log.Debug().Msg("Method not found")
			continue
		}
		txInput, err = e.parseTxInputs(txData, method)
		if err != nil {
			return ParsedTX{}, err
		}
		txOutput, err = e.parseTxOutputs(txData, method)
		if err != nil {
			return ParsedTX{}, err
		}
		if receipt != nil {
			log.Debug().Interface("receipt", receipt).Msg("tx receipt")
			logsValues := make([]types.Log, 0)
			for _, l := range receipt.Logs {
				logsValues = append(logsValues, *l)
			}
			txEvents, err = e.parseContractEvents(logsValues, cd.ABIInstance)
			if err != nil {
				return ParsedTX{}, err
			}
		}
	}
	log.Debug().Interface("Inputs", txInput).Msg("Inputs collected")
	log.Debug().Interface("Outputs", txOutput).Msg("Outputs collected")
	log.Debug().Interface("Events", txEvents).Msg("Events collected")
	return ParsedTX{
		Index:     receipt.TransactionIndex,
		Protected: tx.Protected(),
		Hash:      tx.Hash().String(),
		Input:     txInput,
		Output:    txOutput,
		Receipt:   receipt,
		Events:    txEvents,
	}, nil
}

func (e *EthLog) orderHistory(h *History) {
	sort.Slice(h.Blocks, func(i, j int) bool {
		return h.Blocks[i].Number.Int64() < h.Blocks[j].Number.Int64()
	})
	for _, block := range h.Blocks {
		sort.Slice(block.Transactions, func(i, j int) bool {
			return block.Transactions[i].Index < block.Transactions[j].Index
		})
	}
}

// History parses all blocks with inputs/outputs, receipts and events for all contracts
func (e *EthLog) History(bh *HistoryConfig) (*History, error) {
	if err := e.prepare(bh); err != nil {
		return nil, err
	}
	history := &History{}
	bmu := &sync.Mutex{}
	beg := errgroup.Group{}
	for i := bh.FromBlock.Int64(); i < bh.ToBlock.Int64(); i++ {
		i := i
		beg.Go(func() error {
			log.Info().Int64("Number", i).Msg("Parsing block")
			e.rl.Take()
			tn := time.Now()
			b, err := e.client.BlockByNumber(context.Background(), big.NewInt(i))
			if err != nil {
				if strings.Contains(err.Error(), "not found") {
					log.Err(err).Msg("missing block")
				}
				return err
			}
			getBlockHist.Observe(float64(time.Since(tn).Milliseconds()))
			parsedBlock := ParsedBlock{
				Number:            b.Number(),
				Hash:              b.Hash().String(),
				TransactionsCount: b.Transactions().Len(),
				GasLimit:          b.GasLimit(),
				GasUsed:           b.GasUsed(),
				Time:              b.Time(),
				ReceivedAt:        b.ReceivedAt,
			}
			mu := &sync.Mutex{}
			eg := errgroup.Group{}
			for _, tx := range b.Transactions() {
				tx := tx
				eg.Go(func() error {
					parsedTx, err := e.parseTX(tx, bh)
					if err != nil {
						return err
					}
					mu.Lock()
					parsedBlock.Transactions = append(parsedBlock.Transactions, parsedTx)
					mu.Unlock()
					return nil
				})
			}
			if err := eg.Wait(); err != nil {
				return err
			}
			bmu.Lock()
			history.Blocks = append(history.Blocks, parsedBlock)
			bmu.Unlock()
			return nil
		})
	}
	if err := beg.Wait(); err != nil {
		return nil, err
	}
	e.orderHistory(history)
	if bh.OutputFile != "" {
		if err := e.writeBlockHistory(bh.OutputFile, bh.Format, history); err != nil {
			return nil, err
		}
	}
	return history, nil
}

func (e *EthLog) NewFile(name string) *os.File {
	f, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		log.Fatal().Err(err)
	}
	return f
}

func (e *EthLog) formatBlob(hr interface{}, format HistoryFormat) ([]byte, string, error) {
	var fileSuffix string
	var blob []byte
	var err error
	switch format {
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

func (e *EthLog) writeBlockHistory(name string, format HistoryFormat, hr *History) error {
	blob, fileSuffix, err := e.formatBlob(hr, format)
	if err != nil {
		return err
	}
	if err := os.WriteFile(name+fileSuffix, blob, os.ModePerm); err != nil {
		return err
	}
	return nil
}

func (e *EthLog) DumpEventsByFile(bh *HistoryConfig, eventsByContract [][]ParsedEvent) error {
	for _, events := range eventsByContract {
		if len(events) == 0 {
			log.Info().Msg("no events to dump")
			continue
		}
		blob, fileSuffix, err := e.formatBlob(events, bh.Format)
		if err != nil {
			return err
		}
		if err := os.WriteFile(events[0].FileTag+"_events"+fileSuffix, blob, os.ModePerm); err != nil {
			return err
		}
	}
	return nil
}

func (e *EthLog) prepare(bh *HistoryConfig) error {
	if err := e.SetABI(bh); err != nil {
		return err
	}
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
	return nil
}

// EventsHistory parse all events for all contracts
func (e *EthLog) EventsHistory(bh *HistoryConfig) ([][]ParsedEvent, error) {
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
		events, err := e.parseContractEvents(logs, d.ABIInstance)
		if err != nil {
			return allEvents, err
		}
		if len(events) > 0 {
			events[0].FileTag = d.Name
		}
		allEvents = append(allEvents, events)
	}
	return allEvents, nil
}

// parseEventFromLog parses event body and additional index fields
func (e *EthLog) parseEventFromLog(a abi.ABI, eventABISpec abi.Event, l types.Log) (map[string]interface{}, map[string]interface{}, error) {
	eventsMap := make(map[string]interface{})
	topicsMap := make(map[string]interface{})
	// if no data event has only indexed fields
	if len(l.Data) != 0 {
		err := a.UnpackIntoMap(eventsMap, eventABISpec.RawName, l.Data)
		if err != nil {
			log.Err(err).Msg("error parsing non indexed data")
			return nil, nil, err
		}
		log.Debug().Interface("Non-indexed", eventsMap).Send()
	}
	// might have up to 3 additional indexed fields
	if len(l.Topics) > 1 {
		topics := l.Topics[1:]
		var indexed []abi.Argument
		indexedTopics := make([]common.Hash, 0)
		for idx, topic := range topics {
			arg := eventABISpec.Inputs[idx]
			if arg.Indexed {
				indexed = append(indexed, arg)
				indexedTopics = append(indexedTopics, topic)
			}
		}
		log.Debug().Int("Topics", len(l.Topics[1:])).Int("Arguments", len(indexed)).Send()
		log.Debug().Interface("AllTopics", l.Topics).Send()
		log.Debug().Interface("HashOfName", eventABISpec.ID.Hex()).Send()
		log.Debug().Interface("FirstTopic", l.Topics[0]).Send()
		log.Debug().Interface("Topics", l.Topics[1:]).Send()
		log.Debug().Interface("Arguments", eventABISpec.Inputs).Send()
		log.Debug().Interface("Indexed", indexed).Send()
		err := abi.ParseTopicsIntoMap(topicsMap, indexed, indexedTopics)
		if err != nil {
			return nil, nil, err
		}
		log.Debug().Interface("Indexed", topicsMap).Send()
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
func (e *EthLog) parseContractEvents(logs []types.Log, a abi.ABI) ([]ParsedEvent, error) {
	log.Debug().
		Msg("Parsing events")
	var eventsParsed []ParsedEvent
	for _, l := range logs {
		for _, evSpec := range a.Events {
			// evSpec.ID.Hex() == crypto.Keccak256Hash([]byte(evSpec.Sig)).Hex() ?
			// sig := crypto.Keccak256Hash([]byte(evSpec.Sig)).Hex()
			// first topic is always 4 bytes keccak256 signature
			if evSpec.ID.Hex() == l.Topics[0].Hex() {
				log.Debug().Str("Name", evSpec.RawName).Str("Signature", evSpec.Sig).Msg("Unpacking event")
				eventsMap, topicsMap, err := e.parseEventFromLog(a, evSpec, l)
				if err != nil {
					return nil, err
				}
				parsedEvent := mergeEventTopicsMap(eventsMap, topicsMap)
				parsedEvent.Signature = evSpec.Sig
				e.mergeLogMeta(&parsedEvent, l)
				eventsParsed = append(eventsParsed, parsedEvent)
				log.Debug().Interface("Event", parsedEvent).Msg("Parsed event map")
			}
		}
	}
	return eventsParsed, nil
}

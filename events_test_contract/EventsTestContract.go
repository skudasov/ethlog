// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package events_test_contract

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// EventsTestContractAccount is an auto generated low-level Go binding around an user-defined struct.
type EventsTestContractAccount struct {
	Name       string
	Balance    uint64
	Dailylimit *big.Int
}

// EventsTestContractABI is the input ABI used to generate the binding from.
const EventsTestContractABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"NoIndexEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"balance\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"dailylimit\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structEventsTestContract.Account\",\"name\":\"a\",\"type\":\"tuple\"}],\"name\":\"NoIndexStructEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"a\",\"type\":\"string\"}],\"name\":\"OneIndexEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"startedBy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"}],\"name\":\"ThreeIndexEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"startedBy\",\"type\":\"address\"}],\"name\":\"TwoIndexEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"ValueEvent\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"fireMultipleEvents\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fireNoIndexEvent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fireNoIndexStructEvent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fireOneIndexEvent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fireThreeIndexEvent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fireTwoIndexEvent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"}],\"name\":\"set\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// EventsTestContractBin is the compiled bytecode used for deploying new contracts.
var EventsTestContractBin = "0x608060405234801561001057600080fd5b50610753806100206000396000f3fe608060405234801561001057600080fd5b50600436106100875760003560e01c80636d28a01f1161005b5780636d28a01f146100da5780636d4ce63c146100e457806395a777bf14610102578063eac9dc871461010c57610087565b80625fd3511461008c578063425aa72914610096578063577dd294146100a057806360fe47b1146100aa575b600080fd5b610094610116565b005b61009e6101aa565b005b6100a861024e565b005b6100c460048036038101906100bf91906103fc565b610295565b6040516100d1919061058d565b60405180910390f35b6100e2610303565b005b6100ec610345565b6040516100f9919061058d565b60405180910390f35b61010a61035b565b005b610114610394565b005b7febe3ff7e2071d351bf2e65b4fccd24e3ae99485f02468f1feecf7d64dc04418860405180606001604052806040518060400160405280600481526020017f4a6f686e000000000000000000000000000000000000000000000000000000008152508152602001600567ffffffffffffffff168152602001600a8152506040516101a0919061056b565b60405180910390a1565b3373ffffffffffffffffffffffffffffffffffffffff1660017f5660e8f93f0146f45abcd659e026b75995db50053cbbca4d7f365934ade68bf360036040516101f39190610550565b60405180910390a33373ffffffffffffffffffffffffffffffffffffffff1660017f5660e8f93f0146f45abcd659e026b75995db50053cbbca4d7f365934ade68bf360036040516102449190610550565b60405180910390a3565b3373ffffffffffffffffffffffffffffffffffffffff1660017f33b47a1cd66813164ec00800d74296f57415217c22505ee380594a712936a0b560405160405180910390a3565b6000816000819055503373ffffffffffffffffffffffffffffffffffffffff167f834ba0b35948e5b0a41be62c5c77104608331d60ee912667365f453865a6dbaf836040516102e4919061058d565b60405180910390a26005826102f991906105cf565b9150819050919050565b60405161030f90610520565b60405180910390207f164b5941f5ac984fc056425d6baefece85522d55cef52e1902da35cec5c575ad60405160405180910390a2565b6000600a60005461035691906105cf565b905090565b7f33bc9bae48dbe1e057f264b3fc6a1dacdcceacb3ba28d937231c70e068a02f1c3360405161038a9190610535565b60405180910390a1565b3373ffffffffffffffffffffffffffffffffffffffff1660017f5660e8f93f0146f45abcd659e026b75995db50053cbbca4d7f365934ade68bf360036040516103dd9190610550565b60405180910390a3565b6000813590506103f681610706565b92915050565b600060208284031215610412576104116106c7565b5b6000610420848285016103e7565b91505092915050565b61043281610603565b82525050565b61044181610653565b82525050565b6000610452826105a8565b61045c81856105b3565b935061046c818560208601610665565b610475816106cc565b840191505092915050565b600061048d6011836105c4565b9150610498826106dd565b601182019050919050565b600060608301600083015184820360008601526104c08282610447565b91505060208301516104d56020860182610511565b5060408301516104e860408601826104f3565b508091505092915050565b6104fc81610635565b82525050565b61050b81610635565b82525050565b61051a8161063f565b82525050565b600061052b82610480565b9150819050919050565b600060208201905061054a6000830184610429565b92915050565b60006020820190506105656000830184610438565b92915050565b6000602082019050818103600083015261058581846104a3565b905092915050565b60006020820190506105a26000830184610502565b92915050565b600081519050919050565b600082825260208201905092915050565b600081905092915050565b60006105da82610635565b91506105e583610635565b9250828210156105f8576105f7610698565b5b828203905092915050565b600061060e82610615565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600067ffffffffffffffff82169050919050565b600061065e82610635565b9050919050565b60005b83811015610683578082015181840152602081019050610668565b83811115610692576000848401525b50505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600080fd5b6000601f19601f8301169050919050565b7f6576656e745f646174615f737472696e67000000000000000000000000000000600082015250565b61070f81610635565b811461071a57600080fd5b5056fea26469706673582212206bc22637aec15c2a0a7eb0489412843b746f00259c9eca4dbb369cfe81bbd5eb64736f6c63430008060033"

// DeployEventsTestContract deploys a new Ethereum contract, binding an instance of EventsTestContract to it.
func DeployEventsTestContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *EventsTestContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EventsTestContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(EventsTestContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EventsTestContract{EventsTestContractCaller: EventsTestContractCaller{contract: contract}, EventsTestContractTransactor: EventsTestContractTransactor{contract: contract}, EventsTestContractFilterer: EventsTestContractFilterer{contract: contract}}, nil
}

// EventsTestContract is an auto generated Go binding around an Ethereum contract.
type EventsTestContract struct {
	EventsTestContractCaller     // Read-only binding to the contract
	EventsTestContractTransactor // Write-only binding to the contract
	EventsTestContractFilterer   // Log filterer for contract events
}

// EventsTestContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type EventsTestContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EventsTestContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EventsTestContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EventsTestContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EventsTestContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EventsTestContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EventsTestContractSession struct {
	Contract     *EventsTestContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// EventsTestContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EventsTestContractCallerSession struct {
	Contract *EventsTestContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// EventsTestContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EventsTestContractTransactorSession struct {
	Contract     *EventsTestContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// EventsTestContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type EventsTestContractRaw struct {
	Contract *EventsTestContract // Generic contract binding to access the raw methods on
}

// EventsTestContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EventsTestContractCallerRaw struct {
	Contract *EventsTestContractCaller // Generic read-only contract binding to access the raw methods on
}

// EventsTestContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EventsTestContractTransactorRaw struct {
	Contract *EventsTestContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEventsTestContract creates a new instance of EventsTestContract, bound to a specific deployed contract.
func NewEventsTestContract(address common.Address, backend bind.ContractBackend) (*EventsTestContract, error) {
	contract, err := bindEventsTestContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EventsTestContract{EventsTestContractCaller: EventsTestContractCaller{contract: contract}, EventsTestContractTransactor: EventsTestContractTransactor{contract: contract}, EventsTestContractFilterer: EventsTestContractFilterer{contract: contract}}, nil
}

// NewEventsTestContractCaller creates a new read-only instance of EventsTestContract, bound to a specific deployed contract.
func NewEventsTestContractCaller(address common.Address, caller bind.ContractCaller) (*EventsTestContractCaller, error) {
	contract, err := bindEventsTestContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EventsTestContractCaller{contract: contract}, nil
}

// NewEventsTestContractTransactor creates a new write-only instance of EventsTestContract, bound to a specific deployed contract.
func NewEventsTestContractTransactor(address common.Address, transactor bind.ContractTransactor) (*EventsTestContractTransactor, error) {
	contract, err := bindEventsTestContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EventsTestContractTransactor{contract: contract}, nil
}

// NewEventsTestContractFilterer creates a new log filterer instance of EventsTestContract, bound to a specific deployed contract.
func NewEventsTestContractFilterer(address common.Address, filterer bind.ContractFilterer) (*EventsTestContractFilterer, error) {
	contract, err := bindEventsTestContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EventsTestContractFilterer{contract: contract}, nil
}

// bindEventsTestContract binds a generic wrapper to an already deployed contract.
func bindEventsTestContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EventsTestContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EventsTestContract *EventsTestContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EventsTestContract.Contract.EventsTestContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EventsTestContract *EventsTestContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EventsTestContract.Contract.EventsTestContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EventsTestContract *EventsTestContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EventsTestContract.Contract.EventsTestContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EventsTestContract *EventsTestContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EventsTestContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EventsTestContract *EventsTestContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EventsTestContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EventsTestContract *EventsTestContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EventsTestContract.Contract.contract.Transact(opts, method, params...)
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns(uint256)
func (_EventsTestContract *EventsTestContractCaller) Get(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EventsTestContract.contract.Call(opts, &out, "get")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns(uint256)
func (_EventsTestContract *EventsTestContractSession) Get() (*big.Int, error) {
	return _EventsTestContract.Contract.Get(&_EventsTestContract.CallOpts)
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns(uint256)
func (_EventsTestContract *EventsTestContractCallerSession) Get() (*big.Int, error) {
	return _EventsTestContract.Contract.Get(&_EventsTestContract.CallOpts)
}

// FireMultipleEvents is a paid mutator transaction binding the contract method 0x425aa729.
//
// Solidity: function fireMultipleEvents() returns()
func (_EventsTestContract *EventsTestContractTransactor) FireMultipleEvents(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EventsTestContract.contract.Transact(opts, "fireMultipleEvents")
}

// FireMultipleEvents is a paid mutator transaction binding the contract method 0x425aa729.
//
// Solidity: function fireMultipleEvents() returns()
func (_EventsTestContract *EventsTestContractSession) FireMultipleEvents() (*types.Transaction, error) {
	return _EventsTestContract.Contract.FireMultipleEvents(&_EventsTestContract.TransactOpts)
}

// FireMultipleEvents is a paid mutator transaction binding the contract method 0x425aa729.
//
// Solidity: function fireMultipleEvents() returns()
func (_EventsTestContract *EventsTestContractTransactorSession) FireMultipleEvents() (*types.Transaction, error) {
	return _EventsTestContract.Contract.FireMultipleEvents(&_EventsTestContract.TransactOpts)
}

// FireNoIndexEvent is a paid mutator transaction binding the contract method 0x95a777bf.
//
// Solidity: function fireNoIndexEvent() returns()
func (_EventsTestContract *EventsTestContractTransactor) FireNoIndexEvent(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EventsTestContract.contract.Transact(opts, "fireNoIndexEvent")
}

// FireNoIndexEvent is a paid mutator transaction binding the contract method 0x95a777bf.
//
// Solidity: function fireNoIndexEvent() returns()
func (_EventsTestContract *EventsTestContractSession) FireNoIndexEvent() (*types.Transaction, error) {
	return _EventsTestContract.Contract.FireNoIndexEvent(&_EventsTestContract.TransactOpts)
}

// FireNoIndexEvent is a paid mutator transaction binding the contract method 0x95a777bf.
//
// Solidity: function fireNoIndexEvent() returns()
func (_EventsTestContract *EventsTestContractTransactorSession) FireNoIndexEvent() (*types.Transaction, error) {
	return _EventsTestContract.Contract.FireNoIndexEvent(&_EventsTestContract.TransactOpts)
}

// FireNoIndexStructEvent is a paid mutator transaction binding the contract method 0x005fd351.
//
// Solidity: function fireNoIndexStructEvent() returns()
func (_EventsTestContract *EventsTestContractTransactor) FireNoIndexStructEvent(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EventsTestContract.contract.Transact(opts, "fireNoIndexStructEvent")
}

// FireNoIndexStructEvent is a paid mutator transaction binding the contract method 0x005fd351.
//
// Solidity: function fireNoIndexStructEvent() returns()
func (_EventsTestContract *EventsTestContractSession) FireNoIndexStructEvent() (*types.Transaction, error) {
	return _EventsTestContract.Contract.FireNoIndexStructEvent(&_EventsTestContract.TransactOpts)
}

// FireNoIndexStructEvent is a paid mutator transaction binding the contract method 0x005fd351.
//
// Solidity: function fireNoIndexStructEvent() returns()
func (_EventsTestContract *EventsTestContractTransactorSession) FireNoIndexStructEvent() (*types.Transaction, error) {
	return _EventsTestContract.Contract.FireNoIndexStructEvent(&_EventsTestContract.TransactOpts)
}

// FireOneIndexEvent is a paid mutator transaction binding the contract method 0x6d28a01f.
//
// Solidity: function fireOneIndexEvent() returns()
func (_EventsTestContract *EventsTestContractTransactor) FireOneIndexEvent(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EventsTestContract.contract.Transact(opts, "fireOneIndexEvent")
}

// FireOneIndexEvent is a paid mutator transaction binding the contract method 0x6d28a01f.
//
// Solidity: function fireOneIndexEvent() returns()
func (_EventsTestContract *EventsTestContractSession) FireOneIndexEvent() (*types.Transaction, error) {
	return _EventsTestContract.Contract.FireOneIndexEvent(&_EventsTestContract.TransactOpts)
}

// FireOneIndexEvent is a paid mutator transaction binding the contract method 0x6d28a01f.
//
// Solidity: function fireOneIndexEvent() returns()
func (_EventsTestContract *EventsTestContractTransactorSession) FireOneIndexEvent() (*types.Transaction, error) {
	return _EventsTestContract.Contract.FireOneIndexEvent(&_EventsTestContract.TransactOpts)
}

// FireThreeIndexEvent is a paid mutator transaction binding the contract method 0xeac9dc87.
//
// Solidity: function fireThreeIndexEvent() returns()
func (_EventsTestContract *EventsTestContractTransactor) FireThreeIndexEvent(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EventsTestContract.contract.Transact(opts, "fireThreeIndexEvent")
}

// FireThreeIndexEvent is a paid mutator transaction binding the contract method 0xeac9dc87.
//
// Solidity: function fireThreeIndexEvent() returns()
func (_EventsTestContract *EventsTestContractSession) FireThreeIndexEvent() (*types.Transaction, error) {
	return _EventsTestContract.Contract.FireThreeIndexEvent(&_EventsTestContract.TransactOpts)
}

// FireThreeIndexEvent is a paid mutator transaction binding the contract method 0xeac9dc87.
//
// Solidity: function fireThreeIndexEvent() returns()
func (_EventsTestContract *EventsTestContractTransactorSession) FireThreeIndexEvent() (*types.Transaction, error) {
	return _EventsTestContract.Contract.FireThreeIndexEvent(&_EventsTestContract.TransactOpts)
}

// FireTwoIndexEvent is a paid mutator transaction binding the contract method 0x577dd294.
//
// Solidity: function fireTwoIndexEvent() returns()
func (_EventsTestContract *EventsTestContractTransactor) FireTwoIndexEvent(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EventsTestContract.contract.Transact(opts, "fireTwoIndexEvent")
}

// FireTwoIndexEvent is a paid mutator transaction binding the contract method 0x577dd294.
//
// Solidity: function fireTwoIndexEvent() returns()
func (_EventsTestContract *EventsTestContractSession) FireTwoIndexEvent() (*types.Transaction, error) {
	return _EventsTestContract.Contract.FireTwoIndexEvent(&_EventsTestContract.TransactOpts)
}

// FireTwoIndexEvent is a paid mutator transaction binding the contract method 0x577dd294.
//
// Solidity: function fireTwoIndexEvent() returns()
func (_EventsTestContract *EventsTestContractTransactorSession) FireTwoIndexEvent() (*types.Transaction, error) {
	return _EventsTestContract.Contract.FireTwoIndexEvent(&_EventsTestContract.TransactOpts)
}

// Set is a paid mutator transaction binding the contract method 0x60fe47b1.
//
// Solidity: function set(uint256 x) returns(uint256)
func (_EventsTestContract *EventsTestContractTransactor) Set(opts *bind.TransactOpts, x *big.Int) (*types.Transaction, error) {
	return _EventsTestContract.contract.Transact(opts, "set", x)
}

// Set is a paid mutator transaction binding the contract method 0x60fe47b1.
//
// Solidity: function set(uint256 x) returns(uint256)
func (_EventsTestContract *EventsTestContractSession) Set(x *big.Int) (*types.Transaction, error) {
	return _EventsTestContract.Contract.Set(&_EventsTestContract.TransactOpts, x)
}

// Set is a paid mutator transaction binding the contract method 0x60fe47b1.
//
// Solidity: function set(uint256 x) returns(uint256)
func (_EventsTestContract *EventsTestContractTransactorSession) Set(x *big.Int) (*types.Transaction, error) {
	return _EventsTestContract.Contract.Set(&_EventsTestContract.TransactOpts, x)
}

// EventsTestContractNoIndexEventIterator is returned from FilterNoIndexEvent and is used to iterate over the raw logs and unpacked data for NoIndexEvent events raised by the EventsTestContract contract.
type EventsTestContractNoIndexEventIterator struct {
	Event *EventsTestContractNoIndexEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EventsTestContractNoIndexEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventsTestContractNoIndexEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EventsTestContractNoIndexEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EventsTestContractNoIndexEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventsTestContractNoIndexEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventsTestContractNoIndexEvent represents a NoIndexEvent event raised by the EventsTestContract contract.
type EventsTestContractNoIndexEvent struct {
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNoIndexEvent is a free log retrieval operation binding the contract event 0x33bc9bae48dbe1e057f264b3fc6a1dacdcceacb3ba28d937231c70e068a02f1c.
//
// Solidity: event NoIndexEvent(address sender)
func (_EventsTestContract *EventsTestContractFilterer) FilterNoIndexEvent(opts *bind.FilterOpts) (*EventsTestContractNoIndexEventIterator, error) {

	logs, sub, err := _EventsTestContract.contract.FilterLogs(opts, "NoIndexEvent")
	if err != nil {
		return nil, err
	}
	return &EventsTestContractNoIndexEventIterator{contract: _EventsTestContract.contract, event: "NoIndexEvent", logs: logs, sub: sub}, nil
}

// WatchNoIndexEvent is a free log subscription operation binding the contract event 0x33bc9bae48dbe1e057f264b3fc6a1dacdcceacb3ba28d937231c70e068a02f1c.
//
// Solidity: event NoIndexEvent(address sender)
func (_EventsTestContract *EventsTestContractFilterer) WatchNoIndexEvent(opts *bind.WatchOpts, sink chan<- *EventsTestContractNoIndexEvent) (event.Subscription, error) {

	logs, sub, err := _EventsTestContract.contract.WatchLogs(opts, "NoIndexEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventsTestContractNoIndexEvent)
				if err := _EventsTestContract.contract.UnpackLog(event, "NoIndexEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNoIndexEvent is a log parse operation binding the contract event 0x33bc9bae48dbe1e057f264b3fc6a1dacdcceacb3ba28d937231c70e068a02f1c.
//
// Solidity: event NoIndexEvent(address sender)
func (_EventsTestContract *EventsTestContractFilterer) ParseNoIndexEvent(log types.Log) (*EventsTestContractNoIndexEvent, error) {
	event := new(EventsTestContractNoIndexEvent)
	if err := _EventsTestContract.contract.UnpackLog(event, "NoIndexEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EventsTestContractNoIndexStructEventIterator is returned from FilterNoIndexStructEvent and is used to iterate over the raw logs and unpacked data for NoIndexStructEvent events raised by the EventsTestContract contract.
type EventsTestContractNoIndexStructEventIterator struct {
	Event *EventsTestContractNoIndexStructEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EventsTestContractNoIndexStructEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventsTestContractNoIndexStructEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EventsTestContractNoIndexStructEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EventsTestContractNoIndexStructEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventsTestContractNoIndexStructEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventsTestContractNoIndexStructEvent represents a NoIndexStructEvent event raised by the EventsTestContract contract.
type EventsTestContractNoIndexStructEvent struct {
	A   EventsTestContractAccount
	Raw types.Log // Blockchain specific contextual infos
}

// FilterNoIndexStructEvent is a free log retrieval operation binding the contract event 0xebe3ff7e2071d351bf2e65b4fccd24e3ae99485f02468f1feecf7d64dc044188.
//
// Solidity: event NoIndexStructEvent((string,uint64,uint256) a)
func (_EventsTestContract *EventsTestContractFilterer) FilterNoIndexStructEvent(opts *bind.FilterOpts) (*EventsTestContractNoIndexStructEventIterator, error) {

	logs, sub, err := _EventsTestContract.contract.FilterLogs(opts, "NoIndexStructEvent")
	if err != nil {
		return nil, err
	}
	return &EventsTestContractNoIndexStructEventIterator{contract: _EventsTestContract.contract, event: "NoIndexStructEvent", logs: logs, sub: sub}, nil
}

// WatchNoIndexStructEvent is a free log subscription operation binding the contract event 0xebe3ff7e2071d351bf2e65b4fccd24e3ae99485f02468f1feecf7d64dc044188.
//
// Solidity: event NoIndexStructEvent((string,uint64,uint256) a)
func (_EventsTestContract *EventsTestContractFilterer) WatchNoIndexStructEvent(opts *bind.WatchOpts, sink chan<- *EventsTestContractNoIndexStructEvent) (event.Subscription, error) {

	logs, sub, err := _EventsTestContract.contract.WatchLogs(opts, "NoIndexStructEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventsTestContractNoIndexStructEvent)
				if err := _EventsTestContract.contract.UnpackLog(event, "NoIndexStructEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNoIndexStructEvent is a log parse operation binding the contract event 0xebe3ff7e2071d351bf2e65b4fccd24e3ae99485f02468f1feecf7d64dc044188.
//
// Solidity: event NoIndexStructEvent((string,uint64,uint256) a)
func (_EventsTestContract *EventsTestContractFilterer) ParseNoIndexStructEvent(log types.Log) (*EventsTestContractNoIndexStructEvent, error) {
	event := new(EventsTestContractNoIndexStructEvent)
	if err := _EventsTestContract.contract.UnpackLog(event, "NoIndexStructEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EventsTestContractOneIndexEventIterator is returned from FilterOneIndexEvent and is used to iterate over the raw logs and unpacked data for OneIndexEvent events raised by the EventsTestContract contract.
type EventsTestContractOneIndexEventIterator struct {
	Event *EventsTestContractOneIndexEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EventsTestContractOneIndexEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventsTestContractOneIndexEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EventsTestContractOneIndexEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EventsTestContractOneIndexEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventsTestContractOneIndexEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventsTestContractOneIndexEvent represents a OneIndexEvent event raised by the EventsTestContract contract.
type EventsTestContractOneIndexEvent struct {
	A   common.Hash
	Raw types.Log // Blockchain specific contextual infos
}

// FilterOneIndexEvent is a free log retrieval operation binding the contract event 0x164b5941f5ac984fc056425d6baefece85522d55cef52e1902da35cec5c575ad.
//
// Solidity: event OneIndexEvent(string indexed a)
func (_EventsTestContract *EventsTestContractFilterer) FilterOneIndexEvent(opts *bind.FilterOpts, a []string) (*EventsTestContractOneIndexEventIterator, error) {

	var aRule []interface{}
	for _, aItem := range a {
		aRule = append(aRule, aItem)
	}

	logs, sub, err := _EventsTestContract.contract.FilterLogs(opts, "OneIndexEvent", aRule)
	if err != nil {
		return nil, err
	}
	return &EventsTestContractOneIndexEventIterator{contract: _EventsTestContract.contract, event: "OneIndexEvent", logs: logs, sub: sub}, nil
}

// WatchOneIndexEvent is a free log subscription operation binding the contract event 0x164b5941f5ac984fc056425d6baefece85522d55cef52e1902da35cec5c575ad.
//
// Solidity: event OneIndexEvent(string indexed a)
func (_EventsTestContract *EventsTestContractFilterer) WatchOneIndexEvent(opts *bind.WatchOpts, sink chan<- *EventsTestContractOneIndexEvent, a []string) (event.Subscription, error) {

	var aRule []interface{}
	for _, aItem := range a {
		aRule = append(aRule, aItem)
	}

	logs, sub, err := _EventsTestContract.contract.WatchLogs(opts, "OneIndexEvent", aRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventsTestContractOneIndexEvent)
				if err := _EventsTestContract.contract.UnpackLog(event, "OneIndexEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOneIndexEvent is a log parse operation binding the contract event 0x164b5941f5ac984fc056425d6baefece85522d55cef52e1902da35cec5c575ad.
//
// Solidity: event OneIndexEvent(string indexed a)
func (_EventsTestContract *EventsTestContractFilterer) ParseOneIndexEvent(log types.Log) (*EventsTestContractOneIndexEvent, error) {
	event := new(EventsTestContractOneIndexEvent)
	if err := _EventsTestContract.contract.UnpackLog(event, "OneIndexEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EventsTestContractThreeIndexEventIterator is returned from FilterThreeIndexEvent and is used to iterate over the raw logs and unpacked data for ThreeIndexEvent events raised by the EventsTestContract contract.
type EventsTestContractThreeIndexEventIterator struct {
	Event *EventsTestContractThreeIndexEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EventsTestContractThreeIndexEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventsTestContractThreeIndexEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EventsTestContractThreeIndexEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EventsTestContractThreeIndexEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventsTestContractThreeIndexEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventsTestContractThreeIndexEvent represents a ThreeIndexEvent event raised by the EventsTestContract contract.
type EventsTestContractThreeIndexEvent struct {
	RoundId   *big.Int
	StartedBy common.Address
	StartedAt *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterThreeIndexEvent is a free log retrieval operation binding the contract event 0x5660e8f93f0146f45abcd659e026b75995db50053cbbca4d7f365934ade68bf3.
//
// Solidity: event ThreeIndexEvent(uint256 indexed roundId, address indexed startedBy, uint256 startedAt)
func (_EventsTestContract *EventsTestContractFilterer) FilterThreeIndexEvent(opts *bind.FilterOpts, roundId []*big.Int, startedBy []common.Address) (*EventsTestContractThreeIndexEventIterator, error) {

	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}
	var startedByRule []interface{}
	for _, startedByItem := range startedBy {
		startedByRule = append(startedByRule, startedByItem)
	}

	logs, sub, err := _EventsTestContract.contract.FilterLogs(opts, "ThreeIndexEvent", roundIdRule, startedByRule)
	if err != nil {
		return nil, err
	}
	return &EventsTestContractThreeIndexEventIterator{contract: _EventsTestContract.contract, event: "ThreeIndexEvent", logs: logs, sub: sub}, nil
}

// WatchThreeIndexEvent is a free log subscription operation binding the contract event 0x5660e8f93f0146f45abcd659e026b75995db50053cbbca4d7f365934ade68bf3.
//
// Solidity: event ThreeIndexEvent(uint256 indexed roundId, address indexed startedBy, uint256 startedAt)
func (_EventsTestContract *EventsTestContractFilterer) WatchThreeIndexEvent(opts *bind.WatchOpts, sink chan<- *EventsTestContractThreeIndexEvent, roundId []*big.Int, startedBy []common.Address) (event.Subscription, error) {

	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}
	var startedByRule []interface{}
	for _, startedByItem := range startedBy {
		startedByRule = append(startedByRule, startedByItem)
	}

	logs, sub, err := _EventsTestContract.contract.WatchLogs(opts, "ThreeIndexEvent", roundIdRule, startedByRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventsTestContractThreeIndexEvent)
				if err := _EventsTestContract.contract.UnpackLog(event, "ThreeIndexEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseThreeIndexEvent is a log parse operation binding the contract event 0x5660e8f93f0146f45abcd659e026b75995db50053cbbca4d7f365934ade68bf3.
//
// Solidity: event ThreeIndexEvent(uint256 indexed roundId, address indexed startedBy, uint256 startedAt)
func (_EventsTestContract *EventsTestContractFilterer) ParseThreeIndexEvent(log types.Log) (*EventsTestContractThreeIndexEvent, error) {
	event := new(EventsTestContractThreeIndexEvent)
	if err := _EventsTestContract.contract.UnpackLog(event, "ThreeIndexEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EventsTestContractTwoIndexEventIterator is returned from FilterTwoIndexEvent and is used to iterate over the raw logs and unpacked data for TwoIndexEvent events raised by the EventsTestContract contract.
type EventsTestContractTwoIndexEventIterator struct {
	Event *EventsTestContractTwoIndexEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EventsTestContractTwoIndexEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventsTestContractTwoIndexEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EventsTestContractTwoIndexEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EventsTestContractTwoIndexEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventsTestContractTwoIndexEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventsTestContractTwoIndexEvent represents a TwoIndexEvent event raised by the EventsTestContract contract.
type EventsTestContractTwoIndexEvent struct {
	RoundId   *big.Int
	StartedBy common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTwoIndexEvent is a free log retrieval operation binding the contract event 0x33b47a1cd66813164ec00800d74296f57415217c22505ee380594a712936a0b5.
//
// Solidity: event TwoIndexEvent(uint256 indexed roundId, address indexed startedBy)
func (_EventsTestContract *EventsTestContractFilterer) FilterTwoIndexEvent(opts *bind.FilterOpts, roundId []*big.Int, startedBy []common.Address) (*EventsTestContractTwoIndexEventIterator, error) {

	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}
	var startedByRule []interface{}
	for _, startedByItem := range startedBy {
		startedByRule = append(startedByRule, startedByItem)
	}

	logs, sub, err := _EventsTestContract.contract.FilterLogs(opts, "TwoIndexEvent", roundIdRule, startedByRule)
	if err != nil {
		return nil, err
	}
	return &EventsTestContractTwoIndexEventIterator{contract: _EventsTestContract.contract, event: "TwoIndexEvent", logs: logs, sub: sub}, nil
}

// WatchTwoIndexEvent is a free log subscription operation binding the contract event 0x33b47a1cd66813164ec00800d74296f57415217c22505ee380594a712936a0b5.
//
// Solidity: event TwoIndexEvent(uint256 indexed roundId, address indexed startedBy)
func (_EventsTestContract *EventsTestContractFilterer) WatchTwoIndexEvent(opts *bind.WatchOpts, sink chan<- *EventsTestContractTwoIndexEvent, roundId []*big.Int, startedBy []common.Address) (event.Subscription, error) {

	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}
	var startedByRule []interface{}
	for _, startedByItem := range startedBy {
		startedByRule = append(startedByRule, startedByItem)
	}

	logs, sub, err := _EventsTestContract.contract.WatchLogs(opts, "TwoIndexEvent", roundIdRule, startedByRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventsTestContractTwoIndexEvent)
				if err := _EventsTestContract.contract.UnpackLog(event, "TwoIndexEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTwoIndexEvent is a log parse operation binding the contract event 0x33b47a1cd66813164ec00800d74296f57415217c22505ee380594a712936a0b5.
//
// Solidity: event TwoIndexEvent(uint256 indexed roundId, address indexed startedBy)
func (_EventsTestContract *EventsTestContractFilterer) ParseTwoIndexEvent(log types.Log) (*EventsTestContractTwoIndexEvent, error) {
	event := new(EventsTestContractTwoIndexEvent)
	if err := _EventsTestContract.contract.UnpackLog(event, "TwoIndexEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EventsTestContractValueEventIterator is returned from FilterValueEvent and is used to iterate over the raw logs and unpacked data for ValueEvent events raised by the EventsTestContract contract.
type EventsTestContractValueEventIterator struct {
	Event *EventsTestContractValueEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EventsTestContractValueEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventsTestContractValueEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EventsTestContractValueEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EventsTestContractValueEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventsTestContractValueEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventsTestContractValueEvent represents a ValueEvent event raised by the EventsTestContract contract.
type EventsTestContractValueEvent struct {
	Sender common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterValueEvent is a free log retrieval operation binding the contract event 0x834ba0b35948e5b0a41be62c5c77104608331d60ee912667365f453865a6dbaf.
//
// Solidity: event ValueEvent(address indexed sender, uint256 value)
func (_EventsTestContract *EventsTestContractFilterer) FilterValueEvent(opts *bind.FilterOpts, sender []common.Address) (*EventsTestContractValueEventIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _EventsTestContract.contract.FilterLogs(opts, "ValueEvent", senderRule)
	if err != nil {
		return nil, err
	}
	return &EventsTestContractValueEventIterator{contract: _EventsTestContract.contract, event: "ValueEvent", logs: logs, sub: sub}, nil
}

// WatchValueEvent is a free log subscription operation binding the contract event 0x834ba0b35948e5b0a41be62c5c77104608331d60ee912667365f453865a6dbaf.
//
// Solidity: event ValueEvent(address indexed sender, uint256 value)
func (_EventsTestContract *EventsTestContractFilterer) WatchValueEvent(opts *bind.WatchOpts, sink chan<- *EventsTestContractValueEvent, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _EventsTestContract.contract.WatchLogs(opts, "ValueEvent", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventsTestContractValueEvent)
				if err := _EventsTestContract.contract.UnpackLog(event, "ValueEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValueEvent is a log parse operation binding the contract event 0x834ba0b35948e5b0a41be62c5c77104608331d60ee912667365f453865a6dbaf.
//
// Solidity: event ValueEvent(address indexed sender, uint256 value)
func (_EventsTestContract *EventsTestContractFilterer) ParseValueEvent(log types.Log) (*EventsTestContractValueEvent, error) {
	event := new(EventsTestContractValueEvent)
	if err := _EventsTestContract.contract.UnpackLog(event, "ValueEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

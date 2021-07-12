// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package another_events_test_contract

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

// AnotherEventsTestContractAccount is an auto generated low-level Go binding around an user-defined struct.
type AnotherEventsTestContractAccount struct {
	Name       string
	Balance    uint64
	Dailylimit *big.Int
}

// AnotherEventsTestContractABI is the input ABI used to generate the binding from.
const AnotherEventsTestContractABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"NoIndexEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"balance\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"dailylimit\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structAnotherEventsTestContract.Account\",\"name\":\"a\",\"type\":\"tuple\"}],\"name\":\"NoIndexStructEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"a\",\"type\":\"string\"}],\"name\":\"OneIndexEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"startedBy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"}],\"name\":\"ThreeIndexEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"startedBy\",\"type\":\"address\"}],\"name\":\"TwoIndexEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"ValueEvent\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"fireNoIndexEvent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fireNoIndexStructEvent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fireOneIndexEvent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fireThreeIndexEvent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fireTwoIndexEvent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"}],\"name\":\"set\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// AnotherEventsTestContractBin is the compiled bytecode used for deploying new contracts.
var AnotherEventsTestContractBin = "0x608060405234801561001057600080fd5b5061069a806100206000396000f3fe608060405234801561001057600080fd5b506004361061007c5760003560e01c80636d28a01f1161005b5780636d28a01f146100c55780636d4ce63c146100cf57806395a777bf146100ed578063eac9dc87146100f75761007c565b80625fd35114610081578063577dd2941461008b57806360fe47b114610095575b600080fd5b610089610101565b005b610093610195565b005b6100af60048036038101906100aa9190610343565b6101dc565b6040516100bc91906104d4565b60405180910390f35b6100cd61024a565b005b6100d761028c565b6040516100e491906104d4565b60405180910390f35b6100f56102a2565b005b6100ff6102db565b005b7febe3ff7e2071d351bf2e65b4fccd24e3ae99485f02468f1feecf7d64dc04418860405180606001604052806040518060400160405280600481526020017f4a6f686e000000000000000000000000000000000000000000000000000000008152508152602001600567ffffffffffffffff168152602001600a81525060405161018b91906104b2565b60405180910390a1565b3373ffffffffffffffffffffffffffffffffffffffff1660017f33b47a1cd66813164ec00800d74296f57415217c22505ee380594a712936a0b560405160405180910390a3565b6000816000819055503373ffffffffffffffffffffffffffffffffffffffff167f834ba0b35948e5b0a41be62c5c77104608331d60ee912667365f453865a6dbaf8360405161022b91906104d4565b60405180910390a26005826102409190610516565b9150819050919050565b60405161025690610467565b60405180910390207f164b5941f5ac984fc056425d6baefece85522d55cef52e1902da35cec5c575ad60405160405180910390a2565b6000600a60005461029d9190610516565b905090565b7f33bc9bae48dbe1e057f264b3fc6a1dacdcceacb3ba28d937231c70e068a02f1c336040516102d1919061047c565b60405180910390a1565b3373ffffffffffffffffffffffffffffffffffffffff1660017f5660e8f93f0146f45abcd659e026b75995db50053cbbca4d7f365934ade68bf360036040516103249190610497565b60405180910390a3565b60008135905061033d8161064d565b92915050565b6000602082840312156103595761035861060e565b5b60006103678482850161032e565b91505092915050565b6103798161054a565b82525050565b6103888161059a565b82525050565b6000610399826104ef565b6103a381856104fa565b93506103b38185602086016105ac565b6103bc81610613565b840191505092915050565b60006103d460118361050b565b91506103df82610624565b601182019050919050565b60006060830160008301518482036000860152610407828261038e565b915050602083015161041c6020860182610458565b50604083015161042f604086018261043a565b508091505092915050565b6104438161057c565b82525050565b6104528161057c565b82525050565b61046181610586565b82525050565b6000610472826103c7565b9150819050919050565b60006020820190506104916000830184610370565b92915050565b60006020820190506104ac600083018461037f565b92915050565b600060208201905081810360008301526104cc81846103ea565b905092915050565b60006020820190506104e96000830184610449565b92915050565b600081519050919050565b600082825260208201905092915050565b600081905092915050565b60006105218261057c565b915061052c8361057c565b92508282101561053f5761053e6105df565b5b828203905092915050565b60006105558261055c565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600067ffffffffffffffff82169050919050565b60006105a58261057c565b9050919050565b60005b838110156105ca5780820151818401526020810190506105af565b838111156105d9576000848401525b50505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600080fd5b6000601f19601f8301169050919050565b7f6576656e745f646174615f737472696e67000000000000000000000000000000600082015250565b6106568161057c565b811461066157600080fd5b5056fea2646970667358221220823e13d3d4b4ac85cf93a6776d476a927170acf4e81026e7b2638c8f41c6993d64736f6c63430008060033"

// DeployAnotherEventsTestContract deploys a new Ethereum contract, binding an instance of AnotherEventsTestContract to it.
func DeployAnotherEventsTestContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AnotherEventsTestContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AnotherEventsTestContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AnotherEventsTestContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AnotherEventsTestContract{AnotherEventsTestContractCaller: AnotherEventsTestContractCaller{contract: contract}, AnotherEventsTestContractTransactor: AnotherEventsTestContractTransactor{contract: contract}, AnotherEventsTestContractFilterer: AnotherEventsTestContractFilterer{contract: contract}}, nil
}

// AnotherEventsTestContract is an auto generated Go binding around an Ethereum contract.
type AnotherEventsTestContract struct {
	AnotherEventsTestContractCaller     // Read-only binding to the contract
	AnotherEventsTestContractTransactor // Write-only binding to the contract
	AnotherEventsTestContractFilterer   // Log filterer for contract events
}

// AnotherEventsTestContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type AnotherEventsTestContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AnotherEventsTestContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AnotherEventsTestContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AnotherEventsTestContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AnotherEventsTestContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AnotherEventsTestContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AnotherEventsTestContractSession struct {
	Contract     *AnotherEventsTestContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// AnotherEventsTestContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AnotherEventsTestContractCallerSession struct {
	Contract *AnotherEventsTestContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// AnotherEventsTestContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AnotherEventsTestContractTransactorSession struct {
	Contract     *AnotherEventsTestContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// AnotherEventsTestContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type AnotherEventsTestContractRaw struct {
	Contract *AnotherEventsTestContract // Generic contract binding to access the raw methods on
}

// AnotherEventsTestContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AnotherEventsTestContractCallerRaw struct {
	Contract *AnotherEventsTestContractCaller // Generic read-only contract binding to access the raw methods on
}

// AnotherEventsTestContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AnotherEventsTestContractTransactorRaw struct {
	Contract *AnotherEventsTestContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAnotherEventsTestContract creates a new instance of AnotherEventsTestContract, bound to a specific deployed contract.
func NewAnotherEventsTestContract(address common.Address, backend bind.ContractBackend) (*AnotherEventsTestContract, error) {
	contract, err := bindAnotherEventsTestContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AnotherEventsTestContract{AnotherEventsTestContractCaller: AnotherEventsTestContractCaller{contract: contract}, AnotherEventsTestContractTransactor: AnotherEventsTestContractTransactor{contract: contract}, AnotherEventsTestContractFilterer: AnotherEventsTestContractFilterer{contract: contract}}, nil
}

// NewAnotherEventsTestContractCaller creates a new read-only instance of AnotherEventsTestContract, bound to a specific deployed contract.
func NewAnotherEventsTestContractCaller(address common.Address, caller bind.ContractCaller) (*AnotherEventsTestContractCaller, error) {
	contract, err := bindAnotherEventsTestContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AnotherEventsTestContractCaller{contract: contract}, nil
}

// NewAnotherEventsTestContractTransactor creates a new write-only instance of AnotherEventsTestContract, bound to a specific deployed contract.
func NewAnotherEventsTestContractTransactor(address common.Address, transactor bind.ContractTransactor) (*AnotherEventsTestContractTransactor, error) {
	contract, err := bindAnotherEventsTestContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AnotherEventsTestContractTransactor{contract: contract}, nil
}

// NewAnotherEventsTestContractFilterer creates a new log filterer instance of AnotherEventsTestContract, bound to a specific deployed contract.
func NewAnotherEventsTestContractFilterer(address common.Address, filterer bind.ContractFilterer) (*AnotherEventsTestContractFilterer, error) {
	contract, err := bindAnotherEventsTestContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AnotherEventsTestContractFilterer{contract: contract}, nil
}

// bindAnotherEventsTestContract binds a generic wrapper to an already deployed contract.
func bindAnotherEventsTestContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AnotherEventsTestContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AnotherEventsTestContract *AnotherEventsTestContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AnotherEventsTestContract.Contract.AnotherEventsTestContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AnotherEventsTestContract *AnotherEventsTestContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AnotherEventsTestContract.Contract.AnotherEventsTestContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AnotherEventsTestContract *AnotherEventsTestContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AnotherEventsTestContract.Contract.AnotherEventsTestContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AnotherEventsTestContract *AnotherEventsTestContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AnotherEventsTestContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AnotherEventsTestContract *AnotherEventsTestContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AnotherEventsTestContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AnotherEventsTestContract *AnotherEventsTestContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AnotherEventsTestContract.Contract.contract.Transact(opts, method, params...)
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns(uint256)
func (_AnotherEventsTestContract *AnotherEventsTestContractCaller) Get(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AnotherEventsTestContract.contract.Call(opts, &out, "get")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns(uint256)
func (_AnotherEventsTestContract *AnotherEventsTestContractSession) Get() (*big.Int, error) {
	return _AnotherEventsTestContract.Contract.Get(&_AnotherEventsTestContract.CallOpts)
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns(uint256)
func (_AnotherEventsTestContract *AnotherEventsTestContractCallerSession) Get() (*big.Int, error) {
	return _AnotherEventsTestContract.Contract.Get(&_AnotherEventsTestContract.CallOpts)
}

// FireNoIndexEvent is a paid mutator transaction binding the contract method 0x95a777bf.
//
// Solidity: function fireNoIndexEvent() returns()
func (_AnotherEventsTestContract *AnotherEventsTestContractTransactor) FireNoIndexEvent(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AnotherEventsTestContract.contract.Transact(opts, "fireNoIndexEvent")
}

// FireNoIndexEvent is a paid mutator transaction binding the contract method 0x95a777bf.
//
// Solidity: function fireNoIndexEvent() returns()
func (_AnotherEventsTestContract *AnotherEventsTestContractSession) FireNoIndexEvent() (*types.Transaction, error) {
	return _AnotherEventsTestContract.Contract.FireNoIndexEvent(&_AnotherEventsTestContract.TransactOpts)
}

// FireNoIndexEvent is a paid mutator transaction binding the contract method 0x95a777bf.
//
// Solidity: function fireNoIndexEvent() returns()
func (_AnotherEventsTestContract *AnotherEventsTestContractTransactorSession) FireNoIndexEvent() (*types.Transaction, error) {
	return _AnotherEventsTestContract.Contract.FireNoIndexEvent(&_AnotherEventsTestContract.TransactOpts)
}

// FireNoIndexStructEvent is a paid mutator transaction binding the contract method 0x005fd351.
//
// Solidity: function fireNoIndexStructEvent() returns()
func (_AnotherEventsTestContract *AnotherEventsTestContractTransactor) FireNoIndexStructEvent(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AnotherEventsTestContract.contract.Transact(opts, "fireNoIndexStructEvent")
}

// FireNoIndexStructEvent is a paid mutator transaction binding the contract method 0x005fd351.
//
// Solidity: function fireNoIndexStructEvent() returns()
func (_AnotherEventsTestContract *AnotherEventsTestContractSession) FireNoIndexStructEvent() (*types.Transaction, error) {
	return _AnotherEventsTestContract.Contract.FireNoIndexStructEvent(&_AnotherEventsTestContract.TransactOpts)
}

// FireNoIndexStructEvent is a paid mutator transaction binding the contract method 0x005fd351.
//
// Solidity: function fireNoIndexStructEvent() returns()
func (_AnotherEventsTestContract *AnotherEventsTestContractTransactorSession) FireNoIndexStructEvent() (*types.Transaction, error) {
	return _AnotherEventsTestContract.Contract.FireNoIndexStructEvent(&_AnotherEventsTestContract.TransactOpts)
}

// FireOneIndexEvent is a paid mutator transaction binding the contract method 0x6d28a01f.
//
// Solidity: function fireOneIndexEvent() returns()
func (_AnotherEventsTestContract *AnotherEventsTestContractTransactor) FireOneIndexEvent(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AnotherEventsTestContract.contract.Transact(opts, "fireOneIndexEvent")
}

// FireOneIndexEvent is a paid mutator transaction binding the contract method 0x6d28a01f.
//
// Solidity: function fireOneIndexEvent() returns()
func (_AnotherEventsTestContract *AnotherEventsTestContractSession) FireOneIndexEvent() (*types.Transaction, error) {
	return _AnotherEventsTestContract.Contract.FireOneIndexEvent(&_AnotherEventsTestContract.TransactOpts)
}

// FireOneIndexEvent is a paid mutator transaction binding the contract method 0x6d28a01f.
//
// Solidity: function fireOneIndexEvent() returns()
func (_AnotherEventsTestContract *AnotherEventsTestContractTransactorSession) FireOneIndexEvent() (*types.Transaction, error) {
	return _AnotherEventsTestContract.Contract.FireOneIndexEvent(&_AnotherEventsTestContract.TransactOpts)
}

// FireThreeIndexEvent is a paid mutator transaction binding the contract method 0xeac9dc87.
//
// Solidity: function fireThreeIndexEvent() returns()
func (_AnotherEventsTestContract *AnotherEventsTestContractTransactor) FireThreeIndexEvent(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AnotherEventsTestContract.contract.Transact(opts, "fireThreeIndexEvent")
}

// FireThreeIndexEvent is a paid mutator transaction binding the contract method 0xeac9dc87.
//
// Solidity: function fireThreeIndexEvent() returns()
func (_AnotherEventsTestContract *AnotherEventsTestContractSession) FireThreeIndexEvent() (*types.Transaction, error) {
	return _AnotherEventsTestContract.Contract.FireThreeIndexEvent(&_AnotherEventsTestContract.TransactOpts)
}

// FireThreeIndexEvent is a paid mutator transaction binding the contract method 0xeac9dc87.
//
// Solidity: function fireThreeIndexEvent() returns()
func (_AnotherEventsTestContract *AnotherEventsTestContractTransactorSession) FireThreeIndexEvent() (*types.Transaction, error) {
	return _AnotherEventsTestContract.Contract.FireThreeIndexEvent(&_AnotherEventsTestContract.TransactOpts)
}

// FireTwoIndexEvent is a paid mutator transaction binding the contract method 0x577dd294.
//
// Solidity: function fireTwoIndexEvent() returns()
func (_AnotherEventsTestContract *AnotherEventsTestContractTransactor) FireTwoIndexEvent(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AnotherEventsTestContract.contract.Transact(opts, "fireTwoIndexEvent")
}

// FireTwoIndexEvent is a paid mutator transaction binding the contract method 0x577dd294.
//
// Solidity: function fireTwoIndexEvent() returns()
func (_AnotherEventsTestContract *AnotherEventsTestContractSession) FireTwoIndexEvent() (*types.Transaction, error) {
	return _AnotherEventsTestContract.Contract.FireTwoIndexEvent(&_AnotherEventsTestContract.TransactOpts)
}

// FireTwoIndexEvent is a paid mutator transaction binding the contract method 0x577dd294.
//
// Solidity: function fireTwoIndexEvent() returns()
func (_AnotherEventsTestContract *AnotherEventsTestContractTransactorSession) FireTwoIndexEvent() (*types.Transaction, error) {
	return _AnotherEventsTestContract.Contract.FireTwoIndexEvent(&_AnotherEventsTestContract.TransactOpts)
}

// Set is a paid mutator transaction binding the contract method 0x60fe47b1.
//
// Solidity: function set(uint256 x) returns(uint256)
func (_AnotherEventsTestContract *AnotherEventsTestContractTransactor) Set(opts *bind.TransactOpts, x *big.Int) (*types.Transaction, error) {
	return _AnotherEventsTestContract.contract.Transact(opts, "set", x)
}

// Set is a paid mutator transaction binding the contract method 0x60fe47b1.
//
// Solidity: function set(uint256 x) returns(uint256)
func (_AnotherEventsTestContract *AnotherEventsTestContractSession) Set(x *big.Int) (*types.Transaction, error) {
	return _AnotherEventsTestContract.Contract.Set(&_AnotherEventsTestContract.TransactOpts, x)
}

// Set is a paid mutator transaction binding the contract method 0x60fe47b1.
//
// Solidity: function set(uint256 x) returns(uint256)
func (_AnotherEventsTestContract *AnotherEventsTestContractTransactorSession) Set(x *big.Int) (*types.Transaction, error) {
	return _AnotherEventsTestContract.Contract.Set(&_AnotherEventsTestContract.TransactOpts, x)
}

// AnotherEventsTestContractNoIndexEventIterator is returned from FilterNoIndexEvent and is used to iterate over the raw logs and unpacked data for NoIndexEvent events raised by the AnotherEventsTestContract contract.
type AnotherEventsTestContractNoIndexEventIterator struct {
	Event *AnotherEventsTestContractNoIndexEvent // Event containing the contract specifics and raw log

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
func (it *AnotherEventsTestContractNoIndexEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AnotherEventsTestContractNoIndexEvent)
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
		it.Event = new(AnotherEventsTestContractNoIndexEvent)
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
func (it *AnotherEventsTestContractNoIndexEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AnotherEventsTestContractNoIndexEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AnotherEventsTestContractNoIndexEvent represents a NoIndexEvent event raised by the AnotherEventsTestContract contract.
type AnotherEventsTestContractNoIndexEvent struct {
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNoIndexEvent is a free log retrieval operation binding the contract event 0x33bc9bae48dbe1e057f264b3fc6a1dacdcceacb3ba28d937231c70e068a02f1c.
//
// Solidity: event NoIndexEvent(address sender)
func (_AnotherEventsTestContract *AnotherEventsTestContractFilterer) FilterNoIndexEvent(opts *bind.FilterOpts) (*AnotherEventsTestContractNoIndexEventIterator, error) {

	logs, sub, err := _AnotherEventsTestContract.contract.FilterLogs(opts, "NoIndexEvent")
	if err != nil {
		return nil, err
	}
	return &AnotherEventsTestContractNoIndexEventIterator{contract: _AnotherEventsTestContract.contract, event: "NoIndexEvent", logs: logs, sub: sub}, nil
}

// WatchNoIndexEvent is a free log subscription operation binding the contract event 0x33bc9bae48dbe1e057f264b3fc6a1dacdcceacb3ba28d937231c70e068a02f1c.
//
// Solidity: event NoIndexEvent(address sender)
func (_AnotherEventsTestContract *AnotherEventsTestContractFilterer) WatchNoIndexEvent(opts *bind.WatchOpts, sink chan<- *AnotherEventsTestContractNoIndexEvent) (event.Subscription, error) {

	logs, sub, err := _AnotherEventsTestContract.contract.WatchLogs(opts, "NoIndexEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AnotherEventsTestContractNoIndexEvent)
				if err := _AnotherEventsTestContract.contract.UnpackLog(event, "NoIndexEvent", log); err != nil {
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
func (_AnotherEventsTestContract *AnotherEventsTestContractFilterer) ParseNoIndexEvent(log types.Log) (*AnotherEventsTestContractNoIndexEvent, error) {
	event := new(AnotherEventsTestContractNoIndexEvent)
	if err := _AnotherEventsTestContract.contract.UnpackLog(event, "NoIndexEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AnotherEventsTestContractNoIndexStructEventIterator is returned from FilterNoIndexStructEvent and is used to iterate over the raw logs and unpacked data for NoIndexStructEvent events raised by the AnotherEventsTestContract contract.
type AnotherEventsTestContractNoIndexStructEventIterator struct {
	Event *AnotherEventsTestContractNoIndexStructEvent // Event containing the contract specifics and raw log

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
func (it *AnotherEventsTestContractNoIndexStructEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AnotherEventsTestContractNoIndexStructEvent)
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
		it.Event = new(AnotherEventsTestContractNoIndexStructEvent)
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
func (it *AnotherEventsTestContractNoIndexStructEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AnotherEventsTestContractNoIndexStructEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AnotherEventsTestContractNoIndexStructEvent represents a NoIndexStructEvent event raised by the AnotherEventsTestContract contract.
type AnotherEventsTestContractNoIndexStructEvent struct {
	A   AnotherEventsTestContractAccount
	Raw types.Log // Blockchain specific contextual infos
}

// FilterNoIndexStructEvent is a free log retrieval operation binding the contract event 0xebe3ff7e2071d351bf2e65b4fccd24e3ae99485f02468f1feecf7d64dc044188.
//
// Solidity: event NoIndexStructEvent((string,uint64,uint256) a)
func (_AnotherEventsTestContract *AnotherEventsTestContractFilterer) FilterNoIndexStructEvent(opts *bind.FilterOpts) (*AnotherEventsTestContractNoIndexStructEventIterator, error) {

	logs, sub, err := _AnotherEventsTestContract.contract.FilterLogs(opts, "NoIndexStructEvent")
	if err != nil {
		return nil, err
	}
	return &AnotherEventsTestContractNoIndexStructEventIterator{contract: _AnotherEventsTestContract.contract, event: "NoIndexStructEvent", logs: logs, sub: sub}, nil
}

// WatchNoIndexStructEvent is a free log subscription operation binding the contract event 0xebe3ff7e2071d351bf2e65b4fccd24e3ae99485f02468f1feecf7d64dc044188.
//
// Solidity: event NoIndexStructEvent((string,uint64,uint256) a)
func (_AnotherEventsTestContract *AnotherEventsTestContractFilterer) WatchNoIndexStructEvent(opts *bind.WatchOpts, sink chan<- *AnotherEventsTestContractNoIndexStructEvent) (event.Subscription, error) {

	logs, sub, err := _AnotherEventsTestContract.contract.WatchLogs(opts, "NoIndexStructEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AnotherEventsTestContractNoIndexStructEvent)
				if err := _AnotherEventsTestContract.contract.UnpackLog(event, "NoIndexStructEvent", log); err != nil {
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
func (_AnotherEventsTestContract *AnotherEventsTestContractFilterer) ParseNoIndexStructEvent(log types.Log) (*AnotherEventsTestContractNoIndexStructEvent, error) {
	event := new(AnotherEventsTestContractNoIndexStructEvent)
	if err := _AnotherEventsTestContract.contract.UnpackLog(event, "NoIndexStructEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AnotherEventsTestContractOneIndexEventIterator is returned from FilterOneIndexEvent and is used to iterate over the raw logs and unpacked data for OneIndexEvent events raised by the AnotherEventsTestContract contract.
type AnotherEventsTestContractOneIndexEventIterator struct {
	Event *AnotherEventsTestContractOneIndexEvent // Event containing the contract specifics and raw log

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
func (it *AnotherEventsTestContractOneIndexEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AnotherEventsTestContractOneIndexEvent)
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
		it.Event = new(AnotherEventsTestContractOneIndexEvent)
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
func (it *AnotherEventsTestContractOneIndexEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AnotherEventsTestContractOneIndexEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AnotherEventsTestContractOneIndexEvent represents a OneIndexEvent event raised by the AnotherEventsTestContract contract.
type AnotherEventsTestContractOneIndexEvent struct {
	A   common.Hash
	Raw types.Log // Blockchain specific contextual infos
}

// FilterOneIndexEvent is a free log retrieval operation binding the contract event 0x164b5941f5ac984fc056425d6baefece85522d55cef52e1902da35cec5c575ad.
//
// Solidity: event OneIndexEvent(string indexed a)
func (_AnotherEventsTestContract *AnotherEventsTestContractFilterer) FilterOneIndexEvent(opts *bind.FilterOpts, a []string) (*AnotherEventsTestContractOneIndexEventIterator, error) {

	var aRule []interface{}
	for _, aItem := range a {
		aRule = append(aRule, aItem)
	}

	logs, sub, err := _AnotherEventsTestContract.contract.FilterLogs(opts, "OneIndexEvent", aRule)
	if err != nil {
		return nil, err
	}
	return &AnotherEventsTestContractOneIndexEventIterator{contract: _AnotherEventsTestContract.contract, event: "OneIndexEvent", logs: logs, sub: sub}, nil
}

// WatchOneIndexEvent is a free log subscription operation binding the contract event 0x164b5941f5ac984fc056425d6baefece85522d55cef52e1902da35cec5c575ad.
//
// Solidity: event OneIndexEvent(string indexed a)
func (_AnotherEventsTestContract *AnotherEventsTestContractFilterer) WatchOneIndexEvent(opts *bind.WatchOpts, sink chan<- *AnotherEventsTestContractOneIndexEvent, a []string) (event.Subscription, error) {

	var aRule []interface{}
	for _, aItem := range a {
		aRule = append(aRule, aItem)
	}

	logs, sub, err := _AnotherEventsTestContract.contract.WatchLogs(opts, "OneIndexEvent", aRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AnotherEventsTestContractOneIndexEvent)
				if err := _AnotherEventsTestContract.contract.UnpackLog(event, "OneIndexEvent", log); err != nil {
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
func (_AnotherEventsTestContract *AnotherEventsTestContractFilterer) ParseOneIndexEvent(log types.Log) (*AnotherEventsTestContractOneIndexEvent, error) {
	event := new(AnotherEventsTestContractOneIndexEvent)
	if err := _AnotherEventsTestContract.contract.UnpackLog(event, "OneIndexEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AnotherEventsTestContractThreeIndexEventIterator is returned from FilterThreeIndexEvent and is used to iterate over the raw logs and unpacked data for ThreeIndexEvent events raised by the AnotherEventsTestContract contract.
type AnotherEventsTestContractThreeIndexEventIterator struct {
	Event *AnotherEventsTestContractThreeIndexEvent // Event containing the contract specifics and raw log

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
func (it *AnotherEventsTestContractThreeIndexEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AnotherEventsTestContractThreeIndexEvent)
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
		it.Event = new(AnotherEventsTestContractThreeIndexEvent)
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
func (it *AnotherEventsTestContractThreeIndexEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AnotherEventsTestContractThreeIndexEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AnotherEventsTestContractThreeIndexEvent represents a ThreeIndexEvent event raised by the AnotherEventsTestContract contract.
type AnotherEventsTestContractThreeIndexEvent struct {
	RoundId   *big.Int
	StartedBy common.Address
	StartedAt *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterThreeIndexEvent is a free log retrieval operation binding the contract event 0x5660e8f93f0146f45abcd659e026b75995db50053cbbca4d7f365934ade68bf3.
//
// Solidity: event ThreeIndexEvent(uint256 indexed roundId, address indexed startedBy, uint256 startedAt)
func (_AnotherEventsTestContract *AnotherEventsTestContractFilterer) FilterThreeIndexEvent(opts *bind.FilterOpts, roundId []*big.Int, startedBy []common.Address) (*AnotherEventsTestContractThreeIndexEventIterator, error) {

	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}
	var startedByRule []interface{}
	for _, startedByItem := range startedBy {
		startedByRule = append(startedByRule, startedByItem)
	}

	logs, sub, err := _AnotherEventsTestContract.contract.FilterLogs(opts, "ThreeIndexEvent", roundIdRule, startedByRule)
	if err != nil {
		return nil, err
	}
	return &AnotherEventsTestContractThreeIndexEventIterator{contract: _AnotherEventsTestContract.contract, event: "ThreeIndexEvent", logs: logs, sub: sub}, nil
}

// WatchThreeIndexEvent is a free log subscription operation binding the contract event 0x5660e8f93f0146f45abcd659e026b75995db50053cbbca4d7f365934ade68bf3.
//
// Solidity: event ThreeIndexEvent(uint256 indexed roundId, address indexed startedBy, uint256 startedAt)
func (_AnotherEventsTestContract *AnotherEventsTestContractFilterer) WatchThreeIndexEvent(opts *bind.WatchOpts, sink chan<- *AnotherEventsTestContractThreeIndexEvent, roundId []*big.Int, startedBy []common.Address) (event.Subscription, error) {

	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}
	var startedByRule []interface{}
	for _, startedByItem := range startedBy {
		startedByRule = append(startedByRule, startedByItem)
	}

	logs, sub, err := _AnotherEventsTestContract.contract.WatchLogs(opts, "ThreeIndexEvent", roundIdRule, startedByRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AnotherEventsTestContractThreeIndexEvent)
				if err := _AnotherEventsTestContract.contract.UnpackLog(event, "ThreeIndexEvent", log); err != nil {
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
func (_AnotherEventsTestContract *AnotherEventsTestContractFilterer) ParseThreeIndexEvent(log types.Log) (*AnotherEventsTestContractThreeIndexEvent, error) {
	event := new(AnotherEventsTestContractThreeIndexEvent)
	if err := _AnotherEventsTestContract.contract.UnpackLog(event, "ThreeIndexEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AnotherEventsTestContractTwoIndexEventIterator is returned from FilterTwoIndexEvent and is used to iterate over the raw logs and unpacked data for TwoIndexEvent events raised by the AnotherEventsTestContract contract.
type AnotherEventsTestContractTwoIndexEventIterator struct {
	Event *AnotherEventsTestContractTwoIndexEvent // Event containing the contract specifics and raw log

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
func (it *AnotherEventsTestContractTwoIndexEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AnotherEventsTestContractTwoIndexEvent)
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
		it.Event = new(AnotherEventsTestContractTwoIndexEvent)
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
func (it *AnotherEventsTestContractTwoIndexEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AnotherEventsTestContractTwoIndexEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AnotherEventsTestContractTwoIndexEvent represents a TwoIndexEvent event raised by the AnotherEventsTestContract contract.
type AnotherEventsTestContractTwoIndexEvent struct {
	RoundId   *big.Int
	StartedBy common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTwoIndexEvent is a free log retrieval operation binding the contract event 0x33b47a1cd66813164ec00800d74296f57415217c22505ee380594a712936a0b5.
//
// Solidity: event TwoIndexEvent(uint256 indexed roundId, address indexed startedBy)
func (_AnotherEventsTestContract *AnotherEventsTestContractFilterer) FilterTwoIndexEvent(opts *bind.FilterOpts, roundId []*big.Int, startedBy []common.Address) (*AnotherEventsTestContractTwoIndexEventIterator, error) {

	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}
	var startedByRule []interface{}
	for _, startedByItem := range startedBy {
		startedByRule = append(startedByRule, startedByItem)
	}

	logs, sub, err := _AnotherEventsTestContract.contract.FilterLogs(opts, "TwoIndexEvent", roundIdRule, startedByRule)
	if err != nil {
		return nil, err
	}
	return &AnotherEventsTestContractTwoIndexEventIterator{contract: _AnotherEventsTestContract.contract, event: "TwoIndexEvent", logs: logs, sub: sub}, nil
}

// WatchTwoIndexEvent is a free log subscription operation binding the contract event 0x33b47a1cd66813164ec00800d74296f57415217c22505ee380594a712936a0b5.
//
// Solidity: event TwoIndexEvent(uint256 indexed roundId, address indexed startedBy)
func (_AnotherEventsTestContract *AnotherEventsTestContractFilterer) WatchTwoIndexEvent(opts *bind.WatchOpts, sink chan<- *AnotherEventsTestContractTwoIndexEvent, roundId []*big.Int, startedBy []common.Address) (event.Subscription, error) {

	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}
	var startedByRule []interface{}
	for _, startedByItem := range startedBy {
		startedByRule = append(startedByRule, startedByItem)
	}

	logs, sub, err := _AnotherEventsTestContract.contract.WatchLogs(opts, "TwoIndexEvent", roundIdRule, startedByRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AnotherEventsTestContractTwoIndexEvent)
				if err := _AnotherEventsTestContract.contract.UnpackLog(event, "TwoIndexEvent", log); err != nil {
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
func (_AnotherEventsTestContract *AnotherEventsTestContractFilterer) ParseTwoIndexEvent(log types.Log) (*AnotherEventsTestContractTwoIndexEvent, error) {
	event := new(AnotherEventsTestContractTwoIndexEvent)
	if err := _AnotherEventsTestContract.contract.UnpackLog(event, "TwoIndexEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AnotherEventsTestContractValueEventIterator is returned from FilterValueEvent and is used to iterate over the raw logs and unpacked data for ValueEvent events raised by the AnotherEventsTestContract contract.
type AnotherEventsTestContractValueEventIterator struct {
	Event *AnotherEventsTestContractValueEvent // Event containing the contract specifics and raw log

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
func (it *AnotherEventsTestContractValueEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AnotherEventsTestContractValueEvent)
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
		it.Event = new(AnotherEventsTestContractValueEvent)
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
func (it *AnotherEventsTestContractValueEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AnotherEventsTestContractValueEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AnotherEventsTestContractValueEvent represents a ValueEvent event raised by the AnotherEventsTestContract contract.
type AnotherEventsTestContractValueEvent struct {
	Sender common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterValueEvent is a free log retrieval operation binding the contract event 0x834ba0b35948e5b0a41be62c5c77104608331d60ee912667365f453865a6dbaf.
//
// Solidity: event ValueEvent(address indexed sender, uint256 value)
func (_AnotherEventsTestContract *AnotherEventsTestContractFilterer) FilterValueEvent(opts *bind.FilterOpts, sender []common.Address) (*AnotherEventsTestContractValueEventIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AnotherEventsTestContract.contract.FilterLogs(opts, "ValueEvent", senderRule)
	if err != nil {
		return nil, err
	}
	return &AnotherEventsTestContractValueEventIterator{contract: _AnotherEventsTestContract.contract, event: "ValueEvent", logs: logs, sub: sub}, nil
}

// WatchValueEvent is a free log subscription operation binding the contract event 0x834ba0b35948e5b0a41be62c5c77104608331d60ee912667365f453865a6dbaf.
//
// Solidity: event ValueEvent(address indexed sender, uint256 value)
func (_AnotherEventsTestContract *AnotherEventsTestContractFilterer) WatchValueEvent(opts *bind.WatchOpts, sink chan<- *AnotherEventsTestContractValueEvent, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AnotherEventsTestContract.contract.WatchLogs(opts, "ValueEvent", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AnotherEventsTestContractValueEvent)
				if err := _AnotherEventsTestContract.contract.UnpackLog(event, "ValueEvent", log); err != nil {
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
func (_AnotherEventsTestContract *AnotherEventsTestContractFilterer) ParseValueEvent(log types.Log) (*AnotherEventsTestContractValueEvent, error) {
	event := new(AnotherEventsTestContractValueEvent)
	if err := _AnotherEventsTestContract.contract.UnpackLog(event, "ValueEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

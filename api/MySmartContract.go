// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package api

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ApiMetaData contains all meta data concerning the Api contract.
var ApiMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"Balance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"Deposite\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amt\",\"type\":\"uint256\"}],\"name\":\"Withdrawl\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60806040526000805534801561001457600080fd5b50600180546001600160a01b0319163317905561002f610034565b610072565b34600080828254610045919061004c565b9091555050565b6000821982111561006d57634e487b7160e01b600052601160045260246000fd5b500190565b6101e5806100816000396000f3fe6080604052600436106100435760003560e01c80630ef678871461005757806342002bc81461007b578063e615c1a01461009b578063f851a440146100bb57600080fd5b36610052576100506100f3565b005b600080fd5b34801561006357600080fd5b506000545b6040519081526020015b60405180910390f35b34801561008757600080fd5b50610068610096366004610151565b61010b565b3480156100a757600080fd5b506100506100b6366004610151565b610126565b3480156100c757600080fd5b506001546100db906001600160a01b031681565b6040516001600160a01b039091168152602001610072565b346000808282546101049190610180565b9091555050565b60008160005461011b9190610180565b600081905592915050565b6001546001600160a01b0316331461013d57600080fd5b8060005461014b9190610198565b60005550565b60006020828403121561016357600080fd5b5035919050565b634e487b7160e01b600052601160045260246000fd5b600082198211156101935761019361016a565b500190565b6000828210156101aa576101aa61016a565b50039056fea26469706673582212202a886a16db3a2519c47f56f6c0a9aa72d1ffa8b18909f2c5e7edebc433eceaab64736f6c634300080a0033",
}

// ApiABI is the input ABI used to generate the binding from.
// Deprecated: Use ApiMetaData.ABI instead.
var ApiABI = ApiMetaData.ABI

// ApiBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ApiMetaData.Bin instead.
var ApiBin = ApiMetaData.Bin

// DeployApi deploys a new Ethereum contract, binding an instance of Api to it.
func DeployApi(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Api, error) {
	parsed, err := ApiMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ApiBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Api{ApiCaller: ApiCaller{contract: contract}, ApiTransactor: ApiTransactor{contract: contract}, ApiFilterer: ApiFilterer{contract: contract}}, nil
}

// Api is an auto generated Go binding around an Ethereum contract.
type Api struct {
	ApiCaller     // Read-only binding to the contract
	ApiTransactor // Write-only binding to the contract
	ApiFilterer   // Log filterer for contract events
}

// ApiCaller is an auto generated read-only Go binding around an Ethereum contract.
type ApiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ApiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ApiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ApiSession struct {
	Contract     *Api              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ApiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ApiCallerSession struct {
	Contract *ApiCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ApiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ApiTransactorSession struct {
	Contract     *ApiTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ApiRaw is an auto generated low-level Go binding around an Ethereum contract.
type ApiRaw struct {
	Contract *Api // Generic contract binding to access the raw methods on
}

// ApiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ApiCallerRaw struct {
	Contract *ApiCaller // Generic read-only contract binding to access the raw methods on
}

// ApiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ApiTransactorRaw struct {
	Contract *ApiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewApi creates a new instance of Api, bound to a specific deployed contract.
func NewApi(address common.Address, backend bind.ContractBackend) (*Api, error) {
	contract, err := bindApi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Api{ApiCaller: ApiCaller{contract: contract}, ApiTransactor: ApiTransactor{contract: contract}, ApiFilterer: ApiFilterer{contract: contract}}, nil
}

// NewApiCaller creates a new read-only instance of Api, bound to a specific deployed contract.
func NewApiCaller(address common.Address, caller bind.ContractCaller) (*ApiCaller, error) {
	contract, err := bindApi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ApiCaller{contract: contract}, nil
}

// NewApiTransactor creates a new write-only instance of Api, bound to a specific deployed contract.
func NewApiTransactor(address common.Address, transactor bind.ContractTransactor) (*ApiTransactor, error) {
	contract, err := bindApi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ApiTransactor{contract: contract}, nil
}

// NewApiFilterer creates a new log filterer instance of Api, bound to a specific deployed contract.
func NewApiFilterer(address common.Address, filterer bind.ContractFilterer) (*ApiFilterer, error) {
	contract, err := bindApi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ApiFilterer{contract: contract}, nil
}

// bindApi binds a generic wrapper to an already deployed contract.
func bindApi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ApiABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Api *ApiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Api.Contract.ApiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Api *ApiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.Contract.ApiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Api *ApiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Api.Contract.ApiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Api *ApiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Api.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Api *ApiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Api *ApiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Api.Contract.contract.Transact(opts, method, params...)
}

// Balance is a free data retrieval call binding the contract method 0x0ef67887.
//
// Solidity: function Balance() view returns(uint256)
func (_Api *ApiCaller) Balance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "Balance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balance is a free data retrieval call binding the contract method 0x0ef67887.
//
// Solidity: function Balance() view returns(uint256)
func (_Api *ApiSession) Balance() (*big.Int, error) {
	return _Api.Contract.Balance(&_Api.CallOpts)
}

// Balance is a free data retrieval call binding the contract method 0x0ef67887.
//
// Solidity: function Balance() view returns(uint256)
func (_Api *ApiCallerSession) Balance() (*big.Int, error) {
	return _Api.Contract.Balance(&_Api.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Api *ApiCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Api *ApiSession) Admin() (common.Address, error) {
	return _Api.Contract.Admin(&_Api.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Api *ApiCallerSession) Admin() (common.Address, error) {
	return _Api.Contract.Admin(&_Api.CallOpts)
}

// Deposite is a paid mutator transaction binding the contract method 0x42002bc8.
//
// Solidity: function Deposite(uint256 amt) returns(uint256)
func (_Api *ApiTransactor) Deposite(opts *bind.TransactOpts, amt *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "Deposite", amt)
}

// Deposite is a paid mutator transaction binding the contract method 0x42002bc8.
//
// Solidity: function Deposite(uint256 amt) returns(uint256)
func (_Api *ApiSession) Deposite(amt *big.Int) (*types.Transaction, error) {
	return _Api.Contract.Deposite(&_Api.TransactOpts, amt)
}

// Deposite is a paid mutator transaction binding the contract method 0x42002bc8.
//
// Solidity: function Deposite(uint256 amt) returns(uint256)
func (_Api *ApiTransactorSession) Deposite(amt *big.Int) (*types.Transaction, error) {
	return _Api.Contract.Deposite(&_Api.TransactOpts, amt)
}

// Withdrawl is a paid mutator transaction binding the contract method 0xe615c1a0.
//
// Solidity: function Withdrawl(uint256 _amt) returns()
func (_Api *ApiTransactor) Withdrawl(opts *bind.TransactOpts, _amt *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "Withdrawl", _amt)
}

// Withdrawl is a paid mutator transaction binding the contract method 0xe615c1a0.
//
// Solidity: function Withdrawl(uint256 _amt) returns()
func (_Api *ApiSession) Withdrawl(_amt *big.Int) (*types.Transaction, error) {
	return _Api.Contract.Withdrawl(&_Api.TransactOpts, _amt)
}

// Withdrawl is a paid mutator transaction binding the contract method 0xe615c1a0.
//
// Solidity: function Withdrawl(uint256 _amt) returns()
func (_Api *ApiTransactorSession) Withdrawl(_amt *big.Int) (*types.Transaction, error) {
	return _Api.Contract.Withdrawl(&_Api.TransactOpts, _amt)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Api *ApiTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Api *ApiSession) Receive() (*types.Transaction, error) {
	return _Api.Contract.Receive(&_Api.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Api *ApiTransactorSession) Receive() (*types.Transaction, error) {
	return _Api.Contract.Receive(&_Api.TransactOpts)
}

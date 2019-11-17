// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ContractABI is the input ABI used to generate the binding from.
const ContractABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"doc\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"name\":\"appendValidated\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"docMedOps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"docSuccesses\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"doc\",\"type\":\"string\"}],\"name\":\"getDocMedOps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"doc\",\"type\":\"string\"}],\"name\":\"getDocSuccess\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"doc\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"updateValidated\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"validated\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ContractBin is the compiled bytecode used for deploying new contracts.
var ContractBin = "0x608060405234801561001057600080fd5b50610a2b806100206000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c80636bd5bbc21161005b5780636bd5bbc2146102e7578063db8ddced146103b6578063e23125e114610485578063e6ea34cd146105625761007d565b806309cd5a20146100825780631c2aeed814610151578063460f2e6c14610222575b600080fd5b61013b6004803603602081101561009857600080fd5b81019080803590602001906401000000008111156100b557600080fd5b8201836020820111156100c757600080fd5b803590602001918460018302840111640100000000831117156100e957600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050610631565b6040518082815260200191505060405180910390f35b6102206004803603606081101561016757600080fd5b810190808035906020019064010000000081111561018457600080fd5b82018360208201111561019657600080fd5b803590602001918460018302840111640100000000831117156101b857600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290803590602001909291908035151590602001909291905050506106a4565b005b6102e56004803603604081101561023857600080fd5b810190808035906020019064010000000081111561025557600080fd5b82018360208201111561026757600080fd5b8035906020019184600183028401116401000000008311171561028957600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505091929192908035906020019092919050505061074e565b005b6103a0600480360360208110156102fd57600080fd5b810190808035906020019064010000000081111561031a57600080fd5b82018360208201111561032c57600080fd5b8035906020019184600183028401116401000000008311171561034e57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505091929192905050506107e5565b6040518082815260200191505060405180910390f35b61046f600480360360208110156103cc57600080fd5b81019080803590602001906401000000008111156103e957600080fd5b8201836020820111156103fb57600080fd5b8035906020019184600183028401116401000000008311171561041d57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050610813565b6040518082815260200191505060405180910390f35b6105486004803603604081101561049b57600080fd5b81019080803590602001906401000000008111156104b857600080fd5b8201836020820111156104ca57600080fd5b803590602001918460018302840111640100000000831117156104ec57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050919291929080359060200190929190505050610886565b604051808215151515815260200191505060405180910390f35b61061b6004803603602081101561057857600080fd5b810190808035906020019064010000000081111561059557600080fd5b8201836020820111156105a757600080fd5b803590602001918460018302840111640100000000831117156105c957600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505091929192905050506108cb565b6040518082815260200191505060405180910390f35b60006002826040518082805190602001908083835b602083106106695780518252602082019150602081019050602083039250610646565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040518091039020549050919050565b600080846040518082805190602001908083835b602083106106db57805182526020820191506020810190506020830392506106b8565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040518091039020600084815260200190815260200160002060006101000a81548160ff02191690831515021790555061073f836108f9565b6107498382610973565b505050565b60016000836040518082805190602001908083835b602083106107865780518252602082019150602081019050602083039250610763565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040518091039020600083815260200190815260200160002060006101000a81548160ff0219169083151502179055505050565b6001818051602081018201805184825260208301602085012081835280955050505050506000915090505481565b60006001826040518082805190602001908083835b6020831061084b5780518252602082019150602081019050602083039250610828565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040518091039020549050919050565b6000828051602081018201805184825260208301602085012081835280955050505050506020528060005260406000206000915091509054906101000a900460ff1681565b6002818051602081018201805184825260208301602085012081835280955050505050506000915090505481565b600180826040518082805190602001908083835b60208310610930578051825260208201915060208101905060208303925061090d565b6001836020036101000a03801982511681845116808217855250505050505090500191505090815260200160405180910390206000828254019250508190555050565b80156109f25760016002836040518082805190602001908083835b602083106109b1578051825260208201915060208101905060208303925061098e565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040518091039020600082825401925050819055505b505056fea265627a7a7231582075ccf4016f41a5e14d2a8627d6e7123aa1463ef724f508cf298d1876011d57ab64736f6c634300050d0032"

// DeployContract deploys a new Ethereum contract, binding an instance of Contract to it.
func DeployContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Contract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// DocMedOps is a free data retrieval call binding the contract method 0x6bd5bbc2.
//
// Solidity: function docMedOps(string ) constant returns(uint256)
func (_Contract *ContractCaller) DocMedOps(opts *bind.CallOpts, arg0 string) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "docMedOps", arg0)
	return *ret0, err
}

// DocMedOps is a free data retrieval call binding the contract method 0x6bd5bbc2.
//
// Solidity: function docMedOps(string ) constant returns(uint256)
func (_Contract *ContractSession) DocMedOps(arg0 string) (*big.Int, error) {
	return _Contract.Contract.DocMedOps(&_Contract.CallOpts, arg0)
}

// DocMedOps is a free data retrieval call binding the contract method 0x6bd5bbc2.
//
// Solidity: function docMedOps(string ) constant returns(uint256)
func (_Contract *ContractCallerSession) DocMedOps(arg0 string) (*big.Int, error) {
	return _Contract.Contract.DocMedOps(&_Contract.CallOpts, arg0)
}

// DocSuccesses is a free data retrieval call binding the contract method 0xe6ea34cd.
//
// Solidity: function docSuccesses(string ) constant returns(uint256)
func (_Contract *ContractCaller) DocSuccesses(opts *bind.CallOpts, arg0 string) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "docSuccesses", arg0)
	return *ret0, err
}

// DocSuccesses is a free data retrieval call binding the contract method 0xe6ea34cd.
//
// Solidity: function docSuccesses(string ) constant returns(uint256)
func (_Contract *ContractSession) DocSuccesses(arg0 string) (*big.Int, error) {
	return _Contract.Contract.DocSuccesses(&_Contract.CallOpts, arg0)
}

// DocSuccesses is a free data retrieval call binding the contract method 0xe6ea34cd.
//
// Solidity: function docSuccesses(string ) constant returns(uint256)
func (_Contract *ContractCallerSession) DocSuccesses(arg0 string) (*big.Int, error) {
	return _Contract.Contract.DocSuccesses(&_Contract.CallOpts, arg0)
}

// GetDocMedOps is a free data retrieval call binding the contract method 0xdb8ddced.
//
// Solidity: function getDocMedOps(string doc) constant returns(uint256)
func (_Contract *ContractCaller) GetDocMedOps(opts *bind.CallOpts, doc string) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "getDocMedOps", doc)
	return *ret0, err
}

// GetDocMedOps is a free data retrieval call binding the contract method 0xdb8ddced.
//
// Solidity: function getDocMedOps(string doc) constant returns(uint256)
func (_Contract *ContractSession) GetDocMedOps(doc string) (*big.Int, error) {
	return _Contract.Contract.GetDocMedOps(&_Contract.CallOpts, doc)
}

// GetDocMedOps is a free data retrieval call binding the contract method 0xdb8ddced.
//
// Solidity: function getDocMedOps(string doc) constant returns(uint256)
func (_Contract *ContractCallerSession) GetDocMedOps(doc string) (*big.Int, error) {
	return _Contract.Contract.GetDocMedOps(&_Contract.CallOpts, doc)
}

// GetDocSuccess is a free data retrieval call binding the contract method 0x09cd5a20.
//
// Solidity: function getDocSuccess(string doc) constant returns(uint256)
func (_Contract *ContractCaller) GetDocSuccess(opts *bind.CallOpts, doc string) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "getDocSuccess", doc)
	return *ret0, err
}

// GetDocSuccess is a free data retrieval call binding the contract method 0x09cd5a20.
//
// Solidity: function getDocSuccess(string doc) constant returns(uint256)
func (_Contract *ContractSession) GetDocSuccess(doc string) (*big.Int, error) {
	return _Contract.Contract.GetDocSuccess(&_Contract.CallOpts, doc)
}

// GetDocSuccess is a free data retrieval call binding the contract method 0x09cd5a20.
//
// Solidity: function getDocSuccess(string doc) constant returns(uint256)
func (_Contract *ContractCallerSession) GetDocSuccess(doc string) (*big.Int, error) {
	return _Contract.Contract.GetDocSuccess(&_Contract.CallOpts, doc)
}

// Validated is a free data retrieval call binding the contract method 0xe23125e1.
//
// Solidity: function validated(string , bytes32 ) constant returns(bool)
func (_Contract *ContractCaller) Validated(opts *bind.CallOpts, arg0 string, arg1 [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "validated", arg0, arg1)
	return *ret0, err
}

// Validated is a free data retrieval call binding the contract method 0xe23125e1.
//
// Solidity: function validated(string , bytes32 ) constant returns(bool)
func (_Contract *ContractSession) Validated(arg0 string, arg1 [32]byte) (bool, error) {
	return _Contract.Contract.Validated(&_Contract.CallOpts, arg0, arg1)
}

// Validated is a free data retrieval call binding the contract method 0xe23125e1.
//
// Solidity: function validated(string , bytes32 ) constant returns(bool)
func (_Contract *ContractCallerSession) Validated(arg0 string, arg1 [32]byte) (bool, error) {
	return _Contract.Contract.Validated(&_Contract.CallOpts, arg0, arg1)
}

// AppendValidated is a paid mutator transaction binding the contract method 0x1c2aeed8.
//
// Solidity: function appendValidated(string doc, bytes32 hash, bool success) returns()
func (_Contract *ContractTransactor) AppendValidated(opts *bind.TransactOpts, doc string, hash [32]byte, success bool) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "appendValidated", doc, hash, success)
}

// AppendValidated is a paid mutator transaction binding the contract method 0x1c2aeed8.
//
// Solidity: function appendValidated(string doc, bytes32 hash, bool success) returns()
func (_Contract *ContractSession) AppendValidated(doc string, hash [32]byte, success bool) (*types.Transaction, error) {
	return _Contract.Contract.AppendValidated(&_Contract.TransactOpts, doc, hash, success)
}

// AppendValidated is a paid mutator transaction binding the contract method 0x1c2aeed8.
//
// Solidity: function appendValidated(string doc, bytes32 hash, bool success) returns()
func (_Contract *ContractTransactorSession) AppendValidated(doc string, hash [32]byte, success bool) (*types.Transaction, error) {
	return _Contract.Contract.AppendValidated(&_Contract.TransactOpts, doc, hash, success)
}

// UpdateValidated is a paid mutator transaction binding the contract method 0x460f2e6c.
//
// Solidity: function updateValidated(string doc, bytes32 hash) returns()
func (_Contract *ContractTransactor) UpdateValidated(opts *bind.TransactOpts, doc string, hash [32]byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "updateValidated", doc, hash)
}

// UpdateValidated is a paid mutator transaction binding the contract method 0x460f2e6c.
//
// Solidity: function updateValidated(string doc, bytes32 hash) returns()
func (_Contract *ContractSession) UpdateValidated(doc string, hash [32]byte) (*types.Transaction, error) {
	return _Contract.Contract.UpdateValidated(&_Contract.TransactOpts, doc, hash)
}

// UpdateValidated is a paid mutator transaction binding the contract method 0x460f2e6c.
//
// Solidity: function updateValidated(string doc, bytes32 hash) returns()
func (_Contract *ContractTransactorSession) UpdateValidated(doc string, hash [32]byte) (*types.Transaction, error) {
	return _Contract.Contract.UpdateValidated(&_Contract.TransactOpts, doc, hash)
}

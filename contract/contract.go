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
const ContractABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"doc\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"name\":\"appendValidated\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"docMedOps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"docSuccesses\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"validated\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ContractBin is the compiled bytecode used for deploying new contracts.
var ContractBin = "0x608060405234801561001057600080fd5b50610781806100206000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c80631c2aeed81461005c578063460f2e6c1461012d5780636bd5bbc2146101f2578063e23125e1146102c1578063e6ea34cd1461039e575b600080fd5b61012b6004803603606081101561007257600080fd5b810190808035906020019064010000000081111561008f57600080fd5b8201836020820111156100a157600080fd5b803590602001918460018302840111640100000000831117156100c357600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505091929192908035906020019092919080351515906020019092919050505061046d565b005b6101f06004803603604081101561014357600080fd5b810190808035906020019064010000000081111561016057600080fd5b82018360208201111561017257600080fd5b8035906020019184600183028401116401000000008311171561019457600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050919291929080359060200190929190505050610517565b005b6102ab6004803603602081101561020857600080fd5b810190808035906020019064010000000081111561022557600080fd5b82018360208201111561023757600080fd5b8035906020019184600183028401116401000000008311171561025957600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505091929192905050506105ae565b6040518082815260200191505060405180910390f35b610384600480360360408110156102d757600080fd5b81019080803590602001906401000000008111156102f457600080fd5b82018360208201111561030657600080fd5b8035906020019184600183028401116401000000008311171561032857600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290803590602001909291905050506105dc565b604051808215151515815260200191505060405180910390f35b610457600480360360208110156103b457600080fd5b81019080803590602001906401000000008111156103d157600080fd5b8201836020820111156103e357600080fd5b8035906020019184600183028401116401000000008311171561040557600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050610621565b6040518082815260200191505060405180910390f35b600080846040518082805190602001908083835b602083106104a45780518252602082019150602081019050602083039250610481565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040518091039020600084815260200190815260200160002060006101000a81548160ff0219169083151502179055506105088361064f565b61051283826106c9565b505050565b60016000836040518082805190602001908083835b6020831061054f578051825260208201915060208101905060208303925061052c565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040518091039020600083815260200190815260200160002060006101000a81548160ff0219169083151502179055505050565b6001818051602081018201805184825260208301602085012081835280955050505050506000915090505481565b6000828051602081018201805184825260208301602085012081835280955050505050506020528060005260406000206000915091509054906101000a900460ff1681565b6002818051602081018201805184825260208301602085012081835280955050505050506000915090505481565b600180826040518082805190602001908083835b602083106106865780518252602082019150602081019050602083039250610663565b6001836020036101000a03801982511681845116808217855250505050505090500191505090815260200160405180910390206000828254019250508190555050565b80156107485760016002836040518082805190602001908083835b6020831061070757805182526020820191506020810190506020830392506106e4565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040518091039020600082825401925050819055505b505056fea265627a7a72315820e81e9b29a5becf3bb96b033615b13b8199ced4728deaf1b86065046b3261b03a64736f6c634300050d0032"

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

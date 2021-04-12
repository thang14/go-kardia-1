// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package validator_set

import (
	"math/big"
	"strings"

	kardia "github.com/kardiachain/go-kardia"
	"github.com/kardiachain/go-kardia/lib/abi"
	"github.com/kardiachain/go-kardia/lib/abi/bind"
	"github.com/kardiachain/go-kardia/lib/common"
	"github.com/kardiachain/go-kardia/lib/event"
	"github.com/kardiachain/go-kardia/types"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = bind.Bind
	_ = kardia.NotFound
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ValidatorSetDescription is an auto generated low-level Go binding around an user-defined struct.
type ValidatorSetDescription struct {
	Moniker  string
	Identity string
	Website  string
	Email    string
	Details  string
}

// ValidatorSetABI is the input ABI used to generate the binding from.
const ValidatorSetABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"valAddr\",\"type\":\"address\"}],\"name\":\"ValidatorActivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"moniker\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"identity\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"website\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"details\",\"type\":\"string\"}],\"name\":\"ValidatorCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"valAddr\",\"type\":\"address\"}],\"name\":\"ValidatorDeActivated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"valAddrs\",\"type\":\"address[]\"}],\"name\":\"activate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"activeValidators\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signatures\",\"type\":\"bytes\"}],\"name\":\"checkSignatures\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"moniker\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"identity\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"website\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"details\",\"type\":\"string\"}],\"name\":\"createOrEdit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"valAddrs\",\"type\":\"address[]\"}],\"name\":\"deActivate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getActiveValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"addrs\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAtiveValidatorsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"val\",\"type\":\"address\"}],\"name\":\"getValidatorDescription\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"vals_\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"timelockContract_\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timelockContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"validators\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"moniker\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"identity\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"website\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"details\",\"type\":\"string\"}],\"internalType\":\"structValidatorSet.Description\",\"name\":\"description\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]=======lib/BytesLib.sol:BytesLib=======[]=======lib/ECDSA.sol:ECDSA=======[]=======lib/Initializable.sol:Initializable=======[]=======lib/SafeMath.sol:SafeMath=======[]"

// ValidatorSet is an auto generated Go binding around an Ethereum contract.
type ValidatorSet struct {
	ValidatorSetCaller     // Read-only binding to the contract
	ValidatorSetTransactor // Write-only binding to the contract
	ValidatorSetFilterer   // Log filterer for contract events
}

// ValidatorSetCaller is an auto generated read-only Go binding around an Ethereum contract.
type ValidatorSetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorSetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ValidatorSetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorSetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ValidatorSetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorSetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ValidatorSetSession struct {
	Contract     *ValidatorSet     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ValidatorSetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ValidatorSetCallerSession struct {
	Contract *ValidatorSetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ValidatorSetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ValidatorSetTransactorSession struct {
	Contract     *ValidatorSetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ValidatorSetRaw is an auto generated low-level Go binding around an Ethereum contract.
type ValidatorSetRaw struct {
	Contract *ValidatorSet // Generic contract binding to access the raw methods on
}

// ValidatorSetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ValidatorSetCallerRaw struct {
	Contract *ValidatorSetCaller // Generic read-only contract binding to access the raw methods on
}

// ValidatorSetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ValidatorSetTransactorRaw struct {
	Contract *ValidatorSetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewValidatorSet creates a new instance of ValidatorSet, bound to a specific deployed contract.
func NewValidatorSet(address common.Address, backend bind.ContractBackend) (*ValidatorSet, error) {
	contract, err := bindValidatorSet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ValidatorSet{ValidatorSetCaller: ValidatorSetCaller{contract: contract}, ValidatorSetTransactor: ValidatorSetTransactor{contract: contract}, ValidatorSetFilterer: ValidatorSetFilterer{contract: contract}}, nil
}

// NewValidatorSetCaller creates a new read-only instance of ValidatorSet, bound to a specific deployed contract.
func NewValidatorSetCaller(address common.Address, caller bind.ContractCaller) (*ValidatorSetCaller, error) {
	contract, err := bindValidatorSet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorSetCaller{contract: contract}, nil
}

// NewValidatorSetTransactor creates a new write-only instance of ValidatorSet, bound to a specific deployed contract.
func NewValidatorSetTransactor(address common.Address, transactor bind.ContractTransactor) (*ValidatorSetTransactor, error) {
	contract, err := bindValidatorSet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorSetTransactor{contract: contract}, nil
}

// NewValidatorSetFilterer creates a new log filterer instance of ValidatorSet, bound to a specific deployed contract.
func NewValidatorSetFilterer(address common.Address, filterer bind.ContractFilterer) (*ValidatorSetFilterer, error) {
	contract, err := bindValidatorSet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ValidatorSetFilterer{contract: contract}, nil
}

// bindValidatorSet binds a generic wrapper to an already deployed contract.
func bindValidatorSet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ValidatorSetABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValidatorSet *ValidatorSetRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ValidatorSet.Contract.ValidatorSetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValidatorSet *ValidatorSetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorSet.Contract.ValidatorSetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValidatorSet *ValidatorSetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValidatorSet.Contract.ValidatorSetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValidatorSet *ValidatorSetCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ValidatorSet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValidatorSet *ValidatorSetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorSet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValidatorSet *ValidatorSetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValidatorSet.Contract.contract.Transact(opts, method, params...)
}

// ActiveValidators is a free data retrieval call binding the contract method 0x14f64c78.
//
// Solidity: function activeValidators(uint256 ) view returns(address)
func (_ValidatorSet *ValidatorSetCaller) ActiveValidators(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ValidatorSet.contract.Call(opts, out, "activeValidators", arg0)
	return *ret0, err
}

// ActiveValidators is a free data retrieval call binding the contract method 0x14f64c78.
//
// Solidity: function activeValidators(uint256 ) view returns(address)
func (_ValidatorSet *ValidatorSetSession) ActiveValidators(arg0 *big.Int) (common.Address, error) {
	return _ValidatorSet.Contract.ActiveValidators(&_ValidatorSet.CallOpts, arg0)
}

// ActiveValidators is a free data retrieval call binding the contract method 0x14f64c78.
//
// Solidity: function activeValidators(uint256 ) view returns(address)
func (_ValidatorSet *ValidatorSetCallerSession) ActiveValidators(arg0 *big.Int) (common.Address, error) {
	return _ValidatorSet.Contract.ActiveValidators(&_ValidatorSet.CallOpts, arg0)
}

// CheckSignatures is a free data retrieval call binding the contract method 0xed516d51.
//
// Solidity: function checkSignatures(bytes32 digest, bytes signatures) view returns(bool)
func (_ValidatorSet *ValidatorSetCaller) CheckSignatures(opts *bind.CallOpts, digest [32]byte, signatures []byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ValidatorSet.contract.Call(opts, out, "checkSignatures", digest, signatures)
	return *ret0, err
}

// CheckSignatures is a free data retrieval call binding the contract method 0xed516d51.
//
// Solidity: function checkSignatures(bytes32 digest, bytes signatures) view returns(bool)
func (_ValidatorSet *ValidatorSetSession) CheckSignatures(digest [32]byte, signatures []byte) (bool, error) {
	return _ValidatorSet.Contract.CheckSignatures(&_ValidatorSet.CallOpts, digest, signatures)
}

// CheckSignatures is a free data retrieval call binding the contract method 0xed516d51.
//
// Solidity: function checkSignatures(bytes32 digest, bytes signatures) view returns(bool)
func (_ValidatorSet *ValidatorSetCallerSession) CheckSignatures(digest [32]byte, signatures []byte) (bool, error) {
	return _ValidatorSet.Contract.CheckSignatures(&_ValidatorSet.CallOpts, digest, signatures)
}

// GetActiveValidators is a free data retrieval call binding the contract method 0x9de70258.
//
// Solidity: function getActiveValidators() view returns(address[] addrs)
func (_ValidatorSet *ValidatorSetCaller) GetActiveValidators(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _ValidatorSet.contract.Call(opts, out, "getActiveValidators")
	return *ret0, err
}

// GetActiveValidators is a free data retrieval call binding the contract method 0x9de70258.
//
// Solidity: function getActiveValidators() view returns(address[] addrs)
func (_ValidatorSet *ValidatorSetSession) GetActiveValidators() ([]common.Address, error) {
	return _ValidatorSet.Contract.GetActiveValidators(&_ValidatorSet.CallOpts)
}

// GetActiveValidators is a free data retrieval call binding the contract method 0x9de70258.
//
// Solidity: function getActiveValidators() view returns(address[] addrs)
func (_ValidatorSet *ValidatorSetCallerSession) GetActiveValidators() ([]common.Address, error) {
	return _ValidatorSet.Contract.GetActiveValidators(&_ValidatorSet.CallOpts)
}

// GetAtiveValidatorsLength is a free data retrieval call binding the contract method 0xb60ae944.
//
// Solidity: function getAtiveValidatorsLength() view returns(uint256)
func (_ValidatorSet *ValidatorSetCaller) GetAtiveValidatorsLength(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ValidatorSet.contract.Call(opts, out, "getAtiveValidatorsLength")
	return *ret0, err
}

// GetAtiveValidatorsLength is a free data retrieval call binding the contract method 0xb60ae944.
//
// Solidity: function getAtiveValidatorsLength() view returns(uint256)
func (_ValidatorSet *ValidatorSetSession) GetAtiveValidatorsLength() (*big.Int, error) {
	return _ValidatorSet.Contract.GetAtiveValidatorsLength(&_ValidatorSet.CallOpts)
}

// GetAtiveValidatorsLength is a free data retrieval call binding the contract method 0xb60ae944.
//
// Solidity: function getAtiveValidatorsLength() view returns(uint256)
func (_ValidatorSet *ValidatorSetCallerSession) GetAtiveValidatorsLength() (*big.Int, error) {
	return _ValidatorSet.Contract.GetAtiveValidatorsLength(&_ValidatorSet.CallOpts)
}

// GetValidatorDescription is a free data retrieval call binding the contract method 0xa43569b3.
//
// Solidity: function getValidatorDescription(address val) view returns(string, string, string, string, string)
func (_ValidatorSet *ValidatorSetCaller) GetValidatorDescription(opts *bind.CallOpts, val common.Address) (string, string, string, string, string, error) {
	var (
		ret0 = new(string)
		ret1 = new(string)
		ret2 = new(string)
		ret3 = new(string)
		ret4 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
	}
	err := _ValidatorSet.contract.Call(opts, out, "getValidatorDescription", val)
	return *ret0, *ret1, *ret2, *ret3, *ret4, err
}

// GetValidatorDescription is a free data retrieval call binding the contract method 0xa43569b3.
//
// Solidity: function getValidatorDescription(address val) view returns(string, string, string, string, string)
func (_ValidatorSet *ValidatorSetSession) GetValidatorDescription(val common.Address) (string, string, string, string, string, error) {
	return _ValidatorSet.Contract.GetValidatorDescription(&_ValidatorSet.CallOpts, val)
}

// GetValidatorDescription is a free data retrieval call binding the contract method 0xa43569b3.
//
// Solidity: function getValidatorDescription(address val) view returns(string, string, string, string, string)
func (_ValidatorSet *ValidatorSetCallerSession) GetValidatorDescription(val common.Address) (string, string, string, string, string, error) {
	return _ValidatorSet.Contract.GetValidatorDescription(&_ValidatorSet.CallOpts, val)
}

// IsActive is a free data retrieval call binding the contract method 0x9f8a13d7.
//
// Solidity: function isActive(address ) view returns(bool)
func (_ValidatorSet *ValidatorSetCaller) IsActive(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ValidatorSet.contract.Call(opts, out, "isActive", arg0)
	return *ret0, err
}

// IsActive is a free data retrieval call binding the contract method 0x9f8a13d7.
//
// Solidity: function isActive(address ) view returns(bool)
func (_ValidatorSet *ValidatorSetSession) IsActive(arg0 common.Address) (bool, error) {
	return _ValidatorSet.Contract.IsActive(&_ValidatorSet.CallOpts, arg0)
}

// IsActive is a free data retrieval call binding the contract method 0x9f8a13d7.
//
// Solidity: function isActive(address ) view returns(bool)
func (_ValidatorSet *ValidatorSetCallerSession) IsActive(arg0 common.Address) (bool, error) {
	return _ValidatorSet.Contract.IsActive(&_ValidatorSet.CallOpts, arg0)
}

// TimelockContract is a free data retrieval call binding the contract method 0xa9181cc4.
//
// Solidity: function timelockContract() view returns(address)
func (_ValidatorSet *ValidatorSetCaller) TimelockContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ValidatorSet.contract.Call(opts, out, "timelockContract")
	return *ret0, err
}

// TimelockContract is a free data retrieval call binding the contract method 0xa9181cc4.
//
// Solidity: function timelockContract() view returns(address)
func (_ValidatorSet *ValidatorSetSession) TimelockContract() (common.Address, error) {
	return _ValidatorSet.Contract.TimelockContract(&_ValidatorSet.CallOpts)
}

// TimelockContract is a free data retrieval call binding the contract method 0xa9181cc4.
//
// Solidity: function timelockContract() view returns(address)
func (_ValidatorSet *ValidatorSetCallerSession) TimelockContract() (common.Address, error) {
	return _ValidatorSet.Contract.TimelockContract(&_ValidatorSet.CallOpts)
}

// Validators is a free data retrieval call binding the contract method 0xfa52c7d8.
//
// Solidity: function validators(address ) view returns((string,string,string,string,string) description)
func (_ValidatorSet *ValidatorSetCaller) Validators(opts *bind.CallOpts, arg0 common.Address) (ValidatorSetDescription, error) {
	var (
		ret0 = new(ValidatorSetDescription)
	)
	out := ret0
	err := _ValidatorSet.contract.Call(opts, out, "validators", arg0)
	return *ret0, err
}

// Validators is a free data retrieval call binding the contract method 0xfa52c7d8.
//
// Solidity: function validators(address ) view returns((string,string,string,string,string) description)
func (_ValidatorSet *ValidatorSetSession) Validators(arg0 common.Address) (ValidatorSetDescription, error) {
	return _ValidatorSet.Contract.Validators(&_ValidatorSet.CallOpts, arg0)
}

// Validators is a free data retrieval call binding the contract method 0xfa52c7d8.
//
// Solidity: function validators(address ) view returns((string,string,string,string,string) description)
func (_ValidatorSet *ValidatorSetCallerSession) Validators(arg0 common.Address) (ValidatorSetDescription, error) {
	return _ValidatorSet.Contract.Validators(&_ValidatorSet.CallOpts, arg0)
}

// Activate is a paid mutator transaction binding the contract method 0x5293840d.
//
// Solidity: function activate(address[] valAddrs) returns()
func (_ValidatorSet *ValidatorSetTransactor) Activate(opts *bind.TransactOpts, valAddrs []common.Address) (*types.Transaction, error) {
	return _ValidatorSet.contract.Transact(opts, "activate", valAddrs)
}

// Activate is a paid mutator transaction binding the contract method 0x5293840d.
//
// Solidity: function activate(address[] valAddrs) returns()
func (_ValidatorSet *ValidatorSetSession) Activate(valAddrs []common.Address) (*types.Transaction, error) {
	return _ValidatorSet.Contract.Activate(&_ValidatorSet.TransactOpts, valAddrs)
}

// Activate is a paid mutator transaction binding the contract method 0x5293840d.
//
// Solidity: function activate(address[] valAddrs) returns()
func (_ValidatorSet *ValidatorSetTransactorSession) Activate(valAddrs []common.Address) (*types.Transaction, error) {
	return _ValidatorSet.Contract.Activate(&_ValidatorSet.TransactOpts, valAddrs)
}

// CreateOrEdit is a paid mutator transaction binding the contract method 0xfcfc7653.
//
// Solidity: function createOrEdit(string moniker, string identity, string website, string email, string details) returns(bool)
func (_ValidatorSet *ValidatorSetTransactor) CreateOrEdit(opts *bind.TransactOpts, moniker string, identity string, website string, email string, details string) (*types.Transaction, error) {
	return _ValidatorSet.contract.Transact(opts, "createOrEdit", moniker, identity, website, email, details)
}

// CreateOrEdit is a paid mutator transaction binding the contract method 0xfcfc7653.
//
// Solidity: function createOrEdit(string moniker, string identity, string website, string email, string details) returns(bool)
func (_ValidatorSet *ValidatorSetSession) CreateOrEdit(moniker string, identity string, website string, email string, details string) (*types.Transaction, error) {
	return _ValidatorSet.Contract.CreateOrEdit(&_ValidatorSet.TransactOpts, moniker, identity, website, email, details)
}

// CreateOrEdit is a paid mutator transaction binding the contract method 0xfcfc7653.
//
// Solidity: function createOrEdit(string moniker, string identity, string website, string email, string details) returns(bool)
func (_ValidatorSet *ValidatorSetTransactorSession) CreateOrEdit(moniker string, identity string, website string, email string, details string) (*types.Transaction, error) {
	return _ValidatorSet.Contract.CreateOrEdit(&_ValidatorSet.TransactOpts, moniker, identity, website, email, details)
}

// DeActivate is a paid mutator transaction binding the contract method 0x62e00a95.
//
// Solidity: function deActivate(address[] valAddrs) returns()
func (_ValidatorSet *ValidatorSetTransactor) DeActivate(opts *bind.TransactOpts, valAddrs []common.Address) (*types.Transaction, error) {
	return _ValidatorSet.contract.Transact(opts, "deActivate", valAddrs)
}

// DeActivate is a paid mutator transaction binding the contract method 0x62e00a95.
//
// Solidity: function deActivate(address[] valAddrs) returns()
func (_ValidatorSet *ValidatorSetSession) DeActivate(valAddrs []common.Address) (*types.Transaction, error) {
	return _ValidatorSet.Contract.DeActivate(&_ValidatorSet.TransactOpts, valAddrs)
}

// DeActivate is a paid mutator transaction binding the contract method 0x62e00a95.
//
// Solidity: function deActivate(address[] valAddrs) returns()
func (_ValidatorSet *ValidatorSetTransactorSession) DeActivate(valAddrs []common.Address) (*types.Transaction, error) {
	return _ValidatorSet.Contract.DeActivate(&_ValidatorSet.TransactOpts, valAddrs)
}

// Initialize is a paid mutator transaction binding the contract method 0x462d0b2e.
//
// Solidity: function initialize(address[] vals_, address timelockContract_) returns()
func (_ValidatorSet *ValidatorSetTransactor) Initialize(opts *bind.TransactOpts, vals_ []common.Address, timelockContract_ common.Address) (*types.Transaction, error) {
	return _ValidatorSet.contract.Transact(opts, "initialize", vals_, timelockContract_)
}

// Initialize is a paid mutator transaction binding the contract method 0x462d0b2e.
//
// Solidity: function initialize(address[] vals_, address timelockContract_) returns()
func (_ValidatorSet *ValidatorSetSession) Initialize(vals_ []common.Address, timelockContract_ common.Address) (*types.Transaction, error) {
	return _ValidatorSet.Contract.Initialize(&_ValidatorSet.TransactOpts, vals_, timelockContract_)
}

// Initialize is a paid mutator transaction binding the contract method 0x462d0b2e.
//
// Solidity: function initialize(address[] vals_, address timelockContract_) returns()
func (_ValidatorSet *ValidatorSetTransactorSession) Initialize(vals_ []common.Address, timelockContract_ common.Address) (*types.Transaction, error) {
	return _ValidatorSet.Contract.Initialize(&_ValidatorSet.TransactOpts, vals_, timelockContract_)
}

// ValidatorSetValidatorActivatedIterator is returned from FilterValidatorActivated and is used to iterate over the raw logs and unpacked data for ValidatorActivated events raised by the ValidatorSet contract.
type ValidatorSetValidatorActivatedIterator struct {
	Event *ValidatorSetValidatorActivated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log     // Log channel receiving the found contract events
	sub  event.Subscription // Subscription for errors, completion and termination
	done bool               // Whether the subscription completed delivering logs
	fail error              // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ValidatorSetValidatorActivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorSetValidatorActivated)
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
		it.Event = new(ValidatorSetValidatorActivated)
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
func (it *ValidatorSetValidatorActivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorSetValidatorActivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorSetValidatorActivated represents a ValidatorActivated event raised by the ValidatorSet contract.
type ValidatorSetValidatorActivated struct {
	ValAddr common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterValidatorActivated is a free log retrieval operation binding the contract event 0xb2b8c5f713f27009f7a78ec2167999c472e5ee8bc591102da17bc4a6551236c8.
//
// Solidity: event ValidatorActivated(address valAddr)
func (_ValidatorSet *ValidatorSetFilterer) FilterValidatorActivated(opts *bind.FilterOpts) (*ValidatorSetValidatorActivatedIterator, error) {

	logs, sub, err := _ValidatorSet.contract.FilterLogs(opts, "ValidatorActivated")
	if err != nil {
		return nil, err
	}
	return &ValidatorSetValidatorActivatedIterator{contract: _ValidatorSet.contract, event: "ValidatorActivated", logs: logs, sub: sub}, nil
}

// WatchValidatorActivated is a free log subscription operation binding the contract event 0xb2b8c5f713f27009f7a78ec2167999c472e5ee8bc591102da17bc4a6551236c8.
//
// Solidity: event ValidatorActivated(address valAddr)
func (_ValidatorSet *ValidatorSetFilterer) WatchValidatorActivated(opts *bind.WatchOpts, sink chan<- *ValidatorSetValidatorActivated) (event.Subscription, error) {

	logs, sub, err := _ValidatorSet.contract.WatchLogs(opts, "ValidatorActivated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorSetValidatorActivated)
				if err := _ValidatorSet.contract.UnpackLog(event, "ValidatorActivated", log); err != nil {
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

// ParseValidatorActivated is a log parse operation binding the contract event 0xb2b8c5f713f27009f7a78ec2167999c472e5ee8bc591102da17bc4a6551236c8.
//
// Solidity: event ValidatorActivated(address valAddr)
func (_ValidatorSet *ValidatorSetFilterer) ParseValidatorActivated(log types.Log) (*ValidatorSetValidatorActivated, error) {
	event := new(ValidatorSetValidatorActivated)
	if err := _ValidatorSet.contract.UnpackLog(event, "ValidatorActivated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ValidatorSetValidatorCreatedIterator is returned from FilterValidatorCreated and is used to iterate over the raw logs and unpacked data for ValidatorCreated events raised by the ValidatorSet contract.
type ValidatorSetValidatorCreatedIterator struct {
	Event *ValidatorSetValidatorCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log     // Log channel receiving the found contract events
	sub  event.Subscription // Subscription for errors, completion and termination
	done bool               // Whether the subscription completed delivering logs
	fail error              // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ValidatorSetValidatorCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorSetValidatorCreated)
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
		it.Event = new(ValidatorSetValidatorCreated)
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
func (it *ValidatorSetValidatorCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorSetValidatorCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorSetValidatorCreated represents a ValidatorCreated event raised by the ValidatorSet contract.
type ValidatorSetValidatorCreated struct {
	Owner    common.Address
	Moniker  string
	Identity string
	Website  string
	Email    string
	Details  string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterValidatorCreated is a free log retrieval operation binding the contract event 0x8f8bf971aa9ce1b3f2ee714af8895aa9e9d3a650ad2aaa45ea0a60548784f81a.
//
// Solidity: event ValidatorCreated(address owner, string moniker, string identity, string website, string email, string details)
func (_ValidatorSet *ValidatorSetFilterer) FilterValidatorCreated(opts *bind.FilterOpts) (*ValidatorSetValidatorCreatedIterator, error) {

	logs, sub, err := _ValidatorSet.contract.FilterLogs(opts, "ValidatorCreated")
	if err != nil {
		return nil, err
	}
	return &ValidatorSetValidatorCreatedIterator{contract: _ValidatorSet.contract, event: "ValidatorCreated", logs: logs, sub: sub}, nil
}

// WatchValidatorCreated is a free log subscription operation binding the contract event 0x8f8bf971aa9ce1b3f2ee714af8895aa9e9d3a650ad2aaa45ea0a60548784f81a.
//
// Solidity: event ValidatorCreated(address owner, string moniker, string identity, string website, string email, string details)
func (_ValidatorSet *ValidatorSetFilterer) WatchValidatorCreated(opts *bind.WatchOpts, sink chan<- *ValidatorSetValidatorCreated) (event.Subscription, error) {

	logs, sub, err := _ValidatorSet.contract.WatchLogs(opts, "ValidatorCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorSetValidatorCreated)
				if err := _ValidatorSet.contract.UnpackLog(event, "ValidatorCreated", log); err != nil {
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

// ParseValidatorCreated is a log parse operation binding the contract event 0x8f8bf971aa9ce1b3f2ee714af8895aa9e9d3a650ad2aaa45ea0a60548784f81a.
//
// Solidity: event ValidatorCreated(address owner, string moniker, string identity, string website, string email, string details)
func (_ValidatorSet *ValidatorSetFilterer) ParseValidatorCreated(log types.Log) (*ValidatorSetValidatorCreated, error) {
	event := new(ValidatorSetValidatorCreated)
	if err := _ValidatorSet.contract.UnpackLog(event, "ValidatorCreated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ValidatorSetValidatorDeActivatedIterator is returned from FilterValidatorDeActivated and is used to iterate over the raw logs and unpacked data for ValidatorDeActivated events raised by the ValidatorSet contract.
type ValidatorSetValidatorDeActivatedIterator struct {
	Event *ValidatorSetValidatorDeActivated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log     // Log channel receiving the found contract events
	sub  event.Subscription // Subscription for errors, completion and termination
	done bool               // Whether the subscription completed delivering logs
	fail error              // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ValidatorSetValidatorDeActivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorSetValidatorDeActivated)
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
		it.Event = new(ValidatorSetValidatorDeActivated)
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
func (it *ValidatorSetValidatorDeActivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorSetValidatorDeActivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorSetValidatorDeActivated represents a ValidatorDeActivated event raised by the ValidatorSet contract.
type ValidatorSetValidatorDeActivated struct {
	ValAddr common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterValidatorDeActivated is a free log retrieval operation binding the contract event 0xbbce43c6752a966a51f2b032b00fe474b61d2473dd046e553d4d8575d0167e0f.
//
// Solidity: event ValidatorDeActivated(address valAddr)
func (_ValidatorSet *ValidatorSetFilterer) FilterValidatorDeActivated(opts *bind.FilterOpts) (*ValidatorSetValidatorDeActivatedIterator, error) {

	logs, sub, err := _ValidatorSet.contract.FilterLogs(opts, "ValidatorDeActivated")
	if err != nil {
		return nil, err
	}
	return &ValidatorSetValidatorDeActivatedIterator{contract: _ValidatorSet.contract, event: "ValidatorDeActivated", logs: logs, sub: sub}, nil
}

// WatchValidatorDeActivated is a free log subscription operation binding the contract event 0xbbce43c6752a966a51f2b032b00fe474b61d2473dd046e553d4d8575d0167e0f.
//
// Solidity: event ValidatorDeActivated(address valAddr)
func (_ValidatorSet *ValidatorSetFilterer) WatchValidatorDeActivated(opts *bind.WatchOpts, sink chan<- *ValidatorSetValidatorDeActivated) (event.Subscription, error) {

	logs, sub, err := _ValidatorSet.contract.WatchLogs(opts, "ValidatorDeActivated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorSetValidatorDeActivated)
				if err := _ValidatorSet.contract.UnpackLog(event, "ValidatorDeActivated", log); err != nil {
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

// ParseValidatorDeActivated is a log parse operation binding the contract event 0xbbce43c6752a966a51f2b032b00fe474b61d2473dd046e553d4d8575d0167e0f.
//
// Solidity: event ValidatorDeActivated(address valAddr)
func (_ValidatorSet *ValidatorSetFilterer) ParseValidatorDeActivated(log types.Log) (*ValidatorSetValidatorDeActivated, error) {
	event := new(ValidatorSetValidatorDeActivated)
	if err := _ValidatorSet.contract.UnpackLog(event, "ValidatorDeActivated", log); err != nil {
		return nil, err
	}
	return event, nil
}

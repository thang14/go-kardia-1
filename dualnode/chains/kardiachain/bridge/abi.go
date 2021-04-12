// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bridge

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

// BridgeABI is the input ABI used to generate the binding from.
const BridgeABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"recipient\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"token\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"symbol\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumBridge.TokenLockType\",\"name\":\"locktype\",\"type\":\"uint8\"}],\"name\":\"TokenAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"symbol\",\"type\":\"bytes32\"}],\"name\":\"TokenRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"depositor\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"token\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEPOSIT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"symbol\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"enumBridge.TokenLockType\",\"name\":\"lockType\",\"type\":\"uint8\"}],\"name\":\"addToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"token\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"recipient\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChainId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validatorSet_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"timeLockContract_\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"symbol\",\"type\":\"bytes32\"}],\"name\":\"removeToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timeLockContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"tokenInfos\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"enumBridge.TokenLockType\",\"name\":\"lockType\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorSet\",\"outputs\":[{\"internalType\":\"contractIValidatorSet\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depositId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"depositor\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"token\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signatures\",\"type\":\"bytes\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]=======interfaces/IMintBurnToken.sol:IMintBurnToken=======[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]=======interfaces/IValidatorSet.sol:IValidatorSet=======[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"val\",\"type\":\"address\"}],\"name\":\"active\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signatures\",\"type\":\"bytes\"}],\"name\":\"checkSignatures\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"val\",\"type\":\"address\"}],\"name\":\"deactive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"val\",\"type\":\"address\"}],\"name\":\"isActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]=======lib/Initializable.sol:Initializable=======[]=======lib/TransferHelper.sol:TransferHelper=======[]"

// Bridge is an auto generated Go binding around an Ethereum contract.
type Bridge struct {
	BridgeCaller     // Read-only binding to the contract
	BridgeTransactor // Write-only binding to the contract
	BridgeFilterer   // Log filterer for contract events
}

// BridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeSession struct {
	Contract     *Bridge           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeCallerSession struct {
	Contract *BridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeTransactorSession struct {
	Contract     *BridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeRaw struct {
	Contract *Bridge // Generic contract binding to access the raw methods on
}

// BridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeCallerRaw struct {
	Contract *BridgeCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeTransactorRaw struct {
	Contract *BridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridge creates a new instance of Bridge, bound to a specific deployed contract.
func NewBridge(address common.Address, backend bind.ContractBackend) (*Bridge, error) {
	contract, err := bindBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bridge{BridgeCaller: BridgeCaller{contract: contract}, BridgeTransactor: BridgeTransactor{contract: contract}, BridgeFilterer: BridgeFilterer{contract: contract}}, nil
}

// NewBridgeCaller creates a new read-only instance of Bridge, bound to a specific deployed contract.
func NewBridgeCaller(address common.Address, caller bind.ContractCaller) (*BridgeCaller, error) {
	contract, err := bindBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeCaller{contract: contract}, nil
}

// NewBridgeTransactor creates a new write-only instance of Bridge, bound to a specific deployed contract.
func NewBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeTransactor, error) {
	contract, err := bindBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeTransactor{contract: contract}, nil
}

// NewBridgeFilterer creates a new log filterer instance of Bridge, bound to a specific deployed contract.
func NewBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeFilterer, error) {
	contract, err := bindBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeFilterer{contract: contract}, nil
}

// bindBridge binds a generic wrapper to an already deployed contract.
func bindBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BridgeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridge *BridgeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Bridge.Contract.BridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridge *BridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.Contract.BridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridge *BridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridge.Contract.BridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridge *BridgeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Bridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridge *BridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridge *BridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridge.Contract.contract.Transact(opts, method, params...)
}

// DEPOSITTYPEHASH is a free data retrieval call binding the contract method 0x48825e94.
//
// Solidity: function DEPOSIT_TYPEHASH() view returns(bytes32)
func (_Bridge *BridgeCaller) DEPOSITTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Bridge.contract.Call(opts, out, "DEPOSIT_TYPEHASH")
	return *ret0, err
}

// DEPOSITTYPEHASH is a free data retrieval call binding the contract method 0x48825e94.
//
// Solidity: function DEPOSIT_TYPEHASH() view returns(bytes32)
func (_Bridge *BridgeSession) DEPOSITTYPEHASH() ([32]byte, error) {
	return _Bridge.Contract.DEPOSITTYPEHASH(&_Bridge.CallOpts)
}

// DEPOSITTYPEHASH is a free data retrieval call binding the contract method 0x48825e94.
//
// Solidity: function DEPOSIT_TYPEHASH() view returns(bytes32)
func (_Bridge *BridgeCallerSession) DEPOSITTYPEHASH() ([32]byte, error) {
	return _Bridge.Contract.DEPOSITTYPEHASH(&_Bridge.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Bridge *BridgeCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Bridge.contract.Call(opts, out, "DOMAIN_SEPARATOR")
	return *ret0, err
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Bridge *BridgeSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Bridge.Contract.DOMAINSEPARATOR(&_Bridge.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Bridge *BridgeCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Bridge.Contract.DOMAINSEPARATOR(&_Bridge.CallOpts)
}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256)
func (_Bridge *BridgeCaller) GetChainId(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Bridge.contract.Call(opts, out, "getChainId")
	return *ret0, err
}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256)
func (_Bridge *BridgeSession) GetChainId() (*big.Int, error) {
	return _Bridge.Contract.GetChainId(&_Bridge.CallOpts)
}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256)
func (_Bridge *BridgeCallerSession) GetChainId() (*big.Int, error) {
	return _Bridge.Contract.GetChainId(&_Bridge.CallOpts)
}

// TimeLockContract is a free data retrieval call binding the contract method 0x2f660362.
//
// Solidity: function timeLockContract() view returns(address)
func (_Bridge *BridgeCaller) TimeLockContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Bridge.contract.Call(opts, out, "timeLockContract")
	return *ret0, err
}

// TimeLockContract is a free data retrieval call binding the contract method 0x2f660362.
//
// Solidity: function timeLockContract() view returns(address)
func (_Bridge *BridgeSession) TimeLockContract() (common.Address, error) {
	return _Bridge.Contract.TimeLockContract(&_Bridge.CallOpts)
}

// TimeLockContract is a free data retrieval call binding the contract method 0x2f660362.
//
// Solidity: function timeLockContract() view returns(address)
func (_Bridge *BridgeCallerSession) TimeLockContract() (common.Address, error) {
	return _Bridge.Contract.TimeLockContract(&_Bridge.CallOpts)
}

// TokenInfos is a free data retrieval call binding the contract method 0x90f1f17b.
//
// Solidity: function tokenInfos(bytes32 ) view returns(address addr, uint8 lockType)
func (_Bridge *BridgeCaller) TokenInfos(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Addr     common.Address
	LockType uint8
}, error) {
	ret := new(struct {
		Addr     common.Address
		LockType uint8
	})
	out := ret
	err := _Bridge.contract.Call(opts, out, "tokenInfos", arg0)
	return *ret, err
}

// TokenInfos is a free data retrieval call binding the contract method 0x90f1f17b.
//
// Solidity: function tokenInfos(bytes32 ) view returns(address addr, uint8 lockType)
func (_Bridge *BridgeSession) TokenInfos(arg0 [32]byte) (struct {
	Addr     common.Address
	LockType uint8
}, error) {
	return _Bridge.Contract.TokenInfos(&_Bridge.CallOpts, arg0)
}

// TokenInfos is a free data retrieval call binding the contract method 0x90f1f17b.
//
// Solidity: function tokenInfos(bytes32 ) view returns(address addr, uint8 lockType)
func (_Bridge *BridgeCallerSession) TokenInfos(arg0 [32]byte) (struct {
	Addr     common.Address
	LockType uint8
}, error) {
	return _Bridge.Contract.TokenInfos(&_Bridge.CallOpts, arg0)
}

// ValidatorSet is a free data retrieval call binding the contract method 0x9426e226.
//
// Solidity: function validatorSet() view returns(address)
func (_Bridge *BridgeCaller) ValidatorSet(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Bridge.contract.Call(opts, out, "validatorSet")
	return *ret0, err
}

// ValidatorSet is a free data retrieval call binding the contract method 0x9426e226.
//
// Solidity: function validatorSet() view returns(address)
func (_Bridge *BridgeSession) ValidatorSet() (common.Address, error) {
	return _Bridge.Contract.ValidatorSet(&_Bridge.CallOpts)
}

// ValidatorSet is a free data retrieval call binding the contract method 0x9426e226.
//
// Solidity: function validatorSet() view returns(address)
func (_Bridge *BridgeCallerSession) ValidatorSet() (common.Address, error) {
	return _Bridge.Contract.ValidatorSet(&_Bridge.CallOpts)
}

// AddToken is a paid mutator transaction binding the contract method 0xb5949091.
//
// Solidity: function addToken(bytes32 symbol, address tokenAddr, uint8 lockType) returns()
func (_Bridge *BridgeTransactor) AddToken(opts *bind.TransactOpts, symbol [32]byte, tokenAddr common.Address, lockType uint8) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "addToken", symbol, tokenAddr, lockType)
}

// AddToken is a paid mutator transaction binding the contract method 0xb5949091.
//
// Solidity: function addToken(bytes32 symbol, address tokenAddr, uint8 lockType) returns()
func (_Bridge *BridgeSession) AddToken(symbol [32]byte, tokenAddr common.Address, lockType uint8) (*types.Transaction, error) {
	return _Bridge.Contract.AddToken(&_Bridge.TransactOpts, symbol, tokenAddr, lockType)
}

// AddToken is a paid mutator transaction binding the contract method 0xb5949091.
//
// Solidity: function addToken(bytes32 symbol, address tokenAddr, uint8 lockType) returns()
func (_Bridge *BridgeTransactorSession) AddToken(symbol [32]byte, tokenAddr common.Address, lockType uint8) (*types.Transaction, error) {
	return _Bridge.Contract.AddToken(&_Bridge.TransactOpts, symbol, tokenAddr, lockType)
}

// Deposit is a paid mutator transaction binding the contract method 0xa9d620ee.
//
// Solidity: function deposit(bytes32 token, uint256 destChainId, bytes32 recipient, uint256 amount) returns()
func (_Bridge *BridgeTransactor) Deposit(opts *bind.TransactOpts, token [32]byte, destChainId *big.Int, recipient [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "deposit", token, destChainId, recipient, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xa9d620ee.
//
// Solidity: function deposit(bytes32 token, uint256 destChainId, bytes32 recipient, uint256 amount) returns()
func (_Bridge *BridgeSession) Deposit(token [32]byte, destChainId *big.Int, recipient [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.Deposit(&_Bridge.TransactOpts, token, destChainId, recipient, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xa9d620ee.
//
// Solidity: function deposit(bytes32 token, uint256 destChainId, bytes32 recipient, uint256 amount) returns()
func (_Bridge *BridgeTransactorSession) Deposit(token [32]byte, destChainId *big.Int, recipient [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.Deposit(&_Bridge.TransactOpts, token, destChainId, recipient, amount)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address validatorSet_, address timeLockContract_) returns()
func (_Bridge *BridgeTransactor) Initialize(opts *bind.TransactOpts, validatorSet_ common.Address, timeLockContract_ common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "initialize", validatorSet_, timeLockContract_)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address validatorSet_, address timeLockContract_) returns()
func (_Bridge *BridgeSession) Initialize(validatorSet_ common.Address, timeLockContract_ common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.Initialize(&_Bridge.TransactOpts, validatorSet_, timeLockContract_)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address validatorSet_, address timeLockContract_) returns()
func (_Bridge *BridgeTransactorSession) Initialize(validatorSet_ common.Address, timeLockContract_ common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.Initialize(&_Bridge.TransactOpts, validatorSet_, timeLockContract_)
}

// RemoveToken is a paid mutator transaction binding the contract method 0x224e7462.
//
// Solidity: function removeToken(bytes32 symbol) returns()
func (_Bridge *BridgeTransactor) RemoveToken(opts *bind.TransactOpts, symbol [32]byte) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "removeToken", symbol)
}

// RemoveToken is a paid mutator transaction binding the contract method 0x224e7462.
//
// Solidity: function removeToken(bytes32 symbol) returns()
func (_Bridge *BridgeSession) RemoveToken(symbol [32]byte) (*types.Transaction, error) {
	return _Bridge.Contract.RemoveToken(&_Bridge.TransactOpts, symbol)
}

// RemoveToken is a paid mutator transaction binding the contract method 0x224e7462.
//
// Solidity: function removeToken(bytes32 symbol) returns()
func (_Bridge *BridgeTransactorSession) RemoveToken(symbol [32]byte) (*types.Transaction, error) {
	return _Bridge.Contract.RemoveToken(&_Bridge.TransactOpts, symbol)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51614f7f.
//
// Solidity: function withdraw(uint256 sourceChainId, uint256 depositId, bytes32 depositor, address recipient, bytes32 token, uint256 amount, bytes signatures) returns()
func (_Bridge *BridgeTransactor) Withdraw(opts *bind.TransactOpts, sourceChainId *big.Int, depositId *big.Int, depositor [32]byte, recipient common.Address, token [32]byte, amount *big.Int, signatures []byte) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "withdraw", sourceChainId, depositId, depositor, recipient, token, amount, signatures)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51614f7f.
//
// Solidity: function withdraw(uint256 sourceChainId, uint256 depositId, bytes32 depositor, address recipient, bytes32 token, uint256 amount, bytes signatures) returns()
func (_Bridge *BridgeSession) Withdraw(sourceChainId *big.Int, depositId *big.Int, depositor [32]byte, recipient common.Address, token [32]byte, amount *big.Int, signatures []byte) (*types.Transaction, error) {
	return _Bridge.Contract.Withdraw(&_Bridge.TransactOpts, sourceChainId, depositId, depositor, recipient, token, amount, signatures)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51614f7f.
//
// Solidity: function withdraw(uint256 sourceChainId, uint256 depositId, bytes32 depositor, address recipient, bytes32 token, uint256 amount, bytes signatures) returns()
func (_Bridge *BridgeTransactorSession) Withdraw(sourceChainId *big.Int, depositId *big.Int, depositor [32]byte, recipient common.Address, token [32]byte, amount *big.Int, signatures []byte) (*types.Transaction, error) {
	return _Bridge.Contract.Withdraw(&_Bridge.TransactOpts, sourceChainId, depositId, depositor, recipient, token, amount, signatures)
}

// BridgeDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the Bridge contract.
type BridgeDepositedIterator struct {
	Event *BridgeDeposited // Event containing the contract specifics and raw log

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
func (it *BridgeDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeDeposited)
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
		it.Event = new(BridgeDeposited)
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
func (it *BridgeDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeDeposited represents a Deposited event raised by the Bridge contract.
type BridgeDeposited struct {
	DestChainId *big.Int
	DepositId   *big.Int
	Depositor   common.Address
	Recipient   [32]byte
	Token       [32]byte
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x6d66e2e035c394faa164b4c1d2a539ae97944ce0d276b372042d921c2b951f11.
//
// Solidity: event Deposited(uint256 destChainId, uint256 depositId, address depositor, bytes32 recipient, bytes32 token, uint256 amount)
func (_Bridge *BridgeFilterer) FilterDeposited(opts *bind.FilterOpts) (*BridgeDepositedIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "Deposited")
	if err != nil {
		return nil, err
	}
	return &BridgeDepositedIterator{contract: _Bridge.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x6d66e2e035c394faa164b4c1d2a539ae97944ce0d276b372042d921c2b951f11.
//
// Solidity: event Deposited(uint256 destChainId, uint256 depositId, address depositor, bytes32 recipient, bytes32 token, uint256 amount)
func (_Bridge *BridgeFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *BridgeDeposited) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "Deposited")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeDeposited)
				if err := _Bridge.contract.UnpackLog(event, "Deposited", log); err != nil {
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

// ParseDeposited is a log parse operation binding the contract event 0x6d66e2e035c394faa164b4c1d2a539ae97944ce0d276b372042d921c2b951f11.
//
// Solidity: event Deposited(uint256 destChainId, uint256 depositId, address depositor, bytes32 recipient, bytes32 token, uint256 amount)
func (_Bridge *BridgeFilterer) ParseDeposited(log types.Log) (*BridgeDeposited, error) {
	event := new(BridgeDeposited)
	if err := _Bridge.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BridgeTokenAddedIterator is returned from FilterTokenAdded and is used to iterate over the raw logs and unpacked data for TokenAdded events raised by the Bridge contract.
type BridgeTokenAddedIterator struct {
	Event *BridgeTokenAdded // Event containing the contract specifics and raw log

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
func (it *BridgeTokenAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeTokenAdded)
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
		it.Event = new(BridgeTokenAdded)
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
func (it *BridgeTokenAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeTokenAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeTokenAdded represents a TokenAdded event raised by the Bridge contract.
type BridgeTokenAdded struct {
	Symbol    [32]byte
	TokenAddr common.Address
	Locktype  uint8
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTokenAdded is a free log retrieval operation binding the contract event 0xd9c6008de31d775aae4523eb2e6e6565ea965cabb3eac06153885621c3ddb035.
//
// Solidity: event TokenAdded(bytes32 symbol, address tokenAddr, uint8 locktype)
func (_Bridge *BridgeFilterer) FilterTokenAdded(opts *bind.FilterOpts) (*BridgeTokenAddedIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "TokenAdded")
	if err != nil {
		return nil, err
	}
	return &BridgeTokenAddedIterator{contract: _Bridge.contract, event: "TokenAdded", logs: logs, sub: sub}, nil
}

// WatchTokenAdded is a free log subscription operation binding the contract event 0xd9c6008de31d775aae4523eb2e6e6565ea965cabb3eac06153885621c3ddb035.
//
// Solidity: event TokenAdded(bytes32 symbol, address tokenAddr, uint8 locktype)
func (_Bridge *BridgeFilterer) WatchTokenAdded(opts *bind.WatchOpts, sink chan<- *BridgeTokenAdded) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "TokenAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeTokenAdded)
				if err := _Bridge.contract.UnpackLog(event, "TokenAdded", log); err != nil {
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

// ParseTokenAdded is a log parse operation binding the contract event 0xd9c6008de31d775aae4523eb2e6e6565ea965cabb3eac06153885621c3ddb035.
//
// Solidity: event TokenAdded(bytes32 symbol, address tokenAddr, uint8 locktype)
func (_Bridge *BridgeFilterer) ParseTokenAdded(log types.Log) (*BridgeTokenAdded, error) {
	event := new(BridgeTokenAdded)
	if err := _Bridge.contract.UnpackLog(event, "TokenAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BridgeTokenRemovedIterator is returned from FilterTokenRemoved and is used to iterate over the raw logs and unpacked data for TokenRemoved events raised by the Bridge contract.
type BridgeTokenRemovedIterator struct {
	Event *BridgeTokenRemoved // Event containing the contract specifics and raw log

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
func (it *BridgeTokenRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeTokenRemoved)
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
		it.Event = new(BridgeTokenRemoved)
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
func (it *BridgeTokenRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeTokenRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeTokenRemoved represents a TokenRemoved event raised by the Bridge contract.
type BridgeTokenRemoved struct {
	Symbol [32]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTokenRemoved is a free log retrieval operation binding the contract event 0x27b51627532fe8cd0df946e5a9590efb06cc367d3d611c1e4084c7e1dccb3ab4.
//
// Solidity: event TokenRemoved(bytes32 symbol)
func (_Bridge *BridgeFilterer) FilterTokenRemoved(opts *bind.FilterOpts) (*BridgeTokenRemovedIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "TokenRemoved")
	if err != nil {
		return nil, err
	}
	return &BridgeTokenRemovedIterator{contract: _Bridge.contract, event: "TokenRemoved", logs: logs, sub: sub}, nil
}

// WatchTokenRemoved is a free log subscription operation binding the contract event 0x27b51627532fe8cd0df946e5a9590efb06cc367d3d611c1e4084c7e1dccb3ab4.
//
// Solidity: event TokenRemoved(bytes32 symbol)
func (_Bridge *BridgeFilterer) WatchTokenRemoved(opts *bind.WatchOpts, sink chan<- *BridgeTokenRemoved) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "TokenRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeTokenRemoved)
				if err := _Bridge.contract.UnpackLog(event, "TokenRemoved", log); err != nil {
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

// ParseTokenRemoved is a log parse operation binding the contract event 0x27b51627532fe8cd0df946e5a9590efb06cc367d3d611c1e4084c7e1dccb3ab4.
//
// Solidity: event TokenRemoved(bytes32 symbol)
func (_Bridge *BridgeFilterer) ParseTokenRemoved(log types.Log) (*BridgeTokenRemoved, error) {
	event := new(BridgeTokenRemoved)
	if err := _Bridge.contract.UnpackLog(event, "TokenRemoved", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BridgeWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Bridge contract.
type BridgeWithdrawIterator struct {
	Event *BridgeWithdraw // Event containing the contract specifics and raw log

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
func (it *BridgeWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeWithdraw)
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
		it.Event = new(BridgeWithdraw)
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
func (it *BridgeWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeWithdraw represents a Withdraw event raised by the Bridge contract.
type BridgeWithdraw struct {
	SourceChainId *big.Int
	DepositId     *big.Int
	Depositor     [32]byte
	Recipient     common.Address
	Token         [32]byte
	Amount        *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x30b3ee8e1271b2c5575cfcd522428ebe2a38a57eece3fd08c5aa17a996671d01.
//
// Solidity: event Withdraw(uint256 sourceChainId, uint256 depositId, bytes32 depositor, address recipient, bytes32 token, uint256 amount)
func (_Bridge *BridgeFilterer) FilterWithdraw(opts *bind.FilterOpts) (*BridgeWithdrawIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return &BridgeWithdrawIterator{contract: _Bridge.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x30b3ee8e1271b2c5575cfcd522428ebe2a38a57eece3fd08c5aa17a996671d01.
//
// Solidity: event Withdraw(uint256 sourceChainId, uint256 depositId, bytes32 depositor, address recipient, bytes32 token, uint256 amount)
func (_Bridge *BridgeFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *BridgeWithdraw) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeWithdraw)
				if err := _Bridge.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x30b3ee8e1271b2c5575cfcd522428ebe2a38a57eece3fd08c5aa17a996671d01.
//
// Solidity: event Withdraw(uint256 sourceChainId, uint256 depositId, bytes32 depositor, address recipient, bytes32 token, uint256 amount)
func (_Bridge *BridgeFilterer) ParseWithdraw(log types.Log) (*BridgeWithdraw, error) {
	event := new(BridgeWithdraw)
	if err := _Bridge.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	return event, nil
}

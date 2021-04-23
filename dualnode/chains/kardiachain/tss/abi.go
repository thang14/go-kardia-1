// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tss

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

// TssABI is the input ABI used to generate the binding from.
const TssABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"symbol\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumTss.TokenLockType\",\"name\":\"locktype\",\"type\":\"uint8\"}],\"name\":\"TokenAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"symbol\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"TokenRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"VaultChainAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"VaultChainEdited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"VaultChainRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"pubKey\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"members\",\"type\":\"bytes[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"chainIds\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"contractAddresses\",\"type\":\"address[]\"}],\"name\":\"VaultCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"pubKey\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"members\",\"type\":\"bytes[]\"}],\"name\":\"VaultUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"chainIds\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"chainContracts\",\"type\":\"address[]\"}],\"name\":\"addChains\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"symbol\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"enumTss.TokenLockType\",\"name\":\"tokenLockType\",\"type\":\"uint8\"}],\"name\":\"addToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"pubKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"members\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256[]\",\"name\":\"chainIds\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"chainContracts\",\"type\":\"address[]\"}],\"name\":\"createVault\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"getTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"tokenAddrs\",\"type\":\"address[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"symbols\",\"type\":\"bytes32[]\"},{\"internalType\":\"enumTss.TokenLockType[]\",\"name\":\"locktypes\",\"type\":\"uint8[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"}],\"name\":\"getVault\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"pubKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"members\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256[]\",\"name\":\"chainIds\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"chainContracts\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"idx\",\"type\":\"uint256[]\"}],\"name\":\"removeChains\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"removeToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenInfos\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"symbol\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"enumTss.TokenLockType\",\"name\":\"locktype\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"pubKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"members\",\"type\":\"bytes[]\"}],\"name\":\"updateVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vaultCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"vaults\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"pubKey\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]=======lib/Context.sol:Context=======[]=======lib/Ownable.sol:Ownable=======[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Tss is an auto generated Go binding around an Ethereum contract.
type Tss struct {
	TssCaller     // Read-only binding to the contract
	TssTransactor // Write-only binding to the contract
	TssFilterer   // Log filterer for contract events
}

// TssCaller is an auto generated read-only Go binding around an Ethereum contract.
type TssCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TssTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TssTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TssFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TssFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TssSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TssSession struct {
	Contract     *Tss              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TssCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TssCallerSession struct {
	Contract *TssCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TssTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TssTransactorSession struct {
	Contract     *TssTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TssRaw is an auto generated low-level Go binding around an Ethereum contract.
type TssRaw struct {
	Contract *Tss // Generic contract binding to access the raw methods on
}

// TssCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TssCallerRaw struct {
	Contract *TssCaller // Generic read-only contract binding to access the raw methods on
}

// TssTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TssTransactorRaw struct {
	Contract *TssTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTss creates a new instance of Tss, bound to a specific deployed contract.
func NewTss(address common.Address, backend bind.ContractBackend) (*Tss, error) {
	contract, err := bindTss(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Tss{TssCaller: TssCaller{contract: contract}, TssTransactor: TssTransactor{contract: contract}, TssFilterer: TssFilterer{contract: contract}}, nil
}

// NewTssCaller creates a new read-only instance of Tss, bound to a specific deployed contract.
func NewTssCaller(address common.Address, caller bind.ContractCaller) (*TssCaller, error) {
	contract, err := bindTss(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TssCaller{contract: contract}, nil
}

// NewTssTransactor creates a new write-only instance of Tss, bound to a specific deployed contract.
func NewTssTransactor(address common.Address, transactor bind.ContractTransactor) (*TssTransactor, error) {
	contract, err := bindTss(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TssTransactor{contract: contract}, nil
}

// NewTssFilterer creates a new log filterer instance of Tss, bound to a specific deployed contract.
func NewTssFilterer(address common.Address, filterer bind.ContractFilterer) (*TssFilterer, error) {
	contract, err := bindTss(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TssFilterer{contract: contract}, nil
}

// bindTss binds a generic wrapper to an already deployed contract.
func bindTss(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TssABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tss *TssRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Tss.Contract.TssCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tss *TssRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tss.Contract.TssTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tss *TssRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tss.Contract.TssTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tss *TssCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Tss.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tss *TssTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tss.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tss *TssTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tss.Contract.contract.Transact(opts, method, params...)
}

// GetTokens is a free data retrieval call binding the contract method 0x494cfc6c.
//
// Solidity: function getTokens(uint256 vaultId, uint256 chainId) view returns(address[] tokenAddrs, bytes32[] symbols, uint8[] locktypes)
func (_Tss *TssCaller) GetTokens(opts *bind.CallOpts, vaultId *big.Int, chainId *big.Int) (struct {
	TokenAddrs []common.Address
	Symbols    [][32]byte
	Locktypes  []uint8
}, error) {
	ret := new(struct {
		TokenAddrs []common.Address
		Symbols    [][32]byte
		Locktypes  []uint8
	})
	out := ret
	err := _Tss.contract.Call(opts, out, "getTokens", vaultId, chainId)
	return *ret, err
}

// GetTokens is a free data retrieval call binding the contract method 0x494cfc6c.
//
// Solidity: function getTokens(uint256 vaultId, uint256 chainId) view returns(address[] tokenAddrs, bytes32[] symbols, uint8[] locktypes)
func (_Tss *TssSession) GetTokens(vaultId *big.Int, chainId *big.Int) (struct {
	TokenAddrs []common.Address
	Symbols    [][32]byte
	Locktypes  []uint8
}, error) {
	return _Tss.Contract.GetTokens(&_Tss.CallOpts, vaultId, chainId)
}

// GetTokens is a free data retrieval call binding the contract method 0x494cfc6c.
//
// Solidity: function getTokens(uint256 vaultId, uint256 chainId) view returns(address[] tokenAddrs, bytes32[] symbols, uint8[] locktypes)
func (_Tss *TssCallerSession) GetTokens(vaultId *big.Int, chainId *big.Int) (struct {
	TokenAddrs []common.Address
	Symbols    [][32]byte
	Locktypes  []uint8
}, error) {
	return _Tss.Contract.GetTokens(&_Tss.CallOpts, vaultId, chainId)
}

// GetVault is a free data retrieval call binding the contract method 0x9403b634.
//
// Solidity: function getVault(uint256 vaultId) view returns(bytes pubKey, bytes[] members, uint256[] chainIds, address[] chainContracts)
func (_Tss *TssCaller) GetVault(opts *bind.CallOpts, vaultId *big.Int) (struct {
	PubKey         []byte
	Members        [][]byte
	ChainIds       []*big.Int
	ChainContracts []common.Address
}, error) {
	ret := new(struct {
		PubKey         []byte
		Members        [][]byte
		ChainIds       []*big.Int
		ChainContracts []common.Address
	})
	out := ret
	err := _Tss.contract.Call(opts, out, "getVault", vaultId)
	return *ret, err
}

// GetVault is a free data retrieval call binding the contract method 0x9403b634.
//
// Solidity: function getVault(uint256 vaultId) view returns(bytes pubKey, bytes[] members, uint256[] chainIds, address[] chainContracts)
func (_Tss *TssSession) GetVault(vaultId *big.Int) (struct {
	PubKey         []byte
	Members        [][]byte
	ChainIds       []*big.Int
	ChainContracts []common.Address
}, error) {
	return _Tss.Contract.GetVault(&_Tss.CallOpts, vaultId)
}

// GetVault is a free data retrieval call binding the contract method 0x9403b634.
//
// Solidity: function getVault(uint256 vaultId) view returns(bytes pubKey, bytes[] members, uint256[] chainIds, address[] chainContracts)
func (_Tss *TssCallerSession) GetVault(vaultId *big.Int) (struct {
	PubKey         []byte
	Members        [][]byte
	ChainIds       []*big.Int
	ChainContracts []common.Address
}, error) {
	return _Tss.Contract.GetVault(&_Tss.CallOpts, vaultId)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Tss *TssCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Tss.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Tss *TssSession) Owner() (common.Address, error) {
	return _Tss.Contract.Owner(&_Tss.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Tss *TssCallerSession) Owner() (common.Address, error) {
	return _Tss.Contract.Owner(&_Tss.CallOpts)
}

// TokenInfos is a free data retrieval call binding the contract method 0xfe8103f6.
//
// Solidity: function tokenInfos(uint256 , uint256 , uint256 ) view returns(bytes32 symbol, address tokenAddress, uint8 locktype)
func (_Tss *TssCaller) TokenInfos(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int, arg2 *big.Int) (struct {
	Symbol       [32]byte
	TokenAddress common.Address
	Locktype     uint8
}, error) {
	ret := new(struct {
		Symbol       [32]byte
		TokenAddress common.Address
		Locktype     uint8
	})
	out := ret
	err := _Tss.contract.Call(opts, out, "tokenInfos", arg0, arg1, arg2)
	return *ret, err
}

// TokenInfos is a free data retrieval call binding the contract method 0xfe8103f6.
//
// Solidity: function tokenInfos(uint256 , uint256 , uint256 ) view returns(bytes32 symbol, address tokenAddress, uint8 locktype)
func (_Tss *TssSession) TokenInfos(arg0 *big.Int, arg1 *big.Int, arg2 *big.Int) (struct {
	Symbol       [32]byte
	TokenAddress common.Address
	Locktype     uint8
}, error) {
	return _Tss.Contract.TokenInfos(&_Tss.CallOpts, arg0, arg1, arg2)
}

// TokenInfos is a free data retrieval call binding the contract method 0xfe8103f6.
//
// Solidity: function tokenInfos(uint256 , uint256 , uint256 ) view returns(bytes32 symbol, address tokenAddress, uint8 locktype)
func (_Tss *TssCallerSession) TokenInfos(arg0 *big.Int, arg1 *big.Int, arg2 *big.Int) (struct {
	Symbol       [32]byte
	TokenAddress common.Address
	Locktype     uint8
}, error) {
	return _Tss.Contract.TokenInfos(&_Tss.CallOpts, arg0, arg1, arg2)
}

// VaultCount is a free data retrieval call binding the contract method 0xa7c6a100.
//
// Solidity: function vaultCount() view returns(uint256)
func (_Tss *TssCaller) VaultCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Tss.contract.Call(opts, out, "vaultCount")
	return *ret0, err
}

// VaultCount is a free data retrieval call binding the contract method 0xa7c6a100.
//
// Solidity: function vaultCount() view returns(uint256)
func (_Tss *TssSession) VaultCount() (*big.Int, error) {
	return _Tss.Contract.VaultCount(&_Tss.CallOpts)
}

// VaultCount is a free data retrieval call binding the contract method 0xa7c6a100.
//
// Solidity: function vaultCount() view returns(uint256)
func (_Tss *TssCallerSession) VaultCount() (*big.Int, error) {
	return _Tss.Contract.VaultCount(&_Tss.CallOpts)
}

// Vaults is a free data retrieval call binding the contract method 0x8c64ea4a.
//
// Solidity: function vaults(uint256 ) view returns(uint256 id, uint256 blockHeight, bytes pubKey)
func (_Tss *TssCaller) Vaults(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id          *big.Int
	BlockHeight *big.Int
	PubKey      []byte
}, error) {
	ret := new(struct {
		Id          *big.Int
		BlockHeight *big.Int
		PubKey      []byte
	})
	out := ret
	err := _Tss.contract.Call(opts, out, "vaults", arg0)
	return *ret, err
}

// Vaults is a free data retrieval call binding the contract method 0x8c64ea4a.
//
// Solidity: function vaults(uint256 ) view returns(uint256 id, uint256 blockHeight, bytes pubKey)
func (_Tss *TssSession) Vaults(arg0 *big.Int) (struct {
	Id          *big.Int
	BlockHeight *big.Int
	PubKey      []byte
}, error) {
	return _Tss.Contract.Vaults(&_Tss.CallOpts, arg0)
}

// Vaults is a free data retrieval call binding the contract method 0x8c64ea4a.
//
// Solidity: function vaults(uint256 ) view returns(uint256 id, uint256 blockHeight, bytes pubKey)
func (_Tss *TssCallerSession) Vaults(arg0 *big.Int) (struct {
	Id          *big.Int
	BlockHeight *big.Int
	PubKey      []byte
}, error) {
	return _Tss.Contract.Vaults(&_Tss.CallOpts, arg0)
}

// AddChains is a paid mutator transaction binding the contract method 0xbd2df917.
//
// Solidity: function addChains(uint256 vaultId, uint256[] chainIds, address[] chainContracts) returns()
func (_Tss *TssTransactor) AddChains(opts *bind.TransactOpts, vaultId *big.Int, chainIds []*big.Int, chainContracts []common.Address) (*types.Transaction, error) {
	return _Tss.contract.Transact(opts, "addChains", vaultId, chainIds, chainContracts)
}

// AddChains is a paid mutator transaction binding the contract method 0xbd2df917.
//
// Solidity: function addChains(uint256 vaultId, uint256[] chainIds, address[] chainContracts) returns()
func (_Tss *TssSession) AddChains(vaultId *big.Int, chainIds []*big.Int, chainContracts []common.Address) (*types.Transaction, error) {
	return _Tss.Contract.AddChains(&_Tss.TransactOpts, vaultId, chainIds, chainContracts)
}

// AddChains is a paid mutator transaction binding the contract method 0xbd2df917.
//
// Solidity: function addChains(uint256 vaultId, uint256[] chainIds, address[] chainContracts) returns()
func (_Tss *TssTransactorSession) AddChains(vaultId *big.Int, chainIds []*big.Int, chainContracts []common.Address) (*types.Transaction, error) {
	return _Tss.Contract.AddChains(&_Tss.TransactOpts, vaultId, chainIds, chainContracts)
}

// AddToken is a paid mutator transaction binding the contract method 0x22f00a17.
//
// Solidity: function addToken(uint256 vaultId, uint256 chainId, bytes32 symbol, address tokenAddress, uint8 tokenLockType) returns()
func (_Tss *TssTransactor) AddToken(opts *bind.TransactOpts, vaultId *big.Int, chainId *big.Int, symbol [32]byte, tokenAddress common.Address, tokenLockType uint8) (*types.Transaction, error) {
	return _Tss.contract.Transact(opts, "addToken", vaultId, chainId, symbol, tokenAddress, tokenLockType)
}

// AddToken is a paid mutator transaction binding the contract method 0x22f00a17.
//
// Solidity: function addToken(uint256 vaultId, uint256 chainId, bytes32 symbol, address tokenAddress, uint8 tokenLockType) returns()
func (_Tss *TssSession) AddToken(vaultId *big.Int, chainId *big.Int, symbol [32]byte, tokenAddress common.Address, tokenLockType uint8) (*types.Transaction, error) {
	return _Tss.Contract.AddToken(&_Tss.TransactOpts, vaultId, chainId, symbol, tokenAddress, tokenLockType)
}

// AddToken is a paid mutator transaction binding the contract method 0x22f00a17.
//
// Solidity: function addToken(uint256 vaultId, uint256 chainId, bytes32 symbol, address tokenAddress, uint8 tokenLockType) returns()
func (_Tss *TssTransactorSession) AddToken(vaultId *big.Int, chainId *big.Int, symbol [32]byte, tokenAddress common.Address, tokenLockType uint8) (*types.Transaction, error) {
	return _Tss.Contract.AddToken(&_Tss.TransactOpts, vaultId, chainId, symbol, tokenAddress, tokenLockType)
}

// CreateVault is a paid mutator transaction binding the contract method 0x3ff884dc.
//
// Solidity: function createVault(bytes pubKey, bytes[] members, uint256[] chainIds, address[] chainContracts) returns(uint256)
func (_Tss *TssTransactor) CreateVault(opts *bind.TransactOpts, pubKey []byte, members [][]byte, chainIds []*big.Int, chainContracts []common.Address) (*types.Transaction, error) {
	return _Tss.contract.Transact(opts, "createVault", pubKey, members, chainIds, chainContracts)
}

// CreateVault is a paid mutator transaction binding the contract method 0x3ff884dc.
//
// Solidity: function createVault(bytes pubKey, bytes[] members, uint256[] chainIds, address[] chainContracts) returns(uint256)
func (_Tss *TssSession) CreateVault(pubKey []byte, members [][]byte, chainIds []*big.Int, chainContracts []common.Address) (*types.Transaction, error) {
	return _Tss.Contract.CreateVault(&_Tss.TransactOpts, pubKey, members, chainIds, chainContracts)
}

// CreateVault is a paid mutator transaction binding the contract method 0x3ff884dc.
//
// Solidity: function createVault(bytes pubKey, bytes[] members, uint256[] chainIds, address[] chainContracts) returns(uint256)
func (_Tss *TssTransactorSession) CreateVault(pubKey []byte, members [][]byte, chainIds []*big.Int, chainContracts []common.Address) (*types.Transaction, error) {
	return _Tss.Contract.CreateVault(&_Tss.TransactOpts, pubKey, members, chainIds, chainContracts)
}

// RemoveChains is a paid mutator transaction binding the contract method 0xc9e17e6d.
//
// Solidity: function removeChains(uint256 vaultId, uint256[] idx) returns()
func (_Tss *TssTransactor) RemoveChains(opts *bind.TransactOpts, vaultId *big.Int, idx []*big.Int) (*types.Transaction, error) {
	return _Tss.contract.Transact(opts, "removeChains", vaultId, idx)
}

// RemoveChains is a paid mutator transaction binding the contract method 0xc9e17e6d.
//
// Solidity: function removeChains(uint256 vaultId, uint256[] idx) returns()
func (_Tss *TssSession) RemoveChains(vaultId *big.Int, idx []*big.Int) (*types.Transaction, error) {
	return _Tss.Contract.RemoveChains(&_Tss.TransactOpts, vaultId, idx)
}

// RemoveChains is a paid mutator transaction binding the contract method 0xc9e17e6d.
//
// Solidity: function removeChains(uint256 vaultId, uint256[] idx) returns()
func (_Tss *TssTransactorSession) RemoveChains(vaultId *big.Int, idx []*big.Int) (*types.Transaction, error) {
	return _Tss.Contract.RemoveChains(&_Tss.TransactOpts, vaultId, idx)
}

// RemoveToken is a paid mutator transaction binding the contract method 0x7510199b.
//
// Solidity: function removeToken(uint256 vaultId, uint256 chainId, uint256 index) returns()
func (_Tss *TssTransactor) RemoveToken(opts *bind.TransactOpts, vaultId *big.Int, chainId *big.Int, index *big.Int) (*types.Transaction, error) {
	return _Tss.contract.Transact(opts, "removeToken", vaultId, chainId, index)
}

// RemoveToken is a paid mutator transaction binding the contract method 0x7510199b.
//
// Solidity: function removeToken(uint256 vaultId, uint256 chainId, uint256 index) returns()
func (_Tss *TssSession) RemoveToken(vaultId *big.Int, chainId *big.Int, index *big.Int) (*types.Transaction, error) {
	return _Tss.Contract.RemoveToken(&_Tss.TransactOpts, vaultId, chainId, index)
}

// RemoveToken is a paid mutator transaction binding the contract method 0x7510199b.
//
// Solidity: function removeToken(uint256 vaultId, uint256 chainId, uint256 index) returns()
func (_Tss *TssTransactorSession) RemoveToken(vaultId *big.Int, chainId *big.Int, index *big.Int) (*types.Transaction, error) {
	return _Tss.Contract.RemoveToken(&_Tss.TransactOpts, vaultId, chainId, index)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Tss *TssTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tss.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Tss *TssSession) RenounceOwnership() (*types.Transaction, error) {
	return _Tss.Contract.RenounceOwnership(&_Tss.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Tss *TssTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Tss.Contract.RenounceOwnership(&_Tss.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Tss *TssTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Tss.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Tss *TssSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Tss.Contract.TransferOwnership(&_Tss.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Tss *TssTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Tss.Contract.TransferOwnership(&_Tss.TransactOpts, newOwner)
}

// UpdateVault is a paid mutator transaction binding the contract method 0x9943755d.
//
// Solidity: function updateVault(uint256 id, bytes pubKey, bytes[] members) returns()
func (_Tss *TssTransactor) UpdateVault(opts *bind.TransactOpts, id *big.Int, pubKey []byte, members [][]byte) (*types.Transaction, error) {
	return _Tss.contract.Transact(opts, "updateVault", id, pubKey, members)
}

// UpdateVault is a paid mutator transaction binding the contract method 0x9943755d.
//
// Solidity: function updateVault(uint256 id, bytes pubKey, bytes[] members) returns()
func (_Tss *TssSession) UpdateVault(id *big.Int, pubKey []byte, members [][]byte) (*types.Transaction, error) {
	return _Tss.Contract.UpdateVault(&_Tss.TransactOpts, id, pubKey, members)
}

// UpdateVault is a paid mutator transaction binding the contract method 0x9943755d.
//
// Solidity: function updateVault(uint256 id, bytes pubKey, bytes[] members) returns()
func (_Tss *TssTransactorSession) UpdateVault(id *big.Int, pubKey []byte, members [][]byte) (*types.Transaction, error) {
	return _Tss.Contract.UpdateVault(&_Tss.TransactOpts, id, pubKey, members)
}

// TssOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Tss contract.
type TssOwnershipTransferredIterator struct {
	Event *TssOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TssOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TssOwnershipTransferred)
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
		it.Event = new(TssOwnershipTransferred)
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
func (it *TssOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TssOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TssOwnershipTransferred represents a OwnershipTransferred event raised by the Tss contract.
type TssOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Tss *TssFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TssOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Tss.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TssOwnershipTransferredIterator{contract: _Tss.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Tss *TssFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TssOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Tss.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TssOwnershipTransferred)
				if err := _Tss.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Tss *TssFilterer) ParseOwnershipTransferred(log types.Log) (*TssOwnershipTransferred, error) {
	event := new(TssOwnershipTransferred)
	if err := _Tss.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TssTokenAddedIterator is returned from FilterTokenAdded and is used to iterate over the raw logs and unpacked data for TokenAdded events raised by the Tss contract.
type TssTokenAddedIterator struct {
	Event *TssTokenAdded // Event containing the contract specifics and raw log

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
func (it *TssTokenAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TssTokenAdded)
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
		it.Event = new(TssTokenAdded)
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
func (it *TssTokenAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TssTokenAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TssTokenAdded represents a TokenAdded event raised by the Tss contract.
type TssTokenAdded struct {
	VaultId   *big.Int
	ChainId   *big.Int
	Symbol    [32]byte
	TokenAddr common.Address
	Locktype  uint8
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTokenAdded is a free log retrieval operation binding the contract event 0x7210726ed9c60cabffe8ea7561ae74fc72774bcac094c9584de572b62514684f.
//
// Solidity: event TokenAdded(uint256 vaultId, uint256 chainId, bytes32 symbol, address tokenAddr, uint8 locktype)
func (_Tss *TssFilterer) FilterTokenAdded(opts *bind.FilterOpts) (*TssTokenAddedIterator, error) {

	logs, sub, err := _Tss.contract.FilterLogs(opts, "TokenAdded")
	if err != nil {
		return nil, err
	}
	return &TssTokenAddedIterator{contract: _Tss.contract, event: "TokenAdded", logs: logs, sub: sub}, nil
}

// WatchTokenAdded is a free log subscription operation binding the contract event 0x7210726ed9c60cabffe8ea7561ae74fc72774bcac094c9584de572b62514684f.
//
// Solidity: event TokenAdded(uint256 vaultId, uint256 chainId, bytes32 symbol, address tokenAddr, uint8 locktype)
func (_Tss *TssFilterer) WatchTokenAdded(opts *bind.WatchOpts, sink chan<- *TssTokenAdded) (event.Subscription, error) {

	logs, sub, err := _Tss.contract.WatchLogs(opts, "TokenAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TssTokenAdded)
				if err := _Tss.contract.UnpackLog(event, "TokenAdded", log); err != nil {
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

// ParseTokenAdded is a log parse operation binding the contract event 0x7210726ed9c60cabffe8ea7561ae74fc72774bcac094c9584de572b62514684f.
//
// Solidity: event TokenAdded(uint256 vaultId, uint256 chainId, bytes32 symbol, address tokenAddr, uint8 locktype)
func (_Tss *TssFilterer) ParseTokenAdded(log types.Log) (*TssTokenAdded, error) {
	event := new(TssTokenAdded)
	if err := _Tss.contract.UnpackLog(event, "TokenAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TssTokenRemovedIterator is returned from FilterTokenRemoved and is used to iterate over the raw logs and unpacked data for TokenRemoved events raised by the Tss contract.
type TssTokenRemovedIterator struct {
	Event *TssTokenRemoved // Event containing the contract specifics and raw log

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
func (it *TssTokenRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TssTokenRemoved)
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
		it.Event = new(TssTokenRemoved)
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
func (it *TssTokenRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TssTokenRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TssTokenRemoved represents a TokenRemoved event raised by the Tss contract.
type TssTokenRemoved struct {
	VaultId *big.Int
	ChainId *big.Int
	Symbol  [32]byte
	Index   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTokenRemoved is a free log retrieval operation binding the contract event 0x9bb313ef61286446ce0990aa915c2148137ff120dde4b6bffc4f9212925f7f3f.
//
// Solidity: event TokenRemoved(uint256 vaultId, uint256 chainId, bytes32 symbol, uint256 index)
func (_Tss *TssFilterer) FilterTokenRemoved(opts *bind.FilterOpts) (*TssTokenRemovedIterator, error) {

	logs, sub, err := _Tss.contract.FilterLogs(opts, "TokenRemoved")
	if err != nil {
		return nil, err
	}
	return &TssTokenRemovedIterator{contract: _Tss.contract, event: "TokenRemoved", logs: logs, sub: sub}, nil
}

// WatchTokenRemoved is a free log subscription operation binding the contract event 0x9bb313ef61286446ce0990aa915c2148137ff120dde4b6bffc4f9212925f7f3f.
//
// Solidity: event TokenRemoved(uint256 vaultId, uint256 chainId, bytes32 symbol, uint256 index)
func (_Tss *TssFilterer) WatchTokenRemoved(opts *bind.WatchOpts, sink chan<- *TssTokenRemoved) (event.Subscription, error) {

	logs, sub, err := _Tss.contract.WatchLogs(opts, "TokenRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TssTokenRemoved)
				if err := _Tss.contract.UnpackLog(event, "TokenRemoved", log); err != nil {
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

// ParseTokenRemoved is a log parse operation binding the contract event 0x9bb313ef61286446ce0990aa915c2148137ff120dde4b6bffc4f9212925f7f3f.
//
// Solidity: event TokenRemoved(uint256 vaultId, uint256 chainId, bytes32 symbol, uint256 index)
func (_Tss *TssFilterer) ParseTokenRemoved(log types.Log) (*TssTokenRemoved, error) {
	event := new(TssTokenRemoved)
	if err := _Tss.contract.UnpackLog(event, "TokenRemoved", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TssVaultChainAddedIterator is returned from FilterVaultChainAdded and is used to iterate over the raw logs and unpacked data for VaultChainAdded events raised by the Tss contract.
type TssVaultChainAddedIterator struct {
	Event *TssVaultChainAdded // Event containing the contract specifics and raw log

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
func (it *TssVaultChainAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TssVaultChainAdded)
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
		it.Event = new(TssVaultChainAdded)
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
func (it *TssVaultChainAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TssVaultChainAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TssVaultChainAdded represents a VaultChainAdded event raised by the Tss contract.
type TssVaultChainAdded struct {
	Id              *big.Int
	ChainId         *big.Int
	ContractAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterVaultChainAdded is a free log retrieval operation binding the contract event 0x0bf1d76ef9cf87088990b532cfb98cba32ded4542cd31463e339031526970c4f.
//
// Solidity: event VaultChainAdded(uint256 id, uint256 chainId, address contractAddress)
func (_Tss *TssFilterer) FilterVaultChainAdded(opts *bind.FilterOpts) (*TssVaultChainAddedIterator, error) {

	logs, sub, err := _Tss.contract.FilterLogs(opts, "VaultChainAdded")
	if err != nil {
		return nil, err
	}
	return &TssVaultChainAddedIterator{contract: _Tss.contract, event: "VaultChainAdded", logs: logs, sub: sub}, nil
}

// WatchVaultChainAdded is a free log subscription operation binding the contract event 0x0bf1d76ef9cf87088990b532cfb98cba32ded4542cd31463e339031526970c4f.
//
// Solidity: event VaultChainAdded(uint256 id, uint256 chainId, address contractAddress)
func (_Tss *TssFilterer) WatchVaultChainAdded(opts *bind.WatchOpts, sink chan<- *TssVaultChainAdded) (event.Subscription, error) {

	logs, sub, err := _Tss.contract.WatchLogs(opts, "VaultChainAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TssVaultChainAdded)
				if err := _Tss.contract.UnpackLog(event, "VaultChainAdded", log); err != nil {
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

// ParseVaultChainAdded is a log parse operation binding the contract event 0x0bf1d76ef9cf87088990b532cfb98cba32ded4542cd31463e339031526970c4f.
//
// Solidity: event VaultChainAdded(uint256 id, uint256 chainId, address contractAddress)
func (_Tss *TssFilterer) ParseVaultChainAdded(log types.Log) (*TssVaultChainAdded, error) {
	event := new(TssVaultChainAdded)
	if err := _Tss.contract.UnpackLog(event, "VaultChainAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TssVaultChainEditedIterator is returned from FilterVaultChainEdited and is used to iterate over the raw logs and unpacked data for VaultChainEdited events raised by the Tss contract.
type TssVaultChainEditedIterator struct {
	Event *TssVaultChainEdited // Event containing the contract specifics and raw log

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
func (it *TssVaultChainEditedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TssVaultChainEdited)
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
		it.Event = new(TssVaultChainEdited)
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
func (it *TssVaultChainEditedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TssVaultChainEditedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TssVaultChainEdited represents a VaultChainEdited event raised by the Tss contract.
type TssVaultChainEdited struct {
	Id              *big.Int
	ChainId         *big.Int
	ContractAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterVaultChainEdited is a free log retrieval operation binding the contract event 0x6ffad9fe8a51626151a33a417952d8ebe39ddc2cfaa6738c09d74fb62d3cc548.
//
// Solidity: event VaultChainEdited(uint256 id, uint256 chainId, address contractAddress)
func (_Tss *TssFilterer) FilterVaultChainEdited(opts *bind.FilterOpts) (*TssVaultChainEditedIterator, error) {

	logs, sub, err := _Tss.contract.FilterLogs(opts, "VaultChainEdited")
	if err != nil {
		return nil, err
	}
	return &TssVaultChainEditedIterator{contract: _Tss.contract, event: "VaultChainEdited", logs: logs, sub: sub}, nil
}

// WatchVaultChainEdited is a free log subscription operation binding the contract event 0x6ffad9fe8a51626151a33a417952d8ebe39ddc2cfaa6738c09d74fb62d3cc548.
//
// Solidity: event VaultChainEdited(uint256 id, uint256 chainId, address contractAddress)
func (_Tss *TssFilterer) WatchVaultChainEdited(opts *bind.WatchOpts, sink chan<- *TssVaultChainEdited) (event.Subscription, error) {

	logs, sub, err := _Tss.contract.WatchLogs(opts, "VaultChainEdited")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TssVaultChainEdited)
				if err := _Tss.contract.UnpackLog(event, "VaultChainEdited", log); err != nil {
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

// ParseVaultChainEdited is a log parse operation binding the contract event 0x6ffad9fe8a51626151a33a417952d8ebe39ddc2cfaa6738c09d74fb62d3cc548.
//
// Solidity: event VaultChainEdited(uint256 id, uint256 chainId, address contractAddress)
func (_Tss *TssFilterer) ParseVaultChainEdited(log types.Log) (*TssVaultChainEdited, error) {
	event := new(TssVaultChainEdited)
	if err := _Tss.contract.UnpackLog(event, "VaultChainEdited", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TssVaultChainRemovedIterator is returned from FilterVaultChainRemoved and is used to iterate over the raw logs and unpacked data for VaultChainRemoved events raised by the Tss contract.
type TssVaultChainRemovedIterator struct {
	Event *TssVaultChainRemoved // Event containing the contract specifics and raw log

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
func (it *TssVaultChainRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TssVaultChainRemoved)
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
		it.Event = new(TssVaultChainRemoved)
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
func (it *TssVaultChainRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TssVaultChainRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TssVaultChainRemoved represents a VaultChainRemoved event raised by the Tss contract.
type TssVaultChainRemoved struct {
	Id              *big.Int
	ChainId         *big.Int
	ContractAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterVaultChainRemoved is a free log retrieval operation binding the contract event 0xd9b1841df91b154dfe2cc2d10b346bb5298e38fc4e76a76e9d6ad05bceebbbbd.
//
// Solidity: event VaultChainRemoved(uint256 id, uint256 chainId, address contractAddress)
func (_Tss *TssFilterer) FilterVaultChainRemoved(opts *bind.FilterOpts) (*TssVaultChainRemovedIterator, error) {

	logs, sub, err := _Tss.contract.FilterLogs(opts, "VaultChainRemoved")
	if err != nil {
		return nil, err
	}
	return &TssVaultChainRemovedIterator{contract: _Tss.contract, event: "VaultChainRemoved", logs: logs, sub: sub}, nil
}

// WatchVaultChainRemoved is a free log subscription operation binding the contract event 0xd9b1841df91b154dfe2cc2d10b346bb5298e38fc4e76a76e9d6ad05bceebbbbd.
//
// Solidity: event VaultChainRemoved(uint256 id, uint256 chainId, address contractAddress)
func (_Tss *TssFilterer) WatchVaultChainRemoved(opts *bind.WatchOpts, sink chan<- *TssVaultChainRemoved) (event.Subscription, error) {

	logs, sub, err := _Tss.contract.WatchLogs(opts, "VaultChainRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TssVaultChainRemoved)
				if err := _Tss.contract.UnpackLog(event, "VaultChainRemoved", log); err != nil {
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

// ParseVaultChainRemoved is a log parse operation binding the contract event 0xd9b1841df91b154dfe2cc2d10b346bb5298e38fc4e76a76e9d6ad05bceebbbbd.
//
// Solidity: event VaultChainRemoved(uint256 id, uint256 chainId, address contractAddress)
func (_Tss *TssFilterer) ParseVaultChainRemoved(log types.Log) (*TssVaultChainRemoved, error) {
	event := new(TssVaultChainRemoved)
	if err := _Tss.contract.UnpackLog(event, "VaultChainRemoved", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TssVaultCreatedIterator is returned from FilterVaultCreated and is used to iterate over the raw logs and unpacked data for VaultCreated events raised by the Tss contract.
type TssVaultCreatedIterator struct {
	Event *TssVaultCreated // Event containing the contract specifics and raw log

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
func (it *TssVaultCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TssVaultCreated)
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
		it.Event = new(TssVaultCreated)
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
func (it *TssVaultCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TssVaultCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TssVaultCreated represents a VaultCreated event raised by the Tss contract.
type TssVaultCreated struct {
	Id                *big.Int
	PubKey            []byte
	Members           [][]byte
	ChainIds          []*big.Int
	ContractAddresses []common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterVaultCreated is a free log retrieval operation binding the contract event 0x0faf49e2d45fde9c990679020d75e87eb6c3e5f7149dec58f1cb09afbb496356.
//
// Solidity: event VaultCreated(uint256 id, bytes pubKey, bytes[] members, uint256[] chainIds, address[] contractAddresses)
func (_Tss *TssFilterer) FilterVaultCreated(opts *bind.FilterOpts) (*TssVaultCreatedIterator, error) {

	logs, sub, err := _Tss.contract.FilterLogs(opts, "VaultCreated")
	if err != nil {
		return nil, err
	}
	return &TssVaultCreatedIterator{contract: _Tss.contract, event: "VaultCreated", logs: logs, sub: sub}, nil
}

// WatchVaultCreated is a free log subscription operation binding the contract event 0x0faf49e2d45fde9c990679020d75e87eb6c3e5f7149dec58f1cb09afbb496356.
//
// Solidity: event VaultCreated(uint256 id, bytes pubKey, bytes[] members, uint256[] chainIds, address[] contractAddresses)
func (_Tss *TssFilterer) WatchVaultCreated(opts *bind.WatchOpts, sink chan<- *TssVaultCreated) (event.Subscription, error) {

	logs, sub, err := _Tss.contract.WatchLogs(opts, "VaultCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TssVaultCreated)
				if err := _Tss.contract.UnpackLog(event, "VaultCreated", log); err != nil {
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

// ParseVaultCreated is a log parse operation binding the contract event 0x0faf49e2d45fde9c990679020d75e87eb6c3e5f7149dec58f1cb09afbb496356.
//
// Solidity: event VaultCreated(uint256 id, bytes pubKey, bytes[] members, uint256[] chainIds, address[] contractAddresses)
func (_Tss *TssFilterer) ParseVaultCreated(log types.Log) (*TssVaultCreated, error) {
	event := new(TssVaultCreated)
	if err := _Tss.contract.UnpackLog(event, "VaultCreated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TssVaultUpdatedIterator is returned from FilterVaultUpdated and is used to iterate over the raw logs and unpacked data for VaultUpdated events raised by the Tss contract.
type TssVaultUpdatedIterator struct {
	Event *TssVaultUpdated // Event containing the contract specifics and raw log

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
func (it *TssVaultUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TssVaultUpdated)
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
		it.Event = new(TssVaultUpdated)
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
func (it *TssVaultUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TssVaultUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TssVaultUpdated represents a VaultUpdated event raised by the Tss contract.
type TssVaultUpdated struct {
	Id      *big.Int
	PubKey  []byte
	Members [][]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterVaultUpdated is a free log retrieval operation binding the contract event 0x46bb55f72ccd86f76e250fdfd4e11bf90791a760e05d643f9cadc6d2cd81d187.
//
// Solidity: event VaultUpdated(uint256 id, bytes pubKey, bytes[] members)
func (_Tss *TssFilterer) FilterVaultUpdated(opts *bind.FilterOpts) (*TssVaultUpdatedIterator, error) {

	logs, sub, err := _Tss.contract.FilterLogs(opts, "VaultUpdated")
	if err != nil {
		return nil, err
	}
	return &TssVaultUpdatedIterator{contract: _Tss.contract, event: "VaultUpdated", logs: logs, sub: sub}, nil
}

// WatchVaultUpdated is a free log subscription operation binding the contract event 0x46bb55f72ccd86f76e250fdfd4e11bf90791a760e05d643f9cadc6d2cd81d187.
//
// Solidity: event VaultUpdated(uint256 id, bytes pubKey, bytes[] members)
func (_Tss *TssFilterer) WatchVaultUpdated(opts *bind.WatchOpts, sink chan<- *TssVaultUpdated) (event.Subscription, error) {

	logs, sub, err := _Tss.contract.WatchLogs(opts, "VaultUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TssVaultUpdated)
				if err := _Tss.contract.UnpackLog(event, "VaultUpdated", log); err != nil {
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

// ParseVaultUpdated is a log parse operation binding the contract event 0x46bb55f72ccd86f76e250fdfd4e11bf90791a760e05d643f9cadc6d2cd81d187.
//
// Solidity: event VaultUpdated(uint256 id, bytes pubKey, bytes[] members)
func (_Tss *TssFilterer) ParseVaultUpdated(log types.Log) (*TssVaultUpdated, error) {
	event := new(TssVaultUpdated)
	if err := _Tss.contract.UnpackLog(event, "VaultUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}

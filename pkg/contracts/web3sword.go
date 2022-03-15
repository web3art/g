// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// Web3SwordMetaData contains all meta data concerning the Web3Sword contract.
var Web3SwordMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"BuySuccess\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPrice\",\"type\":\"uint256\"}],\"name\":\"ResetPrice\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SelledCountUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"claimer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"t\",\"type\":\"uint8\"}],\"name\":\"SocialClaimSuccess\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"TransferBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TransferSingle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"URI\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Withdrawal\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"balanceOfBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"blockById\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"ownerAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"imgURL\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"number\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"currentPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeBatchTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"selledCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"swordMatrix\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\",\"payable\":true},{\"inputs\":[{\"internalType\":\"contractIMetadata\",\"name\":\"_metadataGenertaor\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"buy\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\",\"payable\":true},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"resetCurrentPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"t\",\"type\":\"uint8\"}],\"name\":\"socialClaim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"t\",\"type\":\"uint8\"}],\"name\":\"socialCanClaimTokenIds\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_uri\",\"type\":\"string\"}],\"name\":\"upload\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"uri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true}]",
}

// Web3SwordABI is the input ABI used to generate the binding from.
// Deprecated: Use Web3SwordMetaData.ABI instead.
var Web3SwordABI = Web3SwordMetaData.ABI

// Web3Sword is an auto generated Go binding around an Ethereum contract.
type Web3Sword struct {
	Web3SwordCaller     // Read-only binding to the contract
	Web3SwordTransactor // Write-only binding to the contract
	Web3SwordFilterer   // Log filterer for contract events
}

// Web3SwordCaller is an auto generated read-only Go binding around an Ethereum contract.
type Web3SwordCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Web3SwordTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Web3SwordTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Web3SwordFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Web3SwordFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Web3SwordSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Web3SwordSession struct {
	Contract     *Web3Sword        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Web3SwordCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Web3SwordCallerSession struct {
	Contract *Web3SwordCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// Web3SwordTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Web3SwordTransactorSession struct {
	Contract     *Web3SwordTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// Web3SwordRaw is an auto generated low-level Go binding around an Ethereum contract.
type Web3SwordRaw struct {
	Contract *Web3Sword // Generic contract binding to access the raw methods on
}

// Web3SwordCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Web3SwordCallerRaw struct {
	Contract *Web3SwordCaller // Generic read-only contract binding to access the raw methods on
}

// Web3SwordTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Web3SwordTransactorRaw struct {
	Contract *Web3SwordTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWeb3Sword creates a new instance of Web3Sword, bound to a specific deployed contract.
func NewWeb3Sword(address common.Address, backend bind.ContractBackend) (*Web3Sword, error) {
	contract, err := bindWeb3Sword(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Web3Sword{Web3SwordCaller: Web3SwordCaller{contract: contract}, Web3SwordTransactor: Web3SwordTransactor{contract: contract}, Web3SwordFilterer: Web3SwordFilterer{contract: contract}}, nil
}

// NewWeb3SwordCaller creates a new read-only instance of Web3Sword, bound to a specific deployed contract.
func NewWeb3SwordCaller(address common.Address, caller bind.ContractCaller) (*Web3SwordCaller, error) {
	contract, err := bindWeb3Sword(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Web3SwordCaller{contract: contract}, nil
}

// NewWeb3SwordTransactor creates a new write-only instance of Web3Sword, bound to a specific deployed contract.
func NewWeb3SwordTransactor(address common.Address, transactor bind.ContractTransactor) (*Web3SwordTransactor, error) {
	contract, err := bindWeb3Sword(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Web3SwordTransactor{contract: contract}, nil
}

// NewWeb3SwordFilterer creates a new log filterer instance of Web3Sword, bound to a specific deployed contract.
func NewWeb3SwordFilterer(address common.Address, filterer bind.ContractFilterer) (*Web3SwordFilterer, error) {
	contract, err := bindWeb3Sword(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Web3SwordFilterer{contract: contract}, nil
}

// bindWeb3Sword binds a generic wrapper to an already deployed contract.
func bindWeb3Sword(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Web3SwordABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Web3Sword *Web3SwordRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Web3Sword.Contract.Web3SwordCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Web3Sword *Web3SwordRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Web3Sword.Contract.Web3SwordTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Web3Sword *Web3SwordRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Web3Sword.Contract.Web3SwordTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Web3Sword *Web3SwordCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Web3Sword.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Web3Sword *Web3SwordTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Web3Sword.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Web3Sword *Web3SwordTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Web3Sword.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_Web3Sword *Web3SwordCaller) BalanceOf(opts *bind.CallOpts, account common.Address, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Web3Sword.contract.Call(opts, &out, "balanceOf", account, id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_Web3Sword *Web3SwordSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _Web3Sword.Contract.BalanceOf(&_Web3Sword.CallOpts, account, id)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_Web3Sword *Web3SwordCallerSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _Web3Sword.Contract.BalanceOf(&_Web3Sword.CallOpts, account, id)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_Web3Sword *Web3SwordCaller) BalanceOfBatch(opts *bind.CallOpts, accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _Web3Sword.contract.Call(opts, &out, "balanceOfBatch", accounts, ids)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_Web3Sword *Web3SwordSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _Web3Sword.Contract.BalanceOfBatch(&_Web3Sword.CallOpts, accounts, ids)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_Web3Sword *Web3SwordCallerSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _Web3Sword.Contract.BalanceOfBatch(&_Web3Sword.CallOpts, accounts, ids)
}

// BlockById is a free data retrieval call binding the contract method 0x62ba2e56.
//
// Solidity: function blockById(uint256 ) view returns(uint256 tokenId, address ownerAddress, string imgURL, uint256 number)
func (_Web3Sword *Web3SwordCaller) BlockById(opts *bind.CallOpts, arg0 *big.Int) (struct {
	TokenId      *big.Int
	OwnerAddress common.Address
	ImgURL       string
	Number       *big.Int
}, error) {
	var out []interface{}
	err := _Web3Sword.contract.Call(opts, &out, "blockById", arg0)

	outstruct := new(struct {
		TokenId      *big.Int
		OwnerAddress common.Address
		ImgURL       string
		Number       *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TokenId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.OwnerAddress = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.ImgURL = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Number = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// BlockById is a free data retrieval call binding the contract method 0x62ba2e56.
//
// Solidity: function blockById(uint256 ) view returns(uint256 tokenId, address ownerAddress, string imgURL, uint256 number)
func (_Web3Sword *Web3SwordSession) BlockById(arg0 *big.Int) (struct {
	TokenId      *big.Int
	OwnerAddress common.Address
	ImgURL       string
	Number       *big.Int
}, error) {
	return _Web3Sword.Contract.BlockById(&_Web3Sword.CallOpts, arg0)
}

// BlockById is a free data retrieval call binding the contract method 0x62ba2e56.
//
// Solidity: function blockById(uint256 ) view returns(uint256 tokenId, address ownerAddress, string imgURL, uint256 number)
func (_Web3Sword *Web3SwordCallerSession) BlockById(arg0 *big.Int) (struct {
	TokenId      *big.Int
	OwnerAddress common.Address
	ImgURL       string
	Number       *big.Int
}, error) {
	return _Web3Sword.Contract.BlockById(&_Web3Sword.CallOpts, arg0)
}

// CurrentPrice is a free data retrieval call binding the contract method 0x9d1b464a.
//
// Solidity: function currentPrice() view returns(uint256)
func (_Web3Sword *Web3SwordCaller) CurrentPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Web3Sword.contract.Call(opts, &out, "currentPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentPrice is a free data retrieval call binding the contract method 0x9d1b464a.
//
// Solidity: function currentPrice() view returns(uint256)
func (_Web3Sword *Web3SwordSession) CurrentPrice() (*big.Int, error) {
	return _Web3Sword.Contract.CurrentPrice(&_Web3Sword.CallOpts)
}

// CurrentPrice is a free data retrieval call binding the contract method 0x9d1b464a.
//
// Solidity: function currentPrice() view returns(uint256)
func (_Web3Sword *Web3SwordCallerSession) CurrentPrice() (*big.Int, error) {
	return _Web3Sword.Contract.CurrentPrice(&_Web3Sword.CallOpts)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_Web3Sword *Web3SwordCaller) IsApprovedForAll(opts *bind.CallOpts, account common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Web3Sword.contract.Call(opts, &out, "isApprovedForAll", account, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_Web3Sword *Web3SwordSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _Web3Sword.Contract.IsApprovedForAll(&_Web3Sword.CallOpts, account, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_Web3Sword *Web3SwordCallerSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _Web3Sword.Contract.IsApprovedForAll(&_Web3Sword.CallOpts, account, operator)
}

// Name is a free data retrieval call binding the contract method 0x00ad800c.
//
// Solidity: function name(uint256 ) view returns(string)
func (_Web3Sword *Web3SwordCaller) Name(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _Web3Sword.contract.Call(opts, &out, "name", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x00ad800c.
//
// Solidity: function name(uint256 ) view returns(string)
func (_Web3Sword *Web3SwordSession) Name(arg0 *big.Int) (string, error) {
	return _Web3Sword.Contract.Name(&_Web3Sword.CallOpts, arg0)
}

// Name is a free data retrieval call binding the contract method 0x00ad800c.
//
// Solidity: function name(uint256 ) view returns(string)
func (_Web3Sword *Web3SwordCallerSession) Name(arg0 *big.Int) (string, error) {
	return _Web3Sword.Contract.Name(&_Web3Sword.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Web3Sword *Web3SwordCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Web3Sword.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Web3Sword *Web3SwordSession) Owner() (common.Address, error) {
	return _Web3Sword.Contract.Owner(&_Web3Sword.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Web3Sword *Web3SwordCallerSession) Owner() (common.Address, error) {
	return _Web3Sword.Contract.Owner(&_Web3Sword.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Web3Sword *Web3SwordCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Web3Sword.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Web3Sword *Web3SwordSession) ProxiableUUID() ([32]byte, error) {
	return _Web3Sword.Contract.ProxiableUUID(&_Web3Sword.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Web3Sword *Web3SwordCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Web3Sword.Contract.ProxiableUUID(&_Web3Sword.CallOpts)
}

// SelledCount is a free data retrieval call binding the contract method 0xf567f405.
//
// Solidity: function selledCount() view returns(uint256)
func (_Web3Sword *Web3SwordCaller) SelledCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Web3Sword.contract.Call(opts, &out, "selledCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SelledCount is a free data retrieval call binding the contract method 0xf567f405.
//
// Solidity: function selledCount() view returns(uint256)
func (_Web3Sword *Web3SwordSession) SelledCount() (*big.Int, error) {
	return _Web3Sword.Contract.SelledCount(&_Web3Sword.CallOpts)
}

// SelledCount is a free data retrieval call binding the contract method 0xf567f405.
//
// Solidity: function selledCount() view returns(uint256)
func (_Web3Sword *Web3SwordCallerSession) SelledCount() (*big.Int, error) {
	return _Web3Sword.Contract.SelledCount(&_Web3Sword.CallOpts)
}

// SocialCanClaimTokenIds is a free data retrieval call binding the contract method 0x5f347c4f.
//
// Solidity: function socialCanClaimTokenIds(uint8 t) view returns(uint256[])
func (_Web3Sword *Web3SwordCaller) SocialCanClaimTokenIds(opts *bind.CallOpts, t uint8) ([]*big.Int, error) {
	var out []interface{}
	err := _Web3Sword.contract.Call(opts, &out, "socialCanClaimTokenIds", t)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// SocialCanClaimTokenIds is a free data retrieval call binding the contract method 0x5f347c4f.
//
// Solidity: function socialCanClaimTokenIds(uint8 t) view returns(uint256[])
func (_Web3Sword *Web3SwordSession) SocialCanClaimTokenIds(t uint8) ([]*big.Int, error) {
	return _Web3Sword.Contract.SocialCanClaimTokenIds(&_Web3Sword.CallOpts, t)
}

// SocialCanClaimTokenIds is a free data retrieval call binding the contract method 0x5f347c4f.
//
// Solidity: function socialCanClaimTokenIds(uint8 t) view returns(uint256[])
func (_Web3Sword *Web3SwordCallerSession) SocialCanClaimTokenIds(t uint8) ([]*big.Int, error) {
	return _Web3Sword.Contract.SocialCanClaimTokenIds(&_Web3Sword.CallOpts, t)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Web3Sword *Web3SwordCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Web3Sword.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Web3Sword *Web3SwordSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Web3Sword.Contract.SupportsInterface(&_Web3Sword.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Web3Sword *Web3SwordCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Web3Sword.Contract.SupportsInterface(&_Web3Sword.CallOpts, interfaceId)
}

// SwordMatrix is a free data retrieval call binding the contract method 0xe8a083b1.
//
// Solidity: function swordMatrix(uint256 , uint256 ) view returns(uint8)
func (_Web3Sword *Web3SwordCaller) SwordMatrix(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (uint8, error) {
	var out []interface{}
	err := _Web3Sword.contract.Call(opts, &out, "swordMatrix", arg0, arg1)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// SwordMatrix is a free data retrieval call binding the contract method 0xe8a083b1.
//
// Solidity: function swordMatrix(uint256 , uint256 ) view returns(uint8)
func (_Web3Sword *Web3SwordSession) SwordMatrix(arg0 *big.Int, arg1 *big.Int) (uint8, error) {
	return _Web3Sword.Contract.SwordMatrix(&_Web3Sword.CallOpts, arg0, arg1)
}

// SwordMatrix is a free data retrieval call binding the contract method 0xe8a083b1.
//
// Solidity: function swordMatrix(uint256 , uint256 ) view returns(uint8)
func (_Web3Sword *Web3SwordCallerSession) SwordMatrix(arg0 *big.Int, arg1 *big.Int) (uint8, error) {
	return _Web3Sword.Contract.SwordMatrix(&_Web3Sword.CallOpts, arg0, arg1)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 id) view returns(string)
func (_Web3Sword *Web3SwordCaller) Uri(opts *bind.CallOpts, id *big.Int) (string, error) {
	var out []interface{}
	err := _Web3Sword.contract.Call(opts, &out, "uri", id)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 id) view returns(string)
func (_Web3Sword *Web3SwordSession) Uri(id *big.Int) (string, error) {
	return _Web3Sword.Contract.Uri(&_Web3Sword.CallOpts, id)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 id) view returns(string)
func (_Web3Sword *Web3SwordCallerSession) Uri(id *big.Int) (string, error) {
	return _Web3Sword.Contract.Uri(&_Web3Sword.CallOpts, id)
}

// Buy is a paid mutator transaction binding the contract method 0xd96a094a.
//
// Solidity: function buy(uint256 tokenId) payable returns()
func (_Web3Sword *Web3SwordTransactor) Buy(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _Web3Sword.contract.Transact(opts, "buy", tokenId)
}

// Buy is a paid mutator transaction binding the contract method 0xd96a094a.
//
// Solidity: function buy(uint256 tokenId) payable returns()
func (_Web3Sword *Web3SwordSession) Buy(tokenId *big.Int) (*types.Transaction, error) {
	return _Web3Sword.Contract.Buy(&_Web3Sword.TransactOpts, tokenId)
}

// Buy is a paid mutator transaction binding the contract method 0xd96a094a.
//
// Solidity: function buy(uint256 tokenId) payable returns()
func (_Web3Sword *Web3SwordTransactorSession) Buy(tokenId *big.Int) (*types.Transaction, error) {
	return _Web3Sword.Contract.Buy(&_Web3Sword.TransactOpts, tokenId)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _metadataGenertaor) returns()
func (_Web3Sword *Web3SwordTransactor) Initialize(opts *bind.TransactOpts, _metadataGenertaor common.Address) (*types.Transaction, error) {
	return _Web3Sword.contract.Transact(opts, "initialize", _metadataGenertaor)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _metadataGenertaor) returns()
func (_Web3Sword *Web3SwordSession) Initialize(_metadataGenertaor common.Address) (*types.Transaction, error) {
	return _Web3Sword.Contract.Initialize(&_Web3Sword.TransactOpts, _metadataGenertaor)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _metadataGenertaor) returns()
func (_Web3Sword *Web3SwordTransactorSession) Initialize(_metadataGenertaor common.Address) (*types.Transaction, error) {
	return _Web3Sword.Contract.Initialize(&_Web3Sword.TransactOpts, _metadataGenertaor)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Web3Sword *Web3SwordTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Web3Sword.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Web3Sword *Web3SwordSession) RenounceOwnership() (*types.Transaction, error) {
	return _Web3Sword.Contract.RenounceOwnership(&_Web3Sword.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Web3Sword *Web3SwordTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Web3Sword.Contract.RenounceOwnership(&_Web3Sword.TransactOpts)
}

// ResetCurrentPrice is a paid mutator transaction binding the contract method 0xbcf3e129.
//
// Solidity: function resetCurrentPrice(uint256 price) returns()
func (_Web3Sword *Web3SwordTransactor) ResetCurrentPrice(opts *bind.TransactOpts, price *big.Int) (*types.Transaction, error) {
	return _Web3Sword.contract.Transact(opts, "resetCurrentPrice", price)
}

// ResetCurrentPrice is a paid mutator transaction binding the contract method 0xbcf3e129.
//
// Solidity: function resetCurrentPrice(uint256 price) returns()
func (_Web3Sword *Web3SwordSession) ResetCurrentPrice(price *big.Int) (*types.Transaction, error) {
	return _Web3Sword.Contract.ResetCurrentPrice(&_Web3Sword.TransactOpts, price)
}

// ResetCurrentPrice is a paid mutator transaction binding the contract method 0xbcf3e129.
//
// Solidity: function resetCurrentPrice(uint256 price) returns()
func (_Web3Sword *Web3SwordTransactorSession) ResetCurrentPrice(price *big.Int) (*types.Transaction, error) {
	return _Web3Sword.Contract.ResetCurrentPrice(&_Web3Sword.TransactOpts, price)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_Web3Sword *Web3SwordTransactor) SafeBatchTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _Web3Sword.contract.Transact(opts, "safeBatchTransferFrom", from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_Web3Sword *Web3SwordSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _Web3Sword.Contract.SafeBatchTransferFrom(&_Web3Sword.TransactOpts, from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_Web3Sword *Web3SwordTransactorSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _Web3Sword.Contract.SafeBatchTransferFrom(&_Web3Sword.TransactOpts, from, to, ids, amounts, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_Web3Sword *Web3SwordTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Web3Sword.contract.Transact(opts, "safeTransferFrom", from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_Web3Sword *Web3SwordSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Web3Sword.Contract.SafeTransferFrom(&_Web3Sword.TransactOpts, from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_Web3Sword *Web3SwordTransactorSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Web3Sword.Contract.SafeTransferFrom(&_Web3Sword.TransactOpts, from, to, id, amount, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Web3Sword *Web3SwordTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Web3Sword.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Web3Sword *Web3SwordSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Web3Sword.Contract.SetApprovalForAll(&_Web3Sword.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Web3Sword *Web3SwordTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Web3Sword.Contract.SetApprovalForAll(&_Web3Sword.TransactOpts, operator, approved)
}

// SocialClaim is a paid mutator transaction binding the contract method 0xf7f10f69.
//
// Solidity: function socialClaim(address to, uint256 tokenId, uint8 t) returns()
func (_Web3Sword *Web3SwordTransactor) SocialClaim(opts *bind.TransactOpts, to common.Address, tokenId *big.Int, t uint8) (*types.Transaction, error) {
	return _Web3Sword.contract.Transact(opts, "socialClaim", to, tokenId, t)
}

// SocialClaim is a paid mutator transaction binding the contract method 0xf7f10f69.
//
// Solidity: function socialClaim(address to, uint256 tokenId, uint8 t) returns()
func (_Web3Sword *Web3SwordSession) SocialClaim(to common.Address, tokenId *big.Int, t uint8) (*types.Transaction, error) {
	return _Web3Sword.Contract.SocialClaim(&_Web3Sword.TransactOpts, to, tokenId, t)
}

// SocialClaim is a paid mutator transaction binding the contract method 0xf7f10f69.
//
// Solidity: function socialClaim(address to, uint256 tokenId, uint8 t) returns()
func (_Web3Sword *Web3SwordTransactorSession) SocialClaim(to common.Address, tokenId *big.Int, t uint8) (*types.Transaction, error) {
	return _Web3Sword.Contract.SocialClaim(&_Web3Sword.TransactOpts, to, tokenId, t)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Web3Sword *Web3SwordTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Web3Sword.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Web3Sword *Web3SwordSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Web3Sword.Contract.TransferOwnership(&_Web3Sword.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Web3Sword *Web3SwordTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Web3Sword.Contract.TransferOwnership(&_Web3Sword.TransactOpts, newOwner)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Web3Sword *Web3SwordTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _Web3Sword.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Web3Sword *Web3SwordSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Web3Sword.Contract.UpgradeTo(&_Web3Sword.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Web3Sword *Web3SwordTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Web3Sword.Contract.UpgradeTo(&_Web3Sword.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Web3Sword *Web3SwordTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Web3Sword.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Web3Sword *Web3SwordSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Web3Sword.Contract.UpgradeToAndCall(&_Web3Sword.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Web3Sword *Web3SwordTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Web3Sword.Contract.UpgradeToAndCall(&_Web3Sword.TransactOpts, newImplementation, data)
}

// Upload is a paid mutator transaction binding the contract method 0x4d725e95.
//
// Solidity: function upload(uint256 tokenId, string _uri) returns()
func (_Web3Sword *Web3SwordTransactor) Upload(opts *bind.TransactOpts, tokenId *big.Int, _uri string) (*types.Transaction, error) {
	return _Web3Sword.contract.Transact(opts, "upload", tokenId, _uri)
}

// Upload is a paid mutator transaction binding the contract method 0x4d725e95.
//
// Solidity: function upload(uint256 tokenId, string _uri) returns()
func (_Web3Sword *Web3SwordSession) Upload(tokenId *big.Int, _uri string) (*types.Transaction, error) {
	return _Web3Sword.Contract.Upload(&_Web3Sword.TransactOpts, tokenId, _uri)
}

// Upload is a paid mutator transaction binding the contract method 0x4d725e95.
//
// Solidity: function upload(uint256 tokenId, string _uri) returns()
func (_Web3Sword *Web3SwordTransactorSession) Upload(tokenId *big.Int, _uri string) (*types.Transaction, error) {
	return _Web3Sword.Contract.Upload(&_Web3Sword.TransactOpts, tokenId, _uri)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 value) returns()
func (_Web3Sword *Web3SwordTransactor) Withdraw(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _Web3Sword.contract.Transact(opts, "withdraw", value)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 value) returns()
func (_Web3Sword *Web3SwordSession) Withdraw(value *big.Int) (*types.Transaction, error) {
	return _Web3Sword.Contract.Withdraw(&_Web3Sword.TransactOpts, value)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 value) returns()
func (_Web3Sword *Web3SwordTransactorSession) Withdraw(value *big.Int) (*types.Transaction, error) {
	return _Web3Sword.Contract.Withdraw(&_Web3Sword.TransactOpts, value)
}

// Web3SwordAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the Web3Sword contract.
type Web3SwordAdminChangedIterator struct {
	Event *Web3SwordAdminChanged // Event containing the contract specifics and raw log

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
func (it *Web3SwordAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Web3SwordAdminChanged)
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
		it.Event = new(Web3SwordAdminChanged)
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
func (it *Web3SwordAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Web3SwordAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Web3SwordAdminChanged represents a AdminChanged event raised by the Web3Sword contract.
type Web3SwordAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Web3Sword *Web3SwordFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*Web3SwordAdminChangedIterator, error) {

	logs, sub, err := _Web3Sword.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &Web3SwordAdminChangedIterator{contract: _Web3Sword.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Web3Sword *Web3SwordFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *Web3SwordAdminChanged) (event.Subscription, error) {

	logs, sub, err := _Web3Sword.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Web3SwordAdminChanged)
				if err := _Web3Sword.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Web3Sword *Web3SwordFilterer) ParseAdminChanged(log types.Log) (*Web3SwordAdminChanged, error) {
	event := new(Web3SwordAdminChanged)
	if err := _Web3Sword.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Web3SwordApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Web3Sword contract.
type Web3SwordApprovalForAllIterator struct {
	Event *Web3SwordApprovalForAll // Event containing the contract specifics and raw log

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
func (it *Web3SwordApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Web3SwordApprovalForAll)
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
		it.Event = new(Web3SwordApprovalForAll)
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
func (it *Web3SwordApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Web3SwordApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Web3SwordApprovalForAll represents a ApprovalForAll event raised by the Web3Sword contract.
type Web3SwordApprovalForAll struct {
	Account  common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_Web3Sword *Web3SwordFilterer) FilterApprovalForAll(opts *bind.FilterOpts, account []common.Address, operator []common.Address) (*Web3SwordApprovalForAllIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Web3Sword.contract.FilterLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &Web3SwordApprovalForAllIterator{contract: _Web3Sword.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_Web3Sword *Web3SwordFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *Web3SwordApprovalForAll, account []common.Address, operator []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Web3Sword.contract.WatchLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Web3SwordApprovalForAll)
				if err := _Web3Sword.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_Web3Sword *Web3SwordFilterer) ParseApprovalForAll(log types.Log) (*Web3SwordApprovalForAll, error) {
	event := new(Web3SwordApprovalForAll)
	if err := _Web3Sword.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Web3SwordBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the Web3Sword contract.
type Web3SwordBeaconUpgradedIterator struct {
	Event *Web3SwordBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *Web3SwordBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Web3SwordBeaconUpgraded)
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
		it.Event = new(Web3SwordBeaconUpgraded)
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
func (it *Web3SwordBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Web3SwordBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Web3SwordBeaconUpgraded represents a BeaconUpgraded event raised by the Web3Sword contract.
type Web3SwordBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Web3Sword *Web3SwordFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*Web3SwordBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Web3Sword.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &Web3SwordBeaconUpgradedIterator{contract: _Web3Sword.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Web3Sword *Web3SwordFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *Web3SwordBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Web3Sword.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Web3SwordBeaconUpgraded)
				if err := _Web3Sword.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Web3Sword *Web3SwordFilterer) ParseBeaconUpgraded(log types.Log) (*Web3SwordBeaconUpgraded, error) {
	event := new(Web3SwordBeaconUpgraded)
	if err := _Web3Sword.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Web3SwordBuySuccessIterator is returned from FilterBuySuccess and is used to iterate over the raw logs and unpacked data for BuySuccess events raised by the Web3Sword contract.
type Web3SwordBuySuccessIterator struct {
	Event *Web3SwordBuySuccess // Event containing the contract specifics and raw log

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
func (it *Web3SwordBuySuccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Web3SwordBuySuccess)
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
		it.Event = new(Web3SwordBuySuccess)
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
func (it *Web3SwordBuySuccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Web3SwordBuySuccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Web3SwordBuySuccess represents a BuySuccess event raised by the Web3Sword contract.
type Web3SwordBuySuccess struct {
	Buyer   common.Address
	TokenId *big.Int
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBuySuccess is a free log retrieval operation binding the contract event 0xf86cbedb5a8608ff23c1d58ffb721f8b4e1a8f793f510208d89d346b3b24ef28.
//
// Solidity: event BuySuccess(address indexed buyer, uint256 tokenId, uint256 value)
func (_Web3Sword *Web3SwordFilterer) FilterBuySuccess(opts *bind.FilterOpts, buyer []common.Address) (*Web3SwordBuySuccessIterator, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _Web3Sword.contract.FilterLogs(opts, "BuySuccess", buyerRule)
	if err != nil {
		return nil, err
	}
	return &Web3SwordBuySuccessIterator{contract: _Web3Sword.contract, event: "BuySuccess", logs: logs, sub: sub}, nil
}

// WatchBuySuccess is a free log subscription operation binding the contract event 0xf86cbedb5a8608ff23c1d58ffb721f8b4e1a8f793f510208d89d346b3b24ef28.
//
// Solidity: event BuySuccess(address indexed buyer, uint256 tokenId, uint256 value)
func (_Web3Sword *Web3SwordFilterer) WatchBuySuccess(opts *bind.WatchOpts, sink chan<- *Web3SwordBuySuccess, buyer []common.Address) (event.Subscription, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _Web3Sword.contract.WatchLogs(opts, "BuySuccess", buyerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Web3SwordBuySuccess)
				if err := _Web3Sword.contract.UnpackLog(event, "BuySuccess", log); err != nil {
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

// ParseBuySuccess is a log parse operation binding the contract event 0xf86cbedb5a8608ff23c1d58ffb721f8b4e1a8f793f510208d89d346b3b24ef28.
//
// Solidity: event BuySuccess(address indexed buyer, uint256 tokenId, uint256 value)
func (_Web3Sword *Web3SwordFilterer) ParseBuySuccess(log types.Log) (*Web3SwordBuySuccess, error) {
	event := new(Web3SwordBuySuccess)
	if err := _Web3Sword.contract.UnpackLog(event, "BuySuccess", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Web3SwordOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Web3Sword contract.
type Web3SwordOwnershipTransferredIterator struct {
	Event *Web3SwordOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *Web3SwordOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Web3SwordOwnershipTransferred)
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
		it.Event = new(Web3SwordOwnershipTransferred)
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
func (it *Web3SwordOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Web3SwordOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Web3SwordOwnershipTransferred represents a OwnershipTransferred event raised by the Web3Sword contract.
type Web3SwordOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Web3Sword *Web3SwordFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*Web3SwordOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Web3Sword.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &Web3SwordOwnershipTransferredIterator{contract: _Web3Sword.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Web3Sword *Web3SwordFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *Web3SwordOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Web3Sword.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Web3SwordOwnershipTransferred)
				if err := _Web3Sword.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Web3Sword *Web3SwordFilterer) ParseOwnershipTransferred(log types.Log) (*Web3SwordOwnershipTransferred, error) {
	event := new(Web3SwordOwnershipTransferred)
	if err := _Web3Sword.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Web3SwordResetPriceIterator is returned from FilterResetPrice and is used to iterate over the raw logs and unpacked data for ResetPrice events raised by the Web3Sword contract.
type Web3SwordResetPriceIterator struct {
	Event *Web3SwordResetPrice // Event containing the contract specifics and raw log

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
func (it *Web3SwordResetPriceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Web3SwordResetPrice)
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
		it.Event = new(Web3SwordResetPrice)
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
func (it *Web3SwordResetPriceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Web3SwordResetPriceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Web3SwordResetPrice represents a ResetPrice event raised by the Web3Sword contract.
type Web3SwordResetPrice struct {
	NewPrice *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterResetPrice is a free log retrieval operation binding the contract event 0xac8f29f4224d50dcf2e66da564835b17a8c5ae91225eeb020ab82995cddc83b4.
//
// Solidity: event ResetPrice(uint256 newPrice)
func (_Web3Sword *Web3SwordFilterer) FilterResetPrice(opts *bind.FilterOpts) (*Web3SwordResetPriceIterator, error) {

	logs, sub, err := _Web3Sword.contract.FilterLogs(opts, "ResetPrice")
	if err != nil {
		return nil, err
	}
	return &Web3SwordResetPriceIterator{contract: _Web3Sword.contract, event: "ResetPrice", logs: logs, sub: sub}, nil
}

// WatchResetPrice is a free log subscription operation binding the contract event 0xac8f29f4224d50dcf2e66da564835b17a8c5ae91225eeb020ab82995cddc83b4.
//
// Solidity: event ResetPrice(uint256 newPrice)
func (_Web3Sword *Web3SwordFilterer) WatchResetPrice(opts *bind.WatchOpts, sink chan<- *Web3SwordResetPrice) (event.Subscription, error) {

	logs, sub, err := _Web3Sword.contract.WatchLogs(opts, "ResetPrice")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Web3SwordResetPrice)
				if err := _Web3Sword.contract.UnpackLog(event, "ResetPrice", log); err != nil {
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

// ParseResetPrice is a log parse operation binding the contract event 0xac8f29f4224d50dcf2e66da564835b17a8c5ae91225eeb020ab82995cddc83b4.
//
// Solidity: event ResetPrice(uint256 newPrice)
func (_Web3Sword *Web3SwordFilterer) ParseResetPrice(log types.Log) (*Web3SwordResetPrice, error) {
	event := new(Web3SwordResetPrice)
	if err := _Web3Sword.contract.UnpackLog(event, "ResetPrice", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Web3SwordSelledCountUpdateIterator is returned from FilterSelledCountUpdate and is used to iterate over the raw logs and unpacked data for SelledCountUpdate events raised by the Web3Sword contract.
type Web3SwordSelledCountUpdateIterator struct {
	Event *Web3SwordSelledCountUpdate // Event containing the contract specifics and raw log

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
func (it *Web3SwordSelledCountUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Web3SwordSelledCountUpdate)
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
		it.Event = new(Web3SwordSelledCountUpdate)
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
func (it *Web3SwordSelledCountUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Web3SwordSelledCountUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Web3SwordSelledCountUpdate represents a SelledCountUpdate event raised by the Web3Sword contract.
type Web3SwordSelledCountUpdate struct {
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSelledCountUpdate is a free log retrieval operation binding the contract event 0xdbc38d09e60916f3bf282de216054bdaaaa8c1da1c4a5e3c9661163cbe126f62.
//
// Solidity: event SelledCountUpdate(uint256 value)
func (_Web3Sword *Web3SwordFilterer) FilterSelledCountUpdate(opts *bind.FilterOpts) (*Web3SwordSelledCountUpdateIterator, error) {

	logs, sub, err := _Web3Sword.contract.FilterLogs(opts, "SelledCountUpdate")
	if err != nil {
		return nil, err
	}
	return &Web3SwordSelledCountUpdateIterator{contract: _Web3Sword.contract, event: "SelledCountUpdate", logs: logs, sub: sub}, nil
}

// WatchSelledCountUpdate is a free log subscription operation binding the contract event 0xdbc38d09e60916f3bf282de216054bdaaaa8c1da1c4a5e3c9661163cbe126f62.
//
// Solidity: event SelledCountUpdate(uint256 value)
func (_Web3Sword *Web3SwordFilterer) WatchSelledCountUpdate(opts *bind.WatchOpts, sink chan<- *Web3SwordSelledCountUpdate) (event.Subscription, error) {

	logs, sub, err := _Web3Sword.contract.WatchLogs(opts, "SelledCountUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Web3SwordSelledCountUpdate)
				if err := _Web3Sword.contract.UnpackLog(event, "SelledCountUpdate", log); err != nil {
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

// ParseSelledCountUpdate is a log parse operation binding the contract event 0xdbc38d09e60916f3bf282de216054bdaaaa8c1da1c4a5e3c9661163cbe126f62.
//
// Solidity: event SelledCountUpdate(uint256 value)
func (_Web3Sword *Web3SwordFilterer) ParseSelledCountUpdate(log types.Log) (*Web3SwordSelledCountUpdate, error) {
	event := new(Web3SwordSelledCountUpdate)
	if err := _Web3Sword.contract.UnpackLog(event, "SelledCountUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Web3SwordSocialClaimSuccessIterator is returned from FilterSocialClaimSuccess and is used to iterate over the raw logs and unpacked data for SocialClaimSuccess events raised by the Web3Sword contract.
type Web3SwordSocialClaimSuccessIterator struct {
	Event *Web3SwordSocialClaimSuccess // Event containing the contract specifics and raw log

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
func (it *Web3SwordSocialClaimSuccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Web3SwordSocialClaimSuccess)
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
		it.Event = new(Web3SwordSocialClaimSuccess)
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
func (it *Web3SwordSocialClaimSuccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Web3SwordSocialClaimSuccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Web3SwordSocialClaimSuccess represents a SocialClaimSuccess event raised by the Web3Sword contract.
type Web3SwordSocialClaimSuccess struct {
	Claimer common.Address
	TokenId *big.Int
	T       uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSocialClaimSuccess is a free log retrieval operation binding the contract event 0x11dab502fce81d031bd91a4e2af283ca6bfdfb67461321873e517285c5dbecf0.
//
// Solidity: event SocialClaimSuccess(address indexed claimer, uint256 tokenId, uint8 t)
func (_Web3Sword *Web3SwordFilterer) FilterSocialClaimSuccess(opts *bind.FilterOpts, claimer []common.Address) (*Web3SwordSocialClaimSuccessIterator, error) {

	var claimerRule []interface{}
	for _, claimerItem := range claimer {
		claimerRule = append(claimerRule, claimerItem)
	}

	logs, sub, err := _Web3Sword.contract.FilterLogs(opts, "SocialClaimSuccess", claimerRule)
	if err != nil {
		return nil, err
	}
	return &Web3SwordSocialClaimSuccessIterator{contract: _Web3Sword.contract, event: "SocialClaimSuccess", logs: logs, sub: sub}, nil
}

// WatchSocialClaimSuccess is a free log subscription operation binding the contract event 0x11dab502fce81d031bd91a4e2af283ca6bfdfb67461321873e517285c5dbecf0.
//
// Solidity: event SocialClaimSuccess(address indexed claimer, uint256 tokenId, uint8 t)
func (_Web3Sword *Web3SwordFilterer) WatchSocialClaimSuccess(opts *bind.WatchOpts, sink chan<- *Web3SwordSocialClaimSuccess, claimer []common.Address) (event.Subscription, error) {

	var claimerRule []interface{}
	for _, claimerItem := range claimer {
		claimerRule = append(claimerRule, claimerItem)
	}

	logs, sub, err := _Web3Sword.contract.WatchLogs(opts, "SocialClaimSuccess", claimerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Web3SwordSocialClaimSuccess)
				if err := _Web3Sword.contract.UnpackLog(event, "SocialClaimSuccess", log); err != nil {
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

// ParseSocialClaimSuccess is a log parse operation binding the contract event 0x11dab502fce81d031bd91a4e2af283ca6bfdfb67461321873e517285c5dbecf0.
//
// Solidity: event SocialClaimSuccess(address indexed claimer, uint256 tokenId, uint8 t)
func (_Web3Sword *Web3SwordFilterer) ParseSocialClaimSuccess(log types.Log) (*Web3SwordSocialClaimSuccess, error) {
	event := new(Web3SwordSocialClaimSuccess)
	if err := _Web3Sword.contract.UnpackLog(event, "SocialClaimSuccess", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Web3SwordTransferBatchIterator is returned from FilterTransferBatch and is used to iterate over the raw logs and unpacked data for TransferBatch events raised by the Web3Sword contract.
type Web3SwordTransferBatchIterator struct {
	Event *Web3SwordTransferBatch // Event containing the contract specifics and raw log

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
func (it *Web3SwordTransferBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Web3SwordTransferBatch)
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
		it.Event = new(Web3SwordTransferBatch)
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
func (it *Web3SwordTransferBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Web3SwordTransferBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Web3SwordTransferBatch represents a TransferBatch event raised by the Web3Sword contract.
type Web3SwordTransferBatch struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Ids      []*big.Int
	Values   []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferBatch is a free log retrieval operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_Web3Sword *Web3SwordFilterer) FilterTransferBatch(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*Web3SwordTransferBatchIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Web3Sword.contract.FilterLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &Web3SwordTransferBatchIterator{contract: _Web3Sword.contract, event: "TransferBatch", logs: logs, sub: sub}, nil
}

// WatchTransferBatch is a free log subscription operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_Web3Sword *Web3SwordFilterer) WatchTransferBatch(opts *bind.WatchOpts, sink chan<- *Web3SwordTransferBatch, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Web3Sword.contract.WatchLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Web3SwordTransferBatch)
				if err := _Web3Sword.contract.UnpackLog(event, "TransferBatch", log); err != nil {
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

// ParseTransferBatch is a log parse operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_Web3Sword *Web3SwordFilterer) ParseTransferBatch(log types.Log) (*Web3SwordTransferBatch, error) {
	event := new(Web3SwordTransferBatch)
	if err := _Web3Sword.contract.UnpackLog(event, "TransferBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Web3SwordTransferSingleIterator is returned from FilterTransferSingle and is used to iterate over the raw logs and unpacked data for TransferSingle events raised by the Web3Sword contract.
type Web3SwordTransferSingleIterator struct {
	Event *Web3SwordTransferSingle // Event containing the contract specifics and raw log

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
func (it *Web3SwordTransferSingleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Web3SwordTransferSingle)
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
		it.Event = new(Web3SwordTransferSingle)
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
func (it *Web3SwordTransferSingleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Web3SwordTransferSingleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Web3SwordTransferSingle represents a TransferSingle event raised by the Web3Sword contract.
type Web3SwordTransferSingle struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Id       *big.Int
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferSingle is a free log retrieval operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_Web3Sword *Web3SwordFilterer) FilterTransferSingle(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*Web3SwordTransferSingleIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Web3Sword.contract.FilterLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &Web3SwordTransferSingleIterator{contract: _Web3Sword.contract, event: "TransferSingle", logs: logs, sub: sub}, nil
}

// WatchTransferSingle is a free log subscription operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_Web3Sword *Web3SwordFilterer) WatchTransferSingle(opts *bind.WatchOpts, sink chan<- *Web3SwordTransferSingle, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Web3Sword.contract.WatchLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Web3SwordTransferSingle)
				if err := _Web3Sword.contract.UnpackLog(event, "TransferSingle", log); err != nil {
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

// ParseTransferSingle is a log parse operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_Web3Sword *Web3SwordFilterer) ParseTransferSingle(log types.Log) (*Web3SwordTransferSingle, error) {
	event := new(Web3SwordTransferSingle)
	if err := _Web3Sword.contract.UnpackLog(event, "TransferSingle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Web3SwordURIIterator is returned from FilterURI and is used to iterate over the raw logs and unpacked data for URI events raised by the Web3Sword contract.
type Web3SwordURIIterator struct {
	Event *Web3SwordURI // Event containing the contract specifics and raw log

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
func (it *Web3SwordURIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Web3SwordURI)
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
		it.Event = new(Web3SwordURI)
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
func (it *Web3SwordURIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Web3SwordURIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Web3SwordURI represents a URI event raised by the Web3Sword contract.
type Web3SwordURI struct {
	Value string
	Id    *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterURI is a free log retrieval operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_Web3Sword *Web3SwordFilterer) FilterURI(opts *bind.FilterOpts, id []*big.Int) (*Web3SwordURIIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Web3Sword.contract.FilterLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return &Web3SwordURIIterator{contract: _Web3Sword.contract, event: "URI", logs: logs, sub: sub}, nil
}

// WatchURI is a free log subscription operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_Web3Sword *Web3SwordFilterer) WatchURI(opts *bind.WatchOpts, sink chan<- *Web3SwordURI, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Web3Sword.contract.WatchLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Web3SwordURI)
				if err := _Web3Sword.contract.UnpackLog(event, "URI", log); err != nil {
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

// ParseURI is a log parse operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_Web3Sword *Web3SwordFilterer) ParseURI(log types.Log) (*Web3SwordURI, error) {
	event := new(Web3SwordURI)
	if err := _Web3Sword.contract.UnpackLog(event, "URI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Web3SwordUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Web3Sword contract.
type Web3SwordUpgradedIterator struct {
	Event *Web3SwordUpgraded // Event containing the contract specifics and raw log

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
func (it *Web3SwordUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Web3SwordUpgraded)
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
		it.Event = new(Web3SwordUpgraded)
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
func (it *Web3SwordUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Web3SwordUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Web3SwordUpgraded represents a Upgraded event raised by the Web3Sword contract.
type Web3SwordUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Web3Sword *Web3SwordFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*Web3SwordUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Web3Sword.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &Web3SwordUpgradedIterator{contract: _Web3Sword.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Web3Sword *Web3SwordFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *Web3SwordUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Web3Sword.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Web3SwordUpgraded)
				if err := _Web3Sword.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Web3Sword *Web3SwordFilterer) ParseUpgraded(log types.Log) (*Web3SwordUpgraded, error) {
	event := new(Web3SwordUpgraded)
	if err := _Web3Sword.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Web3SwordWithdrawalIterator is returned from FilterWithdrawal and is used to iterate over the raw logs and unpacked data for Withdrawal events raised by the Web3Sword contract.
type Web3SwordWithdrawalIterator struct {
	Event *Web3SwordWithdrawal // Event containing the contract specifics and raw log

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
func (it *Web3SwordWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Web3SwordWithdrawal)
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
		it.Event = new(Web3SwordWithdrawal)
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
func (it *Web3SwordWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Web3SwordWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Web3SwordWithdrawal represents a Withdrawal event raised by the Web3Sword contract.
type Web3SwordWithdrawal struct {
	Arg0  common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterWithdrawal is a free log retrieval operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address indexed arg0, uint256 value)
func (_Web3Sword *Web3SwordFilterer) FilterWithdrawal(opts *bind.FilterOpts, arg0 []common.Address) (*Web3SwordWithdrawalIterator, error) {

	var arg0Rule []interface{}
	for _, arg0Item := range arg0 {
		arg0Rule = append(arg0Rule, arg0Item)
	}

	logs, sub, err := _Web3Sword.contract.FilterLogs(opts, "Withdrawal", arg0Rule)
	if err != nil {
		return nil, err
	}
	return &Web3SwordWithdrawalIterator{contract: _Web3Sword.contract, event: "Withdrawal", logs: logs, sub: sub}, nil
}

// WatchWithdrawal is a free log subscription operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address indexed arg0, uint256 value)
func (_Web3Sword *Web3SwordFilterer) WatchWithdrawal(opts *bind.WatchOpts, sink chan<- *Web3SwordWithdrawal, arg0 []common.Address) (event.Subscription, error) {

	var arg0Rule []interface{}
	for _, arg0Item := range arg0 {
		arg0Rule = append(arg0Rule, arg0Item)
	}

	logs, sub, err := _Web3Sword.contract.WatchLogs(opts, "Withdrawal", arg0Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Web3SwordWithdrawal)
				if err := _Web3Sword.contract.UnpackLog(event, "Withdrawal", log); err != nil {
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

// ParseWithdrawal is a log parse operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address indexed arg0, uint256 value)
func (_Web3Sword *Web3SwordFilterer) ParseWithdrawal(log types.Log) (*Web3SwordWithdrawal, error) {
	event := new(Web3SwordWithdrawal)
	if err := _Web3Sword.contract.UnpackLog(event, "Withdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

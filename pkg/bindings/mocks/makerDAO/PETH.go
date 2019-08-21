// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package makerDAO

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

// PETHABI is the input ABI used to generate the binding from.
const PETHABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"stop\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"guy\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"owner_\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"dst\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"symbol_\",\"type\":\"bytes32\"}],\"name\":\"DSToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"guy\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"name_\",\"type\":\"bytes32\"}],\"name\":\"setName\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stopped\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"authority_\",\"type\":\"address\"}],\"name\":\"setAuthority\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"guy\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"dst\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"dst\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"push\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"dst\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"move\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"start\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"authority\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"guy\",\"type\":\"address\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"guy\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"pull\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"guy\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"guy\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"authority\",\"type\":\"address\"}],\"name\":\"LogSetAuthority\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"LogSetOwner\",\"type\":\"event\"},{\"anonymous\":true,\"inputs\":[{\"indexed\":true,\"name\":\"sig\",\"type\":\"bytes4\"},{\"indexed\":true,\"name\":\"guy\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"foo\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"bar\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"wad\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"fax\",\"type\":\"bytes\"}],\"name\":\"LogNote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"src\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"guy\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"src\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"dst\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// PETHBin is the compiled bytecode used for deploying new contracts.
var PETHBin = "0x60606040908152601260065560006007819055600160a060020a033316808252600160205282822082905590805560048054600160a060020a03191682179055907fce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed94905160405180910390a2610d598061007a6000396000f3006060604052600436106101535763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306fdde03811461015857806307da68f51461017d578063095ea7b31461019257806313af4035146101c857806318160ddd146101e757806323b872dd146101fa5780632f0573d914610222578063313ce5671461023857806340c10f191461024b57806342966c681461026d5780635ac801fe1461028357806370a082311461029957806375f12b21146102b85780637a9e5e4b146102cb5780638da5cb5b146102ea57806395d89b41146103195780639dc29fac1461032c578063a0712d681461034e578063a9059cbb14610364578063b753a98c14610386578063bb35783b146103a8578063be9a6555146103d0578063bf7e214f146103e3578063daea85c5146103f6578063dd62ed3e14610415578063f2d5d56b1461043a575b600080fd5b341561016357600080fd5b61016b61045c565b60405190815260200160405180910390f35b341561018857600080fd5b610190610462565b005b341561019d57600080fd5b6101b4600160a060020a0360043516602435610501565b604051901515815260200160405180910390f35b34156101d357600080fd5b610190600160a060020a036004351661052e565b34156101f257600080fd5b61016b6105ad565b341561020557600080fd5b6101b4600160a060020a03600435811690602435166044356105b3565b341561022d57600080fd5b610190600435610728565b341561024357600080fd5b61016b61072d565b341561025657600080fd5b610190600160a060020a0360043516602435610733565b341561027857600080fd5b6101906004356107f9565b341561028e57600080fd5b610190600435610806565b34156102a457600080fd5b61016b600160a060020a036004351661082c565b34156102c357600080fd5b6101b4610847565b34156102d657600080fd5b610190600160a060020a0360043516610857565b34156102f557600080fd5b6102fd6108d6565b604051600160a060020a03909116815260200160405180910390f35b341561032457600080fd5b61016b6108e5565b341561033757600080fd5b610190600160a060020a03600435166024356108eb565b341561035957600080fd5b610190600435610a59565b341561036f57600080fd5b6101b4600160a060020a0360043516602435610a63565b341561039157600080fd5b610190600160a060020a0360043516602435610a70565b34156103b357600080fd5b610190600160a060020a0360043581169060243516604435610a80565b34156103db57600080fd5b610190610a91565b34156103ee57600080fd5b6102fd610b2a565b341561040157600080fd5b6101b4600160a060020a0360043516610b39565b341561042057600080fd5b61016b600160a060020a0360043581169060243516610b5f565b341561044557600080fd5b610190600160a060020a0360043516602435610b8a565b60075481565b61047833600035600160e060020a031916610b95565b151561048357600080fd5b600435602435808233600160a060020a031660008035600160e060020a0319169034903660405183815260406020820181815290820183905260608201848480828437820191505094505050505060405180910390a450506004805474ff0000000000000000000000000000000000000000191660a060020a179055565b60045460009060a060020a900460ff161561051b57600080fd5b6105258383610ca1565b90505b92915050565b61054433600035600160e060020a031916610b95565b151561054f57600080fd5b6004805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a038381169190911791829055167fce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed9460405160405180910390a250565b60005490565b60045460009060a060020a900460ff16156105cd57600080fd5b33600160a060020a031684600160a060020a0316141580156106175750600160a060020a038085166000908152600260209081526040808320339094168352929052205460001914155b1561067557600160a060020a038085166000908152600260209081526040808320339094168352929052205461064d9083610d0d565b600160a060020a03808616600090815260026020908152604080832033909416835292905220555b600160a060020a0384166000908152600160205260409020546106989083610d0d565b600160a060020a0380861660009081526001602052604080822093909355908516815220546106c79083610d1d565b600160a060020a03808516600081815260016020526040908190209390935591908616907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a35060019392505050565b600555565b60065481565b61074933600035600160e060020a031916610b95565b151561075457600080fd5b60045460a060020a900460ff161561076b57600080fd5b600160a060020a03821660009081526001602052604090205461078e9082610d1d565b600160a060020a038316600090815260016020526040812091909155546107b59082610d1d565b600055600160a060020a0382167f0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d41213968858260405190815260200160405180910390a25050565b61080333826108eb565b50565b61081c33600035600160e060020a031916610b95565b151561082757600080fd5b600755565b600160a060020a031660009081526001602052604090205490565b60045460a060020a900460ff1681565b61086d33600035600160e060020a031916610b95565b151561087857600080fd5b6003805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a038381169190911791829055167f1abebea81bfa2637f28358c371278fb15ede7ea8dd28d2e03b112ff6d936ada460405160405180910390a250565b600454600160a060020a031681565b60055481565b61090133600035600160e060020a031916610b95565b151561090c57600080fd5b60045460a060020a900460ff161561092357600080fd5b33600160a060020a031682600160a060020a03161415801561096d5750600160a060020a038083166000908152600260209081526040808320339094168352929052205460001914155b156109cb57600160a060020a03808316600090815260026020908152604080832033909416835292905220546109a39082610d0d565b600160a060020a03808416600090815260026020908152604080832033909416835292905220555b600160a060020a0382166000908152600160205260409020546109ee9082610d0d565b600160a060020a03831660009081526001602052604081209190915554610a159082610d0d565b600055600160a060020a0382167fcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca58260405190815260200160405180910390a25050565b6108033382610733565b60006105253384846105b3565b610a7b3383836105b3565b505050565b610a8b8383836105b3565b50505050565b610aa733600035600160e060020a031916610b95565b1515610ab257600080fd5b600435602435808233600160a060020a031660008035600160e060020a0319169034903660405183815260406020820181815290820183905260608201848480828437820191505094505050505060405180910390a450506004805474ff000000000000000000000000000000000000000019169055565b600354600160a060020a031681565b60045460009060a060020a900460ff1615610b5357600080fd5b61052882600019610ca1565b600160a060020a03918216600090815260026020908152604080832093909416825291909152205490565b610a7b8233836105b3565b600030600160a060020a031683600160a060020a03161415610bb957506001610528565b600454600160a060020a0384811691161415610bd757506001610528565b600354600160a060020a03161515610bf157506000610528565b600354600160a060020a031663b70096138430856000604051602001526040517c010000000000000000000000000000000000000000000000000000000063ffffffff8616028152600160a060020a039384166004820152919092166024820152600160e060020a03199091166044820152606401602060405180830381600087803b1515610c7f57600080fd5b6102c65a03f11515610c9057600080fd5b505050604051805190509050610528565b600160a060020a03338116600081815260026020908152604080832094871680845294909152808220859055909291907f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259085905190815260200160405180910390a350600192915050565b8082038281111561052857600080fd5b8082018281101561052857600080fd00a165627a7a72305820b518f8c07bf669bba68db0ff1c747b887d4e8335782174f82650d4572fe0cdf20029"

// DeployPETH deploys a new Ethereum contract, binding an instance of PETH to it.
func DeployPETH(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PETH, error) {
	parsed, err := abi.JSON(strings.NewReader(PETHABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PETHBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PETH{PETHCaller: PETHCaller{contract: contract}, PETHTransactor: PETHTransactor{contract: contract}, PETHFilterer: PETHFilterer{contract: contract}}, nil
}

// PETH is an auto generated Go binding around an Ethereum contract.
type PETH struct {
	PETHCaller     // Read-only binding to the contract
	PETHTransactor // Write-only binding to the contract
	PETHFilterer   // Log filterer for contract events
}

// PETHCaller is an auto generated read-only Go binding around an Ethereum contract.
type PETHCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PETHTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PETHTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PETHFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PETHFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PETHSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PETHSession struct {
	Contract     *PETH             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PETHCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PETHCallerSession struct {
	Contract *PETHCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PETHTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PETHTransactorSession struct {
	Contract     *PETHTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PETHRaw is an auto generated low-level Go binding around an Ethereum contract.
type PETHRaw struct {
	Contract *PETH // Generic contract binding to access the raw methods on
}

// PETHCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PETHCallerRaw struct {
	Contract *PETHCaller // Generic read-only contract binding to access the raw methods on
}

// PETHTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PETHTransactorRaw struct {
	Contract *PETHTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPETH creates a new instance of PETH, bound to a specific deployed contract.
func NewPETH(address common.Address, backend bind.ContractBackend) (*PETH, error) {
	contract, err := bindPETH(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PETH{PETHCaller: PETHCaller{contract: contract}, PETHTransactor: PETHTransactor{contract: contract}, PETHFilterer: PETHFilterer{contract: contract}}, nil
}

// NewPETHCaller creates a new read-only instance of PETH, bound to a specific deployed contract.
func NewPETHCaller(address common.Address, caller bind.ContractCaller) (*PETHCaller, error) {
	contract, err := bindPETH(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PETHCaller{contract: contract}, nil
}

// NewPETHTransactor creates a new write-only instance of PETH, bound to a specific deployed contract.
func NewPETHTransactor(address common.Address, transactor bind.ContractTransactor) (*PETHTransactor, error) {
	contract, err := bindPETH(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PETHTransactor{contract: contract}, nil
}

// NewPETHFilterer creates a new log filterer instance of PETH, bound to a specific deployed contract.
func NewPETHFilterer(address common.Address, filterer bind.ContractFilterer) (*PETHFilterer, error) {
	contract, err := bindPETH(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PETHFilterer{contract: contract}, nil
}

// bindPETH binds a generic wrapper to an already deployed contract.
func bindPETH(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PETHABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PETH *PETHRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PETH.Contract.PETHCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PETH *PETHRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PETH.Contract.PETHTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PETH *PETHRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PETH.Contract.PETHTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PETH *PETHCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PETH.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PETH *PETHTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PETH.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PETH *PETHTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PETH.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address src, address guy) constant returns(uint256)
func (_PETH *PETHCaller) Allowance(opts *bind.CallOpts, src common.Address, guy common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PETH.contract.Call(opts, out, "allowance", src, guy)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address src, address guy) constant returns(uint256)
func (_PETH *PETHSession) Allowance(src common.Address, guy common.Address) (*big.Int, error) {
	return _PETH.Contract.Allowance(&_PETH.CallOpts, src, guy)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address src, address guy) constant returns(uint256)
func (_PETH *PETHCallerSession) Allowance(src common.Address, guy common.Address) (*big.Int, error) {
	return _PETH.Contract.Allowance(&_PETH.CallOpts, src, guy)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_PETH *PETHCaller) Authority(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PETH.contract.Call(opts, out, "authority")
	return *ret0, err
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_PETH *PETHSession) Authority() (common.Address, error) {
	return _PETH.Contract.Authority(&_PETH.CallOpts)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_PETH *PETHCallerSession) Authority() (common.Address, error) {
	return _PETH.Contract.Authority(&_PETH.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address src) constant returns(uint256)
func (_PETH *PETHCaller) BalanceOf(opts *bind.CallOpts, src common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PETH.contract.Call(opts, out, "balanceOf", src)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address src) constant returns(uint256)
func (_PETH *PETHSession) BalanceOf(src common.Address) (*big.Int, error) {
	return _PETH.Contract.BalanceOf(&_PETH.CallOpts, src)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address src) constant returns(uint256)
func (_PETH *PETHCallerSession) BalanceOf(src common.Address) (*big.Int, error) {
	return _PETH.Contract.BalanceOf(&_PETH.CallOpts, src)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_PETH *PETHCaller) Decimals(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PETH.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_PETH *PETHSession) Decimals() (*big.Int, error) {
	return _PETH.Contract.Decimals(&_PETH.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_PETH *PETHCallerSession) Decimals() (*big.Int, error) {
	return _PETH.Contract.Decimals(&_PETH.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(bytes32)
func (_PETH *PETHCaller) Name(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _PETH.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(bytes32)
func (_PETH *PETHSession) Name() ([32]byte, error) {
	return _PETH.Contract.Name(&_PETH.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(bytes32)
func (_PETH *PETHCallerSession) Name() ([32]byte, error) {
	return _PETH.Contract.Name(&_PETH.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_PETH *PETHCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PETH.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_PETH *PETHSession) Owner() (common.Address, error) {
	return _PETH.Contract.Owner(&_PETH.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_PETH *PETHCallerSession) Owner() (common.Address, error) {
	return _PETH.Contract.Owner(&_PETH.CallOpts)
}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() constant returns(bool)
func (_PETH *PETHCaller) Stopped(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PETH.contract.Call(opts, out, "stopped")
	return *ret0, err
}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() constant returns(bool)
func (_PETH *PETHSession) Stopped() (bool, error) {
	return _PETH.Contract.Stopped(&_PETH.CallOpts)
}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() constant returns(bool)
func (_PETH *PETHCallerSession) Stopped() (bool, error) {
	return _PETH.Contract.Stopped(&_PETH.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(bytes32)
func (_PETH *PETHCaller) Symbol(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _PETH.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(bytes32)
func (_PETH *PETHSession) Symbol() ([32]byte, error) {
	return _PETH.Contract.Symbol(&_PETH.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(bytes32)
func (_PETH *PETHCallerSession) Symbol() ([32]byte, error) {
	return _PETH.Contract.Symbol(&_PETH.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_PETH *PETHCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PETH.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_PETH *PETHSession) TotalSupply() (*big.Int, error) {
	return _PETH.Contract.TotalSupply(&_PETH.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_PETH *PETHCallerSession) TotalSupply() (*big.Int, error) {
	return _PETH.Contract.TotalSupply(&_PETH.CallOpts)
}

// DSToken is a paid mutator transaction binding the contract method 0x2f0573d9.
//
// Solidity: function DSToken(bytes32 symbol_) returns()
func (_PETH *PETHTransactor) DSToken(opts *bind.TransactOpts, symbol_ [32]byte) (*types.Transaction, error) {
	return _PETH.contract.Transact(opts, "DSToken", symbol_)
}

// DSToken is a paid mutator transaction binding the contract method 0x2f0573d9.
//
// Solidity: function DSToken(bytes32 symbol_) returns()
func (_PETH *PETHSession) DSToken(symbol_ [32]byte) (*types.Transaction, error) {
	return _PETH.Contract.DSToken(&_PETH.TransactOpts, symbol_)
}

// DSToken is a paid mutator transaction binding the contract method 0x2f0573d9.
//
// Solidity: function DSToken(bytes32 symbol_) returns()
func (_PETH *PETHTransactorSession) DSToken(symbol_ [32]byte) (*types.Transaction, error) {
	return _PETH.Contract.DSToken(&_PETH.TransactOpts, symbol_)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address guy, uint256 wad) returns(bool)
func (_PETH *PETHTransactor) Approve(opts *bind.TransactOpts, guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _PETH.contract.Transact(opts, "approve", guy, wad)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address guy, uint256 wad) returns(bool)
func (_PETH *PETHSession) Approve(guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _PETH.Contract.Approve(&_PETH.TransactOpts, guy, wad)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address guy, uint256 wad) returns(bool)
func (_PETH *PETHTransactorSession) Approve(guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _PETH.Contract.Approve(&_PETH.TransactOpts, guy, wad)
}

// Approve0 is a paid mutator transaction binding the contract method 0xdaea85c5.
//
// Solidity: function approve(address guy) returns(bool)
func (_PETH *PETHTransactor) Approve0(opts *bind.TransactOpts, guy common.Address) (*types.Transaction, error) {
	return _PETH.contract.Transact(opts, "approve0", guy)
}

// Approve0 is a paid mutator transaction binding the contract method 0xdaea85c5.
//
// Solidity: function approve(address guy) returns(bool)
func (_PETH *PETHSession) Approve0(guy common.Address) (*types.Transaction, error) {
	return _PETH.Contract.Approve0(&_PETH.TransactOpts, guy)
}

// Approve0 is a paid mutator transaction binding the contract method 0xdaea85c5.
//
// Solidity: function approve(address guy) returns(bool)
func (_PETH *PETHTransactorSession) Approve0(guy common.Address) (*types.Transaction, error) {
	return _PETH.Contract.Approve0(&_PETH.TransactOpts, guy)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 wad) returns()
func (_PETH *PETHTransactor) Burn(opts *bind.TransactOpts, wad *big.Int) (*types.Transaction, error) {
	return _PETH.contract.Transact(opts, "burn", wad)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 wad) returns()
func (_PETH *PETHSession) Burn(wad *big.Int) (*types.Transaction, error) {
	return _PETH.Contract.Burn(&_PETH.TransactOpts, wad)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 wad) returns()
func (_PETH *PETHTransactorSession) Burn(wad *big.Int) (*types.Transaction, error) {
	return _PETH.Contract.Burn(&_PETH.TransactOpts, wad)
}

// Burn0 is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address guy, uint256 wad) returns()
func (_PETH *PETHTransactor) Burn0(opts *bind.TransactOpts, guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _PETH.contract.Transact(opts, "burn0", guy, wad)
}

// Burn0 is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address guy, uint256 wad) returns()
func (_PETH *PETHSession) Burn0(guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _PETH.Contract.Burn0(&_PETH.TransactOpts, guy, wad)
}

// Burn0 is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address guy, uint256 wad) returns()
func (_PETH *PETHTransactorSession) Burn0(guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _PETH.Contract.Burn0(&_PETH.TransactOpts, guy, wad)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address guy, uint256 wad) returns()
func (_PETH *PETHTransactor) Mint(opts *bind.TransactOpts, guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _PETH.contract.Transact(opts, "mint", guy, wad)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address guy, uint256 wad) returns()
func (_PETH *PETHSession) Mint(guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _PETH.Contract.Mint(&_PETH.TransactOpts, guy, wad)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address guy, uint256 wad) returns()
func (_PETH *PETHTransactorSession) Mint(guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _PETH.Contract.Mint(&_PETH.TransactOpts, guy, wad)
}

// Mint0 is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 wad) returns()
func (_PETH *PETHTransactor) Mint0(opts *bind.TransactOpts, wad *big.Int) (*types.Transaction, error) {
	return _PETH.contract.Transact(opts, "mint0", wad)
}

// Mint0 is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 wad) returns()
func (_PETH *PETHSession) Mint0(wad *big.Int) (*types.Transaction, error) {
	return _PETH.Contract.Mint0(&_PETH.TransactOpts, wad)
}

// Mint0 is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 wad) returns()
func (_PETH *PETHTransactorSession) Mint0(wad *big.Int) (*types.Transaction, error) {
	return _PETH.Contract.Mint0(&_PETH.TransactOpts, wad)
}

// Move is a paid mutator transaction binding the contract method 0xbb35783b.
//
// Solidity: function move(address src, address dst, uint256 wad) returns()
func (_PETH *PETHTransactor) Move(opts *bind.TransactOpts, src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _PETH.contract.Transact(opts, "move", src, dst, wad)
}

// Move is a paid mutator transaction binding the contract method 0xbb35783b.
//
// Solidity: function move(address src, address dst, uint256 wad) returns()
func (_PETH *PETHSession) Move(src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _PETH.Contract.Move(&_PETH.TransactOpts, src, dst, wad)
}

// Move is a paid mutator transaction binding the contract method 0xbb35783b.
//
// Solidity: function move(address src, address dst, uint256 wad) returns()
func (_PETH *PETHTransactorSession) Move(src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _PETH.Contract.Move(&_PETH.TransactOpts, src, dst, wad)
}

// Pull is a paid mutator transaction binding the contract method 0xf2d5d56b.
//
// Solidity: function pull(address src, uint256 wad) returns()
func (_PETH *PETHTransactor) Pull(opts *bind.TransactOpts, src common.Address, wad *big.Int) (*types.Transaction, error) {
	return _PETH.contract.Transact(opts, "pull", src, wad)
}

// Pull is a paid mutator transaction binding the contract method 0xf2d5d56b.
//
// Solidity: function pull(address src, uint256 wad) returns()
func (_PETH *PETHSession) Pull(src common.Address, wad *big.Int) (*types.Transaction, error) {
	return _PETH.Contract.Pull(&_PETH.TransactOpts, src, wad)
}

// Pull is a paid mutator transaction binding the contract method 0xf2d5d56b.
//
// Solidity: function pull(address src, uint256 wad) returns()
func (_PETH *PETHTransactorSession) Pull(src common.Address, wad *big.Int) (*types.Transaction, error) {
	return _PETH.Contract.Pull(&_PETH.TransactOpts, src, wad)
}

// Push is a paid mutator transaction binding the contract method 0xb753a98c.
//
// Solidity: function push(address dst, uint256 wad) returns()
func (_PETH *PETHTransactor) Push(opts *bind.TransactOpts, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _PETH.contract.Transact(opts, "push", dst, wad)
}

// Push is a paid mutator transaction binding the contract method 0xb753a98c.
//
// Solidity: function push(address dst, uint256 wad) returns()
func (_PETH *PETHSession) Push(dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _PETH.Contract.Push(&_PETH.TransactOpts, dst, wad)
}

// Push is a paid mutator transaction binding the contract method 0xb753a98c.
//
// Solidity: function push(address dst, uint256 wad) returns()
func (_PETH *PETHTransactorSession) Push(dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _PETH.Contract.Push(&_PETH.TransactOpts, dst, wad)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(address authority_) returns()
func (_PETH *PETHTransactor) SetAuthority(opts *bind.TransactOpts, authority_ common.Address) (*types.Transaction, error) {
	return _PETH.contract.Transact(opts, "setAuthority", authority_)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(address authority_) returns()
func (_PETH *PETHSession) SetAuthority(authority_ common.Address) (*types.Transaction, error) {
	return _PETH.Contract.SetAuthority(&_PETH.TransactOpts, authority_)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(address authority_) returns()
func (_PETH *PETHTransactorSession) SetAuthority(authority_ common.Address) (*types.Transaction, error) {
	return _PETH.Contract.SetAuthority(&_PETH.TransactOpts, authority_)
}

// SetName is a paid mutator transaction binding the contract method 0x5ac801fe.
//
// Solidity: function setName(bytes32 name_) returns()
func (_PETH *PETHTransactor) SetName(opts *bind.TransactOpts, name_ [32]byte) (*types.Transaction, error) {
	return _PETH.contract.Transact(opts, "setName", name_)
}

// SetName is a paid mutator transaction binding the contract method 0x5ac801fe.
//
// Solidity: function setName(bytes32 name_) returns()
func (_PETH *PETHSession) SetName(name_ [32]byte) (*types.Transaction, error) {
	return _PETH.Contract.SetName(&_PETH.TransactOpts, name_)
}

// SetName is a paid mutator transaction binding the contract method 0x5ac801fe.
//
// Solidity: function setName(bytes32 name_) returns()
func (_PETH *PETHTransactorSession) SetName(name_ [32]byte) (*types.Transaction, error) {
	return _PETH.Contract.SetName(&_PETH.TransactOpts, name_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address owner_) returns()
func (_PETH *PETHTransactor) SetOwner(opts *bind.TransactOpts, owner_ common.Address) (*types.Transaction, error) {
	return _PETH.contract.Transact(opts, "setOwner", owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address owner_) returns()
func (_PETH *PETHSession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _PETH.Contract.SetOwner(&_PETH.TransactOpts, owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address owner_) returns()
func (_PETH *PETHTransactorSession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _PETH.Contract.SetOwner(&_PETH.TransactOpts, owner_)
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns()
func (_PETH *PETHTransactor) Start(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PETH.contract.Transact(opts, "start")
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns()
func (_PETH *PETHSession) Start() (*types.Transaction, error) {
	return _PETH.Contract.Start(&_PETH.TransactOpts)
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns()
func (_PETH *PETHTransactorSession) Start() (*types.Transaction, error) {
	return _PETH.Contract.Start(&_PETH.TransactOpts)
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_PETH *PETHTransactor) Stop(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PETH.contract.Transact(opts, "stop")
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_PETH *PETHSession) Stop() (*types.Transaction, error) {
	return _PETH.Contract.Stop(&_PETH.TransactOpts)
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_PETH *PETHTransactorSession) Stop() (*types.Transaction, error) {
	return _PETH.Contract.Stop(&_PETH.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 wad) returns(bool)
func (_PETH *PETHTransactor) Transfer(opts *bind.TransactOpts, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _PETH.contract.Transact(opts, "transfer", dst, wad)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 wad) returns(bool)
func (_PETH *PETHSession) Transfer(dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _PETH.Contract.Transfer(&_PETH.TransactOpts, dst, wad)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 wad) returns(bool)
func (_PETH *PETHTransactorSession) Transfer(dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _PETH.Contract.Transfer(&_PETH.TransactOpts, dst, wad)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 wad) returns(bool)
func (_PETH *PETHTransactor) TransferFrom(opts *bind.TransactOpts, src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _PETH.contract.Transact(opts, "transferFrom", src, dst, wad)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 wad) returns(bool)
func (_PETH *PETHSession) TransferFrom(src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _PETH.Contract.TransferFrom(&_PETH.TransactOpts, src, dst, wad)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 wad) returns(bool)
func (_PETH *PETHTransactorSession) TransferFrom(src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _PETH.Contract.TransferFrom(&_PETH.TransactOpts, src, dst, wad)
}

// PETHApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the PETH contract.
type PETHApprovalIterator struct {
	Event *PETHApproval // Event containing the contract specifics and raw log

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
func (it *PETHApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PETHApproval)
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
		it.Event = new(PETHApproval)
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
func (it *PETHApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PETHApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PETHApproval represents a Approval event raised by the PETH contract.
type PETHApproval struct {
	Src common.Address
	Guy common.Address
	Wad *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed src, address indexed guy, uint256 wad)
func (_PETH *PETHFilterer) FilterApproval(opts *bind.FilterOpts, src []common.Address, guy []common.Address) (*PETHApprovalIterator, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var guyRule []interface{}
	for _, guyItem := range guy {
		guyRule = append(guyRule, guyItem)
	}

	logs, sub, err := _PETH.contract.FilterLogs(opts, "Approval", srcRule, guyRule)
	if err != nil {
		return nil, err
	}
	return &PETHApprovalIterator{contract: _PETH.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed src, address indexed guy, uint256 wad)
func (_PETH *PETHFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *PETHApproval, src []common.Address, guy []common.Address) (event.Subscription, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var guyRule []interface{}
	for _, guyItem := range guy {
		guyRule = append(guyRule, guyItem)
	}

	logs, sub, err := _PETH.contract.WatchLogs(opts, "Approval", srcRule, guyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PETHApproval)
				if err := _PETH.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed src, address indexed guy, uint256 wad)
func (_PETH *PETHFilterer) ParseApproval(log types.Log) (*PETHApproval, error) {
	event := new(PETHApproval)
	if err := _PETH.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PETHBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the PETH contract.
type PETHBurnIterator struct {
	Event *PETHBurn // Event containing the contract specifics and raw log

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
func (it *PETHBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PETHBurn)
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
		it.Event = new(PETHBurn)
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
func (it *PETHBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PETHBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PETHBurn represents a Burn event raised by the PETH contract.
type PETHBurn struct {
	Guy common.Address
	Wad *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed guy, uint256 wad)
func (_PETH *PETHFilterer) FilterBurn(opts *bind.FilterOpts, guy []common.Address) (*PETHBurnIterator, error) {

	var guyRule []interface{}
	for _, guyItem := range guy {
		guyRule = append(guyRule, guyItem)
	}

	logs, sub, err := _PETH.contract.FilterLogs(opts, "Burn", guyRule)
	if err != nil {
		return nil, err
	}
	return &PETHBurnIterator{contract: _PETH.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed guy, uint256 wad)
func (_PETH *PETHFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *PETHBurn, guy []common.Address) (event.Subscription, error) {

	var guyRule []interface{}
	for _, guyItem := range guy {
		guyRule = append(guyRule, guyItem)
	}

	logs, sub, err := _PETH.contract.WatchLogs(opts, "Burn", guyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PETHBurn)
				if err := _PETH.contract.UnpackLog(event, "Burn", log); err != nil {
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

// ParseBurn is a log parse operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed guy, uint256 wad)
func (_PETH *PETHFilterer) ParseBurn(log types.Log) (*PETHBurn, error) {
	event := new(PETHBurn)
	if err := _PETH.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PETHLogSetAuthorityIterator is returned from FilterLogSetAuthority and is used to iterate over the raw logs and unpacked data for LogSetAuthority events raised by the PETH contract.
type PETHLogSetAuthorityIterator struct {
	Event *PETHLogSetAuthority // Event containing the contract specifics and raw log

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
func (it *PETHLogSetAuthorityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PETHLogSetAuthority)
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
		it.Event = new(PETHLogSetAuthority)
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
func (it *PETHLogSetAuthorityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PETHLogSetAuthorityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PETHLogSetAuthority represents a LogSetAuthority event raised by the PETH contract.
type PETHLogSetAuthority struct {
	Authority common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLogSetAuthority is a free log retrieval operation binding the contract event 0x1abebea81bfa2637f28358c371278fb15ede7ea8dd28d2e03b112ff6d936ada4.
//
// Solidity: event LogSetAuthority(address indexed authority)
func (_PETH *PETHFilterer) FilterLogSetAuthority(opts *bind.FilterOpts, authority []common.Address) (*PETHLogSetAuthorityIterator, error) {

	var authorityRule []interface{}
	for _, authorityItem := range authority {
		authorityRule = append(authorityRule, authorityItem)
	}

	logs, sub, err := _PETH.contract.FilterLogs(opts, "LogSetAuthority", authorityRule)
	if err != nil {
		return nil, err
	}
	return &PETHLogSetAuthorityIterator{contract: _PETH.contract, event: "LogSetAuthority", logs: logs, sub: sub}, nil
}

// WatchLogSetAuthority is a free log subscription operation binding the contract event 0x1abebea81bfa2637f28358c371278fb15ede7ea8dd28d2e03b112ff6d936ada4.
//
// Solidity: event LogSetAuthority(address indexed authority)
func (_PETH *PETHFilterer) WatchLogSetAuthority(opts *bind.WatchOpts, sink chan<- *PETHLogSetAuthority, authority []common.Address) (event.Subscription, error) {

	var authorityRule []interface{}
	for _, authorityItem := range authority {
		authorityRule = append(authorityRule, authorityItem)
	}

	logs, sub, err := _PETH.contract.WatchLogs(opts, "LogSetAuthority", authorityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PETHLogSetAuthority)
				if err := _PETH.contract.UnpackLog(event, "LogSetAuthority", log); err != nil {
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

// ParseLogSetAuthority is a log parse operation binding the contract event 0x1abebea81bfa2637f28358c371278fb15ede7ea8dd28d2e03b112ff6d936ada4.
//
// Solidity: event LogSetAuthority(address indexed authority)
func (_PETH *PETHFilterer) ParseLogSetAuthority(log types.Log) (*PETHLogSetAuthority, error) {
	event := new(PETHLogSetAuthority)
	if err := _PETH.contract.UnpackLog(event, "LogSetAuthority", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PETHLogSetOwnerIterator is returned from FilterLogSetOwner and is used to iterate over the raw logs and unpacked data for LogSetOwner events raised by the PETH contract.
type PETHLogSetOwnerIterator struct {
	Event *PETHLogSetOwner // Event containing the contract specifics and raw log

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
func (it *PETHLogSetOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PETHLogSetOwner)
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
		it.Event = new(PETHLogSetOwner)
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
func (it *PETHLogSetOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PETHLogSetOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PETHLogSetOwner represents a LogSetOwner event raised by the PETH contract.
type PETHLogSetOwner struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterLogSetOwner is a free log retrieval operation binding the contract event 0xce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed94.
//
// Solidity: event LogSetOwner(address indexed owner)
func (_PETH *PETHFilterer) FilterLogSetOwner(opts *bind.FilterOpts, owner []common.Address) (*PETHLogSetOwnerIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _PETH.contract.FilterLogs(opts, "LogSetOwner", ownerRule)
	if err != nil {
		return nil, err
	}
	return &PETHLogSetOwnerIterator{contract: _PETH.contract, event: "LogSetOwner", logs: logs, sub: sub}, nil
}

// WatchLogSetOwner is a free log subscription operation binding the contract event 0xce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed94.
//
// Solidity: event LogSetOwner(address indexed owner)
func (_PETH *PETHFilterer) WatchLogSetOwner(opts *bind.WatchOpts, sink chan<- *PETHLogSetOwner, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _PETH.contract.WatchLogs(opts, "LogSetOwner", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PETHLogSetOwner)
				if err := _PETH.contract.UnpackLog(event, "LogSetOwner", log); err != nil {
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

// ParseLogSetOwner is a log parse operation binding the contract event 0xce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed94.
//
// Solidity: event LogSetOwner(address indexed owner)
func (_PETH *PETHFilterer) ParseLogSetOwner(log types.Log) (*PETHLogSetOwner, error) {
	event := new(PETHLogSetOwner)
	if err := _PETH.contract.UnpackLog(event, "LogSetOwner", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PETHMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the PETH contract.
type PETHMintIterator struct {
	Event *PETHMint // Event containing the contract specifics and raw log

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
func (it *PETHMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PETHMint)
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
		it.Event = new(PETHMint)
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
func (it *PETHMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PETHMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PETHMint represents a Mint event raised by the PETH contract.
type PETHMint struct {
	Guy common.Address
	Wad *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed guy, uint256 wad)
func (_PETH *PETHFilterer) FilterMint(opts *bind.FilterOpts, guy []common.Address) (*PETHMintIterator, error) {

	var guyRule []interface{}
	for _, guyItem := range guy {
		guyRule = append(guyRule, guyItem)
	}

	logs, sub, err := _PETH.contract.FilterLogs(opts, "Mint", guyRule)
	if err != nil {
		return nil, err
	}
	return &PETHMintIterator{contract: _PETH.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed guy, uint256 wad)
func (_PETH *PETHFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *PETHMint, guy []common.Address) (event.Subscription, error) {

	var guyRule []interface{}
	for _, guyItem := range guy {
		guyRule = append(guyRule, guyItem)
	}

	logs, sub, err := _PETH.contract.WatchLogs(opts, "Mint", guyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PETHMint)
				if err := _PETH.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed guy, uint256 wad)
func (_PETH *PETHFilterer) ParseMint(log types.Log) (*PETHMint, error) {
	event := new(PETHMint)
	if err := _PETH.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PETHTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the PETH contract.
type PETHTransferIterator struct {
	Event *PETHTransfer // Event containing the contract specifics and raw log

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
func (it *PETHTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PETHTransfer)
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
		it.Event = new(PETHTransfer)
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
func (it *PETHTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PETHTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PETHTransfer represents a Transfer event raised by the PETH contract.
type PETHTransfer struct {
	Src common.Address
	Dst common.Address
	Wad *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed src, address indexed dst, uint256 wad)
func (_PETH *PETHFilterer) FilterTransfer(opts *bind.FilterOpts, src []common.Address, dst []common.Address) (*PETHTransferIterator, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _PETH.contract.FilterLogs(opts, "Transfer", srcRule, dstRule)
	if err != nil {
		return nil, err
	}
	return &PETHTransferIterator{contract: _PETH.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed src, address indexed dst, uint256 wad)
func (_PETH *PETHFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *PETHTransfer, src []common.Address, dst []common.Address) (event.Subscription, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _PETH.contract.WatchLogs(opts, "Transfer", srcRule, dstRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PETHTransfer)
				if err := _PETH.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed src, address indexed dst, uint256 wad)
func (_PETH *PETHFilterer) ParseTransfer(log types.Log) (*PETHTransfer, error) {
	event := new(PETHTransfer)
	if err := _PETH.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

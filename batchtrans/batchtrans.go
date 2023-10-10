// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package batchtrans

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
	_ = abi.ConvertType
)

// BatchtransMetaData contains all meta data concerning the Batchtrans contract.
var BatchtransMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"handOut\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061002d61002261003260201b60201c565b61003a60201b60201c565b6100fe565b600033905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b610a308061010d6000396000f3fe6080604052600436106100595760003560e01c806312065fe0146100655780633ccfd60b146100905780635da4e7b4146100a7578063715018a6146100c35780638da5cb5b146100da578063f2fde38b1461010557610060565b3661006057005b600080fd5b34801561007157600080fd5b5061007a61012e565b60405161008791906104ef565b60405180910390f35b34801561009c57600080fd5b506100a561013e565b005b6100c160048036038101906100bc91906105cf565b61018f565b005b3480156100cf57600080fd5b506100d861029d565b005b3480156100e657600080fd5b506100ef6102e0565b6040516100fc9190610691565b60405180910390f35b34801561011157600080fd5b5061012c600480360381019061012791906106d8565b610309565b005b600061013861038c565b47905090565b61014661038c565b3373ffffffffffffffffffffffffffffffffffffffff166108fc479081150290604051600060405180830381858888f1935050505015801561018c573d6000803e3d6000fd5b50565b61019761038c565b600084849050116101dd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101d490610788565b60405180910390fd5b60005b84849050811015610296578484828181106101fe576101fd6107a8565b5b905060200201602081019061021391906106d8565b73ffffffffffffffffffffffffffffffffffffffff166108fc655af3107a4000858585818110610246576102456107a8565b5b905060200201356102579190610806565b9081150290604051600060405180830381858888f19350505050158015610282573d6000803e3d6000fd5b50808061028e90610848565b9150506101e0565b5050505050565b6102a561038c565b6040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102d7906108dc565b60405180910390fd5b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b61031161038c565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610380576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103779061096e565b60405180910390fd5b6103898161040a565b50565b6103946104ce565b73ffffffffffffffffffffffffffffffffffffffff166103b26102e0565b73ffffffffffffffffffffffffffffffffffffffff1614610408576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103ff906109da565b60405180910390fd5b565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600033905090565b6000819050919050565b6104e9816104d6565b82525050565b600060208201905061050460008301846104e0565b92915050565b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b60008083601f84011261053957610538610514565b5b8235905067ffffffffffffffff81111561055657610555610519565b5b6020830191508360208202830111156105725761057161051e565b5b9250929050565b60008083601f84011261058f5761058e610514565b5b8235905067ffffffffffffffff8111156105ac576105ab610519565b5b6020830191508360208202830111156105c8576105c761051e565b5b9250929050565b600080600080604085870312156105e9576105e861050a565b5b600085013567ffffffffffffffff8111156106075761060661050f565b5b61061387828801610523565b9450945050602085013567ffffffffffffffff8111156106365761063561050f565b5b61064287828801610579565b925092505092959194509250565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061067b82610650565b9050919050565b61068b81610670565b82525050565b60006020820190506106a66000830184610682565b92915050565b6106b581610670565b81146106c057600080fd5b50565b6000813590506106d2816106ac565b92915050565b6000602082840312156106ee576106ed61050a565b5b60006106fc848285016106c3565b91505092915050565b600082825260208201905092915050565b7f706c6561736520656e7465722074686520616363657074616e6365206164647260008201527f6573730000000000000000000000000000000000000000000000000000000000602082015250565b6000610772602383610705565b915061077d82610716565b604082019050919050565b600060208201905081810360008301526107a181610765565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000610811826104d6565b915061081c836104d6565b925082820261082a816104d6565b91508282048414831517610841576108406107d7565b5b5092915050565b6000610853826104d6565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610885576108846107d7565b5b600182019050919050565b7f72656e6f756e63696e67206f776e6572736869702069732064697361626c6564600082015250565b60006108c6602083610705565b91506108d182610890565b602082019050919050565b600060208201905081810360008301526108f5816108b9565b9050919050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b6000610958602683610705565b9150610963826108fc565b604082019050919050565b600060208201905081810360008301526109878161094b565b9050919050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b60006109c4602083610705565b91506109cf8261098e565b602082019050919050565b600060208201905081810360008301526109f3816109b7565b905091905056fea2646970667358221220ac83569652f3f760f468920a477707b4af49b541d0b47c8d7f69d9deba54a00964736f6c63430008120033",
}

// BatchtransABI is the input ABI used to generate the binding from.
// Deprecated: Use BatchtransMetaData.ABI instead.
var BatchtransABI = BatchtransMetaData.ABI

// BatchtransBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BatchtransMetaData.Bin instead.
var BatchtransBin = BatchtransMetaData.Bin

// DeployBatchtrans deploys a new Ethereum contract, binding an instance of Batchtrans to it.
func DeployBatchtrans(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Batchtrans, error) {
	parsed, err := BatchtransMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BatchtransBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Batchtrans{BatchtransCaller: BatchtransCaller{contract: contract}, BatchtransTransactor: BatchtransTransactor{contract: contract}, BatchtransFilterer: BatchtransFilterer{contract: contract}}, nil
}

// Batchtrans is an auto generated Go binding around an Ethereum contract.
type Batchtrans struct {
	BatchtransCaller     // Read-only binding to the contract
	BatchtransTransactor // Write-only binding to the contract
	BatchtransFilterer   // Log filterer for contract events
}

// BatchtransCaller is an auto generated read-only Go binding around an Ethereum contract.
type BatchtransCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BatchtransTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BatchtransTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BatchtransFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BatchtransFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BatchtransSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BatchtransSession struct {
	Contract     *Batchtrans       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BatchtransCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BatchtransCallerSession struct {
	Contract *BatchtransCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// BatchtransTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BatchtransTransactorSession struct {
	Contract     *BatchtransTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// BatchtransRaw is an auto generated low-level Go binding around an Ethereum contract.
type BatchtransRaw struct {
	Contract *Batchtrans // Generic contract binding to access the raw methods on
}

// BatchtransCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BatchtransCallerRaw struct {
	Contract *BatchtransCaller // Generic read-only contract binding to access the raw methods on
}

// BatchtransTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BatchtransTransactorRaw struct {
	Contract *BatchtransTransactor // Generic write-only contract binding to access the raw methods on
}



// NewBatchtrans creates a new instance of Batchtrans, bound to a specific deployed contract.
func NewBatchtrans(address common.Address, backend bind.ContractBackend) (*Batchtrans, error) {
	contract, err := bindBatchtrans(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Batchtrans{BatchtransCaller: BatchtransCaller{contract: contract}, BatchtransTransactor: BatchtransTransactor{contract: contract}, BatchtransFilterer: BatchtransFilterer{contract: contract}}, nil
}

// NewBatchtransCaller creates a new read-only instance of Batchtrans, bound to a specific deployed contract.
func NewBatchtransCaller(address common.Address, caller bind.ContractCaller) (*BatchtransCaller, error) {
	contract, err := bindBatchtrans(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BatchtransCaller{contract: contract}, nil
}

// NewBatchtransTransactor creates a new write-only instance of Batchtrans, bound to a specific deployed contract.
func NewBatchtransTransactor(address common.Address, transactor bind.ContractTransactor) (*BatchtransTransactor, error) {
	contract, err := bindBatchtrans(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BatchtransTransactor{contract: contract}, nil
}

// NewBatchtransFilterer creates a new log filterer instance of Batchtrans, bound to a specific deployed contract.
func NewBatchtransFilterer(address common.Address, filterer bind.ContractFilterer) (*BatchtransFilterer, error) {
	contract, err := bindBatchtrans(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BatchtransFilterer{contract: contract}, nil
}

// bindBatchtrans binds a generic wrapper to an already deployed contract.
func bindBatchtrans(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BatchtransMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Batchtrans *BatchtransRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Batchtrans.Contract.BatchtransCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Batchtrans *BatchtransRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Batchtrans.Contract.BatchtransTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Batchtrans *BatchtransRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Batchtrans.Contract.BatchtransTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Batchtrans *BatchtransCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Batchtrans.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Batchtrans *BatchtransTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Batchtrans.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Batchtrans *BatchtransTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Batchtrans.Contract.contract.Transact(opts, method, params...)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Batchtrans *BatchtransCaller) GetBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Batchtrans.contract.Call(opts, &out, "getBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Batchtrans *BatchtransSession) GetBalance() (*big.Int, error) {
	return _Batchtrans.Contract.GetBalance(&_Batchtrans.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Batchtrans *BatchtransCallerSession) GetBalance() (*big.Int, error) {
	return _Batchtrans.Contract.GetBalance(&_Batchtrans.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Batchtrans *BatchtransCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Batchtrans.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Batchtrans *BatchtransSession) Owner() (common.Address, error) {
	return _Batchtrans.Contract.Owner(&_Batchtrans.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Batchtrans *BatchtransCallerSession) Owner() (common.Address, error) {
	return _Batchtrans.Contract.Owner(&_Batchtrans.CallOpts)
}

// HandOut is a paid mutator transaction binding the contract method 0x5da4e7b4.
//
// Solidity: function handOut(address[] addresses, uint256[] amounts) payable returns()
func (_Batchtrans *BatchtransTransactor) HandOut(opts *bind.TransactOpts, addresses []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Batchtrans.contract.Transact(opts, "handOut", addresses, amounts)
}

// HandOut is a paid mutator transaction binding the contract method 0x5da4e7b4.
//
// Solidity: function handOut(address[] addresses, uint256[] amounts) payable returns()
func (_Batchtrans *BatchtransSession) HandOut(addresses []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Batchtrans.Contract.HandOut(&_Batchtrans.TransactOpts, addresses, amounts)
}

// HandOut is a paid mutator transaction binding the contract method 0x5da4e7b4.
//
// Solidity: function handOut(address[] addresses, uint256[] amounts) payable returns()
func (_Batchtrans *BatchtransTransactorSession) HandOut(addresses []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Batchtrans.Contract.HandOut(&_Batchtrans.TransactOpts, addresses, amounts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Batchtrans *BatchtransTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Batchtrans.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Batchtrans *BatchtransSession) RenounceOwnership() (*types.Transaction, error) {
	return _Batchtrans.Contract.RenounceOwnership(&_Batchtrans.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Batchtrans *BatchtransTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Batchtrans.Contract.RenounceOwnership(&_Batchtrans.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Batchtrans *BatchtransTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Batchtrans.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Batchtrans *BatchtransSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Batchtrans.Contract.TransferOwnership(&_Batchtrans.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Batchtrans *BatchtransTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Batchtrans.Contract.TransferOwnership(&_Batchtrans.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Batchtrans *BatchtransTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Batchtrans.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Batchtrans *BatchtransSession) Withdraw() (*types.Transaction, error) {
	return _Batchtrans.Contract.Withdraw(&_Batchtrans.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Batchtrans *BatchtransTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Batchtrans.Contract.Withdraw(&_Batchtrans.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Batchtrans *BatchtransTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Batchtrans.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Batchtrans *BatchtransSession) Receive() (*types.Transaction, error) {
	return _Batchtrans.Contract.Receive(&_Batchtrans.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Batchtrans *BatchtransTransactorSession) Receive() (*types.Transaction, error) {
	return _Batchtrans.Contract.Receive(&_Batchtrans.TransactOpts)
}

// BatchtransOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Batchtrans contract.
type BatchtransOwnershipTransferredIterator struct {
	Event *BatchtransOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BatchtransOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BatchtransOwnershipTransferred)
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
		it.Event = new(BatchtransOwnershipTransferred)
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
func (it *BatchtransOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BatchtransOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BatchtransOwnershipTransferred represents a OwnershipTransferred event raised by the Batchtrans contract.
type BatchtransOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Batchtrans *BatchtransFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BatchtransOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Batchtrans.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BatchtransOwnershipTransferredIterator{contract: _Batchtrans.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Batchtrans *BatchtransFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BatchtransOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Batchtrans.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BatchtransOwnershipTransferred)
				if err := _Batchtrans.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Batchtrans *BatchtransFilterer) ParseOwnershipTransferred(log types.Log) (*BatchtransOwnershipTransferred, error) {
	event := new(BatchtransOwnershipTransferred)
	if err := _Batchtrans.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

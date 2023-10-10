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

// MarketplaceMetaData contains all meta data concerning the Marketplace contract.
var MarketplaceMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"groupId\",\"type\":\"uint256\"}],\"name\":\"Buy\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"groupId\",\"type\":\"uint256\"}],\"name\":\"BuyFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"CreateFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"status\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"groupName\",\"type\":\"bytes\"}],\"name\":\"CreateGroupFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"groupName\",\"type\":\"bytes\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"CreateGroupSuccess\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"CreateSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"CreateSuccess\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"DeleteFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"status\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"DeleteGroupFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"DeleteGroupSuccess\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"DeleteSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"DeleteSuccess\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"groupId\",\"type\":\"uint256\"}],\"name\":\"Delist\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"groupId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"List\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"failReason\",\"type\":\"bytes\"}],\"name\":\"MirrorFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"MirrorSuccess\",\"type\":\"event\"}]",
}

// MarketplaceABI is the input ABI used to generate the binding from.
// Deprecated: Use MarketplaceMetaData.ABI instead.
var MarketplaceABI = MarketplaceMetaData.ABI

// Marketplace is an auto generated Go binding around an Ethereum contract.
type Marketplace struct {
	MarketplaceCaller     // Read-only binding to the contract
	MarketplaceTransactor // Write-only binding to the contract
	MarketplaceFilterer   // Log filterer for contract events
}

// MarketplaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type MarketplaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketplaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MarketplaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketplaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MarketplaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketplaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MarketplaceSession struct {
	Contract     *Marketplace      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MarketplaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MarketplaceCallerSession struct {
	Contract *MarketplaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// MarketplaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MarketplaceTransactorSession struct {
	Contract     *MarketplaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// MarketplaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type MarketplaceRaw struct {
	Contract *Marketplace // Generic contract binding to access the raw methods on
}

// MarketplaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MarketplaceCallerRaw struct {
	Contract *MarketplaceCaller // Generic read-only contract binding to access the raw methods on
}

// MarketplaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MarketplaceTransactorRaw struct {
	Contract *MarketplaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMarketplace creates a new instance of Marketplace, bound to a specific deployed contract.
func NewMarketplace(address common.Address, backend bind.ContractBackend) (*Marketplace, error) {
	contract, err := bindMarketplace(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Marketplace{MarketplaceCaller: MarketplaceCaller{contract: contract}, MarketplaceTransactor: MarketplaceTransactor{contract: contract}, MarketplaceFilterer: MarketplaceFilterer{contract: contract}}, nil
}

// NewMarketplaceCaller creates a new read-only instance of Marketplace, bound to a specific deployed contract.
func NewMarketplaceCaller(address common.Address, caller bind.ContractCaller) (*MarketplaceCaller, error) {
	contract, err := bindMarketplace(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MarketplaceCaller{contract: contract}, nil
}

// NewMarketplaceTransactor creates a new write-only instance of Marketplace, bound to a specific deployed contract.
func NewMarketplaceTransactor(address common.Address, transactor bind.ContractTransactor) (*MarketplaceTransactor, error) {
	contract, err := bindMarketplace(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MarketplaceTransactor{contract: contract}, nil
}

// NewMarketplaceFilterer creates a new log filterer instance of Marketplace, bound to a specific deployed contract.
func NewMarketplaceFilterer(address common.Address, filterer bind.ContractFilterer) (*MarketplaceFilterer, error) {
	contract, err := bindMarketplace(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MarketplaceFilterer{contract: contract}, nil
}

// bindMarketplace binds a generic wrapper to an already deployed contract.
func bindMarketplace(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MarketplaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Marketplace *MarketplaceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Marketplace.Contract.MarketplaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Marketplace *MarketplaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Marketplace.Contract.MarketplaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Marketplace *MarketplaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Marketplace.Contract.MarketplaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Marketplace *MarketplaceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Marketplace.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Marketplace *MarketplaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Marketplace.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Marketplace *MarketplaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Marketplace.Contract.contract.Transact(opts, method, params...)
}

// MarketplaceBuyIterator is returned from FilterBuy and is used to iterate over the raw logs and unpacked data for Buy events raised by the Marketplace contract.
type MarketplaceBuyIterator struct {
	Event *MarketplaceBuy // Event containing the contract specifics and raw log

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
func (it *MarketplaceBuyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceBuy)
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
		it.Event = new(MarketplaceBuy)
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
func (it *MarketplaceBuyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceBuyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceBuy represents a Buy event raised by the Marketplace contract.
type MarketplaceBuy struct {
	Buyer   common.Address
	GroupId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBuy is a free log retrieval operation binding the contract event 0xe3d4187f6ca4248660cc0ac8b8056515bac4a8132be2eca31d6d0cc170722a7e.
//
// Solidity: event Buy(address indexed buyer, uint256 indexed groupId)
func (_Marketplace *MarketplaceFilterer) FilterBuy(opts *bind.FilterOpts, buyer []common.Address, groupId []*big.Int) (*MarketplaceBuyIterator, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var groupIdRule []interface{}
	for _, groupIdItem := range groupId {
		groupIdRule = append(groupIdRule, groupIdItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "Buy", buyerRule, groupIdRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceBuyIterator{contract: _Marketplace.contract, event: "Buy", logs: logs, sub: sub}, nil
}

// WatchBuy is a free log subscription operation binding the contract event 0xe3d4187f6ca4248660cc0ac8b8056515bac4a8132be2eca31d6d0cc170722a7e.
//
// Solidity: event Buy(address indexed buyer, uint256 indexed groupId)
func (_Marketplace *MarketplaceFilterer) WatchBuy(opts *bind.WatchOpts, sink chan<- *MarketplaceBuy, buyer []common.Address, groupId []*big.Int) (event.Subscription, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var groupIdRule []interface{}
	for _, groupIdItem := range groupId {
		groupIdRule = append(groupIdRule, groupIdItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "Buy", buyerRule, groupIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceBuy)
				if err := _Marketplace.contract.UnpackLog(event, "Buy", log); err != nil {
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

// ParseBuy is a log parse operation binding the contract event 0xe3d4187f6ca4248660cc0ac8b8056515bac4a8132be2eca31d6d0cc170722a7e.
//
// Solidity: event Buy(address indexed buyer, uint256 indexed groupId)
func (_Marketplace *MarketplaceFilterer) ParseBuy(log types.Log) (*MarketplaceBuy, error) {
	event := new(MarketplaceBuy)
	if err := _Marketplace.contract.UnpackLog(event, "Buy", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceBuyFailedIterator is returned from FilterBuyFailed and is used to iterate over the raw logs and unpacked data for BuyFailed events raised by the Marketplace contract.
type MarketplaceBuyFailedIterator struct {
	Event *MarketplaceBuyFailed // Event containing the contract specifics and raw log

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
func (it *MarketplaceBuyFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceBuyFailed)
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
		it.Event = new(MarketplaceBuyFailed)
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
func (it *MarketplaceBuyFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceBuyFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceBuyFailed represents a BuyFailed event raised by the Marketplace contract.
type MarketplaceBuyFailed struct {
	Buyer   common.Address
	GroupId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBuyFailed is a free log retrieval operation binding the contract event 0x6cb8a65861aadde2ff0a249436b975cacdda3930978595076da8e8011a6fe215.
//
// Solidity: event BuyFailed(address indexed buyer, uint256 indexed groupId)
func (_Marketplace *MarketplaceFilterer) FilterBuyFailed(opts *bind.FilterOpts, buyer []common.Address, groupId []*big.Int) (*MarketplaceBuyFailedIterator, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var groupIdRule []interface{}
	for _, groupIdItem := range groupId {
		groupIdRule = append(groupIdRule, groupIdItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "BuyFailed", buyerRule, groupIdRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceBuyFailedIterator{contract: _Marketplace.contract, event: "BuyFailed", logs: logs, sub: sub}, nil
}

// WatchBuyFailed is a free log subscription operation binding the contract event 0x6cb8a65861aadde2ff0a249436b975cacdda3930978595076da8e8011a6fe215.
//
// Solidity: event BuyFailed(address indexed buyer, uint256 indexed groupId)
func (_Marketplace *MarketplaceFilterer) WatchBuyFailed(opts *bind.WatchOpts, sink chan<- *MarketplaceBuyFailed, buyer []common.Address, groupId []*big.Int) (event.Subscription, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var groupIdRule []interface{}
	for _, groupIdItem := range groupId {
		groupIdRule = append(groupIdRule, groupIdItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "BuyFailed", buyerRule, groupIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceBuyFailed)
				if err := _Marketplace.contract.UnpackLog(event, "BuyFailed", log); err != nil {
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

// ParseBuyFailed is a log parse operation binding the contract event 0x6cb8a65861aadde2ff0a249436b975cacdda3930978595076da8e8011a6fe215.
//
// Solidity: event BuyFailed(address indexed buyer, uint256 indexed groupId)
func (_Marketplace *MarketplaceFilterer) ParseBuyFailed(log types.Log) (*MarketplaceBuyFailed, error) {
	event := new(MarketplaceBuyFailed)
	if err := _Marketplace.contract.UnpackLog(event, "BuyFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceCreateFailedIterator is returned from FilterCreateFailed and is used to iterate over the raw logs and unpacked data for CreateFailed events raised by the Marketplace contract.
type MarketplaceCreateFailedIterator struct {
	Event *MarketplaceCreateFailed // Event containing the contract specifics and raw log

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
func (it *MarketplaceCreateFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceCreateFailed)
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
		it.Event = new(MarketplaceCreateFailed)
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
func (it *MarketplaceCreateFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceCreateFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceCreateFailed represents a CreateFailed event raised by the Marketplace contract.
type MarketplaceCreateFailed struct {
	Creator common.Address
	Id      *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterCreateFailed is a free log retrieval operation binding the contract event 0x8fcc3c7c07d973811455683d6926125ec03d4e5a8f3dae28a45dc3eca5ec1cf6.
//
// Solidity: event CreateFailed(address indexed creator, uint256 indexed id)
func (_Marketplace *MarketplaceFilterer) FilterCreateFailed(opts *bind.FilterOpts, creator []common.Address, id []*big.Int) (*MarketplaceCreateFailedIterator, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "CreateFailed", creatorRule, idRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceCreateFailedIterator{contract: _Marketplace.contract, event: "CreateFailed", logs: logs, sub: sub}, nil
}

// WatchCreateFailed is a free log subscription operation binding the contract event 0x8fcc3c7c07d973811455683d6926125ec03d4e5a8f3dae28a45dc3eca5ec1cf6.
//
// Solidity: event CreateFailed(address indexed creator, uint256 indexed id)
func (_Marketplace *MarketplaceFilterer) WatchCreateFailed(opts *bind.WatchOpts, sink chan<- *MarketplaceCreateFailed, creator []common.Address, id []*big.Int) (event.Subscription, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "CreateFailed", creatorRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceCreateFailed)
				if err := _Marketplace.contract.UnpackLog(event, "CreateFailed", log); err != nil {
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

// ParseCreateFailed is a log parse operation binding the contract event 0x8fcc3c7c07d973811455683d6926125ec03d4e5a8f3dae28a45dc3eca5ec1cf6.
//
// Solidity: event CreateFailed(address indexed creator, uint256 indexed id)
func (_Marketplace *MarketplaceFilterer) ParseCreateFailed(log types.Log) (*MarketplaceCreateFailed, error) {
	event := new(MarketplaceCreateFailed)
	if err := _Marketplace.contract.UnpackLog(event, "CreateFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceCreateGroupFailedIterator is returned from FilterCreateGroupFailed and is used to iterate over the raw logs and unpacked data for CreateGroupFailed events raised by the Marketplace contract.
type MarketplaceCreateGroupFailedIterator struct {
	Event *MarketplaceCreateGroupFailed // Event containing the contract specifics and raw log

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
func (it *MarketplaceCreateGroupFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceCreateGroupFailed)
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
		it.Event = new(MarketplaceCreateGroupFailed)
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
func (it *MarketplaceCreateGroupFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceCreateGroupFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceCreateGroupFailed represents a CreateGroupFailed event raised by the Marketplace contract.
type MarketplaceCreateGroupFailed struct {
	Status    uint32
	GroupName []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCreateGroupFailed is a free log retrieval operation binding the contract event 0xac28ffffbbf8321968aab9d10c38db9b3ca47dc521bf7dc3b33819ee03ae7be2.
//
// Solidity: event CreateGroupFailed(uint32 status, bytes groupName)
func (_Marketplace *MarketplaceFilterer) FilterCreateGroupFailed(opts *bind.FilterOpts) (*MarketplaceCreateGroupFailedIterator, error) {

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "CreateGroupFailed")
	if err != nil {
		return nil, err
	}
	return &MarketplaceCreateGroupFailedIterator{contract: _Marketplace.contract, event: "CreateGroupFailed", logs: logs, sub: sub}, nil
}

// WatchCreateGroupFailed is a free log subscription operation binding the contract event 0xac28ffffbbf8321968aab9d10c38db9b3ca47dc521bf7dc3b33819ee03ae7be2.
//
// Solidity: event CreateGroupFailed(uint32 status, bytes groupName)
func (_Marketplace *MarketplaceFilterer) WatchCreateGroupFailed(opts *bind.WatchOpts, sink chan<- *MarketplaceCreateGroupFailed) (event.Subscription, error) {

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "CreateGroupFailed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceCreateGroupFailed)
				if err := _Marketplace.contract.UnpackLog(event, "CreateGroupFailed", log); err != nil {
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

// ParseCreateGroupFailed is a log parse operation binding the contract event 0xac28ffffbbf8321968aab9d10c38db9b3ca47dc521bf7dc3b33819ee03ae7be2.
//
// Solidity: event CreateGroupFailed(uint32 status, bytes groupName)
func (_Marketplace *MarketplaceFilterer) ParseCreateGroupFailed(log types.Log) (*MarketplaceCreateGroupFailed, error) {
	event := new(MarketplaceCreateGroupFailed)
	if err := _Marketplace.contract.UnpackLog(event, "CreateGroupFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceCreateGroupSuccessIterator is returned from FilterCreateGroupSuccess and is used to iterate over the raw logs and unpacked data for CreateGroupSuccess events raised by the Marketplace contract.
type MarketplaceCreateGroupSuccessIterator struct {
	Event *MarketplaceCreateGroupSuccess // Event containing the contract specifics and raw log

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
func (it *MarketplaceCreateGroupSuccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceCreateGroupSuccess)
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
		it.Event = new(MarketplaceCreateGroupSuccess)
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
func (it *MarketplaceCreateGroupSuccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceCreateGroupSuccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceCreateGroupSuccess represents a CreateGroupSuccess event raised by the Marketplace contract.
type MarketplaceCreateGroupSuccess struct {
	GroupName []byte
	TokenId   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCreateGroupSuccess is a free log retrieval operation binding the contract event 0xc389e58637ae2cdf73c61105637a66283b5690b8764500cf30316e92b314ecda.
//
// Solidity: event CreateGroupSuccess(bytes groupName, uint256 indexed tokenId)
func (_Marketplace *MarketplaceFilterer) FilterCreateGroupSuccess(opts *bind.FilterOpts, tokenId []*big.Int) (*MarketplaceCreateGroupSuccessIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "CreateGroupSuccess", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceCreateGroupSuccessIterator{contract: _Marketplace.contract, event: "CreateGroupSuccess", logs: logs, sub: sub}, nil
}

// WatchCreateGroupSuccess is a free log subscription operation binding the contract event 0xc389e58637ae2cdf73c61105637a66283b5690b8764500cf30316e92b314ecda.
//
// Solidity: event CreateGroupSuccess(bytes groupName, uint256 indexed tokenId)
func (_Marketplace *MarketplaceFilterer) WatchCreateGroupSuccess(opts *bind.WatchOpts, sink chan<- *MarketplaceCreateGroupSuccess, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "CreateGroupSuccess", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceCreateGroupSuccess)
				if err := _Marketplace.contract.UnpackLog(event, "CreateGroupSuccess", log); err != nil {
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

// ParseCreateGroupSuccess is a log parse operation binding the contract event 0xc389e58637ae2cdf73c61105637a66283b5690b8764500cf30316e92b314ecda.
//
// Solidity: event CreateGroupSuccess(bytes groupName, uint256 indexed tokenId)
func (_Marketplace *MarketplaceFilterer) ParseCreateGroupSuccess(log types.Log) (*MarketplaceCreateGroupSuccess, error) {
	event := new(MarketplaceCreateGroupSuccess)
	if err := _Marketplace.contract.UnpackLog(event, "CreateGroupSuccess", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceCreateSubmittedIterator is returned from FilterCreateSubmitted and is used to iterate over the raw logs and unpacked data for CreateSubmitted events raised by the Marketplace contract.
type MarketplaceCreateSubmittedIterator struct {
	Event *MarketplaceCreateSubmitted // Event containing the contract specifics and raw log

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
func (it *MarketplaceCreateSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceCreateSubmitted)
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
		it.Event = new(MarketplaceCreateSubmitted)
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
func (it *MarketplaceCreateSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceCreateSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceCreateSubmitted represents a CreateSubmitted event raised by the Marketplace contract.
type MarketplaceCreateSubmitted struct {
	Owner    common.Address
	Operator common.Address
	Name     string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCreateSubmitted is a free log retrieval operation binding the contract event 0x03618de88ddbe65bdfeecf02937d8a478d24563b1a24783f2c4471fed5bde3b8.
//
// Solidity: event CreateSubmitted(address indexed owner, address indexed operator, string name)
func (_Marketplace *MarketplaceFilterer) FilterCreateSubmitted(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*MarketplaceCreateSubmittedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "CreateSubmitted", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceCreateSubmittedIterator{contract: _Marketplace.contract, event: "CreateSubmitted", logs: logs, sub: sub}, nil
}

// WatchCreateSubmitted is a free log subscription operation binding the contract event 0x03618de88ddbe65bdfeecf02937d8a478d24563b1a24783f2c4471fed5bde3b8.
//
// Solidity: event CreateSubmitted(address indexed owner, address indexed operator, string name)
func (_Marketplace *MarketplaceFilterer) WatchCreateSubmitted(opts *bind.WatchOpts, sink chan<- *MarketplaceCreateSubmitted, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "CreateSubmitted", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceCreateSubmitted)
				if err := _Marketplace.contract.UnpackLog(event, "CreateSubmitted", log); err != nil {
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

// ParseCreateSubmitted is a log parse operation binding the contract event 0x03618de88ddbe65bdfeecf02937d8a478d24563b1a24783f2c4471fed5bde3b8.
//
// Solidity: event CreateSubmitted(address indexed owner, address indexed operator, string name)
func (_Marketplace *MarketplaceFilterer) ParseCreateSubmitted(log types.Log) (*MarketplaceCreateSubmitted, error) {
	event := new(MarketplaceCreateSubmitted)
	if err := _Marketplace.contract.UnpackLog(event, "CreateSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceCreateSuccessIterator is returned from FilterCreateSuccess and is used to iterate over the raw logs and unpacked data for CreateSuccess events raised by the Marketplace contract.
type MarketplaceCreateSuccessIterator struct {
	Event *MarketplaceCreateSuccess // Event containing the contract specifics and raw log

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
func (it *MarketplaceCreateSuccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceCreateSuccess)
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
		it.Event = new(MarketplaceCreateSuccess)
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
func (it *MarketplaceCreateSuccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceCreateSuccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceCreateSuccess represents a CreateSuccess event raised by the Marketplace contract.
type MarketplaceCreateSuccess struct {
	Creator common.Address
	Id      *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterCreateSuccess is a free log retrieval operation binding the contract event 0x272354479408536e7d87366f3ae6c7beac79cf606b31522fb7b1582c8b29a8fc.
//
// Solidity: event CreateSuccess(address indexed creator, uint256 indexed id)
func (_Marketplace *MarketplaceFilterer) FilterCreateSuccess(opts *bind.FilterOpts, creator []common.Address, id []*big.Int) (*MarketplaceCreateSuccessIterator, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "CreateSuccess", creatorRule, idRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceCreateSuccessIterator{contract: _Marketplace.contract, event: "CreateSuccess", logs: logs, sub: sub}, nil
}

// WatchCreateSuccess is a free log subscription operation binding the contract event 0x272354479408536e7d87366f3ae6c7beac79cf606b31522fb7b1582c8b29a8fc.
//
// Solidity: event CreateSuccess(address indexed creator, uint256 indexed id)
func (_Marketplace *MarketplaceFilterer) WatchCreateSuccess(opts *bind.WatchOpts, sink chan<- *MarketplaceCreateSuccess, creator []common.Address, id []*big.Int) (event.Subscription, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "CreateSuccess", creatorRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceCreateSuccess)
				if err := _Marketplace.contract.UnpackLog(event, "CreateSuccess", log); err != nil {
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

// ParseCreateSuccess is a log parse operation binding the contract event 0x272354479408536e7d87366f3ae6c7beac79cf606b31522fb7b1582c8b29a8fc.
//
// Solidity: event CreateSuccess(address indexed creator, uint256 indexed id)
func (_Marketplace *MarketplaceFilterer) ParseCreateSuccess(log types.Log) (*MarketplaceCreateSuccess, error) {
	event := new(MarketplaceCreateSuccess)
	if err := _Marketplace.contract.UnpackLog(event, "CreateSuccess", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceDeleteFailedIterator is returned from FilterDeleteFailed and is used to iterate over the raw logs and unpacked data for DeleteFailed events raised by the Marketplace contract.
type MarketplaceDeleteFailedIterator struct {
	Event *MarketplaceDeleteFailed // Event containing the contract specifics and raw log

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
func (it *MarketplaceDeleteFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceDeleteFailed)
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
		it.Event = new(MarketplaceDeleteFailed)
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
func (it *MarketplaceDeleteFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceDeleteFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceDeleteFailed represents a DeleteFailed event raised by the Marketplace contract.
type MarketplaceDeleteFailed struct {
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDeleteFailed is a free log retrieval operation binding the contract event 0x0a1ce445987f88624924ba9be18aa8a4541a3989b0f608dce437a027efbb3ab4.
//
// Solidity: event DeleteFailed(uint256 indexed id)
func (_Marketplace *MarketplaceFilterer) FilterDeleteFailed(opts *bind.FilterOpts, id []*big.Int) (*MarketplaceDeleteFailedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "DeleteFailed", idRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceDeleteFailedIterator{contract: _Marketplace.contract, event: "DeleteFailed", logs: logs, sub: sub}, nil
}

// WatchDeleteFailed is a free log subscription operation binding the contract event 0x0a1ce445987f88624924ba9be18aa8a4541a3989b0f608dce437a027efbb3ab4.
//
// Solidity: event DeleteFailed(uint256 indexed id)
func (_Marketplace *MarketplaceFilterer) WatchDeleteFailed(opts *bind.WatchOpts, sink chan<- *MarketplaceDeleteFailed, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "DeleteFailed", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceDeleteFailed)
				if err := _Marketplace.contract.UnpackLog(event, "DeleteFailed", log); err != nil {
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

// ParseDeleteFailed is a log parse operation binding the contract event 0x0a1ce445987f88624924ba9be18aa8a4541a3989b0f608dce437a027efbb3ab4.
//
// Solidity: event DeleteFailed(uint256 indexed id)
func (_Marketplace *MarketplaceFilterer) ParseDeleteFailed(log types.Log) (*MarketplaceDeleteFailed, error) {
	event := new(MarketplaceDeleteFailed)
	if err := _Marketplace.contract.UnpackLog(event, "DeleteFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceDeleteGroupFailedIterator is returned from FilterDeleteGroupFailed and is used to iterate over the raw logs and unpacked data for DeleteGroupFailed events raised by the Marketplace contract.
type MarketplaceDeleteGroupFailedIterator struct {
	Event *MarketplaceDeleteGroupFailed // Event containing the contract specifics and raw log

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
func (it *MarketplaceDeleteGroupFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceDeleteGroupFailed)
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
		it.Event = new(MarketplaceDeleteGroupFailed)
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
func (it *MarketplaceDeleteGroupFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceDeleteGroupFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceDeleteGroupFailed represents a DeleteGroupFailed event raised by the Marketplace contract.
type MarketplaceDeleteGroupFailed struct {
	Status  uint32
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDeleteGroupFailed is a free log retrieval operation binding the contract event 0x4b6deb7c34a6970ae46dab01610fe16ad6cbff1a0c42299445a3161087ad5e70.
//
// Solidity: event DeleteGroupFailed(uint32 status, uint256 indexed tokenId)
func (_Marketplace *MarketplaceFilterer) FilterDeleteGroupFailed(opts *bind.FilterOpts, tokenId []*big.Int) (*MarketplaceDeleteGroupFailedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "DeleteGroupFailed", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceDeleteGroupFailedIterator{contract: _Marketplace.contract, event: "DeleteGroupFailed", logs: logs, sub: sub}, nil
}

// WatchDeleteGroupFailed is a free log subscription operation binding the contract event 0x4b6deb7c34a6970ae46dab01610fe16ad6cbff1a0c42299445a3161087ad5e70.
//
// Solidity: event DeleteGroupFailed(uint32 status, uint256 indexed tokenId)
func (_Marketplace *MarketplaceFilterer) WatchDeleteGroupFailed(opts *bind.WatchOpts, sink chan<- *MarketplaceDeleteGroupFailed, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "DeleteGroupFailed", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceDeleteGroupFailed)
				if err := _Marketplace.contract.UnpackLog(event, "DeleteGroupFailed", log); err != nil {
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

// ParseDeleteGroupFailed is a log parse operation binding the contract event 0x4b6deb7c34a6970ae46dab01610fe16ad6cbff1a0c42299445a3161087ad5e70.
//
// Solidity: event DeleteGroupFailed(uint32 status, uint256 indexed tokenId)
func (_Marketplace *MarketplaceFilterer) ParseDeleteGroupFailed(log types.Log) (*MarketplaceDeleteGroupFailed, error) {
	event := new(MarketplaceDeleteGroupFailed)
	if err := _Marketplace.contract.UnpackLog(event, "DeleteGroupFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceDeleteGroupSuccessIterator is returned from FilterDeleteGroupSuccess and is used to iterate over the raw logs and unpacked data for DeleteGroupSuccess events raised by the Marketplace contract.
type MarketplaceDeleteGroupSuccessIterator struct {
	Event *MarketplaceDeleteGroupSuccess // Event containing the contract specifics and raw log

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
func (it *MarketplaceDeleteGroupSuccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceDeleteGroupSuccess)
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
		it.Event = new(MarketplaceDeleteGroupSuccess)
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
func (it *MarketplaceDeleteGroupSuccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceDeleteGroupSuccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceDeleteGroupSuccess represents a DeleteGroupSuccess event raised by the Marketplace contract.
type MarketplaceDeleteGroupSuccess struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDeleteGroupSuccess is a free log retrieval operation binding the contract event 0xf7440fb2c5cf7e7276db53eea9dcf574e7fea21dbe5cf5ef4efd2498e67682d4.
//
// Solidity: event DeleteGroupSuccess(uint256 indexed tokenId)
func (_Marketplace *MarketplaceFilterer) FilterDeleteGroupSuccess(opts *bind.FilterOpts, tokenId []*big.Int) (*MarketplaceDeleteGroupSuccessIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "DeleteGroupSuccess", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceDeleteGroupSuccessIterator{contract: _Marketplace.contract, event: "DeleteGroupSuccess", logs: logs, sub: sub}, nil
}

// WatchDeleteGroupSuccess is a free log subscription operation binding the contract event 0xf7440fb2c5cf7e7276db53eea9dcf574e7fea21dbe5cf5ef4efd2498e67682d4.
//
// Solidity: event DeleteGroupSuccess(uint256 indexed tokenId)
func (_Marketplace *MarketplaceFilterer) WatchDeleteGroupSuccess(opts *bind.WatchOpts, sink chan<- *MarketplaceDeleteGroupSuccess, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "DeleteGroupSuccess", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceDeleteGroupSuccess)
				if err := _Marketplace.contract.UnpackLog(event, "DeleteGroupSuccess", log); err != nil {
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

// ParseDeleteGroupSuccess is a log parse operation binding the contract event 0xf7440fb2c5cf7e7276db53eea9dcf574e7fea21dbe5cf5ef4efd2498e67682d4.
//
// Solidity: event DeleteGroupSuccess(uint256 indexed tokenId)
func (_Marketplace *MarketplaceFilterer) ParseDeleteGroupSuccess(log types.Log) (*MarketplaceDeleteGroupSuccess, error) {
	event := new(MarketplaceDeleteGroupSuccess)
	if err := _Marketplace.contract.UnpackLog(event, "DeleteGroupSuccess", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceDeleteSubmittedIterator is returned from FilterDeleteSubmitted and is used to iterate over the raw logs and unpacked data for DeleteSubmitted events raised by the Marketplace contract.
type MarketplaceDeleteSubmittedIterator struct {
	Event *MarketplaceDeleteSubmitted // Event containing the contract specifics and raw log

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
func (it *MarketplaceDeleteSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceDeleteSubmitted)
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
		it.Event = new(MarketplaceDeleteSubmitted)
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
func (it *MarketplaceDeleteSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceDeleteSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceDeleteSubmitted represents a DeleteSubmitted event raised by the Marketplace contract.
type MarketplaceDeleteSubmitted struct {
	Owner    common.Address
	Operator common.Address
	Id       *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDeleteSubmitted is a free log retrieval operation binding the contract event 0x52af988231d4d3725f70aa24b8eb7fd0cbb921a08289f7cea967eb6887f661a1.
//
// Solidity: event DeleteSubmitted(address indexed owner, address indexed operator, uint256 indexed id)
func (_Marketplace *MarketplaceFilterer) FilterDeleteSubmitted(opts *bind.FilterOpts, owner []common.Address, operator []common.Address, id []*big.Int) (*MarketplaceDeleteSubmittedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "DeleteSubmitted", ownerRule, operatorRule, idRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceDeleteSubmittedIterator{contract: _Marketplace.contract, event: "DeleteSubmitted", logs: logs, sub: sub}, nil
}

// WatchDeleteSubmitted is a free log subscription operation binding the contract event 0x52af988231d4d3725f70aa24b8eb7fd0cbb921a08289f7cea967eb6887f661a1.
//
// Solidity: event DeleteSubmitted(address indexed owner, address indexed operator, uint256 indexed id)
func (_Marketplace *MarketplaceFilterer) WatchDeleteSubmitted(opts *bind.WatchOpts, sink chan<- *MarketplaceDeleteSubmitted, owner []common.Address, operator []common.Address, id []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "DeleteSubmitted", ownerRule, operatorRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceDeleteSubmitted)
				if err := _Marketplace.contract.UnpackLog(event, "DeleteSubmitted", log); err != nil {
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

// ParseDeleteSubmitted is a log parse operation binding the contract event 0x52af988231d4d3725f70aa24b8eb7fd0cbb921a08289f7cea967eb6887f661a1.
//
// Solidity: event DeleteSubmitted(address indexed owner, address indexed operator, uint256 indexed id)
func (_Marketplace *MarketplaceFilterer) ParseDeleteSubmitted(log types.Log) (*MarketplaceDeleteSubmitted, error) {
	event := new(MarketplaceDeleteSubmitted)
	if err := _Marketplace.contract.UnpackLog(event, "DeleteSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceDeleteSuccessIterator is returned from FilterDeleteSuccess and is used to iterate over the raw logs and unpacked data for DeleteSuccess events raised by the Marketplace contract.
type MarketplaceDeleteSuccessIterator struct {
	Event *MarketplaceDeleteSuccess // Event containing the contract specifics and raw log

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
func (it *MarketplaceDeleteSuccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceDeleteSuccess)
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
		it.Event = new(MarketplaceDeleteSuccess)
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
func (it *MarketplaceDeleteSuccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceDeleteSuccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceDeleteSuccess represents a DeleteSuccess event raised by the Marketplace contract.
type MarketplaceDeleteSuccess struct {
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDeleteSuccess is a free log retrieval operation binding the contract event 0x3cb3b8d70d2584f191448f9b28dd657549ec03b348603e5eb630c285e3b6ce37.
//
// Solidity: event DeleteSuccess(uint256 indexed id)
func (_Marketplace *MarketplaceFilterer) FilterDeleteSuccess(opts *bind.FilterOpts, id []*big.Int) (*MarketplaceDeleteSuccessIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "DeleteSuccess", idRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceDeleteSuccessIterator{contract: _Marketplace.contract, event: "DeleteSuccess", logs: logs, sub: sub}, nil
}

// WatchDeleteSuccess is a free log subscription operation binding the contract event 0x3cb3b8d70d2584f191448f9b28dd657549ec03b348603e5eb630c285e3b6ce37.
//
// Solidity: event DeleteSuccess(uint256 indexed id)
func (_Marketplace *MarketplaceFilterer) WatchDeleteSuccess(opts *bind.WatchOpts, sink chan<- *MarketplaceDeleteSuccess, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "DeleteSuccess", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceDeleteSuccess)
				if err := _Marketplace.contract.UnpackLog(event, "DeleteSuccess", log); err != nil {
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

// ParseDeleteSuccess is a log parse operation binding the contract event 0x3cb3b8d70d2584f191448f9b28dd657549ec03b348603e5eb630c285e3b6ce37.
//
// Solidity: event DeleteSuccess(uint256 indexed id)
func (_Marketplace *MarketplaceFilterer) ParseDeleteSuccess(log types.Log) (*MarketplaceDeleteSuccess, error) {
	event := new(MarketplaceDeleteSuccess)
	if err := _Marketplace.contract.UnpackLog(event, "DeleteSuccess", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceDelistIterator is returned from FilterDelist and is used to iterate over the raw logs and unpacked data for Delist events raised by the Marketplace contract.
type MarketplaceDelistIterator struct {
	Event *MarketplaceDelist // Event containing the contract specifics and raw log

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
func (it *MarketplaceDelistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceDelist)
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
		it.Event = new(MarketplaceDelist)
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
func (it *MarketplaceDelistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceDelistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceDelist represents a Delist event raised by the Marketplace contract.
type MarketplaceDelist struct {
	Owner   common.Address
	GroupId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDelist is a free log retrieval operation binding the contract event 0x8fcc1d45240b67aa8f5859c01c295e240be99a9d5e4c11873bb82cf40be7533c.
//
// Solidity: event Delist(address indexed owner, uint256 indexed groupId)
func (_Marketplace *MarketplaceFilterer) FilterDelist(opts *bind.FilterOpts, owner []common.Address, groupId []*big.Int) (*MarketplaceDelistIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var groupIdRule []interface{}
	for _, groupIdItem := range groupId {
		groupIdRule = append(groupIdRule, groupIdItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "Delist", ownerRule, groupIdRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceDelistIterator{contract: _Marketplace.contract, event: "Delist", logs: logs, sub: sub}, nil
}

// WatchDelist is a free log subscription operation binding the contract event 0x8fcc1d45240b67aa8f5859c01c295e240be99a9d5e4c11873bb82cf40be7533c.
//
// Solidity: event Delist(address indexed owner, uint256 indexed groupId)
func (_Marketplace *MarketplaceFilterer) WatchDelist(opts *bind.WatchOpts, sink chan<- *MarketplaceDelist, owner []common.Address, groupId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var groupIdRule []interface{}
	for _, groupIdItem := range groupId {
		groupIdRule = append(groupIdRule, groupIdItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "Delist", ownerRule, groupIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceDelist)
				if err := _Marketplace.contract.UnpackLog(event, "Delist", log); err != nil {
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

// ParseDelist is a log parse operation binding the contract event 0x8fcc1d45240b67aa8f5859c01c295e240be99a9d5e4c11873bb82cf40be7533c.
//
// Solidity: event Delist(address indexed owner, uint256 indexed groupId)
func (_Marketplace *MarketplaceFilterer) ParseDelist(log types.Log) (*MarketplaceDelist, error) {
	event := new(MarketplaceDelist)
	if err := _Marketplace.contract.UnpackLog(event, "Delist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceListIterator is returned from FilterList and is used to iterate over the raw logs and unpacked data for List events raised by the Marketplace contract.
type MarketplaceListIterator struct {
	Event *MarketplaceList // Event containing the contract specifics and raw log

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
func (it *MarketplaceListIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceList)
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
		it.Event = new(MarketplaceList)
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
func (it *MarketplaceListIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceListIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceList represents a List event raised by the Marketplace contract.
type MarketplaceList struct {
	Owner   common.Address
	GroupId *big.Int
	Price   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterList is a free log retrieval operation binding the contract event 0x1518ffac149698f404e82117efa8b67d99c365490eefe4e0f91856a93ca8d9c1.
//
// Solidity: event List(address indexed owner, uint256 indexed groupId, uint256 price)
func (_Marketplace *MarketplaceFilterer) FilterList(opts *bind.FilterOpts, owner []common.Address, groupId []*big.Int) (*MarketplaceListIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var groupIdRule []interface{}
	for _, groupIdItem := range groupId {
		groupIdRule = append(groupIdRule, groupIdItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "List", ownerRule, groupIdRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceListIterator{contract: _Marketplace.contract, event: "List", logs: logs, sub: sub}, nil
}

// WatchList is a free log subscription operation binding the contract event 0x1518ffac149698f404e82117efa8b67d99c365490eefe4e0f91856a93ca8d9c1.
//
// Solidity: event List(address indexed owner, uint256 indexed groupId, uint256 price)
func (_Marketplace *MarketplaceFilterer) WatchList(opts *bind.WatchOpts, sink chan<- *MarketplaceList, owner []common.Address, groupId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var groupIdRule []interface{}
	for _, groupIdItem := range groupId {
		groupIdRule = append(groupIdRule, groupIdItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "List", ownerRule, groupIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceList)
				if err := _Marketplace.contract.UnpackLog(event, "List", log); err != nil {
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

// ParseList is a log parse operation binding the contract event 0x1518ffac149698f404e82117efa8b67d99c365490eefe4e0f91856a93ca8d9c1.
//
// Solidity: event List(address indexed owner, uint256 indexed groupId, uint256 price)
func (_Marketplace *MarketplaceFilterer) ParseList(log types.Log) (*MarketplaceList, error) {
	event := new(MarketplaceList)
	if err := _Marketplace.contract.UnpackLog(event, "List", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceMirrorFailedIterator is returned from FilterMirrorFailed and is used to iterate over the raw logs and unpacked data for MirrorFailed events raised by the Marketplace contract.
type MarketplaceMirrorFailedIterator struct {
	Event *MarketplaceMirrorFailed // Event containing the contract specifics and raw log

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
func (it *MarketplaceMirrorFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceMirrorFailed)
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
		it.Event = new(MarketplaceMirrorFailed)
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
func (it *MarketplaceMirrorFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceMirrorFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceMirrorFailed represents a MirrorFailed event raised by the Marketplace contract.
type MarketplaceMirrorFailed struct {
	Id         *big.Int
	Owner      common.Address
	FailReason []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMirrorFailed is a free log retrieval operation binding the contract event 0x42bf66fb325bb74dce540406916688333b68d163b61dc9fc22a6892fe73dda01.
//
// Solidity: event MirrorFailed(uint256 indexed id, address indexed owner, bytes failReason)
func (_Marketplace *MarketplaceFilterer) FilterMirrorFailed(opts *bind.FilterOpts, id []*big.Int, owner []common.Address) (*MarketplaceMirrorFailedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "MirrorFailed", idRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceMirrorFailedIterator{contract: _Marketplace.contract, event: "MirrorFailed", logs: logs, sub: sub}, nil
}

// WatchMirrorFailed is a free log subscription operation binding the contract event 0x42bf66fb325bb74dce540406916688333b68d163b61dc9fc22a6892fe73dda01.
//
// Solidity: event MirrorFailed(uint256 indexed id, address indexed owner, bytes failReason)
func (_Marketplace *MarketplaceFilterer) WatchMirrorFailed(opts *bind.WatchOpts, sink chan<- *MarketplaceMirrorFailed, id []*big.Int, owner []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "MirrorFailed", idRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceMirrorFailed)
				if err := _Marketplace.contract.UnpackLog(event, "MirrorFailed", log); err != nil {
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

// ParseMirrorFailed is a log parse operation binding the contract event 0x42bf66fb325bb74dce540406916688333b68d163b61dc9fc22a6892fe73dda01.
//
// Solidity: event MirrorFailed(uint256 indexed id, address indexed owner, bytes failReason)
func (_Marketplace *MarketplaceFilterer) ParseMirrorFailed(log types.Log) (*MarketplaceMirrorFailed, error) {
	event := new(MarketplaceMirrorFailed)
	if err := _Marketplace.contract.UnpackLog(event, "MirrorFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceMirrorSuccessIterator is returned from FilterMirrorSuccess and is used to iterate over the raw logs and unpacked data for MirrorSuccess events raised by the Marketplace contract.
type MarketplaceMirrorSuccessIterator struct {
	Event *MarketplaceMirrorSuccess // Event containing the contract specifics and raw log

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
func (it *MarketplaceMirrorSuccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceMirrorSuccess)
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
		it.Event = new(MarketplaceMirrorSuccess)
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
func (it *MarketplaceMirrorSuccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceMirrorSuccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceMirrorSuccess represents a MirrorSuccess event raised by the Marketplace contract.
type MarketplaceMirrorSuccess struct {
	Id    *big.Int
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterMirrorSuccess is a free log retrieval operation binding the contract event 0x1dddad5022f520826425a8233f0f024533427d10121900971e2521e4294bed46.
//
// Solidity: event MirrorSuccess(uint256 indexed id, address indexed owner)
func (_Marketplace *MarketplaceFilterer) FilterMirrorSuccess(opts *bind.FilterOpts, id []*big.Int, owner []common.Address) (*MarketplaceMirrorSuccessIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "MirrorSuccess", idRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceMirrorSuccessIterator{contract: _Marketplace.contract, event: "MirrorSuccess", logs: logs, sub: sub}, nil
}

// WatchMirrorSuccess is a free log subscription operation binding the contract event 0x1dddad5022f520826425a8233f0f024533427d10121900971e2521e4294bed46.
//
// Solidity: event MirrorSuccess(uint256 indexed id, address indexed owner)
func (_Marketplace *MarketplaceFilterer) WatchMirrorSuccess(opts *bind.WatchOpts, sink chan<- *MarketplaceMirrorSuccess, id []*big.Int, owner []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "MirrorSuccess", idRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceMirrorSuccess)
				if err := _Marketplace.contract.UnpackLog(event, "MirrorSuccess", log); err != nil {
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

// ParseMirrorSuccess is a log parse operation binding the contract event 0x1dddad5022f520826425a8233f0f024533427d10121900971e2521e4294bed46.
//
// Solidity: event MirrorSuccess(uint256 indexed id, address indexed owner)
func (_Marketplace *MarketplaceFilterer) ParseMirrorSuccess(log types.Log) (*MarketplaceMirrorSuccess, error) {
	event := new(MarketplaceMirrorSuccess)
	if err := _Marketplace.contract.UnpackLog(event, "MirrorSuccess", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

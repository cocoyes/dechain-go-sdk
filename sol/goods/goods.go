// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package goods

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

// GoodsMetaData contains all meta data concerning the Goods contract.
var GoodsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_master\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"amount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_otherGoods\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_myOrderId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_otherOrderId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_myNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_otherNum\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_myPay\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_payAmount\",\"type\":\"uint256\"}],\"name\":\"createOrder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_myOrderId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_orderStatus\",\"type\":\"uint256\"}],\"name\":\"doOrderStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBaseInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"own\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"p\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rem\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"am\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_myOrderId\",\"type\":\"string\"}],\"name\":\"getOrderInfo\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"_myId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_otherOrderId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_myNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_otherNum\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_myPay\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_payAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"own\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"orderMap\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"otherGoods\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"myOrderId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"otherOrderId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"myNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"otherNum\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"myPay\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"payAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"orderStatus\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"used\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"price\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"remainAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405260006003556000600455600060055534801561001f57600080fd5b506040516113d43803806113d483398101604081905261003e91610088565b600080546001600160a01b039586166001600160a01b03199182161790915560028054821633179055600180549490951693169290921790925560048290556003556005556100e7565b6000806000806080858703121561009d578384fd5b84516100a8816100cf565b60208601519094506100b9816100cf565b6040860151606090960151949790965092505050565b6001600160a01b03811681146100e457600080fd5b50565b6112de806100f66000396000f3fe6080604052600436106100915760003560e01c8063a035b1fe11610059578063a035b1fe1461014f578063aa528e7614610164578063aa8c217c14610184578063bc7f253d14610199578063e7b0272e146101ce57610091565b806315eb3f301461009657806325e7514b146100d1578063893d20e8146100f35780638da5cb5b146101155780639551ae441461012a575b600080fd5b3480156100a257600080fd5b506100b66100b1366004610eec565b6101ee565b6040516100c8969594939291906110a4565b60405180910390f35b3480156100dd57600080fd5b506100e6610370565b6040516100c891906111bb565b3480156100ff57600080fd5b50610108610376565b6040516100c89190610fb2565b34801561012157600080fd5b50610108610385565b34801561013657600080fd5b5061013f610394565b6040516100c89493929190611073565b34801561015b57600080fd5b506100e66103b2565b610177610172366004610f27565b6103b8565b6040516100c89190611099565b34801561019057600080fd5b506100e661083b565b3480156101a557600080fd5b506101b96101b4366004610eec565b610841565b6040516100c899989796959493929190610fea565b3480156101da57600080fd5b506101776101e9366004610e43565b6109b5565b60608060008060008060006006886040516102099190610f96565b908152602001604051809103902090508060010181600201826003015483600401548460050160009054906101000a900460ff16856006015485805461024e9061123f565b80601f016020809104026020016040519081016040528092919081815260200182805461027a9061123f565b80156102c75780601f1061029c576101008083540402835291602001916102c7565b820191906000526020600020905b8154815290600101906020018083116102aa57829003601f168201915b505050505095508480546102da9061123f565b80601f01602080910402602001604051908101604052809291908181526020018280546103069061123f565b80156103535780601f1061032857610100808354040283529160200191610353565b820191906000526020600020905b81548152906001019060200180831161033657829003601f168201915b505050505094509650965096509650965096505091939550919395565b60055481565b6002546001600160a01b031690565b6002546001600160a01b031681565b6002546003546005546004546001600160a01b039093169290919293565b60035481565b600080546001600160a01b031633146103ec5760405162461bcd60e51b81526004016103e39061116d565b60405180910390fd5b6006836040516103fc9190610f96565b9081526040519081900360200190206008015460ff1661042e5760405162461bcd60e51b81526004016103e390611192565b60026006846040516104409190610f96565b9081526020016040518091039020600701541061046f5760405162461bcd60e51b81526004016103e390611136565b816006846040516104809190610f96565b90815260200160405180910390206007018190555060006006846040516104a79190610f96565b908152604051908190036020019020546001600160a01b031690506002831480156104f457506006846040516104dd9190610f96565b9081526040519081900360200190206005015460ff165b80156105215750600060068560405161050d9190610f96565b908152602001604051809103902060060154115b156105b9576001546002546040516001600160a01b039283169263a9059cbb921690600690610551908990610f96565b908152604051908190036020018120600601546001600160e01b031960e085901b168252610582929160040161105a565b600060405180830381600087803b15801561059c57600080fd5b505af11580156105b0573d6000803e3d6000fd5b5050505061082f565b8260031480156105eb57506006846040516105d49190610f96565b9081526040519081900360200190206005015460ff165b8015610618575060006006856040516106049190610f96565b908152602001604051809103902060060154115b156106c257600160009054906101000a90046001600160a01b03166001600160a01b031663a9059cbb826001600160a01b031663893d20e86040518163ffffffff1660e01b815260040160206040518083038186803b15801561067a57600080fd5b505afa15801561068e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106b29190610e27565b6006876040516105519190610f96565b8260041480156106f457506006846040516106dd9190610f96565b9081526040519081900360200190206005015460ff165b80156107215750600060068560405161070d9190610f96565b908152602001604051809103902060060154115b1561082f57600160009054906101000a90046001600160a01b03166001600160a01b031663a9059cbb826001600160a01b031663893d20e86040518163ffffffff1660e01b815260040160206040518083038186803b15801561078357600080fd5b505afa158015610797573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107bb9190610e27565b6006876040516107cb9190610f96565b908152604051908190036020018120600601546001600160e01b031960e085901b1682526107fc929160040161105a565b600060405180830381600087803b15801561081657600080fd5b505af115801561082a573d6000803e3d6000fd5b505050505b60019150505b92915050565b60045481565b8051602081830181018051600682529282019190930120915280546001820180546001600160a01b0390921692916108789061123f565b80601f01602080910402602001604051908101604052809291908181526020018280546108a49061123f565b80156108f15780601f106108c6576101008083540402835291602001916108f1565b820191906000526020600020905b8154815290600101906020018083116108d457829003601f168201915b5050505050908060020180546109069061123f565b80601f01602080910402602001604051908101604052809291908181526020018280546109329061123f565b801561097f5780601f106109545761010080835404028352916020019161097f565b820191906000526020600020905b81548152906001019060200180831161096257829003601f168201915b505050600384015460048501546005860154600687015460078801546008909801549697939692955060ff918216945092911689565b6002546000906001600160a01b031633146109e25760405162461bcd60e51b81526004016103e39061116d565b6006876040516109f29190610f96565b9081526040519081900360200190206008015460ff1615610a255760405162461bcd60e51b81526004016103e390611111565b84600688604051610a369190610f96565b90815260200160405180910390206003018190555086600688604051610a5c9190610f96565b90815260200160405180910390206001019080519060200190610a80929190610d0d565b5082600688604051610a929190610f96565b908152604051908190036020018120600501805492151560ff19909316929092179091558890600690610ac6908a90610f96565b90815260405190819003602001812080546001600160a01b03939093166001600160a01b0319909316929092179091558690600690610b06908a90610f96565b90815260200160405180910390206002019080519060200190610b2a929190610d0d565b5083600688604051610b3c9190610f96565b90815260200160405180910390206004018190555081600688604051610b629190610f96565b9081526020016040518091039020600601819055506000600688604051610b899190610f96565b908152604051908190036020019020600701558215610c09576001546040516323b872dd60e01b81526001600160a01b03909116906323b872dd90610bd690339030908790600401610fc6565b600060405180830381600087803b158015610bf057600080fd5b505af1158015610c04573d6000803e3d6000fd5b505050505b6001600688604051610c1b9190610f96565b908152604051908190036020019020600801805491151560ff199092169190911790556000546003546001600160a01b0390911690819063cbcfd2a9908a90610c65908a90610cc3565b6040518363ffffffff1660e01b8152600401610c829291906110ef565b600060405180830381600087803b158015610c9c57600080fd5b505af1158015610cb0573d6000803e3d6000fd5b5060019c9b505050505050505050505050565b600082610cd257506000610835565b6000610cde83856111e4565b905082610ceb85836111c4565b14610d0657634e487b7160e01b600052600160045260246000fd5b9392505050565b828054610d199061123f565b90600052602060002090601f016020900481019282610d3b5760008555610d81565b82601f10610d5457805160ff1916838001178555610d81565b82800160010185558215610d81579182015b82811115610d81578251825591602001919060010190610d66565b50610d8d929150610d91565b5090565b5b80821115610d8d5760008155600101610d92565b600082601f830112610db6578081fd5b813567ffffffffffffffff80821115610dd157610dd161127a565b604051601f8301601f191681016020018281118282101715610df557610df561127a565b604052828152848301602001861015610e0c578384fd5b82602086016020830137918201602001929092529392505050565b600060208284031215610e38578081fd5b8151610d0681611290565b600080600080600080600060e0888a031215610e5d578283fd5b8735610e6881611290565b9650602088013567ffffffffffffffff80821115610e84578485fd5b610e908b838c01610da6565b975060408a0135915080821115610ea5578485fd5b50610eb28a828b01610da6565b955050606088013593506080880135925060a08801358015158114610ed5578283fd5b8092505060c0880135905092959891949750929550565b600060208284031215610efd578081fd5b813567ffffffffffffffff811115610f13578182fd5b610f1f84828501610da6565b949350505050565b60008060408385031215610f39578182fd5b823567ffffffffffffffff811115610f4f578283fd5b610f5b85828601610da6565b95602094909401359450505050565b60008151808452610f8281602086016020860161120f565b601f01601f19169290920160200192915050565b60008251610fa881846020870161120f565b9190910192915050565b6001600160a01b0391909116815260200190565b6001600160a01b039384168152919092166020820152604081019190915260600190565b6001600160a01b038a1681526101206020820181905260009061100f8382018c610f6a565b90508281036040840152611023818b610f6a565b60608401999099525050608081019590955292151560a085015260c084019190915260e08301521515610100909101529392505050565b6001600160a01b03929092168252602082015260400190565b6001600160a01b0394909416845260208401929092526040830152606082015260800190565b901515815260200190565b600060c082526110b760c0830189610f6a565b82810360208401526110c98189610f6a565b604084019790975250506060810193909352901515608083015260a09091015292915050565b6000604082526111026040830185610f6a565b90508260208301529392505050565b6020808252600b908201526a1bdc99195c88195e1a5cdd60aa1b604082015260600190565b60208082526017908201527f6f72646572207374617475732063616e206e6f74206f70000000000000000000604082015260600190565b6020808252600b908201526a0796f752063616e74206f760ac1b604082015260600190565b6020808252600f908201526e1bdc99195c881b9bdd08195e1a5cdd608a1b604082015260600190565b90815260200190565b6000826111df57634e487b7160e01b81526012600452602481fd5b500490565b600081600019048311821515161561120a57634e487b7160e01b81526011600452602481fd5b500290565b60005b8381101561122a578181015183820152602001611212565b83811115611239576000848401525b50505050565b60028104600182168061125357607f821691505b6020821081141561127457634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052604160045260246000fd5b6001600160a01b03811681146112a557600080fd5b5056fea26469706673582212200bcbf97f5d5bff63f3ab87d83e3f4e03fe1127d801b694f6014939186b9d861e64736f6c63430008000033",
}

// GoodsABI is the input ABI used to generate the binding from.
// Deprecated: Use GoodsMetaData.ABI instead.
var GoodsABI = GoodsMetaData.ABI

// GoodsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GoodsMetaData.Bin instead.
var GoodsBin = GoodsMetaData.Bin

// DeployGoods deploys a new Ethereum contract, binding an instance of Goods to it.
func DeployGoods(auth *bind.TransactOpts, backend bind.ContractBackend, _master common.Address, _token common.Address, _amount *big.Int, _price *big.Int) (common.Address, *types.Transaction, *Goods, error) {
	parsed, err := GoodsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GoodsBin), backend, _master, _token, _amount, _price)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Goods{GoodsCaller: GoodsCaller{contract: contract}, GoodsTransactor: GoodsTransactor{contract: contract}, GoodsFilterer: GoodsFilterer{contract: contract}}, nil
}

// Goods is an auto generated Go binding around an Ethereum contract.
type Goods struct {
	GoodsCaller     // Read-only binding to the contract
	GoodsTransactor // Write-only binding to the contract
	GoodsFilterer   // Log filterer for contract events
}

// GoodsCaller is an auto generated read-only Go binding around an Ethereum contract.
type GoodsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GoodsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GoodsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GoodsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GoodsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GoodsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GoodsSession struct {
	Contract     *Goods            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GoodsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GoodsCallerSession struct {
	Contract *GoodsCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// GoodsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GoodsTransactorSession struct {
	Contract     *GoodsTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GoodsRaw is an auto generated low-level Go binding around an Ethereum contract.
type GoodsRaw struct {
	Contract *Goods // Generic contract binding to access the raw methods on
}

// GoodsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GoodsCallerRaw struct {
	Contract *GoodsCaller // Generic read-only contract binding to access the raw methods on
}

// GoodsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GoodsTransactorRaw struct {
	Contract *GoodsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGoods creates a new instance of Goods, bound to a specific deployed contract.
func NewGoods(address common.Address, backend bind.ContractBackend) (*Goods, error) {
	contract, err := bindGoods(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Goods{GoodsCaller: GoodsCaller{contract: contract}, GoodsTransactor: GoodsTransactor{contract: contract}, GoodsFilterer: GoodsFilterer{contract: contract}}, nil
}

// NewGoodsCaller creates a new read-only instance of Goods, bound to a specific deployed contract.
func NewGoodsCaller(address common.Address, caller bind.ContractCaller) (*GoodsCaller, error) {
	contract, err := bindGoods(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GoodsCaller{contract: contract}, nil
}

// NewGoodsTransactor creates a new write-only instance of Goods, bound to a specific deployed contract.
func NewGoodsTransactor(address common.Address, transactor bind.ContractTransactor) (*GoodsTransactor, error) {
	contract, err := bindGoods(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GoodsTransactor{contract: contract}, nil
}

// NewGoodsFilterer creates a new log filterer instance of Goods, bound to a specific deployed contract.
func NewGoodsFilterer(address common.Address, filterer bind.ContractFilterer) (*GoodsFilterer, error) {
	contract, err := bindGoods(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GoodsFilterer{contract: contract}, nil
}

// bindGoods binds a generic wrapper to an already deployed contract.
func bindGoods(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GoodsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Goods *GoodsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Goods.Contract.GoodsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Goods *GoodsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Goods.Contract.GoodsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Goods *GoodsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Goods.Contract.GoodsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Goods *GoodsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Goods.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Goods *GoodsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Goods.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Goods *GoodsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Goods.Contract.contract.Transact(opts, method, params...)
}

// Amount is a free data retrieval call binding the contract method 0xaa8c217c.
//
// Solidity: function amount() view returns(uint256)
func (_Goods *GoodsCaller) Amount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Goods.contract.Call(opts, &out, "amount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Amount is a free data retrieval call binding the contract method 0xaa8c217c.
//
// Solidity: function amount() view returns(uint256)
func (_Goods *GoodsSession) Amount() (*big.Int, error) {
	return _Goods.Contract.Amount(&_Goods.CallOpts)
}

// Amount is a free data retrieval call binding the contract method 0xaa8c217c.
//
// Solidity: function amount() view returns(uint256)
func (_Goods *GoodsCallerSession) Amount() (*big.Int, error) {
	return _Goods.Contract.Amount(&_Goods.CallOpts)
}

// GetBaseInfo is a free data retrieval call binding the contract method 0x9551ae44.
//
// Solidity: function getBaseInfo() view returns(address own, uint256 p, uint256 rem, uint256 am)
func (_Goods *GoodsCaller) GetBaseInfo(opts *bind.CallOpts) (struct {
	Own common.Address
	P   *big.Int
	Rem *big.Int
	Am  *big.Int
}, error) {
	var out []interface{}
	err := _Goods.contract.Call(opts, &out, "getBaseInfo")

	outstruct := new(struct {
		Own common.Address
		P   *big.Int
		Rem *big.Int
		Am  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Own = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.P = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Rem = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Am = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetBaseInfo is a free data retrieval call binding the contract method 0x9551ae44.
//
// Solidity: function getBaseInfo() view returns(address own, uint256 p, uint256 rem, uint256 am)
func (_Goods *GoodsSession) GetBaseInfo() (struct {
	Own common.Address
	P   *big.Int
	Rem *big.Int
	Am  *big.Int
}, error) {
	return _Goods.Contract.GetBaseInfo(&_Goods.CallOpts)
}

// GetBaseInfo is a free data retrieval call binding the contract method 0x9551ae44.
//
// Solidity: function getBaseInfo() view returns(address own, uint256 p, uint256 rem, uint256 am)
func (_Goods *GoodsCallerSession) GetBaseInfo() (struct {
	Own common.Address
	P   *big.Int
	Rem *big.Int
	Am  *big.Int
}, error) {
	return _Goods.Contract.GetBaseInfo(&_Goods.CallOpts)
}

// GetOrderInfo is a free data retrieval call binding the contract method 0x15eb3f30.
//
// Solidity: function getOrderInfo(string _myOrderId) view returns(string _myId, string _otherOrderId, uint256 _myNum, uint256 _otherNum, bool _myPay, uint256 _payAmount)
func (_Goods *GoodsCaller) GetOrderInfo(opts *bind.CallOpts, _myOrderId string) (struct {
	MyId         string
	OtherOrderId string
	MyNum        *big.Int
	OtherNum     *big.Int
	MyPay        bool
	PayAmount    *big.Int
}, error) {
	var out []interface{}
	err := _Goods.contract.Call(opts, &out, "getOrderInfo", _myOrderId)

	outstruct := new(struct {
		MyId         string
		OtherOrderId string
		MyNum        *big.Int
		OtherNum     *big.Int
		MyPay        bool
		PayAmount    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MyId = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.OtherOrderId = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.MyNum = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.OtherNum = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.MyPay = *abi.ConvertType(out[4], new(bool)).(*bool)
	outstruct.PayAmount = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetOrderInfo is a free data retrieval call binding the contract method 0x15eb3f30.
//
// Solidity: function getOrderInfo(string _myOrderId) view returns(string _myId, string _otherOrderId, uint256 _myNum, uint256 _otherNum, bool _myPay, uint256 _payAmount)
func (_Goods *GoodsSession) GetOrderInfo(_myOrderId string) (struct {
	MyId         string
	OtherOrderId string
	MyNum        *big.Int
	OtherNum     *big.Int
	MyPay        bool
	PayAmount    *big.Int
}, error) {
	return _Goods.Contract.GetOrderInfo(&_Goods.CallOpts, _myOrderId)
}

// GetOrderInfo is a free data retrieval call binding the contract method 0x15eb3f30.
//
// Solidity: function getOrderInfo(string _myOrderId) view returns(string _myId, string _otherOrderId, uint256 _myNum, uint256 _otherNum, bool _myPay, uint256 _payAmount)
func (_Goods *GoodsCallerSession) GetOrderInfo(_myOrderId string) (struct {
	MyId         string
	OtherOrderId string
	MyNum        *big.Int
	OtherNum     *big.Int
	MyPay        bool
	PayAmount    *big.Int
}, error) {
	return _Goods.Contract.GetOrderInfo(&_Goods.CallOpts, _myOrderId)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address own)
func (_Goods *GoodsCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Goods.contract.Call(opts, &out, "getOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address own)
func (_Goods *GoodsSession) GetOwner() (common.Address, error) {
	return _Goods.Contract.GetOwner(&_Goods.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address own)
func (_Goods *GoodsCallerSession) GetOwner() (common.Address, error) {
	return _Goods.Contract.GetOwner(&_Goods.CallOpts)
}

// OrderMap is a free data retrieval call binding the contract method 0xbc7f253d.
//
// Solidity: function orderMap(string ) view returns(address otherGoods, string myOrderId, string otherOrderId, uint256 myNum, uint256 otherNum, bool myPay, uint256 payAmount, uint256 orderStatus, bool used)
func (_Goods *GoodsCaller) OrderMap(opts *bind.CallOpts, arg0 string) (struct {
	OtherGoods   common.Address
	MyOrderId    string
	OtherOrderId string
	MyNum        *big.Int
	OtherNum     *big.Int
	MyPay        bool
	PayAmount    *big.Int
	OrderStatus  *big.Int
	Used         bool
}, error) {
	var out []interface{}
	err := _Goods.contract.Call(opts, &out, "orderMap", arg0)

	outstruct := new(struct {
		OtherGoods   common.Address
		MyOrderId    string
		OtherOrderId string
		MyNum        *big.Int
		OtherNum     *big.Int
		MyPay        bool
		PayAmount    *big.Int
		OrderStatus  *big.Int
		Used         bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.OtherGoods = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.MyOrderId = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.OtherOrderId = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.MyNum = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.OtherNum = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.MyPay = *abi.ConvertType(out[5], new(bool)).(*bool)
	outstruct.PayAmount = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.OrderStatus = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.Used = *abi.ConvertType(out[8], new(bool)).(*bool)

	return *outstruct, err

}

// OrderMap is a free data retrieval call binding the contract method 0xbc7f253d.
//
// Solidity: function orderMap(string ) view returns(address otherGoods, string myOrderId, string otherOrderId, uint256 myNum, uint256 otherNum, bool myPay, uint256 payAmount, uint256 orderStatus, bool used)
func (_Goods *GoodsSession) OrderMap(arg0 string) (struct {
	OtherGoods   common.Address
	MyOrderId    string
	OtherOrderId string
	MyNum        *big.Int
	OtherNum     *big.Int
	MyPay        bool
	PayAmount    *big.Int
	OrderStatus  *big.Int
	Used         bool
}, error) {
	return _Goods.Contract.OrderMap(&_Goods.CallOpts, arg0)
}

// OrderMap is a free data retrieval call binding the contract method 0xbc7f253d.
//
// Solidity: function orderMap(string ) view returns(address otherGoods, string myOrderId, string otherOrderId, uint256 myNum, uint256 otherNum, bool myPay, uint256 payAmount, uint256 orderStatus, bool used)
func (_Goods *GoodsCallerSession) OrderMap(arg0 string) (struct {
	OtherGoods   common.Address
	MyOrderId    string
	OtherOrderId string
	MyNum        *big.Int
	OtherNum     *big.Int
	MyPay        bool
	PayAmount    *big.Int
	OrderStatus  *big.Int
	Used         bool
}, error) {
	return _Goods.Contract.OrderMap(&_Goods.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Goods *GoodsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Goods.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Goods *GoodsSession) Owner() (common.Address, error) {
	return _Goods.Contract.Owner(&_Goods.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Goods *GoodsCallerSession) Owner() (common.Address, error) {
	return _Goods.Contract.Owner(&_Goods.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() view returns(uint256)
func (_Goods *GoodsCaller) Price(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Goods.contract.Call(opts, &out, "price")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() view returns(uint256)
func (_Goods *GoodsSession) Price() (*big.Int, error) {
	return _Goods.Contract.Price(&_Goods.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() view returns(uint256)
func (_Goods *GoodsCallerSession) Price() (*big.Int, error) {
	return _Goods.Contract.Price(&_Goods.CallOpts)
}

// RemainAmount is a free data retrieval call binding the contract method 0x25e7514b.
//
// Solidity: function remainAmount() view returns(uint256)
func (_Goods *GoodsCaller) RemainAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Goods.contract.Call(opts, &out, "remainAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RemainAmount is a free data retrieval call binding the contract method 0x25e7514b.
//
// Solidity: function remainAmount() view returns(uint256)
func (_Goods *GoodsSession) RemainAmount() (*big.Int, error) {
	return _Goods.Contract.RemainAmount(&_Goods.CallOpts)
}

// RemainAmount is a free data retrieval call binding the contract method 0x25e7514b.
//
// Solidity: function remainAmount() view returns(uint256)
func (_Goods *GoodsCallerSession) RemainAmount() (*big.Int, error) {
	return _Goods.Contract.RemainAmount(&_Goods.CallOpts)
}

// CreateOrder is a paid mutator transaction binding the contract method 0xe7b0272e.
//
// Solidity: function createOrder(address _otherGoods, string _myOrderId, string _otherOrderId, uint256 _myNum, uint256 _otherNum, bool _myPay, uint256 _payAmount) returns(bool)
func (_Goods *GoodsTransactor) CreateOrder(opts *bind.TransactOpts, _otherGoods common.Address, _myOrderId string, _otherOrderId string, _myNum *big.Int, _otherNum *big.Int, _myPay bool, _payAmount *big.Int) (*types.Transaction, error) {
	return _Goods.contract.Transact(opts, "createOrder", _otherGoods, _myOrderId, _otherOrderId, _myNum, _otherNum, _myPay, _payAmount)
}

// CreateOrder is a paid mutator transaction binding the contract method 0xe7b0272e.
//
// Solidity: function createOrder(address _otherGoods, string _myOrderId, string _otherOrderId, uint256 _myNum, uint256 _otherNum, bool _myPay, uint256 _payAmount) returns(bool)
func (_Goods *GoodsSession) CreateOrder(_otherGoods common.Address, _myOrderId string, _otherOrderId string, _myNum *big.Int, _otherNum *big.Int, _myPay bool, _payAmount *big.Int) (*types.Transaction, error) {
	return _Goods.Contract.CreateOrder(&_Goods.TransactOpts, _otherGoods, _myOrderId, _otherOrderId, _myNum, _otherNum, _myPay, _payAmount)
}

// CreateOrder is a paid mutator transaction binding the contract method 0xe7b0272e.
//
// Solidity: function createOrder(address _otherGoods, string _myOrderId, string _otherOrderId, uint256 _myNum, uint256 _otherNum, bool _myPay, uint256 _payAmount) returns(bool)
func (_Goods *GoodsTransactorSession) CreateOrder(_otherGoods common.Address, _myOrderId string, _otherOrderId string, _myNum *big.Int, _otherNum *big.Int, _myPay bool, _payAmount *big.Int) (*types.Transaction, error) {
	return _Goods.Contract.CreateOrder(&_Goods.TransactOpts, _otherGoods, _myOrderId, _otherOrderId, _myNum, _otherNum, _myPay, _payAmount)
}

// DoOrderStatus is a paid mutator transaction binding the contract method 0xaa528e76.
//
// Solidity: function doOrderStatus(string _myOrderId, uint256 _orderStatus) payable returns(bool)
func (_Goods *GoodsTransactor) DoOrderStatus(opts *bind.TransactOpts, _myOrderId string, _orderStatus *big.Int) (*types.Transaction, error) {
	return _Goods.contract.Transact(opts, "doOrderStatus", _myOrderId, _orderStatus)
}

// DoOrderStatus is a paid mutator transaction binding the contract method 0xaa528e76.
//
// Solidity: function doOrderStatus(string _myOrderId, uint256 _orderStatus) payable returns(bool)
func (_Goods *GoodsSession) DoOrderStatus(_myOrderId string, _orderStatus *big.Int) (*types.Transaction, error) {
	return _Goods.Contract.DoOrderStatus(&_Goods.TransactOpts, _myOrderId, _orderStatus)
}

// DoOrderStatus is a paid mutator transaction binding the contract method 0xaa528e76.
//
// Solidity: function doOrderStatus(string _myOrderId, uint256 _orderStatus) payable returns(bool)
func (_Goods *GoodsTransactorSession) DoOrderStatus(_myOrderId string, _orderStatus *big.Int) (*types.Transaction, error) {
	return _Goods.Contract.DoOrderStatus(&_Goods.TransactOpts, _myOrderId, _orderStatus)
}

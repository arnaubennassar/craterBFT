// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package cratercontract

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

// CraterComet is an auto generated low-level Go binding around an user-defined struct.
type CraterComet struct {
	Message       []byte
	Nonce         *big.Int
	BlockDeadline *big.Int
}

// CratercontractMetaData contains all meta data concerning the Cratercontract contract.
var CratercontractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"BlockDeadlineExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CommitteeAddressDoesntExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyURLNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooManyRequiredSignatures\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnexpectedAddrsAndSignaturesSize\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnexpectedAddrsBytesLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnexpectedCommitteeHash\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WrongAddrOrder\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"committeeHash\",\"type\":\"bytes32\"}],\"name\":\"CommitteeUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"committeeHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"craters\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAmountOfMembers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"members\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requiredAmountOfSignatures\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockDeadline\",\"type\":\"uint256\"}],\"internalType\":\"structCrater.Comet\",\"name\":\"comet\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signaturesAndAddrs\",\"type\":\"bytes\"}],\"name\":\"sendComet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requiredAmountOfSignatures\",\"type\":\"uint256\"},{\"internalType\":\"string[]\",\"name\":\"urls\",\"type\":\"string[]\"},{\"internalType\":\"bytes\",\"name\":\"addrsBytes\",\"type\":\"bytes\"}],\"name\":\"setupCommittee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"signedHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signaturesAndAddrs\",\"type\":\"bytes\"}],\"name\":\"verifySignatures\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061159b806100206000396000f3fe608060405234801561001057600080fd5b50600436106100a35760003560e01c8063609d454411610076578063affed0e01161005b578063affed0e01461013a578063c7a823e014610143578063dce1e2b61461015657600080fd5b8063609d45441461011a5780636beedd391461013157600080fd5b8063078fba2a146100a857806337d62dd9146100bd578063461eadcb146100d05780635daf08ca146100f9575b600080fd5b6100bb6100b6366004610da1565b61015e565b005b6100bb6100cb366004610e4c565b61045a565b6100e36100de366004610ebd565b6105c9565b6040516100f09190610f3a565b60405180910390f35b61010c610107366004610ebd565b610675565b6040516100f0929190610f54565b61012360015481565b6040519081526020016100f0565b61012360005481565b61012360035481565b6100bb610151366004610f8c565b610747565b600254610123565b8285811015610199576040517f2e7dcd6e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6101a4601482610ffa565b82146101dc576040517f2ab6a12900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6101e860026000610cad565b6000805b828110156103fe576000610201601483610ffa565b90506000868287610213601483611011565b9261022093929190611024565b6102299161104e565b60601c905088888481811061024057610240611096565b905060200281019061025291906110c5565b905060000361028d576040517fb54b70e400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff16106102f2576040517fd53cfbe000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b809350600260405180604001604052808b8b8781811061031457610314611096565b905060200281019061032691906110c5565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092018290525093855250505073ffffffffffffffffffffffffffffffffffffffff8516602092830152835460018101855593815220815191926002020190819061039a90826111fb565b5060209190910151600190910180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff909216919091179055508190506103f681611315565b9150506101ec565b50838360405161040f92919061134d565b6040519081900381206001819055600089905581527f831403fd381b3e6ac875d912ec2eee0e0203d0d29f7b3e0c96fc8f582d6db6579060200160405180910390a150505050505050565b600354610468906001611011565b8360200135146104a4576040517f756688fe00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b43836040013511156104e2576040517ff68e590d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006104ee84806110c5565b6040516104fc92919061134d565b6040519081900381207fc7a823e00000000000000000000000000000000000000000000000000000000082529150309063c7a823e0906105449084908790879060040161135d565b60006040518083038186803b15801561055c57600080fd5b505afa158015610570573d6000803e3d6000fd5b5061058192508691508190506110c5565b60046003548154811061059657610596611096565b9060005260206000200191826105ad9291906113b1565b50600380549060006105be83611315565b919050555050505050565b600481815481106105d957600080fd5b9060005260206000200160009150905080546105f490611159565b80601f016020809104026020016040519081016040528092919081815260200182805461062090611159565b801561066d5780601f106106425761010080835404028352916020019161066d565b820191906000526020600020905b81548152906001019060200180831161065057829003601f168201915b505050505081565b6002818154811061068557600080fd5b90600052602060002090600202016000915090508060000180546106a890611159565b80601f01602080910402602001604051908101604052809291908181526020018280546106d490611159565b80156107215780601f106106f657610100808354040283529160200191610721565b820191906000526020600020905b81548152906001019060200180831161070457829003601f168201915b5050506001909301549192505073ffffffffffffffffffffffffffffffffffffffff1682565b60008054610756906041610ffa565b90508082108061077a5750601461076d82846114cc565b610777919061150e565b15155b156107b1576040517f6b8eec4600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001546107c083838187611024565b6040516107ce92919061134d565b60405180910390201461080d576040517f6b156b2800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600080601461081c84866114cc565b6108269190611522565b905060005b60005481101561098f5760006108a6888888610848604187610ffa565b9060416108558189610ffa565b61085f9190611011565b9261086c93929190611024565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061099892505050565b90506000845b848110156109425760006108c1601483610ffa565b6108cb9089611011565b905060008a828b6108dd601483611011565b926108ea93929190611024565b6108f39161104e565b60601c905073ffffffffffffffffffffffffffffffffffffffff8516810361092d57610920836001611011565b9750600193505050610942565b5050808061093a90611315565b9150506108ac565b508061097a576040517f8431721300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5050808061098790611315565b91505061082b565b50505050505050565b60008060006109a785856109be565b915091506109b481610a03565b5090505b92915050565b60008082516041036109f45760208301516040840151606085015160001a6109e887828585610bbe565b945094505050506109fc565b506000905060025b9250929050565b6000816004811115610a1757610a17611536565b03610a1f5750565b6001816004811115610a3357610a33611536565b03610a9f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f45434453413a20696e76616c6964207369676e6174757265000000000000000060448201526064015b60405180910390fd5b6002816004811115610ab357610ab3611536565b03610b1a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e677468006044820152606401610a96565b6003816004811115610b2e57610b2e611536565b03610bbb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c60448201527f75650000000000000000000000000000000000000000000000000000000000006064820152608401610a96565b50565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0831115610bf55750600090506003610ca4565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015610c49573d6000803e3d6000fd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0015191505073ffffffffffffffffffffffffffffffffffffffff8116610c9d57600060019250925050610ca4565b9150600090505b94509492505050565b5080546000825560020290600052602060002090810190610bbb91905b80821115610d11576000610cde8282610d15565b506001810180547fffffffffffffffffffffffff0000000000000000000000000000000000000000169055600201610cca565b5090565b508054610d2190611159565b6000825580601f10610d31575050565b601f016020900490600052602060002090810190610bbb91905b80821115610d115760008155600101610d4b565b60008083601f840112610d7157600080fd5b50813567ffffffffffffffff811115610d8957600080fd5b6020830191508360208285010111156109fc57600080fd5b600080600080600060608688031215610db957600080fd5b85359450602086013567ffffffffffffffff80821115610dd857600080fd5b818801915088601f830112610dec57600080fd5b813581811115610dfb57600080fd5b8960208260051b8501011115610e1057600080fd5b602083019650809550506040880135915080821115610e2e57600080fd5b50610e3b88828901610d5f565b969995985093965092949392505050565b600080600060408486031215610e6157600080fd5b833567ffffffffffffffff80821115610e7957600080fd5b9085019060608288031215610e8d57600080fd5b90935060208501359080821115610ea357600080fd5b50610eb086828701610d5f565b9497909650939450505050565b600060208284031215610ecf57600080fd5b5035919050565b6000815180845260005b81811015610efc57602081850181015186830182015201610ee0565b5060006020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b602081526000610f4d6020830184610ed6565b9392505050565b604081526000610f676040830185610ed6565b905073ffffffffffffffffffffffffffffffffffffffff831660208301529392505050565b600080600060408486031215610fa157600080fd5b83359250602084013567ffffffffffffffff811115610fbf57600080fd5b610eb086828701610d5f565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b80820281158282048414176109b8576109b8610fcb565b808201808211156109b8576109b8610fcb565b6000808585111561103457600080fd5b8386111561104157600080fd5b5050820193919092039150565b7fffffffffffffffffffffffffffffffffffffffff000000000000000000000000813581811691601485101561108e5780818660140360031b1b83161692505b505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18436030181126110fa57600080fd5b83018035915067ffffffffffffffff82111561111557600080fd5b6020019150368190038213156109fc57600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600181811c9082168061116d57607f821691505b6020821081036111a6577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b601f8211156111f657600081815260208120601f850160051c810160208610156111d35750805b601f850160051c820191505b818110156111f2578281556001016111df565b5050505b505050565b815167ffffffffffffffff8111156112155761121561112a565b611229816112238454611159565b846111ac565b602080601f83116001811461127c57600084156112465750858301515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600386901b1c1916600185901b1785556111f2565b6000858152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08616915b828110156112c9578886015182559484019460019091019084016112aa565b508582101561130557878501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600388901b60f8161c191681555b5050505050600190811b01905550565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361134657611346610fcb565b5060010190565b8183823760009101908152919050565b83815260406020820152816040820152818360608301376000818301606090810191909152601f9092017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016010192915050565b67ffffffffffffffff8311156113c9576113c961112a565b6113dd836113d78354611159565b836111ac565b6000601f84116001811461142f57600085156113f95750838201355b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600387901b1c1916600186901b1783556114c5565b6000838152602090207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0861690835b8281101561147e578685013582556020948501946001909201910161145e565b50868210156114b9577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88860031b161c19848701351681555b505060018560011b0183555b5050505050565b818103818111156109b8576109b8610fcb565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b60008261151d5761151d6114df565b500690565b600082611531576115316114df565b500490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fdfea26469706673582212201448d3a33ef8523770040d590f1622a4fee24a09e1703f73b9c92350699379c764736f6c63430008140033",
}

// CratercontractABI is the input ABI used to generate the binding from.
// Deprecated: Use CratercontractMetaData.ABI instead.
var CratercontractABI = CratercontractMetaData.ABI

// CratercontractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CratercontractMetaData.Bin instead.
var CratercontractBin = CratercontractMetaData.Bin

// DeployCratercontract deploys a new Ethereum contract, binding an instance of Cratercontract to it.
func DeployCratercontract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Cratercontract, error) {
	parsed, err := CratercontractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CratercontractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Cratercontract{CratercontractCaller: CratercontractCaller{contract: contract}, CratercontractTransactor: CratercontractTransactor{contract: contract}, CratercontractFilterer: CratercontractFilterer{contract: contract}}, nil
}

// Cratercontract is an auto generated Go binding around an Ethereum contract.
type Cratercontract struct {
	CratercontractCaller     // Read-only binding to the contract
	CratercontractTransactor // Write-only binding to the contract
	CratercontractFilterer   // Log filterer for contract events
}

// CratercontractCaller is an auto generated read-only Go binding around an Ethereum contract.
type CratercontractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CratercontractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CratercontractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CratercontractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CratercontractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CratercontractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CratercontractSession struct {
	Contract     *Cratercontract   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CratercontractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CratercontractCallerSession struct {
	Contract *CratercontractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// CratercontractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CratercontractTransactorSession struct {
	Contract     *CratercontractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// CratercontractRaw is an auto generated low-level Go binding around an Ethereum contract.
type CratercontractRaw struct {
	Contract *Cratercontract // Generic contract binding to access the raw methods on
}

// CratercontractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CratercontractCallerRaw struct {
	Contract *CratercontractCaller // Generic read-only contract binding to access the raw methods on
}

// CratercontractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CratercontractTransactorRaw struct {
	Contract *CratercontractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCratercontract creates a new instance of Cratercontract, bound to a specific deployed contract.
func NewCratercontract(address common.Address, backend bind.ContractBackend) (*Cratercontract, error) {
	contract, err := bindCratercontract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Cratercontract{CratercontractCaller: CratercontractCaller{contract: contract}, CratercontractTransactor: CratercontractTransactor{contract: contract}, CratercontractFilterer: CratercontractFilterer{contract: contract}}, nil
}

// NewCratercontractCaller creates a new read-only instance of Cratercontract, bound to a specific deployed contract.
func NewCratercontractCaller(address common.Address, caller bind.ContractCaller) (*CratercontractCaller, error) {
	contract, err := bindCratercontract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CratercontractCaller{contract: contract}, nil
}

// NewCratercontractTransactor creates a new write-only instance of Cratercontract, bound to a specific deployed contract.
func NewCratercontractTransactor(address common.Address, transactor bind.ContractTransactor) (*CratercontractTransactor, error) {
	contract, err := bindCratercontract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CratercontractTransactor{contract: contract}, nil
}

// NewCratercontractFilterer creates a new log filterer instance of Cratercontract, bound to a specific deployed contract.
func NewCratercontractFilterer(address common.Address, filterer bind.ContractFilterer) (*CratercontractFilterer, error) {
	contract, err := bindCratercontract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CratercontractFilterer{contract: contract}, nil
}

// bindCratercontract binds a generic wrapper to an already deployed contract.
func bindCratercontract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CratercontractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Cratercontract *CratercontractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Cratercontract.Contract.CratercontractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Cratercontract *CratercontractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cratercontract.Contract.CratercontractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Cratercontract *CratercontractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Cratercontract.Contract.CratercontractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Cratercontract *CratercontractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Cratercontract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Cratercontract *CratercontractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cratercontract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Cratercontract *CratercontractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Cratercontract.Contract.contract.Transact(opts, method, params...)
}

// CommitteeHash is a free data retrieval call binding the contract method 0x609d4544.
//
// Solidity: function committeeHash() view returns(bytes32)
func (_Cratercontract *CratercontractCaller) CommitteeHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Cratercontract.contract.Call(opts, &out, "committeeHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CommitteeHash is a free data retrieval call binding the contract method 0x609d4544.
//
// Solidity: function committeeHash() view returns(bytes32)
func (_Cratercontract *CratercontractSession) CommitteeHash() ([32]byte, error) {
	return _Cratercontract.Contract.CommitteeHash(&_Cratercontract.CallOpts)
}

// CommitteeHash is a free data retrieval call binding the contract method 0x609d4544.
//
// Solidity: function committeeHash() view returns(bytes32)
func (_Cratercontract *CratercontractCallerSession) CommitteeHash() ([32]byte, error) {
	return _Cratercontract.Contract.CommitteeHash(&_Cratercontract.CallOpts)
}

// Craters is a free data retrieval call binding the contract method 0x461eadcb.
//
// Solidity: function craters(uint256 ) view returns(bytes)
func (_Cratercontract *CratercontractCaller) Craters(opts *bind.CallOpts, arg0 *big.Int) ([]byte, error) {
	var out []interface{}
	err := _Cratercontract.contract.Call(opts, &out, "craters", arg0)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// Craters is a free data retrieval call binding the contract method 0x461eadcb.
//
// Solidity: function craters(uint256 ) view returns(bytes)
func (_Cratercontract *CratercontractSession) Craters(arg0 *big.Int) ([]byte, error) {
	return _Cratercontract.Contract.Craters(&_Cratercontract.CallOpts, arg0)
}

// Craters is a free data retrieval call binding the contract method 0x461eadcb.
//
// Solidity: function craters(uint256 ) view returns(bytes)
func (_Cratercontract *CratercontractCallerSession) Craters(arg0 *big.Int) ([]byte, error) {
	return _Cratercontract.Contract.Craters(&_Cratercontract.CallOpts, arg0)
}

// GetAmountOfMembers is a free data retrieval call binding the contract method 0xdce1e2b6.
//
// Solidity: function getAmountOfMembers() view returns(uint256)
func (_Cratercontract *CratercontractCaller) GetAmountOfMembers(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cratercontract.contract.Call(opts, &out, "getAmountOfMembers")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAmountOfMembers is a free data retrieval call binding the contract method 0xdce1e2b6.
//
// Solidity: function getAmountOfMembers() view returns(uint256)
func (_Cratercontract *CratercontractSession) GetAmountOfMembers() (*big.Int, error) {
	return _Cratercontract.Contract.GetAmountOfMembers(&_Cratercontract.CallOpts)
}

// GetAmountOfMembers is a free data retrieval call binding the contract method 0xdce1e2b6.
//
// Solidity: function getAmountOfMembers() view returns(uint256)
func (_Cratercontract *CratercontractCallerSession) GetAmountOfMembers() (*big.Int, error) {
	return _Cratercontract.Contract.GetAmountOfMembers(&_Cratercontract.CallOpts)
}

// Members is a free data retrieval call binding the contract method 0x5daf08ca.
//
// Solidity: function members(uint256 ) view returns(string url, address addr)
func (_Cratercontract *CratercontractCaller) Members(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Url  string
	Addr common.Address
}, error) {
	var out []interface{}
	err := _Cratercontract.contract.Call(opts, &out, "members", arg0)

	outstruct := new(struct {
		Url  string
		Addr common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Url = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Addr = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Members is a free data retrieval call binding the contract method 0x5daf08ca.
//
// Solidity: function members(uint256 ) view returns(string url, address addr)
func (_Cratercontract *CratercontractSession) Members(arg0 *big.Int) (struct {
	Url  string
	Addr common.Address
}, error) {
	return _Cratercontract.Contract.Members(&_Cratercontract.CallOpts, arg0)
}

// Members is a free data retrieval call binding the contract method 0x5daf08ca.
//
// Solidity: function members(uint256 ) view returns(string url, address addr)
func (_Cratercontract *CratercontractCallerSession) Members(arg0 *big.Int) (struct {
	Url  string
	Addr common.Address
}, error) {
	return _Cratercontract.Contract.Members(&_Cratercontract.CallOpts, arg0)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_Cratercontract *CratercontractCaller) Nonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cratercontract.contract.Call(opts, &out, "nonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_Cratercontract *CratercontractSession) Nonce() (*big.Int, error) {
	return _Cratercontract.Contract.Nonce(&_Cratercontract.CallOpts)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_Cratercontract *CratercontractCallerSession) Nonce() (*big.Int, error) {
	return _Cratercontract.Contract.Nonce(&_Cratercontract.CallOpts)
}

// RequiredAmountOfSignatures is a free data retrieval call binding the contract method 0x6beedd39.
//
// Solidity: function requiredAmountOfSignatures() view returns(uint256)
func (_Cratercontract *CratercontractCaller) RequiredAmountOfSignatures(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cratercontract.contract.Call(opts, &out, "requiredAmountOfSignatures")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RequiredAmountOfSignatures is a free data retrieval call binding the contract method 0x6beedd39.
//
// Solidity: function requiredAmountOfSignatures() view returns(uint256)
func (_Cratercontract *CratercontractSession) RequiredAmountOfSignatures() (*big.Int, error) {
	return _Cratercontract.Contract.RequiredAmountOfSignatures(&_Cratercontract.CallOpts)
}

// RequiredAmountOfSignatures is a free data retrieval call binding the contract method 0x6beedd39.
//
// Solidity: function requiredAmountOfSignatures() view returns(uint256)
func (_Cratercontract *CratercontractCallerSession) RequiredAmountOfSignatures() (*big.Int, error) {
	return _Cratercontract.Contract.RequiredAmountOfSignatures(&_Cratercontract.CallOpts)
}

// VerifySignatures is a free data retrieval call binding the contract method 0xc7a823e0.
//
// Solidity: function verifySignatures(bytes32 signedHash, bytes signaturesAndAddrs) view returns()
func (_Cratercontract *CratercontractCaller) VerifySignatures(opts *bind.CallOpts, signedHash [32]byte, signaturesAndAddrs []byte) error {
	var out []interface{}
	err := _Cratercontract.contract.Call(opts, &out, "verifySignatures", signedHash, signaturesAndAddrs)

	if err != nil {
		return err
	}

	return err

}

// VerifySignatures is a free data retrieval call binding the contract method 0xc7a823e0.
//
// Solidity: function verifySignatures(bytes32 signedHash, bytes signaturesAndAddrs) view returns()
func (_Cratercontract *CratercontractSession) VerifySignatures(signedHash [32]byte, signaturesAndAddrs []byte) error {
	return _Cratercontract.Contract.VerifySignatures(&_Cratercontract.CallOpts, signedHash, signaturesAndAddrs)
}

// VerifySignatures is a free data retrieval call binding the contract method 0xc7a823e0.
//
// Solidity: function verifySignatures(bytes32 signedHash, bytes signaturesAndAddrs) view returns()
func (_Cratercontract *CratercontractCallerSession) VerifySignatures(signedHash [32]byte, signaturesAndAddrs []byte) error {
	return _Cratercontract.Contract.VerifySignatures(&_Cratercontract.CallOpts, signedHash, signaturesAndAddrs)
}

// SendComet is a paid mutator transaction binding the contract method 0x37d62dd9.
//
// Solidity: function sendComet((bytes,uint256,uint256) comet, bytes signaturesAndAddrs) returns()
func (_Cratercontract *CratercontractTransactor) SendComet(opts *bind.TransactOpts, comet CraterComet, signaturesAndAddrs []byte) (*types.Transaction, error) {
	return _Cratercontract.contract.Transact(opts, "sendComet", comet, signaturesAndAddrs)
}

// SendComet is a paid mutator transaction binding the contract method 0x37d62dd9.
//
// Solidity: function sendComet((bytes,uint256,uint256) comet, bytes signaturesAndAddrs) returns()
func (_Cratercontract *CratercontractSession) SendComet(comet CraterComet, signaturesAndAddrs []byte) (*types.Transaction, error) {
	return _Cratercontract.Contract.SendComet(&_Cratercontract.TransactOpts, comet, signaturesAndAddrs)
}

// SendComet is a paid mutator transaction binding the contract method 0x37d62dd9.
//
// Solidity: function sendComet((bytes,uint256,uint256) comet, bytes signaturesAndAddrs) returns()
func (_Cratercontract *CratercontractTransactorSession) SendComet(comet CraterComet, signaturesAndAddrs []byte) (*types.Transaction, error) {
	return _Cratercontract.Contract.SendComet(&_Cratercontract.TransactOpts, comet, signaturesAndAddrs)
}

// SetupCommittee is a paid mutator transaction binding the contract method 0x078fba2a.
//
// Solidity: function setupCommittee(uint256 _requiredAmountOfSignatures, string[] urls, bytes addrsBytes) returns()
func (_Cratercontract *CratercontractTransactor) SetupCommittee(opts *bind.TransactOpts, _requiredAmountOfSignatures *big.Int, urls []string, addrsBytes []byte) (*types.Transaction, error) {
	return _Cratercontract.contract.Transact(opts, "setupCommittee", _requiredAmountOfSignatures, urls, addrsBytes)
}

// SetupCommittee is a paid mutator transaction binding the contract method 0x078fba2a.
//
// Solidity: function setupCommittee(uint256 _requiredAmountOfSignatures, string[] urls, bytes addrsBytes) returns()
func (_Cratercontract *CratercontractSession) SetupCommittee(_requiredAmountOfSignatures *big.Int, urls []string, addrsBytes []byte) (*types.Transaction, error) {
	return _Cratercontract.Contract.SetupCommittee(&_Cratercontract.TransactOpts, _requiredAmountOfSignatures, urls, addrsBytes)
}

// SetupCommittee is a paid mutator transaction binding the contract method 0x078fba2a.
//
// Solidity: function setupCommittee(uint256 _requiredAmountOfSignatures, string[] urls, bytes addrsBytes) returns()
func (_Cratercontract *CratercontractTransactorSession) SetupCommittee(_requiredAmountOfSignatures *big.Int, urls []string, addrsBytes []byte) (*types.Transaction, error) {
	return _Cratercontract.Contract.SetupCommittee(&_Cratercontract.TransactOpts, _requiredAmountOfSignatures, urls, addrsBytes)
}

// CratercontractCommitteeUpdatedIterator is returned from FilterCommitteeUpdated and is used to iterate over the raw logs and unpacked data for CommitteeUpdated events raised by the Cratercontract contract.
type CratercontractCommitteeUpdatedIterator struct {
	Event *CratercontractCommitteeUpdated // Event containing the contract specifics and raw log

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
func (it *CratercontractCommitteeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CratercontractCommitteeUpdated)
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
		it.Event = new(CratercontractCommitteeUpdated)
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
func (it *CratercontractCommitteeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CratercontractCommitteeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CratercontractCommitteeUpdated represents a CommitteeUpdated event raised by the Cratercontract contract.
type CratercontractCommitteeUpdated struct {
	CommitteeHash [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCommitteeUpdated is a free log retrieval operation binding the contract event 0x831403fd381b3e6ac875d912ec2eee0e0203d0d29f7b3e0c96fc8f582d6db657.
//
// Solidity: event CommitteeUpdated(bytes32 committeeHash)
func (_Cratercontract *CratercontractFilterer) FilterCommitteeUpdated(opts *bind.FilterOpts) (*CratercontractCommitteeUpdatedIterator, error) {

	logs, sub, err := _Cratercontract.contract.FilterLogs(opts, "CommitteeUpdated")
	if err != nil {
		return nil, err
	}
	return &CratercontractCommitteeUpdatedIterator{contract: _Cratercontract.contract, event: "CommitteeUpdated", logs: logs, sub: sub}, nil
}

// WatchCommitteeUpdated is a free log subscription operation binding the contract event 0x831403fd381b3e6ac875d912ec2eee0e0203d0d29f7b3e0c96fc8f582d6db657.
//
// Solidity: event CommitteeUpdated(bytes32 committeeHash)
func (_Cratercontract *CratercontractFilterer) WatchCommitteeUpdated(opts *bind.WatchOpts, sink chan<- *CratercontractCommitteeUpdated) (event.Subscription, error) {

	logs, sub, err := _Cratercontract.contract.WatchLogs(opts, "CommitteeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CratercontractCommitteeUpdated)
				if err := _Cratercontract.contract.UnpackLog(event, "CommitteeUpdated", log); err != nil {
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

// ParseCommitteeUpdated is a log parse operation binding the contract event 0x831403fd381b3e6ac875d912ec2eee0e0203d0d29f7b3e0c96fc8f582d6db657.
//
// Solidity: event CommitteeUpdated(bytes32 committeeHash)
func (_Cratercontract *CratercontractFilterer) ParseCommitteeUpdated(log types.Log) (*CratercontractCommitteeUpdated, error) {
	event := new(CratercontractCommitteeUpdated)
	if err := _Cratercontract.contract.UnpackLog(event, "CommitteeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

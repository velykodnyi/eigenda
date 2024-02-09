// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractMockRollkit

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

// BN254G1Point is an auto generated low-level Go binding around an user-defined struct.
type BN254G1Point struct {
	X *big.Int
	Y *big.Int
}

// BlobData is an auto generated low-level Go binding around an user-defined struct.
type BlobData struct {
	BlobIndex           uint32
	BlobBatchHeaderHash []byte
}

// EigenDARollupUtilsBlobVerificationProof is an auto generated low-level Go binding around an user-defined struct.
type EigenDARollupUtilsBlobVerificationProof struct {
	BatchId                uint32
	BlobIndex              uint8
	BatchMetadata          IEigenDAServiceManagerBatchMetadata
	InclusionProof         []byte
	QuorumThresholdIndexes []byte
}

// IEigenDAServiceManagerBatchHeader is an auto generated low-level Go binding around an user-defined struct.
type IEigenDAServiceManagerBatchHeader struct {
	BlobHeadersRoot            [32]byte
	QuorumNumbers              []byte
	QuorumThresholdPercentages []byte
	ReferenceBlockNumber       uint32
}

// IEigenDAServiceManagerBatchMetadata is an auto generated low-level Go binding around an user-defined struct.
type IEigenDAServiceManagerBatchMetadata struct {
	BatchHeader             IEigenDAServiceManagerBatchHeader
	SignatoryRecordHash     [32]byte
	Fee                     *big.Int
	ConfirmationBlockNumber uint32
}

// IEigenDAServiceManagerBlobHeader is an auto generated low-level Go binding around an user-defined struct.
type IEigenDAServiceManagerBlobHeader struct {
	Commitment       BN254G1Point
	DataLength       uint32
	QuorumBlobParams []IEigenDAServiceManagerQuorumBlobParam
}

// IEigenDAServiceManagerQuorumBlobParam is an auto generated low-level Go binding around an user-defined struct.
type IEigenDAServiceManagerQuorumBlobParam struct {
	QuorumNumber                 uint8
	AdversaryThresholdPercentage uint8
	QuorumThresholdPercentage    uint8
	ChunkLength                  uint32
}

// ContractMockRollkitMetaData contains all meta data concerning the ContractMockRollkit contract.
var ContractMockRollkitMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_eigenDAServiceManager\",\"type\":\"address\",\"internalType\":\"contractIEigenDAServiceManager\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"eigenDAServiceManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIEigenDAServiceManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBlobIdentificationData\",\"inputs\":[{\"name\":\"blobId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structBlobData\",\"components\":[{\"name\":\"blobIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"blobBatchHeaderHash\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLatestBlobId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"submitBlobs\",\"inputs\":[{\"name\":\"blobIds\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"_blobHeaders\",\"type\":\"tuple[]\",\"internalType\":\"structIEigenDAServiceManager.BlobHeader[]\",\"components\":[{\"name\":\"commitment\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"dataLength\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumBlobParams\",\"type\":\"tuple[]\",\"internalType\":\"structIEigenDAServiceManager.QuorumBlobParam[]\",\"components\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"adversaryThresholdPercentage\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"chunkLength\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}]},{\"name\":\"_blobVerificationProofs\",\"type\":\"tuple[]\",\"internalType\":\"structEigenDARollupUtils.BlobVerificationProof[]\",\"components\":[{\"name\":\"batchId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"blobIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"batchMetadata\",\"type\":\"tuple\",\"internalType\":\"structIEigenDAServiceManager.BatchMetadata\",\"components\":[{\"name\":\"batchHeader\",\"type\":\"tuple\",\"internalType\":\"structIEigenDAServiceManager.BatchHeader\",\"components\":[{\"name\":\"blobHeadersRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentages\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"signatoryRecordHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"fee\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"confirmationBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"inclusionProof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdIndexes\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"_batchHeaderHashes\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"}]",
	Bin: "0x608060405234801561001057600080fd5b50604051610cb7380380610cb783398101604081905261002f91610054565b600080546001600160a01b0319166001600160a01b0392909216919091179055610084565b60006020828403121561006657600080fd5b81516001600160a01b038116811461007d57600080fd5b9392505050565b610c24806100936000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c80635059cdac146100515780637c565b4a14610072578063b33baf8214610092578063fc30cad0146100a7575b600080fd5b60025460405167ffffffffffffffff90911681526020015b60405180910390f35b6100856100803660046104f1565b6100d2565b6040516100699190610522565b6100a56100a0366004610622565b6101b1565b005b6000546100ba906001600160a01b031681565b6040516001600160a01b039091168152602001610069565b60408051808201909152600081526060602082015267ffffffffffffffff82166000908152600160208181526040928390208351808501909452805463ffffffff168452918201805491840191610128906107c8565b80601f0160208091040260200160405190810160405280929190818152602001828054610154906107c8565b80156101a15780601f10610176576101008083540402835291602001916101a1565b820191906000526020600020905b81548152906001019060200180831161018457829003601f168201915b5050505050815250509050919050565b85841480156101bf57508382145b6102475760405162461bcd60e51b815260206004820152604960248201527f4c656e677468206f6620626c6f624964732c20626c6f62486561646572732c2060448201527f616e6420626c6f62566572696669636174696f6e50726f6f6673206d7573742060648201526831329032b8bab0b61760b91b608482015260a40160405180910390fd5b60005b8481101561044e5773__$32f04d18c688c2c57b0347c20a77f3d6c9$__638651838487878481811061027e5761027e610803565b90506020028101906102909190610819565b6000546001600160a01b03168787868181106102ae576102ae610803565b90506020028101906102c09190610839565b6040518463ffffffff1660e01b81526004016102de93929190610af8565b60006040518083038186803b1580156102f657600080fd5b505af415801561030a573d6000803e3d6000fd5b50505050604051806040016040528085858481811061032b5761032b610803565b905060200281019061033d9190610839565b61034e906040810190602001610baa565b60ff1663ffffffff16815260200183838151811061036e5761036e610803565b6020026020010151815250600160008a8a8581811061038f5761038f610803565b90506020020160208101906103a491906104f1565b67ffffffffffffffff1681526020808201929092526040016000208251815463ffffffff191663ffffffff90911617815582820151805191926103ef92600185019290910190610458565b5090505087878281811061040557610405610803565b905060200201602081019061041a91906104f1565b6002805467ffffffffffffffff191667ffffffffffffffff929092169190911790558061044681610bc5565b91505061024a565b5050505050505050565b828054610464906107c8565b90600052602060002090601f01602090048101928261048657600085556104cc565b82601f1061049f57805160ff19168380011785556104cc565b828001600101855582156104cc579182015b828111156104cc5782518255916020019190600101906104b1565b506104d89291506104dc565b5090565b5b808211156104d857600081556001016104dd565b60006020828403121561050357600080fd5b813567ffffffffffffffff8116811461051b57600080fd5b9392505050565b6000602080835263ffffffff8451168184015280840151604080850152805180606086015260005b818110156105665782810184015186820160800152830161054a565b81811115610578576000608083880101525b50601f01601f191693909301608001949350505050565b60008083601f8401126105a157600080fd5b50813567ffffffffffffffff8111156105b957600080fd5b6020830191508360208260051b85010111156105d457600080fd5b9250929050565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff8111828210171561061a5761061a6105db565b604052919050565b60008060008060008060006080888a03121561063d57600080fd5b67ffffffffffffffff808935111561065457600080fd5b6106618a8a358b0161058f565b9098509650602089013581101561067757600080fd5b6106878a60208b01358b0161058f565b9096509450604089013581101561069d57600080fd5b6106ad8a60408b01358b0161058f565b909450925060608901358110156106c357600080fd5b606089013589018a601f8201126106d957600080fd5b81813511156106ea576106ea6105db565b6106fa6020823560051b016105f1565b81358082526020808301929160051b8401018d101561071857600080fd5b602083015b6020843560051b8501018110156107b457848135111561073c57600080fd5b8d603f82358601011261074e57600080fd5b602081358501013585811115610766576107666105db565b610779601f8201601f19166020016105f1565b8181528f604083853589010101111561079157600080fd5b81604084358801016020830137600060209282018301528452928301920161071d565b508094505050505092959891949750929550565b600181811c908216806107dc57607f821691505b602082108114156107fd57634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052603260045260246000fd5b60008235607e1983360301811261082f57600080fd5b9190910192915050565b60008235609e1983360301811261082f57600080fd5b803563ffffffff8116811461086357600080fd5b919050565b803560ff8116811461086357600080fd5b8183526000602080850194508260005b858110156108ef5760ff8061089d84610868565b168852806108ac858501610868565b16848901526040816108bf828601610868565b169089015250606063ffffffff6108d784830161084f565b16908801526080968701969190910190600101610889565b509495945050505050565b60008235607e1983360301811261091057600080fd5b90910192915050565b6000808335601e1984360301811261093057600080fd5b830160208101925035905067ffffffffffffffff81111561095057600080fd5b8036038313156105d457600080fd5b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b80356bffffffffffffffffffffffff8116811461086357600080fd5b600063ffffffff806109b58461084f565b16845260ff6109c660208501610868565b1660208501526109d960408401846108fa565b60a060408601526109ea81826108fa565b608060a08701528035610120870152610a066020820182610919565b6080610140890152610a1d6101a08901828461095f565b915050610a2d6040830183610919565b88830361011f19016101608a0152610a4683828461095f565b9250505083610a576060840161084f565b16610180880152602083013560c0880152610a7460408401610988565b6bffffffffffffffffffffffff811660e08901529350610a966060840161084f565b63ffffffff81166101008901529350610ab26060870187610919565b945092508681036060880152610ac981858561095f565b9350505050610adb6080840184610919565b8583036080870152610aee83828461095f565b9695505050505050565b60608152833560608201526020840135608082015263ffffffff610b1e6040860161084f565b1660a082015260006060850135601e19863603018112610b3d57600080fd5b8501803567ffffffffffffffff811115610b5657600080fd5b8060071b3603871315610b6857600080fd5b608060c0850152610b8060e085018260208501610879565b915050610b9860208401866001600160a01b03169052565b8281036040840152610aee81856109a4565b600060208284031215610bbc57600080fd5b61051b82610868565b6000600019821415610be757634e487b7160e01b600052601160045260246000fd5b506001019056fea2646970667358221220ad798aa90f1db5b13789c0e640533fdc6063e543097d10b6f696d1086fa6203264736f6c634300080c0033",
}

// ContractMockRollkitABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMockRollkitMetaData.ABI instead.
var ContractMockRollkitABI = ContractMockRollkitMetaData.ABI

// ContractMockRollkitBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractMockRollkitMetaData.Bin instead.
var ContractMockRollkitBin = ContractMockRollkitMetaData.Bin

// DeployContractMockRollkit deploys a new Ethereum contract, binding an instance of ContractMockRollkit to it.
func DeployContractMockRollkit(auth *bind.TransactOpts, backend bind.ContractBackend, _eigenDAServiceManager common.Address) (common.Address, *types.Transaction, *ContractMockRollkit, error) {
	parsed, err := ContractMockRollkitMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractMockRollkitBin), backend, _eigenDAServiceManager)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractMockRollkit{ContractMockRollkitCaller: ContractMockRollkitCaller{contract: contract}, ContractMockRollkitTransactor: ContractMockRollkitTransactor{contract: contract}, ContractMockRollkitFilterer: ContractMockRollkitFilterer{contract: contract}}, nil
}

// ContractMockRollkit is an auto generated Go binding around an Ethereum contract.
type ContractMockRollkit struct {
	ContractMockRollkitCaller     // Read-only binding to the contract
	ContractMockRollkitTransactor // Write-only binding to the contract
	ContractMockRollkitFilterer   // Log filterer for contract events
}

// ContractMockRollkitCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractMockRollkitCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractMockRollkitTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractMockRollkitTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractMockRollkitFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractMockRollkitFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractMockRollkitSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractMockRollkitSession struct {
	Contract     *ContractMockRollkit // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ContractMockRollkitCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractMockRollkitCallerSession struct {
	Contract *ContractMockRollkitCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// ContractMockRollkitTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractMockRollkitTransactorSession struct {
	Contract     *ContractMockRollkitTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// ContractMockRollkitRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractMockRollkitRaw struct {
	Contract *ContractMockRollkit // Generic contract binding to access the raw methods on
}

// ContractMockRollkitCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractMockRollkitCallerRaw struct {
	Contract *ContractMockRollkitCaller // Generic read-only contract binding to access the raw methods on
}

// ContractMockRollkitTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractMockRollkitTransactorRaw struct {
	Contract *ContractMockRollkitTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractMockRollkit creates a new instance of ContractMockRollkit, bound to a specific deployed contract.
func NewContractMockRollkit(address common.Address, backend bind.ContractBackend) (*ContractMockRollkit, error) {
	contract, err := bindContractMockRollkit(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractMockRollkit{ContractMockRollkitCaller: ContractMockRollkitCaller{contract: contract}, ContractMockRollkitTransactor: ContractMockRollkitTransactor{contract: contract}, ContractMockRollkitFilterer: ContractMockRollkitFilterer{contract: contract}}, nil
}

// NewContractMockRollkitCaller creates a new read-only instance of ContractMockRollkit, bound to a specific deployed contract.
func NewContractMockRollkitCaller(address common.Address, caller bind.ContractCaller) (*ContractMockRollkitCaller, error) {
	contract, err := bindContractMockRollkit(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractMockRollkitCaller{contract: contract}, nil
}

// NewContractMockRollkitTransactor creates a new write-only instance of ContractMockRollkit, bound to a specific deployed contract.
func NewContractMockRollkitTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractMockRollkitTransactor, error) {
	contract, err := bindContractMockRollkit(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractMockRollkitTransactor{contract: contract}, nil
}

// NewContractMockRollkitFilterer creates a new log filterer instance of ContractMockRollkit, bound to a specific deployed contract.
func NewContractMockRollkitFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractMockRollkitFilterer, error) {
	contract, err := bindContractMockRollkit(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractMockRollkitFilterer{contract: contract}, nil
}

// bindContractMockRollkit binds a generic wrapper to an already deployed contract.
func bindContractMockRollkit(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractMockRollkitMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractMockRollkit *ContractMockRollkitRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractMockRollkit.Contract.ContractMockRollkitCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractMockRollkit *ContractMockRollkitRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractMockRollkit.Contract.ContractMockRollkitTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractMockRollkit *ContractMockRollkitRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractMockRollkit.Contract.ContractMockRollkitTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractMockRollkit *ContractMockRollkitCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractMockRollkit.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractMockRollkit *ContractMockRollkitTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractMockRollkit.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractMockRollkit *ContractMockRollkitTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractMockRollkit.Contract.contract.Transact(opts, method, params...)
}

// EigenDAServiceManager is a free data retrieval call binding the contract method 0xfc30cad0.
//
// Solidity: function eigenDAServiceManager() view returns(address)
func (_ContractMockRollkit *ContractMockRollkitCaller) EigenDAServiceManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractMockRollkit.contract.Call(opts, &out, "eigenDAServiceManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EigenDAServiceManager is a free data retrieval call binding the contract method 0xfc30cad0.
//
// Solidity: function eigenDAServiceManager() view returns(address)
func (_ContractMockRollkit *ContractMockRollkitSession) EigenDAServiceManager() (common.Address, error) {
	return _ContractMockRollkit.Contract.EigenDAServiceManager(&_ContractMockRollkit.CallOpts)
}

// EigenDAServiceManager is a free data retrieval call binding the contract method 0xfc30cad0.
//
// Solidity: function eigenDAServiceManager() view returns(address)
func (_ContractMockRollkit *ContractMockRollkitCallerSession) EigenDAServiceManager() (common.Address, error) {
	return _ContractMockRollkit.Contract.EigenDAServiceManager(&_ContractMockRollkit.CallOpts)
}

// GetBlobIdentificationData is a free data retrieval call binding the contract method 0x7c565b4a.
//
// Solidity: function getBlobIdentificationData(uint64 blobId) view returns((uint32,bytes))
func (_ContractMockRollkit *ContractMockRollkitCaller) GetBlobIdentificationData(opts *bind.CallOpts, blobId uint64) (BlobData, error) {
	var out []interface{}
	err := _ContractMockRollkit.contract.Call(opts, &out, "getBlobIdentificationData", blobId)

	if err != nil {
		return *new(BlobData), err
	}

	out0 := *abi.ConvertType(out[0], new(BlobData)).(*BlobData)

	return out0, err

}

// GetBlobIdentificationData is a free data retrieval call binding the contract method 0x7c565b4a.
//
// Solidity: function getBlobIdentificationData(uint64 blobId) view returns((uint32,bytes))
func (_ContractMockRollkit *ContractMockRollkitSession) GetBlobIdentificationData(blobId uint64) (BlobData, error) {
	return _ContractMockRollkit.Contract.GetBlobIdentificationData(&_ContractMockRollkit.CallOpts, blobId)
}

// GetBlobIdentificationData is a free data retrieval call binding the contract method 0x7c565b4a.
//
// Solidity: function getBlobIdentificationData(uint64 blobId) view returns((uint32,bytes))
func (_ContractMockRollkit *ContractMockRollkitCallerSession) GetBlobIdentificationData(blobId uint64) (BlobData, error) {
	return _ContractMockRollkit.Contract.GetBlobIdentificationData(&_ContractMockRollkit.CallOpts, blobId)
}

// GetLatestBlobId is a free data retrieval call binding the contract method 0x5059cdac.
//
// Solidity: function getLatestBlobId() view returns(uint64)
func (_ContractMockRollkit *ContractMockRollkitCaller) GetLatestBlobId(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _ContractMockRollkit.contract.Call(opts, &out, "getLatestBlobId")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetLatestBlobId is a free data retrieval call binding the contract method 0x5059cdac.
//
// Solidity: function getLatestBlobId() view returns(uint64)
func (_ContractMockRollkit *ContractMockRollkitSession) GetLatestBlobId() (uint64, error) {
	return _ContractMockRollkit.Contract.GetLatestBlobId(&_ContractMockRollkit.CallOpts)
}

// GetLatestBlobId is a free data retrieval call binding the contract method 0x5059cdac.
//
// Solidity: function getLatestBlobId() view returns(uint64)
func (_ContractMockRollkit *ContractMockRollkitCallerSession) GetLatestBlobId() (uint64, error) {
	return _ContractMockRollkit.Contract.GetLatestBlobId(&_ContractMockRollkit.CallOpts)
}

// SubmitBlobs is a paid mutator transaction binding the contract method 0xb33baf82.
//
// Solidity: function submitBlobs(uint64[] blobIds, ((uint256,uint256),uint32,(uint8,uint8,uint8,uint32)[])[] _blobHeaders, (uint32,uint8,((bytes32,bytes,bytes,uint32),bytes32,uint96,uint32),bytes,bytes)[] _blobVerificationProofs, bytes[] _batchHeaderHashes) returns()
func (_ContractMockRollkit *ContractMockRollkitTransactor) SubmitBlobs(opts *bind.TransactOpts, blobIds []uint64, _blobHeaders []IEigenDAServiceManagerBlobHeader, _blobVerificationProofs []EigenDARollupUtilsBlobVerificationProof, _batchHeaderHashes [][]byte) (*types.Transaction, error) {
	return _ContractMockRollkit.contract.Transact(opts, "submitBlobs", blobIds, _blobHeaders, _blobVerificationProofs, _batchHeaderHashes)
}

// SubmitBlobs is a paid mutator transaction binding the contract method 0xb33baf82.
//
// Solidity: function submitBlobs(uint64[] blobIds, ((uint256,uint256),uint32,(uint8,uint8,uint8,uint32)[])[] _blobHeaders, (uint32,uint8,((bytes32,bytes,bytes,uint32),bytes32,uint96,uint32),bytes,bytes)[] _blobVerificationProofs, bytes[] _batchHeaderHashes) returns()
func (_ContractMockRollkit *ContractMockRollkitSession) SubmitBlobs(blobIds []uint64, _blobHeaders []IEigenDAServiceManagerBlobHeader, _blobVerificationProofs []EigenDARollupUtilsBlobVerificationProof, _batchHeaderHashes [][]byte) (*types.Transaction, error) {
	return _ContractMockRollkit.Contract.SubmitBlobs(&_ContractMockRollkit.TransactOpts, blobIds, _blobHeaders, _blobVerificationProofs, _batchHeaderHashes)
}

// SubmitBlobs is a paid mutator transaction binding the contract method 0xb33baf82.
//
// Solidity: function submitBlobs(uint64[] blobIds, ((uint256,uint256),uint32,(uint8,uint8,uint8,uint32)[])[] _blobHeaders, (uint32,uint8,((bytes32,bytes,bytes,uint32),bytes32,uint96,uint32),bytes,bytes)[] _blobVerificationProofs, bytes[] _batchHeaderHashes) returns()
func (_ContractMockRollkit *ContractMockRollkitTransactorSession) SubmitBlobs(blobIds []uint64, _blobHeaders []IEigenDAServiceManagerBlobHeader, _blobVerificationProofs []EigenDARollupUtilsBlobVerificationProof, _batchHeaderHashes [][]byte) (*types.Transaction, error) {
	return _ContractMockRollkit.Contract.SubmitBlobs(&_ContractMockRollkit.TransactOpts, blobIds, _blobHeaders, _blobVerificationProofs, _batchHeaderHashes)
}

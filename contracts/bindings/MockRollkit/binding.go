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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_eigenDAServiceManager\",\"type\":\"address\",\"internalType\":\"contractIEigenDAServiceManager\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"eigenDAServiceManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIEigenDAServiceManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBlobIdentificationData\",\"inputs\":[{\"name\":\"blobId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structBlobData\",\"components\":[{\"name\":\"blobIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"blobBatchHeaderHash\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"latestBlobId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"submitBlobs\",\"inputs\":[{\"name\":\"blobIds\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"_blobHeaders\",\"type\":\"tuple[]\",\"internalType\":\"structIEigenDAServiceManager.BlobHeader[]\",\"components\":[{\"name\":\"commitment\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"dataLength\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumBlobParams\",\"type\":\"tuple[]\",\"internalType\":\"structIEigenDAServiceManager.QuorumBlobParam[]\",\"components\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"adversaryThresholdPercentage\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"chunkLength\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}]},{\"name\":\"_blobVerificationProofs\",\"type\":\"tuple[]\",\"internalType\":\"structEigenDARollupUtils.BlobVerificationProof[]\",\"components\":[{\"name\":\"batchId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"blobIndex\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"batchMetadata\",\"type\":\"tuple\",\"internalType\":\"structIEigenDAServiceManager.BatchMetadata\",\"components\":[{\"name\":\"batchHeader\",\"type\":\"tuple\",\"internalType\":\"structIEigenDAServiceManager.BatchHeader\",\"components\":[{\"name\":\"blobHeadersRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentages\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"signatoryRecordHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"fee\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"confirmationBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"inclusionProof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdIndexes\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"_batchHeaderHashes\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"}]",
	Bin: "0x6080604052600280546001600160401b031916905534801561002057600080fd5b50604051610cdf380380610cdf83398101604081905261003f91610064565b600080546001600160a01b0319166001600160a01b0392909216919091179055610094565b60006020828403121561007657600080fd5b81516001600160a01b038116811461008d57600080fd5b9392505050565b610c3c806100a36000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c8063546aef8d146100515780637c565b4a14610081578063b33baf82146100a1578063fc30cad0146100c2575b600080fd5b600254610064906001600160401b031681565b6040516001600160401b0390911681526020015b60405180910390f35b61009461008f3660046104e0565b6100ed565b6040516100789190610510565b6100b46100af36600461060e565b6101cb565b604051908152602001610078565b6000546100d5906001600160a01b031681565b6040516001600160a01b039091168152602001610078565b6040805180820190915260008152606060208201526001600160401b0382166000908152600160208181526040928390208351808501909452805463ffffffff168452918201805491840191610142906107b3565b80601f016020809104026020016040519081016040528092919081815260200182805461016e906107b3565b80156101bb5780601f10610190576101008083540402835291602001916101bb565b820191906000526020600020905b81548152906001019060200180831161019e57829003601f168201915b5050505050815250509050919050565b600086851480156101db57508483145b6102635760405162461bcd60e51b815260206004820152604960248201527f4c656e677468206f6620626c6f624964732c20626c6f62486561646572732c2060448201527f616e6420626c6f62566572696669636174696f6e50726f6f6673206d7573742060648201526831329032b8bab0b61760b91b608482015260a40160405180910390fd5b60005b8581101561042e5773__$32f04d18c688c2c57b0347c20a77f3d6c9$__638651838488888481811061029a5761029a6107ee565b90506020028101906102ac9190610804565b6000546001600160a01b03168888868181106102ca576102ca6107ee565b90506020028101906102dc9190610824565b6040518463ffffffff1660e01b81526004016102fa93929190610ae2565b60006040518083038186803b15801561031257600080fd5b505af4158015610326573d6000803e3d6000fd5b505050506040518060400160405280868684818110610347576103476107ee565b90506020028101906103599190610824565b61036a906040810190602001610b93565b60ff1663ffffffff16815260200184838151811061038a5761038a6107ee565b602090810291909101015190526002805460019160009182906103b5906001600160401b0316610bc4565b82546001600160401b039182166101009390930a838102920219161790915581526020808201929092526040016000208251815463ffffffff90911663ffffffff19909116178155828201518051919261041792600185019290910190610447565b50905050808061042690610beb565b915050610266565b50506002546001600160401b0316979650505050505050565b828054610453906107b3565b90600052602060002090601f01602090048101928261047557600085556104bb565b82601f1061048e57805160ff19168380011785556104bb565b828001600101855582156104bb579182015b828111156104bb5782518255916020019190600101906104a0565b506104c79291506104cb565b5090565b5b808211156104c757600081556001016104cc565b6000602082840312156104f257600080fd5b81356001600160401b038116811461050957600080fd5b9392505050565b6000602080835263ffffffff8451168184015280840151604080850152805180606086015260005b8181101561055457828101840151868201608001528301610538565b81811115610566576000608083880101525b50601f01601f191693909301608001949350505050565b60008083601f84011261058f57600080fd5b5081356001600160401b038111156105a657600080fd5b6020830191508360208260051b85010111156105c157600080fd5b9250929050565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f191681016001600160401b0381118282101715610606576106066105c8565b604052919050565b60008060008060008060006080888a03121561062957600080fd5b6001600160401b03808935111561063f57600080fd5b61064c8a8a358b0161057d565b9098509650602089013581101561066257600080fd5b6106728a60208b01358b0161057d565b9096509450604089013581101561068857600080fd5b6106988a60408b01358b0161057d565b909450925060608901358110156106ae57600080fd5b606089013589018a601f8201126106c457600080fd5b81813511156106d5576106d56105c8565b6106e56020823560051b016105de565b81358082526020808301929160051b8401018d101561070357600080fd5b602083015b6020843560051b85010181101561079f57848135111561072757600080fd5b8d603f82358601011261073957600080fd5b602081358501013585811115610751576107516105c8565b610764601f8201601f19166020016105de565b8181528f604083853589010101111561077c57600080fd5b816040843588010160208301376000602092820183015284529283019201610708565b508094505050505092959891949750929550565b600181811c908216806107c757607f821691505b602082108114156107e857634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052603260045260246000fd5b60008235607e1983360301811261081a57600080fd5b9190910192915050565b60008235609e1983360301811261081a57600080fd5b803563ffffffff8116811461084e57600080fd5b919050565b803560ff8116811461084e57600080fd5b8183526000602080850194508260005b858110156108da5760ff8061088884610853565b16885280610897858501610853565b16848901526040816108aa828601610853565b169089015250606063ffffffff6108c284830161083a565b16908801526080968701969190910190600101610874565b509495945050505050565b60008235607e198336030181126108fb57600080fd5b90910192915050565b6000808335601e1984360301811261091b57600080fd5b83016020810192503590506001600160401b0381111561093a57600080fd5b8036038313156105c157600080fd5b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b80356bffffffffffffffffffffffff8116811461084e57600080fd5b600063ffffffff8061099f8461083a565b16845260ff6109b060208501610853565b1660208501526109c360408401846108e5565b60a060408601526109d481826108e5565b608060a087015280356101208701526109f06020820182610904565b6080610140890152610a076101a089018284610949565b915050610a176040830183610904565b88830361011f19016101608a0152610a30838284610949565b9250505083610a416060840161083a565b16610180880152602083013560c0880152610a5e60408401610972565b6bffffffffffffffffffffffff811660e08901529350610a806060840161083a565b63ffffffff81166101008901529350610a9c6060870187610904565b945092508681036060880152610ab3818585610949565b9350505050610ac56080840184610904565b8583036080870152610ad8838284610949565b9695505050505050565b60608152833560608201526020840135608082015263ffffffff610b086040860161083a565b1660a082015260006060850135601e19863603018112610b2757600080fd5b850180356001600160401b03811115610b3f57600080fd5b8060071b3603871315610b5157600080fd5b608060c0850152610b6960e085018260208501610864565b915050610b8160208401866001600160a01b03169052565b8281036040840152610ad8818561098e565b600060208284031215610ba557600080fd5b61050982610853565b634e487b7160e01b600052601160045260246000fd5b60006001600160401b0380831681811415610be157610be1610bae565b6001019392505050565b6000600019821415610bff57610bff610bae565b506001019056fea264697066735822122008ef6c18469dcddb9df00dbb020fa853e2722b3e05d8862c65d8c9509407ba0464736f6c634300080c0033",
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

// LatestBlobId is a free data retrieval call binding the contract method 0x546aef8d.
//
// Solidity: function latestBlobId() view returns(uint64)
func (_ContractMockRollkit *ContractMockRollkitCaller) LatestBlobId(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _ContractMockRollkit.contract.Call(opts, &out, "latestBlobId")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LatestBlobId is a free data retrieval call binding the contract method 0x546aef8d.
//
// Solidity: function latestBlobId() view returns(uint64)
func (_ContractMockRollkit *ContractMockRollkitSession) LatestBlobId() (uint64, error) {
	return _ContractMockRollkit.Contract.LatestBlobId(&_ContractMockRollkit.CallOpts)
}

// LatestBlobId is a free data retrieval call binding the contract method 0x546aef8d.
//
// Solidity: function latestBlobId() view returns(uint64)
func (_ContractMockRollkit *ContractMockRollkitCallerSession) LatestBlobId() (uint64, error) {
	return _ContractMockRollkit.Contract.LatestBlobId(&_ContractMockRollkit.CallOpts)
}

// SubmitBlobs is a paid mutator transaction binding the contract method 0xb33baf82.
//
// Solidity: function submitBlobs(uint64[] blobIds, ((uint256,uint256),uint32,(uint8,uint8,uint8,uint32)[])[] _blobHeaders, (uint32,uint8,((bytes32,bytes,bytes,uint32),bytes32,uint96,uint32),bytes,bytes)[] _blobVerificationProofs, bytes[] _batchHeaderHashes) returns(uint256)
func (_ContractMockRollkit *ContractMockRollkitTransactor) SubmitBlobs(opts *bind.TransactOpts, blobIds []uint64, _blobHeaders []IEigenDAServiceManagerBlobHeader, _blobVerificationProofs []EigenDARollupUtilsBlobVerificationProof, _batchHeaderHashes [][]byte) (*types.Transaction, error) {
	return _ContractMockRollkit.contract.Transact(opts, "submitBlobs", blobIds, _blobHeaders, _blobVerificationProofs, _batchHeaderHashes)
}

// SubmitBlobs is a paid mutator transaction binding the contract method 0xb33baf82.
//
// Solidity: function submitBlobs(uint64[] blobIds, ((uint256,uint256),uint32,(uint8,uint8,uint8,uint32)[])[] _blobHeaders, (uint32,uint8,((bytes32,bytes,bytes,uint32),bytes32,uint96,uint32),bytes,bytes)[] _blobVerificationProofs, bytes[] _batchHeaderHashes) returns(uint256)
func (_ContractMockRollkit *ContractMockRollkitSession) SubmitBlobs(blobIds []uint64, _blobHeaders []IEigenDAServiceManagerBlobHeader, _blobVerificationProofs []EigenDARollupUtilsBlobVerificationProof, _batchHeaderHashes [][]byte) (*types.Transaction, error) {
	return _ContractMockRollkit.Contract.SubmitBlobs(&_ContractMockRollkit.TransactOpts, blobIds, _blobHeaders, _blobVerificationProofs, _batchHeaderHashes)
}

// SubmitBlobs is a paid mutator transaction binding the contract method 0xb33baf82.
//
// Solidity: function submitBlobs(uint64[] blobIds, ((uint256,uint256),uint32,(uint8,uint8,uint8,uint32)[])[] _blobHeaders, (uint32,uint8,((bytes32,bytes,bytes,uint32),bytes32,uint96,uint32),bytes,bytes)[] _blobVerificationProofs, bytes[] _batchHeaderHashes) returns(uint256)
func (_ContractMockRollkit *ContractMockRollkitTransactorSession) SubmitBlobs(blobIds []uint64, _blobHeaders []IEigenDAServiceManagerBlobHeader, _blobVerificationProofs []EigenDARollupUtilsBlobVerificationProof, _batchHeaderHashes [][]byte) (*types.Transaction, error) {
	return _ContractMockRollkit.Contract.SubmitBlobs(&_ContractMockRollkit.TransactOpts, blobIds, _blobHeaders, _blobVerificationProofs, _batchHeaderHashes)
}

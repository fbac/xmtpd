// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package identityupdates

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

// IdentityUpdatesMetaData contains all meta data concerning the IdentityUpdates contract.
var IdentityUpdatesMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"ABSOLUTE_MAX_PAYLOAD_SIZE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ABSOLUTE_MIN_PAYLOAD_SIZE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addIdentityUpdate\",\"inputs\":[{\"name\":\"inboxId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"update\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"admin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"maxPayloadSize\",\"inputs\":[],\"outputs\":[{\"name\":\"size\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minPayloadSize\",\"inputs\":[],\"outputs\":[{\"name\":\"size\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMaxPayloadSize\",\"inputs\":[{\"name\":\"maxPayloadSizeRequest\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMinPayloadSize\",\"inputs\":[{\"name\":\"minPayloadSizeRequest\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"IdentityUpdateCreated\",\"inputs\":[{\"name\":\"inboxId\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"update\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"sequenceId\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MaxPayloadSizeUpdated\",\"inputs\":[{\"name\":\"oldSize\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newSize\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MinPayloadSizeUpdated\",\"inputs\":[{\"name\":\"oldSize\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newSize\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpgradeAuthorized\",\"inputs\":[{\"name\":\"upgrader\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newImplementation\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExpectedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidMaxPayloadSize\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidMinPayloadSize\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPayloadSize\",\"inputs\":[{\"name\":\"actualSize\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minSize\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxSize\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnauthorizedCallContext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnsupportedProxiableUUID\",\"inputs\":[{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ZeroAdminAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroImplementationAddress\",\"inputs\":[]}]",
	Bin: "0x60a0604052306080523480156012575f5ffd5b50608051611a4b6100395f395f8181610e2901528181610e5201526111110152611a4b5ff3fe608060405260043610610161575f3560e01c806358e3e94c116100c6578063ad3cb1cc1161007c578063d547741f11610057578063d547741f1461044b578063f96927ac1461046a578063fe8e37a31461049d575f5ffd5b8063ad3cb1cc146103b8578063ba74fc7c1461040d578063c4d66de81461042c575f5ffd5b80638456cb59116100ac5780638456cb591461032157806391d1485414610335578063a217fddf146103a5575f5ffd5b806358e3e94c146102b85780635c975abb146102eb575f5ffd5b8063314a100e1161011b5780633f4ba83a116101015780633f4ba83a1461027d5780634f1ef2861461029157806352d1902d146102a4575f5ffd5b8063314a100e1461023f57806336568abe1461025e575f5ffd5b80631de015991161014b5780631de01599146101bd578063248a9ca3146101d15780632f2ff15d1461021e575f5ffd5b806209e1271461016557806301ffc9a71461018e575b5f5ffd5b348015610170575f5ffd5b5061017b6240000081565b6040519081526020015b60405180910390f35b348015610199575f5ffd5b506101ad6101a83660046116c9565b6104bc565b6040519015158152602001610185565b3480156101c8575f5ffd5b5061017b604e81565b3480156101dc575f5ffd5b5061017b6101eb366004611708565b5f9081527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015490565b348015610229575f5ffd5b5061023d610238366004611747565b610554565b005b34801561024a575f5ffd5b5061023d610259366004611708565b61059d565b348015610269575f5ffd5b5061023d610278366004611747565b610674565b348015610288575f5ffd5b5061023d6106d2565b61023d61029f36600461179e565b6106e7565b3480156102af575f5ffd5b5061017b610706565b3480156102c3575f5ffd5b507f92f6d7b379434335724ccaa6ce32661f25de0b6cb746fac5f5edaed4b9685e015461017b565b3480156102f6575f5ffd5b507fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff166101ad565b34801561032c575f5ffd5b5061023d610734565b348015610340575f5ffd5b506101ad61034f366004611747565b5f9182527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020908152604080842073ffffffffffffffffffffffffffffffffffffffff93909316845291905290205460ff1690565b3480156103b0575f5ffd5b5061017b5f81565b3480156103c3575f5ffd5b506104006040518060400160405280600581526020017f352e302e3000000000000000000000000000000000000000000000000000000081525081565b604051610185919061189f565b348015610418575f5ffd5b5061023d6104273660046118f2565b610746565b348015610437575f5ffd5b5061023d610446366004611969565b610848565b348015610456575f5ffd5b5061023d610465366004611747565b610a7a565b348015610475575f5ffd5b507f92f6d7b379434335724ccaa6ce32661f25de0b6cb746fac5f5edaed4b9685e005461017b565b3480156104a8575f5ffd5b5061023d6104b7366004611708565b610abd565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b00000000000000000000000000000000000000000000000000000000148061054e57507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316145b92915050565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015461058d81610b70565b6105978383610b7a565b50505050565b5f6105a781610b70565b7f92f6d7b379434335724ccaa6ce32661f25de0b6cb746fac5f5edaed4b9685e01547f92f6d7b379434335724ccaa6ce32661f25de0b6cb746fac5f5edaed4b9685e00908311806105f85750604e83105b1561062f576040517fe219e4f000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b805483825560408051828152602081018690527f1ee836faee0e7c61d20a079d0b5b4e1ee9c536e18268ef6f7c620dcec82f72cd91015b60405180910390a150505050565b73ffffffffffffffffffffffffffffffffffffffff811633146106c3576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6106cd8282610c98565b505050565b5f6106dc81610b70565b6106e4610d74565b50565b6106ef610e11565b6106f882610f17565b6107028282610fc0565b5050565b5f61070f6110f9565b507f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc90565b5f61073e81610b70565b6106e4611168565b61074e6111e1565b7f92f6d7b379434335724ccaa6ce32661f25de0b6cb746fac5f5edaed4b9685e0080548210806107815750600181015482115b156107d457805460018201546040517f93b7abe600000000000000000000000000000000000000000000000000000000815260048101859052602481019290925260448201526064015b60405180910390fd5b6002810180547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000008116600167ffffffffffffffff92831601918216179091556040517fc1a40f292090ec0435e939cdfe248e0322a88566679a90a50c4e9e5ef762dbd5916106669187918791879190611982565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000810460ff16159067ffffffffffffffff165f811580156108925750825b90505f8267ffffffffffffffff1660011480156108ae5750303b155b9050811580156108bc575080155b156108f3576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b84547fffffffffffffffffffffffffffffffffffffffffffffffff000000000000000016600117855583156109545784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff16680100000000000000001785555b73ffffffffffffffffffffffffffffffffffffffff86166109a1576040517f3ef39b8100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6109a961123d565b6109b161123d565b6109b9611245565b604e7f92f6d7b379434335724ccaa6ce32661f25de0b6cb746fac5f5edaed4b9685e00908155624000007f92f6d7b379434335724ccaa6ce32661f25de0b6cb746fac5f5edaed4b9685e0155610a0f5f88610b7a565b50508315610a725784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b505050505050565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020526040902060010154610ab381610b70565b6105978383610c98565b5f610ac781610b70565b7f92f6d7b379434335724ccaa6ce32661f25de0b6cb746fac5f5edaed4b9685e008054831080610af957506240000083115b15610b30576040517f1d8e7a4a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001810180549084905560408051828152602081018690527ff59e99f8f54d2696b7cf184949ab2b4bbd6336ec1816b36f58ae9948d868fe909101610666565b6106e48133611255565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020818152604080842073ffffffffffffffffffffffffffffffffffffffff8616855290915282205460ff16610c8f575f8481526020828152604080832073ffffffffffffffffffffffffffffffffffffffff87168452909152902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055610c2b3390565b73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a4600191505061054e565b5f91505061054e565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020818152604080842073ffffffffffffffffffffffffffffffffffffffff8616855290915282205460ff1615610c8f575f8481526020828152604080832073ffffffffffffffffffffffffffffffffffffffff8716808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a4600191505061054e565b610d7c6112fb565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001681557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390a150565b3073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000161480610ede57507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16610ec57f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff1614155b15610f15576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b5f610f2181610b70565b73ffffffffffffffffffffffffffffffffffffffff8216610f6e576040517fd02c623d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040805133815273ffffffffffffffffffffffffffffffffffffffff841660208201527fd30e1d298bf814ea43d22b4ce8298062b08609cd67496483769d836157dd52fa910160405180910390a15050565b8173ffffffffffffffffffffffffffffffffffffffff166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015611045575060408051601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0168201909252611042918101906119e8565b60015b611093576040517f4c9c8ce300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff831660048201526024016107cb565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc81146110ef576040517faa1d49a4000000000000000000000000000000000000000000000000000000008152600481018290526024016107cb565b6106cd8383611356565b3073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610f15576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6111706111e1565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011781557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25833610de6565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff1615610f15576040517fd93c066500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610f156113b8565b61124d6113b8565b610f1561141f565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff16610702576040517fe2517d3f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82166004820152602481018390526044016107cb565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff16610f15576040517f8dfc202b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61135f82611470565b60405173ffffffffffffffffffffffffffffffffffffffff8316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a28051156113b0576106cd828261153e565b6107026115bd565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005468010000000000000000900460ff16610f15576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6114276113b8565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055565b8073ffffffffffffffffffffffffffffffffffffffff163b5f036114d8576040517f4c9c8ce300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff821660048201526024016107cb565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60605f5f8473ffffffffffffffffffffffffffffffffffffffff168460405161156791906119ff565b5f60405180830381855af49150503d805f811461159f576040519150601f19603f3d011682016040523d82523d5f602084013e6115a4565b606091505b50915091506115b48583836115f5565b95945050505050565b3415610f15576040517fb398979f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60608261160a5761160582611687565b611680565b815115801561162e575073ffffffffffffffffffffffffffffffffffffffff84163b155b1561167d576040517f9996b31500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff851660048201526024016107cb565b50805b9392505050565b8051156116975780518082602001fd5b6040517fd6bda27500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f602082840312156116d9575f5ffd5b81357fffffffff0000000000000000000000000000000000000000000000000000000081168114611680575f5ffd5b5f60208284031215611718575f5ffd5b5035919050565b803573ffffffffffffffffffffffffffffffffffffffff81168114611742575f5ffd5b919050565b5f5f60408385031215611758575f5ffd5b823591506117686020840161171f565b90509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f5f604083850312156117af575f5ffd5b6117b88361171f565b9150602083013567ffffffffffffffff8111156117d3575f5ffd5b8301601f810185136117e3575f5ffd5b803567ffffffffffffffff8111156117fd576117fd611771565b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0603f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8501160116810181811067ffffffffffffffff8211171561186957611869611771565b604052818152828201602001871015611880575f5ffd5b816020840160208301375f602083830101528093505050509250929050565b602081525f82518060208401528060208501604085015e5f6040828501015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505092915050565b5f5f5f60408486031215611904575f5ffd5b83359250602084013567ffffffffffffffff811115611921575f5ffd5b8401601f81018613611931575f5ffd5b803567ffffffffffffffff811115611947575f5ffd5b866020828401011115611958575f5ffd5b939660209190910195509293505050565b5f60208284031215611979575f5ffd5b6116808261171f565b84815260606020820152826060820152828460808301375f608084830101525f60807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f860116830101905067ffffffffffffffff8316604083015295945050505050565b5f602082840312156119f8575f5ffd5b5051919050565b5f82518060208501845e5f92019182525091905056fea26469706673582212203e0e0ca9f564b4fb7ae14d070b30c93d131a9043728b0f0ba504d9f2c563406e64736f6c634300081c0033",
}

// IdentityUpdatesABI is the input ABI used to generate the binding from.
// Deprecated: Use IdentityUpdatesMetaData.ABI instead.
var IdentityUpdatesABI = IdentityUpdatesMetaData.ABI

// IdentityUpdatesBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use IdentityUpdatesMetaData.Bin instead.
var IdentityUpdatesBin = IdentityUpdatesMetaData.Bin

// DeployIdentityUpdates deploys a new Ethereum contract, binding an instance of IdentityUpdates to it.
func DeployIdentityUpdates(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *IdentityUpdates, error) {
	parsed, err := IdentityUpdatesMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(IdentityUpdatesBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IdentityUpdates{IdentityUpdatesCaller: IdentityUpdatesCaller{contract: contract}, IdentityUpdatesTransactor: IdentityUpdatesTransactor{contract: contract}, IdentityUpdatesFilterer: IdentityUpdatesFilterer{contract: contract}}, nil
}

// IdentityUpdates is an auto generated Go binding around an Ethereum contract.
type IdentityUpdates struct {
	IdentityUpdatesCaller     // Read-only binding to the contract
	IdentityUpdatesTransactor // Write-only binding to the contract
	IdentityUpdatesFilterer   // Log filterer for contract events
}

// IdentityUpdatesCaller is an auto generated read-only Go binding around an Ethereum contract.
type IdentityUpdatesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentityUpdatesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IdentityUpdatesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentityUpdatesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IdentityUpdatesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentityUpdatesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IdentityUpdatesSession struct {
	Contract     *IdentityUpdates  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IdentityUpdatesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IdentityUpdatesCallerSession struct {
	Contract *IdentityUpdatesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// IdentityUpdatesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IdentityUpdatesTransactorSession struct {
	Contract     *IdentityUpdatesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// IdentityUpdatesRaw is an auto generated low-level Go binding around an Ethereum contract.
type IdentityUpdatesRaw struct {
	Contract *IdentityUpdates // Generic contract binding to access the raw methods on
}

// IdentityUpdatesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IdentityUpdatesCallerRaw struct {
	Contract *IdentityUpdatesCaller // Generic read-only contract binding to access the raw methods on
}

// IdentityUpdatesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IdentityUpdatesTransactorRaw struct {
	Contract *IdentityUpdatesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIdentityUpdates creates a new instance of IdentityUpdates, bound to a specific deployed contract.
func NewIdentityUpdates(address common.Address, backend bind.ContractBackend) (*IdentityUpdates, error) {
	contract, err := bindIdentityUpdates(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IdentityUpdates{IdentityUpdatesCaller: IdentityUpdatesCaller{contract: contract}, IdentityUpdatesTransactor: IdentityUpdatesTransactor{contract: contract}, IdentityUpdatesFilterer: IdentityUpdatesFilterer{contract: contract}}, nil
}

// NewIdentityUpdatesCaller creates a new read-only instance of IdentityUpdates, bound to a specific deployed contract.
func NewIdentityUpdatesCaller(address common.Address, caller bind.ContractCaller) (*IdentityUpdatesCaller, error) {
	contract, err := bindIdentityUpdates(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IdentityUpdatesCaller{contract: contract}, nil
}

// NewIdentityUpdatesTransactor creates a new write-only instance of IdentityUpdates, bound to a specific deployed contract.
func NewIdentityUpdatesTransactor(address common.Address, transactor bind.ContractTransactor) (*IdentityUpdatesTransactor, error) {
	contract, err := bindIdentityUpdates(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IdentityUpdatesTransactor{contract: contract}, nil
}

// NewIdentityUpdatesFilterer creates a new log filterer instance of IdentityUpdates, bound to a specific deployed contract.
func NewIdentityUpdatesFilterer(address common.Address, filterer bind.ContractFilterer) (*IdentityUpdatesFilterer, error) {
	contract, err := bindIdentityUpdates(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IdentityUpdatesFilterer{contract: contract}, nil
}

// bindIdentityUpdates binds a generic wrapper to an already deployed contract.
func bindIdentityUpdates(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IdentityUpdatesMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IdentityUpdates *IdentityUpdatesRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IdentityUpdates.Contract.IdentityUpdatesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IdentityUpdates *IdentityUpdatesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IdentityUpdates.Contract.IdentityUpdatesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IdentityUpdates *IdentityUpdatesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IdentityUpdates.Contract.IdentityUpdatesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IdentityUpdates *IdentityUpdatesCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IdentityUpdates.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IdentityUpdates *IdentityUpdatesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IdentityUpdates.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IdentityUpdates *IdentityUpdatesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IdentityUpdates.Contract.contract.Transact(opts, method, params...)
}

// ABSOLUTEMAXPAYLOADSIZE is a free data retrieval call binding the contract method 0x0009e127.
//
// Solidity: function ABSOLUTE_MAX_PAYLOAD_SIZE() view returns(uint256)
func (_IdentityUpdates *IdentityUpdatesCaller) ABSOLUTEMAXPAYLOADSIZE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IdentityUpdates.contract.Call(opts, &out, "ABSOLUTE_MAX_PAYLOAD_SIZE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ABSOLUTEMAXPAYLOADSIZE is a free data retrieval call binding the contract method 0x0009e127.
//
// Solidity: function ABSOLUTE_MAX_PAYLOAD_SIZE() view returns(uint256)
func (_IdentityUpdates *IdentityUpdatesSession) ABSOLUTEMAXPAYLOADSIZE() (*big.Int, error) {
	return _IdentityUpdates.Contract.ABSOLUTEMAXPAYLOADSIZE(&_IdentityUpdates.CallOpts)
}

// ABSOLUTEMAXPAYLOADSIZE is a free data retrieval call binding the contract method 0x0009e127.
//
// Solidity: function ABSOLUTE_MAX_PAYLOAD_SIZE() view returns(uint256)
func (_IdentityUpdates *IdentityUpdatesCallerSession) ABSOLUTEMAXPAYLOADSIZE() (*big.Int, error) {
	return _IdentityUpdates.Contract.ABSOLUTEMAXPAYLOADSIZE(&_IdentityUpdates.CallOpts)
}

// ABSOLUTEMINPAYLOADSIZE is a free data retrieval call binding the contract method 0x1de01599.
//
// Solidity: function ABSOLUTE_MIN_PAYLOAD_SIZE() view returns(uint256)
func (_IdentityUpdates *IdentityUpdatesCaller) ABSOLUTEMINPAYLOADSIZE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IdentityUpdates.contract.Call(opts, &out, "ABSOLUTE_MIN_PAYLOAD_SIZE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ABSOLUTEMINPAYLOADSIZE is a free data retrieval call binding the contract method 0x1de01599.
//
// Solidity: function ABSOLUTE_MIN_PAYLOAD_SIZE() view returns(uint256)
func (_IdentityUpdates *IdentityUpdatesSession) ABSOLUTEMINPAYLOADSIZE() (*big.Int, error) {
	return _IdentityUpdates.Contract.ABSOLUTEMINPAYLOADSIZE(&_IdentityUpdates.CallOpts)
}

// ABSOLUTEMINPAYLOADSIZE is a free data retrieval call binding the contract method 0x1de01599.
//
// Solidity: function ABSOLUTE_MIN_PAYLOAD_SIZE() view returns(uint256)
func (_IdentityUpdates *IdentityUpdatesCallerSession) ABSOLUTEMINPAYLOADSIZE() (*big.Int, error) {
	return _IdentityUpdates.Contract.ABSOLUTEMINPAYLOADSIZE(&_IdentityUpdates.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_IdentityUpdates *IdentityUpdatesCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IdentityUpdates.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_IdentityUpdates *IdentityUpdatesSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _IdentityUpdates.Contract.DEFAULTADMINROLE(&_IdentityUpdates.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_IdentityUpdates *IdentityUpdatesCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _IdentityUpdates.Contract.DEFAULTADMINROLE(&_IdentityUpdates.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_IdentityUpdates *IdentityUpdatesCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IdentityUpdates.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_IdentityUpdates *IdentityUpdatesSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _IdentityUpdates.Contract.UPGRADEINTERFACEVERSION(&_IdentityUpdates.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_IdentityUpdates *IdentityUpdatesCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _IdentityUpdates.Contract.UPGRADEINTERFACEVERSION(&_IdentityUpdates.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_IdentityUpdates *IdentityUpdatesCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _IdentityUpdates.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_IdentityUpdates *IdentityUpdatesSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _IdentityUpdates.Contract.GetRoleAdmin(&_IdentityUpdates.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_IdentityUpdates *IdentityUpdatesCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _IdentityUpdates.Contract.GetRoleAdmin(&_IdentityUpdates.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_IdentityUpdates *IdentityUpdatesCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _IdentityUpdates.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_IdentityUpdates *IdentityUpdatesSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _IdentityUpdates.Contract.HasRole(&_IdentityUpdates.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_IdentityUpdates *IdentityUpdatesCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _IdentityUpdates.Contract.HasRole(&_IdentityUpdates.CallOpts, role, account)
}

// MaxPayloadSize is a free data retrieval call binding the contract method 0x58e3e94c.
//
// Solidity: function maxPayloadSize() view returns(uint256 size)
func (_IdentityUpdates *IdentityUpdatesCaller) MaxPayloadSize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IdentityUpdates.contract.Call(opts, &out, "maxPayloadSize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxPayloadSize is a free data retrieval call binding the contract method 0x58e3e94c.
//
// Solidity: function maxPayloadSize() view returns(uint256 size)
func (_IdentityUpdates *IdentityUpdatesSession) MaxPayloadSize() (*big.Int, error) {
	return _IdentityUpdates.Contract.MaxPayloadSize(&_IdentityUpdates.CallOpts)
}

// MaxPayloadSize is a free data retrieval call binding the contract method 0x58e3e94c.
//
// Solidity: function maxPayloadSize() view returns(uint256 size)
func (_IdentityUpdates *IdentityUpdatesCallerSession) MaxPayloadSize() (*big.Int, error) {
	return _IdentityUpdates.Contract.MaxPayloadSize(&_IdentityUpdates.CallOpts)
}

// MinPayloadSize is a free data retrieval call binding the contract method 0xf96927ac.
//
// Solidity: function minPayloadSize() view returns(uint256 size)
func (_IdentityUpdates *IdentityUpdatesCaller) MinPayloadSize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IdentityUpdates.contract.Call(opts, &out, "minPayloadSize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinPayloadSize is a free data retrieval call binding the contract method 0xf96927ac.
//
// Solidity: function minPayloadSize() view returns(uint256 size)
func (_IdentityUpdates *IdentityUpdatesSession) MinPayloadSize() (*big.Int, error) {
	return _IdentityUpdates.Contract.MinPayloadSize(&_IdentityUpdates.CallOpts)
}

// MinPayloadSize is a free data retrieval call binding the contract method 0xf96927ac.
//
// Solidity: function minPayloadSize() view returns(uint256 size)
func (_IdentityUpdates *IdentityUpdatesCallerSession) MinPayloadSize() (*big.Int, error) {
	return _IdentityUpdates.Contract.MinPayloadSize(&_IdentityUpdates.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_IdentityUpdates *IdentityUpdatesCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _IdentityUpdates.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_IdentityUpdates *IdentityUpdatesSession) Paused() (bool, error) {
	return _IdentityUpdates.Contract.Paused(&_IdentityUpdates.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_IdentityUpdates *IdentityUpdatesCallerSession) Paused() (bool, error) {
	return _IdentityUpdates.Contract.Paused(&_IdentityUpdates.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_IdentityUpdates *IdentityUpdatesCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IdentityUpdates.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_IdentityUpdates *IdentityUpdatesSession) ProxiableUUID() ([32]byte, error) {
	return _IdentityUpdates.Contract.ProxiableUUID(&_IdentityUpdates.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_IdentityUpdates *IdentityUpdatesCallerSession) ProxiableUUID() ([32]byte, error) {
	return _IdentityUpdates.Contract.ProxiableUUID(&_IdentityUpdates.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IdentityUpdates *IdentityUpdatesCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _IdentityUpdates.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IdentityUpdates *IdentityUpdatesSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IdentityUpdates.Contract.SupportsInterface(&_IdentityUpdates.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IdentityUpdates *IdentityUpdatesCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IdentityUpdates.Contract.SupportsInterface(&_IdentityUpdates.CallOpts, interfaceId)
}

// AddIdentityUpdate is a paid mutator transaction binding the contract method 0xba74fc7c.
//
// Solidity: function addIdentityUpdate(bytes32 inboxId, bytes update) returns()
func (_IdentityUpdates *IdentityUpdatesTransactor) AddIdentityUpdate(opts *bind.TransactOpts, inboxId [32]byte, update []byte) (*types.Transaction, error) {
	return _IdentityUpdates.contract.Transact(opts, "addIdentityUpdate", inboxId, update)
}

// AddIdentityUpdate is a paid mutator transaction binding the contract method 0xba74fc7c.
//
// Solidity: function addIdentityUpdate(bytes32 inboxId, bytes update) returns()
func (_IdentityUpdates *IdentityUpdatesSession) AddIdentityUpdate(inboxId [32]byte, update []byte) (*types.Transaction, error) {
	return _IdentityUpdates.Contract.AddIdentityUpdate(&_IdentityUpdates.TransactOpts, inboxId, update)
}

// AddIdentityUpdate is a paid mutator transaction binding the contract method 0xba74fc7c.
//
// Solidity: function addIdentityUpdate(bytes32 inboxId, bytes update) returns()
func (_IdentityUpdates *IdentityUpdatesTransactorSession) AddIdentityUpdate(inboxId [32]byte, update []byte) (*types.Transaction, error) {
	return _IdentityUpdates.Contract.AddIdentityUpdate(&_IdentityUpdates.TransactOpts, inboxId, update)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_IdentityUpdates *IdentityUpdatesTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IdentityUpdates.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_IdentityUpdates *IdentityUpdatesSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IdentityUpdates.Contract.GrantRole(&_IdentityUpdates.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_IdentityUpdates *IdentityUpdatesTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IdentityUpdates.Contract.GrantRole(&_IdentityUpdates.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address admin) returns()
func (_IdentityUpdates *IdentityUpdatesTransactor) Initialize(opts *bind.TransactOpts, admin common.Address) (*types.Transaction, error) {
	return _IdentityUpdates.contract.Transact(opts, "initialize", admin)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address admin) returns()
func (_IdentityUpdates *IdentityUpdatesSession) Initialize(admin common.Address) (*types.Transaction, error) {
	return _IdentityUpdates.Contract.Initialize(&_IdentityUpdates.TransactOpts, admin)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address admin) returns()
func (_IdentityUpdates *IdentityUpdatesTransactorSession) Initialize(admin common.Address) (*types.Transaction, error) {
	return _IdentityUpdates.Contract.Initialize(&_IdentityUpdates.TransactOpts, admin)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_IdentityUpdates *IdentityUpdatesTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IdentityUpdates.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_IdentityUpdates *IdentityUpdatesSession) Pause() (*types.Transaction, error) {
	return _IdentityUpdates.Contract.Pause(&_IdentityUpdates.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_IdentityUpdates *IdentityUpdatesTransactorSession) Pause() (*types.Transaction, error) {
	return _IdentityUpdates.Contract.Pause(&_IdentityUpdates.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_IdentityUpdates *IdentityUpdatesTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _IdentityUpdates.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_IdentityUpdates *IdentityUpdatesSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _IdentityUpdates.Contract.RenounceRole(&_IdentityUpdates.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_IdentityUpdates *IdentityUpdatesTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _IdentityUpdates.Contract.RenounceRole(&_IdentityUpdates.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_IdentityUpdates *IdentityUpdatesTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IdentityUpdates.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_IdentityUpdates *IdentityUpdatesSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IdentityUpdates.Contract.RevokeRole(&_IdentityUpdates.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_IdentityUpdates *IdentityUpdatesTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IdentityUpdates.Contract.RevokeRole(&_IdentityUpdates.TransactOpts, role, account)
}

// SetMaxPayloadSize is a paid mutator transaction binding the contract method 0xfe8e37a3.
//
// Solidity: function setMaxPayloadSize(uint256 maxPayloadSizeRequest) returns()
func (_IdentityUpdates *IdentityUpdatesTransactor) SetMaxPayloadSize(opts *bind.TransactOpts, maxPayloadSizeRequest *big.Int) (*types.Transaction, error) {
	return _IdentityUpdates.contract.Transact(opts, "setMaxPayloadSize", maxPayloadSizeRequest)
}

// SetMaxPayloadSize is a paid mutator transaction binding the contract method 0xfe8e37a3.
//
// Solidity: function setMaxPayloadSize(uint256 maxPayloadSizeRequest) returns()
func (_IdentityUpdates *IdentityUpdatesSession) SetMaxPayloadSize(maxPayloadSizeRequest *big.Int) (*types.Transaction, error) {
	return _IdentityUpdates.Contract.SetMaxPayloadSize(&_IdentityUpdates.TransactOpts, maxPayloadSizeRequest)
}

// SetMaxPayloadSize is a paid mutator transaction binding the contract method 0xfe8e37a3.
//
// Solidity: function setMaxPayloadSize(uint256 maxPayloadSizeRequest) returns()
func (_IdentityUpdates *IdentityUpdatesTransactorSession) SetMaxPayloadSize(maxPayloadSizeRequest *big.Int) (*types.Transaction, error) {
	return _IdentityUpdates.Contract.SetMaxPayloadSize(&_IdentityUpdates.TransactOpts, maxPayloadSizeRequest)
}

// SetMinPayloadSize is a paid mutator transaction binding the contract method 0x314a100e.
//
// Solidity: function setMinPayloadSize(uint256 minPayloadSizeRequest) returns()
func (_IdentityUpdates *IdentityUpdatesTransactor) SetMinPayloadSize(opts *bind.TransactOpts, minPayloadSizeRequest *big.Int) (*types.Transaction, error) {
	return _IdentityUpdates.contract.Transact(opts, "setMinPayloadSize", minPayloadSizeRequest)
}

// SetMinPayloadSize is a paid mutator transaction binding the contract method 0x314a100e.
//
// Solidity: function setMinPayloadSize(uint256 minPayloadSizeRequest) returns()
func (_IdentityUpdates *IdentityUpdatesSession) SetMinPayloadSize(minPayloadSizeRequest *big.Int) (*types.Transaction, error) {
	return _IdentityUpdates.Contract.SetMinPayloadSize(&_IdentityUpdates.TransactOpts, minPayloadSizeRequest)
}

// SetMinPayloadSize is a paid mutator transaction binding the contract method 0x314a100e.
//
// Solidity: function setMinPayloadSize(uint256 minPayloadSizeRequest) returns()
func (_IdentityUpdates *IdentityUpdatesTransactorSession) SetMinPayloadSize(minPayloadSizeRequest *big.Int) (*types.Transaction, error) {
	return _IdentityUpdates.Contract.SetMinPayloadSize(&_IdentityUpdates.TransactOpts, minPayloadSizeRequest)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_IdentityUpdates *IdentityUpdatesTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IdentityUpdates.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_IdentityUpdates *IdentityUpdatesSession) Unpause() (*types.Transaction, error) {
	return _IdentityUpdates.Contract.Unpause(&_IdentityUpdates.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_IdentityUpdates *IdentityUpdatesTransactorSession) Unpause() (*types.Transaction, error) {
	return _IdentityUpdates.Contract.Unpause(&_IdentityUpdates.TransactOpts)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_IdentityUpdates *IdentityUpdatesTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _IdentityUpdates.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_IdentityUpdates *IdentityUpdatesSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _IdentityUpdates.Contract.UpgradeToAndCall(&_IdentityUpdates.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_IdentityUpdates *IdentityUpdatesTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _IdentityUpdates.Contract.UpgradeToAndCall(&_IdentityUpdates.TransactOpts, newImplementation, data)
}

// IdentityUpdatesIdentityUpdateCreatedIterator is returned from FilterIdentityUpdateCreated and is used to iterate over the raw logs and unpacked data for IdentityUpdateCreated events raised by the IdentityUpdates contract.
type IdentityUpdatesIdentityUpdateCreatedIterator struct {
	Event *IdentityUpdatesIdentityUpdateCreated // Event containing the contract specifics and raw log

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
func (it *IdentityUpdatesIdentityUpdateCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityUpdatesIdentityUpdateCreated)
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
		it.Event = new(IdentityUpdatesIdentityUpdateCreated)
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
func (it *IdentityUpdatesIdentityUpdateCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityUpdatesIdentityUpdateCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityUpdatesIdentityUpdateCreated represents a IdentityUpdateCreated event raised by the IdentityUpdates contract.
type IdentityUpdatesIdentityUpdateCreated struct {
	InboxId    [32]byte
	Update     []byte
	SequenceId uint64
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterIdentityUpdateCreated is a free log retrieval operation binding the contract event 0xc1a40f292090ec0435e939cdfe248e0322a88566679a90a50c4e9e5ef762dbd5.
//
// Solidity: event IdentityUpdateCreated(bytes32 inboxId, bytes update, uint64 sequenceId)
func (_IdentityUpdates *IdentityUpdatesFilterer) FilterIdentityUpdateCreated(opts *bind.FilterOpts) (*IdentityUpdatesIdentityUpdateCreatedIterator, error) {

	logs, sub, err := _IdentityUpdates.contract.FilterLogs(opts, "IdentityUpdateCreated")
	if err != nil {
		return nil, err
	}
	return &IdentityUpdatesIdentityUpdateCreatedIterator{contract: _IdentityUpdates.contract, event: "IdentityUpdateCreated", logs: logs, sub: sub}, nil
}

// WatchIdentityUpdateCreated is a free log subscription operation binding the contract event 0xc1a40f292090ec0435e939cdfe248e0322a88566679a90a50c4e9e5ef762dbd5.
//
// Solidity: event IdentityUpdateCreated(bytes32 inboxId, bytes update, uint64 sequenceId)
func (_IdentityUpdates *IdentityUpdatesFilterer) WatchIdentityUpdateCreated(opts *bind.WatchOpts, sink chan<- *IdentityUpdatesIdentityUpdateCreated) (event.Subscription, error) {

	logs, sub, err := _IdentityUpdates.contract.WatchLogs(opts, "IdentityUpdateCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityUpdatesIdentityUpdateCreated)
				if err := _IdentityUpdates.contract.UnpackLog(event, "IdentityUpdateCreated", log); err != nil {
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

// ParseIdentityUpdateCreated is a log parse operation binding the contract event 0xc1a40f292090ec0435e939cdfe248e0322a88566679a90a50c4e9e5ef762dbd5.
//
// Solidity: event IdentityUpdateCreated(bytes32 inboxId, bytes update, uint64 sequenceId)
func (_IdentityUpdates *IdentityUpdatesFilterer) ParseIdentityUpdateCreated(log types.Log) (*IdentityUpdatesIdentityUpdateCreated, error) {
	event := new(IdentityUpdatesIdentityUpdateCreated)
	if err := _IdentityUpdates.contract.UnpackLog(event, "IdentityUpdateCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdentityUpdatesInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the IdentityUpdates contract.
type IdentityUpdatesInitializedIterator struct {
	Event *IdentityUpdatesInitialized // Event containing the contract specifics and raw log

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
func (it *IdentityUpdatesInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityUpdatesInitialized)
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
		it.Event = new(IdentityUpdatesInitialized)
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
func (it *IdentityUpdatesInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityUpdatesInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityUpdatesInitialized represents a Initialized event raised by the IdentityUpdates contract.
type IdentityUpdatesInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_IdentityUpdates *IdentityUpdatesFilterer) FilterInitialized(opts *bind.FilterOpts) (*IdentityUpdatesInitializedIterator, error) {

	logs, sub, err := _IdentityUpdates.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &IdentityUpdatesInitializedIterator{contract: _IdentityUpdates.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_IdentityUpdates *IdentityUpdatesFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *IdentityUpdatesInitialized) (event.Subscription, error) {

	logs, sub, err := _IdentityUpdates.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityUpdatesInitialized)
				if err := _IdentityUpdates.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_IdentityUpdates *IdentityUpdatesFilterer) ParseInitialized(log types.Log) (*IdentityUpdatesInitialized, error) {
	event := new(IdentityUpdatesInitialized)
	if err := _IdentityUpdates.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdentityUpdatesMaxPayloadSizeUpdatedIterator is returned from FilterMaxPayloadSizeUpdated and is used to iterate over the raw logs and unpacked data for MaxPayloadSizeUpdated events raised by the IdentityUpdates contract.
type IdentityUpdatesMaxPayloadSizeUpdatedIterator struct {
	Event *IdentityUpdatesMaxPayloadSizeUpdated // Event containing the contract specifics and raw log

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
func (it *IdentityUpdatesMaxPayloadSizeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityUpdatesMaxPayloadSizeUpdated)
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
		it.Event = new(IdentityUpdatesMaxPayloadSizeUpdated)
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
func (it *IdentityUpdatesMaxPayloadSizeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityUpdatesMaxPayloadSizeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityUpdatesMaxPayloadSizeUpdated represents a MaxPayloadSizeUpdated event raised by the IdentityUpdates contract.
type IdentityUpdatesMaxPayloadSizeUpdated struct {
	OldSize *big.Int
	NewSize *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMaxPayloadSizeUpdated is a free log retrieval operation binding the contract event 0xf59e99f8f54d2696b7cf184949ab2b4bbd6336ec1816b36f58ae9948d868fe90.
//
// Solidity: event MaxPayloadSizeUpdated(uint256 oldSize, uint256 newSize)
func (_IdentityUpdates *IdentityUpdatesFilterer) FilterMaxPayloadSizeUpdated(opts *bind.FilterOpts) (*IdentityUpdatesMaxPayloadSizeUpdatedIterator, error) {

	logs, sub, err := _IdentityUpdates.contract.FilterLogs(opts, "MaxPayloadSizeUpdated")
	if err != nil {
		return nil, err
	}
	return &IdentityUpdatesMaxPayloadSizeUpdatedIterator{contract: _IdentityUpdates.contract, event: "MaxPayloadSizeUpdated", logs: logs, sub: sub}, nil
}

// WatchMaxPayloadSizeUpdated is a free log subscription operation binding the contract event 0xf59e99f8f54d2696b7cf184949ab2b4bbd6336ec1816b36f58ae9948d868fe90.
//
// Solidity: event MaxPayloadSizeUpdated(uint256 oldSize, uint256 newSize)
func (_IdentityUpdates *IdentityUpdatesFilterer) WatchMaxPayloadSizeUpdated(opts *bind.WatchOpts, sink chan<- *IdentityUpdatesMaxPayloadSizeUpdated) (event.Subscription, error) {

	logs, sub, err := _IdentityUpdates.contract.WatchLogs(opts, "MaxPayloadSizeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityUpdatesMaxPayloadSizeUpdated)
				if err := _IdentityUpdates.contract.UnpackLog(event, "MaxPayloadSizeUpdated", log); err != nil {
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

// ParseMaxPayloadSizeUpdated is a log parse operation binding the contract event 0xf59e99f8f54d2696b7cf184949ab2b4bbd6336ec1816b36f58ae9948d868fe90.
//
// Solidity: event MaxPayloadSizeUpdated(uint256 oldSize, uint256 newSize)
func (_IdentityUpdates *IdentityUpdatesFilterer) ParseMaxPayloadSizeUpdated(log types.Log) (*IdentityUpdatesMaxPayloadSizeUpdated, error) {
	event := new(IdentityUpdatesMaxPayloadSizeUpdated)
	if err := _IdentityUpdates.contract.UnpackLog(event, "MaxPayloadSizeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdentityUpdatesMinPayloadSizeUpdatedIterator is returned from FilterMinPayloadSizeUpdated and is used to iterate over the raw logs and unpacked data for MinPayloadSizeUpdated events raised by the IdentityUpdates contract.
type IdentityUpdatesMinPayloadSizeUpdatedIterator struct {
	Event *IdentityUpdatesMinPayloadSizeUpdated // Event containing the contract specifics and raw log

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
func (it *IdentityUpdatesMinPayloadSizeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityUpdatesMinPayloadSizeUpdated)
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
		it.Event = new(IdentityUpdatesMinPayloadSizeUpdated)
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
func (it *IdentityUpdatesMinPayloadSizeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityUpdatesMinPayloadSizeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityUpdatesMinPayloadSizeUpdated represents a MinPayloadSizeUpdated event raised by the IdentityUpdates contract.
type IdentityUpdatesMinPayloadSizeUpdated struct {
	OldSize *big.Int
	NewSize *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMinPayloadSizeUpdated is a free log retrieval operation binding the contract event 0x1ee836faee0e7c61d20a079d0b5b4e1ee9c536e18268ef6f7c620dcec82f72cd.
//
// Solidity: event MinPayloadSizeUpdated(uint256 oldSize, uint256 newSize)
func (_IdentityUpdates *IdentityUpdatesFilterer) FilterMinPayloadSizeUpdated(opts *bind.FilterOpts) (*IdentityUpdatesMinPayloadSizeUpdatedIterator, error) {

	logs, sub, err := _IdentityUpdates.contract.FilterLogs(opts, "MinPayloadSizeUpdated")
	if err != nil {
		return nil, err
	}
	return &IdentityUpdatesMinPayloadSizeUpdatedIterator{contract: _IdentityUpdates.contract, event: "MinPayloadSizeUpdated", logs: logs, sub: sub}, nil
}

// WatchMinPayloadSizeUpdated is a free log subscription operation binding the contract event 0x1ee836faee0e7c61d20a079d0b5b4e1ee9c536e18268ef6f7c620dcec82f72cd.
//
// Solidity: event MinPayloadSizeUpdated(uint256 oldSize, uint256 newSize)
func (_IdentityUpdates *IdentityUpdatesFilterer) WatchMinPayloadSizeUpdated(opts *bind.WatchOpts, sink chan<- *IdentityUpdatesMinPayloadSizeUpdated) (event.Subscription, error) {

	logs, sub, err := _IdentityUpdates.contract.WatchLogs(opts, "MinPayloadSizeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityUpdatesMinPayloadSizeUpdated)
				if err := _IdentityUpdates.contract.UnpackLog(event, "MinPayloadSizeUpdated", log); err != nil {
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

// ParseMinPayloadSizeUpdated is a log parse operation binding the contract event 0x1ee836faee0e7c61d20a079d0b5b4e1ee9c536e18268ef6f7c620dcec82f72cd.
//
// Solidity: event MinPayloadSizeUpdated(uint256 oldSize, uint256 newSize)
func (_IdentityUpdates *IdentityUpdatesFilterer) ParseMinPayloadSizeUpdated(log types.Log) (*IdentityUpdatesMinPayloadSizeUpdated, error) {
	event := new(IdentityUpdatesMinPayloadSizeUpdated)
	if err := _IdentityUpdates.contract.UnpackLog(event, "MinPayloadSizeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdentityUpdatesPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the IdentityUpdates contract.
type IdentityUpdatesPausedIterator struct {
	Event *IdentityUpdatesPaused // Event containing the contract specifics and raw log

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
func (it *IdentityUpdatesPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityUpdatesPaused)
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
		it.Event = new(IdentityUpdatesPaused)
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
func (it *IdentityUpdatesPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityUpdatesPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityUpdatesPaused represents a Paused event raised by the IdentityUpdates contract.
type IdentityUpdatesPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_IdentityUpdates *IdentityUpdatesFilterer) FilterPaused(opts *bind.FilterOpts) (*IdentityUpdatesPausedIterator, error) {

	logs, sub, err := _IdentityUpdates.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &IdentityUpdatesPausedIterator{contract: _IdentityUpdates.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_IdentityUpdates *IdentityUpdatesFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *IdentityUpdatesPaused) (event.Subscription, error) {

	logs, sub, err := _IdentityUpdates.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityUpdatesPaused)
				if err := _IdentityUpdates.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_IdentityUpdates *IdentityUpdatesFilterer) ParsePaused(log types.Log) (*IdentityUpdatesPaused, error) {
	event := new(IdentityUpdatesPaused)
	if err := _IdentityUpdates.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdentityUpdatesRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the IdentityUpdates contract.
type IdentityUpdatesRoleAdminChangedIterator struct {
	Event *IdentityUpdatesRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *IdentityUpdatesRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityUpdatesRoleAdminChanged)
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
		it.Event = new(IdentityUpdatesRoleAdminChanged)
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
func (it *IdentityUpdatesRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityUpdatesRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityUpdatesRoleAdminChanged represents a RoleAdminChanged event raised by the IdentityUpdates contract.
type IdentityUpdatesRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_IdentityUpdates *IdentityUpdatesFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*IdentityUpdatesRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _IdentityUpdates.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &IdentityUpdatesRoleAdminChangedIterator{contract: _IdentityUpdates.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_IdentityUpdates *IdentityUpdatesFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *IdentityUpdatesRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _IdentityUpdates.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityUpdatesRoleAdminChanged)
				if err := _IdentityUpdates.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_IdentityUpdates *IdentityUpdatesFilterer) ParseRoleAdminChanged(log types.Log) (*IdentityUpdatesRoleAdminChanged, error) {
	event := new(IdentityUpdatesRoleAdminChanged)
	if err := _IdentityUpdates.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdentityUpdatesRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the IdentityUpdates contract.
type IdentityUpdatesRoleGrantedIterator struct {
	Event *IdentityUpdatesRoleGranted // Event containing the contract specifics and raw log

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
func (it *IdentityUpdatesRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityUpdatesRoleGranted)
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
		it.Event = new(IdentityUpdatesRoleGranted)
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
func (it *IdentityUpdatesRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityUpdatesRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityUpdatesRoleGranted represents a RoleGranted event raised by the IdentityUpdates contract.
type IdentityUpdatesRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_IdentityUpdates *IdentityUpdatesFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*IdentityUpdatesRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IdentityUpdates.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &IdentityUpdatesRoleGrantedIterator{contract: _IdentityUpdates.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_IdentityUpdates *IdentityUpdatesFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *IdentityUpdatesRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IdentityUpdates.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityUpdatesRoleGranted)
				if err := _IdentityUpdates.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_IdentityUpdates *IdentityUpdatesFilterer) ParseRoleGranted(log types.Log) (*IdentityUpdatesRoleGranted, error) {
	event := new(IdentityUpdatesRoleGranted)
	if err := _IdentityUpdates.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdentityUpdatesRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the IdentityUpdates contract.
type IdentityUpdatesRoleRevokedIterator struct {
	Event *IdentityUpdatesRoleRevoked // Event containing the contract specifics and raw log

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
func (it *IdentityUpdatesRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityUpdatesRoleRevoked)
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
		it.Event = new(IdentityUpdatesRoleRevoked)
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
func (it *IdentityUpdatesRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityUpdatesRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityUpdatesRoleRevoked represents a RoleRevoked event raised by the IdentityUpdates contract.
type IdentityUpdatesRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_IdentityUpdates *IdentityUpdatesFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*IdentityUpdatesRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IdentityUpdates.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &IdentityUpdatesRoleRevokedIterator{contract: _IdentityUpdates.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_IdentityUpdates *IdentityUpdatesFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *IdentityUpdatesRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IdentityUpdates.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityUpdatesRoleRevoked)
				if err := _IdentityUpdates.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_IdentityUpdates *IdentityUpdatesFilterer) ParseRoleRevoked(log types.Log) (*IdentityUpdatesRoleRevoked, error) {
	event := new(IdentityUpdatesRoleRevoked)
	if err := _IdentityUpdates.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdentityUpdatesUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the IdentityUpdates contract.
type IdentityUpdatesUnpausedIterator struct {
	Event *IdentityUpdatesUnpaused // Event containing the contract specifics and raw log

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
func (it *IdentityUpdatesUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityUpdatesUnpaused)
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
		it.Event = new(IdentityUpdatesUnpaused)
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
func (it *IdentityUpdatesUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityUpdatesUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityUpdatesUnpaused represents a Unpaused event raised by the IdentityUpdates contract.
type IdentityUpdatesUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_IdentityUpdates *IdentityUpdatesFilterer) FilterUnpaused(opts *bind.FilterOpts) (*IdentityUpdatesUnpausedIterator, error) {

	logs, sub, err := _IdentityUpdates.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &IdentityUpdatesUnpausedIterator{contract: _IdentityUpdates.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_IdentityUpdates *IdentityUpdatesFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *IdentityUpdatesUnpaused) (event.Subscription, error) {

	logs, sub, err := _IdentityUpdates.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityUpdatesUnpaused)
				if err := _IdentityUpdates.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_IdentityUpdates *IdentityUpdatesFilterer) ParseUnpaused(log types.Log) (*IdentityUpdatesUnpaused, error) {
	event := new(IdentityUpdatesUnpaused)
	if err := _IdentityUpdates.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdentityUpdatesUpgradeAuthorizedIterator is returned from FilterUpgradeAuthorized and is used to iterate over the raw logs and unpacked data for UpgradeAuthorized events raised by the IdentityUpdates contract.
type IdentityUpdatesUpgradeAuthorizedIterator struct {
	Event *IdentityUpdatesUpgradeAuthorized // Event containing the contract specifics and raw log

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
func (it *IdentityUpdatesUpgradeAuthorizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityUpdatesUpgradeAuthorized)
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
		it.Event = new(IdentityUpdatesUpgradeAuthorized)
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
func (it *IdentityUpdatesUpgradeAuthorizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityUpdatesUpgradeAuthorizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityUpdatesUpgradeAuthorized represents a UpgradeAuthorized event raised by the IdentityUpdates contract.
type IdentityUpdatesUpgradeAuthorized struct {
	Upgrader          common.Address
	NewImplementation common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterUpgradeAuthorized is a free log retrieval operation binding the contract event 0xd30e1d298bf814ea43d22b4ce8298062b08609cd67496483769d836157dd52fa.
//
// Solidity: event UpgradeAuthorized(address upgrader, address newImplementation)
func (_IdentityUpdates *IdentityUpdatesFilterer) FilterUpgradeAuthorized(opts *bind.FilterOpts) (*IdentityUpdatesUpgradeAuthorizedIterator, error) {

	logs, sub, err := _IdentityUpdates.contract.FilterLogs(opts, "UpgradeAuthorized")
	if err != nil {
		return nil, err
	}
	return &IdentityUpdatesUpgradeAuthorizedIterator{contract: _IdentityUpdates.contract, event: "UpgradeAuthorized", logs: logs, sub: sub}, nil
}

// WatchUpgradeAuthorized is a free log subscription operation binding the contract event 0xd30e1d298bf814ea43d22b4ce8298062b08609cd67496483769d836157dd52fa.
//
// Solidity: event UpgradeAuthorized(address upgrader, address newImplementation)
func (_IdentityUpdates *IdentityUpdatesFilterer) WatchUpgradeAuthorized(opts *bind.WatchOpts, sink chan<- *IdentityUpdatesUpgradeAuthorized) (event.Subscription, error) {

	logs, sub, err := _IdentityUpdates.contract.WatchLogs(opts, "UpgradeAuthorized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityUpdatesUpgradeAuthorized)
				if err := _IdentityUpdates.contract.UnpackLog(event, "UpgradeAuthorized", log); err != nil {
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

// ParseUpgradeAuthorized is a log parse operation binding the contract event 0xd30e1d298bf814ea43d22b4ce8298062b08609cd67496483769d836157dd52fa.
//
// Solidity: event UpgradeAuthorized(address upgrader, address newImplementation)
func (_IdentityUpdates *IdentityUpdatesFilterer) ParseUpgradeAuthorized(log types.Log) (*IdentityUpdatesUpgradeAuthorized, error) {
	event := new(IdentityUpdatesUpgradeAuthorized)
	if err := _IdentityUpdates.contract.UnpackLog(event, "UpgradeAuthorized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdentityUpdatesUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the IdentityUpdates contract.
type IdentityUpdatesUpgradedIterator struct {
	Event *IdentityUpdatesUpgraded // Event containing the contract specifics and raw log

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
func (it *IdentityUpdatesUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityUpdatesUpgraded)
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
		it.Event = new(IdentityUpdatesUpgraded)
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
func (it *IdentityUpdatesUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityUpdatesUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityUpdatesUpgraded represents a Upgraded event raised by the IdentityUpdates contract.
type IdentityUpdatesUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_IdentityUpdates *IdentityUpdatesFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*IdentityUpdatesUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _IdentityUpdates.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &IdentityUpdatesUpgradedIterator{contract: _IdentityUpdates.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_IdentityUpdates *IdentityUpdatesFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *IdentityUpdatesUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _IdentityUpdates.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityUpdatesUpgraded)
				if err := _IdentityUpdates.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_IdentityUpdates *IdentityUpdatesFilterer) ParseUpgraded(log types.Log) (*IdentityUpdatesUpgraded, error) {
	event := new(IdentityUpdatesUpgraded)
	if err := _IdentityUpdates.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

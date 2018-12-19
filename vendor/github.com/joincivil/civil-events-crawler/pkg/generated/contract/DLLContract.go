// Code generated by 'gen/libgen'  DO NOT EDIT.
// IT SHOULD NOT BE EDITED BY HAND AS ANY CHANGES MAY BE OVERWRITTEN
// Please reference 'gen/libgen' for more details
// File was generated at 2018-11-06 00:06:57.664792 +0000 UTC
package contract

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

const DLLContractABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"self\",\"type\":\"string\"}],\"name\":\"isEmpty\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"self\",\"type\":\"string\"},{\"name\":\"_curr\",\"type\":\"uint256\"}],\"name\":\"contains\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"self\",\"type\":\"string\"},{\"name\":\"_curr\",\"type\":\"uint256\"}],\"name\":\"getNext\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"self\",\"type\":\"string\"},{\"name\":\"_curr\",\"type\":\"uint256\"}],\"name\":\"getPrev\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"self\",\"type\":\"string\"}],\"name\":\"getStart\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"self\",\"type\":\"string\"}],\"name\":\"getEnd\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"self\",\"type\":\"string\"},{\"name\":\"_prev\",\"type\":\"uint256\"},{\"name\":\"_curr\",\"type\":\"uint256\"},{\"name\":\"_next\",\"type\":\"uint256\"}],\"name\":\"insert\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"self\",\"type\":\"string\"},{\"name\":\"_curr\",\"type\":\"uint256\"}],\"name\":\"remove\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

const DLLContractBin = `0x610369610030600b82828239805160001a6073146000811461002057610022565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600436106100a45763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166307d29ac981146100a957806330fe0a0a146100c9578063366a5ba2146100d75780636d900ed0146100f95780637c11cf64146101165780639735c51b14610121578063c426b00614610142578063ee4f1ac41461014d575b600080fd5b6100b7600435602435610158565b60408051918252519081900360200190f35b6100b760043560243561016b565b6100e5600435602435610181565b604080519115158252519081900360200190f35b81801561010557600080fd5b506101146004356024356101ff565b005b6100b7600435610261565b81801561012d57600080fd5b50610114600435602435604435606435610274565b6100e560043561031d565b6100b7600435610330565b6000908152602091909152604090205490565b6000908152602091909152604090206001015490565b600080600061018f8561031d565b80610198575083155b156101a657600092506101f7565b836101b086610330565b1480156101c45750836101c286610261565b145b915060006101d28686610158565b1480156101e8575060006101e6868661016b565b145b905081806101f4575080155b92505b505092915050565b60008061020c8484610181565b15156102175761025b565b6102218484610158565b915061022d848461016b565b6000838152602086905260408082206001908101849055838352818320869055868352908220828155015590505b50505050565b600061026e82600061016b565b92915050565b81151561028057600080fd5b61028a84836101ff565b82158061029c575061029c8484610181565b15156102a757600080fd5b8015806102b957506102b98482610181565b15156102c457600080fd5b806102cf8585610158565b146102d957600080fd5b826102e4858361016b565b146102ee57600080fd5b600082815260209490945260408085206001808201869055908390559385528085208390559084529092200155565b60008061032983610330565b1492915050565b600061026e8260006101585600a165627a7a723058203247076ce09af43d593d872c28d2d9395139cf10e913560f6243c6bb114939da0029`

func DeployDLLContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DLLContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DLLContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, contract, nil
}

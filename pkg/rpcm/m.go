package rpcm

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

type Network string

var Network_local = Network("local")
var Network_polygon = Network("polygon")
var Network_mumbai = Network("mumbai")

var rpcMap = map[Network]*ethclient.Client{}

var _init bool = false

func Init() error {
	if _init {
		return nil
	}

	rpc, err := rpc.DialWebsocket(context.Background(), "ws://polygon-fast2-bor.web3-system:8546", "")

	if err != nil {
		return err
	}

	rpcMap[Network_polygon] = ethclient.NewClient(rpc)

	println("rpcm: polygon network initialized")
	_init = true
	return nil
}

func GetEthClient(network Network) (*ethclient.Client, error) {
	if _init {
		return rpcMap[network], nil
	}

	return nil, fmt.Errorf("rpcm: not initialized")
}

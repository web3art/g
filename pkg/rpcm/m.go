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
var Network_ethereum = Network("ethereum")

var rpcMap = map[Network]*ethclient.Client{}

var _init bool = false

func initPolygon() error {
	rpc, err := rpc.DialWebsocket(context.Background(), "ws://polygon-fast2-bor.web3-system:8546", "")

	if err != nil {
		return err
	}

	rpcMap[Network_polygon] = ethclient.NewClient(rpc)
	return nil
}

func initMumbai() error {
	rpc, err := rpc.DialWebsocket(context.Background(), "ws://polygon-mumbai-fast-bor.web3-system:8546", "")

	if err != nil {
		return err
	}

	rpcMap[Network_mumbai] = ethclient.NewClient(rpc)
	return nil
}

func initEthereum() error {
	rpc, err := rpc.DialWebsocket(context.Background(), "wss://speedy-nodes-nyc.moralis.io/036db9847a8d6409b3dbcddc/eth/mainnet/ws", "")

	if err != nil {
		return err
	}

	rpcMap[Network_ethereum] = ethclient.NewClient(rpc)
	return nil
}

func Init() error {
	if _init {
		return nil
	}

	if err := initPolygon(); err != nil {
		return err
	}

	if err := initMumbai(); err != nil {
		return err
	}

	if err := initEthereum(); err != nil {
		return err
	}

	println("rpcm: network initialized")
	_init = true
	return nil
}

func GetEthClient(network Network) (*ethclient.Client, error) {
	if _init {
		return rpcMap[network], nil
	}

	return nil, fmt.Errorf("rpcm: not initialized")
}

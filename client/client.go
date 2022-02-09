package client

import (
	"context"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"math/big"
	"strings"

	"github.com/cosmos/cosmos-sdk/codec"
)


type rpcClient struct {
	*rpc.Client
}

type ethClient struct {
	*ethclient.Client
	*rpcClient
}
func NewEthClient(ctx context.Context, rawurl string) (*ethClient, error) {
	c, err := rpc.DialContext(ctx, rawurl)
	if err != nil {
		return nil, err
	}
	return &ethClient{ethclient.NewClient(c), &rpcClient{c}}, nil
}

// Client - structure of the main client of ExChain GoSDK
type Client struct {

	cdc     *codec.Codec

}

var EthClient *ethClient


func InitClient(ip string)  {
	sdk.GetConfig().SetBech32PrefixForAccount("de", "depub")
	EthClient,_=NewEthClient(context.Background(),"http://"+ip+":8545")
}


func(cli *ethClient) GetContractBalance(tokenAddress string, address string) (*big.Int,error) {
	tokenAddressHash := common.HexToAddress(tokenAddress)
	// 生成交易
	contractAbi, err := abi.JSON(strings.NewReader(EthABI))
	if err != nil {
		return nil, err
	}
	input, err := contractAbi.Pack(
		"balanceOf",
		common.HexToAddress(address),
	)
	if err != nil {
		return nil, err
	}
	msg := ethereum.CallMsg{
		From:  common.HexToAddress(address),
		To:    &tokenAddressHash,
		Value: nil,
		Data:  input,
	}
	out, err := EthClient.CallContract(context.Background(), msg, nil)
	if err != nil {
		return nil, err
	}
	res, err := contractAbi.Unpack("balanceOf", out)
	if err != nil {
		return nil, err
	}
	if len(res) != 1 {
		return nil, fmt.Errorf("error call res")
	}
	out0, ok := res[0].(*big.Int)
	if !ok {
		return nil, fmt.Errorf("error call res")
	}
	return out0, nil

}

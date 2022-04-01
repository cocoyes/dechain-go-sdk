package face

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
	"math/big"
	"testing"
)

func TestBuildTrans(t *testing.T) {
	// 一、ABI编码请求参数
	methodId := crypto.Keccak256([]byte("setA(uint256)"))[:4]
	fmt.Println("methodId: ", common.Bytes2Hex(methodId))
	paramValue := math.U256Bytes(new(big.Int).Set(big.NewInt(123)))
	fmt.Println("paramValue: ", common.Bytes2Hex(paramValue))
	input := append(methodId, paramValue...)
	fmt.Println("input: ", common.Bytes2Hex(input))

	// 二、构造交易对象
	nonce := uint64(24)
	value := big.NewInt(0)
	gasLimit := uint64(3000000)
	gasPrice := big.NewInt(20000000000)
	rawTx := types.NewTransaction(nonce, common.HexToAddress("0x05e56888360ae54acf2a389bab39bd41e3934d2b"), value, gasLimit, gasPrice, input)
	jsonRawTx, _ := rawTx.MarshalJSON()
	fmt.Println("rawTx: ", string(jsonRawTx))

	// 三、交易签名
	signer := types.NewEIP155Signer(big.NewInt(1))
	key, err := crypto.HexToECDSA("e8e14120bb5c085622253540e886527d24746cd42d764a5974be47090d3cbc42")
	if err != nil {
		fmt.Println("crypto.HexToECDSA failed: ", err.Error())
		return
	}
	sigTransaction, err := types.SignTx(rawTx, signer, key)
	if err != nil {
		fmt.Println("types.SignTx failed: ", err.Error())
		return
	}
	jsonSigTx, _ := sigTransaction.MarshalJSON()
	fmt.Println("sigTransaction: ", string(jsonSigTx))

	// 四、发送交易
	ethClient, err := ethclient.Dial("http://127.0.0.1:7545")
	if err != nil {
		fmt.Println("ethclient.Dial failed: ", err.Error())
		return
	}
	err = ethClient.SendTransaction(context.Background(), sigTransaction)
	if err != nil {
		fmt.Println("ethClient.SendTransaction failed: ", err.Error())
		return
	}
	fmt.Println("send transaction success,tx: ", sigTransaction.Hash().Hex())
}

func TestDecode(t *testing.T) {
	// 还原交易对象
	encodedTxStr := "0xf889188504a817c800832dc6c09405e56888360ae54acf2a389bab39bd41e3934d2b80a4ee919d50000000000000000000000000000000000000000000000000000000000000007b25a041c4a2eb073e6df89c3f467b3516e9c313590d8d57f7c217fe7e72a7b4a6b8eda05f20a758396a5e681ce1ab4cec749f8560e28c9eb91072ec7a8acc002a11bb1d"
	encodedTx, err := hexutil.Decode(encodedTxStr)
	if err != nil {
		fmt.Println("hexutil.Decode failed: ", err.Error())
		return
	}
	// rlp解码
	tx := new(types.Transaction)
	if err := rlp.DecodeBytes(encodedTx, tx); err != nil {
		fmt.Println("rlp.DecodeBytes failed: ", err.Error())
		return
	}
	// chainId为1的EIP155签名器
	signer := types.NewEIP155Signer(big.NewInt(1))
	// 使用签名器从已签名的交易中还原账户公钥
	from, err := types.Sender(signer, tx)
	if err != nil {
		fmt.Println("types.Sender: ", err.Error())
		return
	}
	fmt.Println("from: ", from.Hex())
	jsonTx, _ := tx.MarshalJSON()
	fmt.Println("tx: ", string(jsonTx))
}

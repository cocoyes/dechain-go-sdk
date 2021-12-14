package java


import (
	"bytes"
	"context"
	"crypto/md5"
	"dechain-go-sdk/client"
	"dechain-go-sdk/utils"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/shopspring/decimal"
	"math/big"
	"strings"
)

func main() {

}



//初始化链接
func CreateClient(ip string) string {
	client.InitClient(ip)
	return "success"
}


//生成助记词、私钥、地址
func GenKey() string {
	mnWallet,_:= utils.CreateAccount("","")
	st := struct {
		Prikey     string `json:"prikey"`
		Mnemo 	   string  `json:"mnemo"`
		Address     string `json:"address"`
		HexAddr	 	string `json:"hexAddr"`
	}{
		Prikey:     mnWallet.Prikey,
		Mnemo: 		mnWallet.Mnemo,
		Address:    mnWallet.Address,
		HexAddr:    mnWallet.HexAddress,
	}
	return utils.ToJson(st)
}

//导入助记，返回私钥、地址
func ImportMnemo(mnemo string) string {
	mn:=mnemo
	mnWallet,_:=utils.ImportWallet(mn)
	st := struct {
		Prikey     string `json:"prikey"`
		Mnemo 	   string  `json:"mnemo"`
		Address     string `json:"address"`
		HexAddr	 	string `json:"hexAddr"`
	}{
		Prikey:     mnWallet.Prikey,
		Mnemo: 		mnWallet.Mnemo,
		Address:    mnWallet.Address,
		HexAddr:    mnWallet.HexAddress,
	}
	return utils.ToJson(st)
}

// 普通地址地址转0x地址
func AddrToHex(address string) string {
	addr:=address
	commonAddress,_:=utils.ToHexAddress(addr)
	return commonAddress.String()
}

//0x地址转普通地址
func HexToAddr(hexAddr string) string {
	hx:=hexAddr
	addr,_:=utils.ToCosmosAddress(hx)
	return addr.String()
}

// 主币余额（活力值余额）
func BalanceMain(address string) string {
	addr:=address
	if !strings.HasPrefix(addr,"0x") {
		addr=utils.AddrToHex(addr)
	}
	ba,err:=client.EthClient.BalanceAt(context.Background(),common.HexToAddress(addr),nil)
	if err!=nil {
		return "0"
	}else {
		return ba.String()
	}
}

// 代币余额（通证余额）
func BalanceContract(contractAddr string, address string) string {
	caddr:=contractAddr
	addr:=address
	if !strings.HasPrefix(addr,"0x") {
		addr=utils.AddrToHex(addr)
	}
	if !strings.HasPrefix(caddr,"0x") {
		caddr=utils.AddrToHex(caddr)
	}
	ba,err:=client.EthClient.GetContractBalance(caddr,addr)
	if err!=nil {
		return "0"
	}else {
		return ba.String()
	}
}


// 主币转账（活力值转账）
func TransferMain(priKey string,toAddress string,amount string) string {
	taddr:=toAddress
	pri:=priKey
	amo:=amount

	res := struct {
		Code     int `json:"code"`
		Message 	   string  `json:"message"`
		Data     string `json:"data"`
	}{
		Code:     0,
		Message: 		"success",
		Data:    "",
	}
	number:=big.NewInt(1000000000000000000)
	if !strings.HasPrefix(taddr,"0x") {
		taddr=utils.AddrToHex(taddr)
	}
	cryKey, _  := crypto.HexToECDSA(pri)
	chainID, _ := client.EthClient.NetworkID(context.Background())
	fromAddr:=crypto.PubkeyToAddress(cryKey.PublicKey)

	penNonce, err := client.EthClient.PendingNonceAt(context.Background(), fromAddr)

	gasPrice, err := client.EthClient.SuggestGasPrice(context.Background())
	// Create transaction
	value,_:=new(big.Int).SetString(amo,10)
	tx := types.NewTransaction(penNonce, common.HexToAddress(taddr),number.Mul(number,value) , client.EthMaxGasLimit, gasPrice, nil)
	signer := types.LatestSignerForChainID(chainID)
	signature, err := crypto.Sign(signer.Hash(tx).Bytes(), cryKey)
	if err != nil {
		res.Code=1
		res.Message="key sign error"
		return utils.ToJson(res)
	}
	signedTx, err := tx.WithSignature(signer, signature)
	if err != nil {
		res.Code=1
		res.Message="tx sign error"
		return utils.ToJson(res)
	}
	// Send transaction
	err = client.EthClient.SendTransaction(context.Background(), signedTx)
	if err != nil {
		res.Code=1
		res.Message="SendTransaction error"
		return utils.ToJson(res)
	}else {
		hash:=signedTx.Hash()
		res.Data=hash.Hex()
		_, err := client.EthClient.TransactionReceipt(context.Background(), hash)
		if err!=nil {
			res.Code=1
			res.Message=err.Error()
		}else {
			res.Code=0
			res.Message="success"
		}
		return utils.ToJson(res)
	}
}

//代币转账（通证转账）contractAddr通证地址、amount数量对应单位长度
func TransferContract(priKey string,contractAddr string,toAddress string,amount string) string {
	taddr:=toAddress
	pri:=priKey
	amo:=amount
	caddr:=contractAddr
	res := struct {
		Code     int `json:"code"`
		Message 	   string  `json:"message"`
		Data     string `json:"data"`
	}{
		Code:     0,
		Message: 		"success",
		Data:    "",
	}

	if !strings.HasPrefix(taddr,"0x") {
		taddr=utils.AddrToHex(taddr)
	}
	if !strings.HasPrefix(caddr,"0x") {
		caddr=utils.AddrToHex(caddr)
	}
	cryKey, _  := crypto.HexToECDSA(pri)
	chainID, _ := client.EthClient.NetworkID(context.Background())
	fromAddr:=crypto.PubkeyToAddress(cryKey.PublicKey)

	penNonce, err := client.EthClient.PendingNonceAt(context.Background(), fromAddr)


	// Create transaction
	value,_:=new(big.Int).SetString(amo,10)
	paddedAddress := common.LeftPadBytes(common.HexToAddress(taddr).Bytes(), 32)
	paddedAmount := common.LeftPadBytes(value.Bytes(), 32)
	transferFnSignature := []byte("transfer(address,uint256)")
	methodID := crypto.Keccak256(transferFnSignature)[:4]
	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)
	tx := types.NewTransaction(penNonce, common.HexToAddress(caddr), big.NewInt(0), client.ContractGasLimit, big.NewInt(client.GasPrice), data)
	signer := types.LatestSignerForChainID(chainID)
	signature, err := crypto.Sign(signer.Hash(tx).Bytes(), cryKey)
	if err != nil {
		res.Code=1
		res.Message="key sign error"
		return utils.ToJson(res)
	}
	signedTx, err := tx.WithSignature(signer, signature)
	if err != nil {
		res.Code=1
		res.Message="tx sign error"
		return utils.ToJson(res)
	}
	// Send transaction
	err = client.EthClient.SendTransaction(context.Background(), signedTx)
	if err != nil {
		res.Code=1
		res.Message="SendTransaction error"
		return utils.ToJson(res)
	}else {
		hash:=signedTx.Hash()
		res.Data=hash.Hex()
		_, err := client.EthClient.TransactionReceipt(context.Background(), hash)
		if err!=nil {
			res.Code=1
			res.Message=err.Error()
		}else {
			res.Code=0
			res.Message="success"
		}
		return utils.ToJson(res)
	}
}

//计算手续费
func CalFee() string {
	res := struct {
		Common     string `json:"common"`
		Contract 	   string  `json:"contract"`
		GasPrice   string `json:"gasPrice"`
	}{
		Common:     "",
		Contract: 		"",
		GasPrice:	"",
	}
	gs:=decimal.NewFromInt(client.GasPrice)
	base:=decimal.NewFromInt(1).Mul(decimal.NewFromInt(10).Pow(decimal.NewFromInt(18)))
	cheng1:=gs.Mul(decimal.NewFromInt(client.EthMaxGasLimit)).DivRound(base,18)
	cheng2:=gs.Mul(decimal.NewFromInt(client.ContractGasLimit)).DivRound(base,18)

	// 20 * 21000 /x = 0.00042
	res.Common=cheng1.String()
	res.Contract=cheng2.String()
	return utils.ToJson(res)
}


//检查交易状态  参数_hash为交易流水  返回值status 为1才表示成功，0则表示未成功（或即将成功），通常间隔一秒调用此方法10次后，可认为失败
func CheckStatus(_hash string) string {
	hash:=_hash
	receipt, err := client.EthClient.TransactionReceipt(context.Background(), common.HexToHash(hash))
	if err != nil {
		return err.Error()
	}
	return utils.ToJson(receipt)
}



// --------------------------------------- 智能合约操作部分 -------------------------------------------------------

//授权 （此处的红包（或支付平台）的智能合约跟 原本的代币（通证）合约是两种不同的合约，红包合约负责处理红包业务逻辑，代币（通证）合约 负责记账，因此需要向代币（通证）合约授权指定的操作金额）  其中_coinContract表示代币（通证）的合约地址，_approveContract表示授权给谁去操作，这里通常指授权给 红包/支付 合约去操作你的代币（通证） 该接口需消耗手续费，因此会返回交易hash，调用checkStatus函数可查询是否授权成功，手续费不足会失败
func Approve(_priKey string,_coinContract string,_approveContract string,_amount string) string {
	coinContract:=_coinContract
	pri:=_priKey
	amo:=_amount
	approveContract:=_approveContract
	res := struct {
		Code     int `json:"code"`
		Message 	   string  `json:"message"`
		Data     string `json:"data"`
	}{
		Code:     0,
		Message: 		"success",
		Data:    "",
	}

	inputParams:=make([]string,2)
	inputParams[0]=approveContract
	inputParams[1]=amo
	signedTx, err := utils.CallContractMethod(pri,coinContract,inputParams,"approve",client.EthABI)
	err = client.EthClient.SendTransaction(context.Background(), signedTx)
	if err != nil {
		res.Code=1
		res.Message="SendTransaction error"
		return utils.ToJson(res)
	}else {
		hash:=signedTx.Hash()
		res.Data=hash.Hex()
		_, err := client.EthClient.TransactionReceipt(context.Background(), hash)
		if err!=nil {
			res.Code=1
			res.Message=err.Error()
		}else {
			res.Code=0
			res.Message="success"
		}
		return utils.ToJson(res)
	}
}


//创建红包 授权成功后可进行红包创建，注意，若是没有授权就进行红包创建，交易hash会返回，但是一直不会成功，是失败的 _redContract是红包合约，红包合约已自动绑定上通证合约 _redId 红包id，需自行本地生成随机数，建议不超过15位，若跟之前的重复，会失败 _count 红包数量，必须大于1 ，_amount红包发送的总额 该接口需消耗手续费，因此会返回交易hash，调用checkStatus函数可查询是否授权成功，手续费不足会失败
func CreateRedPack(_priKey string,_redContract string,_amount string,_count string,_redId string) string {
	pri:=_priKey
	amo:=_amount
	redContract:=_redContract
	count:=_count
	redId:=_redId

	res := struct {
		Code     int `json:"code"`
		Message 	   string  `json:"message"`
		Data     string `json:"data"`
	}{
		Code:     0,
		Message: 		"success",
		Data:    "",
	}


	inputParams:=make([]string,3)
	inputParams[0]=count
	inputParams[1]=amo
	inputParams[2]=redId
	signedTx, err := utils.CallContractMethod(pri,redContract,inputParams,"toll",client.ReDPackAbi)

	err = client.EthClient.SendTransaction(context.Background(), signedTx)
	if err != nil {
		res.Code=1
		res.Message="SendTransaction error"
		return utils.ToJson(res)
	}else {
		hash:=signedTx.Hash()
		res.Data=hash.Hex()
		_, err := client.EthClient.TransactionReceipt(context.Background(), hash)
		if err!=nil {
			res.Code=1
			res.Message=err.Error()
		}else {
			res.Code=0
			res.Message="success"
		}
		return utils.ToJson(res)
	}
}

//抢红包 _redId 红包的id _redContract是红包合约，红包合约已自动绑定上通证合约 该接口需消耗手续费，因此会返回交易hash，调用checkStatus函数可查询是否授权成功，手续费不足会失败
func HuntingRedPack(_priKey string,_redContract string,_redId string) string {
	pri:=_priKey
	redContract:=_redContract
	redId:=_redId

	res := struct {
		Code     int `json:"code"`
		Message 	   string  `json:"message"`
		Data     string `json:"data"`
	}{
		Code:     0,
		Message: 		"success",
		Data:    "",
	}



	inputParams:=make([]string,1)
	inputParams[0]=redId
	signedTx, err := utils.CallContractMethod(pri,redContract,inputParams,"hunting",client.ReDPackAbi)

	err = client.EthClient.SendTransaction(context.Background(), signedTx)
	if err != nil {
		res.Code=1
		res.Message="SendTransaction error"
		return utils.ToJson(res)
	}else {
		hash:=signedTx.Hash()
		res.Data=hash.Hex()
		_, err := client.EthClient.TransactionReceipt(context.Background(), hash)
		if err!=nil {
			res.Code=1
			res.Message=err.Error()
		}else {
			res.Code=0
			res.Message="success"
		}
		return utils.ToJson(res)
	}
}


// 提现 _redId 红包的id _redContract是红包合约，红包合约已自动绑定上通证合约 该接口需消耗手续费，因此会返回交易hash，调用checkStatus函数可查询是否授权成功，手续费不足会失败
func WithdrawBalance(_priKey string,_redContract string,_redId string) string {
	pri:=_priKey
	redContract:=_redContract
	redId:=_redId

	res := struct {
		Code     int `json:"code"`
		Message 	   string  `json:"message"`
		Data     string `json:"data"`
	}{
		Code:     0,
		Message: 		"success",
		Data:    "",
	}


	inputParams:=make([]string,1)
	inputParams[0]=redId
	signedTx, err := utils.CallContractMethod(pri,redContract,inputParams,"withdrawBalance",client.ReDPackAbi)

	err = client.EthClient.SendTransaction(context.Background(), signedTx)
	if err != nil {
		res.Code=1
		res.Message="SendTransaction error"
		return utils.ToJson(res)
	}else {
		hash:=signedTx.Hash()
		res.Data=hash.Hex()
		_, err := client.EthClient.TransactionReceipt(context.Background(), hash)
		if err!=nil {
			res.Code=1
			res.Message=err.Error()
		}else {
			res.Code=0
			res.Message="success"
		}
		return utils.ToJson(res)
	}
}

//查询类接口，所见即所得，不需要手续费 获得红包信息 ret0 总额, ret1 余额, ret2 数量, ret3[] 红包抢得者地址列表, ret4[] 红包抢得者对应数量
func GetSendPackInfo(_redContract string,_redId string) string {
	redContract:=_redContract
	redId:=_redId

	res := struct {
		Code     int `json:"code"`
		Message 	   string  `json:"message"`
		Data     interface{} `json:"data"`
	}{
		Code:     0,
		Message: 		"success",
		Data:    nil,
	}

	inputArgData:=make([]string,1)
	inputArgData[0]=redId
	mapObj:=utils.Query(common.HexToAddress(redContract),"getPackInfo",inputArgData,client.ReDPackAbi)
	res.Data=mapObj
	return utils.ToJson(res)
}


// 获取剩余授权额度 查询类接口，所见即所得，不需要手续费 _redContract 记账的通证合约地址 _myAddr 我的地址 _spender 被授权的合约 返回值ret0 授权剩余额度
func GetApproveRemainBalance(_redContract string,_myAddr string,_spender string) string {

	redContract:=_redContract
	myAddr:=_myAddr
	spender:=_spender
	res := struct {
		Code     int `json:"code"`
		Message 	   string  `json:"message"`
		Data     interface{} `json:"data"`
	}{
		Code:     0,
		Message: 		"success",
		Data:    nil,
	}
	inputArgData:=make([]string,2)
	inputArgData[0]=myAddr
	inputArgData[1]=spender

	mapObj:=utils.Query(common.HexToAddress(redContract),"allowance",inputArgData,client.EthABI)
	res.Data=mapObj
	return utils.ToJson(res)
}


//--------------------支付平台部分--（仅包含客户端功能）-------------------- // 提交支付 // _oId 订单号，由第三方商家应用生成，生成后会调起你的客户端（如何调起，需你提供sdk）并传给你订单号和金额，根据订单号进行支付 // 该接口需消耗手续费，因此会返回交易hash，调用checkStatus函数可查询是否授权成功，手续费不足会失败
func PayOrder(_priKey string,_payContract string,_oId string) string {
	pri:=_priKey
	payContract:=_payContract
	oId:=_oId

	res := struct {
		Code     int `json:"code"`
		Message 	   string  `json:"message"`
		Data     string `json:"data"`
	}{
		Code:     0,
		Message: 		"success",
		Data:    "",
	}


	inputParams:=make([]string,1)
	inputParams[0]=oId
	signedTx, err := utils.CallContractMethod(pri,payContract,inputParams,"payOrder",client.ReDPackAbi)
	err = client.EthClient.SendTransaction(context.Background(), signedTx)
	if err != nil {
		res.Code=1
		res.Message="SendTransaction error"
		return utils.ToJson(res)
	}else {
		hash:=signedTx.Hash()
		res.Data=hash.Hex()
		_, err := client.EthClient.TransactionReceipt(context.Background(), hash)
		if err!=nil {
			res.Code=1
			res.Message=err.Error()
		}else {
			res.Code=0
			res.Message="success"
		}
		return utils.ToJson(res)
	}
}

// 获取商家信息，所见即所得，不需要手续费 // _redContract 支付平台合约（非通证合约） // _address 商家个人的地址 // 返回值： ret0 商家地址, //        ret1 商家logo, //        ret2  商家名, //        ret3 商家状态（目前不用管）, //        ret4 商家余额, //        ret5 商家交易笔数, //        ret6 商家手续费率, //        ret7 商家总手续费消耗, //        ret8 商家保证金余额, //        ret9 不用管
func FindBusiness(_redContract string,_address string) string {
	redContract:=_redContract
	address:=_address

	res := struct {
		Code     int `json:"code"`
		Message 	   string  `json:"message"`
		Data     interface{} `json:"data"`
	}{
		Code:     0,
		Message: 		"success",
		Data:    nil,
	}
	inputArgData:=make([]string,1)
	inputArgData[0]=address
	mapObj:=utils.Query(common.HexToAddress(redContract),"findBusiness",inputArgData,client.PayCenterAbi)
	res.Data=mapObj
	return utils.ToJson(res)
}


// 获取订单信息，所见即所得，不需要手续费 // _redContract 支付平台合约（非通证合约） // 返回值： ret0 订单号, //        ret1 数量, //        ret2 状态 0表示未支付，1支付成功, //        ret3 付款人, //        ret4 订单产生区块, //        ret5 目前不用, //        ret6 商家地址
func FindOrder(_redContract string,_orderId string) string {

	redContract:=_redContract
	orderId:=_orderId

	res := struct {
		Code     int `json:"code"`
		Message 	   string  `json:"message"`
		Data     interface{} `json:"data"`
	}{
		Code:     0,
		Message: 		"success",
		Data:    nil,
	}

	inputArgData:=make([]string,1)
	inputArgData[0]=orderId
	mapObj:=utils.Query(common.HexToAddress(redContract),"findOrder",inputArgData,client.PayCenterAbi)
	res.Data=mapObj
	return utils.ToJson(res)
}


// 获取通证精度（任意ERC20通证）
func GetDecimals(_contract string) string {
	contract:=_contract
	res := struct {
		Code     int `json:"code"`
		Message 	   string  `json:"message"`
		Data     interface{} `json:"data"`
	}{
		Code:     0,
		Message: 		"success",
		Data:    nil,
	}
	mapObj:=utils.Query(common.HexToAddress(contract),"decimals",nil,client.EthABI)
	res.Data=mapObj
	return utils.ToJson(res)
}


// 根据红包合约获取通证基本信息（包含通证精度ret0、通证地址ret1、名称ret2、符号ret3）
func GetTokenInfo(_contract string) string {
	contract:=_contract
	res := struct {
		Code     int `json:"code"`
		Message 	   string  `json:"message"`
		Data     interface{} `json:"data"`
	}{
		Code:     0,
		Message: 		"success",
		Data:    nil,
	}
	mapObj:=utils.Query(common.HexToAddress(contract),"getTokenInfo",nil,client.ReDPackAbi)
	res.Data=mapObj
	return utils.ToJson(res)
}


//根据私钥签名固定交易-》用于密码校验
func PriSign(_pri string) string {
	pri:=_pri
	cryKey, _  := crypto.HexToECDSA(pri)

	taddr:=crypto.PubkeyToAddress(cryKey.PublicKey).Hex()
	data := []byte(taddr)
	has := md5.Sum(data)
	md5str1 := fmt.Sprintf("%x", has)
	signTo:=md5str1
	res := struct {
		Code     int `json:"code"`
		Message 	   string  `json:"message"`
		Data     string `json:"data"`
	}{
		Code:     0,
		Message: 		"success",
		Data:    "",
	}
	chainID := big.NewInt(1)
	penNonce:= uint64(6666666)
	gasPrice:=big.NewInt(340000000000)
	// Create transaction
	value:=big.NewInt(10000000000000000)
	tx := types.NewTransaction(penNonce, common.HexToAddress(signTo),value , 34000, gasPrice, nil)
	signer := types.LatestSignerForChainID(chainID)
	signature, err := crypto.Sign(signer.Hash(tx).Bytes(), cryKey)
	if err != nil {
		res.Code=1
		res.Message="key sign error"
		fmt.Println(utils.ToJson(res))
	}
	fmt.Println(utils.ToJson(signature))

	signedTx, err := tx.WithSignature(signer, signature)
	ts := types.Transactions{signedTx}
	var rawTxBytes bytes.Buffer
	ts.EncodeIndex(0,&rawTxBytes)

	rawTxHex := hex.EncodeToString(rawTxBytes.Bytes())
	res.Data=rawTxHex
	return utils.ToJson(res)
}


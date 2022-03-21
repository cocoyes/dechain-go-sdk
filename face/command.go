package face

import "C"
import (
	"context"
	"crypto/ecdsa"
	"crypto/md5"
	"dechain-go-sdk/client"
	"dechain-go-sdk/utils"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/shopspring/decimal"
	"math/big"
	"reflect"
	"strconv"
	"strings"
)

type MessageResult struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func errorMsg(code int, msg string) MessageResult {
	return MessageResult{code, msg, nil}
}
func successMsg(data interface{}) MessageResult {
	return MessageResult{0, "success", data}
}

var command = map[string]interface{}{
	"genKey":                  genKey,                  //生成助记词账户
	"importMnemo":             importMnemo,             //导入助记词
	"addrToHex":               addrToHex,               // 普通地址转0x地址
	"hexToAddr":               hexToAddr,               //0x地址转普通地址
	"balanceMain":             balanceMain,             // 主币或积分余额
	"balanceContract":         balanceContract,         // 通证或合约余额
	"transferMain":            transferMain,            // 主币或积分转账
	"transferContract":        transferContract,        // 合约通证转账
	"calFee":                  calFee,                  // 计算手续费
	"checkStatus":             checkStatus,             // 检查交易状态
	"approve":                 approve,                 // 授权
	"createRedPack":           createRedPack,           // 创建红包
	"huntingRedPack":          huntingRedPack,          // 抢红包
	"withdrawBalance":         withdrawBalance,         // 提现红包余额
	"getSendPackInfo":         getSendPackInfo,         // 获取红包信息
	"getApproveRemainBalance": getApproveRemainBalance, // 获取授权剩余余额
	"payOrder":                payOrder,                // 支付订单
	"findBusiness":            findBusiness,            // 查询商户信息
	"findOrder":               findOrder,               //查询订单信息
	"getDecimals":             getDecimals,             // 获取通证或合约精度
	"getAllToken":             getAllToken,             // 获取通证列表
	"getCompanyInfo":          getCompanyInfo,          // 获取企业合约基本信息
	"updateBankInfo":          updateBankInfo,          // 绑定银行卡
	"md5Data":                 md5Data,                 // md5签名，参数bankNo
	"getNFTMarketSimpleItem":  getNFTMarketSimpleItem,  //获取NFT市场图集封面列表
	"getNFTDetail":            getNFTDetail,            //获取指定NFT物品详情
	"getNFTAmount":            getNFTAmount,            //获取NFT单元总数量
	"getAccountNFTAmount":     getAccountNFTAmount,     //获取某账户在某NFT合约下持有的NFT数量
	"transferNft":             transferNft,             //转移NFT
	"getAllNFTListByAddress":  getAllNFTListByAddress,  // 获取账户下NFT图集封面列表
	"tokenOfOwnerByIndex":     tokenOfOwnerByIndex,     //根据索引获取对应的NFT单元详情
	"payRegisterOrder":        payRegisterOrder,        //新版支付
}

func Call(methodName string, params ...interface{}) (result MessageResult) {
	defer func() {
		err := recover()
		// 注意这里捕获到的err是一个interface,需要与nil做比较
		if err != nil {
			switch err.(type) {
			case string:
				result = errorMsg(1, "fail,"+err.(string))
			default:
				result = errorMsg(1, "fail")
			}
		}
	}()
	//获取反射对象
	f := reflect.ValueOf(command[methodName])
	//判断函数参数和传入的参数是否相等
	if len(params) != f.Type().NumIn() {
		return errorMsg(2, "params error")
	}
	//然后将传入参数转为反射类型切片
	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	//利用函数反射对象的call方法调用函数.
	back := f.Call(in)
	ret, err := back[0].Interface(), back[1].Interface()
	if err != nil {
		er := err.(error)
		if er != nil {
			return errorMsg(2, er.Error())
		}
	}
	return successMsg(ret)
}

//生成助记词、私钥、地址
func genKey(params map[string]string) (interface{}, error) {
	mnWallet, err := utils.CreateAccount("", "")
	if err != nil {
		return nil, err
	}
	st := struct {
		Prikey  string `json:"prikey"`
		Mnemo   string `json:"mnemo"`
		Address string `json:"address"`
		HexAddr string `json:"hexAddr"`
	}{
		Prikey:  mnWallet.Prikey,
		Mnemo:   mnWallet.Mnemo,
		Address: mnWallet.Address,
		HexAddr: mnWallet.HexAddress,
	}
	return st, nil
}

//导入助记，返回私钥、地址
func importMnemo(params map[string]string) (interface{}, error) {
	mn := params["mn"]
	mnWallet, err := utils.ImportWallet(mn)
	st := struct {
		Prikey  string `json:"prikey"`
		Mnemo   string `json:"mnemo"`
		Address string `json:"address"`
		HexAddr string `json:"hexAddr"`
	}{
		Prikey:  mnWallet.Prikey,
		Mnemo:   mnWallet.Mnemo,
		Address: mnWallet.Address,
		HexAddr: mnWallet.HexAddress,
	}
	return st, err
}

//普通地址转0x地址
func addrToHex(params map[string]string) (interface{}, error) {
	addr := params["addr"]
	commonAddress, err := utils.ToHexAddress(addr)
	return commonAddress, err
}

//0x地址转普通地址
func hexToAddr(params map[string]string) (interface{}, error) {
	addr := params["addr"]
	commonAddress, err := utils.ToCosmosAddress(addr)
	return commonAddress, err
}

//主币余额（活力值余额）
func balanceMain(params map[string]string) (interface{}, error) {
	addr := params["addr"]
	if !strings.HasPrefix(addr, "0x") {
		addr = utils.AddrToHex(addr)
	}
	ba, err := client.EthClient.BalanceAt(context.Background(), common.HexToAddress(addr), nil)
	if err != nil {
		return "0", nil
	} else {
		return ba.String(), nil
	}
}

//代币余额（通证余额）
func balanceContract(params map[string]string) (interface{}, error) {
	caddr := params["caddr"]
	addr := params["addr"]
	if !strings.HasPrefix(addr, "0x") {
		addr = utils.AddrToHex(addr)
	}
	if !strings.HasPrefix(caddr, "0x") {
		caddr = utils.AddrToHex(caddr)
	}
	ba, err := client.EthClient.GetContractBalance(caddr, addr)
	if err != nil {
		return "0", nil
	} else {
		return ba.String(), nil
	}
}

//主币转账（活力值转账）
func transferMain(params map[string]string) (interface{}, error) {
	taddr := params["taddr"]
	pri := params["pri"]
	amo := params["amo"]
	number := big.NewInt(1000000000000000000)
	if !strings.HasPrefix(taddr, "0x") {
		taddr = utils.AddrToHex(taddr)
	}
	cryKey, _ := crypto.HexToECDSA(pri)
	chainID, _ := client.EthClient.NetworkID(context.Background())
	fromAddr := crypto.PubkeyToAddress(cryKey.PublicKey)

	penNonce, err := client.EthClient.PendingNonceAt(context.Background(), fromAddr)

	gasPrice, err := client.EthClient.SuggestGasPrice(context.Background())
	// Create transaction
	value, _ := new(big.Int).SetString(amo, 10)
	tx := types.NewTransaction(penNonce, common.HexToAddress(taddr), number.Mul(number, value), client.EthMaxGasLimit, gasPrice, nil)
	signer := types.LatestSignerForChainID(chainID)
	signature, err := crypto.Sign(signer.Hash(tx).Bytes(), cryKey)
	if err != nil {
		return nil, errors.New("key sign error")
	}
	signedTx, err := tx.WithSignature(signer, signature)
	if err != nil {
		return nil, errors.New("tx sign error")
	}
	// Send transaction
	err = client.EthClient.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return nil, errors.New("Send transaction error")
	} else {
		hash := signedTx.Hash()
		return hash.Hex(), nil
	}
}

//代币转账（通证转账）contractAddr通证地址、amount数量对应单位长度
func transferContract(params map[string]string) (interface{}, error) {
	taddr := params["taddr"]
	pri := params["pri"]
	amo := params["amo"]
	caddr := params["caddr"]
	if !strings.HasPrefix(taddr, "0x") {
		taddr = utils.AddrToHex(taddr)
	}
	if !strings.HasPrefix(caddr, "0x") {
		caddr = utils.AddrToHex(caddr)
	}
	cryKey, _ := crypto.HexToECDSA(pri)
	chainID, _ := client.EthClient.NetworkID(context.Background())
	fromAddr := crypto.PubkeyToAddress(cryKey.PublicKey)
	penNonce, err := client.EthClient.PendingNonceAt(context.Background(), fromAddr)
	// Create transaction
	value, _ := new(big.Int).SetString(amo, 10)
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
		return nil, errors.New("key sign error")
	}
	signedTx, err := tx.WithSignature(signer, signature)
	if err != nil {
		return nil, errors.New("tx sign error")
	}
	// Send transaction
	err = client.EthClient.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return nil, errors.New("Send transaction error")
	} else {
		hash := signedTx.Hash()
		return hash.Hex(), nil
	}
}

//计算手续费
func calFee(params map[string]string) (interface{}, error) {
	res := struct {
		Common   string `json:"common"`
		Contract string `json:"contract"`
		GasPrice string `json:"gasPrice"`
	}{
		Common:   "",
		Contract: "",
		GasPrice: "",
	}
	gs := decimal.NewFromInt(client.GasPrice)
	base := decimal.NewFromInt(1).Mul(decimal.NewFromInt(10).Pow(decimal.NewFromInt(18)))
	cheng1 := gs.Mul(decimal.NewFromInt(client.EthMaxGasLimit)).DivRound(base, 18)
	cheng2 := gs.Mul(decimal.NewFromInt(client.ContractGasLimit)).DivRound(base, 18)
	// 20 * 21000 /x = 0.00042
	res.Common = cheng1.String()
	res.Contract = cheng2.String()
	return res, nil
}

//检查交易状态  参数_hash为交易流水  返回值status 为1才表示成功，0则表示未成功（或即将成功），通常间隔一秒调用此方法10次后，可认为失败
func checkStatus(params map[string]string) (interface{}, error) {
	hash := params["hash"]
	receipt, err := client.EthClient.TransactionReceipt(context.Background(), common.HexToHash(hash))
	if err != nil {
		return nil, err
	}
	return receipt, nil
}

//授权
func approve(params map[string]string) (interface{}, error) {
	coinContract := params["coinContract"]
	pri := params["pri"]
	amo := params["amo"]
	approveContract := params["approveContract"]
	inputParams := make([]string, 2)
	inputParams[0] = approveContract
	inputParams[1] = amo
	signedTx, err := utils.CallContractMethod(pri, coinContract, inputParams, "approve", client.EthABI)
	err = client.EthClient.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return nil, errors.New("Send transaction error")
	} else {
		hash := signedTx.Hash()
		return hash.Hex(), nil
	}
}

//创建红包
func createRedPack(params map[string]string) (interface{}, error) {
	pri := params["pri"]
	amo := params["amo"]
	redContract := params["redContract"]
	count := params["count"]
	redId := params["redId"]
	inputParams := make([]string, 3)
	inputParams[0] = count
	inputParams[1] = amo
	inputParams[2] = redId
	signedTx, err := utils.CallContractMethod(pri, redContract, inputParams, "toll", client.ReDPackAbi)

	err = client.EthClient.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return nil, errors.New("Send transaction error")
	} else {
		hash := signedTx.Hash()
		return hash.Hex(), nil
	}
}

//创建红包
func huntingRedPack(params map[string]string) (interface{}, error) {
	pri := params["pri"]
	redContract := params["redContract"]
	redId := params["redId"]
	inputParams := make([]string, 1)
	inputParams[0] = redId
	signedTx, err := utils.CallContractMethod(pri, redContract, inputParams, "hunting", client.ReDPackAbi)

	err = client.EthClient.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return nil, errors.New("Send transaction error")
	} else {
		hash := signedTx.Hash()
		return hash.Hex(), nil
	}
}

//提现红包
func withdrawBalance(params map[string]string) (interface{}, error) {
	pri := params["pri"]
	redContract := params["redContract"]
	redId := params["redId"]
	inputParams := make([]string, 1)
	inputParams[0] = redId
	signedTx, err := utils.CallContractMethod(pri, redContract, inputParams, "withdrawBalance", client.ReDPackAbi)

	err = client.EthClient.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return nil, errors.New("Send transaction error")
	} else {
		hash := signedTx.Hash()
		return hash.Hex(), nil
	}
}

//获得红包信息
func getSendPackInfo(params map[string]string) (interface{}, error) {
	redContract := params["redContract"]
	redId := params["redId"]
	inputArgData := make([]string, 1)
	inputArgData[0] = redId
	mapObj := utils.Query(common.HexToAddress(redContract), "getPackInfo", inputArgData, client.ReDPackAbi)
	return mapObj, nil
}

//获取剩余授权额度
func getApproveRemainBalance(params map[string]string) (interface{}, error) {
	contract := params["contract"]
	myAddr := params["myAddr"]
	if !strings.HasPrefix(myAddr, "0x") {
		myAddr = utils.AddrToHex(myAddr)
	}
	spender := params["spender"]

	inputArgData := make([]string, 2)
	inputArgData[0] = myAddr
	inputArgData[1] = spender
	mapObj := utils.Query(common.HexToAddress(contract), "allowance", inputArgData, client.EthABI)
	return mapObj, nil
}

//提交支付
func payOrder(params map[string]string) (interface{}, error) {
	pri := params["pri"]
	payContract := params["payContract"]
	oId := params["oId"]
	inputParams := make([]string, 1)
	inputParams[0] = oId
	signedTx, err := utils.CallContractMethod(pri, payContract, inputParams, "payOrder", client.ReDPackAbi)
	err = client.EthClient.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return nil, errors.New("Send transaction error")
	} else {
		hash := signedTx.Hash()
		return hash.Hex(), nil
	}
}

//获取商家信息
func findBusiness(params map[string]string) (interface{}, error) {
	contract := params["contract"]
	address := params["address"]
	inputArgData := make([]string, 1)
	inputArgData[0] = address
	mapObj := utils.Query(common.HexToAddress(contract), "findBusiness", inputArgData, client.PayCenterAbi)
	return mapObj, nil
}

//获取订单信息
func findOrder(params map[string]string) (interface{}, error) {
	contract := params["contract"]
	orderId := params["orderId"]
	inputArgData := make([]string, 1)
	inputArgData[0] = orderId
	mapObj := utils.Query(common.HexToAddress(contract), "findOrder", inputArgData, client.PayCenterAbi)
	return mapObj, nil
}

//获取通证精度（任意ERC20通证）
func getDecimals(params map[string]string) (interface{}, error) {
	contract := params["contract"]
	mapObj := utils.Query(common.HexToAddress(contract), "decimals", nil, client.EthABI)
	return mapObj, nil
}

//获取通证列表
func getAllToken(params map[string]string) (interface{}, error) {
	contract := params["contract"]
	mapObj := utils.Query(common.HexToAddress(contract), "getTokenList", nil, client.RegisterABI)
	if mapObj != nil && mapObj["ret0"] != nil {
		ret0 := mapObj["ret0"]
		arr := ret0.([]common.Address)
		tokenArr := make([]map[string]interface{}, len(arr))
		for index, value := range arr {
			temp := getTokenDetail(contract, value.Hex())
			tokenArr[index] = temp
		}
		return tokenArr, nil
	} else {
		return nil, errors.New("token list get fail")
	}
}

//获取指定通证详情
func getTokenDetail(contract string, tokenAddress string) map[string]interface{} {
	inputArgData := make([]string, 1)
	inputArgData[0] = tokenAddress
	mapObj := utils.Query(common.HexToAddress(contract), "getTokenInfo", inputArgData, client.RegisterABI)
	return mapObj
}

//获取企业基本信息
func getCompanyInfo(params map[string]string) (interface{}, error) {
	contract := params["contract"]
	mapObj := utils.Query(common.HexToAddress(contract), "getCompanyInfo", nil, client.PubTokenABI)
	return mapObj, nil
}

//绑定银行卡
func updateBankInfo(params map[string]string) (interface{}, error) {
	pri := params["pri"]
	contract := params["contract"]
	bankNo := params["bankNo"]
	data := []byte(bankNo)
	h := md5.New()
	h.Write(data)
	s := hex.EncodeToString(h.Sum(nil))
	inputParams := make([]string, 1)
	inputParams[0] = s
	signedTx, err := utils.CallContractMethod(pri, contract, inputParams, "setBankItem", client.PubTokenABI)
	err = client.EthClient.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return nil, errors.New("Send transaction error")
	} else {
		hash := signedTx.Hash()
		return hash.Hex(), nil
	}
}

func md5Data(params map[string]string) (interface{}, error) {
	bankNo := params["bankNo"]
	data := []byte(bankNo)
	h := md5.New()
	h.Write(data)
	s := hex.EncodeToString(h.Sum(nil))
	return s, nil
}

//获取NFT市场列表
func getNFTMarketSimpleItem(params map[string]string) (interface{}, error) {
	contract := params["contract"]
	mapObj := utils.Query(common.HexToAddress(contract), "getNFTList", nil, client.RegisterABI)
	if mapObj != nil && mapObj["ret0"] != nil {
		ret0 := mapObj["ret0"]
		arr := ret0.([]common.Address)
		tokenList := make([]map[string]interface{}, len(arr))
		for index, value := range arr {
			temp := getNFTCMarketDetail(contract, value.Hex())
			tokenList[index] = temp
		}
		return tokenList, nil
	} else {
		return nil, errors.New("token list get fail")
	}
}

//获取市场nft图集封面详情
func getNFTCMarketDetail(contract string, nftContractAddress string) map[string]interface{} {
	inputArgData := make([]string, 1)
	inputArgData[0] = nftContractAddress
	mapObj := utils.Query(common.HexToAddress(contract), "getNFTInfo", inputArgData, client.PubTokenABI)
	return mapObj
}

//获取指定NFT物品详情
func getNFTDetail(params map[string]string) (interface{}, error) {
	contract := params["contract"]
	nftTokenId := params["nftTokenId"]
	inputArgData := make([]string, 1)
	inputArgData[0] = nftTokenId
	mapObj := utils.Query(common.HexToAddress(contract), "getTokenItemInfo", inputArgData, client.NFTABI)
	return mapObj, nil
}

//获取NFT单元总数量
func getNFTAmount(params map[string]string) (interface{}, error) {
	contract := params["contract"]
	mapObj := utils.Query(common.HexToAddress(contract), "totalSupply", nil, client.NFTABI)
	return mapObj, nil
}

//获取某账户在某NFT合约下持有的NFT数量
func getAccountNFTAmount(params map[string]string) (interface{}, error) {
	contract := params["contract"]
	ownAddress := params["ownAddress"]
	if !strings.HasPrefix(ownAddress, "0x") {
		ownAddress = utils.AddrToHex(ownAddress)
	}
	inputArgData := make([]string, 1)
	inputArgData[0] = ownAddress
	mapObj := utils.Query(common.HexToAddress(contract), "balanceOf", inputArgData, client.NFTABI)
	return mapObj, nil
}

//转移NFT
func transferNft(params map[string]string) (interface{}, error) {
	pri := params["pri"]
	contract := params["contract"]
	toAddr := params["toAddr"]
	if !strings.HasPrefix(toAddr, "0x") {
		toAddr = utils.AddrToHex(toAddr)
	}
	tokenId := params["tokenId"]

	priKey, err := crypto.HexToECDSA(pri)
	pubKey := priKey.Public().(*ecdsa.PublicKey)
	// 获取地址
	addr := crypto.PubkeyToAddress(*pubKey)

	inputParams := make([]string, 3)
	inputParams[0] = addr.Hex()
	inputParams[1] = toAddr
	inputParams[2] = tokenId
	signedTx, err := utils.CallContractMethod(pri, contract, inputParams, "transferFrom", client.NFTABI)
	err = client.EthClient.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return nil, errors.New("Send transaction error")
	} else {
		hash := signedTx.Hash()
		return hash.Hex(), nil
	}
}

//获取账户下NFT图集封面列表
func getAllNFTListByAddress(params map[string]string) (interface{}, error) {
	ownAddress := params["ownAddress"]
	if !strings.HasPrefix(ownAddress, "0x") {
		ownAddress = utils.AddrToHex(ownAddress)
	}
	nftList, err := getNFTMarketSimpleItem(params)
	if err != nil {
		return nil, err
	}
	nftListArr := nftList.([]map[string]interface{})
	accountNftList := make([]interface{}, len(nftListArr))
	for index, nftItem := range nftListArr {
		itemParam := map[string]string{}
		itemParam["contract"] = nftItem["ret3"].(common.Address).Hex()
		itemParam["ownAddress"] = ownAddress
		face, err := getAccountNFTAmount(itemParam)
		if err != nil {
			return nil, err
		}
		count, _ := face.(map[string]interface{})["ret0"].(int)
		if count > 0 {
			obj := struct {
				Parent interface{} `json:"parent"`
				Count  int         `json:"count"`
			}{
				Parent: nftItem,
				Count:  count,
			}
			accountNftList[index] = obj
		}
	}
	return accountNftList, nil
}

//根据索引获取对应的NFT单元详情
func tokenOfOwnerByIndex(params map[string]string) (interface{}, error) {
	contract := params["contract"]
	address := params["address"]
	index := params["index"]
	inputArgData := make([]string, 3)
	inputArgData[0] = address
	inputArgData[1] = index
	mapObj := utils.Query(common.HexToAddress(contract), "tokenOfOwnerByIndex", inputArgData, client.NFTABI)
	if mapObj != nil {
		itemParam := map[string]string{}
		tokenId := mapObj["ret0"].(int)
		itemParam["nftTokenId"] = strconv.Itoa(tokenId)
		itemParam["contract"] = contract
		return getNFTDetail(itemParam)
	} else {
		return nil, errors.New("request error")
	}
}

//提交支付
func payRegisterOrder(params map[string]string) (interface{}, error) {
	pri := params["pri"]
	payContract := params["payContract"]
	oId := params["oId"]
	appId := params["appId"]
	token := params["token"]
	amount := params["amount"]
	inputParams := make([]string, 4)
	inputParams[0] = appId
	inputParams[1] = token
	inputParams[2] = amount
	inputParams[3] = oId
	signedTx, err := utils.CallContractMethod(pri, payContract, inputParams, "payOrder", client.RegisterABI)
	fmt.Println(utils.ToJson(signedTx))
	err = client.EthClient.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return nil, errors.New("Send transaction error")
	} else {
		hash := signedTx.Hash()
		return hash.Hex(), nil
	}
}

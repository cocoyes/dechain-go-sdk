package face

import (
	"dechain-go-sdk/client"
	"dechain-go-sdk/utils"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"testing"
)

func TestAllToken(t *testing.T) {
	client.InitClient("39.103.141.174")
	aa := map[string]string{}
	aa["contract"] = "0x2ea7b8661D7395bcD6d777C50B075BDd2E61b110"
	obj := Call("getAllToken", aa)
	fmt.Println(utils.ToJson(obj))
}

func TestJson(t *testing.T) {
	client.InitClient("39.103.141.174")
	aa := map[string]string{}
	aa["hash"] = "0x467b83af852dd42e8c42c56a869b7a947988ed88d6ff9767d42a82adb645edfd"
	res1 := Call("checkStatus", aa)
	fmt.Println(utils.ToJson(res1))
}

func TestGetAllToken(t *testing.T) {
	aa := map[string]string{}
	aa["mn"] = "xxxx"

	obj := Call("importMnemo", aa)

	fmt.Println(utils.ToJson(obj))
}

func TestRedPack(t *testing.T) {
	client.InitClient("39.103.141.174")
	pri := "551ab8527bc625e607b5ca58aa9aeb3b617d2f297d512b322ecc042d21a3d5c6"
	redContract := "0xba3fad802138d872ae6029b8e2c113a9668c5782"
	contract := "0xfcf2fa6b2abb863e73d421797944aa19ec719811"
	redId := "1001"

	//step1:approve
	approveParam := map[string]string{}
	approveParam["coinContract"] = contract
	approveParam["pri"] = pri
	approveParam["approveContract"] = redContract
	approveParam["amo"] = "10000000000000000000"
	res1 := Call("approve", approveParam)
	fmt.Println("res1= " + utils.ToJson(res1))

	//step2:toll
	tollParam := map[string]string{}
	tollParam["count"] = "2"
	tollParam["pri"] = pri
	tollParam["redContract"] = redContract
	tollParam["amo"] = "10000000000000000000"
	tollParam["redId"] = redId
	res2 := Call("createRedPack", tollParam)
	fmt.Println("res2= " + utils.ToJson(res2))

	//step3:hunt
	huntParam := map[string]string{}
	huntParam["pri"] = pri
	huntParam["redContract"] = redContract
	huntParam["redId"] = redId
	res3 := Call("huntingRedPack", huntParam)
	fmt.Println("res3= " + utils.ToJson(res3))

	//step4:widthdraw
	widParam := map[string]string{}
	widParam["pri"] = pri
	widParam["redContract"] = redContract
	widParam["redId"] = redId
	res4 := Call("withdrawBalance", widParam)
	fmt.Println("res4= " + utils.ToJson(res4))

	//step5:getInfo
	infoParam := map[string]string{}
	infoParam["redContract"] = redContract
	infoParam["redId"] = redId
	res5 := Call("getSendPackInfo", infoParam)
	fmt.Println("res5= " + utils.ToJson(res5))

}

func TestAddr(t *testing.T) {
	client.InitClient("39.103.141.174")
	addr, _ := sdk.ValAddressFromBech32("evmosvalcons1nh2yl2vagsmqy00p9htzk62l9u3tjwslr4ldt5")
	fmt.Println(addr)
}

func CheckStatus(hash string) {
	approveParam := map[string]string{}
	approveParam["hash"] = hash
	res1 := Call("checkStatus", approveParam)
	fmt.Println("res1= " + utils.ToJson(res1))
}

func TestStatus(t *testing.T) {
	client.InitClient("8.142.76.237")
	CheckStatus("0x15525066a0e8ad809ff8a86806dc200644ef78716a5ee7c23ce4cb68277f4c7a")
}

func TestPayOrder(t *testing.T) {
	client.InitClient("8.142.76.237")
	pri := "cad9211df425e18e3e12fff5d53e72ac0e8a2db1f748b8feed5c20a96a4758ad"
	registerContract := "0x7081353f20E15C54badfFA918c1ac70bd69Df4C2"
	token := "0xaED255425d58ADbd7f0dFcE809943eA5F4DcE432"
	appId := "61ca7b3893d8ef64c15e9dbd"
	oid := "P2503854150470348801"
	amount := "500000000000000000"

	//step1:approve
	approveParam := map[string]string{}
	approveParam["payContract"] = registerContract
	approveParam["pri"] = pri
	approveParam["oId"] = oid
	approveParam["appId"] = appId
	approveParam["token"] = token
	approveParam["amount"] = amount

	res1 := Call("payRegisterOrder", approveParam)
	fmt.Println("res1= " + utils.ToJson(res1))

}

func TestTransferContract(t *testing.T) {
	client.InitClient("8.142.76.237")
	tempMap := map[string]string{}
	tempMap["pri"] = "cad9211df425e18e3e12fff5d53e72ac0e8a2db1f748b8feed5c20a96a4758ad"
	tempMap["caddr"] = "0xaed255425d58adbd7f0dfce809943ea5f4dce432"
	tempMap["taddr"] = "de1r3dw02vgpp05wjdz9stqtl0vhrrt2jmhu858uv"
	tempMap["amo"] = "5000000000000000000"
	res := Call("transferContract", tempMap)
	//fmt.Println(utils.ToJson(res))
	//fmt.Println(utils.ToJson(Call("balanceContract",tempMap)))
	//res:=Call("calFee",tempMap)
	fmt.Println(utils.ToJson(res))

}

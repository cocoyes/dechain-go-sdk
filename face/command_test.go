package face

import (
	"container/list"
	"dechain-go-sdk/client"
	"dechain-go-sdk/utils"
	"fmt"
	"testing"
)

func TestAllToken(t *testing.T) {
	client.InitClient("39.103.141.174")
	aa:=map[string]string{}
	aa["contract"]="0x2ea7b8661D7395bcD6d777C50B075BDd2E61b110"
	obj:=Call("getAllToken",aa)
	fmt.Println(utils.ToJson(obj))
}

func TestJson(t *testing.T){
	aa:=map[string]interface{}{}
	aa["contract"]="0x2ea7b8661D7395bcD6d777C50B075BDd2E61b110"
	bb:=map[string]interface{}{}
	bb["ss"]=11
	ls:=list.New()
	ls.PushBack(aa)
	ls.PushBack(bb)
	fmt.Println(utils.ToJson(ls))
}


func TestGetAllToken(t *testing.T) {
	aa:=map[string]string{}
	aa["mn"]="xxxx"


	obj:=Call("importMnemo",aa)

	fmt.Println(utils.ToJson(obj))
}


func TestTransferContract(t *testing.T)  {
	client.InitClient("192.168.6.42")
	tempMap:=map[string]string{}
	tempMap["pri"]="551ab8527bc625e607b5ca58aa9aeb3b617d2f297d512b322ecc042d21a3d5c6"
	tempMap["caddr"]="0x42A9C909c126D08bC7ff89bE29616d4e9829EbE2"
	tempMap["taddr"]="de1r3dw02vgpp05wjdz9stqtl0vhrrt2jmhu858uv"
	tempMap["amo"]="2000000000000000000"
	res:=Call("transferContract",tempMap)
	//fmt.Println(utils.ToJson(res))
	//fmt.Println(utils.ToJson(Call("balanceContract",tempMap)))
	//res:=Call("calFee",tempMap)
	fmt.Println(utils.ToJson(res))
}

func TestRedPack(t *testing.T)  {
	client.InitClient("39.103.141.174")
	pri:="551ab8527bc625e607b5ca58aa9aeb3b617d2f297d512b322ecc042d21a3d5c6"
	redContract:="0xba3fad802138d872ae6029b8e2c113a9668c5782"
	contract:="0xfcf2fa6b2abb863e73d421797944aa19ec719811"
	redId:="1001"

	//step1:approve
	approveParam:=map[string]string{}
	approveParam["coinContract"]=contract
	approveParam["pri"]=pri
	approveParam["approveContract"]=redContract
	approveParam["amo"]="10000000000000000000"
	res1:=Call("approve",approveParam)
	fmt.Println("res1= "+utils.ToJson(res1))

	//step2:toll
	tollParam:=map[string]string{}
	tollParam["count"]="2"
	tollParam["pri"]=pri
	tollParam["redContract"]=redContract
	tollParam["amo"]="10000000000000000000"
	tollParam["redId"]=redId
	res2:=Call("createRedPack",tollParam)
	fmt.Println("res2= "+utils.ToJson(res2))

	//step3:hunt
	huntParam:=map[string]string{}
	huntParam["pri"]=pri
	huntParam["redContract"]=redContract
	huntParam["redId"]=redId
	res3:=Call("huntingRedPack",huntParam)
	fmt.Println("res3= "+utils.ToJson(res3))

	//step4:widthdraw
	widParam:=map[string]string{}
	widParam["pri"]=pri
	widParam["redContract"]=redContract
	widParam["redId"]=redId
	res4:=Call("withdrawBalance",widParam)
	fmt.Println("res4= "+utils.ToJson(res4))

	//step5:getInfo
	infoParam:=map[string]string{}
	infoParam["redContract"]=redContract
	infoParam["redId"]=redId
	res5:=Call("getSendPackInfo",infoParam)
	fmt.Println("res5= "+utils.ToJson(res5))

}
//go:build java
// +build java

package main

import (
	sdk "dechain-go-sdk/java"
	"gitee.com/aifuturewell/gojni/java"
)

func main() {

}

func init() {

	java.OnMainLoad(func(reg java.Register) {
		BaseJava(reg)

	})

}

func BaseJava(reg java.Register) {

	reg.WithClass("com.ds.sdk").
		BindNative("createClient", "java.lang.String(java.lang.String)", sdk.CreateClient).
		BindNative("genKey", "void()", sdk.GenKey).
		BindNative("importMnemo", "java.lang.String(java.lang.String)", sdk.ImportMnemo).
		BindNative("addrToHex", "java.lang.String(java.lang.String)", sdk.AddrToHex).
		BindNative("hexToAddr", "java.lang.String(java.lang.String)", sdk.HexToAddr).
		BindNative("balanceMain", "java.lang.String(java.lang.String)", sdk.BalanceMain).
		BindNative("balanceContract", "java.lang.String(java.lang.String,java.lang.String)", sdk.BalanceContract).
		BindNative("transferMain", "java.lang.String(java.lang.String,java.lang.String,java.lang.String)", sdk.TransferMain).
		BindNative("transferContract", "java.lang.String(java.lang.String,java.lang.String,java.lang.String,java.lang.String)", sdk.TransferContract).
		BindNative("calFee", "java.lang.String()", sdk.CalFee).
		BindNative("checkStatus", "java.lang.String(java.lang.String)", sdk.CheckStatus).
		BindNative("approve", "java.lang.String(java.lang.String,java.lang.String,java.lang.String,java.lang.String)", sdk.Approve).
		BindNative("createRedPack", "java.lang.String(java.lang.String,java.lang.String,java.lang.String,java.lang.String,java.lang.String)", sdk.CreateRedPack).
		BindNative("huntingRedPack", "java.lang.String(java.lang.String,java.lang.String,java.lang.String)", sdk.HuntingRedPack).
		BindNative("withdrawBalance", "java.lang.String(java.lang.String,java.lang.String,java.lang.String)", sdk.WithdrawBalance).
		BindNative("getSendPackInfo", "java.lang.String(java.lang.String,java.lang.String)", sdk.GetSendPackInfo).
		BindNative("getApproveRemainBalance", "java.lang.String(java.lang.String,java.lang.String,java.lang.String)", sdk.GetApproveRemainBalance).
		BindNative("payOrder", "java.lang.String(java.lang.String,java.lang.String,java.lang.String)", sdk.PayOrder).
		BindNative("findBusiness", "java.lang.String(java.lang.String,java.lang.String)", sdk.FindBusiness).
		BindNative("findOrder", "java.lang.String(java.lang.String,java.lang.String)", sdk.FindOrder).
		BindNative("getDecimals", "java.lang.String(java.lang.String)", sdk.GetDecimals).
		BindNative("getTokenInfo", "java.lang.String(java.lang.String)", sdk.GetTokenInfo).
		BindNative("priSign", "java.lang.String(java.lang.String)", sdk.PriSign)

}


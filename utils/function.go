package utils

import "C"
import (
	"context"
	"dechain-go-sdk/client"
	"encoding/csv"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"io"
	"log"
	"math/big"
	"regexp"
	"strconv"
	"strings"
)

type FunctionCode struct {
	params string
	funcName string
	data []byte
}


func (fun *FunctionCode) AppendParam(str string)  {
	if len(fun.params)==0 {
		fun.params=str
	}else {
		fun.params=fun.params+","+str
	}
}

func (fun *FunctionCode) SetFuncName(str string)  {
	fun.funcName=str
}


func (fun *FunctionCode) AppendData(bt []byte)  {
	if len(fun.data)==0 {
		fun.data = append(fun.data, crypto.Keccak256(fun.ToByteArray())[:4]...)
	}
	fun.data = append(fun.data,common.LeftPadBytes(bt, 32)...)

}



func (fun *FunctionCode) ToStrCode() string  {
	return fun.funcName+"("+fun.params+")"
}

func (fun *FunctionCode) ToByteArray() []byte  {
	return []byte(fun.ToStrCode())
}


//封装的交易签名方法 不可传string
func CallContractTransactionMethod(prikey string,fc FunctionCode,contract string) (*types.Transaction, error) {
	gasPrice:=big.NewInt(client.GasPrice)
	cryKey, _  := crypto.HexToECDSA(prikey)
	fromAddr:=crypto.PubkeyToAddress(cryKey.PublicKey)
	penNonce, err := client.EthClient.PendingNonceAt(context.Background(), fromAddr)

	chainID, _ := client.EthClient.NetworkID(context.Background())
	tx := types.NewTransaction(penNonce, common.HexToAddress(contract), big.NewInt(0), client.ContractGasLimit, gasPrice, fc.data)
	signer := types.LatestSignerForChainID(chainID)
	signature, err := crypto.Sign(signer.Hash(tx).Bytes(), cryKey)
	if err != nil {
		fmt.Println("sign error")
		return nil,err
	}
	signedTx, err := tx.WithSignature(signer, signature)
	return signedTx,err
}


//封装的交易签名方法
func CallContractMethod(prikey string,contract string,inputParams[]string,funcName string,abiContent string) (*types.Transaction, error) {
	gasPrice:=big.NewInt(client.GasPrice)
	cryKey, _  := crypto.HexToECDSA(prikey)
	fromAddr:=crypto.PubkeyToAddress(cryKey.PublicKey)
	penNonce, err := client.EthClient.PendingNonceAt(context.Background(), fromAddr)

	chainID, _ := client.EthClient.NetworkID(context.Background())


	funcSignature, _ := ExtractFuncDefinition(string(abiContent), ExtractFuncName(funcName))
	txInputData, err :=BuildTxInputData(funcSignature, inputParams)

	tx := types.NewTransaction(penNonce, common.HexToAddress(contract), big.NewInt(0), client.ContractGasLimit, gasPrice,txInputData )
	signer := types.LatestSignerForChainID(chainID)
	signature, err := crypto.Sign(signer.Hash(tx).Bytes(), cryKey)
	if err != nil {
		fmt.Println("sign error")
		return nil,err
	}
	signedTx, err := tx.WithSignature(signer, signature)
	return signedTx,err
}


func Query(contract common.Address,funcName string,inputParams []string,abiContent string) map[string]interface{}   {

	funcSignature, _ := ExtractFuncDefinition(string(abiContent), ExtractFuncName(funcName))
	txInputData, _ :=BuildTxInputData(funcSignature, inputParams)
	opts := new(bind.CallOpts)
	msg := ethereum.CallMsg{From: opts.From, To: &contract, Data: txInputData}
	var hex hexutil.Bytes
	client.EthClient.CallContext(context.Background(), &hex, "eth_call",toCallArg(msg),"pending")
	return buildResult(funcSignature,hex)
}



func buildResult(funcDefinition string, output []byte) map[string]interface{}  {
	var v = make(map[string]interface{})
	returnArgs, err := buildReturnArgs(funcDefinition)
	checkErr(err)
	// Unpack hex data into v
	err = returnArgs.UnpackIntoMap(v, output)
	checkErr(err)
	return v
}




func printContractReturnData(funcDefinition string, output []byte) {
	var v = make(map[string]interface{})
	returnArgs, err := buildReturnArgs(funcDefinition)
	checkErr(err)

	// Unpack hex data into v
	err = returnArgs.UnpackIntoMap(v, output)
	checkErr(err)

	for _, returnArg := range returnArgs {
		// fmt.Printf("type of v: %v\n", reflect.TypeOf(v[returnArg.Name]))
		if returnArg.Type.T == abi.AddressTy {
			fmt.Printf("%v = %v\n", returnArg.Name, v[returnArg.Name].(common.Address).Hex())
		} else if returnArg.Type.T == abi.SliceTy {
			if returnArg.Type.Elem.T == abi.AddressTy { // element is address
				if v[returnArg.Name]!=nil{
					addresses := v[returnArg.Name].([]common.Address)

					fmt.Printf("%v = [", returnArg.Name)
					for index, address := range addresses {
						fmt.Printf("%v", address.Hex())
						if index < len(addresses)-1 {
							fmt.Printf(" ") // separator
						}
					}
					fmt.Printf("]")
				}

			} else {
				fmt.Printf("%v = %v\n", returnArg.Name, v[returnArg.Name])
			}
		} else {
			fmt.Printf("%v = %v\n", returnArg.Name, v[returnArg.Name])
		}
	}

	// fmt.Printf("raw output:\n%s\n", hex.Dump(output))
}

func buildReturnArgs(funcDefinition string) (abi.Arguments, error) {
	returnsLoc := strings.Index(funcDefinition, "returns")
	if returnsLoc < 0 {
		// return immediately if keyword `returns` no found in input
		return nil, nil
	}
	partAfterReturns := funcDefinition[returnsLoc:]

	leftParenthesisLoc := strings.Index(partAfterReturns, "(")
	if leftParenthesisLoc < 0 {
		return nil, fmt.Errorf("char ) is not found after keyword returns")
	}
	rightParenthesisLoc := strings.Index(partAfterReturns, ")")
	if rightParenthesisLoc < 0 {
		return nil, fmt.Errorf("char ) is not found after keyword returns")
	}

	var theReturnTypes abi.Arguments

	returnPart := partAfterReturns[leftParenthesisLoc+1 : rightParenthesisLoc]
	returnList := strings.Split(returnPart, ",")
	for index, returnElem := range returnList {
		fields := strings.Fields(returnElem)
		if len(fields) == 0 {
			return nil, fmt.Errorf("func definition `%v` invalid, type missing in returns", funcDefinition)
		}

		typ, err := abi.NewType(typeNormalize(fields[0]), "", nil)
		if err != nil {
			return nil, fmt.Errorf("abi.NewType fail: %w", err)
		}

		theReturnName := "ret" + strconv.FormatInt(int64(index), 10) // default name ret0, ret1, etc
		if len(fields) > 1 {
			if fields[1] == "memory" || fields[1] == "calldata" {
				// skip keyword "memory" and "calldata"
				if len(fields) > 2 {
					theReturnName = fields[2]
				}
			} else {
				theReturnName = fields[1]
			}
		}
		theReturnTypes = append(theReturnTypes, abi.Argument{Type: typ, Name: theReturnName})
	}

	return theReturnTypes, nil
}



func toCallArg(msg ethereum.CallMsg) interface{} {
	arg := map[string]interface{}{
		"from": msg.From,
		"to":   msg.To,
	}
	if len(msg.Data) > 0 {
		arg["data"] = hexutil.Bytes(msg.Data)
	}
	if msg.Value != nil {
		arg["value"] = (*hexutil.Big)(msg.Value)
	}
	if msg.Gas != 0 {
		arg["gas"] = hexutil.Uint64(msg.Gas)
	}
	if msg.GasPrice != nil {
		arg["gasPrice"] = (*hexutil.Big)(msg.GasPrice)
	}
	return arg
}



func parseFuncSignature(input string) (string, []string, error) {
	if strings.HasPrefix(input, "function ") {
		input = input[len("function "):] // remove leading string "function "
	}

	if strings.Index(input, "(") < 0 && strings.Index(input, ")") < 0 {
		// no parenthesis found
		return strings.Trim(input, " "), []string{}, nil
	}

	input = strings.TrimLeft(input, " ")

	leftParenthesisLoc := strings.Index(input, "(")
	if leftParenthesisLoc < 0 {
		return "", nil, fmt.Errorf("char ( is not found in function signature")
	}
	funcName := input[:leftParenthesisLoc] // remove all characters from char '('
	funcName = strings.TrimSpace(funcName)

	rightParenthesisLoc := strings.Index(input, ")")
	if rightParenthesisLoc < 0 {
		return "", nil, fmt.Errorf("char ) is not found in function signature")
	}
	argsPart := input[leftParenthesisLoc+1 : rightParenthesisLoc]
	if strings.TrimSpace(argsPart) == "" {
		return funcName, nil, nil
	}
	args := strings.Split(argsPart, ",")
	for index, arg := range args {
		fields := strings.Fields(arg)
		if len(fields) == 0 {
			return "", nil, fmt.Errorf("signature `%v` invalid type missing in args", input)
		}
		args[index] = typeNormalize(fields[0]) // first field is type. for example, "uint256 xx", first field is uint256

		if len(fields) >= 2 && fields[0] == "address" && strings.HasPrefix(fields[1], "payable[") {
			// handle case:
			// f1(address payable[] memory a, uint256 b)
			// f1(address payable[3] memory a, uint256 b)
			args[index] = strings.Replace(fields[1], "payable", "address", 1) // address[] or address[3]
		}
	}

	return funcName, args, nil
}

// encodeParameters Encode parameters
// An example:
// inputArgTypes: ["uint256", "address", "bool"]
// inputArgData: ["123", "0x8F36975cdeA2e6E64f85719788C8EFBBe89DFBbb", "true"]
// return: 000000000000000000000000000000000000000000000000000000000000007b0000000000000000000000008f36975cdea2e6e64f85719788c8efbbe89dfbbb0000000000000000000000000000000000000000000000000000000000000001
func encodeParameters(inputArgTypes, inputArgData []string) ([]byte, error) {
	var theTypes abi.Arguments
	var theArgData []interface{}
	theTypes, theArgData, err := buildArgumentAndData(inputArgTypes, inputArgData)
	if err != nil {
		return nil, fmt.Errorf("buildArgumentAndData fail: %s", err)
	}
	bytes, err := theTypes.Pack(theArgData...)
	if err != nil {
		return nil, fmt.Errorf("pack fail: %s", err)
	}
	return bytes, nil
}

func buildArgumentAndData(inputArgTypes, inputArgData []string) (abi.Arguments, []interface{}, error) {
	var theTypes abi.Arguments
	var theArgData []interface{}
	for index, inputType := range inputArgTypes {
		typ, err := abi.NewType(typeNormalize(inputType), "", nil)
		if err != nil {
			return nil, nil, fmt.Errorf("abi.NewType fail: %w", err)
		}
		theTypes = append(theTypes, abi.Argument{Type: typ})
		if strings.Count(inputType, "[") == 1 && strings.Count(inputType, "]") == 1 { // handle array type
			var arrayElementType string
			leftParenthesisLoc := strings.Index(inputType, "[")
			arrayElementType = inputType[:leftParenthesisLoc] // remove all characters from char '['

			var arrayOfTypes, arrayOfData []string
			arrayOfData, err := parseArrayData(inputArgData[index])
			if err != nil {
				return nil, nil, fmt.Errorf("parseArrayData fail: %w", err)
			}
			for _ = range arrayOfData {
				arrayOfTypes = append(arrayOfTypes, typeNormalize(arrayElementType)) // `address[3]`  -> `[address, address, address]`
			}

			_, datas, err := buildArgumentAndData(arrayOfTypes, arrayOfData)
			if err != nil {
				return nil, nil, fmt.Errorf("buildArgumentAndData fail: %w", err)
			}
			if arrayElementType == "string" {
				// datas ([]interface {})   --->  elementData ([]string)
				var elementData []string
				for _, data := range datas {
					elementData = append(elementData, data.(string))
				}
				theArgData = append(theArgData, elementData)
			} else if arrayElementType == "int8" {
				// datas ([]interface {})   --->  elementData ([]int8)
				var elementData []int8
				for _, data := range datas {
					elementData = append(elementData, data.(int8))
				}
				theArgData = append(theArgData, elementData)
			} else if arrayElementType == "int16" {
				var elementData []int16
				for _, data := range datas {
					elementData = append(elementData, data.(int16))
				}
				theArgData = append(theArgData, elementData)
			} else if arrayElementType == "int32" {
				var elementData []int32
				for _, data := range datas {
					elementData = append(elementData, data.(int32))
				}
				theArgData = append(theArgData, elementData)
			} else if arrayElementType == "int64" {
				var elementData []int64
				for _, data := range datas {
					elementData = append(elementData, data.(int64))
				}
				theArgData = append(theArgData, elementData)
			} else if arrayElementType == "uint8" {
				var elementData []uint8
				for _, data := range datas {
					elementData = append(elementData, data.(uint8))
				}
				theArgData = append(theArgData, elementData)
			} else if arrayElementType == "uint16" {
				var elementData []uint16
				for _, data := range datas {
					elementData = append(elementData, data.(uint16))
				}
				theArgData = append(theArgData, elementData)
			} else if arrayElementType == "uint32" {
				var elementData []uint32
				for _, data := range datas {
					elementData = append(elementData, data.(uint32))
				}
				theArgData = append(theArgData, elementData)
			} else if arrayElementType == "uint64" {
				var elementData []uint64
				for _, data := range datas {
					elementData = append(elementData, data.(uint64))
				}
				theArgData = append(theArgData, elementData)
			} else if strings.Contains(arrayElementType, "int") {
				var elementData []*big.Int
				for _, data := range datas {
					elementData = append(elementData, data.(*big.Int))
				}
				theArgData = append(theArgData, elementData)
			} else if arrayElementType == "bool" {
				var elementData []bool
				for _, data := range datas {
					elementData = append(elementData, data.(bool))
				}
				theArgData = append(theArgData, elementData)
			} else if arrayElementType == "address" {
				var elementData []common.Address
				for _, data := range datas {
					elementData = append(elementData, data.(common.Address))
				}
				theArgData = append(theArgData, elementData)
			} else if arrayElementType == "bytes" {
				var elementData [][]byte
				for _, data := range datas {
					elementData = append(elementData, data.([]byte))
				}
				theArgData = append(theArgData, elementData)
			} else if arrayElementType == "bytes1" {
				var elementData [][1]byte
				for _, data := range datas {
					elementData = append(elementData, data.([1]byte))
				}
				theArgData = append(theArgData, elementData)
			} else if arrayElementType == "bytes2" {
				var elementData [][2]byte
				for _, data := range datas {
					elementData = append(elementData, data.([2]byte))
				}
				theArgData = append(theArgData, elementData)
			} else if arrayElementType == "bytes3" {
				var elementData [][3]byte
				for _, data := range datas {
					elementData = append(elementData, data.([3]byte))
				}
				theArgData = append(theArgData, elementData)
			} else if arrayElementType == "bytes4" {
				var elementData [][4]byte
				for _, data := range datas {
					elementData = append(elementData, data.([4]byte))
				}
				theArgData = append(theArgData, elementData)
			} else if arrayElementType == "bytes5" {
				var elementData [][5]byte
				for _, data := range datas {
					elementData = append(elementData, data.([5]byte))
				}
				theArgData = append(theArgData, elementData)
			} else if arrayElementType == "bytes6" {
				var elementData [][6]byte
				for _, data := range datas {
					elementData = append(elementData, data.([6]byte))
				}
				theArgData = append(theArgData, elementData)
			} else if arrayElementType == "bytes7" {
				var elementData [][7]byte
				for _, data := range datas {
					elementData = append(elementData, data.([7]byte))
				}
				theArgData = append(theArgData, elementData)
			} else if arrayElementType == "bytes8" {
				var elementData [][8]byte
				for _, data := range datas {
					elementData = append(elementData, data.([8]byte))
				}
				theArgData = append(theArgData, elementData)
			} else if arrayElementType == "bytes9" {
				var elementData [][9]byte
				for _, data := range datas {
					elementData = append(elementData, data.([9]byte))
				}
				theArgData = append(theArgData, elementData)
			} else if arrayElementType == "bytes10" {
				var elementData [][10]byte
				for _, data := range datas {
					elementData = append(elementData, data.([10]byte))
				}
				theArgData = append(theArgData, elementData)
			} else if arrayElementType == "bytes11" {
				var elementData [][11]byte
				for _, data := range datas {
					elementData = append(elementData, data.([11]byte))
				}
				theArgData = append(theArgData, elementData)
			} else if arrayElementType == "bytes12" {
				var elementData [][12]byte
				for _, data := range datas {
					elementData = append(elementData, data.([12]byte))
				}
				theArgData = append(theArgData, elementData)
			} else if arrayElementType == "bytes13" {
				var elementData [][13]byte
				for _, data := range datas {
					elementData = append(elementData, data.([13]byte))
				}
				theArgData = append(theArgData, elementData)
			} else if arrayElementType == "bytes14" {
				var elementData [][14]byte
				for _, data := range datas {
					elementData = append(elementData, data.([14]byte))
				}
				theArgData = append(theArgData, elementData)
			} else if arrayElementType == "bytes15" {
				var elementData [][15]byte
				for _, data := range datas {
					elementData = append(elementData, data.([15]byte))
				}
				theArgData = append(theArgData, elementData)
			} else if arrayElementType == "bytes16" {
				var elementData [][16]byte
				for _, data := range datas {
					elementData = append(elementData, data.([16]byte))
				}
				theArgData = append(theArgData, elementData)
			} else if arrayElementType == "bytes17" {
				var elementData [][17]byte
				for _, data := range datas {
					elementData = append(elementData, data.([17]byte))
				}
				theArgData = append(theArgData, elementData)
			} else if arrayElementType == "bytes18" {
				var elementData [][18]byte
				for _, data := range datas {
					elementData = append(elementData, data.([18]byte))
				}
				theArgData = append(theArgData, elementData)
			} else if arrayElementType == "bytes19" {
				var elementData [][19]byte
				for _, data := range datas {
					elementData = append(elementData, data.([19]byte))
				}
				theArgData = append(theArgData, elementData)
			} else if arrayElementType == "bytes20" {
				var elementData [][20]byte
				for _, data := range datas {
					elementData = append(elementData, data.([20]byte))
				}
				theArgData = append(theArgData, elementData)
			} else if arrayElementType == "bytes21" {
				var elementData [][21]byte
				for _, data := range datas {
					elementData = append(elementData, data.([21]byte))
				}
				theArgData = append(theArgData, elementData)
			} else if arrayElementType == "bytes22" {
				var elementData [][22]byte
				for _, data := range datas {
					elementData = append(elementData, data.([22]byte))
				}
				theArgData = append(theArgData, elementData)
			} else if arrayElementType == "bytes23" {
				var elementData [][23]byte
				for _, data := range datas {
					elementData = append(elementData, data.([23]byte))
				}
				theArgData = append(theArgData, elementData)
			} else if arrayElementType == "bytes24" {
				var elementData [][24]byte
				for _, data := range datas {
					elementData = append(elementData, data.([24]byte))
				}
				theArgData = append(theArgData, elementData)
			} else if arrayElementType == "bytes25" {
				var elementData [][25]byte
				for _, data := range datas {
					elementData = append(elementData, data.([25]byte))
				}
				theArgData = append(theArgData, elementData)
			} else if arrayElementType == "bytes26" {
				var elementData [][26]byte
				for _, data := range datas {
					elementData = append(elementData, data.([26]byte))
				}
				theArgData = append(theArgData, elementData)
			} else if arrayElementType == "bytes27" {
				var elementData [][27]byte
				for _, data := range datas {
					elementData = append(elementData, data.([27]byte))
				}
				theArgData = append(theArgData, elementData)
			} else if arrayElementType == "bytes28" {
				var elementData [][28]byte
				for _, data := range datas {
					elementData = append(elementData, data.([28]byte))
				}
				theArgData = append(theArgData, elementData)
			} else if arrayElementType == "bytes29" {
				var elementData [][29]byte
				for _, data := range datas {
					elementData = append(elementData, data.([29]byte))
				}
				theArgData = append(theArgData, elementData)
			} else if arrayElementType == "bytes30" {
				var elementData [][30]byte
				for _, data := range datas {
					elementData = append(elementData, data.([30]byte))
				}
				theArgData = append(theArgData, elementData)
			} else if arrayElementType == "bytes31" {
				var elementData [][31]byte
				for _, data := range datas {
					elementData = append(elementData, data.([31]byte))
				}
				theArgData = append(theArgData, elementData)
			} else if arrayElementType == "bytes32" {
				var elementData [][32]byte
				for _, data := range datas {
					elementData = append(elementData, data.([32]byte))
				}
				theArgData = append(theArgData, elementData)
			} else {
				return nil, nil, fmt.Errorf("type %v not implemented in array type currently", inputType)
			}
		} else if inputType == "string" {
			theArgData = append(theArgData, inputArgData[index])
		} else if inputType == "int8" {
			i, err := strconv.ParseInt(inputArgData[index], 10, 64)
			if err != nil {
				return nil, nil, fmt.Errorf("arg (position %v) invalid, %s cannot covert to type %v", index, inputArgData[index], inputType)
			}
			theArgData = append(theArgData, int8(i))
		} else if inputType == "int16" {
			i, err := strconv.ParseInt(inputArgData[index], 10, 64)
			if err != nil {
				return nil, nil, fmt.Errorf("arg (position %v) invalid, %s cannot covert to type %v", index, inputArgData[index], inputType)
			}
			theArgData = append(theArgData, int16(i))
		} else if inputType == "int32" {
			i, err := strconv.ParseInt(inputArgData[index], 10, 64)
			if err != nil {
				return nil, nil, fmt.Errorf("arg (position %v) invalid, %s cannot covert to type %v", index, inputArgData[index], inputType)
			}
			theArgData = append(theArgData, int32(i))
		} else if inputType == "int64" {
			i, err := strconv.ParseInt(inputArgData[index], 10, 64)
			if err != nil {
				return nil, nil, fmt.Errorf("arg (position %v) invalid, %s cannot covert to type %v", index, inputArgData[index], inputType)
			}
			theArgData = append(theArgData, int64(i))
		} else if inputType == "uint8" {
			i, err := strconv.ParseInt(inputArgData[index], 10, 64)
			if err != nil {
				return nil, nil, fmt.Errorf("arg (position %v) invalid, %s cannot covert to type %v", index, inputArgData[index], inputType)
			}
			theArgData = append(theArgData, uint8(i))
		} else if inputType == "uint16" {
			i, err := strconv.ParseInt(inputArgData[index], 10, 64)
			if err != nil {
				return nil, nil, fmt.Errorf("arg (position %v) invalid, %s cannot covert to type %v", index, inputArgData[index], inputType)
			}
			theArgData = append(theArgData, uint16(i))
		} else if inputType == "uint32" {
			i, err := strconv.ParseInt(inputArgData[index], 10, 64)
			if err != nil {
				return nil, nil, fmt.Errorf("arg (position %v) invalid, %s cannot covert to type %v", index, inputArgData[index], inputType)
			}
			theArgData = append(theArgData, uint32(i))
		} else if inputType == "uint64" {
			i, err := strconv.ParseUint(inputArgData[index], 10, 64)
			if err != nil {
				return nil, nil, fmt.Errorf("arg (position %v) invalid, %s cannot covert to type %v", index, inputArgData[index], inputType)
			}
			theArgData = append(theArgData, uint64(i))
		} else if strings.Contains(inputType, "int") { // other cases: int24, int40, ..., int256, uint24, uint40, ..., uint256, etc
			argData := inputArgData[index]

			if !isValidInt(inputType) {
				return nil, nil, fmt.Errorf("type %v not a valid type", inputType)
			}

			if (inputType == "uint256" || inputType == "uint") && strings.Contains(argData, "e") {
				// example:
				// convert 1e18 to 1000000000000000000
				argData, err = scientificNotation2Decimal(argData)
				checkErr(err)
			}

			n := new(big.Int)
			n, ok := n.SetString(argData, 10)
			if !ok {
				return nil, nil, fmt.Errorf("arg (position %v) invalid, %s cannot covert to type %v", index, argData, inputType)
			}
			theArgData = append(theArgData, n)

		} else if inputType == "bool" {
			if strings.EqualFold(inputArgData[index], "true") {
				theArgData = append(theArgData, true)
			} else if strings.EqualFold(inputArgData[index], "false") {
				theArgData = append(theArgData, false)
			} else {
				return nil, nil, fmt.Errorf("arg (position %v) invalid, %s cannot covert to type %v", index, inputArgData[index], inputType)
			}
		} else if inputType == "address" {
			theArgData = append(theArgData, common.HexToAddress(inputArgData[index]))
		} else if inputType == "bytes" {
			var inputHex = inputArgData[index]
			if strings.HasPrefix(inputArgData[index], "0x") {
				inputHex = inputArgData[index][2:]
			}
			decoded, err := hex.DecodeString(inputHex)
			if err != nil {
				return nil, nil, fmt.Errorf("arg (position %v) invalid, %s cannot covert to type %v", index, inputArgData[index], inputType)
			}
			theArgData = append(theArgData, decoded)
		} else if strings.Contains(inputType, "bytes") { // bytes1, bytes2, ..., bytes32
			var inputHex = inputArgData[index]
			if strings.HasPrefix(inputArgData[index], "0x") {
				inputHex = inputArgData[index][2:]
			}
			decoded, err := hex.DecodeString(inputHex)
			if err != nil {
				return nil, nil, fmt.Errorf("arg (position %v) invalid, %s cannot covert to type %v", index, inputArgData[index], inputType)
			}
			if inputType == "bytes1" {
				var data [1]byte
				copy(data[:], decoded)
				theArgData = append(theArgData, data)
			} else if inputType == "bytes2" {
				var data [2]byte
				copy(data[:], decoded)
				theArgData = append(theArgData, data)
			} else if inputType == "bytes3" {
				var data [3]byte
				copy(data[:], decoded)
				theArgData = append(theArgData, data)
			} else if inputType == "bytes4" {
				var data [4]byte
				copy(data[:], decoded)
				theArgData = append(theArgData, data)
			} else if inputType == "bytes5" {
				var data [5]byte
				copy(data[:], decoded)
				theArgData = append(theArgData, data)
			} else if inputType == "bytes6" {
				var data [6]byte
				copy(data[:], decoded)
				theArgData = append(theArgData, data)
			} else if inputType == "bytes7" {
				var data [7]byte
				copy(data[:], decoded)
				theArgData = append(theArgData, data)
			} else if inputType == "bytes8" {
				var data [8]byte
				copy(data[:], decoded)
				theArgData = append(theArgData, data)
			} else if inputType == "bytes9" {
				var data [9]byte
				copy(data[:], decoded)
				theArgData = append(theArgData, data)
			} else if inputType == "bytes10" {
				var data [10]byte
				copy(data[:], decoded)
				theArgData = append(theArgData, data)
			} else if inputType == "bytes11" {
				var data [11]byte
				copy(data[:], decoded)
				theArgData = append(theArgData, data)
			} else if inputType == "bytes12" {
				var data [12]byte
				copy(data[:], decoded)
				theArgData = append(theArgData, data)
			} else if inputType == "bytes13" {
				var data [13]byte
				copy(data[:], decoded)
				theArgData = append(theArgData, data)
			} else if inputType == "bytes14" {
				var data [14]byte
				copy(data[:], decoded)
				theArgData = append(theArgData, data)
			} else if inputType == "bytes15" {
				var data [15]byte
				copy(data[:], decoded)
				theArgData = append(theArgData, data)
			} else if inputType == "bytes16" {
				var data [16]byte
				copy(data[:], decoded)
				theArgData = append(theArgData, data)
			} else if inputType == "bytes17" {
				var data [17]byte
				copy(data[:], decoded)
				theArgData = append(theArgData, data)
			} else if inputType == "bytes18" {
				var data [18]byte
				copy(data[:], decoded)
				theArgData = append(theArgData, data)
			} else if inputType == "bytes19" {
				var data [19]byte
				copy(data[:], decoded)
				theArgData = append(theArgData, data)
			} else if inputType == "bytes20" {
				var data [20]byte
				copy(data[:], decoded)
				theArgData = append(theArgData, data)
			} else if inputType == "bytes21" {
				var data [21]byte
				copy(data[:], decoded)
				theArgData = append(theArgData, data)
			} else if inputType == "bytes22" {
				var data [22]byte
				copy(data[:], decoded)
				theArgData = append(theArgData, data)
			} else if inputType == "bytes23" {
				var data [23]byte
				copy(data[:], decoded)
				theArgData = append(theArgData, data)
			} else if inputType == "bytes24" {
				var data [24]byte
				copy(data[:], decoded)
				theArgData = append(theArgData, data)
			} else if inputType == "bytes25" {
				var data [25]byte
				copy(data[:], decoded)
				theArgData = append(theArgData, data)
			} else if inputType == "bytes26" {
				var data [26]byte
				copy(data[:], decoded)
				theArgData = append(theArgData, data)
			} else if inputType == "bytes27" {
				var data [27]byte
				copy(data[:], decoded)
				theArgData = append(theArgData, data)
			} else if inputType == "bytes28" {
				var data [28]byte
				copy(data[:], decoded)
				theArgData = append(theArgData, data)
			} else if inputType == "bytes29" {
				var data [29]byte
				copy(data[:], decoded)
				theArgData = append(theArgData, data)
			} else if inputType == "bytes30" {
				var data [30]byte
				copy(data[:], decoded)
				theArgData = append(theArgData, data)
			} else if inputType == "bytes31" {
				var data [31]byte
				copy(data[:], decoded)
				theArgData = append(theArgData, data)
			} else if inputType == "bytes32" {
				var data [32]byte
				copy(data[:], decoded)
				theArgData = append(theArgData, data)
			} else {
				return nil, nil, fmt.Errorf("type %v not implemented currently", inputType)
			}
		} else {
			return nil, nil, fmt.Errorf("type %v not implemented currently", inputType)
		}
	}

	return theTypes, theArgData, nil
}

// `["abc", "xyz"]`     ----> abc, xyz
// `[abc, xyz]`         ----> abc, xyz
// `["a(a,a)", "abcd"]` ----> a(a,a), abcd
// `abc, xyz`           ----> abc, xyz
func parseArrayData(input string) ([]string, error) {
	input = strings.TrimSpace(input)
	input = strings.TrimLeft(input, "[")
	input = strings.TrimRight(input, "]")
	r := csv.NewReader(strings.NewReader(input))
	r.LazyQuotes = true
	r.TrimLeadingSpace = true
	records, err := r.Read()
	if err != nil {
		if err == io.EOF {
			return []string{}, nil
		} else {
			println(err.Error())
		}
	}
	return records, nil
}

// uint -> uint256
// int -> int256
// uint[] -> uint256[]
// int[] -> int256[]
func typeNormalize(input string) string {
	re := regexp.MustCompile(`\b([u]int)\b`)
	return re.ReplaceAllString(input, "${1}256")
}

// ABI example:
// [
//  {
//      "inputs": [],
//      "stateMutability": "nonpayable",
//      "type": "constructor"
//  },
//	{
//		"inputs": [
//			{
//				"internalType": "uint256[]",
//				"name": "_a",
//				"type": "uint256[]"
//			},
//			{
//				"internalType": "address[]",
//				"name": "_addr",
//				"type": "address[]"
//			}
//		],
//		"name": "f1",
//		"outputs": [],
//		"stateMutability": "nonpayable",
//		"type": "function"
//	},
//	{
//		"inputs": [],
//		"name": "f2",
//		"outputs": [
//			{
//				"internalType": "uint256",
//				"name": "",
//				"type": "uint256"
//			}
//		],
//		"stateMutability": "view",
//		"type": "function"
//	},
// ......
// ]
type AbiData struct {
	Inputs []struct {
		Name string `json:"name"`
		Type string `json:"type"`
	} `json:"inputs"`
	Name    string `json:"name"`
	Type    string `json:"type"` // constructor, function, etc.
	Outputs []struct {
		Name string `json:"name"`
		Type string `json:"type"`
	} `json:"outputs"`
}

type AbiJSONData struct {
	ABI []AbiData `json:"abi"`
}

func ExtractFuncDefinition(abi string, funcName string) (string, error) {
	// log.Printf("abi = %s\nfuncName = %s", abi, funcName)
	abi = strings.TrimSpace(abi)
	if len(abi) == 0 {
		return "", fmt.Errorf("abi is empty")
	}

	var parsedABI []AbiData

	if abi[0:1] == "[" {
		if err := json.Unmarshal([]byte(abi), &parsedABI); err != nil {
			return "", fmt.Errorf("unmarshal fail: %w", err)
		}
	} else if abi[0:1] == "{" {
		var abiJSONData AbiJSONData
		if err := json.Unmarshal([]byte(abi), &abiJSONData); err != nil {
			return "", fmt.Errorf("unmarshal fail: %w", err)
		}
		parsedABI = abiJSONData.ABI
	} else {
		return "", fmt.Errorf("abi invalid")
	}

	var ret = funcName + "("

	if len(parsedABI) == 0 {
		return "", fmt.Errorf("parsedABI is empty")
	}

	var foundFunc = false
	for _, item := range parsedABI {
		if funcName == "constructor" { // constructor
			if item.Type == "constructor" {
				foundFunc = true
			}
		} else { // normal function
			if item.Type == "function" && item.Name == funcName {
				foundFunc = true
			}
		}
		if foundFunc == true {
			for index, input := range item.Inputs {
				ret += input.Type

				if index < len(item.Inputs)-1 { // not the last input
					ret += ", "
				}
			}

			ret += ")"

			if len(item.Outputs) > 0 {
				ret += " returns ("
				for index, output := range item.Outputs {
					ret += output.Type

					if index < len(item.Outputs)-1 { // not the last input
						ret += ", "
					}
				}

				ret += ")"
			}

			break
		}
	}

	if !foundFunc {
		return "", fmt.Errorf("function %v not found in ABI", funcName)
	}

	// Example of ret: `f1(uint256[], address[]) returns (uint256)`
	return ret, nil
}

// isValidInt return true if intType is valid solidity int type
func isValidInt(intType string) bool {
	switch intType {
	case
		"int",
		"int8",
		"int16",
		"int24",
		"int32",
		"int40",
		"int48",
		"int56",
		"int64",
		"int72",
		"int80",
		"int88",
		"int96",
		"int104",
		"int112",
		"int120",
		"int128",
		"int136",
		"int144",
		"int152",
		"int160",
		"int168",
		"int176",
		"int184",
		"int192",
		"int200",
		"int208",
		"int216",
		"int224",
		"int232",
		"int240",
		"int248",
		"int256",
		"uint",
		"uint8",
		"uint16",
		"uint24",
		"uint32",
		"uint40",
		"uint48",
		"uint56",
		"uint64",
		"uint72",
		"uint80",
		"uint88",
		"uint96",
		"uint104",
		"uint112",
		"uint120",
		"uint128",
		"uint136",
		"uint144",
		"uint152",
		"uint160",
		"uint168",
		"uint176",
		"uint184",
		"uint192",
		"uint200",
		"uint208",
		"uint216",
		"uint224",
		"uint232",
		"uint240",
		"uint248",
		"uint256":
		return true
	}
	return false
}

func scientificNotation2Decimal(input string) (string, error) {
	r := regexp.MustCompile(`^([0-9]*)([.]?)([0-9]+)e([0-9]+)$`)
	matches := r.FindStringSubmatch(input)

	part1 := matches[1] // group 1
	part2 := matches[2] // group 2
	part3 := matches[3] // group 3
	part4 := matches[4] // group 4

	part4Int, err := strconv.ParseInt(part4, 10, 64)
	checkErr(err)

	var result = ""
	if part2 == "." {
		// has dot, for example 12.1e3
		if part1 == "0" {
			// for example 0.3e5
			result = part3 + strings.Repeat("0", int(part4Int)-1)
		} else {
			result = part1 + part3 + strings.Repeat("0", int(part4Int)-1)
		}
	} else {
		// no dot
		result = part3 + strings.Repeat("0", int(part4Int))
	}

	log.Printf("convert %v to %v", input, result)
	return result, nil
}

func checkErr(err error)  {
	fmt.Println(err)
}
func ExtractFuncName(input string) string {
	if strings.HasPrefix(input, "function ") {
		input = input[len("function "):] // remove leading string "function "
	}
	funcName := strings.TrimLeft(input, " ")

	leftParenthesisLoc := strings.Index(funcName, "(")
	if leftParenthesisLoc >= 0 { // ( found
		funcName := funcName[:leftParenthesisLoc] // remove all characters from char '('
		funcName = strings.TrimSpace(funcName)
	}
	return funcName
}
// BuildTxInputData build tx input data
func BuildTxInputData(funcSignature string, inputArgData []string) ([]byte, error) {
	funcName, funcArgTypes, err := parseFuncSignature(funcSignature)
	if err != nil {
		return nil, err
	}

	functionSelector := make([]byte, 0)
	if len(funcName) > 0 {
		funcSign := funcName + "(" + strings.Join(funcArgTypes, ",") + ")"
		functionSelector = crypto.Keccak256([]byte(funcSign))[0:4]
	} else {
		// log.Printf("function name is not found, only encode arguments")
	}

	if len(funcArgTypes) != len(inputArgData) {
		return nil, fmt.Errorf("invalid input, there are %v args in signature, but %v args are provided", len(funcArgTypes), len(inputArgData))
	}
	data, err := encodeParameters(funcArgTypes, inputArgData)
	if err != nil {
		return nil, fmt.Errorf("encodeParameters fail: %v", err)
	}

	return append(functionSelector, data...), nil
}

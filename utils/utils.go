package utils

import (
	"encoding/json"
	"fmt"
	"reflect"

	sdk "github.com/cosmos/cosmos-sdk/types"

)

// ParseValAddresses parses validator address string to types.ValAddress
func ParseValAddresses(valAddrsStr []string) ([]sdk.ValAddress, error) {
	valLen := len(valAddrsStr)
	valAddrs := make([]sdk.ValAddress, valLen)
	var err error
	for i := 0; i < valLen; i++ {
		valAddrs[i], err = sdk.ValAddressFromBech32(valAddrsStr[i])
		if err != nil {
			return nil, fmt.Errorf("invalid validator address: %s", valAddrsStr[i])
		}
	}
	return valAddrs, nil
}



func ToJson(f interface{}) string {
	bytes, err := json.Marshal(f)
	Require(err == nil, err)
	return string(bytes)
}
func Require(require bool, msg interface{}) {
	if !require {
		if e, b := msg.(error); b {
			panic(e)
		}
		kind := reflect.TypeOf(msg).Kind()
		if kind == reflect.String {
			panic(fmt.Errorf(msg.(string)))
		}
		if kind == reflect.Func {
			v := reflect.ValueOf(msg).Call(nil)
			if len(v) > 0 && v[0].Type().Kind() == reflect.String {
				fmt.Println(v[0].Interface().(string))
			}
		}
	}
}

package utils
import "C"

func HexToAddr(hexAddr string) string {
	addr,_:=ToCosmosAddress(hexAddr)
	return addr.String()
}
func AddrToHex(addr string) string {
	add,_:=ToHexAddress(addr)
	return add.String()
}

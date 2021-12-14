package utils

import (
	"errors"
	"fmt"

	"log"

	"github.com/bartekn/go-bip39"
	"github.com/miguelmota/go-ethereum-hdwallet"

)

const (
	defaultName         = "alice"
	defaultPassWd       = "12345678"
	mnemonicEntropySize = 128
	defaultCointype     = 60
)

type MnWallet struct {
	HexAddress string
	Address string
	Prikey string
	Mnemo string

}

// CreateAccount creates a random key info with the given name and password
func CreateAccount(name, passWd string) (mnWallet MnWallet, err error) {
	if len(name) == 0 {
		name = defaultName
		log.Printf("Default name: \"%s\"\n", name)
	}

	if len(passWd) == 0 {
		passWd = defaultPassWd
		log.Printf("Default password: \"%s\"\n", passWd)
	}

	mnemo, err := GenerateMnemonic()
	if err != nil {
		return
	}

	wallet,_:=hdwallet.NewFromMnemonic(mnemo)
	path := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/0")
	account, err := wallet.Derive(path, false)
	privateKey, err := wallet.PrivateKeyHex(account)
	bechAddr,_:=ToCosmosAddress(account.Address.Hex())
	st:=struct {
		HexAddress string
		Address string
		Prikey string
		Mnemo string
	}{
		account.Address.Hex(),
		bechAddr.String(),
		privateKey,
		mnemo,
	}
	return st,nil
}





func ImportWallet(mn string) (mnWallet MnWallet, err error) {
	var name=""
	var passWd=""
	if len(name) == 0 {
		name = defaultName
		log.Printf("Default name: \"%s\"\n", name)
	}

	if len(passWd) == 0 {
		passWd = defaultPassWd
		log.Printf("Default password: \"%s\"\n", passWd)
	}
	wallet,_:=hdwallet.NewFromMnemonic(mn)
	path := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/0")
	account, err := wallet.Derive(path, false)
	privateKey, err := wallet.PrivateKeyHex(account)
	bechAddr,_:=ToCosmosAddress(account.Address.Hex())
	st:=struct {
		HexAddress string
		Address string
		Prikey string
		Mnemo string
	}{
		account.Address.Hex(),
		bechAddr.String(),
		privateKey,
		mn,
	}
	return st,nil
}



// GenerateMnemonic creates a random mnemonic
func GenerateMnemonic() (mnemo string, err error) {
	entropySeed, err := bip39.NewEntropy(mnemonicEntropySize)
	if err != nil {
		return mnemo, fmt.Errorf("failed. bip39.NewEntropy err : %s", err.Error())
	}
	//seed:="8cc480a661c7ad3c0fc19158f6b4b6cb6a59e57f5ebd184a1b7c576e02a394205e9f5d971fc537aa7f0de6ebc9887b55edd55717fd2393ba1f783457815b11f5"
	mnemo, err = bip39.NewMnemonic( []byte(entropySeed)[:])
	if err != nil {
		return mnemo, fmt.Errorf("failed. bip39.NewMnemonic err : %s", err.Error())
	}

	log.Printf("New mnemonic: \"%s\". Be sure to remember that!\n", mnemo)

	return
}

// GeneratePrivateKeyFromMnemo converts mnemonic to private key
func GeneratePrivateKeyFromMnemo(mnemonic string) (privKey string, err error) {
	if len(mnemonic) == 0 {
		return privKey, errors.New("failed. no mnemonic input")
	}

	if !bip39.IsMnemonicValid(mnemonic) {
		return privKey, errors.New("failed. mnemonic is invalid")
	}

	wallet,_:=hdwallet.NewFromMnemonic(mnemonic)
	path := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/0")
	account, err := wallet.Derive(path, false)
	privateKey, err := wallet.PrivateKeyHex(account)
	return privateKey,nil
}

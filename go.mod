module dechain-go-sdk

go 1.17

require (
	gitee.com/aifuturewell/gojni v0.0.0-20210925002626-b8dffd1c1d72
	github.com/bartekn/go-bip39 v0.0.0-20171116152956-a05967ea095d
	github.com/cosmos/cosmos-sdk v0.44.5
	github.com/ethereum/go-ethereum v1.10.8
	github.com/miguelmota/go-ethereum-hdwallet v0.1.1
	github.com/shopspring/decimal v1.2.0
	github.com/stretchr/testify v1.7.0
	github.com/tendermint/tendermint v0.34.14

)



replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1

// TODO: remove once v0.45 has been released
replace github.com/cosmos/cosmos-sdk => github.com/tharsis/cosmos-sdk v0.44.3-olympus

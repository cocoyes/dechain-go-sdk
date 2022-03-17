package cosmos

import (
	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/simapp/helpers"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/crypto/secp256k1"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	"testing"
)

func TestStaking(t *testing.T) {
	pri := secp256k1.GenPrivKeySecp256k1([]byte(""))
	bondTokens := sdk.TokensFromConsensusPower(10, sdk.NewIntFromUint64(1000000000000000000))
	bondCoin := sdk.NewCoin("", bondTokens)
	valAddr, _ := sdk.ValAddressFromBech32("")
	addr, _ := sdk.AccAddressFromBech32("")
	delegateMsg := types.NewMsgDelegate(addr, valAddr, bondCoin)
	app := simapp.Setup(true)
	header := tmproto.Header{Height: app.LastBlockHeight() + 1}
	txGen := simapp.MakeTestEncodingConfig().TxConfig

	tx, err := helpers.GenTx(
		txGen,
		[]sdk.Msg{delegateMsg},
		sdk.Coins{sdk.NewInt64Coin(sdk.DefaultBondDenom, 0)},
		helpers.DefaultGenTxGas,
		chainID,
		accNums,
		accSeqs,
		pri...,
	)

	txBytes, err := cdc.MarshalBinaryLengthPrefixed(tx)
	require.Nil(t, err)

	// Must simulate now as CheckTx doesn't run Msgs anymore
	_, res, err := app.Simulate(txBytes, tx)

	if expSimPass {
		require.NoError(t, err)
		require.NotNil(t, res)
	} else {
		require.Error(t, err)
		require.Nil(t, res)
	}

	// Simulate a sending a transaction and committing a block
	app.BeginBlock(abci.RequestBeginBlock{Header: header})
	gInfo, res, err := app.Deliver(tx)

	if expPass {
		require.NoError(t, err)
		require.NotNil(t, res)
	} else {
		require.Error(t, err)
		require.Nil(t, res)
	}

	app.EndBlock(abci.RequestEndBlock{})
	app.Commit()

}

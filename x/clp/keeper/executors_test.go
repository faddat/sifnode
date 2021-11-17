package keeper_test

import (
	"fmt"
	"testing"

	"github.com/Sifchain/sifnode/x/clp/test"
	"github.com/Sifchain/sifnode/x/clp/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestKeeper_CreatePool_And_AddLiquidity_RemoveLiquidity(t *testing.T) {
	// nativeAssetAmount sdk.Uint, externalAssetAmount
	ctx, app := test.CreateTestAppClp(false)
	signer := test.GenerateAddress(test.AddressKey1)
	signer2 := test.GenerateAddress("")
	//Parameters for create pool
	nativeAssetAmount := sdk.NewUintFromString("998")
	externalAssetAmount := sdk.NewUintFromString("998")
	// nativeAssetAmount2 := sdk.NewUintFromString("")
	// externalAssetAmount2 := sdk.NewUintFromString("")
	asset := types.NewAsset("eth")
	asset2 := types.NewAsset("xxx")
	externalCoin := sdk.NewCoin(asset.Symbol, sdk.Int(sdk.NewUint(10000)))
	nativeCoin := sdk.NewCoin(types.NativeSymbol, sdk.Int(sdk.NewUint(10000)))
	_ = app.ClpKeeper.GetBankKeeper().AddCoins(ctx, signer, sdk.NewCoins(externalCoin, nativeCoin))

	msgCreatePool := types.NewMsgCreatePool(nil, asset, nativeAssetAmount, externalAssetAmount)
	// Create Pool
	pool, err := app.ClpKeeper.CreatePool(ctx, sdk.NewUint(1), &msgCreatePool)
	assert.Error(t, err, "empty address string is not allowed")

	/*
		msgCreatePool = types.NewMsgCreatePool(signer, asset, nativeAssetAmount2, externalAssetAmount)
		// Create Pool
		pool, err = app.ClpKeeper.CreatePool(ctx, sdk.NewUint(1), &msgCreatePool)
		assert.Error(t, err, "Unable to parse to Int")

		msgCreatePool = types.NewMsgCreatePool(signer, asset, nativeAssetAmount, externalAssetAmount2)
		// Create Pool
		pool, err = app.ClpKeeper.CreatePool(ctx, sdk.NewUint(1), &msgCreatePool)
		assert.Error(t, err, "Unable to parse to Int")
	*/

	msgCreatePool = types.NewMsgCreatePool(signer2, asset2, nativeAssetAmount, externalAssetAmount)
	// Create Pool
	pool, err = app.ClpKeeper.CreatePool(ctx, sdk.NewUint(1), &msgCreatePool)
	assert.Error(t, err, "user does not have enough balance of the required coin")

	msgCreatePool = types.NewMsgCreatePool(signer, asset, sdk.NewUintFromString("99800000000"), sdk.NewUintFromString("998000000"))
	// Create Pool
	pool, err = app.ClpKeeper.CreatePool(ctx, sdk.NewUint(1), &msgCreatePool)
	assert.Error(t, err, "user does not have enough balance of the required coin")

	msgCreatePool = types.NewMsgCreatePool(signer, asset, nativeAssetAmount, externalAssetAmount)
	// Create Pool
	pool, err = app.ClpKeeper.CreatePool(ctx, sdk.NewUint(1), &msgCreatePool)
	if err != nil {
		fmt.Println("Error Generating new pool :", err)
	}
	_, err = app.ClpKeeper.CreatePool(ctx, sdk.NewUint(1), nil)
	assert.Error(t, err, "MsgCreatePool can not be nil")
	msg := types.NewMsgAddLiquidity(signer, asset, nativeAssetAmount, externalAssetAmount)
	app.ClpKeeper.CreateLiquidityProvider(ctx, &asset, sdk.NewUint(1), signer)
	lp, err := app.ClpKeeper.AddLiquidity(ctx, &msg, *pool, sdk.NewUint(1), sdk.NewUint(998))
	assert.Equal(t, lp.LiquidityProviderAddress, "sif15ky9du8a2wlstz6fpx3p4mqpjyrm5cgqhns3lt")
	assert.NoError(t, err)
	assert.Equal(t, pool.ExternalAssetBalance, externalAssetAmount)
	assert.Equal(t, pool.NativeAssetBalance, nativeAssetAmount)
	msg = types.NewMsgAddLiquidity(signer2, asset2, nativeAssetAmount, externalAssetAmount)
	_, err = app.ClpKeeper.AddLiquidity(ctx, &msg, *pool, sdk.NewUint(1), sdk.NewUint(998))
	assert.Error(t, err, "insufficient funds")
	msg = types.NewMsgAddLiquidity(signer2, asset2, sdk.NewUintFromString("998"), sdk.NewUintFromString("998"))
	_, err = app.ClpKeeper.AddLiquidity(ctx, &msg, *pool, sdk.NewUint(1), sdk.NewUint(998))
	assert.Error(t, err, "Unable to parse to Int")
	subCoin := sdk.NewCoin(asset.Symbol, sdk.Int(sdk.NewUint(100)))
	errorRemoveLiquidity := app.ClpKeeper.RemoveLiquidity(ctx, *pool, subCoin, subCoin, *lp, sdk.NewUint(989), sdk.NewUint(10001), sdk.NewUint(10001))
	assert.NoError(t, errorRemoveLiquidity)
	ok := app.ClpKeeper.HasBalance(ctx, signer, subCoin)
	assert.True(t, ok, "")
	subCoin = sdk.NewCoin(asset.Symbol, sdk.Int(sdk.NewUint(100)))
	errorRemoveLiquidity = app.ClpKeeper.RemoveLiquidity(ctx, *pool, subCoin, subCoin, *lp, sdk.NewUint(989), sdk.NewUint(10001), sdk.NewUint(10001))
	assert.NoError(t, errorRemoveLiquidity)
	res := app.ClpKeeper.HasBalance(ctx, signer, subCoin)
	assert.True(t, res, "Cannot withdraw pool is too shallow")
	subCoin = sdk.NewCoin(asset.Symbol, sdk.Int(sdk.NewUint(100)))
	errorRemoveLiquidity = app.ClpKeeper.RemoveLiquidity(ctx, *pool, subCoin, subCoin, *lp, sdk.NewUint(0), sdk.NewUint(10001), sdk.NewUint(10001))
	assert.NoError(t, errorRemoveLiquidity)
}

func TestKeeper_CreateLiquidityProvider(t *testing.T) {
	ctx, app := test.CreateTestAppClp(false)
	asset := types.NewAsset("eth")
	lpAddess, err := sdk.AccAddressFromBech32("sif1azpar20ck9lpys89r8x7zc8yu0qzgvtp48ng5v")
	if err != nil {
		fmt.Println("Error Creating Liquidity Provider :", err)
	}
	lp := app.ClpKeeper.CreateLiquidityProvider(ctx, &asset, sdk.NewUint(1), lpAddess)
	assert.NoError(t, err)
	assert.Equal(t, lp.LiquidityProviderAddress, "sif1azpar20ck9lpys89r8x7zc8yu0qzgvtp48ng5v")
}

func TestKeeper_RemoveLiquidityProvider(t *testing.T) {
	ctx, app := test.CreateTestAppClp(false)
	signer := test.GenerateAddress(test.AddressKey1)
	//Parameters for create pool
	nativeAssetAmount := sdk.NewUintFromString("998")
	externalAssetAmount := sdk.NewUintFromString("998")
	asset := types.NewAsset("eth")
	asset2 := types.NewAsset("xxx")
	externalCoin := sdk.NewCoin(asset.Symbol, sdk.Int(sdk.NewUint(10000)))
	nativeCoin := sdk.NewCoin(types.NativeSymbol, sdk.Int(sdk.NewUint(10000)))
	subCoin := sdk.NewUintFromString("1")
	newAssetCoin := sdk.NewCoin(asset.Symbol, sdk.Int(subCoin))
	newAssetCoin2 := sdk.NewCoin(asset2.Symbol, sdk.Int(subCoin))
	_ = app.ClpKeeper.GetBankKeeper().AddCoins(ctx, signer, sdk.NewCoins(externalCoin, nativeCoin))
	msgCreatePool := types.NewMsgCreatePool(signer, asset, nativeAssetAmount, externalAssetAmount)
	// Create Pool
	pool, _ := app.ClpKeeper.CreatePool(ctx, sdk.NewUint(1), &msgCreatePool)
	msg := types.NewMsgAddLiquidity(signer, asset, nativeAssetAmount, externalAssetAmount)
	app.ClpKeeper.CreateLiquidityProvider(ctx, &asset, sdk.NewUint(1), signer)
	lp, _ := app.ClpKeeper.AddLiquidity(ctx, &msg, *pool, sdk.NewUint(1), sdk.NewUint(998))
	getlp, _ := app.ClpKeeper.GetLiquidityProvider(ctx, lp.Asset.Symbol, lp.LiquidityProviderAddress)
	assert.True(t, app.ClpKeeper.GetLiquidityProviderIterator(ctx).Valid())
	app.ClpKeeper.DestroyLiquidityProvider(ctx, lp.Asset.Symbol, lp.LiquidityProviderAddress)
	_, err := app.ClpKeeper.GetLiquidityProvider(ctx, lp.Asset.Symbol, lp.LiquidityProviderAddress)
	assert.Error(t, err, "LiquidityProvider has been deleted")
	err = app.ClpKeeper.RemoveLiquidityProvider(ctx, sdk.Coins{newAssetCoin2}.Sort(), getlp)
	assert.Error(t, err, "unable to add balance")
	// This should do nothing
	err = app.ClpKeeper.RemoveLiquidityProvider(ctx, sdk.Coins{newAssetCoin}.Sort(), getlp)
	assert.NoError(t, err)
	assert.False(t, app.ClpKeeper.GetLiquidityProviderIterator(ctx).Valid())
	msg = types.NewMsgAddLiquidity(signer, asset, nativeAssetAmount, externalAssetAmount)
	_, err = app.ClpKeeper.AddLiquidity(ctx, &msg, *pool, sdk.NewUint(1), sdk.NewUint(998))
	assert.NoError(t, err)
}

func TestKeeper_DecommissionPool(t *testing.T) {

	ctx, app := test.CreateTestAppClp(false)
	signer := test.GenerateAddress(test.AddressKey1)
	//Parameters for create pool
	nativeAssetAmount := sdk.NewUintFromString("998")
	externalAssetAmount := sdk.NewUintFromString("998")
	asset := types.NewAsset("eth")
	externalCoin := sdk.NewCoin(asset.Symbol, sdk.Int(sdk.NewUint(10000)))
	nativeCoin := sdk.NewCoin(types.NativeSymbol, sdk.Int(sdk.NewUint(10000)))
	_ = app.ClpKeeper.GetBankKeeper().AddCoins(ctx, signer, sdk.NewCoins(externalCoin, nativeCoin))
	msgCreatePool := types.NewMsgCreatePool(signer, asset, nativeAssetAmount, externalAssetAmount)
	// Create Pool
	pool, err := app.ClpKeeper.CreatePool(ctx, sdk.NewUint(1), &msgCreatePool)
	if err != nil {
		fmt.Println("Error Generating new pool :", err)
	}

	err = app.ClpKeeper.DecommissionPool(ctx, *pool)
	require.NoError(t, err)
	_, err = app.ClpKeeper.GetPool(ctx, pool.ExternalAsset.Symbol)
	assert.Error(t, err, "Pool should be deleted")
}

func TestKeeper_InitiateSwap(t *testing.T) {
	ctx, app := test.CreateTestAppClp(false)
	signer := test.GenerateAddress(test.AddressKey1)
	//Parameters for create pool
	asset := types.NewAsset("eth")
	externalCoin := sdk.NewCoin(asset.Symbol, sdk.Int(sdk.NewUint(10000)))
	nativeCoin := sdk.NewCoin(types.NativeSymbol, sdk.Int(sdk.NewUint(10000)))
	_ = app.ClpKeeper.GetBankKeeper().AddCoins(ctx, signer, sdk.NewCoins(externalCoin, nativeCoin))
	err := app.ClpKeeper.InitiateSwap(ctx, externalCoin, signer)
	require.NoError(t, err)
	ok := app.ClpKeeper.HasBalance(ctx, signer, nativeCoin)
	assert.True(t, ok, "")

}

func TestKeeper_FinalizeSwap(t *testing.T) {
	ctx, app := test.CreateTestAppClp(false)
	signer := test.GenerateAddress(test.AddressKey1)
	//Parameters for create pool
	nativeAssetAmount := sdk.NewUintFromString("998")
	externalAssetAmount := sdk.NewUintFromString("998")
	assetEth := types.NewAsset("eth")
	assetDash := types.NewAsset("dash")
	externalCoin := sdk.NewCoin(assetEth.Symbol, sdk.Int(sdk.NewUint(10000)))
	nativeCoin := sdk.NewCoin(types.NativeSymbol, sdk.Int(sdk.NewUint(10000)))
	_ = app.ClpKeeper.GetBankKeeper().AddCoins(ctx, signer, sdk.NewCoins(externalCoin, nativeCoin))
	msgCreatePool := types.NewMsgCreatePool(signer, assetEth, nativeAssetAmount, externalAssetAmount)
	// Create Pool
	_, err := app.ClpKeeper.CreatePool(ctx, sdk.NewUint(1), &msgCreatePool)
	if err != nil {
		fmt.Println("Error Generating new pool :", err)
	}
	externalCoin = sdk.NewCoin(assetDash.Symbol, sdk.Int(sdk.NewUint(10000)))
	nativeCoin = sdk.NewCoin(types.NativeSymbol, sdk.Int(sdk.NewUint(10000)))
	_ = app.ClpKeeper.GetBankKeeper().AddCoins(ctx, signer, sdk.NewCoins(externalCoin, nativeCoin))

	msgCreatePool = types.NewMsgCreatePool(signer, assetDash, nativeAssetAmount, externalAssetAmount)
	// Create Pool
	pool, err := app.ClpKeeper.CreatePool(ctx, sdk.NewUint(1), &msgCreatePool)
	if err != nil {
		fmt.Println("Error Generating new pool :", err)
	}
	// Test Parameters for swap
	// initialBalance: Initial account balance for all assets created.
	initialBalance := sdk.NewUintFromString("1000000000000000000000")
	// poolBalance: Amount funded to pool. The same amount is used both for native and external asset.
	externalCoin1 := sdk.NewCoin("eth", sdk.Int(initialBalance))
	externalCoin2 := sdk.NewCoin("dash", sdk.Int(initialBalance))
	// Signer is given ETH and RWN (Signer will creat pool and become LP)
	_ = app.ClpKeeper.GetBankKeeper().AddCoins(ctx, signer, sdk.NewCoins(externalCoin1, nativeCoin))
	_ = app.ClpKeeper.GetBankKeeper().AddCoins(ctx, signer, sdk.NewCoins(externalCoin2, nativeCoin))
	msg := types.NewMsgSwap(signer, assetEth, assetDash, sdk.NewUint(1), sdk.NewUint(10))
	err = app.ClpKeeper.FinalizeSwap(ctx, "", *pool, msg)
	assert.Error(t, err, "Unable to parse to Int")

	err = app.ClpKeeper.FinalizeSwap(ctx, "1", *pool, msg)
	require.NoError(t, err)
}

func TestKeeper_ParseToInt(t *testing.T) {
	_, app := test.CreateTestAppClp(false)
	res, boolean := app.ClpKeeper.ParseToInt("1")
	assert.True(t, boolean)
	assert.Equal(t, res.String(), "1")
}

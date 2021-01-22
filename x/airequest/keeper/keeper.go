package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	params "github.com/cosmos/cosmos-sdk/x/params/types"
	staking "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	"github.com/oraichain/orai/x/airequest/types"
	"github.com/oraichain/orai/x/provider"
	"github.com/oraichain/orai/x/wasm"
	"github.com/tendermint/tendermint/libs/log"
)

// always clone keeper to make it immutable
type (
	Keeper struct {
		cdc            codec.Marshaler
		storeKey       sdk.StoreKey
		wasmKeeper     *wasm.Keeper
		paramSpace     params.Subspace
		stakingKeeper  staking.Keeper
		providerKeeper *provider.Keeper
	}
)

// NewKeeper creates a airequest keeper
func NewKeeper(cdc codec.Marshaler, key sdk.StoreKey, wasmKeeper *wasm.Keeper, aiRequestSubspace params.Subspace, stakingKeeper staking.Keeper, providerKeeper *provider.Keeper) *Keeper {
	if !aiRequestSubspace.HasKeyTable() {
		// register parameters of the airequest module into the param space
		aiRequestSubspace = aiRequestSubspace.WithKeyTable(types.ParamKeyTable())
	}
	return &Keeper{
		storeKey:       key,
		cdc:            cdc,
		wasmKeeper:     wasmKeeper,
		paramSpace:     aiRequestSubspace,
		stakingKeeper:  stakingKeeper,
		providerKeeper: providerKeeper,
	}
}

func (k *Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

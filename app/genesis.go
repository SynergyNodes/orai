package app

import (
	"encoding/json"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	crisis "github.com/cosmos/cosmos-sdk/x/crisis/types"
	gov "github.com/cosmos/cosmos-sdk/x/gov/types"
	mint "github.com/cosmos/cosmos-sdk/x/mint/types"
	staking "github.com/cosmos/cosmos-sdk/x/staking/types"
	provider "github.com/oraichain/orai/x/provider/types"
)

// GenesisState default state for the application
type GenesisState map[string]json.RawMessage

// NewDefaultGenesisState generates the default state for the application.
func NewDefaultGenesisState(cdc codec.Codec) GenesisState {
	genesisState := make(GenesisState)
	// Get default genesis states of the modules we are to override.
	stakingGenesis := staking.DefaultGenesisState()
	mintGenesis := mint.DefaultGenesisState()
	govGenesis := gov.DefaultGenesisState()
	crisisGenesis := crisis.DefaultGenesisState()

	stakingGenesis.Params.BondDenom = provider.Denom
	stakingGenesis.Params.HistoricalEntries = 1000

	// TODO: testnet figures only
	stakingGenesis.Params.UnbondingTime = time.Hour * 2
	stakingGenesis.Params.MaxValidators = 100
	genesisState[staking.ModuleName] = cdc.MustMarshalJSON(stakingGenesis)

	mintGenesis.Params.BlocksPerYear = 6311200 // target 5-second block time
	mintGenesis.Params.MintDenom = provider.Denom
	genesisState[mint.ModuleName] = cdc.MustMarshalJSON(mintGenesis)

	govGenesis.DepositParams.MinDeposit = sdk.NewCoins(sdk.NewCoin(provider.Denom, sdk.TokensFromConsensusPower(10, sdk.NewInt(1000000))))
	// TODO: testing
	govGenesis.VotingParams.VotingPeriod = time.Second * 30 // test for 10 mins voting period
	genesisState[gov.ModuleName] = cdc.MustMarshalJSON(govGenesis)

	crisisGenesis.ConstantFee = sdk.NewCoin(provider.Denom, sdk.TokensFromConsensusPower(10, sdk.NewInt(1000000)))
	genesisState[crisis.ModuleName] = cdc.MustMarshalJSON(crisisGenesis)

	// slashingGenesis.Params.SignedBlocksWindow = 30000                         // approximately 1 day
	// slashingGenesis.Params.MinSignedPerWindow = sdk.NewDecWithPrec(5, 2)      // 5%
	// slashingGenesis.Params.DowntimeJailDuration = 60 * 10 * time.Second       // 10 minutes
	// slashingGenesis.Params.SlashFractionDoubleSign = sdk.NewDecWithPrec(5, 2) // 5%
	// slashingGenesis.Params.SlashFractionDowntime = sdk.NewDecWithPrec(1, 4)   // 0.01%
	// Add your modules here for the genesis states

	for _, b := range ModuleBasics {
		name := b.Name()
		if name == staking.ModuleName || name == mint.ModuleName || name == gov.ModuleName || name == crisis.ModuleName {
			continue
		}
		genesisState[b.Name()] = b.DefaultGenesis(cdc)
	}

	return genesisState
}

package adblock_test

import (
	"testing"

	keepertest "adblock/testutil/keeper"
	"adblock/testutil/nullify"
	"adblock/x/adblock"
	"adblock/x/adblock/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.AdblockKeeper(t)
	adblock.InitGenesis(ctx, *k, genesisState)
	got := adblock.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}

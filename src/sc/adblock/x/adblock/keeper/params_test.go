package keeper_test

import (
	"testing"

	testkeeper "adblock/testutil/keeper"
	"adblock/x/adblock/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.AdblockKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}

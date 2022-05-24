package keeper_test

import (
	"testing"

	testkeeper "github.com/chris13524/cosmost-test/testutil/keeper"
	"github.com/chris13524/cosmost-test/x/cosmosttest/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.CosmosttestKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}

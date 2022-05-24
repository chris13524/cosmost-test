package cosmosttest_test

import (
	"testing"

	keepertest "github.com/chris13524/cosmost-test/testutil/keeper"
	"github.com/chris13524/cosmost-test/testutil/nullify"
	"github.com/chris13524/cosmost-test/x/cosmosttest"
	"github.com/chris13524/cosmost-test/x/cosmosttest/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CosmosttestKeeper(t)
	cosmosttest.InitGenesis(ctx, *k, genesisState)
	got := cosmosttest.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}

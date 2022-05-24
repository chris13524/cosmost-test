package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/chris13524/cosmost-test/testutil/keeper"
	"github.com/chris13524/cosmost-test/x/cosmosttest/keeper"
	"github.com/chris13524/cosmost-test/x/cosmosttest/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.CosmosttestKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}

package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/masterfung/checkersCosmos/testutil/keeper"
	"github.com/masterfung/checkersCosmos/x/checkerscosmos/keeper"
	"github.com/masterfung/checkersCosmos/x/checkerscosmos/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.CheckerscosmosKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}

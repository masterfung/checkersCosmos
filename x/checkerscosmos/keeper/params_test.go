package keeper_test

import (
	"testing"

	testkeeper "github.com/masterfung/checkersCosmos/testutil/keeper"
	"github.com/masterfung/checkersCosmos/x/checkerscosmos/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.CheckerscosmosKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}

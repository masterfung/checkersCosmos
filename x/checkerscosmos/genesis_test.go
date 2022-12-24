package checkerscosmos_test

import (
	"testing"

	keepertest "github.com/masterfung/checkersCosmos/testutil/keeper"
	"github.com/masterfung/checkersCosmos/testutil/nullify"
	"github.com/masterfung/checkersCosmos/x/checkerscosmos"
	"github.com/masterfung/checkersCosmos/x/checkerscosmos/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		SystemInfo: &types.SystemInfo{
			NextId: 10,
		},
		StoredGameList: []types.StoredGame{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CheckerscosmosKeeper(t)
	checkerscosmos.InitGenesis(ctx, *k, genesisState)
	got := checkerscosmos.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.SystemInfo, got.SystemInfo)
	require.ElementsMatch(t, genesisState.StoredGameList, got.StoredGameList)
	// this line is used by starport scaffolding # genesis/test/assert
}

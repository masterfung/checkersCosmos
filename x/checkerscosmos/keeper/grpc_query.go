package keeper

import (
	"github.com/masterfung/checkersCosmos/x/checkerscosmos/types"
)

var _ types.QueryServer = Keeper{}

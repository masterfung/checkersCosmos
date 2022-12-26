package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/checkerscosmos module sentinel errors
var (
	ErrInvalidBlack     = sdkerrors.Register(ModuleName, 1100, "Black address is invalid: %s")
	ErrInvalidRed       = sdkerrors.Register(ModuleName, 1101, "Red address is invalid: %s")
	ErrGameNotParseable = sdkerrors.Register(ModuleName, 1102, "Game cannot be parsed")
)

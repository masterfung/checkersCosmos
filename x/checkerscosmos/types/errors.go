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
	ErrGameNotFound = sdkerrors.Register(ModuleName, 1103, "Game by id not found")
	ErrCreatorNotPlayer = sdkerrors.Register(ModuleName, 1104, "Message creator is not a player")
	ErrNotPlayerTurn = sdkerrors.Register(ModuleName, 1105, "Player tried to play out of turn")
	ErrWrongMove = sdkerrors.Register(ModuleName, 1106, "Wrong move")
)

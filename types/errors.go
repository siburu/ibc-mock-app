package types

import (
	errorsmod "cosmossdk.io/errors"
)

// IBC transfer sentinel errors
var (
	ErrInvalidVersion = errorsmod.Register(ModuleName, 1, "invalid MockApp version")
)

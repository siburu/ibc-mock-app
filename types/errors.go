package types

import (
	errorsmod "cosmossdk.io/errors"
)

// IBC transfer sentinel errors
var (
	ErrInvalidVersion            = errorsmod.Register(ModuleName, 1, "invalid MockApp version")
	ErrUnexpectedAcknowledgement = errorsmod.Register(ModuleName, 2, "unexpected MockApp acknowledgement")
	ErrUnexpectedPacket          = errorsmod.Register(ModuleName, 3, "unexpected MockApp packet data")
)

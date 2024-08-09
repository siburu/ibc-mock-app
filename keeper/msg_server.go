package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/datachainlab/ibc-mock-app/types"
)

var _ types.MsgServer = Keeper{}

// SendPacket defines a rpc handler method for MsgSendPacket.
func (k Keeper) SendPacket(goCtx context.Context, msg *types.MsgSendPacket) (*types.MsgSendPacketResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	if err := k.sendPacket(ctx, msg.SourcePort, msg.SourceChannel, msg.Message, sender, msg.TimeoutHeight, msg.TimeoutTimestamp); err != nil {
		return nil, err
	}

	return &types.MsgSendPacketResponse{}, nil
}

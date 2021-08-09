package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSendClose{}

func NewMsgSendClose(
	sender, port, channelId string,
	timeoutTimestamp uint64,
	orderId string,
) *MsgSendClose {
	return &MsgSendClose{
		Sender:           sender,
		Port:             port,
		ChannelId:        channelId,
		TimeoutTimestamp: timeoutTimestamp,
		OrderId:          orderId,
	}
}

func (m *MsgSendClose) Route() string {
	return RouterKey
}

func (m *MsgSendClose) Type() string {
	return "Close"
}

func (m *MsgSendClose) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(m.GetSender())
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (m *MsgSendClose) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(m)
	return sdk.MustSortJSON(bz)
}

func (m *MsgSendClose) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.GetSender()); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

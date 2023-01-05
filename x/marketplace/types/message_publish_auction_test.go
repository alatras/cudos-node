package types

import (
	"testing"
	"time"

	"github.com/CudoVentures/cudos-node/testutil/sample"
	nfttypes "github.com/CudoVentures/cudos-node/x/nft/types"
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgPublishAuction_ValidateBasic(t *testing.T) {
	for _, tc := range []struct {
		desc    string
		arrange func(msg *MsgPublishAuction)
		wantErr error
	}{
		{
			desc:    "valid english auction",
			arrange: func(msg *MsgPublishAuction) {},
		},
		{
			desc: "english auction zero amount",
			arrange: func(msg *MsgPublishAuction) {
				err := msg.SetAuctionType(&EnglishAuction{MinPrice: sdk.NewCoin("acudos", sdk.ZeroInt())})
				require.NoError(t, err)
			},
			wantErr: ErrInvalidPrice,
		},
		{
			desc: "english auction invalid amount denom",
			arrange: func(msg *MsgPublishAuction) {
				err := msg.SetAuctionType(&EnglishAuction{MinPrice: sdk.Coin{"", sdk.OneInt()}})
				require.NoError(t, err)
				sdk.ZeroInt().Sub(sdk.OneInt())
			},
			wantErr: ErrInvalidPrice,
		},
		// todo validate DutchAuction
		{
			desc: "invalid auction type",
			arrange: func(msg *MsgPublishAuction) {
				msg.AuctionType = &types.Any{}
			},
			wantErr: sdkerrors.ErrInvalidType,
		},
		{
			desc: "duration less than 24 hours",
			arrange: func(msg *MsgPublishAuction) {
				msg.Duration = time.Hour * 23
			},
			wantErr: ErrInvalidAuctionDuration,
		},
		{
			desc: "invalid denom id",
			arrange: func(msg *MsgPublishAuction) {
				msg.DenomId = "123"
			},
			wantErr: nfttypes.ErrInvalidNFT,
		},
		{
			desc: "invalid token id",
			arrange: func(msg *MsgPublishAuction) {
				msg.TokenId = "invalid"
			},
			wantErr: nfttypes.ErrInvalidNFT,
		},
		{
			desc: "invalid address",
			arrange: func(msg *MsgPublishAuction) {
				msg.Creator = "invalid"
			},
			wantErr: sdkerrors.ErrInvalidAddress,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			msg, err := NewMsgPublishAuction(
				sample.AccAddress(),
				"test",
				"123",
				time.Hour*24,
				&EnglishAuction{MinPrice: sdk.NewCoin("acudos", sdk.OneInt())},
			)
			require.NoError(t, err)

			tc.arrange(msg)

			err = msg.ValidateBasic()
			require.ErrorIs(t, err, tc.wantErr)
		})
	}
}

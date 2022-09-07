package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgPublishCollection{}, "marketplace/PublishCollection", nil)
	cdc.RegisterConcrete(&MsgPublishNft{}, "marketplace/PublishNft", nil)
	cdc.RegisterConcrete(&MsgBuyNft{}, "marketplace/BuyNft", nil)
	cdc.RegisterConcrete(&MsgMintNft{}, "marketplace/MintNft", nil)
	cdc.RegisterConcrete(&MsgRemoveNft{}, "marketplace/RemoveNft", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgPublishCollection{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgPublishNft{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgBuyNft{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgMintNft{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRemoveNft{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)

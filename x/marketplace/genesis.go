package marketplace

import (
	"github.com/CudoVentures/cudos-node/x/marketplace/keeper"
	"github.com/CudoVentures/cudos-node/x/marketplace/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the collection
	for _, elem := range genState.CollectionList {
		k.SetCollection(ctx, elem)
	}

	// Set collection count
	k.SetCollectionCount(ctx, genState.CollectionCount)
	// Set all the nft
	for _, elem := range genState.NftList {
		k.SetNft(ctx, elem)
	}

	// Set nft count
	k.SetNftCount(ctx, genState.NftCount)

	// Set all the auctions
	auctions, err := types.UnpackAuctions(genState.AuctionList)
	if err != nil {
		panic(err)
	}
	for _, a := range auctions {
		if err := k.SetAuction(ctx, a); err != nil {
			panic(err)
		}
	}

	// Set auction count
	k.SetAuctionCount(ctx, genState.AuctionCount)

	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.CollectionList = k.GetAllCollection(ctx)
	genesis.CollectionCount = k.GetCollectionCount(ctx)
	genesis.NftList = k.GetAllNft(ctx)
	genesis.NftCount = k.GetNftCount(ctx)

	auctions, err := k.GetAllAuction(ctx)
	if err != nil {
		panic(err)
	}

	auctionsAny, err := types.PackAuctions(auctions)
	if err != nil {
		panic(err)
	}

	genesis.AuctionList = auctionsAny
	genesis.AuctionCount = k.GetAuctionCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}

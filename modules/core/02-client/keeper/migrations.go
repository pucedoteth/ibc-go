package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	clientexported "github.com/cosmos/ibc-go/v5/modules/core/02-client/exported"
	v100 "github.com/cosmos/ibc-go/v5/modules/core/02-client/legacy/v100"
)

// Migrator is a struct for handling in-place store migrations.
type Migrator struct {
	keeper clientexported.ClientKeeper
}

// NewMigrator returns a new Migrator.
func NewMigrator(keeper clientexported.ClientKeeper) Migrator {
	return Migrator{keeper: keeper}
}

// Migrate1to2 migrates from version 1 to 2.
// This migration
// - migrates solo machine client states from v1 to v2 protobuf definition
// - prunes solo machine consensus states
// - prunes expired tendermint consensus states
// - adds iteration and processed height keys for unexpired tendermint consensus states
func (m Migrator) Migrate1to2(ctx sdk.Context) error {
	return v100.MigrateStore(ctx, m.keeper.GetStoreKey(), m.keeper.GetCdc())
}

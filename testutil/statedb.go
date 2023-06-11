// Copyright Tharsis Labs Ltd.(Black)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/black/black/blob/main/LICENSE)

package testutil

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/black/black/v13/app/ante/evm"
	"github.com/black/black/v13/x/evm/statedb"
)

// NewStateDB returns a new StateDB for testing purposes.
func NewStateDB(ctx sdk.Context, evmKeeper evm.EVMKeeper) *statedb.StateDB {
	return statedb.New(ctx, evmKeeper, statedb.NewEmptyTxConfig(common.BytesToHash(ctx.HeaderHash().Bytes())))
}

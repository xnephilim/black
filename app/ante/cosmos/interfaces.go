// Copyright Tharsis Labs Ltd.(Black)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/xnephilim/black/blob/main/LICENSE)

package cosmos

import sdk "github.com/cosmos/cosmos-sdk/types"

// BankKeeper defines the exposed interface for using functionality of the bank keeper
// in the context of the cosmos AnteHandler package.
type BankKeeper interface {
	GetBalance(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin
	SendCoins(ctx sdk.Context, from, to sdk.AccAddress, amt sdk.Coins) error
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
}

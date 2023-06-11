// Copyright Tharsis Labs Ltd.(Black)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/xnephilim/black/blob/main/LICENSE)

package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
)

var ModuleCdc = codec.NewProtoCodec(codectypes.NewInterfaceRegistry())

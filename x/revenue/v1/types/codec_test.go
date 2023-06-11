package types

import (
	"testing"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/stretchr/testify/suite"
)

type CodecTestSuite struct {
	suite.Suite
}

func TestCodecSuite(t *testing.T) {
	suite.Run(t, new(CodecTestSuite))
}

func (suite *CodecTestSuite) TestRegisterInterfaces() {
	registry := codectypes.NewInterfaceRegistry()
	registry.RegisterInterface(sdk.MsgInterfaceProtoName, (*sdk.Msg)(nil))
	RegisterInterfaces(registry)

	impls := registry.ListImplementations(sdk.MsgInterfaceProtoName)
	suite.Require().Equal(4, len(impls))
	suite.Require().ElementsMatch([]string{
		"/black.revenue.v1.MsgRegisterRevenue",
		"/black.revenue.v1.MsgCancelRevenue",
		"/black.revenue.v1.MsgUpdateRevenue",
		"/black.revenue.v1.MsgUpdateParams",
	}, impls)
}

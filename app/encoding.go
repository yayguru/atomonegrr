package atomone

import (
	"github.com/cosmos/cosmos-sdk/std"

	"github.com/atomone-hub/atomone/app/params"
)

func RegisterEncodingConfig() params.EncodingConfig {
	encConfig := params.MakeEncodingConfig()

	std.RegisterLegacyAminoCodec(encConfig.Amino)
	std.RegisterInterfaces(encConfig.InterfaceRegistry)
	ModuleBasics.RegisterLegacyAminoCodec(encConfig.Amino)
	ModuleBasics.RegisterInterfaces(encConfig.InterfaceRegistry)

	return encConfig
}

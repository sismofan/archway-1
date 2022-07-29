package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkErrors "github.com/cosmos/cosmos-sdk/types/errors"
	"sigs.k8s.io/yaml"
)

// HasOwnerAddress returns true if the rewards address is set.
func (m ContractMetadata) HasOwnerAddress() bool {
	return m.OwnerAddress != ""
}

// HasRewardsAddress returns true if the rewards address is set.
func (m ContractMetadata) HasRewardsAddress() bool {
	return m.RewardsAddress != ""
}

// MustGetRewardsAddress returns the rewards address.
// CONTRACT: panics in case of an error.
func (m ContractMetadata) MustGetRewardsAddress() sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.RewardsAddress)
	if err != nil {
		panic(fmt.Errorf("parsing rewards address (%s): %s", m.RewardsAddress, err))
	}

	return addr
}

// Validate performs object fields validation.
// genesisValidation flag perform strict validation of the genesis state (some field must be set).
func (m ContractMetadata) Validate(genesisValidation bool) error {
	if genesisValidation || m.OwnerAddress != "" {
		if _, err := sdk.AccAddressFromBech32(m.OwnerAddress); err != nil {
			return sdkErrors.Wrap(sdkErrors.ErrInvalidAddress, "invalid owner address")
		}
	}

	if m.RewardsAddress != "" {
		if _, err := sdk.AccAddressFromBech32(m.RewardsAddress); err != nil {
			return sdkErrors.Wrap(sdkErrors.ErrInvalidAddress, "invalid rewards address")
		}
	}

	return nil
}

// String implements the fmt.Stringer interface.
func (m ContractMetadata) String() string {
	bz, _ := yaml.Marshal(m)
	return string(bz)
}

package types

import (
	errors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/ibc-go/v8/modules/core/02-client/types"
)

var (
	_ sdk.Msg = &MsgClientUpdate{}
)

// Route implements the LegacyMsg interface.
func (m *MsgClientUpdate) Route() string { return sdk.MsgTypeURL(m) }

// Type implements the LegacyMsg interface.
func (m *MsgClientUpdate) Type() string { return sdk.MsgTypeURL(m) }

// GetSignBytes implements the LegacyMsg interface.
func (m *MsgClientUpdate) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(m))
}

// ValidateBasic runs basic stateless validity checks
func (m *MsgClientUpdate) ValidateBasic() error {
	if m.SubjectClientId == m.SubstituteClientId {
		return errors.Wrap(clienttypes.ErrInvalidSubstitute, "subject and substitute client identifiers are equal")
	}
	if _, _, err := clienttypes.ParseClientIdentifier(m.SubjectClientId); err != nil {
		return err
	}
	if _, _, err := clienttypes.ParseClientIdentifier(m.SubstituteClientId); err != nil {
		return err
	}

	return nil
}

// GetSigners returns the expected signers for MsgClientUpdate.
func (m *MsgClientUpdate) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Authority)
	return []sdk.AccAddress{addr}
}

// ValidateBasic runs basic stateless validity checks
func (up *MsgUpgrade) ValidateBasic() error {
	if err := up.Plan.ValidateBasic(); err != nil {
		return err
	}

	if up.UpgradedClientState == nil {
		return errors.Wrap(clienttypes.ErrInvalidUpgradeProposal, "upgraded client state cannot be nil")
	}

	_, err := clienttypes.UnpackClientState(up.UpgradedClientState)
	if err != nil {
		return errors.Wrap(err, "failed to unpack upgraded client state")
	}

	return nil
}

// Route implements the LegacyMsg interface.
func (m *MsgUpgrade) Route() string { return sdk.MsgTypeURL(m) }

// Type implements the LegacyMsg interface.
func (m *MsgUpgrade) Type() string { return sdk.MsgTypeURL(m) }

// GetSignBytes implements the LegacyMsg interface.
func (m *MsgUpgrade) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(m))
}

// GetSigners returns the expected signers for MsgClientUpdate.
func (m *MsgUpgrade) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Authority)
	return []sdk.AccAddress{addr}
}

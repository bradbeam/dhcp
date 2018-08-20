package dhcpv6

import (
	"net"
	"testing"

	"github.com/insomniacslk/dhcp/iana"
	"github.com/stretchr/testify/require"
)

func TestWithClientID(t *testing.T) {
	duid := Duid{
		Type:          DUID_LL,
		HwType:        iana.HwTypeEthernet,
		LinkLayerAddr: net.HardwareAddr([]byte{0xfa, 0xce, 0xb0, 0x00, 0x00, 0x0c}),
	}
	m, err := NewMessage(WithClientID(duid))
	require.NoError(t, err)
	opt := m.GetOneOption(OptionClientID)
	require.NotNil(t, opt)
	cid := opt.(*OptClientId)
	require.Equal(t, cid.Cid, duid)
}

func TestWithServerID(t *testing.T) {
	duid := Duid{
		Type:          DUID_LL,
		HwType:        iana.HwTypeEthernet,
		LinkLayerAddr: net.HardwareAddr([]byte{0xfa, 0xce, 0xb0, 0x00, 0x00, 0x0c}),
	}
	m, err := NewMessage(WithServerID(duid))
	require.NoError(t, err)
	opt := m.GetOneOption(OptionServerID)
	require.NotNil(t, opt)
	sid := opt.(*OptServerId)
	require.Equal(t, sid.Sid, duid)
}

func TestWithRequestedOptions(t *testing.T) {
	// Check if ORO is created when no ORO present
	m, err := NewMessage(WithRequestedOptions(OptionClientID))
	require.NoError(t, err)
	opt := m.GetOneOption(OptionORO)
	require.NotNil(t, opt)
	oro := opt.(*OptRequestedOption)
	require.ElementsMatch(t, oro.RequestedOptions(), []OptionCode{OptionClientID})
	// Check if already set options are preserved
	m = WithRequestedOptions(OptionServerID)(m)
	opt = m.GetOneOption(OptionORO)
	require.NotNil(t, opt)
	oro = opt.(*OptRequestedOption)
	require.ElementsMatch(t, oro.RequestedOptions(), []OptionCode{OptionClientID, OptionServerID})
}

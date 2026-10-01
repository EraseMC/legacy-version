package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

// ClientBoundDataDrivenUIShowScreen is sent by the server to show a data-driven UI screen on the client.
type ClientBoundDataDrivenUIShowScreen struct {
	// ScreenID is the identifier of the screen to show.
	ScreenID string
	// FormID is a unique instance ID for the form, used for scripting to identify specific screen instances.
	FormID uint32
	// DataInstanceID is an optional data ID associated with the screen.
	DataInstanceID protocol.Optional[uint32]
}

// ID ...
func (*ClientBoundDataDrivenUIShowScreen) ID() uint32 {
	return packet.IDClientBoundDataDrivenUIShowScreen
}

func (pk *ClientBoundDataDrivenUIShowScreen) Marshal(io protocol.IO) {
	io.String(&pk.ScreenID)
	if proto.IsProtoGTE(io, proto.ID944) {
		io.Uint32(&pk.FormID)
		protocol.OptionalFunc(io, &pk.DataInstanceID, io.Uint32)
	}
}

// ToLatest ...
func (pk *ClientBoundDataDrivenUIShowScreen) ToLatest() *packet.ClientBoundDataDrivenUIShowScreen {
	return &packet.ClientBoundDataDrivenUIShowScreen{
		ScreenID:       pk.ScreenID,
		FormID:         pk.FormID,
		DataInstanceID: pk.DataInstanceID,
	}
}

// FromLatest ...
func (pk *ClientBoundDataDrivenUIShowScreen) FromLatest(latest *packet.ClientBoundDataDrivenUIShowScreen) *ClientBoundDataDrivenUIShowScreen {
	pk.ScreenID = latest.ScreenID
	pk.FormID = latest.FormID
	pk.DataInstanceID = latest.DataInstanceID
	return pk
}

package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

// PartyChanged is sent by the client to the server to indicate that the player's party ID has changed.
type PartyChanged struct {
	PartyInfo protocol.Optional[packet.PartyInfo]
}

// ID ...
func (*PartyChanged) ID() uint32 {
	return packet.IDPartyChanged
}

func (pk *PartyChanged) Marshal(io protocol.IO) {
	if proto.IsProtoGTE(io, proto.ID975) {
		protocol.OptionalMarshaler(io, &pk.PartyInfo)
	} else {
		var partyID string
		if v, ok := pk.PartyInfo.Value(); ok {
			partyID = v.PartyID
		}
		io.String(&partyID)
		if partyID != "" {
			pk.PartyInfo = protocol.Option(packet.PartyInfo{PartyID: partyID})
		}
	}
}

// ToLatest ...
func (pk *PartyChanged) ToLatest() *packet.PartyChanged {
	return &packet.PartyChanged{
		PartyInfo: pk.PartyInfo,
	}
}

// FromLatest ...
func (pk *PartyChanged) FromLatest(latest *packet.PartyChanged) *PartyChanged {
	pk.PartyInfo = latest.PartyInfo
	return pk
}

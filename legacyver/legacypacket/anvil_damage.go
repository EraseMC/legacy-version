package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

// AnvilDamage is sent by the client to request the dealing damage to an anvil. This packet is completely
// pointless and the server should never listen to it.
type AnvilDamage struct {
	// Damage is the damage that the client requests to be dealt to the anvil.
	Damage uint8
	// AnvilPosition is the position in the world that the anvil can be found at.
	AnvilPosition protocol.BlockPos
}

// ID ...
func (*AnvilDamage) ID() uint32 {
	return packet.IDAnvilDamage
}

func (pk *AnvilDamage) Marshal(io protocol.IO) {
	io.Uint8(&pk.Damage)
	if proto.IsProtoGTE(io, proto.ID944) {
		io.BlockPos(&pk.AnvilPosition)
	} else {
		proto.UBlockPos(io, &pk.AnvilPosition)
	}
}

// ToLatest ...
func (pk *AnvilDamage) ToLatest() *packet.AnvilDamage {
	return &packet.AnvilDamage{
		Damage:        pk.Damage,
		AnvilPosition: pk.AnvilPosition,
	}
}

// FromLatest ...
func (pk *AnvilDamage) FromLatest(latest *packet.AnvilDamage) *AnvilDamage {
	pk.Damage = latest.Damage
	pk.AnvilPosition = latest.AnvilPosition
	return pk
}

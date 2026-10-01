package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

const (
	SpawnTypePlayer = iota
	SpawnTypeWorld
)

// SetSpawnPosition is sent by the server to update the spawn position of a player, for example when sleeping
// in a bed.
type SetSpawnPosition struct {
	// SpawnType is the type of spawn to set. It is either SpawnTypePlayer or SpawnTypeWorld, and specifies
	// the behaviour of the spawn set. If SpawnTypeWorld is set, the position to which compasses will point is
	// also changed.
	SpawnType int32
	// Position is the new position of the spawn that was set. If SpawnType is SpawnTypeWorld, compasses will
	// point to this position. As of 1.16, Position is always the position of the player.
	Position protocol.BlockPos
	// Dimension is the ID of the dimension that had its spawn updated. This is specifically relevant for
	// behaviour added in 1.16 such as the respawn anchor, which allows setting the spawn in a specific
	// dimension.
	Dimension int32
	// SpawnPosition is a new field added in 1.16. It holds the spawn position of the world. This spawn
	// position is {-2147483648, -2147483648, -2147483648} for a default spawn position.
	SpawnPosition protocol.BlockPos
}

// ID ...
func (*SetSpawnPosition) ID() uint32 {
	return packet.IDSetSpawnPosition
}

func (pk *SetSpawnPosition) Marshal(io protocol.IO) {
	io.Varint32(&pk.SpawnType)
	if proto.IsProtoGTE(io, proto.ID944) {
		io.BlockPos(&pk.Position)
	} else {
		proto.UBlockPos(io, &pk.Position)
	}
	io.Varint32(&pk.Dimension)
	if proto.IsProtoGTE(io, proto.ID944) {
		io.BlockPos(&pk.SpawnPosition)
	} else {
		proto.UBlockPos(io, &pk.SpawnPosition)
	}
}

// ToLatest ...
func (pk *SetSpawnPosition) ToLatest() *packet.SetSpawnPosition {
	return &packet.SetSpawnPosition{
		SpawnType:     pk.SpawnType,
		Position:      pk.Position,
		Dimension:     pk.Dimension,
		SpawnPosition: pk.SpawnPosition,
	}
}

// FromLatest ...
func (pk *SetSpawnPosition) FromLatest(latest *packet.SetSpawnPosition) *SetSpawnPosition {
	pk.SpawnType = latest.SpawnType
	pk.Position = latest.Position
	pk.Dimension = latest.Dimension
	pk.SpawnPosition = latest.SpawnPosition
	return pk
}

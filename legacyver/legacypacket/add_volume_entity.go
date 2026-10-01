package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/nbt"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

// AddVolumeEntity sends a volume entity's definition and metadata from server to client.
type AddVolumeEntity struct {
	// EntityRuntimeID is the runtime ID of the volume. The runtime ID is unique for each world session, and
	// entities are generally identified in packets using this runtime ID.
	EntityRuntimeID uint32
	// EntityMetadata is a map of entity metadata, which includes flags and data properties that alter in
	// particular the way the volume functions or looks.
	EntityMetadata map[string]any
	// EncodingIdentifier is the unique identifier for the volume. It must be of the form 'namespace:name', where
	// namespace cannot be 'minecraft'.
	EncodingIdentifier string
	// InstanceIdentifier is the identifier of a fog definition.
	InstanceIdentifier string
	// Bounds represent the volume's bounds. The first value is the minimum bounds, and the second value is the
	// maximum bounds.
	Bounds [2]protocol.BlockPos
	// Dimension is the dimension in which the volume exists.
	Dimension int32
	// EngineVersion is the engine version the entity is using, for example, '1.17.0'.
	EngineVersion string
}

// ID ...
func (*AddVolumeEntity) ID() uint32 {
	return packet.IDAddVolumeEntity
}

func (pk *AddVolumeEntity) Marshal(io protocol.IO) {
	io.Varuint32(&pk.EntityRuntimeID)
	io.NBT(&pk.EntityMetadata, nbt.NetworkLittleEndian)
	io.String(&pk.EncodingIdentifier)
	io.String(&pk.InstanceIdentifier)
	if proto.IsProtoGTE(io, proto.ID944) {
		io.BlockPos(&pk.Bounds[0])
		io.BlockPos(&pk.Bounds[1])
	} else {
		proto.UBlockPos(io, &pk.Bounds[0])
		proto.UBlockPos(io, &pk.Bounds[1])
	}
	io.Varint32(&pk.Dimension)
	io.String(&pk.EngineVersion)
}

// ToLatest ...
func (pk *AddVolumeEntity) ToLatest() *packet.AddVolumeEntity {
	return &packet.AddVolumeEntity{
		EntityRuntimeID:    pk.EntityRuntimeID,
		EntityMetadata:     pk.EntityMetadata,
		EncodingIdentifier: pk.EncodingIdentifier,
		InstanceIdentifier: pk.InstanceIdentifier,
		Bounds:             pk.Bounds,
		Dimension:          pk.Dimension,
		EngineVersion:      pk.EngineVersion,
	}
}

// FromLatest ...
func (pk *AddVolumeEntity) FromLatest(latest *packet.AddVolumeEntity) *AddVolumeEntity {
	pk.EntityRuntimeID = latest.EntityRuntimeID
	pk.EntityMetadata = latest.EntityMetadata
	pk.EncodingIdentifier = latest.EncodingIdentifier
	pk.InstanceIdentifier = latest.InstanceIdentifier
	pk.Bounds = latest.Bounds
	pk.Dimension = latest.Dimension
	pk.EngineVersion = latest.EngineVersion
	return pk
}

package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

// VoxelShapes is sent by the server to send voxel shape data to the client.
type VoxelShapes struct {
	// Shapes is a list of voxel shapes.
	Shapes []protocol.VoxelShape
	// NameMap is a map of shape names to IDs.
	NameMap []protocol.VoxelShapeNameEntry
	// CustomShapeCount is the number of custom shapes.
	CustomShapeCount uint16
}

// ID ...
func (*VoxelShapes) ID() uint32 {
	return packet.IDVoxelShapes
}

// Marshal ...
func (pk *VoxelShapes) Marshal(io protocol.IO) {
	protocol.Slice(io, &pk.Shapes)
	protocol.Slice(io, &pk.NameMap)
	if proto.IsProtoGTE(io, proto.ID944) {
		io.Uint16(&pk.CustomShapeCount)
	}
}

// ToLatest ...
func (pk *VoxelShapes) ToLatest() *packet.VoxelShapes {
	shapes := make([]protocol.VoxelShape, len(pk.Shapes))
	for i, shape := range pk.Shapes {
		shapes[i] = shape
	}
	nameMap := make([]protocol.VoxelShapeNameEntry, len(pk.NameMap))
	for i, entry := range pk.NameMap {
		nameMap[i] = entry
	}
	return &packet.VoxelShapes{
		Shapes:           shapes,
		NameMap:          nameMap,
		CustomShapeCount: pk.CustomShapeCount,
	}
}

// FromLatest ...
func (pk *VoxelShapes) FromLatest(latest *packet.VoxelShapes) *VoxelShapes {
	pk.Shapes = make([]protocol.VoxelShape, len(latest.Shapes))
	for i, shape := range latest.Shapes {
		pk.Shapes[i] = shape
	}
	pk.NameMap = make([]protocol.VoxelShapeNameEntry, len(latest.NameMap))
	for i, entry := range latest.NameMap {
		pk.NameMap[i] = entry
	}
	pk.CustomShapeCount = latest.CustomShapeCount
	return pk
}

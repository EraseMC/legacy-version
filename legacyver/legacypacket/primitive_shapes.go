package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

// PrimitiveShapes ...
type PrimitiveShapes struct {
	// Shapes is a list of shapes to draw on the client-side.
	Shapes []proto.PrimitiveShape
}

// ID ...
func (pk *PrimitiveShapes) ID() uint32 {
	return packet.IDPrimitiveShapes
}

func (pk *PrimitiveShapes) Marshal(io protocol.IO) {
	protocol.Slice(io, &pk.Shapes)
}

// ToLatest ...
func (pk *PrimitiveShapes) ToLatest() *packet.PrimitiveShapes {
	shapes := make([]protocol.PrimitiveShape, len(pk.Shapes))
	for i, shape := range pk.Shapes {
		shapes[i] = shape.ToLatest()
	}
	return &packet.PrimitiveShapes{
		Shapes: shapes,
	}
}

// FromLatest ...
func (pk *PrimitiveShapes) FromLatest(latest *packet.PrimitiveShapes) *PrimitiveShapes {
	pk.Shapes = make([]proto.PrimitiveShape, len(latest.Shapes))
	for i, shape := range latest.Shapes {
		pk.Shapes[i] = (&proto.PrimitiveShape{}).FromLatest(shape)
	}
	return pk
}

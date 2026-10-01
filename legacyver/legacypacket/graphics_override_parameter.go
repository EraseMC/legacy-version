package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

// GraphicsOverrideParameter is sent by the server to override graphics parameters.
type GraphicsOverrideParameter struct {
	// Values is a list of parameter keyframe values.
	Values []protocol.ParameterKeyframeValue
	// FloatValue ...
	FloatValue protocol.Optional[float32]
	// Vec3Value ...
	Vec3Value protocol.Optional[mgl32.Vec3]
	// BiomeIdentifier is the identifier of the biome for which the parameters apply.
	BiomeIdentifier string
	// ParameterType is the type of parameter being overridden.
	ParameterType uint8
	// Reset indicates whether to reset the parameters.
	Reset bool
}

// ID ...
func (*GraphicsOverrideParameter) ID() uint32 {
	return packet.IDGraphicsOverrideParameter
}

func (pk *GraphicsOverrideParameter) Marshal(io protocol.IO) {
	protocol.Slice(io, &pk.Values)
	if proto.IsProtoGTE(io, proto.ID924) {
		protocol.OptionalFunc(io, &pk.FloatValue, io.Float32)
		protocol.OptionalFunc(io, &pk.Vec3Value, io.Vec3)
	}
	io.String(&pk.BiomeIdentifier)
	io.Uint8(&pk.ParameterType)
	io.Bool(&pk.Reset)
}

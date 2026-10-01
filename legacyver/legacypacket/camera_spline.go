package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

// CameraSpline is sent by the server to define camera spline paths.
type CameraSpline struct {
	// Splines is a list of camera spline definitions.
	Splines []proto.CameraSplineDefinition
}

// ID ...
func (*CameraSpline) ID() uint32 {
	return packet.IDCameraSpline
}

func (pk *CameraSpline) Marshal(io protocol.IO) {
	protocol.Slice(io, &pk.Splines)
}

// ToLatest ...
func (pk *CameraSpline) ToLatest() *packet.CameraSpline {
	splines := make([]protocol.CameraSplineDefinition, len(pk.Splines))
	for i, spline := range pk.Splines {
		splines[i] = spline.ToLatest()
	}
	return &packet.CameraSpline{
		Splines: splines,
	}
}

// FromLatest ...
func (pk *CameraSpline) FromLatest(latest *packet.CameraSpline) *CameraSpline {
	pk.Splines = make([]proto.CameraSplineDefinition, len(latest.Splines))
	for i, spline := range latest.Splines {
		pk.Splines[i] = pk.Splines[i].FromLatest(spline)
	}
	return pk
}

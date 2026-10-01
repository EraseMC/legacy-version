package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

const (
	CameraAunAssistPresetOperationSet = iota
	CameraAunAssistPresetOperationAddToExisting
)

// CameraAimAssistPresets is sent by the server to the client to provide a list of categories and presets
// that can be used when sending a CameraAimAssist packet or a CameraInstruction including aim assist.
type CameraAimAssistPresets struct {
	// Categories is a list of groups of categories which can be referenced by one of the Presets.
	Categories []proto.CameraAimAssistCategory
	// Presets is a list of presets which define a base for how aim assist should behave
	Presets []proto.CameraAimAssistPreset
	// Operation is the operation to perform with the presets. It is one of the constants above.
	Operation byte
}

// ID ...
func (*CameraAimAssistPresets) ID() uint32 {
	return packet.IDCameraAimAssistPresets
}

func (pk *CameraAimAssistPresets) Marshal(io protocol.IO) {
	protocol.Slice(io, &pk.Categories)
	protocol.Slice(io, &pk.Presets)
	if proto.IsProtoGTE(io, proto.ID776) {
		io.Uint8(&pk.Operation)
	}
}

// ToLatest ...
func (pk *CameraAimAssistPresets) ToLatest() *packet.CameraAimAssistPresets {
	categories := make([]protocol.CameraAimAssistCategory, len(pk.Categories))
	for i, category := range pk.Categories {
		categories[i] = category.ToLatest()
	}
	presets := make([]protocol.CameraAimAssistPreset, len(pk.Presets))
	for i, preset := range pk.Presets {
		presets[i] = preset.ToLatest()
	}
	return &packet.CameraAimAssistPresets{
		Categories: categories,
		Presets:    presets,
		Operation:  pk.Operation,
	}
}

// FromLatest ...
func (pk *CameraAimAssistPresets) FromLatest(latest *packet.CameraAimAssistPresets) *CameraAimAssistPresets {
	pk.Categories = make([]proto.CameraAimAssistCategory, len(latest.Categories))
	for i, category := range latest.Categories {
		pk.Categories[i] = pk.Categories[i].FromLatest(category)
	}
	pk.Presets = make([]proto.CameraAimAssistPreset, len(latest.Presets))
	for i, preset := range latest.Presets {
		pk.Presets[i] = pk.Presets[i].FromLatest(preset)
	}
	pk.Operation = latest.Operation
	return pk
}

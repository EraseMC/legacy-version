package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

const (
	StructureTemplateRequestExportFromSave = iota + 1
	StructureTemplateRequestExportFromLoad
	StructureTemplateRequestQuerySavedStructure
)

// StructureTemplateDataRequest is sent by the client to request data of a structure.
type StructureTemplateDataRequest struct {
	// StructureName is the name of the structure that was set in the structure block's UI. This is the name
	// used to export the structure to a file.
	StructureName string
	// Position is the position of the structure block that has its template data requested.
	Position protocol.BlockPos
	// Settings is a struct of settings that should be used for exporting the structure. These settings are
	// identical to the last sent in the StructureBlockUpdate packet by the client.
	Settings proto.StructureSettings
	// RequestType specifies the type of template data request that the player sent. It is one of the
	// constants found above.
	RequestType byte
}

// ID ...
func (pk *StructureTemplateDataRequest) ID() uint32 {
	return packet.IDStructureTemplateDataRequest
}

func (pk *StructureTemplateDataRequest) Marshal(io protocol.IO) {
	io.String(&pk.StructureName)
	if proto.IsProtoGTE(io, proto.ID944) {
		io.BlockPos(&pk.Position)
	} else {
		proto.UBlockPos(io, &pk.Position)
	}
	protocol.Single(io, &pk.Settings)
	io.Uint8(&pk.RequestType)
}

// ToLatest ...
func (pk *StructureTemplateDataRequest) ToLatest() *packet.StructureTemplateDataRequest {
	return &packet.StructureTemplateDataRequest{
		StructureName: pk.StructureName,
		Position:      pk.Position,
		Settings:      pk.Settings.ToLatest(),
		RequestType:   pk.RequestType,
	}
}

// FromLatest ...
func (pk *StructureTemplateDataRequest) FromLatest(latest *packet.StructureTemplateDataRequest) *StructureTemplateDataRequest {
	pk.StructureName = latest.StructureName
	pk.Position = latest.Position
	pk.Settings = pk.Settings.FromLatest(latest.Settings)
	pk.RequestType = latest.RequestType
	return pk
}

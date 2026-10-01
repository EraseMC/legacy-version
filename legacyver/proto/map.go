package proto

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

const (
	MapDecorationTypeMarkerWhite = iota
	MapDecorationTypeMarkerGreen
	MapDecorationTypeMarkerRed
	MapDecorationTypeMarkerBlue
	MapDecorationTypeCrossWhite
	MapDecorationTypeTriangleRed
	MapDecorationTypeSquareWhite
	MapDecorationTypeMarkerSign
	MapDecorationTypeMarkerPink
	MapDecorationTypeMarkerOrange
	MapDecorationTypeMarkerYellow
	MapDecorationTypeMarkerTeal
	MapDecorationTypeTriangleGreen
	MapDecorationTypeSmallSquareWhite
	MapDecorationTypeMansion
	MapDecorationTypeMonument
	MapDecorationTypeNoDraw
	MapDecorationTypeVillageDesert
	MapDecorationTypeVillagePlains
	MapDecorationTypeVillageSavanna
	MapDecorationTypeVillageSnowy
	MapDecorationTypeVillageTaiga
	MapDecorationTypeJungleTemple
	MapDecorationTypeWitchHut
)

const (
	MapObjectTypeEntity = iota
	MapObjectTypeBlock
)

// MapTrackedObject is an object on a map that is 'tracked' by the client, such as an entity or a block. This
// object may move, which is handled client-side.
type MapTrackedObject struct {
	// Type is the type of the tracked object. It is either MapObjectTypeEntity or MapObjectTypeBlock.
	Type int32
	// EntityUniqueID is the unique ID of the entity, if the tracked object was an entity. It needs not to be
	// filled out if Type is not MapObjectTypeEntity.
	EntityUniqueID int64
	// BlockPosition is the position of the block, if the tracked object was a block. It needs not to be
	// filled out if Type is not MapObjectTypeBlock.
	BlockPosition protocol.BlockPos
}

// Marshal encodes/decodes a MapTrackedObject.
func (x *MapTrackedObject) Marshal(r protocol.IO) {
	r.Int32(&x.Type)
	switch x.Type {
	case MapObjectTypeEntity:
		r.Varint64(&x.EntityUniqueID)
	case MapObjectTypeBlock:
		if IsProtoGTE(r, ID944) {
			r.BlockPos(&x.BlockPosition)
		} else {
			UBlockPos(r, &x.BlockPosition)
		}
	default:
		r.UnknownEnumOption(x.Type, "map tracked object type")
	}
}

// ToLatest ...
func (x *MapTrackedObject) ToLatest() protocol.MapTrackedObject {
	return protocol.MapTrackedObject{
		Type:           x.Type,
		EntityUniqueID: x.EntityUniqueID,
		BlockPosition:  x.BlockPosition,
	}
}

// FromLatest ...
func (x *MapTrackedObject) FromLatest(latest protocol.MapTrackedObject) MapTrackedObject {
	x.Type = latest.Type
	x.EntityUniqueID = latest.EntityUniqueID
	x.BlockPosition = latest.BlockPosition
	return *x
}

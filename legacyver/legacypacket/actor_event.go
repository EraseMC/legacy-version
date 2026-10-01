package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

// ActorEvent is sent by the server when a particular event happens that has to do with an entity. Some of
// these events are entity-specific, for example a wolf shaking itself dry, but others are used for each
// entity, such as dying.
type ActorEvent struct {
	// EntityRuntimeID is the runtime ID of the entity. The runtime ID is unique for each world session, and
	// entities are generally identified in packets using this runtime ID.
	EntityRuntimeID uint64
	// EventType is the ID of the event to be called. It is one of the constants that can be found above.
	EventType byte
	// EventData is optional data associated with a particular event. The data has a different function for
	// different events, however most events don't use this field at all.
	EventData int32
	// FireAtPosition is the position in the same world at which the event should fire. If this is not present,
	// the position entity will be used instead.
	FireAtPosition protocol.Optional[mgl32.Vec3]
}

// ID ...
func (*ActorEvent) ID() uint32 {
	return packet.IDActorEvent
}

func (pk *ActorEvent) Marshal(io protocol.IO) {
	io.Varuint64(&pk.EntityRuntimeID)
	io.Uint8(&pk.EventType)
	io.Varint32(&pk.EventData)
	if proto.IsProtoGTE(io, proto.ID975) {
		protocol.OptionalFunc(io, &pk.FireAtPosition, io.Vec3)
	}
}

// ToLatest ...
func (pk *ActorEvent) ToLatest() *packet.ActorEvent {
	return &packet.ActorEvent{
		EntityRuntimeID: pk.EntityRuntimeID,
		EventType:       pk.EventType,
		EventData:       pk.EventData,
		FireAtPosition:  pk.FireAtPosition,
	}
}

// FromLatest ...
func (pk *ActorEvent) FromLatest(latest *packet.ActorEvent) *ActorEvent {
	pk.EntityRuntimeID = latest.EntityRuntimeID
	pk.EventType = latest.EventType
	pk.EventData = latest.EventData
	pk.FireAtPosition = latest.FireAtPosition
	return pk
}

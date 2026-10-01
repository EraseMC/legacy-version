package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

// PlayerAction is sent by the client when it executes any action, for example starting to sprint, swim,
// starting the breaking of a block, dropping an item, etc.
type PlayerAction struct {
	// EntityRuntimeID is the runtime ID of the player. The runtime ID is unique for each world session, and
	// entities are generally identified in packets using this runtime ID.
	EntityRuntimeID uint64
	// ActionType is the ID of the action that was executed by the player. It is one of the constants that may
	// be found in protocol/player.go.
	ActionType int32
	// BlockPosition is the position of the target block, if the action with the ActionType set concerned a
	// block. If that is not the case, the block position will be zero.
	BlockPosition protocol.BlockPos
	// ResultPosition is the position of the action's result. When a UseItemOn action is sent, this is the position of
	// the block clicked, but when a block is placed, this is the position at which the block will be placed.
	ResultPosition protocol.BlockPos
	// BlockFace is the face of the target block that was touched. If the action with the ActionType set
	// concerned a block. If not, the face is always 0.
	BlockFace int32
}

// ID ...
func (*PlayerAction) ID() uint32 {
	return packet.IDPlayerAction
}

func (pk *PlayerAction) Marshal(io protocol.IO) {
	io.Varuint64(&pk.EntityRuntimeID)
	io.Varint32(&pk.ActionType)
	if proto.IsProtoGTE(io, proto.ID944) {
		io.BlockPos(&pk.BlockPosition)
		io.BlockPos(&pk.ResultPosition)
	} else {
		proto.UBlockPos(io, &pk.BlockPosition)
		proto.UBlockPos(io, &pk.ResultPosition)
	}
	io.Varint32(&pk.BlockFace)
}

// ToLatest ...
func (pk *PlayerAction) ToLatest() *packet.PlayerAction {
	return &packet.PlayerAction{
		EntityRuntimeID: pk.EntityRuntimeID,
		ActionType:      pk.ActionType,
		BlockPosition:   pk.BlockPosition,
		ResultPosition:  pk.ResultPosition,
		BlockFace:       pk.BlockFace,
	}
}

// FromLatest ...
func (pk *PlayerAction) FromLatest(latest *packet.PlayerAction) *PlayerAction {
	pk.EntityRuntimeID = latest.EntityRuntimeID
	pk.ActionType = latest.ActionType
	pk.BlockPosition = latest.BlockPosition
	pk.ResultPosition = latest.ResultPosition
	pk.BlockFace = latest.BlockFace
	return pk
}

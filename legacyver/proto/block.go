package proto

import "github.com/sandertv/gophertunnel/minecraft/protocol"

// BlockChangeEntry is used by the UpdateSubChunkBlocks packet.
type BlockChangeEntry struct {
	protocol.BlockPos
	// BlockRuntimeID is the runtime ID of the block.
	BlockRuntimeID uint32
	// Flags is a combination of flags that specify the way the block is updated client-side.
	Flags uint32
	// SyncedUpdateEntityUniqueID  is the unique ID of the falling block entity that the block transitions to or that the entity transitions from if the block change entry is synced.
	SyncedUpdateEntityUniqueID uint64
	// SyncedUpdateType is the type of the transition that happened. It is either BlockToEntityTransition, when
	// a block placed becomes a falling entity, or EntityToBlockTransition, when a falling entity hits the
	// ground and becomes a solid block again.
	SyncedUpdateType uint32
}

// Marshal encodes/decodes a BlockChangeEntry.
func (x *BlockChangeEntry) Marshal(r protocol.IO) {
	if IsProtoGTE(r, ID944) {
		r.BlockPos(&x.BlockPos)
	} else {
		UBlockPos(r, &x.BlockPos)
	}
	r.Varuint32(&x.BlockRuntimeID)
	r.Varuint32(&x.Flags)
	r.Varuint64(&x.SyncedUpdateEntityUniqueID)
	r.Varuint32(&x.SyncedUpdateType)
}

// FromLatest ...
func (x *BlockChangeEntry) FromLatest(latest *protocol.BlockChangeEntry) {
	x.BlockPos = latest.BlockPos
	x.BlockRuntimeID = latest.BlockRuntimeID
	x.Flags = latest.Flags
	x.SyncedUpdateEntityUniqueID = latest.SyncedUpdateEntityUniqueID
	x.SyncedUpdateType = latest.SyncedUpdateType
}

// ToLatest ...
func (x *BlockChangeEntry) ToLatest() *protocol.BlockChangeEntry {
	return &protocol.BlockChangeEntry{
		BlockPos:                   x.BlockPos,
		BlockRuntimeID:             x.BlockRuntimeID,
		Flags:                      x.Flags,
		SyncedUpdateEntityUniqueID: x.SyncedUpdateEntityUniqueID,
		SyncedUpdateType:           x.SyncedUpdateType,
	}
}

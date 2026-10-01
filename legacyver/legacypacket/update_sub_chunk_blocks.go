package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

// UpdateSubChunkBlocks is essentially just UpdateBlock packet, however for a set of blocks in a sub-chunk.
type UpdateSubChunkBlocks struct {
	// Position is the block position of the sub-chunk being referred to.
	Position protocol.BlockPos
	// Blocks contains each updated block change entry.
	Blocks []proto.BlockChangeEntry
	// Extra contains each updated block change entry for the second layer, usually for waterlogged blocks.
	Extra []proto.BlockChangeEntry
}

// ID ...
func (*UpdateSubChunkBlocks) ID() uint32 {
	return packet.IDUpdateSubChunkBlocks
}

// Marshal ...
func (pk *UpdateSubChunkBlocks) Marshal(io protocol.IO) {
	if proto.IsProtoGTE(io, proto.ID944) {
		io.BlockPos(&pk.Position)
	} else {
		proto.UBlockPos(io, &pk.Position)
	}
	protocol.Slice(io, &pk.Blocks)
	protocol.Slice(io, &pk.Extra)
}

// ToLatest ...
func (pk *UpdateSubChunkBlocks) ToLatest() *packet.UpdateSubChunkBlocks {
	blocks := make([]protocol.BlockChangeEntry, len(pk.Blocks))
	for i, block := range pk.Blocks {
		blocks[i] = *block.ToLatest()
	}
	extra := make([]protocol.BlockChangeEntry, len(pk.Extra))
	for i, block := range pk.Extra {
		extra[i] = *block.ToLatest()
	}
	return &packet.UpdateSubChunkBlocks{
		Position: pk.Position,
		Blocks:   blocks,
		Extra:    extra,
	}
}

// FromLatest ...
func (pk *UpdateSubChunkBlocks) FromLatest(latest *packet.UpdateSubChunkBlocks) *UpdateSubChunkBlocks {
	pk.Position = latest.Position
	pk.Blocks = make([]proto.BlockChangeEntry, len(latest.Blocks))
	for i, block := range latest.Blocks {
		var entry proto.BlockChangeEntry
		entry.FromLatest(&block)
		pk.Blocks[i] = entry
	}
	pk.Extra = make([]proto.BlockChangeEntry, len(latest.Extra))
	for i, block := range latest.Extra {
		var entry proto.BlockChangeEntry
		entry.FromLatest(&block)
		pk.Extra[i] = entry
	}
	return pk
}
